package proxy

import (
	"context"

	models "github.com/grafana/grafana_plugin_model/go/models"
	plugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type DatasourcePlugin interface {
	Query(ctx context.Context, req *models.DatasourceRequest) (*models.DatasourceResponse, error)
}

type DatasourcePluginImpl struct {
	plugin.NetRPCUnsupportedPlugin
	Plugin DatasourcePlugin
}

func (p *DatasourcePluginImpl) GRPCServer(s *grpc.Server) error {
	models.RegisterDatasourcePluginServer(s, &GRPCServer{p.Plugin})
	return nil
}

func (p *DatasourcePluginImpl) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{models.NewDatasourcePluginClient(c)}, nil
}

type GRPCClient struct {
	models.DatasourcePluginClient
}

func (m *GRPCClient) Query(ctx context.Context, req *models.DatasourceRequest) (*models.DatasourceResponse, error) {
	return m.DatasourcePluginClient.Query(ctx, req)
}

type GRPCServer struct {
	DatasourcePlugin
}

func (m *GRPCServer) Query(ctx context.Context, req *models.DatasourceRequest) (*models.DatasourceResponse, error) {
	return m.DatasourcePlugin.Query(ctx, req)
}
