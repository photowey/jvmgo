package classfile

import (
	"github.com/photowey/jvmgo/internal/classreader"
)

const (
	ConstantUtf8               = 1
	ConstantInteger            = 3
	ConstantFloat              = 4
	ConstantLong               = 5
	ConstantDouble             = 6
	ConstantClass              = 7
	ConstantString             = 8
	ConstantFieldRef           = 9
	ConstantMethodRef          = 10
	ConstantInterfaceMethodRef = 11
	ConstantNameAndType        = 12
	ConstantMethodHandle       = 15
	ConstantMethodType         = 16
	ConstantInvokeDynamic      = 18
)

type ConstantInfo interface {
	readInfo(reader *classreader.ClassReader)
}
