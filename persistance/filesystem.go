package persistance

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

//FSManager provides the list of operations for communicating with the file system.
type FSManager interface {
	LoadBlock(filename string) ([]byte, error)
	SaveBlock(block []byte, height int) (string, error)
}

const extension = "dat"
const blocksFilePath = "blockfiles"

//ErrSavingBlock error when saving raw block to filesystem
var ErrSavingBlock = errors.New("Could not save block from filesystem")

//ErrRetrievingBlock error when retrieving raw block from filesystem
var ErrRetrievingBlock = errors.New("Could not retrieve block from filesystem")

type fsManager struct {
	path string
}

//NewFSManager returns new instance of FSManager
func NewFSManager() FSManager {
	return newFSManager(blocksFilePath)
}

func newFSManager(path string) FSManager {
	return &fsManager{path}
}

//LoadBlock given a block file name ("block" + height) retrieves the serialized
//block instance
func (m *fsManager) LoadBlock(filename string) ([]byte, error) {
	path, _ := filepath.Abs("/" + m.path)
	content, err := ioutil.ReadFile(path + "/" + filename + "." + extension)
	if err != nil {
		log.Panicf("Could not retrieve block from disk, msg: %v\n", err)
		return nil, ErrRetrievingBlock
	}

	return content, nil
}

//SaveBlock given the serialized block  and the height, stores it to the disk
func (m *fsManager) SaveBlock(block []byte, height int) (string, error) {
	path, _ := filepath.Abs("/" + m.path)
	os.MkdirAll(path, os.ModePerm)
	err := ioutil.WriteFile(path+"/block"+strconv.Itoa(height)+"."+extension, block, 0644)

	if err != nil {
		log.Panicf("Could not store block to disk, msg: %v\n", err)
		return "", ErrSavingBlock
	}

	return "block" + strconv.Itoa(height), nil
}
