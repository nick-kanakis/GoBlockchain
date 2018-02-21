package persistance

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
}

type manager struct {
	db DBManager
	fs FSManager
}

//New instantiate a new Manager for the persistance layer
func New() Manager {
	filesystem := NewFSManager()
	database := NewDBManager()

	return &manager{
		db: database,
		fs: filesystem,
	}
}

//SaveBlock is responsible for persisting a block to filesystem & storing the metadata to DB.
func (m *manager) SaveBlock(hash []byte, serializedBlock []byte, blockMetadata *BlockMetadata) error {

	path, err := m.fs.SaveBlock(serializedBlock, blockMetadata.Height)
	if err != nil {
		return err
	}
	blockMetadata.Path = path

	err = m.db.SaveBlock(hash, blockMetadata)
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
	return m.db.LastUsedHash()
}
