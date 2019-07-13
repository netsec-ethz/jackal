/*
 * Copyright (c) 2019 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package pubsubmodel

import (
	"encoding/gob"

	gobserializer "github.com/ortuman/jackal/model/gob"
)

type Node struct {
	Host    string
	Name    string
	Options Options
}

// FromGob deserializes a User entity from it's gob binary representation.
func (n *Node) FromGob(dec *gob.Decoder) error {
	gobserializer.Decode(dec, &n.Host)
	gobserializer.Decode(dec, &n.Name)
	gobserializer.Decode(dec, &n.Options)
	return nil
}

// ToGob converts a User entity to it's gob binary representation.
func (n *Node) ToGob(enc *gob.Encoder) {
	gobserializer.Encode(enc, n.Host)
	gobserializer.Encode(enc, n.Name)
	gobserializer.Encode(enc, n.Options)
}
