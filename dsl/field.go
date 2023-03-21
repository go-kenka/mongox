package dsl

type FieldFn func(f *FieldExpr)

type FieldExpr struct {
	Name       string       // 字段名称
	Type       MongoType    // 字段类型
	ObjectAttr []*FieldExpr // 只在Object和Array类型时生效
	ArrayType  MongoType    // Array类型时生效，如果是对象数组，则依赖ObjectAttr对象属性字段
	Subtype    byte         // binData时使用
	Comment    string       // 备注
}

func Field(name string, fns ...FieldFn) *FieldExpr {
	f := &FieldExpr{
		Name: name,
	}
	for _, o := range fns {
		o(f)
	}
	return f
}

func Type(typeInfo MongoType) FieldFn {
	return func(f *FieldExpr) {
		f.Type = typeInfo
	}
}

func Comment(com string) FieldFn {
	return func(f *FieldExpr) {
		f.Comment = com
	}
}

func ObjectAttr(fs ...*FieldExpr) FieldFn {
	return func(t *FieldExpr) {
		t.ObjectAttr = append(t.ObjectAttr, fs...)
	}
}

func ArrayType(typeInfo MongoType) FieldFn {
	return func(t *FieldExpr) {
		t.ArrayType = typeInfo
	}
}

// Subtype
// BinaryGeneric
// BinaryFunction
// BinaryBinaryOld
// BinaryUUIDOld
// BinaryUUID
// BinaryMD5
// BinaryEncrypted
// BinaryColumn
// BinaryUserDefined
func Subtype(subtype byte) FieldFn {
	return func(t *FieldExpr) {
		t.Subtype = subtype
	}
}
