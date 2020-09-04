package api

import (
	"fmt"
	"net/http"

	"github.com/Marvin9/atlan-collect/process"
	"github.com/Marvin9/atlan-collect/utils"
)

func StartUploadAPI(w http.ResponseWriter, req *http.Request) {
	if utils.GetProcessState(utils.FileToBeWritten) == utils.RUNNING {
		fmt.Fprintf(w, "Another instance for same file is running.")
		return
	}

	if utils.GetProcessState(utils.FileToBeWritten) == utils.PAUSED {
		fmt.Fprintf(w, "Another instance is paused. Please cancel it or resume it before starting new.")
		return
	}

	process.MakeThreadForUploadProcess(utils.FileToBeRead, utils.FileToBeWritten, 0)
	fmt.Fprintf(w, "Process started successfully.")
}
