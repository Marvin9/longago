package api

import (
	"fmt"
	"net/http"

	"github.com/Marvin9/atlan-collect/utils"
)

func PauseAPI(w http.ResponseWriter, req *http.Request) {
	if utils.GetProcessState(utils.FileToBeWritten) == utils.PAUSED {
		fmt.Fprintf(w, "Another instance is already paused. Please resume it or cancel it to start new instance")
		return
	}

	controller, is := utils.GetController(utils.FileToBeWritten)
	if !is {
		fmt.Fprintf(w, "No running process")
		return
	}
	controller.Pause <- true
	fmt.Fprintf(w, "Successfully paused.")
}
