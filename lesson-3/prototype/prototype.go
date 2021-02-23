package prototype

type PrototypeStruct struct {
	SomeFieldString      string
	SomeFieldInt         int
	someUnexportedString string
}

func (p PrototypeStruct) Compare(comparable *PrototypeStruct) bool {
	return p.SomeFieldInt == comparable.SomeFieldInt &&
		p.SomeFieldString == comparable.SomeFieldString &&
		p.someUnexportedString == comparable.someUnexportedString
}

func (p PrototypeStruct) Clone() *PrototypeStruct {
	return &PrototypeStruct{
		SomeFieldString:      p.SomeFieldString,
		SomeFieldInt:         p.SomeFieldInt,
		someUnexportedString: p.someUnexportedString,
	}
}

func NewPrototypeStruct() *PrototypeStruct {
	return &PrototypeStruct{
		SomeFieldInt:         1,
		SomeFieldString:      "1",
		someUnexportedString: "unexp",
	}
}
