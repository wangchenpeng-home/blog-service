package tracer

import (
	"io"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	// jaeger client的配置项
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{ //固定采样、对所有数据都进行采样等
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{ // 是否启用LoggingReporter、刷新缓冲区的频率、上报的Agent地址
			BufferFlushInterval: 1 * time.Second,
			LogSpans:            true,
			LocalAgentHostPort:  agentHostPort,
		},
	}

	// 根据配置项初始化Tacer对象，此处返回的是opentracing.Tracer，并不是某个供应商的追踪系统的对象
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	// 设置全局的Tracer对象，根据实际情况设置即可
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
