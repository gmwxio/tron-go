package lsp

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/golangq/q"
	antlr "github.com/wxio/goantlr"
	"github.com/wxio/tron-go/adl"
	"github.com/wxio/tron-go/internal/adlwi"
	"github.com/wxio/tron-go/internal/ctree"
	"golang.org/x/tools/lsp/protocol"
)

type f_clientMsgLog func(ctx context.Context, ptype protocol.MessageType, msg ...string)

func (svr *server) client_msg_log(ctx context.Context, ptype protocol.MessageType, msg ...string) {
	smsg := msg[0]
	lmsg := msg[0]
	if len(msg) > 1 {
		lmsg = msg[1]
	}
	svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
		Message: smsg,
		Type:    ptype,
	})
	svr.client.LogMessage(ctx, &protocol.LogMessageParams{
		Message: lmsg,
		Type:    ptype,
	})
}
func (svr *server) client_log(ctx context.Context, ptype protocol.MessageType, msg ...string) {
	smsg := strings.Join(msg, " ")
	svr.client.LogMessage(ctx, &protocol.LogMessageParams{
		Message: smsg,
		Type:    ptype,
	})
}
func (svr *server) null_msg_log(ctx context.Context, ptype protocol.MessageType, msg ...string) {
}

func (svr *server) lastFileStable(ctx context.Context, clientMsgLog f_clientMsgLog) (bool, error) {
	fname := svr.lastFileUri
	if strings.HasPrefix(fname, "file://") {
		fname = fname[len("file://"):]
	}
	root := svr.initParams.RootURI
	if strings.HasPrefix(root, "file://") {
		root = root[len("file://"):]
	}
	fby, err := ioutil.ReadFile(fname)
	if err != nil {
		clientMsgLog(ctx, protocol.Warning, "LSP  internal error, file not found", err.Error())
		q.Q(err)
		return false, err
	}
	cur, err := svr.fileCache.get(svr.lastFileUri)
	if err != nil {
		clientMsgLog(ctx, protocol.Warning, "LSP  internal error, unknown file", err.Error())
		q.Q(err)
		return false, err
	}
	if string(fby) != cur {
		clientMsgLog(ctx, protocol.Warning, "Can't compile as file is not saved")
		return false, nil
	}
	return true, nil
}

func (svr *server) adl2ast(ctx context.Context, text string, clientMsgLog f_clientMsgLog) (map[string]adl.Module, error) {
	root := svr.initParams.RootURI
	if strings.HasPrefix(root, "file://") {
		root = root[len("file://"):]
	}
	name := svr.lastFileUri
	if strings.HasPrefix(name, "file://") {
		name = name[len("file://"):]
	}
	temp := filepath.Join(svr.tempDir, filepath.Base(name))
	err := ioutil.WriteFile(temp, []byte(text), os.ModePerm)
	if err != nil {
		clientMsgLog(ctx, protocol.Warning, "File write error", err.Error())
		return nil, err
	}
	args := []string{"ast", "--verbose"}
	combile_adl := ""
	{
		// combile adl
		combile_adl, err = filepath.Abs(filepath.Join(svr.tempDir, "ald.json"))
		if err != nil {
			clientMsgLog(ctx, protocol.Warning, "File path error. See TRON LSP log", err.Error())
			q.Q(err)
			return nil, err
		}
		args = append(args, "--combined-output="+combile_adl)
	}
	for _, inc := range svr.extConfig.Includes {
		abs, err := filepath.Abs(filepath.Join(root, inc))
		if err != nil {
			clientMsgLog(ctx, protocol.Warning, "file path error. See TRON LSP log", err.Error())
			q.Q(err)
			return nil, err
		}
		args = append(args, "-I"+abs)
	}
	args = append(args, temp)
	q.Q("adlc", args)
	cmd := exec.Command("adlc", args...)
	q.Q(svr.tempDir)
	cmd.Dir = svr.tempDir
	by, err := cmd.CombinedOutput()
	q.Q(string(by))
	if err != nil {
		clientMsgLog(ctx, protocol.Warning, "ADL error. See TRON LSP log", string(by))
		q.Q(err)
		return nil, err
	}

	astby, err := ioutil.ReadFile(combile_adl)
	if err != nil {
		clientMsgLog(ctx, protocol.Warning, "File error. See TRON LSP log", err.Error())
		q.Q(err)
		return nil, err
	}
	allmod := map[string]adl.Module{}
	err = json.Unmarshal(astby, &allmod)
	if err != nil {
		clientMsgLog(ctx, protocol.Warning, "JSON error. See TRON LSP log", err.Error())
		q.Q(err)
		return nil, err
	}
	return allmod, nil
}

