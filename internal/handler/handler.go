package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	. "httpservice/hendlers"
	"httpservice/internal/model"
	"httpservice/internal/service"
	"log"
	"net/http"
	"strconv"
)

var _ Handler = &handler{}

const (
	usersURL   = "/users"
	userURL    = "/users/{id}"
	friendsURL = "/users/{id}/friends"
	friendURL  = "/users/{source_id}/friends/{target_id}"
)

type handler struct {
	service service.Service
}

func NewHandler(s service.Service) Handler {
	return &handler{service: s}
}

func (h *handler) Register(r *mux.Router) {
	r.HandleFunc(usersURL, h.GetAll).Methods("GET")
	r.HandleFunc(usersURL, h.CreateUser).Methods("POST")
	r.HandleFunc(userURL, h.AgeUpdate).Methods("PATCH")
	r.HandleFunc(userURL, h.DeleteUser).Methods("DELETE")

	r.HandleFunc(friendsURL, h.GetFriendList).Methods("GET")
	r.HandleFunc(friendsURL, h.AddFriend).Methods("PUT")
	r.HandleFunc(friendURL, h.DeleteFriend).Methods("DELETE")
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(h.service.GetAll())
	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := model.NewUser()
	_ = json.NewDecoder(r.Body).Decode(&user)
	err := h.service.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

func (h *handler) AgeUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	fmt.Println(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err.Error())
	}

	val, err := h.service.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err.Error())
	}

	val.Age = user.Age

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)

}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err.Error())
	}

	err = h.service.DeleteById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err.Error())
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) GetFriendList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err.Error())
	}

	user, err := h.service.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err.Error())
	}

	_ = json.NewEncoder(w).Encode(&user.Friends)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) AddFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err.Error())
	}

	source, err := h.service.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err.Error())
	}

	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	target, err := h.service.GetById(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err.Error())
	}

	if err = source.AddFriend(target); err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err.Error())
	}

	_ = json.NewEncoder(w).Encode(&source.Friends)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) DeleteFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	sourceId, err := strconv.Atoi(params["source_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err.Error())
	}

	targetId, err := strconv.Atoi(params["target_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err.Error())
	}

	source, err := h.service.GetById(sourceId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err.Error())
	}

	target, err := h.service.GetById(targetId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err.Error())
	}

	err = source.DeleteFriend(target.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err.Error())
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte(fmt.Sprintln(target.ID)))
}
