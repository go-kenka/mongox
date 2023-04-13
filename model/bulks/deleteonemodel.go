package bulks

import (
	"github.com/go-kenka/mongox/model/filters"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeleteOneModel struct {
	model *mongo.DeleteOneModel
}

func NewDeleteOneModel() *DeleteOneModel {
	return &DeleteOneModel{
		model: mongo.NewDeleteOneModel(),
	}
}

// SetFilter specifies a filter to use to select the document to delete. The filter must be a document containing query
// operators. It cannot be nil. If the filter matches multiple documents, one will be selected from the matching
// documents.
func (dom *DeleteOneModel) SetFilter(filter filters.Filter) *DeleteOneModel {
	dom.model.Filter = filter.Document()
	return dom
}

// SetCollation specifies a collation to use for string comparisons. The default is nil, meaning no collation will be
// used.
func (dom *DeleteOneModel) SetCollation(collation *options.Collation) *DeleteOneModel {
	dom.model.Collation = collation
	return dom
}

// SetHint specifies the index to use for the operation. This should either be the index name as a string or the index
// specification as a document. This option is only valid for MongoDB versions >= 4.4. Server versions >= 3.4 will
// return an error if this option is specified. For server versions < 3.4, the driver will return a client-side error if
// this option is specified. The driver will return an error if this option is specified during an unacknowledged write
// operation. The driver will return an error if the hint parameter is a multi-key map. The default value is nil, which
// means that no hint will be sent.
func (dom *DeleteOneModel) SetHint(hint interface{}) *DeleteOneModel {
	dom.model.Hint = hint
	return dom
}

func (dom *DeleteOneModel) WriteModel() mongo.WriteModel {
	return dom.model
}
