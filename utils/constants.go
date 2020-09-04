package utils

const (
	small      = "foo.csv"
	medium     = "10.csv"
	large      = "100000.csv"
	BufferSize = 1024
)

var use = large

var FileToBeWritten = "./tmp/uploads/" + use
var FileToBeRead = "./fixtures/" + use