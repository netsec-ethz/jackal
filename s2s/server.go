/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package s2s

import (
	"context"
	"net"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/lucas-clemente/quic-go"
	"github.com/netsec-ethz/scion-apps/lib/scionutil"
	streamerror "github.com/ortuman/jackal/errors"
	"github.com/ortuman/jackal/log"
	"github.com/ortuman/jackal/module"
	"github.com/ortuman/jackal/router"
	"github.com/ortuman/jackal/stream"
	"github.com/ortuman/jackal/transport"
	"github.com/scionproto/scion/go/lib/addr"
	"github.com/scionproto/scion/go/lib/sciond"
	"github.com/scionproto/scion/go/lib/snet"
	"github.com/scionproto/scion/go/lib/snet/squic"
)

var listenerProvider = net.Listen

type server struct {
	cfg            *Config
	router         *router.Router
	mods           *module.Modules
	dialer         *dialer
	inConns        sync.Map
	outConns       sync.Map
	ln             net.Listener
	lnQUIC         quic.Listener
	listening      uint32
	listeningSCION uint32
}

func (s *server) start() {
	bindAddr := s.cfg.Transport.BindAddress
	port := s.cfg.Transport.Port
	address := bindAddr + ":" + strconv.Itoa(port)

	log.Infof("s2s_in: listening at %s", address)

	if err := s.listenConn(address); err != nil {
		log.Fatalf("%v", err)
	}
}

func (s *server) startScion() {
	serverPort := uint16(s.cfg.Scion.Port)
	var address *snet.Addr
	var err error
	if s.cfg.Scion.Address == "localhost" {
		address, err = scionutil.GetLocalhost()
	} else {
		address, err = snet.AddrFromString(s.cfg.Scion.Address)
	}
	if err != nil {
		log.Fatalf("s2s_in: can't get local scion address")
	}
	address.Host.L4 = addr.NewL4UDPInfo(serverPort)

	if err := s.listenScionConn(address); err != nil {
		log.Fatalf("%v", err)
	}
	log.Infof("s2s_in: Listening for SCION s2s on port %d", serverPort)
}

func (s *server) listenScionConn(address *snet.Addr) error {
	var sciondPath string
	var dispatcherPath string = "/run/shm/dispatcher/default.sock"

	sciondPath = sciond.GetDefaultSCIONDPath(nil)
	snet.Init(address.IA, sciondPath, dispatcherPath)
	err := squic.Init(s.cfg.Scion.Key, s.cfg.Scion.Cert)
	if err != nil {
		return err
	}

	listener, err := squic.ListenSCION(nil, address, nil)
	if err != nil {
		return err
	}
	log.Infof("listening at %s", address)
	s.lnQUIC = listener
	atomic.StoreUint32(&s.listeningSCION, 1)
	for atomic.LoadUint32(&s.listeningSCION) == 1 {
		conn, err := s.lnQUIC.Accept()
		if err == nil {
			log.Infof("New SCION connection")
			accStream, err := conn.AcceptStream()
			if err != nil {
				log.Infof("No streams opened by the dialer")
			}
			isScion := true
			go s.startInStream(transport.NewQUICSocketTransport(conn, accStream,
				s.cfg.Scion.KeepAlive, true), isScion)
			continue
		}
	}

	return nil
}

func (s *server) shutdown(ctx context.Context) error {
	if atomic.CompareAndSwapUint32(&s.listening, 1, 0) {
		// stop listening...
		if err := s.ln.Close(); err != nil {
			return err
		}

		c, err := closeConnections(ctx, &s.inConns)
		if err != nil {
			return err
		}
		log.Infof("%s: closed %d in connection(s)", s.cfg.ID, c)
	}
	return nil
}

func (s *server) listenConn(address string) error {
	ln, err := listenerProvider("tcp", address)
	if err != nil {
		return err
	}
	s.ln = ln

	atomic.StoreUint32(&s.listening, 1)
	for atomic.LoadUint32(&s.listening) == 1 {
		conn, err := ln.Accept()
		if err == nil {
			isScion := false
			go s.startInStream(transport.NewSocketTransport(conn,
				s.cfg.Transport.KeepAlive), isScion)
			continue
		}
	}
	return nil
}

func (s *server) getOrDial(localDomain, remoteDomain string) (stream.S2SOut, error) {
	domainPair := localDomain + ":" + remoteDomain
	stm, loaded := s.outConns.LoadOrStore(domainPair, newOutStream(s.router, remoteDomain))
	if !loaded {
		outCfg, err := s.dialer.dial(localDomain, remoteDomain)
		if err != nil {
			log.Error(err)
			s.outConns.Delete(domainPair)
			return nil, err
		}
		outCfg.onOutDisconnect = s.unregisterOutStream

		stm.(*outStream).start(outCfg)
		log.Infof("registered s2s out stream... (domainpair: %s)", domainPair)
	}
	return stm.(*outStream), nil
}

func (s *server) unregisterOutStream(stm stream.S2SOut) {
	domainPair := stm.ID()
	s.outConns.Delete(domainPair)
	log.Infof("unregistered s2s out stream... (domainpair: %s)", domainPair)
}

func (s *server) startInStream(tr transport.Transport, isScion bool) {
	stm := newInStream(&streamConfig{
		keyGen:         &keyGen{s.cfg.DialbackSecret},
		transport:      tr,
		connectTimeout: s.cfg.ConnectTimeout,
		maxStanzaSize:  s.cfg.MaxStanzaSize,
		dialer:         s.dialer,
		onInDisconnect: s.unregisterInStream,
		streamSCION:    isScion,
	}, s.mods, s.router)
	s.registerInStream(stm)
}

func (s *server) registerInStream(stm stream.S2SIn) {
	s.inConns.Store(stm.ID(), stm)
	log.Infof("registered s2s in stream... (id: %s)", stm.ID())
}

func (s *server) unregisterInStream(stm stream.S2SIn) {
	s.inConns.Delete(stm.ID())
	log.Infof("unregistered s2s in stream... (id: %s)", stm.ID())
}

func closeConnections(ctx context.Context, connections *sync.Map) (count int, err error) {
	connections.Range(func(_, v interface{}) bool {
		stm := v.(*inStream)
		select {
		case <-closeConn(stm):
			count++
			return true
		case <-ctx.Done():
			err = ctx.Err()
			return false
		}
	})
	return
}

func closeConn(stm stream.InStream) <-chan bool {
	c := make(chan bool, 1)
	go func() {
		stm.Disconnect(streamerror.ErrSystemShutdown)
		c <- true
	}()
	return c
}
