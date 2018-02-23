package persistance

import (
	"sync"
)

//BlockMetadata stores necessary metadata into the DB in order to help
//retrieving the block from the filesystem
type BlockMetadata struct {
	Height int
	Path   string
}

//Manager exposes the list of methods available in the persistance
//layer (fs & DB) of the application
type Manager interface {
	SaveBlock(hash []byte, serializedBlock []byte, blockMetadata *BlockMetadata) error
	RetrieveBlockByHash(hash []byte) ([]byte, error)
	LastUsedHash() []byte
	ClosePersistanceManager()
}

type manager struct {
	db DBManager
	fs FSManager
}

//NewPersistanceManager instantiate a new Manager for the persistance layer
func NewPersistanceManager() Manager {
	database := NewDBManager()
	filesystem := NewFSManager()

	return newPersistanceManager(database, filesystem)
}

func newPersistanceManager(database DBManager, filesystem FSManager) Manager {
	var once sync.Once
	var instance Manager
	/*
	   These should be only one instance of Manager, since  Bolt obtains a file lock on the data file
	   so multiple processes cannot open the same database at the same time.
	   Opening an already open Bolt database will cause it to hang until the other process closes it
	*/
	once.Do(func() {
		instance = &manager{
			db: database,
			fs: filesystem,
		}

	})
	return instance
}

//SaveBlock is responsible for persisting a block to filesystem & storing the metadata to DB.
func (m *manager) SaveBlock(hash []byte, serializedBlock []byte, blockMetadata *BlockMetadata) error {

	path, err := m.fs.SaveBlock(serializedBlock, blockMetadata.Height)
	if err != nil {
		return err
	}
	blockMetadata.Path = path

	err = m.db.SaveBlockMetadata(hash, blockMetadata)
	if err != nil {
		return err
	}

	return nil
}

/*RetrieveBlockByHash is used to fetch a block given a it's hash value.
First retrieves the metadata from the database, from the metadata retrieve the
file path and retrieve from the filesystem the raw block
*/
func (m *manager) RetrieveBlockByHash(hash []byte) ([]byte, error) {
	path := m.db.RetrieveBlockPathByHash(hash)
	block, err := m.fs.LoadBlock(path)
	if err != nil {
		return nil, err
	}

	return block, nil
}

//LastUsedHash returns the hash of the last stored block.
func (m *manager) LastUsedHash() []byte {
	return m.db.lastUsedHash()
}

//ClosePersistanceManager close any remaining connection
func (m *manager) ClosePersistanceManager() {
	m.db.CloseDb()
}
