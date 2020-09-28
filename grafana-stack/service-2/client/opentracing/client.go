package opentracing

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

var tracer opentracing.Tracer
var closer io.Closer

func init() {
	defcfg := config.Configuration{
		ServiceName: "echo-tracer",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	cfg, err := defcfg.FromEnv()
	if err != nil {
		panic("Could not parse Jaeger env vars: " + err.Error())
	}

	tracer, closer, err = cfg.NewTracer()
	if err != nil {
		panic("Could not initialize jaeger tracer: " + err.Error())
	}

	opentracing.SetGlobalTracer(tracer)
}

func Close() error {
	return closer.Close()
}

func GetTracer() opentracing.Tracer {
	return tracer
}
