package layer

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/Marvin9/atlan-collect/utils"
)

func ExtractFileInfo(req *http.Request) (multipart.File, string, error) {
	req.ParseMultipartForm(20 << 20)
	file, handler, err := req.FormFile(utils.HTMLFileBodyName)
	if err != nil {
		return file, "", err
	}

	filename := handler.Filename
	filename = strings.ReplaceAll(time.Now().String(), " ", "-") + "__" + strings.ReplaceAll(filename, " ", "-")

	return file, filename, nil
}

func ExtractFileToBeWritten(req *http.Request) (string, error) {
	var instance utils.RequestType
	err := json.NewDecoder(req.Body).Decode(&instance)
	if err != nil {
		return "", err
	}
	return instance.InstanceID, nil
}
