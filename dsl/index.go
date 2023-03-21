package dsl

import "time"

type Key struct {
	Key   string
	Value any // 1 升序 -1 降序 "2d" "2dsphere" "geoHaystack" "hashed" "text"
}

type IndexFn func(f *IndexExpr)

type IndexExpr struct {
	Name       string        // 索引名称
	Unique     bool          // 是否唯一
	Background bool          // 是否后台运行
	Sparse     bool          // 是否是稀疏索引（null不包含）
	Expire     time.Duration // 是否过期
	Keys       []Key         // 索引组合
}

func Index(name string, fns ...IndexFn) *IndexExpr {
	f := &IndexExpr{
		Name: name,
	}
	for _, o := range fns {
		o(f)
	}
	return f
}

func Sparse(s bool) IndexFn {
	return func(f *IndexExpr) {
		f.Sparse = s
	}
}

func Background(b bool) IndexFn {
	return func(f *IndexExpr) {
		f.Background = b
	}
}

func Expire(t time.Duration) IndexFn {
	return func(f *IndexExpr) {
		f.Expire = t
	}
}

func Unique(unique bool) IndexFn {
	return func(f *IndexExpr) {
		f.Unique = unique
	}
}

func SortKey(key string, val any) IndexFn {
	return func(f *IndexExpr) {
		f.Keys = append(f.Keys, Key{
			Key:   key,
			Value: val,
		})
	}
}
