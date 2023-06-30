package mongox

import (
	"github.com/go-kenka/mongox/schema/field"
	"github.com/go-kenka/mongox/schema/index"
)

type (
	Interface interface {
		// Fields returns the fields of the schema.
		Fields() []Field
		// Indexes returns the indexes of the schema.
		Indexes() []Index
	}

	Field interface {
		Descriptor() *field.Descriptor
	}

	Index interface {
		Descriptor() *index.Descriptor
	}

	// Schema is the default implementation for the schema Interface.
	// It can be embedded in end-user schemas as follows:
	//
	//	type T struct {
	//		mongox.Schema
	//	}
	//
	Schema struct {
		Interface
	}
)

// Fields of the schema.
func (Schema) Fields() []Field { return nil }

// Indexes of the schema.
func (Schema) Indexes() []Index { return nil }
