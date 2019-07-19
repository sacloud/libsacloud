package dsl

import (
	"net/http"
	"testing"

	"github.com/sacloud/libsacloud/v2/internal/dsl/meta"
	"github.com/stretchr/testify/require"
)

func TestOperation(t *testing.T) {
	resources := Resources{}
	resource := &Resource{
		Name: "Test",
	}
	resources.Define(resource)

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
			operation: &Operation{
				ResourceName: resource.Name,
				Name:         "Create",
				PathFormat:   DefaultPathFormat,
				Method:       http.MethodPost,
				RequestEnvelope: RequestEnvelope(&EnvelopePayloadDesc{
					Type: meta.Static(struct{}{}),
					Name: "Test",
				}),
				Arguments: []*Argument{
					{
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
					},
				},
				ResponseEnvelope: ResponseEnvelope(&EnvelopePayloadDesc{
					Name: "Test",
					Type: meta.Static(struct{}{}),
				}),
				Results: Results{
					{
						SourceField: "Test",
						DestField:   "Test",
						IsPlural:    false,
						Model: &Model{
							Name: "ResponseEnvelope",
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
						},
					},
				},
			},
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
		resource.Operations = []*Operation{tc.operation}
		require.Equal(t, tc.expect.methodName, tc.operation.MethodName())
		require.Equal(t, tc.expect.requestEnvelopeStructName, tc.operation.RequestEnvelopeStructName())
		require.Equal(t, tc.expect.responseEnvelopeStructName, tc.operation.ResponseEnvelopeStructName())
		require.Equal(t, tc.expect.requestPayloadName, tc.operation.RequestPayloads()[0].Name)
		require.Equal(t, tc.expect.responsePayloadName, tc.operation.ResponsePayloads()[0].Name)
	}
}
