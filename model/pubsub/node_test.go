/*
 * Copyright (c) 2019 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package pubsubmodel

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNode_Serialization(t *testing.T) {
	n := Node{}
	n.Name = "playing_lists"
	n.Host = "jackal.im"

	n.Options.Title = "Playing lists"
	n.Options.NotifySub = true
	n.Options.MaxPayloadSize = 1024

	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	n.ToGob(enc)

	n2 := Node{}
	dec := gob.NewDecoder(buf)
	_ = n2.FromGob(dec)

	require.True(t, reflect.DeepEqual(&n, &n2))
}
