package middleware

import (
	"context"
	"fmt"
	"net/http"
)

type Middleware struct {
	Handler http.Handler
}

type Middlewares struct {
	Handler     http.Handler
	Middlewares []Middleware
}

//type Middleware func(http.HandlerFunc) http.HandlerFunc

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	//m.Handlers = append(m.Handlers, AuthHandler)
	//ctx := context.WithValue(r.Context(), "user", "unknown")
	//ctx = context.WithValue(ctx, "__requestStartTime__", time.Now())
	//r = r.WithContext(ctx)
	//r = r.WithContext(r.Context())

	//if len(m.Handlers) < 1 {
	//	fmt.Printf("No middleware handlers set up\n")
	//	m.Handler.ServeHTTP(w, r)
	//	return
	//}

	//m.Handler.ServeHTTP(w, r)

	//start := r.Context().Value("__requestStartTime__").(time.Time)
	//fmt.Printf("Request duration was %s\n", time.Now().Sub(start))

	//wrapped := h

	// loop in reverse to preserve middleware order
	//for i := len(m.Handlers) - 1; i >= 0; i-- {
	//	wrapped = m.Handlers[i](wrapped)
	//}
	//
	//return wrapped

}

type AuthHandler HandlerFunc

//{
//	//Handler  http.Handler
//	//Handlers []http.Handler
//}

// Now we define a type to implement ServeHTTP:
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

//func (f http.Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	f(w, r) // the receiver's a func; call it
//}

// func (h AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
func (ah AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AuthHandler")
	//ctx := context.WithValue(r.Context(), "user", "unknown")
	//ctx = context.WithValue(ctx, "__requestStartTime__", time.Now())
	//r = r.WithContext(ctx)
	//r = r.WithContext(r.Context())

	//if len(h.Handlers) < 1 {
	//	fmt.Printf("No middleware handlers set up\n")
	//	h.Handler.ServeHTTP(w, r)
	//	return
	//}

	//m.Handler.ServeHTTP(w, r)

	//start := r.Context().Value("__requestStartTime__").(time.Time)
	//fmt.Printf("Request duration was %s\n", time.Now().Sub(start))

	//wrapped := h

	// loop in reverse to preserve middleware order
	//for i := len(m.Handlers) - 1; i >= 0; i-- {
	//	wrapped = m.Handlers[i](wrapped)
	//}
	//
	//return wrapped
}

//func runMiddleware(h http.HandlerFunc, m []Middleware) http.HandlerFunc {
//	if len(m) < 1 {
//		return h
//	}
//
//	wrapped := h
//
//	// loop in reverse to preserve middleware order
//	for i := len(m) - 1; i >= 0; i-- {
//		wrapped = m[i](wrapped)
//	}
//
//	return wrapped
//}

type middleware func(http.HandlerFunc) http.HandlerFunc

// buildChain builds the middleware chain recursively, functions are first class
func buildChain(f http.HandlerFunc, m ...middleware) http.HandlerFunc {
	// if our chain is done, use the original handlerfunc
	if len(m) == 0 {
		return f
	}
	// otherwise nest the handlerfuncs
	return m[0](buildChain(f, m[1:cap(m)]...))
}

// AuthMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
var AuthMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	// one time scope setup area for middleware
	return func(w http.ResponseWriter, r *http.Request) {
		// ... pre handler functionality
		fmt.Println("start auth")
		f(w, r)
		fmt.Println("end auth")
		// ... post handler functionality
	}
}

// PrivateMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
var PrivateMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	// one time scope setup area for middleware
	return func(w http.ResponseWriter, r *http.Request) {
		// ... pre handler functionality
		fmt.Println("start private")
		f(w, r)
		fmt.Println("end private")
		// ... post handler functionality
	}
}

// PublicMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
var PublicMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	// one time scope setup area for middleware
	return func(w http.ResponseWriter, r *http.Request) {
		// add a request id..
		//backdrop.Set(r, "id", uuid.NewV4())
		// ... pre handler functionality
		fmt.Println("start public")
		f(w, r)
		fmt.Println("end public")
		// ... post handler functionality
	}
}

// this is the handler func we are wrapping with middlewares
func f(w http.ResponseWriter, r *http.Request) {
	// get the id from the context
	//id, err := backdrop.Get(r, "id")
	//if err != nil {
	//	fmt.Println("err: ", err.Error())
	//}
	id := 0
	fmt.Printf("request id is: %v\n", id)
	// you can also get the entire context if you are more comfortable with that
	ctx := r.Context() // backdrop.GetContext(r)
	ctx = context.WithValue(ctx, "key", "value")
	// and setting the newly created context in backdrop
	//backdrop.SetContext(r, ctx)
}
