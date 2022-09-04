package classfile

import (
	"fmt"

	"github.com/photowey/jvmgo/internal/classreader"
)

type ClassFile struct {
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []*AttributeInfo
}

func New() *ClassFile {
	return &ClassFile{}
}

// ---------------------------------------------------------------- parse

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); !ok {
				err = fmt.Errorf("%v", err)
			}
		}
	}()

	reader := classreader.New(classData)
	cf = New()
	cf.read(reader)

	return
}

// ---------------------------------------------------------------- read

func (cf *ClassFile) read(reader *classreader.ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.ReadUint16()
	cf.thisClass = reader.ReadUint16()
	cf.superClass = reader.ReadUint16()
	cf.interfaces = reader.ReadUint16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)

	return
}

func (cf *ClassFile) readAndCheckMagic(reader *classreader.ClassReader) {
	return
}

func (cf *ClassFile) readAndCheckVersion(reader *classreader.ClassReader) {
	return
}

// ---------------------------------------------------------------- getter

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

// ---------------------------------------------------------------- names

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SupperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}

	// only: java.lang.Object

	return ""
}

func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(cpIndex)
	}

	return interfaceNames
}

// ---------------------------------------------------------------- funcs

func readConstantPool(reader *classreader.ClassReader) ConstantPool {
	return ConstantPool{}
}

func readMembers(reader *classreader.ClassReader, constantPool ConstantPool) []*MemberInfo {
	return nil
}

func readAttributes(reader *classreader.ClassReader, constantPool ConstantPool) []*AttributeInfo {
	return nil
}
