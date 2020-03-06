// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/sacloud/libsacloud/v2/grpc/proto"
	"github.com/sacloud/libsacloud/v2/grpc/sacloud"
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
	sacloud.RegisterAuthStatusAPIServer(s, &server.AuthStatusService{})

	log.Printf("Listening: %s", listen.Addr().String())
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
