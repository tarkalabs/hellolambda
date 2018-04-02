package main

import (
	"flag"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	flag.Parse()
}

var initialized = false
var adapter *gorillamux.GorillaMuxAdapter

func defaultHandler(rw http.ResponseWriter, req *http.Request) {
	log.Info("Serving / ...")
	rw.Write([]byte("Hello lamdba world\n"))
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", defaultHandler).Methods("GET")
	return router
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if !initialized {
		log.Info("Performing cold start...")
		router := setupRouter()
		adapter = gorillamux.New(router)
		initialized = true
	}
	return adapter.Proxy(req)
}

func main() {
	if flag.Arg(0) == "local" {
		log.Info("Serving API Locally ...")
		http.ListenAndServe(":3000", setupRouter())
	} else {
		log.Info("Serving via lambda ...")
		lambda.Start(Handler)
	}
}
