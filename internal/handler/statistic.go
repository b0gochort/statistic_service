package handler

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"resenje.org/logging"
	"strconv"
)

func (h *Handler) SetOnline(ctx *fasthttp.RequestCtx) {
	if !ctx.IsPost() {
		logging.Info("invalid req method")
		ctx.Error("handler SetOnline check method: %v", fasthttp.StatusMethodNotAllowed)
		return
	}

	q := ctx.QueryArgs()

	userId, err := strconv.Atoi(string(q.Peek("user-id")))
	if err != nil {
		logging.Info("handler.strconv.Atoi: %v", err)
		ctx.Error("unprocessable entity", fasthttp.StatusUnprocessableEntity)
		return
	}

	if err := h.services.SetOnline(int64(userId), true); err != nil {
		logging.Info("handler.services.SetOnline: %v", err)
		ctx.Error("set online", fasthttp.StatusInternalServerError)
		return
	}
	ctx.Write([]byte("successfull"))
}

func (h *Handler) GetOnline(ctx *fasthttp.RequestCtx) {
	if !ctx.IsGet() {
		logging.Info("invalid req method")
		ctx.Error("handler GetOnline check method: %v", fasthttp.StatusMethodNotAllowed)
		return
	}

	online, err := h.services.GetOnline()
	if err != nil {
		logging.Info("handler.services.SetOnline: %v", err)
		ctx.Error("set online", fasthttp.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(online)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.GetMessages.Marshal: %v", err))
		ctx.Error("something went wrong ", fasthttp.StatusInternalServerError)

		return
	}

	ctx.Write(res)
	return
}

func (h *Handler) GetStatHourByHour(ctx *fasthttp.RequestCtx) {
	if !ctx.IsGet() {
		logging.Info("invalid req method")
		ctx.Error("handler GetStatHourByHour check method: %v", fasthttp.StatusMethodNotAllowed)
		return
	}

	stat, err := h.services.GetStatisticByHour()
	if err != nil {
		logging.Info("handler.GetStatHourByHour.services.GetStatisticByHour: %v", err)
		ctx.Error("get stat", fasthttp.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(stat)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.GetStatHourByHour.Marshal: %v", err))
		ctx.Error("something went wrong ", fasthttp.StatusInternalServerError)

		return
	}

	ctx.Write(res)
	return
}

func (h *Handler) GetStatCategory(ctx *fasthttp.RequestCtx) {
	if !ctx.IsGet() {
		logging.Info("invalid req method")
		ctx.Error("handler GetStatHourByHour check method: %v", fasthttp.StatusMethodNotAllowed)
		return
	}

	stat, err := h.services.GetCategoryRatio()
	if err != nil {
		logging.Info("handler.GetStatHourByHour.services.GetCategoryRatio: %v", err)
		ctx.Error("get stat", fasthttp.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(stat)
	if err != nil {
		logging.Info(fmt.Sprintf("handler.GetStatHourByHour.Marshal: %v", err))
		ctx.Error("something went wrong ", fasthttp.StatusInternalServerError)

		return
	}

	ctx.Write(res)
	return
}
