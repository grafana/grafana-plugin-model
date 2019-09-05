package renderer

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type RendererPlugin interface {
	Render(ctx context.Context, req *RenderRequest) (*RenderResponse, error)
}

type RendererPluginImpl struct {
	plugin.NetRPCUnsupportedPlugin
	Plugin RendererPlugin
}

func (p *RendererPluginImpl) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterRendererServer(s, &GRPCServer{p.Plugin})
	return nil
}

func (p *RendererPluginImpl) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{NewRendererClient(c)}, nil
}
