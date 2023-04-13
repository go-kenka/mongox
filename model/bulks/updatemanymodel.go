package bulks

import (
	"github.com/go-kenka/mongox/model/filters"
	"github.com/go-kenka/mongox/model/updates"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UpdateManyModel struct {
	model *mongo.UpdateManyModel
}

func NewUpdateManyModel() *UpdateManyModel {
	return &UpdateManyModel{
		model: mongo.NewUpdateManyModel(),
	}
}

// SetHint specifies the index to use for the operation. This should either be the index name as a string or the index
// specification as a document. This option is only valid for MongoDB versions >= 4.2. Server versions >= 3.4 will
// return an error if this option is specified. For server versions < 3.4, the driver will return a client-side error if
// this option is specified. The driver will return an error if this option is specified during an unacknowledged write
// operation. The driver will return an error if the hint parameter is a multi-key map. The default value is nil, which
// means that no hint will be sent.
func (umm *UpdateManyModel) SetHint(hint interface{}) *UpdateManyModel {
	umm.model.Hint = hint
	return umm
}

// SetFilter specifies a filter to use to select documents to update. The filter must be a document containing query
// operators. It cannot be nil.
func (umm *UpdateManyModel) SetFilter(filter filters.Filter) *UpdateManyModel {
	umm.model.Filter = filter.Document()
	return umm
}

// SetUpdate specifies the modifications to be made to the selected documents. The value must be a document containing
// update operators (https://www.mongodb.com/docs/manual/reference/operator/update/). It cannot be nil or empty.
func (umm *UpdateManyModel) SetUpdate(update updates.Update) *UpdateManyModel {
	umm.model.Update = update.Document()
	return umm
}

// SetArrayFilters specifies a set of filters to determine which elements should be modified when updating an array
// field.
func (umm *UpdateManyModel) SetArrayFilters(filters options.ArrayFilters) *UpdateManyModel {
	umm.model.ArrayFilters = &filters
	return umm
}

// SetCollation specifies a collation to use for string comparisons. The default is nil, meaning no collation will be
// used.
func (umm *UpdateManyModel) SetCollation(collation *options.Collation) *UpdateManyModel {
	umm.model.Collation = collation
	return umm
}

// SetUpsert specifies whether or not a new document should be inserted if no document matching the filter is found. If
// an upsert is performed, the _id of the upserted document can be retrieved from the UpsertedIDs field of the
// BulkWriteResult.
func (umm *UpdateManyModel) SetUpsert(upsert bool) *UpdateManyModel {
	umm.model.Upsert = &upsert
	return umm
}

func (umm *UpdateManyModel) WriteModel() mongo.WriteModel {
	return umm.model
}
