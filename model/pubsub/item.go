/*
 * Copyright (c) 2019 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package pubsubmodel

import "github.com/ortuman/jackal/xmpp"

type Item struct {
	ID        string
	Publisher string
	Payload   xmpp.XElement
}
