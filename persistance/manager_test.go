package persistance

import (
	"testing"
)

func TestSaveBlock(t *testing.T) {
	mockedPersistanceManager := newPersistanceManager(&mockedDB{}, &mockedFS{})

	err := mockedPersistanceManager.SaveBlock([]byte{}, []byte{}, &BlockMetadata{1, ""})
	if err != nil {
		t.Error("Could not save block metadata")
	}
}

func TestRetrieveBlockByHash(t *testing.T) {
	mockedPersistanceManager := newPersistanceManager(&mockedDB{}, &mockedFS{})
	hash, err := mockedPersistanceManager.RetrieveBlockByHash([]byte{})

	if err != nil {
		t.Error("Could not load hash")
	}

	if "mocked block" != string(hash[:]) {
		t.Error("Could not correct hash value")
	}
}

func TestLastUsedHash(t *testing.T) {
	mockedPersistanceManager := newPersistanceManager(&mockedDB{}, &mockedFS{})
	hash := mockedPersistanceManager.LastUsedHash()

	if "1234" != string(hash[:]) {
		t.Error("Could not correct hash value")
	}
}

type mockedDB struct{}

func (m *mockedDB) SaveBlockMetadata(hash []byte, blockMetadata *BlockMetadata) error {
	return nil
}

func (m *mockedDB) RetrieveBlockPathByHash(hash []byte) string {
	return "/block1234.dat"
}

func (m *mockedDB) lastUsedHash() []byte {
	return []byte("1234")
}

func (m *mockedDB) CloseDb() {}

type mockedFS struct{}

func (m *mockedFS) LoadBlock(filename string) ([]byte, error) {
	return []byte("mocked block"), nil
}

func (m *mockedFS) SaveBlock(block []byte, height int) (string, error) {
	return "block1234", nil
}
