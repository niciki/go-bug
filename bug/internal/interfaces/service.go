package interfaces

import (
	"context"
)

// @tg http-prefix=v1
// @tg http-server log trace metrics
type FouService interface {
	// @tg http-path=/printstring/:id
	// @tg http-method=POST
	// @tg http-success=200
	PrintString(ctx context.Context, id string) (err error)
	// @tg http-path=/get/:id
	// @tg http-method=GET
	// @tg http-success=200
	Get(ctx context.Context, id string) (answ string, err error)
}
