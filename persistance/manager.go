package persistance

//BlockMetadata stores necessary metadata into the DB in order to help
//retrieving the block from the filesystem 
type BlockMetadata struct{
	Height int
	Path string
}

//Manager exposes the list of methods available in the persistance 
//layer (fs & DB) of the application
type Manager interface{
	AddBlock(hash []byte, serializedBlock []byte ,blockMetadata *BlockMetadata) error
	RetrieveBlockByHash(hash []byte) []byte
}