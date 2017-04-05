package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeffmcnd/clik/repos"
	"github.com/jeffmcnd/clik/web/middleware"
)

func UserHandlers(r *mux.Router) {
	r.Handle("/v1/users/{id}", middleware.NoMiddleware(GetUserHandler)).Methods("GET")
	r.Handle("/v1/users/{id}/edit", middleware.NoMiddleware(EditUserHandler)).Methods("POST")
	r.Handle("/v1/users/{id}/queue", middleware.NoMiddleware(GetUserQueueHandler)).Methods("GET")
	r.Handle("/v1/users/{id}/matches", middleware.NoMiddleware(GetUserMatchesHandler)).Methods("GET")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	tokenString := r.Form.Get("access_token")

	if len(tokenString) == 0 {
		WriteError("No access token provided.", http.StatusBadRequest, w)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		WriteError("Invalid id.", http.StatusBadRequest, w)
		return
	}

	user, err := repos.DbGetUser(id)
	if err != nil {
		WriteError(err.Error(), http.StatusNotFound, w)
		return
	}

	fmt.Println(user)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
}

func GetUserQueueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		WriteError("Invalid id.", http.StatusBadRequest, w)
		return
	}

	userQueue, err := repos.DbGetUserQueue(id)
	if err != nil {
		WriteError(err.Error(), http.StatusNotFound, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userQueue); err != nil {
		panic(err)
	}
}

func GetUserMatchesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		WriteError("Invalid id.", http.StatusBadRequest, w)
		return
	}

	matches, err := repos.DbGetUserMatches(id)
	if err != nil {
		WriteError(err.Error(), http.StatusNotFound, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(matches); err != nil {
		panic(err)
	}
}