package middlewares

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/myambition/ambition/internal/logger"
	"golang.org/x/net/context"
)

func LogError(label string, in endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		resp, err = in(ctx, req)
		defer func(inerr error) {
			if inerr != nil {
				logger.Error().Log("endpoint", label, "err", err)
			}
		}(err)
		return
	}
}
