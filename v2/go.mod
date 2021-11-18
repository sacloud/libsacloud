module github.com/sacloud/libsacloud/v2

go 1.16

require (
	github.com/fatih/structs v1.1.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-playground/validator/v10 v10.2.0
	github.com/hashicorp/go-multierror v1.0.1-0.20190722213833-bdca7bb83f60
	github.com/huandu/xstrings v1.2.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.3.2
	github.com/sacloud/ftps v1.1.0
	github.com/sacloud/go-http v0.0.3
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace v0.15.1
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.15.1
	go.opentelemetry.io/otel v0.15.0
	go.opentelemetry.io/otel/exporters/stdout v0.15.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.15.0
	go.opentelemetry.io/otel/sdk v0.15.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
)
