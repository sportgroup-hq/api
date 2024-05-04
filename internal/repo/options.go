package repo

import (
	"github.com/uptrace/bun"
)

type SelectOption func(query *bun.SelectQuery) *bun.SelectQuery

//func SimilarName(name string) QueryBuilder {
//	return func(query bun.QueryBuilder) bun.QueryBuilder {
//		if name == "" {
//			return query
//		}
//		name = "%" + name + "%"
//
//		return query.Where("first_name LIKE ? OR last_name LIKE ?", name, name)
//	}
//}
//
