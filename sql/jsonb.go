package sql

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONB[T any] struct {
	T T
}

func (ur JSONB[T]) Value() (driver.Value, error) {
	return json.Marshal(ur.T)
}

func (ur *JSONB[T]) Scan(value any) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, &ur.T)
	case string:
		return json.Unmarshal([]byte(v), &ur.T)
	default:
		return fmt.Errorf("unable to scan type %T into JSONB[%T]", value, ur.T)
	}
}
