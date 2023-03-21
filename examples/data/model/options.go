package model

import "go.mongodb.org/mongo-driver/mongo/options"

type (
	AggregateOptions              = options.AggregateOptions
	ArrayFilters                  = options.ArrayFilters
	AutoEncryptionOptions         = options.AutoEncryptionOptions
	BucketOptions                 = options.BucketOptions
	BulkWriteOptions              = options.BulkWriteOptions
	ChangeStreamOptions           = options.ChangeStreamOptions
	ClientEncryptionOptions       = options.ClientEncryptionOptions
	ClientOptions                 = options.ClientOptions
	Collation                     = options.Collation
	CollectionOptions             = options.CollectionOptions
	CountOptions                  = options.CountOptions
	CreateCollectionOptions       = options.CreateCollectionOptions
	CreateIndexesOptions          = options.CreateIndexesOptions
	CreateViewOptions             = options.CreateViewOptions
	DataKeyOptions                = options.DataKeyOptions
	DatabaseOptions               = options.DatabaseOptions
	DefaultIndexOptions           = options.DefaultIndexOptions
	DeleteOptions                 = options.DeleteOptions
	DistinctOptions               = options.DistinctOptions
	DropIndexesOptions            = options.DropIndexesOptions
	EncryptOptions                = options.EncryptOptions
	EstimatedDocumentCountOptions = options.EstimatedDocumentCountOptions
	FindOneAndDeleteOptions       = options.FindOneAndDeleteOptions
	FindOneAndReplaceOptions      = options.FindOneAndReplaceOptions
	FindOneAndUpdateOptions       = options.FindOneAndUpdateOptions
	FindOneOptions                = options.FindOneOptions
	FindOptions                   = options.FindOptions
	GridFSFindOptions             = options.GridFSFindOptions
	IndexOptions                  = options.IndexOptions
	InsertManyOptions             = options.InsertManyOptions
	InsertOneOptions              = options.InsertOneOptions
	ListCollectionsOptions        = options.ListCollectionsOptions
	ListDatabasesOptions          = options.ListDatabasesOptions
	ListIndexesOptions            = options.ListIndexesOptions
	NameOptions                   = options.NameOptions
	ReplaceOptions                = options.ReplaceOptions
	RunCmdOptions                 = options.RunCmdOptions
	ServerAPIOptions              = options.ServerAPIOptions
	SessionOptions                = options.SessionOptions
	TimeSeriesOptions             = options.TimeSeriesOptions
	TransactionOptions            = options.TransactionOptions
	UpdateOptions                 = options.UpdateOptions
	UploadOptions                 = options.UploadOptions
)
