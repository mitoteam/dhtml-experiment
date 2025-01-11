package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mitoteam/dhtml-experiment/app"
	"github.com/mitoteam/mbr"
)

func main() {
	address := "localhost:15664"
	baseContext := context.Background()

	//Graceful shutdown according to https://github.com/gorilla/mux#graceful-shutdown
	httpSrv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 20,
		IdleTimeout:  time.Second * 60,
		Handler:      mbr.Handler(app.RootCtl),
		BaseContext:  func(l net.Listener) context.Context { return baseContext },
	}

	log.Printf("Starting up web server at http://%s\nPress Ctrl + C to stop it.\n", address)

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	cancel_channel := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(cancel_channel, os.Interrupt, os.Kill)

	// Block execution until we receive our signal.
	<-cancel_channel

	log.Println("Shutting down web server")

	// Create a deadline to wait for (10s).
	ctx, cancel := context.WithTimeout(baseContext, 10*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
