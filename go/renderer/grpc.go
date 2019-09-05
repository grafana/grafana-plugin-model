package renderer

import (
	"context"
)

type GRPCClient struct {
	RendererClient
}

func (m *GRPCClient) Render(ctx context.Context, req *RenderRequest) (*RenderResponse, error) {
	return m.RendererClient.Render(ctx, req)
}

type GRPCServer struct {
	RendererPlugin
}

func (m *GRPCServer) Render(ctx context.Context, req *RenderRequest) (*RenderResponse, error) {
	return m.RendererPlugin.Render(ctx, req)
}
