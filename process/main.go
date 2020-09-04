package process

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Marvin9/atlan-collect/utils"
)

func MakeThreadForUploadProcess(inputFilePath, filePathToWrite string, offsetByte int) utils.Controllers {
	ctrl := utils.Controllers{
		Pause:  make(chan bool),
		Cancel: make(chan bool),
	}

	utils.StoreInController(filePathToWrite, ctrl)

	go (func() {
		fileToRead, err := os.Open(inputFilePath)
		if err != nil {
			utils.Log(fmt.Sprintf("Error reading file %v.\n%v", inputFilePath, err))
			return
		}
		defer fileToRead.Close()

		var fileToWrite *os.File
		if offsetByte != 0 {
			fileToWrite, err = os.OpenFile(filePathToWrite, os.O_APPEND|os.O_WRONLY, 0644)
		} else {
			fileToWrite, err = os.Create(filePathToWrite)
		}
		if err != nil {
			utils.Log(fmt.Sprintf("Error creating/opening file %v.\n%v", filePathToWrite, err))
		}
		defer fileToWrite.Close()

		fileReader := bufio.NewReader(fileToRead)
		fileReader.Discard(offsetByte)
		fileWriter := bufio.NewWriter(fileToWrite)

		buffer := make([]byte, utils.BufferSize)
		offset := offsetByte

		instance := utils.Instance{
			OriginalFileWithPath: inputFilePath,
			FileWhereUploaded:    filePathToWrite,
			Offset:               offset,
			State:                utils.RUNNING,
		}
		utils.StoreInProcess(filePathToWrite, instance)

		for {
			time.Sleep(time.Millisecond * 1)
			select {
			case <-ctrl.Pause:
				utils.Log("Pausing instance")
				instance.Offset = offset
				instance.State = utils.PAUSED
				utils.StoreInProcess(filePathToWrite, instance)
				return
			case <-ctrl.Cancel:
				utils.Log("Cancelling instance")
				utils.Clear(filePathToWrite)
				return
			default:
				read, _ := fileReader.Read(buffer)
				offset += read
				if read == 0 {
					utils.Log("Done writing file")
					utils.RemoveController(filePathToWrite)
					utils.RemoveProcess(filePathToWrite)
					return
				}

				fileWriter.Write(buffer[:read])
				fileWriter.Flush()
			}
		}
	})()

	return ctrl
}
