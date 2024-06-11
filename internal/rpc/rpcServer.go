package rpc

import (
	"log"
	"net/http"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

func StartRPCServer() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(RPCService), "")

	http.Handle("/rpc", s)
	log.Fatal(http.ListenAndServe(":8081", nil))
}