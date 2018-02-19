package persistance

import(
	"io/ioutil"
	"log"
	"errors"
)

//FSManager provides the list of operations for communicating with the file system.
type FSManager interface{
	LoadBlock(filename string) ([]byte, error)
	SaveBlock(block []byte, height int) (string, error)
}

const extension ="dat"
var ErrSavingBlock = errors.New("Could not save block from filesystem")
var ErrRetrievingBlock = errors.New("Could not retrieve block from filesystem")

type fsManager struct{
	path string
}

//New returns new instance of Manager interface
func New(path string) FSManager{
	return &fsManager{path}
}

//LoadBlock given a block file name ("block" + height) retrieves the serialized
//block instance
func (m *fsManager) LoadBlock(filename string) ([]byte, error){
	content, err := ioutil.ReadFile(m.path +"/"+ filename +"."+extension)
	if err != nil {
		log.Panicf("Could not retrieve block from disk, msg: %v\n",err)
		return nil, ErrRetrievingBlock
	}

	return content, nil
}

//SaveBlock given the serialized block  and the height, stores it to the disk
func (m *fsManager) SaveBlock(block []byte, height int) (string, error){
	err:= ioutil.WriteFile(m.path+"/block"+string(height)+"."+extension, block, 0644)
	if err != nil {
		log.Panicf("Could not store block to disk, msg: %v\n",err)
		return "", ErrSavingBlock
	}

	return "block"+string(height), nil
}