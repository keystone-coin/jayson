package jayson

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type HttpListener struct {
	server *Server
}

func (listen *HttpListener) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//  405 method not allowed if not POST
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte(http.StatusText(405)))
		return
	}
	// 415 unsupported media type if Content-Type is not correct
	if strings.HasPrefix(r.Header.Get("content-type"), "application/json") == false {
		w.WriteHeader(415)
		w.Write([]byte(http.StatusText(415)))
		return
	}
	// 400 Bad Request if body read all is error
	buf := bytes.NewBuffer(make([]byte, 0, r.ContentLength))
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()
	// analysis of JSONRPC request packet
	if buf.Len() == 0 {
		listen.sendError(w, ErrInvalidRequest())
		return
	}
	symbol, _, err := buf.ReadRune()
	if err != nil {
		listen.sendError(w, ErrInvalidRequest())
		return
	}
	err = buf.UnreadByte()
	if err != nil {
		listen.sendError(w, ErrInvalidRequest())
		return
	}
	rs := []*Request{}
	if symbol == '[' {
		// batch call
		err = json.NewDecoder(buf).Decode(&rs)
		if err != nil {
			listen.sendError(w, ErrParse())
			return
		}
		resp := listen.server.invoke(rs)
		listen.sendResponse(w, resp, true)
		return
	} else {
		// one call
		req := &Request{}
		err = json.NewDecoder(buf).Decode(&req)
		if err != nil {
			listen.sendError(w, ErrParse())
			return
		}
		rs = append(rs, req)
		resp := listen.server.invoke(rs)
		listen.sendResponse(w, resp, false)
		return
	}
}

func (listen *HttpListener) sendResponse(w http.ResponseWriter, resp []*Response, batch bool) error {
	w.Header().Set("content-type", "application/json")
	if len(resp) > 1 || batch {
		return json.NewEncoder(w).Encode(resp)
	} else if len(resp) == 1 {
		return json.NewEncoder(w).Encode(resp[0])
	}
	return nil
}

func (listen *HttpListener) sendError(w http.ResponseWriter, err *Error) error {
	resp := &Response{}
	resp.Version = "2.0"
	resp.Error = err
	return json.NewEncoder(w).Encode(resp)
}
