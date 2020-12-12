package Tool

import (
	"encoding/json"
	"io"
)

type StudentJsonParse struct {

}

func Decode(io io.ReadCloser,v interface{}) error {

	return json.NewDecoder(io).Decode(v)
	
}
