package middlewares

import (
	"fmt"
	"net/http"
)

//api is hosted at www.myapi.com
//frontend server us at www.myfrontend.com

// Allowd origins
var allowedOrigins = []string{
	"https://my-origin-url.com",
	"https://www.myfrontend.com",
	"https://localhost:3000",
}

func Cors(next http.Handler) http.Handler {
	fmt.Println("Cors Middlewares")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("cors Middlewares bieng returned...")
		origin := r.Header.Get("Origin")
		fmt.Println(origin)

		if isOriginAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			http.Error(w, "Not allowed by CORS", http.StatusForbidden)
			return
		}

		//w.Header().Set()
		w.Header().Set("Access-Control-Allow-Header", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "3600")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
		fmt.Println("cors Middlewares bieng returned...")

	})
}

func isOriginAllowed(origin string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}
	return false
}
