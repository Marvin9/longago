package process

import (
	"bufio"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/Marvin9/atlan-collect/utils"
)

// MakeThreadForUploadProcess is used to start/resume the upload process.
// It creates upload thread and return immediately.
// Upload thread has controller stored in memory, it helps to pause/stop in next request(s)
// Thread will clear controller & process from memory once uploading is done
// inputFile is file which came from request
// fileNameToWrite is unique name of input file
// offsetByte helps from where to start reading input file (0 - if init, n - if this call is to resume upload)
func MakeThreadForUploadProcess(inputFile multipart.File, fileNameToWrite string, offsetByte int) utils.Controllers {
	filePathToWrite := utils.StoragePrefix + fileNameToWrite
	ctrl := utils.Controllers{
		Pause:  make(chan bool),
		Cancel: make(chan bool),
	}

	// store controller in memory
	utils.StoreInController(fileNameToWrite, ctrl)

	// start thread
	go (func() {
		fileToRead := inputFile
		defer fileToRead.Close()

		var fileToWrite *os.File
		var err error

		if offsetByte != 0 {
			// if file is already on disk, open it in append mode
			fileToWrite, err = os.OpenFile(filePathToWrite, os.O_APPEND|os.O_WRONLY, 0644)
		} else {
			// create new file
			fileToWrite, err = os.Create(filePathToWrite)
		}
		if err != nil {
			utils.Log(fmt.Sprintf("Error creating/opening file %v.\n%v", filePathToWrite, err))
			return
		}
		defer fileToWrite.Close()

		fileReader := bufio.NewReader(fileToRead)
		// discard offsetByte
		fileReader.Discard(offsetByte)
		fileWriter := bufio.NewWriter(fileToWrite)

		buffer := make([]byte, utils.BufferSize)
		offset := offsetByte

		// make new instance
		instance := utils.Instance{
			FileWhereUploaded: filePathToWrite,
			Offset:            offset,
			State:             utils.RUNNING,
		}
		// store instance in memory
		utils.StoreInProcess(fileNameToWrite, instance)

		// start uploading file
		for {
			// use timer when testing manually, so that you get enough time to pause/resume/stop uploading task
			// time.Sleep(time.Millisecond * 20)
			select {
			case <-ctrl.Pause:
				utils.Log("Pausing instance")
				instance.Offset = offset
				instance.State = utils.PAUSED

				// destroy controller, when resume it is new thread.
				utils.RemoveController(fileNameToWrite)

				// we have to store instance in memory, so next time if resumed it could start where it left
				utils.StoreInProcess(fileNameToWrite, instance)
				return
			case <-ctrl.Cancel:
				utils.Log("Cancelling instance")
				utils.Clear(fileNameToWrite)
				return
			default:
				read, _ := fileReader.Read(buffer)
				// keep eye on offset
				offset += read

				// when uploading task is done
				if read == 0 {
					utils.Log("Done writing file")
					utils.RemoveController(fileNameToWrite)
					utils.RemoveProcess(fileNameToWrite)
					// we can notify using socket
					return
				}

				fileWriter.Write(buffer[:read])
				fileWriter.Flush()
			}
		}
	})()

	return ctrl
}
