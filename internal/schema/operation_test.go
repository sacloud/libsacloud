package schema

import (
	"net/http"
	"testing"

	"github.com/sacloud/libsacloud/v2/internal/schema/meta"
	"github.com/stretchr/testify/require"
)

func TestOperation(t *testing.T) {
	resources := Resources{}
	resource := resources.Define("Test")

	type expectOperationValues struct {
		methodName                 string
		requestEnvelopeStructName  string
		responseEnvelopeStructName string
		requestPayloadName         string
		responsePayloadName        string
	}

	expects := []struct {
		operation *Operation
		expect    *expectOperationValues
	}{
		{
			// TODO あとで直す
			operation: func() *Operation {
				o := resource.DefineOperation("Create").
					PathFormat(DefaultPathFormat).
					RequestEnvelope(&EnvelopePayloadDesc{PayloadType: meta.Static(struct{}{})}).
					Argument(ArgumentZone).
					Argument(&Argument{
						Name:       "arg1",
						MapConvTag: "Destination",
						Type: &Model{
							Name: "Model",
							Fields: []*FieldDesc{
								{
									Name: "Field1",
									Type: meta.Static(""),
								},
								{
									Name: "Field2",
									Type: meta.Static(""),
								},
							},
						},
					}).
					ResultFromEnvelope(&Model{
						Name: "ResultFromEnvelope",
						Fields: []*FieldDesc{
							{
								Name: "Field3",
								Type: meta.Static(""),
							},
							{
								Name: "Field4",
								Type: meta.Static(""),
							},
						},
					}, &EnvelopePayloadDesc{
						PayloadName: "Test",
						PayloadType: meta.Static(struct{}{}),
					}, "Test")
				o.Method = http.MethodPost
				return o
			}(),
			expect: &expectOperationValues{
				methodName:                 "Create",
				requestEnvelopeStructName:  "testCreateRequestEnvelope",
				responseEnvelopeStructName: "testCreateResponseEnvelope",
				requestPayloadName:         "Test",
				responsePayloadName:        "Test",
			},
		},
	}

	for _, tc := range expects {
		resource.AddOperation(tc.operation)
		require.Equal(t, tc.expect.methodName, tc.operation.MethodName())
		require.Equal(t, tc.expect.requestEnvelopeStructName, tc.operation.RequestEnvelopeStructName())
		require.Equal(t, tc.expect.responseEnvelopeStructName, tc.operation.ResponseEnvelopeStructName())
		require.Equal(t, tc.expect.requestPayloadName, tc.operation.RequestPayloads()[0].PayloadName)
		require.Equal(t, tc.expect.responsePayloadName, tc.operation.ResponsePayloads()[0].PayloadName)
	}
}
