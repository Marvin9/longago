package api

import (
	"net/http"

	"github.com/Marvin9/atlan-collect/api/layer"

	"github.com/Marvin9/atlan-collect/utils"
)

func PauseAPI(w http.ResponseWriter, req *http.Request) {
	fileToBeWritten, _ := layer.ExtractFileToBeWritten(req)

	if utils.GetProcessState(fileToBeWritten) == utils.PAUSED {
		w.WriteHeader(http.StatusConflict)
		w.Write(utils.SetResponse(true, "Another instance is already paused. Please resume it or cancel it to start new instance"))
		return
	}

	controller, is := utils.GetController(fileToBeWritten)
	if !is {
		w.WriteHeader(http.StatusNoContent)
		w.Write(utils.SetResponse(true, "No running process"))
		return
	}

	controller.Pause <- true
	w.WriteHeader(http.StatusOK)
	w.Write(utils.SetResponse(false, fileToBeWritten))
}
