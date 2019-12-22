/*
 * Copyright (c) 2018 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package pgsql

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/ortuman/jackal/model"
	"github.com/stretchr/testify/require"
)

// Insert a valid block list item
func TestInsertValidBlockListItem(t *testing.T) {
	s, mock := NewMock()

	mock.ExpectExec("INSERT INTO blocklist_items (.+)").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := s.InsertBlockListItem(&model.BlockListItem{Username: "ortuman", JID: "noelia@jackal.im"})
	require.Nil(t, err)
	require.Nil(t, mock.ExpectationsWereMet())
}

// Insert the same row twice to test for key uniqueness validation
func TestInsertDoubleBlockListItem(t *testing.T) {
	s, mock := NewMock()

	// First insertion will be successful
	mock.ExpectExec("INSERT INTO blocklist_items (.+)").
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Second insertion will fail
	mock.ExpectExec("INSERT INTO blocklist_items (.+)").
		WillReturnError(errGeneric)

	err := s.InsertBlockListItem(&model.BlockListItem{Username: "ortuman", JID: "noelia@jackal.im"})
	require.Nil(t, err)

	err = s.InsertBlockListItem(&model.BlockListItem{Username: "ortuman", JID: "noelia@jackal.im"})
	require.Equal(t, errGeneric, err)
	require.Nil(t, mock.ExpectationsWereMet())
}

// Test fetching block list items
func TestFetchBlockListItems(t *testing.T) {
	var blockListColumns = []string{"username", "jid"}
	s, mock := NewMock()

	mock.ExpectQuery("SELECT (.+) FROM blocklist_items (.+)").
		WithArgs("ortuman").
		WillReturnRows(sqlmock.NewRows(blockListColumns).AddRow("ortuman", "noelia@jackal.im"))

	_, err := s.FetchBlockListItems("ortuman")
	require.Nil(t, err)
	require.Nil(t, mock.ExpectationsWereMet())
}

// Test error handling on fetching block list items
func TestFetchBlockListItemsError(t *testing.T) {
	s, mock := NewMock()

	mock.ExpectQuery("SELECT (.+) FROM blocklist_items (.+)").
		WithArgs("ortuman").
		WillReturnError(errGeneric)

	_, err := s.FetchBlockListItems("ortuman")
	require.Equal(t, errGeneric, err)
	require.Nil(t, mock.ExpectationsWereMet())
}

// Test deleting an item from the block list
func TestDeleteBlockListItems(t *testing.T) {
	s, mock := NewMock()

	mock.ExpectExec("DELETE FROM blocklist_items (.+)").
		WithArgs("ortuman", "noelia@jackal.im").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := s.DeleteBlockListItem(&model.BlockListItem{Username: "ortuman", JID: "noelia@jackal.im"})
	require.Nil(t, err)
	require.Nil(t, mock.ExpectationsWereMet())
}

// Test error handling on deleting a row from the block list
func TestDeleteBlockListItemsError(t *testing.T) {
	s, mock := NewMock()

	mock.ExpectExec("DELETE FROM blocklist_items (.+)").
		WithArgs("ortuman", "noelia@jackal.im").
		WillReturnError(errGeneric)

	err := s.DeleteBlockListItem(&model.BlockListItem{Username: "ortuman", JID: "noelia@jackal.im"})
	require.Equal(t, errGeneric, err)
	require.Nil(t, mock.ExpectationsWereMet())
}
