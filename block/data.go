package block

//StoredData is the interface that abstracts the concrete struct each
//block is storing
type StoredData interface {
	ToString() string
}

//Bike is the underling struct that the blockchain is build around
//at this point is only a string, in the future the Bike will get
//more complicated
type Bike struct {
	SerialNumber string
}

//ToString returns a string direved from all Bike fields,
//for now we just return SerialNumber
func (b *Bike) ToString() string {
	return b.SerialNumber
}
