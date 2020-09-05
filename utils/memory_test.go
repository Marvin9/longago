package utils_test

import (
	"os"
	"testing"

	"github.com/Marvin9/atlan-collect/utils"
)

func TestProcessMemory(t *testing.T) {
	id := "id"
	_, ok := utils.GetProcess(id)
	if ok {
		t.Errorf("Process %v was not stored. Still returned true.", id)
	}

	state := utils.GetProcessState(id)
	if state != utils.NULL {
		t.Errorf("Process %v was not stored. Expected %v state, got %v", id, utils.NULL, state)
	}

	instance := utils.Instance{
		OriginalFileWithPath: "/foo/bar",
		FileWhereUploaded:    "/tmp",
		State:                utils.RUNNING,
		Offset:               0,
	}

	utils.StoreInProcess(id, instance)
	_, ok = utils.GetProcess(id)
	if !ok {
		t.Errorf("Process %v was stored. Still returned false", id)
	}

	state = utils.GetProcessState(id)
	if state != utils.RUNNING {
		t.Errorf("For process %v, Expected %v state got %v state.", id, utils.RUNNING, state)
	}

	utils.RemoveProcess(id)
	_, ok = utils.GetProcess(id)
	if ok {
		t.Errorf("Process %v was supposed to be removed.", id)
	}
}

func TestController(t *testing.T) {
	id := "id"
	_, ok := utils.GetController(id)
	if ok {
		t.Errorf("Controller named %v was not stored. still returned", id)
	}

	ctr := utils.Controllers{
		Pause:  make(chan bool),
		Cancel: make(chan bool),
	}

	utils.StoreInController(id, ctr)

	_, ok = utils.GetController(id)
	if !ok {
		t.Errorf("Controller named %v was stored. still not returned", id)
	}

	utils.RemoveController(id)

	_, ok = utils.GetController(id)
	if ok {
		t.Errorf("Controller named %v was removed. but still returned", id)
	}
}

func TestClear(t *testing.T) {
	fileName := "foo"
	os.Getenv("UPLOAD_STORAGE")
	path := utils.StoragePrefix + fileName
	os.MkdirAll(utils.StoragePrefix, os.ModePerm)
	os.Create(path)

	instance := utils.Instance{}
	ctrl := utils.Controllers{}

	utils.StoreInProcess(fileName, instance)
	utils.StoreInController(fileName, ctrl)

	utils.Clear(fileName)

	_, err := os.Stat(path)
	if err == nil {
		t.Errorf("Clear() should remove file. but path %v exists", path)
	}

	_, ok := utils.GetProcess(fileName)
	if ok {
		t.Errorf("Clear() should remove process in memory. For %v, it was not removed", fileName)
	}

	_, ok = utils.GetController(fileName)
	if ok {
		t.Errorf("Clear() should remove controller in memory. For %v, it was not removed", fileName)
	}
}
