package bulks

import (
	"github.com/go-kenka/mongox/model/filters"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReplaceOneModel struct {
	model *mongo.ReplaceOneModel
}

func NewReplaceOneModel() *ReplaceOneModel {
	return &ReplaceOneModel{
		model: mongo.NewReplaceOneModel(),
	}
}

// SetHint specifies the index to use for the operation. This should either be the index name as a string or the index
// specification as a document. This option is only valid for MongoDB versions >= 4.2. Server versions >= 3.4 will
// return an error if this option is specified. For server versions < 3.4, the driver will return a client-side error if
// this option is specified. The driver will return an error if this option is specified during an unacknowledged write
// operation. The driver will return an error if the hint parameter is a multi-key map. The default value is nil, which
// means that no hint will be sent.
func (rom *ReplaceOneModel) SetHint(hint interface{}) *ReplaceOneModel {
	rom.model.Hint = hint
	return rom
}

// SetFilter specifies a filter to use to select the document to replace. The filter must be a document containing query
// operators. It cannot be nil. If the filter matches multiple documents, one will be selected from the matching
// documents.
func (rom *ReplaceOneModel) SetFilter(filter filters.Filter) *ReplaceOneModel {
	rom.model.Filter = filter.Document()
	return rom
}

// SetReplacement specifies a document that will be used to replace the selected document. It cannot be nil and cannot
// contain any update operators (https://www.mongodb.com/docs/manual/reference/operator/update/).
func (rom *ReplaceOneModel) SetReplacement(rep interface{}) *ReplaceOneModel {
	rom.model.Replacement = rep
	return rom
}

// SetCollation specifies a collation to use for string comparisons. The default is nil, meaning no collation will be
// used.
func (rom *ReplaceOneModel) SetCollation(collation *options.Collation) *ReplaceOneModel {
	rom.model.Collation = collation
	return rom
}

// SetUpsert specifies whether or not the replacement document should be inserted if no document matching the filter is
// found. If an upsert is performed, the _id of the upserted document can be retrieved from the UpsertedIDs field of the
// BulkWriteResult.
func (rom *ReplaceOneModel) SetUpsert(upsert bool) *ReplaceOneModel {
	rom.model.Upsert = &upsert
	return rom
}

func (rom *ReplaceOneModel) WriteModel() mongo.WriteModel {
	return rom.model
}
