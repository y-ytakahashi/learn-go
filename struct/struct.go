package structPointer

import (
	"fmt"
)

type StructPointer struct{
	Point int
	Str string
}
func init(){
	fmt.Println("StructPointer.init()")
	StructPointerReference()
}
func StructPointerReference() *StructPointer{
	var spt = StructPointer{}
	spt.Point =123
	spt.Str = "文字列参照"

	fmt.Println(spt)
	return &spt
}