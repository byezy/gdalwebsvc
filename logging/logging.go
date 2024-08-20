package logging

import (
	"net/http"
)

//import (
//
//)
//
//logFile, err := os.OpenFile("log.txt", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
//if err != nil {
//panic(err)
//}
//mw := io.MultiWriter(os.Stdout, logFile)
//log.SetOutput(mw)

func WithLogging(h http.Handler) http.Handler {
	logFn := func(rw http.ResponseWriter, r *http.Request) {
		//start := time.Now()

		//uri := r.RequestURI
		//method := r.Method
		//h.ServeHTTP(rw, r) // serve the original request
		//
		//duration := time.Since(start)

		//// log request details
		//log.WithFields(log.Fields{
		//	"uri":      uri,
		//	"method":   method,
		//	"duration": duration,
		//})
	}
	return http.HandlerFunc(logFn)
}
