package main

// #include "wrapper.h"
import "C"
import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type JsonSchema struct {
	Data100 []Data100 `json:"data"`
}

type Data100 struct {
	Fix Fix `json:"fix"`
}

type Fix struct {
	ADEP_1 C.int `json:"adep-1"`
	ADEP_2 C.int `json:"adep-2"`
	ADEP_3 C.int `json:"adep-3"`
	ADEP_4 C.int `json:"adep-4"`
	ADEP_5 C.int `json:"adep-5"`
}

func unmarshal() JsonSchema {
	intputString := `{
		"data": [{
			"fix": {
				"adep-1":134,
				"adep-2":40,
				"adep-3":15,
				"adep-4":15,
				"adep-5":30
			}
		}]
	}`
	var schema JsonSchema
	err := json.Unmarshal([]byte(intputString), &schema)
	if err != nil {
		fmt.Println("could not unmarshal json")
	}

	return schema
}

func main() {
	C.wrapper((*C.Data)(unsafe.Pointer(&unmarshal().Data100[0].Fix)))
}
