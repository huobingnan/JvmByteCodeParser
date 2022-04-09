package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
	"math"
)

type FloatConstant struct {
	value float32
}

func (f *FloatConstant) Tag() int {
	return ConstantFloatInfo
}

func (f *FloatConstant) Value() interface{} {
	return f.value
}

func (f *FloatConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Float_info: %f>", f.value)
}

func (f *FloatConstant) GoString() string {
	return f.String()
}

func _newFloatConstant(r *reader.ByteCodeReader) *FloatConstant {
	if float, ok := r.ReadU4(); ok {
		return &FloatConstant{
			value: math.Float32frombits(float),
		}
	} else {
		panic("Read float constant error")
	}
}
