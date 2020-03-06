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

	serverOp := &client.ServerOp{Addr: "unix:////tmp/libsacloud.sock"}
	if err := serverOp.Boot(context.Background(), zone, types.StringID(id)); err != nil {
		log.Fatal(err)
	}
	log.Println("done: server boot")
}
