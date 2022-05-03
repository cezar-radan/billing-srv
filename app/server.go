package app

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/ledgertech/billing-srv/config"
	"github.com/ledgertech/billing-srv/util/clog"
	"github.com/ledgertech/billing-srv/util/metering"
)

type ctxKeyType string

const CtxAppConfKey = ctxKeyType("AppConfig")

func Start() {
	clog.Infof("starting service - %s \n", config.App.ServiceName)

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(notImplementedHandler)
	api.MethodNotAllowedHandler = http.HandlerFunc(notFoundHandler)
	api.HandleFunc("/healthCheck",
		multipleMiddleware(healthCheckHandler, metering.MeteringMiddleware)).Methods("GET")
	api.HandleFunc("/generateInvoice",
		multipleMiddleware(generateInvoiceHandler, metering.MeteringMiddleware)).Methods("POST")

	//Metering area
	metering.MeteringConfig.ServiceName = config.App.ServiceName
	metering.MeteringConfig.Enabled = config.App.Metrics.Enabled
	if config.App.Metrics.Enabled {
		r.Handle(config.App.Metrics.Endpoint, metering.PrometheusHandler())
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, CtxAppConfKey, config.App)
	handler := cors.AllowAll().Handler(r)

	srv := &http.Server{
		Addr:    config.App.HostPort,
		Handler: handler,
		BaseContext: func(net.Listener) context.Context {
			return ctx
		},
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		clog.Infof("will listen on: %s", config.App.HostPort)
		if config.App.TLS.Cert == "" || config.App.TLS.Key == "" {
			clog.Warn("tls not configured. running insecure on http")
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				clog.Errorf("on listen: %s\n", err)
				os.Exit(1)
			}
		} else {
			if err := srv.ListenAndServeTLS(config.App.TLS.Cert, config.App.TLS.Key); err != nil && err != http.ErrServerClosed {
				clog.Errorf("on listen: %s\n", err)
				os.Exit(1)
			}
		}
	}()
	clog.Info("[*] To exit press CTRL+C")

	<-done
	clog.Infof("stopping service - %s \n", config.App.ServiceName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		clog.Errorf("stopping/shutdown service - %s - failed:%+v", config.App.ServiceName, err)
		os.Exit(1)
	}
	clog.Infof("service - %s - stopped\n", config.App.ServiceName)

}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func multipleMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}

	return wrapped
}
