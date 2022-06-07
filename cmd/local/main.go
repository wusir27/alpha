package main

import (
	"github.com/wusir27/alpha/pkg/local"
	"os"
	"os/signal"
	"syscall"
	"log"
	"sync"
)

func main() {
	local.LocalBootstrape()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <- sigChan
	go func(){
		sig := <- sigChan
		log.Printf("v% signal received, closing immediately", sig)
		os.Exit(255)
	}()

	log.Printf("%v signal received, closing alpha local", sig)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	local.Shutdown(wg)
	wg.Wait()
}
