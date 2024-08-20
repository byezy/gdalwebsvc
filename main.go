package main

import (
	"admin"
	"context"
	"fmt"
	"gdal"
	//"logging"
	"middleware"
	"net"
	"net/http"
	"ogr"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const OGRPORT = "8080"
const OGRHOST = "0.0.0.0"

const GDLPORT = "8081"
const GDLHOST = "0.0.0.0"

const ADMPORT = "8082"
const ADMHOST = "0.0.0.0"

var TEN = 10 * time.Second

func main() {

	// create the servers
	ogrServer := Server("OGR", OGRHOST, OGRPORT, TEN, TEN, ogr.SetupRoutes)
	gdlServer := Server("GDAL", GDLHOST, GDLPORT, TEN, TEN, gdal.SetupRoutes)
	admServer := Server("ADMIN", ADMHOST, ADMPORT, TEN, TEN, admin.SetupRoutes)

	// create a context with a Done channel and a cancel function
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a signal channel and listen for signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	fmt.Println("Starting services")

	go func() { startService(ogrServer, "OGR") }()
	go func() { startService(gdlServer, "GDAL") }()
	go func() { startService(admServer, "ADMIN") }()

	// shut down services when an interrupt signal is received

	defer func(sigs chan os.Signal) {
		sig := <-sigs
		fmt.Printf("\nServices shutting down...(%s)\n", sig)
		serviceShutdown(ogrServer, ctx, "OGR")
		serviceShutdown(gdlServer, ctx, "GDAL")
		serviceShutdown(admServer, ctx, "ADMIN")
		fmt.Println("\nAll services are now stopped")
	}(sigs)

}

func startService(srv *http.Server, tag string) {
	fmt.Printf("Starting %s server at %s\n", tag, srv.Addr)
	_ = srv.ListenAndServe()
}

func serviceShutdown(srv *http.Server, ctx context.Context, tag string) {
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Error when shutting down the server:\n%s", err)
	} else {
		fmt.Printf("%s server shut down... ", tag)
	}
}

func Server(name string, host string, port string, writeTimeout time.Duration, readTimeout time.Duration, ServerRoutesSetup func(mux *http.ServeMux)) *http.Server {
	addr := net.JoinHostPort(host, port)
	fmt.Printf("Configuring %s server at %s... ", name, addr)

	mux := http.NewServeMux()

	fmt.Printf("Setting up %s routes... ", name)
	ServerRoutesSetup(mux)
	fmt.Printf("%s ok\n", name)

	return &http.Server{
		Addr:         addr,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		Handler:      middleware.Middleware{Handler: mux},
	}
}
