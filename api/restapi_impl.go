package api

import (
	"encoding/json"
	"net/http"

	"invia/api/restapi"
)

type RestApiServer struct{}

func NewRestApiServer() RestApiServer {
	return RestApiServer{}
}

func (RestApiServer) ListUsers(w http.ResponseWriter, r *http.Request, params restapi.ListUsersParams) {
	resp := restapi.User{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (RestApiServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	resp := restapi.User{}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

func (RestApiServer) DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	w.WriteHeader(http.StatusNoContent)
}

func (RestApiServer) GetUserById(w http.ResponseWriter, r *http.Request, id int) {
	resp := restapi.User{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func (RestApiServer) UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	resp := restapi.User{}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
