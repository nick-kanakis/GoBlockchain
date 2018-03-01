package persistance

import (
	"encoding/json"
	"github.com/coreos/bbolt"
	"log"
)

const dbName = "blockchain.db"
const blockBucketName = "blocksBucket"
const lastHashBucketName = "lastHashBucket"

//DBManager provides the list of operations for communicating with the DB.
type DBManager interface {
	SaveBlockMetadata(hash []byte, blockMetadata *BlockMetadata) error
	RetrieveBlockPathByHash(hash []byte) (string, error)
	lastUsedHash() []byte
	CloseDb() error
}

type dbManager struct {
	db *bolt.DB
}

//NewDBManager instantiate a new DBManager
func NewDBManager() DBManager {
	return newDBManager(dbName)
}

func newDBManager(dbName string) DBManager {
	conn, err := bolt.Open(dbName, 0600, nil)

	if err != nil {
		log.Panicf("Could not open connection with DB msg: %v\n", err)
	}
	return &dbManager{
		db: conn,
	}
}

/*SaveBlockMetadata is responsible for storing block's metadata to DB.
There 2 buckets used in the DB:
	1)	Key = block hash
		Value = metadata of block (currently height of block and path in the filesystem)
	2)	Key = "lastUsedHash"
		Value = last inserted block hash
*/
func (m *dbManager) SaveBlockMetadata(hash []byte, blockMetadata *BlockMetadata) error {

	m.db.Update(func(tx *bolt.Tx) error {
		blockBucket, err := tx.CreateBucketIfNotExists([]byte(blockBucketName))
		if err != nil {
			return err
		}
		lastHashBucket, err := tx.CreateBucketIfNotExists([]byte(lastHashBucketName))
		if err != nil {
			return err
		}

		encoded, err := json.Marshal(blockMetadata)
		var metadata BlockMetadata
		json.Unmarshal(encoded, &metadata)
		if err != nil {
			return err
		}

		err = blockBucket.Put(hash, encoded)
		if err != nil {
			return err
		}

		return lastHashBucket.Put([]byte("lastUsedHash"), hash)
	})
	return nil
}

//RetrieveBlockByHash is used to fetch the block path on the disk
//given a it's hash value.
func (m *dbManager) RetrieveBlockPathByHash(hash []byte) (string, error) {
	var metadataJSON []byte
	var metadata BlockMetadata
	m.db.View(func(tx *bolt.Tx) error {
		blockBucket := tx.Bucket([]byte(blockBucketName))
		metadataJSON = blockBucket.Get(hash)
		return nil
	})

	err := json.Unmarshal(metadataJSON, &metadata)
	if err != nil {
		return "", err
	}
	path := metadata.Path
	return path, nil
}

//LastUsedHash returns the hash of the last stored block.
//The information is stored into "lasthash" bucket
func (m *dbManager) lastUsedHash() []byte {
	var hash []byte
	m.db.View(func(tx *bolt.Tx) error {
		lastHashBucket := tx.Bucket([]byte(lastHashBucketName))
		if lastHashBucket == nil {
			return nil
		}
		hash = lastHashBucket.Get([]byte("lastUsedHash"))
		return nil
	})
	return hash
}

func (m *dbManager) CloseDb() error {
	return m.db.Close()
}
