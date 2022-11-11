package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
	"tw/core"
	"tw/server"
)

func main() {

	server := new(server.JSONServer)

	fmt.Println("==> calling wallet core from go")

	mn := "confirm bleak useless tail chalk destroy horn step bulb genuine attract split"

	fmt.Println("==> mnemonic is valid: ", core.IsMnemonicValid(mn))

	server.MN = mn

	port := ":8080"

	s := rpc.NewServer()                                 // Create a new RPC server
	s.RegisterCodec(json.NewCodec(), "application/json") // Register the type of data requested as JSON
	s.RegisterService(server, "")                        // Register the service by creating a new JSON server

	r := mux.NewRouter()
	r.Handle("/api/v1/sign_transaction", s)
	log.Fatal(http.ListenAndServe(port, r))

}
