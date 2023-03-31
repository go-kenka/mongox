package filters

import (
	"encoding/json"
	"testing"

	"github.com/go-kenka/mongox/bsonx"
)

func TestNot(t *testing.T) {
	n := Not(All("aaa", bsonx.String("aaa")))

	doc := n.Exp().AsDocument().Document()

	data, err := json.Marshal(doc)
	if err != nil {
		t.Fail()
	}
	t.Logf("%s", data)
	// t.Logf("%s", data)
}
