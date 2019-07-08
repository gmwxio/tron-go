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

func (svr *server) compile(ctx context.Context) {
	if svr.lastFileUri != "" {
		q.Q("compile " + svr.lastFileUri)
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
			svr.client.LogMessage(ctx, &protocol.LogMessageParams{
				Message: err.Error(),
				Type:    protocol.Warning,
			})
			q.Q(err)
			return
		}
		cur, err := svr.fileCache.get(svr.lastFileUri)
		if err != nil {
			svr.client.LogMessage(ctx, &protocol.LogMessageParams{
				Message: err.Error(),
				Type:    protocol.Warning,
			})
			q.Q(err)
			return
		}
		if string(fby) != cur {
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "Can't compile as file is not saved",
				Type:    protocol.Warning,
			})
			return
		}
		args := []string{"ast", "--verbose"}
		for _, inc := range svr.extConfig.Includes {
			args = append(args, "-I"+root+"/"+inc)
		}
		args = append(args, fname)
		q.Q("adlc", args)
		cmd := exec.Command("adlc", args...)
		q.Q(svr.tempDir)
		cmd.Dir = svr.tempDir
		by, err := cmd.CombinedOutput()
		q.Q(string(by))
		if err != nil {
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "ADL error. See TRON LSP log",
				Type:    protocol.Warning,
			})
			svr.client.LogMessage(ctx, &protocol.LogMessageParams{
				Message: string(by),
				Type:    protocol.Warning,
			})
			q.Q(err)
			return
		}
		tr, atr, bl, ts, err1 := adl.BuildAdlAST(string(fby))
		_, _, _, _ = tr, atr, bl, ts
		// adl.QTreeToken(ts, bl)
		if err1.Error() != nil {
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "ADL AST error. See TRON LSP log",
				Type:    protocol.Warning,
			})
			svr.client.LogMessage(ctx, &protocol.LogMessageParams{
				Message: err.Error(),
				Type:    protocol.Warning,
			})
			q.Q(err)
			return
		}
		cv := &compileV{}
		adl.VisitADL(tr, cv)
		q.Q(cv.name)
		astby, err := ioutil.ReadFile(svr.tempDir + "/" + cv.name + ".json")
		if err != nil {
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "File error. See TRON LSP log",
				Type:    protocol.Warning,
			})
			svr.client.LogMessage(ctx, &protocol.LogMessageParams{
				Message: err.Error(),
				Type:    protocol.Warning,
			})
			q.Q(err)
			return
		}
		mod := adl.Module{}
		err = json.Unmarshal(astby, &mod)
		if err != nil {
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "JSON error. See TRON LSP log",
				Type:    protocol.Warning,
			})
			svr.client.LogMessage(ctx, &protocol.LogMessageParams{
				Message: err.Error(),
				Type:    protocol.Warning,
			})
			q.Q(err)
			return
		}
		cma := struct {
			scope      adl.ScopedName
			outputFile string
			templateGo string
		}{}
		found := false
		for _, an := range mod.Annotations {
			if an.Key.ModuleName == "valuedriven.devmode" && an.Key.Name == "CompileModuleAnnotation" {
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
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "No module annotation valuedriven.devmode.CompileModuleAnnotation found",
				Type:    protocol.Warning,
			})
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
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "No module annotation " + cma.scope.ModuleName + "." + cma.scope.Name,
				Type:    protocol.Warning,
			})
			return
		}
		te, err := template.New("").Parse(cma.templateGo)
		if err != nil {
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "Template error. See TRON LSP log",
				Type:    protocol.Warning,
			})
			svr.client.LogMessage(ctx, &protocol.LogMessageParams{
				Message: err.Error(),
				Type:    protocol.Warning,
			})
			q.Q(err)
			return
		}
		json, err := json.MarshalIndent(val, "", "  ")
		if err != nil {
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "JSON error. See TRON LSP log",
				Type:    protocol.Warning,
			})
			svr.client.LogMessage(ctx, &protocol.LogMessageParams{
				Message: err.Error(),
				Type:    protocol.Warning,
			})
			q.Q(err)
			return
		}
		data := struct {
			Json string
		}{
			Json: string(json),
		}
		dname := filepath.Dir(fname)
		cma.outputFile, err = filepath.Abs(filepath.Join(dname, cma.outputFile))
		if err != nil {
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "File error. See TRON LSP log",
				Type:    protocol.Warning,
			})
			svr.client.LogMessage(ctx, &protocol.LogMessageParams{
				Message: err.Error(),
				Type:    protocol.Warning,
			})
			q.Q(err)
			return
		}
		f, err := os.OpenFile(cma.outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "File error. See TRON LSP log",
				Type:    protocol.Warning,
			})
			svr.client.LogMessage(ctx, &protocol.LogMessageParams{
				Message: err.Error(),
				Type:    protocol.Warning,
			})
			q.Q(err)
			return
		}
		err = te.Execute(f, data)
		if err != nil {
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "Template exec error. See TRON LSP log",
				Type:    protocol.Warning,
			})
			svr.client.LogMessage(ctx, &protocol.LogMessageParams{
				Message: err.Error(),
				Type:    protocol.Warning,
			})
			q.Q(err)
			return
		}
		svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
			Message: "Compiled to " + cma.outputFile,
			Type:    protocol.Info,
		})
	} else {
		svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
			Message: "No file selected",
			Type:    protocol.Info,
		})
	}
}

type compileV struct {
	*antlr.BaseParseTreeVisitor
	name string
}

func (v *compileV) VisitModule(ctx adlwi.IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	v.name = ctx.GetTok().(*ctree.TreeNode).Val.(adl.Module).Name
	return
}
