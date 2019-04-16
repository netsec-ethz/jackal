/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package rostermodel

import (
	"encoding/gob"

	gobserializer "github.com/ortuman/jackal/model/gob"
)

// Version represents a roster version info.
type Version struct {
	Ver         int
	DeletionVer int
}

// FromGob deserializes a Version entity
// from it's gob binary representation.
func (rv *Version) FromGob(dec *gob.Decoder) {
	gobserializer.Decode(dec, &rv.Ver)
	gobserializer.Decode(dec, &rv.DeletionVer)
}

// ToGob converts a Version entity
// to it's gob binary representation.
func (rv *Version) ToGob(enc *gob.Encoder) {
	gobserializer.Encode(enc, rv.Ver)
	gobserializer.Encode(enc, &rv.DeletionVer)
}
