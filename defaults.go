package micro

import (
	"github.com/blastbao/go-micro/client"
	"github.com/blastbao/go-micro/debug/trace"
	"github.com/blastbao/go-micro/server"
	"github.com/blastbao/go-micro/store"

	// set defaults
	gcli "github.com/blastbao/go-micro/client/grpc"
	memTrace "github.com/blastbao/go-micro/debug/trace/memory"
	gsrv "github.com/blastbao/go-micro/server/grpc"
	memStore "github.com/blastbao/go-micro/store/memory"
)

func init() {
	// default client
	client.DefaultClient = gcli.NewClient()
	// default server
	server.DefaultServer = gsrv.NewServer()
	// default store
	store.DefaultStore = memStore.NewStore()
	// set default trace
	trace.DefaultTracer = memTrace.NewTracer()
}
