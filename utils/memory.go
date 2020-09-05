package utils

import "os"

var processesInMemory = make(map[string]Instance)

var controllersInMemory = make(map[string]Controllers)

// GetProcess is used to get instance for given file
// if file was paused, offset of instance would be useful to restart reading file
// from  where it was paused.
func GetProcess(forFile string) (Instance, bool) {
	i, ok := processesInMemory[forFile]
	return i, ok
}

// GetProcessState is used to get whether the given file is
// already being uploaded
// paused
// or not in memory
func GetProcessState(forFile string) StateType {
	i, ok := GetProcess(forFile)
	if !ok {
		return NULL
	}
	return i.State
}

// StoreInProcess is used to store instance in memory
func StoreInProcess(forFile string, i Instance) {
	processesInMemory[forFile] = i
}

// RemoveProcess is used to remove instance from memory
func RemoveProcess(forFile string) {
	delete(processesInMemory, forFile)
}

// GetController is used to get controller for given file (if exist in memory)
// controller is used to pause/stop running upload thread
func GetController(forFile string) (Controllers, bool) {
	c, ok := controllersInMemory[forFile]
	return c, ok
}

// StoreInController is used to store controller in memory
// used to store once upload process is started
func StoreInController(forFile string, c Controllers) {
	controllersInMemory[forFile] = c
}

// RemoveController is used to remove controller from memory
func RemoveController(forFile string) {
	delete(controllersInMemory, forFile)
}

// Clear is used when uploading process is stopped in middle.
// it will remove that file from storage and remove instances, controller from memory
func Clear(forFile string) {
	os.Remove(StoragePrefix + forFile)
	RemoveProcess(forFile)
	RemoveController(forFile)
}
