package pubsubmodel

import "github.com/ortuman/jackal/xmpp"

type Item struct {
	ID        string
	Publisher string
	Payload   xmpp.Element
}
