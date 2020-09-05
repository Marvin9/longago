package layer

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/Marvin9/atlan-collect/utils"
)

// ExtractFileInfo parse file from request and provide unique name for that file.
func ExtractFileInfo(req *http.Request) (multipart.File, string, error) {
	req.ParseMultipartForm(20 << 20)
	file, handler, err := req.FormFile(utils.HTMLFileBodyName)
	if err != nil {
		return file, "", err
	}

	// unique name => [DATE AS PREFIX]__[ORIGINAL_FILENAME] (EVERYTHING WITHOUT SPACE)
	filename := handler.Filename
	filename = strings.ReplaceAll(time.Now().String(), " ", "-") + "__" + strings.ReplaceAll(filename, " ", "-")

	return file, filename, nil
}

// ExtractFileToBeWritten will extract instance_id(Unique instance that is stored in memory) from request
func ExtractFileToBeWritten(req *http.Request) (string, error) {
	var instance utils.RequestType
	err := json.NewDecoder(req.Body).Decode(&instance)
	if err != nil {
		return "", err
	}
	return instance.InstanceID, nil
}
