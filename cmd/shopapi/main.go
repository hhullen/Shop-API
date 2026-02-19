package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"shopapi/internal/api/v1"
	"shopapi/internal/clients/postgres"
	"shopapi/internal/logger"
	"shopapi/internal/service"
)

// @title           Shop API
// @version         1.0
// @description     Cервер на Golang с OpenAPI документацией.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Maksim
// @contact.url    https://github.com/hhullen
// @contact.email  hhullen@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	apiLog := logger.NewLogger(os.Stdout, "API")
	defer apiLog.Stop()
	serviceLog := logger.NewLogger(os.Stdout, "SERVICE")
	defer serviceLog.Stop()

	conn, err := postgres.NewSQLConn(ctx)
	if err != nil {
		log.Fatal(err)
	}

	db := postgres.NewClient(ctx, conn)

	s := service.NewService(ctx, serviceLog, db, db, db, db)
	api := api.NewAPI(ctx, apiLog, s, s, s, s)

	err = api.Start()
	if err != nil {
		apiLog.InfoKV("service stopped with error", "error", api.Start())
	}
}
