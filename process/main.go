package process

import (
	"bufio"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/Marvin9/atlan-collect/utils"
)

func MakeThreadForUploadProcess(inputFile multipart.File, fileNameToWrite string, offsetByte int) utils.Controllers {
	filePathToWrite := utils.StoragePrefix + fileNameToWrite
	ctrl := utils.Controllers{
		Pause:  make(chan bool),
		Cancel: make(chan bool),
	}

	utils.StoreInController(fileNameToWrite, ctrl)

	go (func() {
		fileToRead := inputFile
		defer fileToRead.Close()

		var fileToWrite *os.File
		var err error
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
			FileWhereUploaded: filePathToWrite,
			Offset:            offset,
			State:             utils.RUNNING,
		}
		utils.StoreInProcess(fileNameToWrite, instance)

		for {
			time.Sleep(time.Millisecond * 20)
			select {
			case <-ctrl.Pause:
				utils.Log("Pausing instance")
				instance.Offset = offset
				instance.State = utils.PAUSED
				utils.StoreInProcess(fileNameToWrite, instance)
				return
			case <-ctrl.Cancel:
				utils.Log("Cancelling instance")
				utils.Clear(fileNameToWrite)
				return
			default:
				read, _ := fileReader.Read(buffer)
				offset += read
				if read == 0 {
					utils.Log("Done writing file")
					utils.RemoveController(fileNameToWrite)
					utils.RemoveProcess(fileNameToWrite)
					return
				}

				fileWriter.Write(buffer[:read])
				fileWriter.Flush()
			}
		}
	})()

	return ctrl
}
