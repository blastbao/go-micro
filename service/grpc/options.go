package grpc

import (
	"crypto/tls"

	gc "github.com/blastbao/go-micro/client/grpc"
	gs "github.com/blastbao/go-micro/server/grpc"
	"github.com/blastbao/go-micro/service"
)

// WithTLS sets the TLS config for the service
func WithTLS(t *tls.Config) service.Option {
	return func(o *service.Options) {
		o.Client.Init(
			gc.AuthTLS(t),
		)
		o.Server.Init(
			gs.AuthTLS(t),
		)
	}
}
