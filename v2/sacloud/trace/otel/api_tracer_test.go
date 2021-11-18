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

package otel

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/sacloud/libsacloud/v2"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/fake"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var outputBuf = bytes.NewBufferString("")

func TestTracer(t *testing.T) {
	// init tracer
	cleanup := setupTracer()
	defer cleanup()

	// set factory func
	Initialize()
	// enable fake mode
	fake.SwitchFactoryFuncToFake()

	caller := &sacloud.Client{}
	archiveOp := sacloud.NewArchiveOp(caller)
	archiveOp.Find(context.Background(), "is1a", nil) // nolint
	require.NotEmpty(t, t, outputBuf.String())

	var sp SpanData
	if err := json.Unmarshal(outputBuf.Bytes(), &sp); err != nil {
		t.Fatal(err)
	}
	require.EqualValues(t, instrumentation.Library{
		Name:    "github.com/sacloud/libsacloud",
		Version: libsacloud.Version,
	}, sp.InstrumentationLibrary)

	require.EqualValues(t, codes.Ok, sp.Status.Code)
	require.EqualValues(t, "", sp.Status.Description)
	require.Len(t, sp.Attributes, 3) // arguments(zone, condition) + results = 3
}

// SpanData from: https://github.com/open-telemetry/opentelemetry-go/blob/ece1879fae1bcea6d32f94fd8f6d174c3e118a6b/sdk/trace/tracetest/span.go#L56-L74
type SpanData struct {
	Name                   string
	SpanContext            trace.SpanContext
	Parent                 trace.SpanContext
	SpanKind               trace.SpanKind
	StartTime              time.Time
	EndTime                time.Time
	Attributes             []attribute.KeyValue
	Events                 []sdktrace.Event
	Links                  []sdktrace.Link
	Status                 sdktrace.Status
	DroppedAttributes      int
	DroppedEvents          int
	DroppedLinks           int
	ChildSpanCount         int
	Resources              []*resource.Resource
	InstrumentationLibrary instrumentation.Library
}

func setupTracer() func() {
	outputBuf.Reset()

	exporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithWriter(outputBuf),
	)
	if err != nil {
		log.Fatalf("failed to initialize stdout export pipeline: %v", err)
	}

	bsp := sdktrace.NewSimpleSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp), sdktrace.WithSampler(sdktrace.AlwaysSample()))
	otel.SetTracerProvider(tp)

	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("stopping tracer provider: %v", err)
		}
	}
}
