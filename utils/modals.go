package utils

import (
	"encoding/json"
)

type StateType string

const (
	RUNNING = "RUNNING_STATE"
	PAUSED  = "PAUSED_STATE"
	NULL    = "NULL_STATE"
)

type Instance struct {
	OriginalFileWithPath string
	FileWhereUploaded    string
	Offset               int
	State                StateType
}

type Controllers struct {
	Pause  chan bool
	Cancel chan bool
}

type RequestType struct {
	InstanceID string `json:"instance_id"`
}

type ResponseType struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data"`
}

func SetResponse(err bool, d interface{}) []byte {
	rt, _ := json.Marshal(ResponseType{
		Error: err,
		Data:  d,
	})
	return rt
}
