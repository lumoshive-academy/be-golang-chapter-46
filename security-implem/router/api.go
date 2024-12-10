package router

import (
	"be-golang-chapter-46-implem/infra"
	"be-golang-chapter-46-implem/middleware"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SetupReouter(ctx infra.Context, middleware middleware.Middleware) {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1/api")
	v1.Use(middleware.Logger())
	v1.Use(middleware.IPWhitelistMiddleware())

	auth := v1.Group("/auth")
	{
		auth.POST("/add-ip", ctx.Handler.Auth.AddIp)
		auth.POST("/register", ctx.Handler.Auth.Register)
		auth.POST("/login", middleware.Limiter(), ctx.Handler.Auth.Login)
	}

	// shipping routes
	customer := v1.Group("/voucher")
	{
		customer.Use(ctx.JWT.AuthJWT())
		customer.POST("/", ctx.Handler.VoucherHandler.Create)
		customer.GET("/", ctx.Handler.VoucherHandler.GetAll)
	}

	srv := &http.Server{
		Addr:    ":" + ctx.Config.Port,
		Handler: router.Handler(),
	}

	go func() {
		// service connections
		ctx.Log.Info("Starting server...", zap.String("address", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ctx.Log.Fatal("Server listen error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx.Log.Info("Shutdown signal received")

	ctxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctxt); err != nil {
		ctx.Log.Error("Server shutdown error", zap.Error(err))
	} else {
		ctx.Log.Info("Server shut down gracefully")
	}
}
