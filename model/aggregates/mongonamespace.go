package aggregates

import "strings"

type MongoNamespace struct {
	databaseName   string
	collectionName string
	fullName       string
}

func NewMongoNamespace(fullName string) MongoNamespace {
	return MongoNamespace{
		databaseName:   getDatatabaseNameFromFullName(fullName),
		collectionName: getCollectionNameFullName(fullName),
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

func getCollectionNameFullName(namespace string) string {
	firstDot := strings.Index(namespace, ".")
	if firstDot == -1 {
		return namespace
	}

	return namespace[firstDot+1:]
}

func getDatatabaseNameFromFullName(namespace string) string {
	firstDot := strings.Index(namespace, ".")
	if firstDot == -1 {
		return ""
	}

	return namespace[:firstDot]
}
