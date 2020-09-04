package api

import (
	"net/http"

	"github.com/Marvin9/atlan-collect/api/layer"
	"github.com/Marvin9/atlan-collect/utils"
)

func StopAPI(w http.ResponseWriter, req *http.Request) {
	fileToBeWritten, _ := layer.ExtractFileToBeWritten(req)

	if utils.GetProcessState(fileToBeWritten) == utils.NULL {
		w.WriteHeader(http.StatusConflict)
		w.Write(utils.SetResponse(true, "No instance to stop. Please start instance first"))
		return
	}
	controller, is := utils.GetController(fileToBeWritten)
	if !is {
		w.WriteHeader(http.StatusNoContent)
		w.Write(utils.SetResponse(true, "No running process to stop."))
		return
	}
	controller.Cancel <- true
	w.WriteHeader(http.StatusOK)
	w.Write(utils.SetResponse(false, "Successfully stopped upload."))
}
