package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sacloud/libsacloud/v2/grpc/client"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("[USAGE] go run main.go ZONE ID")
		os.Exit(1)
	}
	zone := os.Args[1]
	id := os.Args[2]

	serverOp := &client.ServerOp{Addr: "passthrough:///unix:///tmp/libsacloud.sock"}
	if err := serverOp.Boot(context.Background(), zone, types.StringID(id)); err != nil {
		log.Fatal(err)
	}
	log.Println("done: server boot")
}
