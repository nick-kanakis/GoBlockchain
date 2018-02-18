package block

//StoredData is the interface that abstracts the concrete struct each
//block is storing
type StoredData interface {
	ToByteSlices() []byte
}

//Bike is the underling struct that the blockchain is build around
//at this point is only a string, in the future the Bike will get
//more complicated
type Bike struct {
	SerialNumber string
}

//ToByteSlices returns a byte slice derived from all Bike fields,
//for now we just return SerialNumber
func (b *Bike) ToByteSlices() []byte {
	return []byte(b.SerialNumber)
}
