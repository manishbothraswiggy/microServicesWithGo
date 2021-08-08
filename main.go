package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"

	"ImportFiles/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// log := log.New(os.Stdout, "product-api", log.LstdFlags)

	// helloHandler := handlers.NewHello(log)
	// goodbyeHandler := handlers.NewGoodbye(log)
	// productsHandler := handlers.NewProducts(log)

	// serverMux := http.NewServeMux()

	// serverMux.Handle("/", helloHandler)
	// serverMux.Handle("/goodbye", goodbyeHandler)
	// serverMux.Handle("/products/", productsHandler)

	// userDefinedServer := &http.Server{
	// 	Addr:         ":9090",
	// 	Handler:      serverMux,
	// 	IdleTimeout:  120 * time.Second,
	// 	ReadTimeout:  time.Second,
	// 	WriteTimeout: time.Second,
	// }

	// go func() {
	// 	err := userDefinedServer.ListenAndServe()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// }()

	// signalChannel := make(chan os.Signal)
	// signal.Notify(signalChannel, os.Interrupt)
	// signal.Notify(signalChannel, os.Kill)

	// sig := <-signalChannel
	// log.Println("received terminate, graceful shutdown\n", sig)

	// timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// userDefinedServer.Shutdown(timeoutContext)

	//Using gorilla framework
	log := log.New(os.Stdout, "product-api", log.LstdFlags)

	productHandler := handlers.NewProducts(log)

	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productHandler.GetProducts)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productHandler.UpdateProduct)
	putRouter.Use(productHandler.MiddlewareProductValidation)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productHandler.AddProduct)
	postRouter.Use(productHandler.MiddlewareProductValidation)

	//Cross Origin Resource Sharing(CORS)
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:9090"}))

	userDefinedServer := &http.Server{
		Addr:         ":9090",
		Handler:      corsHandler(serveMux),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	go func() {
		err := userDefinedServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	log.Println("received terminate, graceful shutdown\n", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	userDefinedServer.Shutdown(timeoutContext)

}
