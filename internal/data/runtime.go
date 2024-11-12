package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquoted, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquoted, " ")
	if len(parts) != 2 || !strings.HasPrefix(parts[1], "min") {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
