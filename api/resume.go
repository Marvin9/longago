package api

import (
	"net/http"

	"github.com/Marvin9/atlan-collect/process"

	"github.com/Marvin9/atlan-collect/utils"
)

func ResumeAPI(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(20 << 20)
	fileToBeWritten := req.FormValue("instance_id")

	if utils.GetProcessState(fileToBeWritten) != utils.PAUSED {
		w.WriteHeader(http.StatusConflict)
		w.Write(utils.SetResponse(true, "Instance is not paused, either it is running or not started."))
		return
	}

	previousInstance, is := utils.GetProcess(fileToBeWritten)
	if !is {
		w.WriteHeader(http.StatusNoContent)
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
