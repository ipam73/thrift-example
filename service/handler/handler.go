package handler

import (
	"fmt"
	"log"

	"github.com/nikhilpandit/thrift-example/gen-go/hello"
	"github.com/nikhilpandit/thrift-example/service/db"
)

type HelloHandler struct {
	database db.DB
}

func NewHelloHandler(database db.DB) *HelloHandler {
	return &HelloHandler{database}
}

func (h *HelloHandler) Ping() (bool, error) {
	if err := h.database.Ping(); err != nil {
		return false, &hello.HelloError{
			ErrorCode:    hello.HelloErrorCode_PING_ERROR,
			ErrorMessage: err.Error(),
		}
	}
	return true, nil
}

func (h *HelloHandler) Hello(username string) (string, error) {
	log.Println("Called Hello with username: ", username)
	person, err := h.database.GetPerson(username)
	if err != nil {
		return "", &hello.HelloError{
			ErrorCode:    hello.HelloErrorCode_NOT_FOUND,
			ErrorMessage: err.Error(),
		}
	}
	return fmt.Sprintf("Hello, %s", person.FirstName), nil
}
