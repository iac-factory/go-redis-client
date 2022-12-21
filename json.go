package main

import (
	"errors"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/json"
	"log"
	"os"
	"path/filepath"
)

import . "encoding/json"

type Structure struct {
	file string
}

func (structure Structure) evaluate(exception error) {
	if exception == nil {
		return
	} else if errors.Is(exception, os.ErrNotExist) {
		panic(exception)
	} else {
		panic(exception)
	}
}

func (structure Structure) Path() string {
	metadata, exception := os.Stat("test.json")

	structure.evaluate(exception)

	pwd, exception := os.Getwd()
	structure.evaluate(exception)

	return filepath.Join(pwd, metadata.Name())
}

func (structure Structure) Stream() []byte {
	if buffer, exception := os.ReadFile(structure.Path()); exception == nil {
		return buffer
	} else {
		panic(exception)
	}
}

func (structure Structure) Value() cty.Value {
	if structure, exception := json.Unmarshal(structure.Stream(), structure.Type()); exception == nil {
		return structure
	} else {
		panic(exception)
	}
}

func (structure Structure) Type() cty.Type {
	var buffer = structure.Stream()

	if evaluation, exception := json.ImpliedType(buffer); exception == nil {
		return evaluation
	} else {
		panic(exception)
	}
}

// Buffer - Generate (marshal) a JSON-serialized structure according to a `cty.Value` `shape`.
// Generally to be used when storing or writing data as an export; for example,
// to ensure type-information isn't lost, Buffer could be interfaced via a Redis client
// so downstream `go` packages could deserialize the information
func (structure Structure) Buffer(shape cty.Type) []byte {
	if evaluation, exception := json.MarshalType(shape); exception == nil {
		if serial, exception := json.UnmarshalType(evaluation); exception == nil {
			if data, exception := MarshalIndent(serial, "", "    "); exception == nil {
				return data
			} else {
				panic(exception)
			}
		} else {
			panic(exception)
		}
	} else {
		panic(exception)
	}

	return nil
}

func Serialize() {
	var s = Structure{file: "test.json"}

	buffer := s.Buffer(s.Type())

	log.Println(string(buffer))

}
