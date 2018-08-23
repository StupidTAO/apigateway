package main

import (
	"fmt"
	"net/http"
	"os"

	cmn "github.com/DSiSc/apigateway/common"
	"github.com/DSiSc/apigateway/log"
	rpcserver "github.com/DSiSc/apigateway/rpc/lib/server"
	amino "github.com/tendermint/go-amino"
)

var routes = map[string]*rpcserver.RPCFunc{
	"hello_world": rpcserver.NewRPCFunc(HelloWorld, "name,num"),
}

func HelloWorld(name string, num int) (Result, error) {
	return Result{fmt.Sprintf("hi %s %d", name, num)}, nil
}

type Result struct {
	Result string
}

func main() {
	mux := http.NewServeMux()
	cdc := amino.NewCodec()
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	rpcserver.RegisterRPCFuncs(mux, routes, cdc, logger)
	_, err := rpcserver.StartHTTPServer("0.0.0.0:8008", mux, logger, rpcserver.Config{})
	if err != nil {
		cmn.Exit(err.Error())
	}

	// Wait forever
	cmn.TrapSignal(func() {
	})

}
