package app

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"video-chat/internal/config"
	"video-chat/internal/logger"
	"video-chat/internal/middleware"
	"video-chat/internal/routes"

	"github.com/gin-gonic/gin"
)

type App struct {
	Server  *http.Server
	Logger  *logger.ZapLogger
	Config  *config.Config
	Router  *gin.Engine
	Context context.Context
	Cancel  context.CancelFunc
}

func Bootstrap() *App {
	app := &App{
		Router: gin.New(),
		Logger: logger.NewLogger(),
		Config: config.Load(),
	}

	app.Logger.Infof("App: INIT")

	app.Router.Use(gin.Recovery(), middleware.ErrorHandler(), middleware.LoggerHandler(app.Logger))
	routes.SetupRoutes(app.Router)

	app.Logger.Infof("Router: INIT")

	app.Context, app.Cancel = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	app.Logger.Infof("Context, Cancel: INIT")

	return app
}

func (app *App) Run() {
	app.Server = &http.Server{
		Addr:    ":" + app.Config.Port,
		Handler: app.Router,
	}

	go startServer(app)

	<-app.Context.Done()
	app.Logger.Infof("Shutdown signal recoived")
	app.gracefulShutdown()
}

func startServer(app *App) {
	app.Logger.Infof("HTTP server starting on port: %s", app.Config.Port)
	if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		app.Logger.Fatalf("HTTP server error: %v", err)
	}
}

func (app *App) gracefulShutdown() {
	app.Cancel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Server.Shutdown(ctx); err != nil {
		app.Logger.Errorf("Error down http server: &v", err)
	}

	app.Logger.Infof("Server close")
}
