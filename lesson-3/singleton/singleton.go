package singleton

type SingleStruct struct {
	someField string
}

var singleStruct *SingleStruct

func NewSingleStruct() *SingleStruct {
	if singleStruct == nil {
		singleStruct = &SingleStruct{
			someField: "someValue",
		}
	}
	return singleStruct
}
