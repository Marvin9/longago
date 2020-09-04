package api

import (
	"fmt"
	"net/http"

	"github.com/Marvin9/atlan-collect/process"

	"github.com/Marvin9/atlan-collect/utils"
)

func ResumeAPI(w http.ResponseWriter, req *http.Request) {
	if utils.GetProcessState(utils.FileToBeWritten) != utils.PAUSED {
		fmt.Fprintf(w, "Instance is not paused, either it is running or not started.")
		return
	}

	previousInstance, is := utils.GetProcess(utils.FileToBeWritten)
	if !is {
		fmt.Fprintf(w, "No previous found. Please start again.")
		return
	}
	process.MakeThreadForUploadProcess(utils.FileToBeRead, utils.FileToBeWritten, previousInstance.Offset)
	fmt.Fprintf(w, "Success.")
}
