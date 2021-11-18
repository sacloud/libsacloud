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

package otel_test

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/sacloud/libsacloud/v2/helper/api"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/ostype"
	traceotel "github.com/sacloud/libsacloud/v2/sacloud/trace/otel"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// ref: https://github.com/open-telemetry/opentelemetry-go/blob/a65d50a4c6412ebaf4b8a5ffeed8eae2a9417ef1/example/jaeger/main.go

// Example ローカルのJaegerを利用する例
func Example() {
	tp, err := tracerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	tr := tp.Tracer("component-main")
	ctx, span := tr.Start(ctx, "foo")
	defer span.End()

	// サンプルAPIリクエスト
	op(ctx)

	// Jaeger UI( http://localhost:16686/search など)を開くとトレースが確認できるはず
}

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("libsacloud"),
			attribute.String("exporter", "jaeger"),
			attribute.Int64("ID", 1),
		)),
	)
	return tp, nil
}

func op(ctx context.Context) {
	// set factory func
	traceotel.Initialize()

	caller := api.NewCaller(&api.CallerOptions{
		AccessToken:       os.Getenv("SAKURACLOUD_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET"),
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
