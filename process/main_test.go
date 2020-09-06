package process_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"testing"

	"github.com/Marvin9/atlan-collect/process"
	"github.com/Marvin9/atlan-collect/utils"
)

func createMultipartFile(path string) (*bytes.Buffer, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(utils.HTMLFileBodyName, filepath.Base(path))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	writer.Close()

	return body, nil
}

func TestUploadProcess(t *testing.T) {
	// test write file in thread thread
	os.Getenv("UPLOAD_STORAGE")
	relativePath := "../fixtures/100000.csv"
	var f *os.File
	f, _ = os.Open(relativePath)

	filename := "100000.csv"
	process.MakeThreadForUploadProcess(f, filename, 0)
	process.Wg.Wait()

	pth := utils.StoragePrefix + filename
	_, err := os.Stat(pth)
	if err != nil {
		t.Errorf("File %v at %v was not written", filename, pth)
	}

	mainFile, _ := ioutil.ReadFile(relativePath)
	opFile, _ := ioutil.ReadFile(pth)

	if !bytes.Equal(mainFile, opFile) {
		t.Errorf("File %v is not equal as %v.", pth, relativePath)
	}

	// test pause file upload in thread
	f, _ = os.Open(relativePath)
	os.Remove(pth)
	ctrl := process.MakeThreadForUploadProcess(f, filename, 0)
	ctrl.Pause <- true
	process.Wg.Wait()
	opFile, _ = ioutil.ReadFile(pth)

	if bytes.Equal(mainFile, opFile) {
		t.Errorf("File %v was paused but is same as %v.", pth, relativePath)
	}

	// test cancel file upload in thread
	f, _ = os.Open(relativePath)
	os.Remove(pth)
	ctrl = process.MakeThreadForUploadProcess(f, filename, 0)
	ctrl.Cancel <- true
	process.Wg.Wait()
	_, err = os.Stat(pth)
	if err == nil {
		t.Errorf("File %v was cancelled, but still found on disk.", pth)
	}

	// test resume file upload in thread
	f, _ = os.Open(relativePath)
	os.Remove(pth)
	ctrl = process.MakeThreadForUploadProcess(f, filename, 0)
	ctrl.Pause <- true
	process.Wg.Wait()

	pausedInstance, _ := utils.GetProcess(filename)

	f, _ = os.Open(relativePath)
	process.MakeThreadForUploadProcess(f, filename, pausedInstance.Offset)
	process.Wg.Wait()
	opFile, _ = ioutil.ReadFile(pth)

	if !bytes.Equal(mainFile, opFile) {
		t.Errorf("File %v was paused and then resumed, but couldn't complete the upload properly.", pth)
	}
}
