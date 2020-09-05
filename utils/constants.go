package utils

import (
	"os"
)

const (
	small  = "foo.csv"
	medium = "10.csv"
	large  = "100000.csv"
	// BufferSize - size to be read in one loop
	BufferSize = 1024
	// HTMLFileBodyName <input name="file" /> in request
	HTMLFileBodyName = "file"
)

// StoragePrefix will be used as upload storage path in application
var StoragePrefix = os.Getenv("UPLOAD_STORAGE") + "/"

var use = large
