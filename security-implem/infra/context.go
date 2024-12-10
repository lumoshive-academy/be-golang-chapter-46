package infra

import (
	"be-golang-chapter-46-implem/database"
	"be-golang-chapter-46-implem/handler"
	"be-golang-chapter-46-implem/infra/jwt"
	"be-golang-chapter-46-implem/middleware"
	"be-golang-chapter-46-implem/repository"
	"be-golang-chapter-46-implem/service"
	"be-golang-chapter-46-implem/util"

	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type Context struct {
	Log        *zap.Logger
	Config     util.Configuration
	Handler    handler.AllHandler
	Cacher     database.Cacher
	Middleware middleware.Middleware
	JWT        jwt.JWT
}

func NewContext() (Context, error) {
	logger, err := util.LoggerInit()
	if err != nil {
		return Context{}, err
	}

	config, err := util.ReadConfig()
	if err != nil {
		return Context{
			Log: logger,
		}, err
	}

	db, err := database.InitDB(config)
	if err != nil {
		return Context{
			Log: logger,
		}, err
	}

	rdb := database.NewCacher(config, 60*60)

	jwt := jwt.NewJWT(config.PrivateKey, config.PublicKey, logger)
	middleware := middleware.NewMiddleware(rdb, logger, config.PrivateKey, rate.Limit(config.Limiter.RateLimit), config.Limiter.Burst)

	repo := repository.NewAllRepository(db, logger)
	service := service.NewAllService(repo, logger)
	handler := handler.NewAllHandler(service, logger, rdb, jwt)
	return Context{Log: logger, Config: config, Handler: handler, Cacher: rdb, Middleware: middleware, JWT: jwt}, nil
}
