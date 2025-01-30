package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jack-cordery/test-go/db"
	"github.com/jackc/pgx/v5"
)

type StatusResponse struct {
	OldMessage string
	NewMessage string
}

type HealthMessage struct {
	Message string
}

type CreateUserRequest struct {
	Email string `json:"email"`
}

type GetUserRequest struct {
	Email string `json:"email"`
}

func healthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(HealthMessage{
			Message: "NEW",
		})
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return

	}
}
}

func getStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(StatusResponse{
			OldMessage: "Who?",
			NewMessage: "When?",
		})
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}
}

func createUser(queries *db.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateUserRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Oh no", http.StatusInternalServerError)
			return
		}
		_, err = queries.CreateUser(ctx, req.Email)
		if err != nil {
			http.Error(w, "Oh no", http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(http.StatusOK)
		if err != nil {
			http.Error(w, "Oh no", http.StatusInternalServerError)
			return
		}
	}
}

func getUser(queries *db.Queries, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GetUserRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Oh no", http.StatusInternalServerError)
			return
		}
		user, err := queries.GetUserByEmail(ctx, req.Email)
		if err != nil {
			http.Error(w, "Oh no", http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, "Oh no", http.StatusInternalServerError)
			return
		}
	}
}


func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	queries := db.New(conn)

	mux.HandleFunc("GET /go/status", getStatus())
	mux.HandleFunc("GET /go/healthz", healthz())
	mux.HandleFunc("PUT /go/user", createUser(queries, ctx))
	mux.HandleFunc("GET /go/user", getUser(queries, ctx))

	err = http.ListenAndServe(":8001", mux)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

}