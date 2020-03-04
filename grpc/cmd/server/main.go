package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/sacloud/libsacloud/v2/grpc/proto"
	"github.com/sacloud/libsacloud/v2/grpc/server"
	"google.golang.org/grpc"
)

const addr = "/tmp/libsacloud.sock"

func main() {

	listen, err := net.Listen("unix", addr)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		if err := listen.Close(); err != nil {
			log.Println(err)
		}
		if err := os.Remove(addr); err != nil {
			log.Println(err)
		}
	}()

	s := grpc.NewServer()
	proto.RegisterServerOpServer(s, &server.ServerOpService{})

	log.Printf("Listening: %s", listen.Addr().String())
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
