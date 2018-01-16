package proxy

import (
	"context"

	models "github.com/grafana/plugin_model/go/models"
	plugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type DatasourcePlugin interface {
	Query(ctx context.Context, req *models.TsdbQuery) (*models.Response, error)
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
