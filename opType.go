package enthistory

import (
	"database/sql/driver"
	"errors"
	"io"
	"strconv"
)

type OpType string

const (
	OpTypeInsert OpType = "INSERT"
	OpTypeUpdate OpType = "UPDATE"
	OpTypeDelete OpType = "DELETE"
)

var opTypes = []string{
	OpTypeInsert.String(),
	OpTypeUpdate.String(),
	OpTypeDelete.String(),
}

// Values provides list valid values for Enum.
func (OpType) Values() (kinds []string) {
	kinds = append(kinds, opTypes...)
	return
}

func (op OpType) Value() (driver.Value, error) {
	return op.String(), nil
}

func (op OpType) String() string {
	return string(op)
}

func (op *OpType) Scan(v any) error {
	if v == nil {
		*op = OpType("")
		return nil
	}

	switch src := v.(type) {
	case string:
		*op = OpType(src)
	case []byte:
		*op = OpType(string(src))
	case OpType:
		*op = src
	default:
		return errors.New("unsupported type")
	}

	return nil
}

func (op OpType) MarshalGQL(w io.Writer) {
	// graphql ID is a scalar which must be quoted
	io.WriteString(w, strconv.Quote(string(op))) //nolint:errcheck
}

func (op *OpType) UnmarshalGQL(v interface{}) error {
	return op.Scan(v)
}
