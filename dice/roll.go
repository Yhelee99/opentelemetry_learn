package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

// 创建一个tracer对象
var tracer = otel.Tracer("roll")

func roll(w http.ResponseWriter, r *http.Request) {

	// 创建一个子span
	_, span := tracer.Start(r.Context(), "roll")
	defer span.End()

	// 业务逻辑
	num := 1 + rand.Intn(6)
	_, _ = fmt.Fprintf(w, "reslut: %v", num)

	// 往span里记录逻辑
	rollValueAttr := attribute.Int("roll.value", num)
	span.SetAttributes(rollValueAttr)
}
