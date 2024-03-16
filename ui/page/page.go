package page

import (
	"gioui.org/layout"
)

type PagePath string

const (
	PageHome     PagePath = "/"
	PageSettings PagePath = "/settings"

	PageTunnel     PagePath = "/tunnel"
	PageTunnelFile PagePath = "/tunnel/file"
	PageTunnelHTTP PagePath = "/tunnel/http"
	PageTunnelTCP  PagePath = "/tunnel/tcp"
	PageTunnelUDP  PagePath = "/tunnel/udp"

	PageEntrypoint    PagePath = "/entrypoint"
	PageEntrypointTCP PagePath = "/entrypoint/tcp"
	PageEntrypointUDP PagePath = "/entrypoint/udp"
)

type PageOptions struct {
	ID string
}

type PageOption func(opts *PageOptions)

func WithPageID(id string) PageOption {
	return func(opts *PageOptions) {
		opts.ID = id
	}
}

type Page interface {
	Init(opts ...PageOption)
	Layout(gtx layout.Context) layout.Dimensions
}
