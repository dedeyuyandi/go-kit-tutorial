package transport

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/testing/go-kit-tutorial/endpoint"
	"github.com/testing/go-kit-tutorial/model"
)

func NewHTTPServer(ctx context.Context, endpoint endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoint.CreateUser,
		decodeUserReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoint.GetUser,
		decodeGetUserByIDReq,
		encodeResponse,
	))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.WriteUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserByIDReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.GetUserRequest
	vars := mux.Vars(r)

	ids, _ := uuid.FromString(vars["id"])
	req = model.GetUserRequest{
		Id: ids,
	}
	return req, nil
}
