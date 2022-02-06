package server

import (
	_ "github.com/kuzminal/proglog"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
