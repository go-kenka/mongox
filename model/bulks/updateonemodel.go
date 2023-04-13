package bulks

import (
	"github.com/go-kenka/mongox/model/filters"
	"github.com/go-kenka/mongox/model/updates"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UpdateOneModel struct {
	model *mongo.UpdateOneModel
}

func NewUpdateOneModel() *UpdateOneModel {
	return &UpdateOneModel{
		model: mongo.NewUpdateOneModel(),
	}
}

// SetHint specifies the index to use for the operation. This should either be the index name as a string or the index
// specification as a document. This option is only valid for MongoDB versions >= 4.2. Server versions >= 3.4 will
// return an error if this option is specified. For server versions < 3.4, the driver will return a client-side error if
// this option is specified. The driver will return an error if this option is specified during an unacknowledged write
// operation. The driver will return an error if the hint parameter is a multi-key map. The default value is nil, which
// means that no hint will be sent.
func (uom *UpdateOneModel) SetHint(hint interface{}) *UpdateOneModel {
	uom.model.Hint = hint
	return uom
}

// SetFilter specifies a filter to use to select the document to update. The filter must be a document containing query
// operators. It cannot be nil. If the filter matches multiple documents, one will be selected from the matching
// documents.
func (uom *UpdateOneModel) SetFilter(filter filters.Filter) *UpdateOneModel {
	uom.model.Filter = filter.Document()
	return uom
}

// SetUpdate specifies the modifications to be made to the selected document. The value must be a document containing
// update operators (https://www.mongodb.com/docs/manual/reference/operator/update/). It cannot be nil or empty.
func (uom *UpdateOneModel) SetUpdate(update updates.Update) *UpdateOneModel {
	uom.model.Update = update.Document()
	return uom
}

// SetArrayFilters specifies a set of filters to determine which elements should be modified when updating an array
// field.
func (uom *UpdateOneModel) SetArrayFilters(filters options.ArrayFilters) *UpdateOneModel {
	uom.model.ArrayFilters = &filters
	return uom
}

// SetCollation specifies a collation to use for string comparisons. The default is nil, meaning no collation will be
// used.
func (uom *UpdateOneModel) SetCollation(collation *options.Collation) *UpdateOneModel {
	uom.model.Collation = collation
	return uom
}

// SetUpsert specifies whether or not a new document should be inserted if no document matching the filter is found. If
// an upsert is performed, the _id of the upserted document can be retrieved from the UpsertedIDs field of the
// BulkWriteResult.
func (uom *UpdateOneModel) SetUpsert(upsert bool) *UpdateOneModel {
	uom.model.Upsert = &upsert
	return uom
}

func (uom *UpdateOneModel) WriteModel() mongo.WriteModel {
	return uom.model
}
