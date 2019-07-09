package lsp

import (
	"context"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/golangq/q"
	"golang.org/x/tools/lsp/protocol"
)

func (svr *server) serveJSON(rw http.ResponseWriter, req *http.Request) {
	q.Q(req)
	encUri := req.URL.RequestURI()[1:]
	q.Q(encUri)
	furi, err := base64.URLEncoding.DecodeString(encUri)
	if err != nil {
		q.Q(err)
		http.Error(rw, "hu", http.StatusBadRequest)
		return
	}
	text, err := svr.fileCache.get(string(furi))
	if err != nil {
		q.Q(err)
		http.Error(rw, "huu", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, text)
}
func (svr *server) serverBrowser() {
	lst, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		q.Q(err)
		ctx := context.Background()
		err := svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
			Message: "Error starting web server inside LSP",
			Type:    protocol.Warning,
		})
		if err != nil {
			q.Q(err)
		}
		return
	}
	svr.webAddr = lst.Addr()
	q.Q(svr.webAddr)
	mux := &http.ServeMux{}
	mux.HandleFunc("/", svr.serveJSON)
	wsvr := http.Server{
		Handler: mux,
	}
	go func() {
		err = wsvr.Serve(lst)
		if err != nil {
			q.Q(err)
			ctx := context.Background()
			err := svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "Error starting web server inside LSP",
				Type:    protocol.Warning,
			})
			if err != nil {
				q.Q(err)
			}
		}
	}()
}

func (svr *server) openbrowser() (err error) {
	if svr.webAddr == nil {
		return
	}
	b64 := base64.URLEncoding.EncodeToString([]byte(svr.lastFileUri))
	url := "http://" + svr.webAddr.String() + "/" + b64
	ok := browserOpen(url)
	if !ok {
		err = fmt.Errorf("unsupported platform")
	}
	return
}

// Commands returns a list of possible commands to use to open a url.
func browserCommands() [][]string {
	var cmds [][]string
	if exe := os.Getenv("BROWSER"); exe != "" {
		cmds = append(cmds, []string{exe})
	}
	switch runtime.GOOS {
	case "darwin":
		cmds = append(cmds, []string{"/usr/bin/open"})
	case "windows":
		cmds = append(cmds, []string{"cmd", "/c", "start"})
	default:
		if os.Getenv("DISPLAY") != "" {
			// xdg-open is only for use in a desktop environment.
			cmds = append(cmds, []string{"xdg-open"})
		} else {
			q.Q("env DISPLAY not set can't launch on Linux")
		}
	}
	cmds = append(cmds,
		[]string{"chrome"},
		[]string{"google-chrome"},
		[]string{"chromium"},
		[]string{"firefox"},
	)
	return cmds
}

// Open tries to open url in a browser and reports whether it succeeded.
func browserOpen(url string) bool {
	for _, args := range browserCommands() {
		cmd := exec.Command(args[0], append(args[1:], url)...)
		if cmd.Start() == nil && appearsSuccessful(cmd, 3*time.Second) {
			return true
		}
	}
	return false
}

// appearsSuccessful reports whether the command appears to have run successfully.
// If the command runs longer than the timeout, it's deemed successful.
// If the command runs within the timeout, it's deemed successful if it exited cleanly.
func appearsSuccessful(cmd *exec.Cmd, timeout time.Duration) bool {
	errc := make(chan error, 1)
	go func() {
		errc <- cmd.Wait()
	}()

	select {
	case <-time.After(timeout):
		return true
	case err := <-errc:
		return err == nil
	}
}
