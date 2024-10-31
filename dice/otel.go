package main

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
)

// setupOTelSDK 引导 OpenTelemetry pipeline。
// 如果没有返回错误，请确保调用 shutdown 进行适当清理。
func setupOTelSDK(ctx context.Context) (shutdown func(context.Context) error, err error) {
	// 	var shutdownFuncs []func(context.Context) error
	//
	// shutdown 会调用通过 shutdownFuncs 注册的清理函数。
	// 调用产生的错误会被合并。
	// 每个注册的清理函数将被调用一次。
	// 	shutdown = func(ctx context.Context) error {
	// 		var err error
	// 		for _, fn := range shutdownFuncs {
	// 			err = errors.Join(err, fn(ctx))
	// 		}
	// 		shutdownFuncs = nil
	// 		return err
	// 	}
	//
	// 	// handleErr 调用 shutdown 进行清理，并确保返回所有错误信息。
	// 	handleErr := func(inErr error) {
	// 		err = errors.Join(inErr, shutdown(ctx))
	// 	}

	// 设置传播器
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// 设置 trace provider.
	tracerProvider, err := newTraceProvider()
	if err != nil {
		// handleErr(err)
		return
	}
	// shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	shutdown = tracerProvider.Shutdown
	otel.SetTracerProvider(tracerProvider)

	// 设置 meter provider.
	// meterProvider, err := newMeterProvider()
	// if err != nil {
	// 	handleErr(err)
	// 	return
	// }
	// shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	// otel.SetMeterProvider(meterProvider)

	return
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTraceProvider() (*trace.TracerProvider, error) {
	// 创建一个往终端输出的 traceExporter
	traceExporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			// 默认为 5s。为便于演示，设置为 1s。
			trace.WithBatchTimeout(time.Second)),
	)
	return traceProvider, nil
}

// func newMeterProvider() (*metric.MeterProvider, error) {
// 	metricExporter, err := stdoutmetric.New()
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	meterProvider := metric.NewMeterProvider(
// 		metric.WithReader(metric.NewPeriodicReader(metricExporter,
// 			// 默认为 1m。为便于演示，设置为 3s。
// 			metric.WithInterval(3*time.Second))),
// 	)
// 	return meterProvider, nil
// }
