/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package storage

import "github.com/ortuman/jackal/xmpp"

// vCardStorage defines storage operations for vCards
type vCardStorage interface {
	UpsertVCard(vCard xmpp.XElement, username string) error
	FetchVCard(username string) (xmpp.XElement, error)
}

// UpsertVCard inserts a new vCard element into storage,
// or updates it in case it's been previously inserted.
func UpsertVCard(vCard xmpp.XElement, username string) error {
	return instance().UpsertVCard(vCard, username)
}

// FetchVCard retrieves from storage a vCard element associated
// to a given user.
func FetchVCard(username string) (xmpp.XElement, error) {
	return instance().FetchVCard(username)
}
