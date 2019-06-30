package lsp

import (
	"context"
	"io"
	"log"
	"net"
	"os"

	"github.com/jpillora/opts"

	"github.com/golangq/q"
	"golang.org/x/tools/jsonrpc2"
	"golang.org/x/tools/lsp/protocol"
)

func NewTcp() opts.Opts {
	return opts.New(&tcpopt{}).Name("tcp").
		AddCommand(opts.New(&tcpsvr{Addr: "localhost:8080"}).Name("svr")).
		AddCommand(opts.New(&tcpclient{Addr: "localhost:8080"}).Name("client"))
}

type tcpopt struct {
}
type tcpsvr struct {
	Addr string
}
type tcpclient struct {
	Addr string
}

func (svr *tcpsvr) Run() error {
	lst, err := net.Listen("tcp", svr.Addr)
	if err != nil {
		q.Q("failed to listen: %v", err)
		log.Fatalf("failed to listen: %v", err)
	}
	defer lst.Close()
	//
	for {
		conn, err := lst.Accept()
		if err != nil {
			q.Q(err)
			os.Exit(1)
		}
		go handle(conn, conn)
	}
}

func handle(in io.Reader, out io.Writer) {
	q.Q("setting up stream")
	stream := jsonrpc2.NewHeaderStream(in, out)
	ctx, cancel := context.WithCancel(context.Background())
	srv := &server{
		// tcpConn: conn,
		cancel: cancel,
	}
	connLSP, client, _ := protocol.NewServer(stream, srv)
	srv.client = client
	srv.conn = connLSP
	q.Q(connLSP.Run(ctx))
	q.Q("exiting")
	// conn.Close()
}

func (tcpc *tcpclient) Run() error {
	con, err := net.Dial("tcp", tcpc.Addr)
	if err != nil {
		return err
	}
	errChan := make(chan error, 2)
	go func() {
		_, err = io.Copy(con, os.Stdin)
		errChan <- err
	}()
	go func() {
		_, err = io.Copy(os.Stdout, con)
		errChan <- err
	}()
	err = <-errChan
	return err
}
