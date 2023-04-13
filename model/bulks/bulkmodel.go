package bulks

import "go.mongodb.org/mongo-driver/mongo"

type BulkModel interface {
	WriteModel() mongo.WriteModel
}
