package mysql

import (
	"database/sql"
	"strings"

	sq "github.com/Masterminds/squirrel"
	pubsubmodel "github.com/ortuman/jackal/model/pubsub"
	"github.com/ortuman/jackal/xmpp"
)

func (s *Storage) InsertOrUpdatePubSubNode(node *pubsubmodel.Node) error {
	return s.inTransaction(func(tx *sql.Tx) error {
		// if not existing, insert new node
		_, err := sq.Insert("pubsub_nodes").
			Columns("host", "name", "updated_at", "created_at").
			Suffix("ON DUPLICATE KEY UPDATE updated_at = NOW()").
			Values(node.Host, node.Name, nowExpr, nowExpr).
			RunWith(tx).Exec()
		if err != nil {
			return err
		}

		// fetch node identifier
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
				Columns("node_id", "name", "value", "updated_at", "created_at").
				Values(nodeIdentifier, name, value, nowExpr, nowExpr).
				RunWith(tx).Exec()
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *Storage) GetPubSubNode(host, name string) (*pubsubmodel.Node, error) {
	rows, err := sq.Select("name", "value").
		From("pubsub_node_options").
		Where("node_id = (SELECT id FROM pubsub_nodes WHERE host = ? AND name = ?)", host, name).
		RunWith(s.db).Query()
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var optMap = make(map[string]string)
	for rows.Next() {
		var opt, value string
		if err := rows.Scan(&opt, &value); err != nil {
			return nil, err
		}
		optMap[opt] = value
	}
	opts, err := pubsubmodel.NewOptionsFromMap(optMap)
	if err != nil {
		return nil, err
	}
	return &pubsubmodel.Node{
		Host:    host,
		Name:    name,
		Options: *opts,
	}, nil
}

func (s *Storage) InsertOrUpdatePubSubNodeItem(item *pubsubmodel.Item, host, name string, maxNodeItems int) error {
	return s.inTransaction(func(tx *sql.Tx) error {
		// fetch node identifier and item count
		var nodeIdentifier string
		var itemCount int

		err := sq.Select("node_id", "COUNT(*)").
			From("pubsub_items").
			Where("node_id = (SELECT id FROM pubsub_nodes WHERE host = ? AND name = ?)", host, name).
			GroupBy("node_id").
			RunWith(tx).QueryRow().Scan(&nodeIdentifier, &itemCount)
		switch err {
		case nil:
			break
		case sql.ErrNoRows:
			return nil
		default:
			return err
		}

		// check if maximum item count was reached
		if itemCount == maxNodeItems {
			// delete oldest item id
			var oldestItemID string

			err := sq.Select("item_id").
				From("pubsub_items").
				Where("node_id = ? AND created_at = (SELECT MIN(created_at) FROM pubsub_items WHERE node_id = ?)", nodeIdentifier, nodeIdentifier).
				RunWith(tx).QueryRow().Scan(&oldestItemID)
			if err != nil {
				return err
			}
			_, err = sq.Delete("pubsub_items").
				Where(sq.Eq{"item_id": oldestItemID}).
				Exec()
			if err != nil {
				return err
			}
		}

		// upsert new item
		rawPayload := item.Payload.String()

		_, err = sq.Insert("pubsub_items").
			Columns("node_id", "item_id", "payload", "publisher", "updated_at", "created_at").
			Values(nodeIdentifier, item.ID, rawPayload, item.Publisher, nowExpr, nowExpr).
			Suffix("ON DUPLICATE KEY UPDATE payload = ?, publisher = ?, updated_at = NOW()", rawPayload, item.Publisher).
			RunWith(s.db).Exec()
		return err
	})
}

func (s *Storage) GetPubSubNodeItems(host, name string) ([]pubsubmodel.Item, error) {
	rows, err := sq.Select("item_id", "publisher", "payload").
		From("pubsub_items").
		Where("node_id = (SELECT id FROM pubsub_nodes WHERE host = ? AND name = ?)", host, name).
		RunWith(s.db).Query()
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var items []pubsubmodel.Item
	for rows.Next() {
		var payload string
		var item pubsubmodel.Item
		if err := rows.Scan(&item.ID, &item.Publisher, &payload); err != nil {
			return nil, err
		}
		parser := xmpp.NewParser(strings.NewReader(payload), xmpp.DefaultMode, 0)
		item.Payload, err = parser.ParseElement()
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (s *Storage) InsertOrUpdatePubSubNodeAffiliation(affiliation *pubsubmodel.Affiliation, host, name string) error {
	return s.inTransaction(func(tx *sql.Tx) error {

		// fetch node identifier
		var nodeIdentifier string

		err := sq.Select("id").
			From("pubsub_nodes").
			Where(sq.And{sq.Eq{"host": host}, sq.Eq{"name": name}}).
			RunWith(tx).QueryRow().Scan(&nodeIdentifier)
		switch err {
		case nil:
			break
		case sql.ErrNoRows:
			return nil
		default:
			return err

		}

		// insert affiliation
		_, err = sq.Insert("pubsub_affiliations").
			Columns("node_id", "jid", "affiliation", "updated_at", "created_at").
			Values(nodeIdentifier, affiliation.JID, affiliation.Affiliation, nowExpr, nowExpr).
			Suffix("ON DUPLICATE KEY UPDATE affiliation = ?, updated_at = NOW()", affiliation.Affiliation).
			RunWith(s.db).Exec()
		return err
	})
}

func (s *Storage) GetPubSubNodeAffiliations(host, name string) ([]pubsubmodel.Affiliation, error) {
	rows, err := sq.Select("jid", "affiliation").
		From("pubsub_affiliations").
		Where("node_id = (SELECT id FROM pubsub_nodes WHERE host = ? AND name = ?)", host, name).
		RunWith(s.db).Query()
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var affiliations []pubsubmodel.Affiliation
	for rows.Next() {
		var affiliation pubsubmodel.Affiliation
		if err := rows.Scan(&affiliation.JID, &affiliation.Affiliation); err != nil {
			return nil, err
		}
		affiliations = append(affiliations, affiliation)
	}
	return affiliations, nil
}
