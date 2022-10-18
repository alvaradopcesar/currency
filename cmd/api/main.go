package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/currency/internal/logger"
	"github.com/currency/pkg/currency/handler"
	"github.com/currency/pkg/currency/repository"
	"github.com/currency/pkg/currency/service"
	"github.com/currency/pkg/free_currency_api"
	"github.com/currency/pkg/load_free_currency"
	"github.com/currency/pkg/transport"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/currency/internal/config"
	"github.com/currency/internal/db/postgres"
	"github.com/currency/internal/request"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const migrationsRootFolder = "file://migrations"

func doMigrate(databaseURL string) error {
	m, err := migrate.New(
		migrationsRootFolder,
		databaseURL,
	)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func main() {

	log := logger.New("currency", false)

	conf := config.New()
	log.Info("Configuration loaded successfully version:", config.GetVersion())

	err := doMigrate(conf.DbPostgresUrl)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	postgresClient := postgres.NewPostgresClient(conf.DbPostgresUrl)

	cRepository := repository.NewRepository(postgresClient.DB, time.Duration(conf.DbTimeOUT))
	cService := service.NewService(cRepository)
	cHandler := handler.NewHandler(cService)
	httpTransportRouter := transport.NewHTTPRouter(cHandler)

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", conf.Port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      httpTransportRouter,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Info("Transport Start http port: ", conf.Port)
		if err = srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
		log.Info("Transport Stop http port: ", conf.Port)
	}()

	req := request.NewHTTP(conf.FreeCurrencyApi, &http.Client{
		Timeout: time.Duration(conf.TimeoutSeconds) * time.Second,
	}, nil)

	fRepository := free_currency_api.NewRepository(req, free_currency_api.Config{
		FreeCurrencyApiKey: conf.FreeCurrencyApiKey,
	})
	serFree := free_currency_api.NewService(fRepository)
	load_free_currency.NewProvider(conf.FreeCurrencyTime, serFree, cService, log).Load()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("Service gracefully shut down")
	os.Exit(0)
}
