// Package strings String expressions, with the exception of $concat, only have a well-defined behavior for strings of ASCII characters.
package strings

import (
	"github.com/go-kenka/mongox/bsonx"
)

type textOperator struct {
	doc bsonx.Bson
}

func (o textOperator) Exp() bsonx.IBsonValue {
	return o.doc.ToBsonDocument()
}

// Meta Returns the metadata associated with a document, e.g. "textScore" when
// performing text search. A $meta expression has the following syntax: { $meta:
// <metaDataKeyword> } The $meta expression can specify the following values as
// the <metaDataKeyword>
func Meta(metaDataKeyword string) textOperator {
	return textOperator{doc: bsonx.BsonDoc("$meta", bsonx.String(metaDataKeyword))}
}
