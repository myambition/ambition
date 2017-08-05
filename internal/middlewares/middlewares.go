package middlewares

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/myambition/ambition/internal/logger"
)

func LogError(label string, in endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		resp, err = in(ctx, req)
		defer func(inerr error) {
			if inerr != nil {
				logger.LogError(err, "endpoint", label)
			}
		}(err)
		return
	}
}
