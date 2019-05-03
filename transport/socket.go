/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package transport

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io"
	"net"
	"strings"
	"time"

	"github.com/lucas-clemente/quic-go"
	"github.com/ortuman/jackal/log"
	"github.com/ortuman/jackal/transport/compress"
)

const socketBuffSize = 4096

type socketTransport struct {
	conn       net.Conn
	connQUIC   quic.Session
	rw         io.ReadWriter
	br         *bufio.Reader
	bw         *bufio.Writer
	keepAlive  time.Duration
	compressed bool
	isQUIC     bool
}

// NewSocketTransport creates a socket class stream transport.
func NewSocketTransport(conn net.Conn, keepAlive time.Duration) Transport {
	s := &socketTransport{
		conn:      conn,
		rw:        conn,
		br:        bufio.NewReaderSize(conn, socketBuffSize),
		bw:        bufio.NewWriterSize(conn, socketBuffSize),
		keepAlive: keepAlive,
	}
	return s
}

func NewQUICSocketTransport(conn quic.Session, uniStream quic.Stream,
	keepAlive time.Duration, connQUIC bool) Transport {
	s := &socketTransport{
		connQUIC:  conn,
		rw:        uniStream,
		br:        bufio.NewReaderSize(uniStream, socketBuffSize),
		bw:        bufio.NewWriterSize(uniStream, socketBuffSize),
		keepAlive: keepAlive,
		isQUIC:    connQUIC,
	}
	return s
}

func (s *socketTransport) Read(p []byte) (n int, err error) {
	if s.keepAlive > 0 && !s.isQUIC {
		s.conn.SetReadDeadline(time.Now().Add(s.keepAlive))
	}
	n, err = s.br.Read(p)
	if err != nil {
		log.Infof("Error reading the byte stream")
	}
	return n, err
}

func (s *socketTransport) Write(p []byte) (n int, err error) {
	defer s.bw.Flush()
	return s.bw.Write(p)
}

func (s *socketTransport) Close() error {
	if s.isQUIC {
		return s.connQUIC.Close()
	}
	return s.conn.Close()
}

func (s *socketTransport) Type() Type {
	return Socket
}

func (s *socketTransport) WriteString(str string) (int, error) {
	defer s.bw.Flush()
	n, err := io.Copy(s.bw, strings.NewReader(str))
	return int(n), err
}

func (s *socketTransport) StartTLS(cfg *tls.Config, asClient bool) {
	//NOTE!! GET RID OF THIS
	log.Infof("Starting TLS")
	if s.isQUIC {
		return
	}
	if _, ok := s.conn.(*net.TCPConn); ok {
		if asClient {
			s.conn = tls.Client(s.conn, cfg)
		} else {
			s.conn = tls.Server(s.conn, cfg)
		}
		s.rw = s.conn
		s.bw.Reset(s.rw)
		s.br.Reset(s.rw)
	}
}

func (s *socketTransport) EnableCompression(level compress.Level) {
	if !s.compressed {
		s.rw = compress.NewZlibCompressor(s.rw, s.rw, level)
		s.bw.Reset(s.rw)
		s.br.Reset(s.rw)
		s.compressed = true
	}
}

func (s *socketTransport) ChannelBindingBytes(mechanism ChannelBindingMechanism) []byte {
	if conn, ok := s.conn.(tlsStateQueryable); ok {
		switch mechanism {
		case TLSUnique:
			st := conn.ConnectionState()
			return st.TLSUnique
		default:
			break
		}
	}
	return nil
}

func (s *socketTransport) PeerCertificates() []*x509.Certificate {
	if conn, ok := s.conn.(tlsStateQueryable); ok {
		st := conn.ConnectionState()
		return st.PeerCertificates
	}
	return nil
}
