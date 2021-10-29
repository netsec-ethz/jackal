/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package s2s

import (
	"context"
	"net"
	"strconv"
	"strings"

	"github.com/ortuman/jackal/log"

	"github.com/lucas-clemente/quic-go"
	"github.com/netsec-ethz/scion-apps/lib/scionutil"
	"github.com/scionproto/scion/go/lib/sciond"
	"github.com/scionproto/scion/go/lib/snet"
	"github.com/scionproto/scion/go/lib/snet/squic"
	"github.com/scionproto/scion/go/lib/sock/reliable"

	libaddr "github.com/scionproto/scion/go/lib/addr"
)

type Dialer interface {
	DialTCP(ctx context.Context, remoteDomain string) (net.Conn, error)
	DialQUIC(cfg *ScionConfig, remote *snet.Addr, localDomain, remoteDomain string) (quic.Session, error)
}

type srvResolveFunc func(service, proto, name string) (cname string, addrs []*net.SRV, err error)
type dialFunc func(ctx context.Context, network, address string) (net.Conn, error)

type dialer struct {
	srvResolve  srvResolveFunc
	dialContext dialFunc
}

func newDialer() *dialer {
	var d net.Dialer
	return &dialer{
		srvResolve:  net.LookupSRV,
		dialContext: d.DialContext,
	}
}

func (d *dialer) DialTCP(ctx context.Context, remoteDomain string) (net.Conn, error) {
	_, address, err := d.srvResolve("xmpp-server", "tcp", remoteDomain)
	if err != nil {
		log.Warnf("srv lookup error: %v", err)
	}
	var target string

	if err != nil || len(address) == 1 && address[0].Target == "." {
		target = remoteDomain + ":5269"
	} else {
		target = strings.TrimSuffix(address[0].Target, ".") + ":" + strconv.Itoa(int(address[0].Port))
	}
	conn, err := d.dialContext(ctx, "tcp", target)
	if err != nil {
		return nil, err
	}
	return conn, err
}

func (d *dialer) DialQUIC(cfg *ScionConfig, remote *snet.Addr, localDomain, remoteDomain string) (quic.Session, error) {
	var local *snet.Addr
	var err error
	if cfg.Address == "localhost" {
		local, err = scionutil.GetLocalhost()
	} else {
		local, err = snet.AddrFromString(cfg.Address)
	}
	if err != nil {
		return nil, err
	}

	sciondPath := sciond.GetDefaultSCIONDPath(nil)
	dispatcherPath := cfg.Dispatcher
	snet.Init(local.IA, sciondPath, reliable.NewDispatcherService(dispatcherPath))
	quicConfig := &quic.Config{
		KeepAlive: true,
	}
	sess, err := squic.DialSCION(nil, local, remote, quicConfig)
	if err != nil {
		return nil, err
	}
	return sess, err
}

func rainsLookup(remoteDomain string) (bool, *snet.Addr) {
	host, port, err := net.SplitHostPort(remoteDomain)
	if err != nil {
		host = remoteDomain
		port = "52690"
	}
	ia, l3, err := scionutil.GetHostByName(host + ".")
	if err != nil {
		return false, nil
	}

	p, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		p = 52690
	}
	l4 := libaddr.NewL4UDPInfo(uint16(p))
	raddr := &snet.Addr{IA: ia, Host: &libaddr.AppAddr{L3: l3, L4: l4}}

	return true, raddr
}
