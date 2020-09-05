package api

import (
	"net/http"

	"github.com/Marvin9/atlan-collect/api/layer"

	"github.com/Marvin9/atlan-collect/utils"
)

func PauseAPI(w http.ResponseWriter, req *http.Request) {
	fileToBeWritten, err := layer.ExtractFileToBeWritten(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.SetResponse(true, "Invalid body"))
		return
	}

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
