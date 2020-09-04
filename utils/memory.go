package utils

import "os"

var processesInMemory = make(map[string]Instance)

var controllersInMemory = make(map[string]Controllers)

func GetProcess(forFile string) (Instance, bool) {
	i, ok := processesInMemory[forFile]
	return i, ok
}

func GetProcessState(forFile string) StateType {
	i, ok := GetProcess(forFile)
	if !ok {
		return NULL
	}
	return i.State
}

func StoreInProcess(forFile string, i Instance) {
	processesInMemory[forFile] = i
}

func RemoveProcess(forFile string) {
	delete(processesInMemory, forFile)
}

func GetController(forFile string) (Controllers, bool) {
	c, ok := controllersInMemory[forFile]
	return c, ok
}

func StoreInController(forFile string, c Controllers) {
	controllersInMemory[forFile] = c
}

func RemoveController(forFile string) {
	delete(controllersInMemory, forFile)
}

func Clear(forFile string) {
	os.Remove(forFile)
	RemoveProcess(forFile)
	RemoveController(forFile)
}
