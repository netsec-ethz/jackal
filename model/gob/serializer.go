/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package gobserializer

import (
	"encoding/gob"

	"github.com/ortuman/jackal/log"
)

// GobSerializer represents a Gob serializable entity.
type GobSerializer interface {
	ToGob(enc *gob.Encoder)
}

// GobDeserializer represents a Gob deserializable entity.
type GobDeserializer interface {
	FromGob(dec *gob.Decoder)
}

func Encode(enc *gob.Encoder, e interface{}) {
	if err := enc.Encode(e); err != nil {
		log.Warnf("gob encoding error: %v", err)
	}
}

func Decode(dec *gob.Decoder, e interface{}) {
	if err := dec.Decode(e); err != nil {
		log.Warnf("gob decoding error: %v", err)
	}
}
