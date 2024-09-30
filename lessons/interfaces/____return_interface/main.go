package main

type DataInterface interface {
	Print()
	Serialize()
	Deserialize()
}

type Data struct{}

func NewData() DataInterface {
	return &Data{}
}

func (d *Data) Print()       {}
func (d *Data) Serialize()   {}
func (d *Data) Deserialize() {}

type Serializable interface {
	Serialize()
	Deserialize()
}

func main() {
	data := NewData()
	_ = Serializable(data)
}
