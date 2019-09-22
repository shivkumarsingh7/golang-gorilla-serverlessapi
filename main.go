package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
)

var gorillaLambda *gorillamux.GorillaMuxAdapter

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	router := mux.NewRouter()
	router.HandleFunc("/ping", PingHandler).Methods("GET")
	router.HandleFunc("/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/posts", getPosts).Methods("POST")

	gorillaLambda = gorillamux.New(router)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla ping!\n"))
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"name": "shivam", "age": "30"})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla Hello!\n"))
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return gorillaLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
