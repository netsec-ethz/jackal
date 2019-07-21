package pgsql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	pubsubmodel "github.com/ortuman/jackal/model/pubsub"
	"github.com/ortuman/jackal/xmpp"
	"github.com/stretchr/testify/require"
)

func TestStorageInsertOrUpdatePubSubNode(t *testing.T) {
	s, mock := NewMock()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO pubsub_nodes (.+) ON CONFLICT (.+) DO NOTHING").
		WithArgs("host", "name").
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectQuery("SELECT id FROM pubsub_nodes WHERE (.+)").
		WithArgs("host", "name").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	mock.ExpectExec("DELETE FROM pubsub_node_options WHERE (.+)").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	opts := pubsubmodel.Options{}

	for i := 0; i < len(opts.Map()); i++ {
		mock.ExpectExec("INSERT INTO pubsub_node_options (.+)").
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(0, 1))
	}
	mock.ExpectCommit()

	node := pubsubmodel.Node{Host: "host", Name: "name", Options: opts}
	err := s.InsertOrUpdatePubSubNode(&node)

	require.Nil(t, mock.ExpectationsWereMet())

	require.Nil(t, err)
}

func TestStorageGetPubSubNode(t *testing.T) {
	var cols = []string{"name", "value"}

	s, mock := NewMock()
	rows := sqlmock.NewRows(cols)
	rows.AddRow("pubsub#access_model", "presence")
	rows.AddRow("pubsub#publish_model", "publishers")
	rows.AddRow("pubsub#send_last_published_item", "on_sub_and_presence")

	mock.ExpectQuery("SELECT name, value FROM pubsub_node_options WHERE (.+)").
		WithArgs("ortuman@jackal.im", "princely_musings").
		WillReturnRows(rows)

	node, err := s.GetPubSubNode("ortuman@jackal.im", "princely_musings")

	require.Nil(t, mock.ExpectationsWereMet())

	require.Nil(t, err)
	require.NotNil(t, node)
	require.Equal(t, node.Options.AccessModel, pubsubmodel.Presence)
	require.Equal(t, node.Options.PublishModel, pubsubmodel.Publishers)
	require.Equal(t, node.Options.SendLastPublishedItem, pubsubmodel.OnSubAndPresence)
}

func TestStorageGetPubSubNodeError(t *testing.T) {

	s, mock := NewMock()
	mock.ExpectQuery("SELECT name, value FROM pubsub_node_options WHERE (.+)").
		WithArgs("ortuman@jackal.im", "princely_musings").
		WillReturnError(errGeneric)

	_, err := s.GetPubSubNode("ortuman@jackal.im", "princely_musings")

	require.Nil(t, mock.ExpectationsWereMet())

	require.NotNil(t, err)
	require.Equal(t, errGeneric, err)
}

func TestStorage_InsertOrUpdatePubSubNodeItem(t *testing.T) {
	payload := xmpp.NewIQType(uuid.New().String(), xmpp.GetType)

	s, mock := NewMock()

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id FROM pubsub_nodes WHERE (.+)").
		WithArgs("ortuman@jackal.im", "princely_musings").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	mock.ExpectExec("INSERT INTO pubsub_items (.+) ON CONFLICT (.+) DO UPDATE SET payload = (.+), publisher = (.+)").
		WithArgs("1", "abc1234", payload.String(), "ortuman@jackal.im", payload.String(), "ortuman@jackal.im").
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec("DELETE FROM pubsub_items WHERE item_id IN (.+)").
		WithArgs("1", int64(1)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := s.InsertOrUpdatePubSubNodeItem(&pubsubmodel.Item{
		ID:        "abc1234",
		Publisher: "ortuman@jackal.im",
		Payload:   payload,
	}, "ortuman@jackal.im", "princely_musings", 1)

	require.Nil(t, mock.ExpectationsWereMet())

	require.Nil(t, err)
}
