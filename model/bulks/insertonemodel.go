package bulks

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type InsertOneModel struct {
	model *mongo.InsertOneModel
}

func NewInsertOneModel() *InsertOneModel {
	return &InsertOneModel{
		model: mongo.NewInsertOneModel(),
	}
}

func (iom *InsertOneModel) WriteModel() mongo.WriteModel {
	return iom.model
}

// SetDocument specifies the document to be inserted. The document cannot be nil. If it does not have an _id field when
// transformed into BSON, one will be added automatically to the marshalled document. The original document will not be
// modified.
func (iom *InsertOneModel) SetDocument(doc any) *InsertOneModel {
	iom.model.SetDocument(doc)
	return iom
}
