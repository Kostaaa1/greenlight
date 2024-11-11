package data

import (
	"fmt"
	"strconv"
	"time"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

// type Movie struct {
// 	ID        int64     `json:"id"`
// 	CreatedAt time.Time `json:"-"`
// 	Title     string    `json:"title"`
// 	Year      int32     `json:"year,omitempty"`
// 	Runtime   int32     `json:"runtime"`
// 	Genres    []string  `json:"genres,omitempty"`
// 	Version   int32     `json:"version"`
// }

// func (m Movie) MarshalJSON() ([]byte, error) {
// 	var runtime string

// 	if m.Runtime != 0 {
// 		runtime = fmt.Sprintf("%d minutes", m.Runtime)
// 	}

// 	type MovieAlias Movie

// 	aux := struct {
// 		MovieAlias
// 		Runtime string `json:"runtime,omitempty"`
// 	}{
// 		MovieAlias: MovieAlias(m),
// 		Runtime:    runtime,
// 	}

// 	return json.Marshal(aux)
// }
