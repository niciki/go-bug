// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"service/internal/interfaces"
)

type serverFouService struct {
	svc         interfaces.FouService
	printString FouServicePrintString
	get         FouServiceGet
}

type MiddlewareSetFouService interface {
	Wrap(m MiddlewareFouService)
	WrapPrintString(m MiddlewareFouServicePrintString)
	WrapGet(m MiddlewareFouServiceGet)

	WithTrace()
	WithMetrics()
	WithLog()
}

func newServerFouService(svc interfaces.FouService) *serverFouService {
	return &serverFouService{
		get:         svc.Get,
		printString: svc.PrintString,
		svc:         svc,
	}
}

func (srv *serverFouService) Wrap(m MiddlewareFouService) {
	srv.svc = m(srv.svc)
	srv.printString = srv.svc.PrintString
	srv.get = srv.svc.Get
}

func (srv *serverFouService) PrintString(ctx context.Context, id string) (err error) {
	return srv.printString(ctx, id)
}

func (srv *serverFouService) Get(ctx context.Context, id string) (answ string, err error) {
	return srv.get(ctx, id)
}

func (srv *serverFouService) WrapPrintString(m MiddlewareFouServicePrintString) {
	srv.printString = m(srv.printString)
}

func (srv *serverFouService) WrapGet(m MiddlewareFouServiceGet) {
	srv.get = m(srv.get)
}

func (srv *serverFouService) WithTrace() {
	srv.Wrap(traceMiddlewareFouService)
}

func (srv *serverFouService) WithMetrics() {
	srv.Wrap(metricsMiddlewareFouService)
}

func (srv *serverFouService) WithLog() {
	srv.Wrap(loggerMiddlewareFouService())
}