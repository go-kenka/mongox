package expressions

import "github.com/go-kenka/mongox/examples/data/bsonx"

type TExpression interface {
	bsonx.IBsonValue
}

type InExpression interface {
	bsonx.IBsonValue
}

type NExpression interface {
	bsonx.IBsonValue
}

type OutExpression interface {
	bsonx.IBsonValue
}

type TBoundary interface {
	bsonx.IBsonNumber
}

type TItem interface {
	bsonx.IBsonValue
}
