/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package rostermodel

import (
	"encoding/gob"

	gobserializer "github.com/ortuman/jackal/model/gob"
	"github.com/ortuman/jackal/xmpp"
)

// Notification represents a roster subscription pending notification.
type Notification struct {
	Contact  string
	JID      string
	Presence *xmpp.Presence
}

// FromGob deserializes a Notification entity from it's gob binary representation.
func (rn *Notification) FromGob(dec *gob.Decoder) {
	gobserializer.Decode(dec, &rn.Contact)
	gobserializer.Decode(dec, &rn.JID)
	rn.Presence = xmpp.NewPresenceFromGob(dec)
}

// ToGob converts a Notification entity
// to it's gob binary representation.
func (rn *Notification) ToGob(enc *gob.Encoder) {
	gobserializer.Encode(enc, rn.Contact)
	gobserializer.Encode(enc, rn.JID)

	rn.Presence.ToGob(enc)
}
