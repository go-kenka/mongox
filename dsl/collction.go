package dsl

type CollectionExpr struct {
	Database string       // 所属数据库
	Name     string       // 集合名称
	Desc     string       // 备注
	Fields   []*FieldExpr // 表字段集合
	Indexes  []*IndexExpr // 表字段集合
}

type CollectionFn func(t *CollectionExpr)

func Collection(name string, fns ...CollectionFn) *CollectionExpr {
	t := &CollectionExpr{
		Name: name,
	}

	for _, fn := range fns {
		fn(t)
	}
	return t
}

func Database(db string) CollectionFn {
	return func(t *CollectionExpr) {
		t.Database = db
	}
}
func Desc(desc string) CollectionFn {
	return func(t *CollectionExpr) {
		t.Desc = desc
	}
}
func Fields(fs ...*FieldExpr) CollectionFn {
	return func(t *CollectionExpr) {
		t.Fields = append(t.Fields, fs...)
	}
}
func Indexes(is ...*IndexExpr) CollectionFn {
	return func(t *CollectionExpr) {
		t.Indexes = append(t.Indexes, is...)
	}
}
