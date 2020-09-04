package api

import (
	"fmt"
	"net/http"

	"github.com/Marvin9/atlan-collect/utils"
)

func StopAPI(w http.ResponseWriter, req *http.Request) {
	if utils.GetProcessState(utils.FileToBeWritten) == utils.NULL {
		fmt.Fprintf(w, "No instance to stop. Please start instance first")
		return
	}
	controller, is := utils.GetController(utils.FileToBeWritten)
	if !is {
		fmt.Fprintf(w, "No running process.")
		return
	}
	controller.Cancel <- true
	fmt.Fprintf(w, "Successfully stopped.")
}
