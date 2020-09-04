package utils

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
