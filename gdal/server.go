package gdal

import (
	"fmt"
	"middleware"
	"net"
	"net/http"
	"time"
)

func Server(host string, port string, writeTimeout time.Duration, readTimeout time.Duration) (*http.Server, string) {
	addr := net.JoinHostPort(host, port)
	fmt.Printf("Configuring GDAL server at %s\n", addr)

	mux := http.NewServeMux()
	SetupRoutes(mux)

	return &http.Server{
		Addr:         addr,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
		Handler:      middleware.Middleware{Handler: mux},
	}, addr
}
