package expression

import "github.com/go-kenka/mongox/bsonx"

type Variable struct {
	bsonx.BsonValue
	start  string
	varStr string
}

func new(start string, varStr string) Variable {
	return Variable{
		start:  start,
		varStr: varStr,
	}
}

// Let gen User variable
func Let(varStr string) Variable {
	return new("$$", varStr)
}

// Path gen path variable
func Path(path string) Variable {
	return new("$", path)
}

func (v Variable) Exp() bsonx.IBsonValue {
	return bsonx.String(v.start + v.varStr)
}
