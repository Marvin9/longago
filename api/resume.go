package api

import (
	"net/http"

	"github.com/Marvin9/atlan-collect/process"

	"github.com/Marvin9/atlan-collect/utils"
)

// ResumeAPI - endpoint to resume uploading task
func ResumeAPI(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(utils.UploadLimit)
	fileToBeWritten := req.FormValue("instance_id")

	// If process is not paused... either running or not exist
	if utils.GetProcessState(fileToBeWritten) != utils.PAUSED {
		w.WriteHeader(http.StatusConflict)
		w.Write(utils.SetResponse(true, "Instance is not paused, either it is running or not started."))
		return
	}

	previousInstance, is := utils.GetProcess(fileToBeWritten)
	if !is {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.SetResponse(true, "No previous instance found. Please upload again."))
		return
	}

	file, _, err := req.FormFile(utils.HTMLFileBodyName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.SetResponse(true, "Internal error occured. Please try again."))
		return
	}

	process.MakeThreadForUploadProcess(file, fileToBeWritten, previousInstance.Offset)
	w.WriteHeader(http.StatusOK)
	w.Write(utils.SetResponse(false, fileToBeWritten))
}
