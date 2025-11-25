package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/internal/repositories/sqlconnect"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		return
	}

	_, err = sqlconnect.ConnectDB()
	if err != nil {
		fmt.Println("Error----:", err)
		return
	}

	port := os.Getenv("API_PORT")

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	//rl := mw.NewRateLimiter(5, time.Minute)

	//hppOptions := mw.HPPOptions{
	//CheckQuery:               true,
	//CheckBody:                true,
	//CheckBodyOnlyContentType: "application/x-www-form-urlencode",
	//Whitelist:                []string{"sortBy", "sortOrder", "name", "age", "class"},
	//}

	//secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	//secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	router := router.Router()
	secureMux := mw.SecurityHeaders(router)

	//create custom server
	server := &http.Server{
		Addr: port,
		//Handler: mux,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
		//Handler: middlewares.Cors(mux),
	}

	fmt.Println("Server is running on port", port)
	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
