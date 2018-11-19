package jayson

import (
	"sync"
)

type Server struct {
	methods map[string]Method
	mu      sync.Mutex
}

func NewServer() *Server {
	server := new(Server)
	return server
}

func (svr *Server) register(name string, method Method) {
	svr.mu.Lock()
	defer svr.mu.Unlock()
	if svr.methods == nil {
		svr.methods = make(map[string]Method)
	}
	svr.methods[name] = method
	return
}

func (svr *Server) invoke(rs []*Request) []*Response {
	resp := []*Response{}
	for _, req := range rs {
		method, exists := svr.methods[req.Method]
		if exists == false {
			resp = append(resp, &Response{
				Version: "2.0",
				Error:   ErrMethodNotFound(),
			})
			continue
		}
		resp = append(resp, method(req))
	}
	return resp
}

func (svr *Server) Register(name string, method Method) {
	svr.register(name, method)
}

func (svr *Server) Http() *HttpListener {
	return &HttpListener{
		server: svr,
	}
}
