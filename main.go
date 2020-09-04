package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Marvin9/atlan-collect/api/middleware"

	"github.com/Marvin9/atlan-collect/api"
	_ "github.com/joho/godotenv/autoload"
)

func usePort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":8000"
}

type apiType struct {
	path    string
	handler http.HandlerFunc
}

var apis = []apiType{
	apiType{
		path:    "/p/start",
		handler: api.StartUploadAPI,
	},
	apiType{
		path:    "/p/pause",
		handler: api.PauseAPI,
	},
	apiType{
		path:    "/p/stop",
		handler: api.StopAPI,
	},
	apiType{
		path:    "/p/resume",
		handler: api.ResumeAPI,
	},
}

func main() {
	for _, API := range apis {
		http.HandleFunc(API.path, middleware.SetJSONHeaderMiddleware(API.handler))
	}
	log.Fatal(http.ListenAndServe(usePort(), nil))
}
