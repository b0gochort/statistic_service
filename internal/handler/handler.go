package handler

import (
	"fmt"
	"github.com/b0gochort/statistic_service/internal/service"
	"github.com/valyala/fasthttp"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}

}

func (h *Handler) InitRoutes(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	switch string(ctx.Path()) {
	case "/new-chat":
		ping(ctx)
	}
}

func ping(ctx *fasthttp.RequestCtx) {
	fmt.Println("pong")
}
