package pgsql

import (
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	pubsubmodel "github.com/ortuman/jackal/model/pubsub"
)

func (s *Storage) InsertOrUpdatePubSubNode(node *pubsubmodel.Node) error {
	return s.inTransaction(func(tx *sql.Tx) error {
		// if not existing, insert new node
		_, err := sq.Insert("pubsub_nodes").
			Columns("host", "name", "updated_at", "created_at").
			Suffix("ON CONFLICT (host, name) DO NOTHING").
			Values(node.Host, node.Name, nowExpr, nowExpr).
			RunWith(tx).Exec()
		if err != nil {
			return err
		}

		// fetch identifier
		var nodeIdentifier string

		err = sq.Select("id").
			From("pubsub_nodes").
			Where(sq.And{sq.Eq{"host": node.Host}, sq.Eq{"name": node.Name}}).
			RunWith(tx).QueryRow().Scan(&nodeIdentifier)
		if err != nil {
			return err
		}

		// delete previous node options
		_, err = sq.Delete("pubsub_node_options").
			Where(sq.Eq{"node_id": nodeIdentifier}).
			RunWith(tx).Exec()
		if err != nil {
			return err
		}
		// insert new option set
		for name, value := range node.Options.Map() {
			_, err = sq.Insert("pubsub_node_options").
				Columns("node_id", "name", "value").
				Values(nodeIdentifier, name, value).
				RunWith(tx).Exec()
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *Storage) GetPubSubNode(host, name string) (*pubsubmodel.Node, error) {
	return nil, errors.New("unimplemented method")
}
