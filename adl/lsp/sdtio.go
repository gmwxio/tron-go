package lsp

import (
	"context"
	"os"

	"github.com/golangq/q"
	"github.com/jpillora/opts"
	"golang.org/x/tools/jsonrpc2"
	"golang.org/x/tools/lsp/protocol"
)

func NewConsole(version string) opts.Opts {
	return opts.New(&tcpsdtio{version: version}).Name("stdio")
}

type tcpsdtio struct {
	version string
}

func (svr *tcpsdtio) Run() error {
	q.Q("setting up stream")
	stream := jsonrpc2.NewHeaderStream(os.Stdin, os.Stdout)
	ctx, cancel := context.WithCancel(context.Background())
	srv := &server{
		// tcpConn: conn,
		cancel:  cancel,
		version: svr.version,
	}
	connLSP, client, _ := protocol.NewServer(stream, srv)
	srv.client = client
	srv.conn = connLSP
	q.Q(connLSP.Run(ctx))
	q.Q("exiting")
	return nil
}