func (svr *server) compile(ctx context.Context, text string, allmod map[string]adl.Module, clientMsgLog f_clientMsgLog) {
	if svr.lastFileUri != "" {
		q.Q("compile " + svr.lastFileUri)
		var err error
		tr, atr, bl, ts, err1 := adl.BuildAdlAST(text)
		_, _, _, _ = tr, atr, bl, ts
		// adl.QTreeToken(ts, bl)
		if err1.Error() != nil {
			clientMsgLog(ctx, protocol.Warning, "ADL AST error. See TRON LSP log", err.Error())
			q.Q(err)
			return
		}
		cv := &compileV{}
		adl.VisitADLWi(tr, cv)
		q.Q(cv.name)
		cma := struct {
			scope      adl.ScopedName
			outputFile string
			templateGo string
		}{}
		found := false
		mod, ex := allmod[cv.name]
		if !ex {
			clientMsgLog(ctx, protocol.Warning, "ADL processing error. See TRON LSP log", "Module name '"+cv.name+"'")
			return
		}
		for _, an := range mod.Annotations {
			if an.Key.ModuleName == "schemas.valuedriven.devmode" && an.Key.Name == "CompileModuleAnnotation" {
				sm := an.Val.(map[string]interface{})
				sma := sm["annotation"].(map[string]interface{})
				cma.scope.ModuleName = sma["moduleName"].(string)
				cma.scope.Name = sma["name"].(string)
				cma.outputFile = sm["outputFile"].(string)
				cma.templateGo = sm["templateGo"].(string)
				found = true
			}
		}
		if !found {
			clientMsgLog(ctx, protocol.Warning, "No module annotation schemas.valuedriven.devmode.CompileModuleAnnotation found")
			return
		}
		q.Q(cma)
		found = false
		var val interface{}
		for _, an := range mod.Annotations {
			if an.Key.ModuleName == cma.scope.ModuleName && an.Key.Name == cma.scope.Name {
				found = true
				val = an.Val
			}
		}
		if !found {
			clientMsgLog(ctx, protocol.Warning, "No module annotation "+cma.scope.ModuleName+"."+cma.scope.Name)
			return
		}
		te, err := template.New("").Parse(cma.templateGo)
		if err != nil {
			clientMsgLog(ctx, protocol.Warning, "Template error. See TRON LSP log", err.Error())
			q.Q(err)
			return
		}
		json, err := json.MarshalIndent(val, "", "  ")
		if err != nil {
			clientMsgLog(ctx, protocol.Warning, "JSON error. See TRON LSP log", err.Error())
			q.Q(err)
			return
		}
		data := struct {
			Json string
		}{
			Json: string(json),
		}
		name := svr.lastFileUri
		if strings.HasPrefix(name, "file://") {
			name = name[len("file://"):]
		}
		dname := filepath.Dir(name)
		cma.outputFile, err = filepath.Abs(filepath.Join(dname, cma.outputFile))
		if err != nil {
			clientMsgLog(ctx, protocol.Warning, "File error. See TRON LSP log", err.Error())
			q.Q(err)
			return
		}
		f, err := os.OpenFile(cma.outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			clientMsgLog(ctx, protocol.Warning, "File error. See TRON LSP log", err.Error())
			q.Q(err)
			return
		}
		err = te.Execute(f, data)
		if err != nil {
			clientMsgLog(ctx, protocol.Warning, "Template exec error. See TRON LSP log", err.Error())
			q.Q(err)
			return
		}
		clientMsgLog(ctx, protocol.Info, "Compiled to "+cma.outputFile)
	} else {
		clientMsgLog(ctx, protocol.Info, "No file selected")
	}
}

type compileV struct {
	*antlr.BaseParseTreeVisitor
	name string
}

func (v *compileV) VisitModule(ctx adlwi.IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	v.name = ctx.GetTok().(ctree.TreeNode).Val().(adl.Module).Name
	return
}
