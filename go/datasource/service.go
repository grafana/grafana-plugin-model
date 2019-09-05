package datasource

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// GrafanaAPI is the Grafana API interface that allows a datasource plugin to callback and request additional information from Grafana.
type GrafanaAPI interface {
	QueryDatasource(ctx context.Context, req *QueryDatasourceRequest) (*QueryDatasourceResponse, error)
}

// DatasourcePlugin is the Grafana datasource interface.
type DatasourcePlugin interface {
	Query(ctx context.Context, req *DatasourceRequest, api GrafanaAPI) (*DatasourceResponse, error)
}

// DatasourcePluginImpl is the implementation of plugin.Plugin so that it can be served/consumed as a Grafana datasource plugin.
// In addition, it also implements plugin.GRPCPlugin so that the it can be served over gRPC.
type DatasourcePluginImpl struct {
	plugin.NetRPCUnsupportedPlugin
	Impl DatasourcePlugin
}

func (p *DatasourcePluginImpl) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterDatasourcePluginServer(s, &GRPCServer{
		broker: broker,
		Impl:   p.Impl,
	})
	return nil
}

func (p *DatasourcePluginImpl) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		broker: broker,
		client: NewDatasourcePluginClient(c),
	}, nil
}

var _ plugin.GRPCPlugin = &DatasourcePluginImpl{}

// GRPCClient is an implementation of DatasourcePluginClient that talks over RPC.
type GRPCClient struct {
	broker *plugin.GRPCBroker
	client DatasourcePluginClient
}

func (m *GRPCClient) Query(ctx context.Context, req *DatasourceRequest, api GrafanaAPI) (*DatasourceResponse, error) {
	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		RegisterGrafanaAPIServer(s, api)

		return s
	}

	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)

	res, err := m.client.Query(ctx, req)

	s.Stop()
	return res, err
}

// GRPCServer is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	broker *plugin.GRPCBroker
	Impl   DatasourcePlugin
}

func (m *GRPCServer) Query(ctx context.Context, req *DatasourceRequest) (*DatasourceResponse, error) {
	conn, err := m.broker.Dial(req.RequestId)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	api := &GRPCGrafanaAPIClient{NewGrafanaAPIClient(conn)}
	return m.Impl.Query(ctx, req, api)
}

// GRPCGrafanaAPIClient is an implementation of GrafanaAPIClient that talks over RPC.
type GRPCGrafanaAPIClient struct{ client GrafanaAPIClient }

func (m *GRPCGrafanaAPIClient) QueryDatasource(ctx context.Context, req *QueryDatasourceRequest) (*QueryDatasourceResponse, error) {
	resp, err := m.client.QueryDatasource(ctx, req)
	if err != nil {
		hclog.Default().Info("grafana.QueryDatasource", "client", "start", "err", err)
		return nil, err
	}
	return resp, err
}
