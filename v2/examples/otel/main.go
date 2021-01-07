// Copyright 2016-2021 The Libsacloud Authors
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
	"log"
	"net/http"
	"os"

	"github.com/sacloud/libsacloud/v2/helper/api"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	"github.com/sacloud/libsacloud/v2/sacloud/trace/otel"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/label"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// ref: https://github.com/open-telemetry/opentelemetry-go/tree/v0.15.0/example/jaeger

// Example ローカルのJaegerを利用する例
func main() {
	ctx := context.Background()

	flush := initTracer()
	defer flush()

	// サンプルAPIリクエスト
	op(ctx)

	// Jaeger UI( http://localhost:16686/search など)を開くとトレースが確認できるはず
}

func initTracer() func() {
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "libsacloud",
			Tags: []label.KeyValue{
				label.String("exporter", "jaeger"),
				label.Float64("float", 312.23),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return flush
}

func op(ctx context.Context) {
	// set factory func
	otel.Initialize()

	caller := api.NewCaller(&api.CallerOptions{
		AccessToken:       os.Getenv("SAKURACLOUD_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET"),
		HTTPClient:        &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)},
	})
	archiveOp := sacloud.NewArchiveOp(caller)

	// normal operation
	archiveOp.Find(ctx, "is1a", &sacloud.FindCondition{ // nolint
		Count:  1,
		From:   0,
		Filter: ostype.ArchiveCriteria[ostype.Ubuntu],
	})

	// invalid operation(not foundエラーになるはず)
	archiveOp.Read(ctx, "is1a", types.ID(1)) // nolint
}
