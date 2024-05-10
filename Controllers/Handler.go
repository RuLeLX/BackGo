package Controllers

import (
	"httpServer/httpserver"
	"log"
)

type Handler struct {
	url string
}

func InitHandler(url string) *Handler {
	return &Handler{url: url}
}

func (handler Handler) Test(server *httpserver.Server, method string, url string) {

	server.HandlerRequst(method, handler.url+url, func(jsonIn []byte) (jsonOut []byte, err error) {
		log.Printf("[syncDemo] jsonIn: %v", string(jsonIn[:]))

		return jsonIn, nil
	})

}
