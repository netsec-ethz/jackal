/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package cluster

import (
	"encoding/gob"

	gobserializer "github.com/ortuman/jackal/model/gob"

	"github.com/ortuman/jackal/xmpp"
	"github.com/ortuman/jackal/xmpp/jid"
)

const (
	// MsgBind represents a bind cluster message.
	MsgBind = iota

	// MsgBatchBind represents a batch bind cluster message.
	MsgBatchBind

	// MsgUnbind represents a unbind cluster message.
	MsgUnbind

	// MsgUpdatePresence represents an update presence cluster message.
	MsgUpdatePresence

	// MsgUpdateContext represents a context update cluster message.
	MsgUpdateContext

	// MsgRouteStanza represents a route stanza cluster message.
	MsgRouteStanza
)

const (
	messageStanza = iota
	presenceStanza
	iqStanza
)

// MessagePayload represents a message payload type.
type MessagePayload struct {
	JID     *jid.JID
	Context map[string]interface{}
	Stanza  xmpp.Stanza
}

// FromGob reads MessagePayload fields from its gob binary representation.
func (p *MessagePayload) FromGob(dec *gob.Decoder) {
	p.JID = jid.NewFromGob(dec)

	var hasContextMap bool
	gobserializer.Decode(dec, &hasContextMap)
	if hasContextMap {
		var m map[string]interface{}
		gobserializer.Decode(dec, &m)
		p.Context = m
	}

	var hasStanza bool
	gobserializer.Decode(dec, &hasStanza)
	if !hasStanza {
		return
	}
	var stanzaType int
	gobserializer.Decode(dec, &stanzaType)
	switch stanzaType {
	case messageStanza:
		p.Stanza = xmpp.NewMessageFromGob(dec)
	case presenceStanza:
		p.Stanza = xmpp.NewPresenceFromGob(dec)
	case iqStanza:
		p.Stanza = xmpp.NewIQFromGob(dec)
	}
}

// ToGob converts a MessagePayload instance to its gob binary representation.
func (p *MessagePayload) ToGob(enc *gob.Encoder) {
	p.JID.ToGob(enc)

	hasContextMap := p.Context != nil
	gobserializer.Encode(enc, hasContextMap)
	if hasContextMap {
		gobserializer.Encode(enc, &p.Context)
	}

	hasStanza := p.Stanza != nil
	gobserializer.Encode(enc, hasStanza)
	if !hasStanza {
		return
	}
	// store stanza type
	switch p.Stanza.(type) {
	case *xmpp.Message:
		gobserializer.Encode(enc, messageStanza)
	case *xmpp.Presence:
		gobserializer.Encode(enc, presenceStanza)
	case *xmpp.IQ:
		gobserializer.Encode(enc, iqStanza)
	default:
		return
	}
	p.Stanza.ToGob(enc)
}

// Message is the c2s message type.
// A message can contain one or more payloads.
type Message struct {
	Type     int
	Node     string
	Payloads []MessagePayload
}

// FromGob reads Message fields from its gob binary representation.
func (m *Message) FromGob(dec *gob.Decoder) {
	gobserializer.Decode(dec, &m.Type)
	gobserializer.Decode(dec, &m.Node)

	var pLen int
	gobserializer.Decode(dec, &pLen)

	m.Payloads = nil
	for i := 0; i < pLen; i++ {
		var p MessagePayload
		p.FromGob(dec)
		m.Payloads = append(m.Payloads, p)
	}
}

// ToGob converts a Message instance to its gob binary representation.
func (m *Message) ToGob(enc *gob.Encoder) {
	gobserializer.Encode(enc, m.Type)
	gobserializer.Encode(enc, m.Node)
	gobserializer.Encode(enc, len(m.Payloads))
	for _, p := range m.Payloads {
		p.ToGob(enc)
	}
}
