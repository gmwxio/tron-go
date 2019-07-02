package lsp

import (
	"context"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/jpillora/opts"

	"github.com/golangq/q"
	"golang.org/x/tools/jsonrpc2"
	"golang.org/x/tools/lsp/protocol"
)

func NewTcp(ver string) opts.Opts {
	return opts.New(&tcpopt{}).Name("tcp").
		AddCommand(opts.New(&tcpsvr{Addr: "localhost:8080", version: ver}).Name("svr")).
		AddCommand(opts.New(&tcpclient{Addr: "localhost:8080", version: ver}).Name("client"))
}

type tcpopt struct {
}
type tcpsvr struct {
	Addr    string
	version string
}
type tcpclient struct {
	Addr      string
	Reconnect bool
	version   string
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
		go svr.handle(conn, conn)
	}
}

func (svr *tcpsvr) handle(in io.Reader, out io.Writer) {
	q.Q("setting up stream")
	stream := jsonrpc2.NewHeaderStream(in, out)
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
	// conn.Close()
}

func (tcpc *tcpclient) Run() error {
	if tcpc.Reconnect {
		for {
			connect(tcpc.Addr)
			<-time.After(1 * time.Second)
		}
	}
	err := connect(tcpc.Addr)
	return err
}

func connect(addr string) error {
	con, err := net.Dial("tcp", addr)
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
