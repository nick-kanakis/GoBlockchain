package persistance

import (
	"testing"
	"os"
	"path/filepath"
)

func TestSaveBlockLoadBlock(t *testing.T){
	fsManager :=New("test")
	path, _ := filepath.Abs("../../test")
	defer os.RemoveAll(path)

	filename, err:= fsManager.SaveBlock([]byte("Test data"), 12)
	if err!= nil {
		t.Error("Failed to save data")
	}
	data, err := fsManager.LoadBlock(filename)
	if err!= nil {
		t.Error("Failed to load data")
	}

	if "Test data" != string(data[:]){
		t.Error("Not original data")
	}
}