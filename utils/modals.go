package utils

import (
	"encoding/json"
)

// StateType is state of uploading instance
type StateType string

const (
	// RUNNING when upload process is running
	RUNNING = "RUNNING_STATE"
	// PAUSED when upload process is paused
	PAUSED = "PAUSED_STATE"
	// NULL when no upload process is going on
	NULL = "NULL_STATE"
)

// Instance is information of file which is being uploaded
type Instance struct {
	OriginalFileWithPath string
	FileWhereUploaded    string
	Offset               int
	State                StateType
}

// Controllers - to control uploading flow.
type Controllers struct {
	Pause  chan bool
	Cancel chan bool
}

// RequestType is used to json decode request body
type RequestType struct {
	InstanceID string `json:"instance_id"`
}

// ResponseType is used to make consisten json response of any request
type ResponseType struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

// SetResponse will make byte array of json response
func SetResponse(err bool, d interface{}) []byte {
	rt, e := json.Marshal(ResponseType{
		Error: err,
		Data:  d,
	})
	if e != nil {
		Log(e)
	}
	return rt
}
