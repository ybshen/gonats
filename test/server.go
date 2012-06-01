package test

import (
	"bytes"
	"net"
	"testing"
)

type TestServer struct {
	net.Conn
	t *testing.T
}

func NewTestServer(t *testing.T, n net.Conn) *TestServer {
	var s = new(TestServer)

	s.t = t
	s.Conn = n

	return s
}

func (s *TestServer) AssertRead(v string) bool {
	var buf []byte
	var n int
	var e error

	buf = make([]byte, len(v))
	if n, e = s.Conn.Read(buf); e != nil {
		s.t.Errorf("\nerror: %#v\n", e)
		return false
	}

	var a []byte = []byte(v)
	var b []byte = buf[0:n]

	if !bytes.Equal(a, b) {
		s.t.Errorf("\nexpected: %#v\ngot: %#v\n", string(a), string(b))
		return false
	}

	return true
}

func (s *TestServer) AssertWrite(v string) bool {
	var e error

	if _, e = s.Conn.Write([]byte(v)); e != nil {
		s.t.Errorf("\nerror: %#v\n", e)
		return false
	}

	return true
}

func (s *TestServer) Close() {
	var e error

	if e = s.Conn.Close(); e != nil {
		panic(e)
	}
}
