package searchqueryspec

import (
	"encoding/json"
	"errors"

	"github.com/grafadruid/go-druid/builder"
)

type Base struct {
	Typ builder.ComponentType `json:"type,omitempty"`
}

func (b *Base) SetType(typ builder.ComponentType) *Base {
	b.Typ = typ
	return b
}

func (b *Base) Type() builder.ComponentType {
	return b.Typ
}

func Load(data []byte) (builder.SearchQuerySpec, error) {
	var t struct {
		Typ builder.ComponentType `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	var s builder.SearchQuerySpec
	switch t.Typ {
	case "all":
		s = NewAll()
	case "contains":
		s = NewContains()
	case "fragment":
		s = NewFragment()
	case "insensitiveContains":
		s = NewInsensitiveContains()
	case "regex":
		s = NewRegex()
	default:
		return nil, errors.New("unsupported type")
	}
	return s, json.Unmarshal(data, &s)
}
