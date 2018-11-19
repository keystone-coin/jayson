package jayson

import (
	"encoding/json"
)

type Request struct {
	ID      *json.RawMessage `json:"id"`
	Version string           `json:"jsonrpc"`
	Method  string           `json:"method"`
	Params  *json.RawMessage `json:"params"`
}

type Response struct {
	ID      *json.RawMessage `json:"id,omitempty"`
	Version string           `json:"jsonrpc"`
	Result  interface{}      `json:"result,omitempty"`
	Error   *Error           `json:"error,omitempty"`
}

type Method func(req *Request) *Response
