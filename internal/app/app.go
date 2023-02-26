package app

import (
	"fmt"
	"github.com/gorilla/mux"
	. "httpservice/internal/handler"
	. "httpservice/internal/model"
	. "httpservice/internal/service"
	. "httpservice/internal/storage"
	"log"
	"net/http"
)

var users = []*User{
	{0, "User1", 20, Friends{}},
	{0, "User2", 21, Friends{}},
	{0, "User3", 22, Friends{}},
	{0, "User4", 34, Friends{}},
	{0, "User5", 32, Friends{}},
	{0, "User6", 31, Friends{}},
	{0, "User7", 13, Friends{}},
	{0, "User8", 12, Friends{}},
	{0, "User9", 53, Friends{}},
	{0, "User10", 41, Friends{}},
}

func Run() {
	r := mux.NewRouter()
	storage := NewWorkStorage()

	for _, val := range users {
		err := storage.Add(val)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	service := NewService(storage)
	handler := NewHandler(service)
	handler.Register(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
