package main

import (
	"httpServer/Controllers"
	"httpServer/httpserver"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	server := httpserver.NewServer("app", 8081).SetAsyncNum(20)

	server.HandlerRequst("POST", "/sync", syncDemo)

	Controllers.InitHandler("/Controller").Test(server, "POST", "/test")

	// handler async http request
	server.HandlerAsyncRequst("POST", "/async", asyncDemo)

	go func() {
		if err := server.Start(); err != nil {
			log.Printf("server failed: %v", err)
		}
	}()

	EndChannel := make(chan os.Signal)
	signal.Notify(EndChannel, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	select {
	case output := <-EndChannel:
		log.Printf("end http server process by: %s", output)
		server.Stop()
	}
	close(EndChannel)
}

func syncDemo(jsonIn []byte) (jsonOut []byte, err error) {
	log.Printf("[syncDemo] jsonIn: %v", string(jsonIn[:]))

	return jsonIn, nil
}

func asyncDemo(jsonIn []byte) (err error) {
	time.Sleep(5 * time.Second)
	log.Printf("[asyncDemo] jsonIn: %v", string(jsonIn[:]))

	return nil
}
