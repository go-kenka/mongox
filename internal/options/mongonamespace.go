package options

import "strings"

type MongoNamespace struct {
	databaseName   string
	collectionName string
	fullName       string
}

func NewMongoNamespace(fullName string) MongoNamespace {
	return MongoNamespace{
		databaseName:   GetDatatabaseNameFromFullName(fullName),
		collectionName: GetCollectionNameFullName(fullName),
		fullName:       fullName,
	}
}

func NewMongoNamespaceWithDB(databaseName, collectionName string) MongoNamespace {
	return MongoNamespace{
		databaseName:   databaseName,
		collectionName: collectionName,
		fullName:       databaseName + "." + collectionName,
	}
}

func GetCollectionNameFullName(namespace string) string {
	firstDot := strings.Index(namespace, ".")
	if firstDot == -1 {
		return namespace
	}

	return namespace[firstDot+1:]
}

func GetDatatabaseNameFromFullName(namespace string) string {
	firstDot := strings.Index(namespace, ".")
	if firstDot == -1 {
		return ""
	}

	return namespace[:firstDot]
}

func (m *MongoNamespace) DatabaseName() string {
	return m.databaseName
}

func (m *MongoNamespace) CollectionName() string {
	return m.collectionName
}
