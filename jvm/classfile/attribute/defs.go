package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

const (
	Code                      = "Code"                      // 代码属性
	ConstantValue             = "ConstantValue"             // 常量
	Deprecated                = "Deprecated"                // 废弃的字段
	Exceptions                = "Exceptions"                // 异常
	LineNumberTable           = "LineNumberTable"           // 行数表
	LocalVariableTable        = "LocalVariableTable"        // 本地变量表
	SourceFile                = "SourceFile"                // 源文件
	Synthetic                 = "Synthetic"                 // 非用户代码生成
	ExceptionTable            = "Exception Table"           // 异常表
	StackMapTable             = "StackMapTable"             // 栈映射表
	InnerClasses              = "InnerClasses"              // 内部类
	RuntimeVisibleAnnotations = "RuntimeVisibleAnnotations" // 运行时可见注解
	BootstrapMethods          = "BootstrapMethods"          // 启动方法
	NestMembers               = "NestMembers"               // 内部成员
)

func ErrorMsgFmt(body, detail string, offset uint32) string {
	return fmt.Sprintf("[ERROR]:   %s (%s) @%d", body, detail, offset)
}

// Attribute 属性接口定义
type Attribute interface {
	Name() string   // 获取属性名
	Length() uint32 // 获取属性的长度
}

func New(r *reader.ByteCodeReader, cp constantpool.ConstantPool) Attribute {
	var name string
	if idx, ok := r.ReadU2(); ok {
		name = cp[idx].Value().(string)
	} else {
		panic(ErrorMsgFmt("Read attribute error", "can't read attribute_name_index info", r.Offset()))
	}
	switch name {
	case Code:
		return NewCodeAttribute(r, cp)
	case ConstantValue:
		return NewConstantValueAttribute(r, cp)
	case Deprecated:
		return NewDeprecatedAttribute(r, cp)
	case Exceptions:
		return NewExceptionsAttribute(r, cp)
	case LineNumberTable:
		return NewLineNumberTableAttribute(r, cp)
	case LocalVariableTable:
		return NewLocalVariableTableAttribute(r, cp)
	case SourceFile:
		return NewSourceFileAttribute(r, cp)
	case Synthetic:
		return NewSyntheticAttribute(r, cp)
	case StackMapTable:
		return NewStackMapTableAttribute(r, cp)
	case RuntimeVisibleAnnotations:
		return NewRuntimeVisibleAnnotationsAttribute(r, cp)
	case InnerClasses:
		return NewInnerClassAttribute(r, cp)
	case BootstrapMethods:
		return NewBootstrapMethodsAttribute(r, cp)
	case NestMembers:
		return NewNestMembersAttribute(r, cp)
	default:
		panic(ErrorMsgFmt("Unsupported attribute", name, r.Offset()))
	}
}
