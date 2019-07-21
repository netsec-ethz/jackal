package mysql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	pubsubmodel "github.com/ortuman/jackal/model/pubsub"
	"github.com/ortuman/jackal/xmpp"
	"github.com/stretchr/testify/require"
)

func TestMySQLStorageInsertPubSubNode(t *testing.T) {
	s, mock := NewMock()
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO pubsub_nodes (.+) ON DUPLICATE KEY UPDATE (.+)").
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

func TestMySQLStorageGetPubSubNode(t *testing.T) {
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

func TestMySQLStorageGetPubSubNodeError(t *testing.T) {

	s, mock := NewMock()
	mock.ExpectQuery("SELECT name, value FROM pubsub_node_options WHERE (.+)").
		WithArgs("ortuman@jackal.im", "princely_musings").
		WillReturnError(errMySQLStorage)

	_, err := s.GetPubSubNode("ortuman@jackal.im", "princely_musings")

	require.Nil(t, mock.ExpectationsWereMet())

	require.NotNil(t, err)
	require.Equal(t, errMySQLStorage, err)
}

func TestMySQLStorageInsertOrUpdatePubSubNodeItem(t *testing.T) {
	payload := xmpp.NewIQType(uuid.New().String(), xmpp.GetType)

	s, mock := NewMock()

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT node_id, COUNT(.+) FROM pubsub_items WHERE (.+) GROUP BY (node_id)").
		WithArgs("ortuman@jackal.im", "princely_musings").
		WillReturnRows(sqlmock.NewRows([]string{"node_id", "COUNT(*)"}).AddRow("1", "1"))

	mock.ExpectQuery("SELECT item_id FROM pubsub_items WHERE node_id = (.+) AND created_at = (.+)").
		WillReturnRows(sqlmock.NewRows([]string{"item_id"}).AddRow("xyz999"))

	mock.ExpectExec("DELETE FROM pubsub_items WHERE (.+)").
		WithArgs("xyz999").
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec("INSERT INTO pubsub_items (.+) ON DUPLICATE KEY UPDATE payload = (.+), publisher = (.+)").
		WithArgs("1", "abc1234", payload.String(), "ortuman@jackal.im", payload.String(), "ortuman@jackal.im").
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
