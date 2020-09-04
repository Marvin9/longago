package main

import (
	"log"
	"net/http"

	"github.com/Marvin9/atlan-collect/api"
)

func main() {
	http.HandleFunc("/p/start", api.StartUploadAPI)
	http.HandleFunc("/p/pause", api.PauseAPI)
	http.HandleFunc("/p/stop", api.StopAPI)
	http.HandleFunc("/p/resume", api.ResumeAPI)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
