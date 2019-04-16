/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package model

import (
	"encoding/gob"

	gobserializer "github.com/ortuman/jackal/model/gob"
)

// BlockListItem represents block list item storage entity.
type BlockListItem struct {
	Username string
	JID      string
}

// FromGob deserializes a BlockListItem entity from it's gob binary representation.
func (bli *BlockListItem) FromGob(dec *gob.Decoder) {
	gobserializer.Decode(dec, &bli.Username)
	gobserializer.Decode(dec, &bli.JID)
}

// ToGob converts a BlockListItem entity
// to it's gob binary representation.
func (bli *BlockListItem) ToGob(enc *gob.Encoder) {
	gobserializer.Encode(enc, bli.Username)
	gobserializer.Encode(enc, bli.JID)
}
