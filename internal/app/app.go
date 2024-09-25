package app

import (
	"context"
	"fmt"
	"log/slog"

	"invia/api"
	"invia/api/restapi"
	"invia/internal/app/handler"
	"invia/internal/app/health"
	"invia/internal/app/services"
	"invia/internal/app/web"
	"invia/internal/lib/middleware"

	"github.com/gorilla/mux"
)

type App struct {
	HTTPServer *web.HTTPServer
	service    services.IService

	health.LivenessChecker
	health.ReadinessChecker
}

// NewService creates a new instance of the service.
func NewService(
	log *slog.Logger,
	connectString string,
	port int,
	useTracing bool,
	tracingAddress string,
	serviceName string,
) (*App, error) {
	const op = "app.NewService"
	ctx := context.Background()

	app := &App{}
	srv, err := services.NewService(log, connectString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	telemetryMiddleware, err := addTelemetryMiddleware(ctx, useTracing, tracingAddress, serviceName)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	restApiServer := api.NewRestApiServer()

	router := mux.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(telemetryMiddleware)

	router.HandleFunc("/live", health.LivenessHandler(app)).Methods("GET")
	router.HandleFunc("/ready", health.ReadinessHandler(app)).Methods("GET")
	//
	// router.HandleFunc("/api/file/{id}", handler.GetFileItem(srv)).Methods("GET")
	// router.HandleFunc("/api/file", handler.PutFileItem(srv)).Methods("PUT")

	h := restapi.HandlerFromMux(restApiServer, router)

	server, err := web.New(log, port, h)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	app.HTTPServer = server
	app.service = srv

	return app, nil
}

// Start starts the application.
func (a *App) Start() {
	a.HTTPServer.Start()
}

// Stop stops the application.
func (a *App) Stop() {
	if a != nil && a.service != nil {
		err := a.service.Close()
		if err != nil {
			fmt.Println("An error occurred closing service" + err.Error())

			return
		}
	}
}

func addTelemetryMiddleware(
	ctx context.Context, useTracing bool, tracingAddress string, serviceName string,
) (mux.MiddlewareFunc, error) {
	var telemetryMiddleware mux.MiddlewareFunc
	var err error
	if useTracing {
		telemetryMiddleware, err = handler.AddTelemetryMiddleware(ctx, tracingAddress, serviceName)
		if err != nil {
			return nil, err
		}
	}

	return telemetryMiddleware, nil
}

func (a *App) LivenessCheck() bool {
	return a.service.LivenessCheck()
}

func (a *App) ReadinessCheck() bool {
	return a.service.ReadinessCheck()
}
