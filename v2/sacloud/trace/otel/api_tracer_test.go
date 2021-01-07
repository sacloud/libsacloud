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

	"github.com/sacloud/libsacloud/v2"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/fake"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/stdout"
	export "go.opentelemetry.io/otel/sdk/export/trace"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric/controller/push"
	"go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
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

	var spanData []*SpanData
	if err := json.Unmarshal(outputBuf.Bytes(), &spanData); err != nil {
		t.Fatal(err)
	}
	require.Len(t, spanData, 1)
	sp := spanData[0]

	require.EqualValues(t, instrumentation.Library{
		Name:    "github.com/sacloud/libsacloud",
		Version: libsacloud.Version,
	}, sp.InstrumentationLibrary)

	require.EqualValues(t, codes.Ok, sp.StatusCode)
	require.EqualValues(t, "", sp.StatusMessage)
	require.Len(t, sp.Attributes, 3) // arguments(zone, condition) + results = 3
}

type SpanData struct {
	SpanContext  map[string]interface{}
	ParentSpanID interface{}
	export.SpanData
}

func setupTracer() func() {
	outputBuf.Reset()

	exporter, err := stdout.NewExporter(
		stdout.WithPrettyPrint(),
		stdout.WithWriter(outputBuf),
		stdout.WithoutMetricExport(),
	)
	if err != nil {
		log.Fatalf("failed to initialize stdout export pipeline: %v", err)
	}

	bsp := sdktrace.NewSimpleSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp), sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}))
	defer func() { _ = tp.Shutdown(context.Background()) }()
	pusher := push.New(
		basic.New(
			simple.NewWithExactDistribution(),
			exporter,
		),
		exporter,
	)
	pusher.Start()
	otel.SetTracerProvider(tp)
	return pusher.Stop
}
