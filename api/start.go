package api

import (
	"net/http"

	"github.com/Marvin9/atlan-collect/api/layer"

	"github.com/Marvin9/atlan-collect/process"
	"github.com/Marvin9/atlan-collect/utils"
)

func StartUploadAPI(w http.ResponseWriter, req *http.Request) {
	file, fileToBeWritten, err := layer.ExtractFileInfo(req)
	if err != nil {
		utils.Log(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(utils.SetResponse(true, "Internal error occured. Please try again."))
		return
	}

	if utils.GetProcessState(fileToBeWritten) == utils.RUNNING {
		w.WriteHeader(http.StatusConflict)
		w.Write(utils.SetResponse(true, "Another instance for same file is running."))
		return
	}

	if utils.GetProcessState(fileToBeWritten) == utils.PAUSED {
		w.WriteHeader(http.StatusConflict)
		w.Write(utils.SetResponse(true, "Another instance is paused. Please cancel it or resume it before starting new."))
		return
	}

	process.MakeThreadForUploadProcess(file, fileToBeWritten, 0)

	w.WriteHeader(http.StatusOK)
	w.Write(utils.SetResponse(false, fileToBeWritten))
}
