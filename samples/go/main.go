package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"tw/core"
	serv "tw/server"
)

func main() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	address, exists := os.LookupEnv("ADDRESS")

	if !exists {
		address = ""
	}

	port, exists := os.LookupEnv("PORT")

	if !exists {
		port = "8080"
	}

	server := new(serv.JSONServer)

	fmt.Println("==> calling wallet core from go")

	mn, exists := os.LookupEnv("MNEMONIC")
	if !exists {
		log.Fatal("can't find mnemonic")
	}

	fmt.Println("==> mnemonic is valid: ", core.IsMnemonicValid(mn))

	server.MN = mn

	s := rpc.NewServer()                                 // Create a new RPC server
	s.RegisterCodec(json.NewCodec(), "application/json") // Register the type of data requested as JSON
	err := s.RegisterService(server, "")                 // Register the service by creating a new JSON server
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.Handle("/api/v1/sign_transaction", s)
	log.Fatal(http.ListenAndServe(address+":"+port, r))

}
