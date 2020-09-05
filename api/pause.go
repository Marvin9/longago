package api

import (
	"net/http"

	"github.com/Marvin9/atlan-collect/api/layer"

	"github.com/Marvin9/atlan-collect/utils"
)

// PauseAPI - endpoint to pause uploading process
func PauseAPI(w http.ResponseWriter, req *http.Request) {
	fileToBeWritten, err := layer.ExtractFileToBeWritten(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.SetResponse(true, "Invalid body"))
		return
	}

	// If uploading process is not running... either it is already paused, or no instance exists
	if utils.GetProcessState(fileToBeWritten) != utils.RUNNING {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.SetResponse(true, "No instance to pause. or already paused one instance"))
		return
	}

	controller, is := utils.GetController(fileToBeWritten)
	if !is {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.SetResponse(true, "No running process"))
		return
	}

	controller.Pause <- true
	w.WriteHeader(http.StatusOK)
	w.Write(utils.SetResponse(false, fileToBeWritten))
}
