// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"github.com/gofiber/fiber/v2"
	otg "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func (http *httpFouService) printString(ctx context.Context, request requestFouServicePrintString) (response responseFouServicePrintString, err error) {

	span := otg.SpanFromContext(ctx)
	span.SetTag("method", "printString")

	err = http.svc.PrintString(ctx, request.Id)
	if err != nil {
		if http.errorHandler != nil {
			err = http.errorHandler(err)
		}
		errData := toString(err)
		ext.Error.Set(span, true)
		span.SetTag("msg", err.Error())
		if errData != "{}" {
			span.SetTag("errData", errData)
		}
	}
	return
}
func (http *httpFouService) servePrintString(ctx *fiber.Ctx) (err error) {

	span := otg.SpanFromContext(ctx.UserContext())
	span.SetTag("method", "printString")

	var request requestFouServicePrintString
	ctx.Response().SetStatusCode(200)

	if _id := ctx.Params("id"); _id != "" {
		var id string
		id = _id
		request.Id = id
	}

	var response responseFouServicePrintString
	if response, err = http.printString(ctx.UserContext(), request); err == nil {
		return sendResponse(ctx, response)
	}
	if errCoder, ok := err.(withErrorCode); ok {
		ctx.Status(errCoder.Code())
	} else {
		ctx.Status(fiber.StatusInternalServerError)
	}
	return sendResponse(ctx, err)
}
func (http *httpFouService) get(ctx context.Context, request requestFouServiceGet) (response responseFouServiceGet, err error) {

	span := otg.SpanFromContext(ctx)
	span.SetTag("method", "get")

	response.Answ, err = http.svc.Get(ctx, request.Id)
	if err != nil {
		if http.errorHandler != nil {
			err = http.errorHandler(err)
		}
		errData := toString(err)
		ext.Error.Set(span, true)
		span.SetTag("msg", err.Error())
		if errData != "{}" {
			span.SetTag("errData", errData)
		}
	}
	return
}
func (http *httpFouService) serveGet(ctx *fiber.Ctx) (err error) {

	span := otg.SpanFromContext(ctx.UserContext())
	span.SetTag("method", "get")

	var request requestFouServiceGet
	ctx.Response().SetStatusCode(200)

	if _id := ctx.Params("id"); _id != "" {
		var id string
		id = _id
		request.Id = id
	}

	var response responseFouServiceGet
	if response, err = http.get(ctx.UserContext(), request); err == nil {
		return sendResponse(ctx, response)
	}
	if errCoder, ok := err.(withErrorCode); ok {
		ctx.Status(errCoder.Code())
	} else {
		ctx.Status(fiber.StatusInternalServerError)
	}
	return sendResponse(ctx, err)
}
