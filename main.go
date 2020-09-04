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

// Controller -
type Controller struct {
	pause  chan bool
	cancel chan bool
}

var pausedInstances = make(map[string]Instance)

var wg sync.WaitGroup

func instance(p Instance) Controller {
	ctr := Controller{
		pause:  make(chan bool),
		cancel: make(chan bool),
	}

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
			case <-ctr.pause:
				logger("Pausing instance")
				pausedInstances[uniqueFileToBeWritten] = Instance{
					originalFilePath: p.originalFilePath,
					alreadyWritten:   offset,
				}
				return
			case <-ctr.cancel:
				logger("Cancelling instance")
				_, ok := pausedInstances[uniqueFileToBeWritten]
				if ok {
					delete(pausedInstances, uniqueFileToBeWritten)
				}
				os.Remove(uniqueFileToBeWritten)
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

	return ctr
}

func main() {
	wg.Add(1)
	ctr := instance(Instance{
		originalFilePath: "./fixtures/" + use,
	})
	time.Sleep(time.Millisecond * 10)
	ctr.pause <- true
	wg.Wait()
	wg.Add(1)
	ctr = instance(pausedInstances[uniqueFileToBeWritten])
	time.Sleep(time.Millisecond * 10)
	ctr.pause <- true
	wg.Wait()
	wg.Add(1)
	ctr = instance(pausedInstances[uniqueFileToBeWritten])
	time.Sleep(time.Millisecond * 10)
	ctr.cancel <- true
	wg.Wait()
}
