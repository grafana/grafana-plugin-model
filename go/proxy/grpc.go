package proxy

import (
	"context"

	models "github.com/grafana/plugin_model/go/models"
)

type GRPCClient struct {
	models.DatasourcePluginClient
}

func (m *GRPCClient) Query(ctx context.Context, req *models.TsdbQuery) (*models.Response, error) {
	return m.DatasourcePluginClient.Query(ctx, req)
}

type GRPCServer struct {
	DatasourcePlugin
}

func (m *GRPCServer) Query(ctx context.Context, req *models.TsdbQuery) (*models.Response, error) {
	return m.DatasourcePlugin.Query(ctx, req)
}
