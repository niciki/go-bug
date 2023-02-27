// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"service/internal/interfaces"
	"time"
)

type metricsFouService struct {
	next            interfaces.FouService
	requestCount    metrics.Counter
	requestCountAll metrics.Counter
	requestLatency  metrics.Histogram
}

func metricsMiddlewareFouService(next interfaces.FouService) interfaces.FouService {
	return &metricsFouService{
		next:            next,
		requestCount:    RequestCount.With("service", "FouService"),
		requestCountAll: RequestCountAll.With("service", "FouService"),
		requestLatency:  RequestLatency.With("service", "FouService"),
	}
}

func (m metricsFouService) PrintString(ctx context.Context, id string) (err error) {

	defer func(begin time.Time) {
		m.requestLatency.With("method", "printString", "success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
	}(time.Now())

	defer m.requestCount.With("method", "printString", "success", fmt.Sprint(err == nil)).Add(1)

	m.requestCountAll.With("method", "printString").Add(1)

	return m.next.PrintString(ctx, id)
}

func (m metricsFouService) Get(ctx context.Context, id string) (answ string, err error) {

	defer func(begin time.Time) {
		m.requestLatency.With("method", "get", "success", fmt.Sprint(err == nil)).Observe(time.Since(begin).Seconds())
	}(time.Now())

	defer m.requestCount.With("method", "get", "success", fmt.Sprint(err == nil)).Add(1)

	m.requestCountAll.With("method", "get").Add(1)

	return m.next.Get(ctx, id)
}