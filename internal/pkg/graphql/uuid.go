package graphql

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid/v5"
)

func MarshalUUID(t uuid.UUID) graphql.Marshaler {
	if t.IsNil() {
		return graphql.Null
	}
	return graphql.MarshalString(t.String())
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	if v, ok := v.(string); ok {
		return uuid.FromString(v)
	}
	return uuid.Nil, fmt.Errorf("%T is not a uuid", v)
}
