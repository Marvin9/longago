package main

import (
	"bufio"
	"log"
	"os"
	"sync"
	"time"
)

func logger(m interface{}) {
	log.Printf("\n\t%v\n", m)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	small      = "foo.csv"
	medium     = "10.csv"
	large      = "100000.csv"
	bufferSize = 1024
)

var use = large

var uniqueFileToBeWritten = "./uploads/" + use

// Instance -
type Instance struct {
	originalFilePath string
	alreadyWritten   int
}

var pausedInstances = make(map[string]Instance)

var wg sync.WaitGroup

func instance(p Instance) chan bool {
	pause := make(chan bool)

	go (func() {
		defer wg.Done()
		fileToRead, err := os.Open(p.originalFilePath)
		checkErr(err)
		defer fileToRead.Close()

		var fileToWrite *os.File
		if p.alreadyWritten != 0 {
			fileToWrite, err = os.OpenFile(uniqueFileToBeWritten, os.O_APPEND|os.O_WRONLY, 0644)
		} else {
			fileToWrite, err = os.Create(uniqueFileToBeWritten)
		}
		checkErr(err)
		defer fileToWrite.Close()

		reader := bufio.NewReader(fileToRead)
		reader.Discard(p.alreadyWritten)
		writer := bufio.NewWriter(fileToWrite)

		buffer := make([]byte, bufferSize)
		offset := p.alreadyWritten
		for {
			select {
			case <-pause:
				pausedInstances[uniqueFileToBeWritten] = Instance{
					originalFilePath: p.originalFilePath,
					alreadyWritten:   offset,
				}
				return
			default:
				read, _ := reader.Read(buffer)
				offset += read
				if read == 0 {
					return
				}

				writer.Write(buffer[:read])
				writer.Flush()
			}
		}
	})()

	return pause
}

func main() {
	wg.Add(1)
	pause := instance(Instance{
		originalFilePath: "./fixtures/" + use,
	})
	time.Sleep(time.Millisecond * 10)
	pause <- true
	wg.Add(1)
	pause = instance(pausedInstances[uniqueFileToBeWritten])
	wg.Wait()
}
