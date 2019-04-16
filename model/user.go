/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package model

import (
	"encoding/gob"
	"time"

	gobserializer "github.com/ortuman/jackal/model/gob"

	"github.com/ortuman/jackal/xmpp"
)

// User represents a user storage entity.
type User struct {
	Username       string
	Password       string
	LastPresence   *xmpp.Presence
	LastPresenceAt time.Time
}

// FromGob deserializes a User entity from it's gob binary representation.
func (u *User) FromGob(dec *gob.Decoder) error {
	gobserializer.Decode(dec, &u.Username)
	gobserializer.Decode(dec, &u.Password)
	var hasPresence bool
	gobserializer.Decode(dec, &hasPresence)
	if hasPresence {
		p, err := xmpp.NewPresenceFromGob(dec)
		if err != nil {
			return err
		}
		u.LastPresence = p
		gobserializer.Decode(dec, &u.LastPresenceAt)
	}
	return nil
}

// ToGob converts a User entity to it's gob binary representation.
func (u *User) ToGob(enc *gob.Encoder) {
	gobserializer.Encode(enc, u.Username)
	gobserializer.Encode(enc, u.Password)
	hasPresence := u.LastPresence != nil
	gobserializer.Encode(enc, hasPresence)
	if hasPresence {
		u.LastPresence.ToGob(enc)
		u.LastPresenceAt = time.Now()
		gobserializer.Encode(enc, &u.LastPresenceAt)
	}
}
