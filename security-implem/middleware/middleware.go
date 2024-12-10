package middleware

import (
	"be-golang-chapter-46-implem/database"

	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type Middleware struct {
	Cacher       database.Cacher
	Log          *zap.Logger
	LoginLimiter *LoginLimiter
}

func NewMiddleware(cacher database.Cacher, log *zap.Logger, privateKey string, r rate.Limit, b int) Middleware {
	return Middleware{
		Cacher: cacher,
		Log:    log,
		LoginLimiter: &LoginLimiter{
			limiters: make(map[string]*rate.Limiter),
			rate:     r,
			burst:    b,
		},
	}
}
