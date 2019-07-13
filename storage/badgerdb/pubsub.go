package badgerdb

import (
	"errors"

	"github.com/dgraph-io/badger"
	pubsubmodel "github.com/ortuman/jackal/model/pubsub"
)

func (b *Storage) InsertOrUpdatePubSubNode(node *pubsubmodel.Node) error {
	return b.db.Update(func(tx *badger.Txn) error {
		return b.insertOrUpdate(node, b.pubSubStorageKey(node.Host, node.Name), tx)
	})
}

func (b *Storage) GetPubSubNode(host, name string) (*pubsubmodel.Node, error) {
	var node pubsubmodel.Node
	err := b.fetch(&node, b.pubSubStorageKey(host, name))
	switch err {
	case nil:
		return &node, nil
	case errBadgerDBEntityNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (b *Storage) InsertPubSubNodeItem(item *pubsubmodel.Item, host, name string, maxNodeItems int) error {
	// TODO(ortuman): implement me!
	return errors.New("unimplemented method")
}

func (b *Storage) GetPubSubNodeItems(host, name string) ([]pubsubmodel.Item, error) {
	// TODO(ortuman): implement me!
	return nil, errors.New("unimplemented method")
}

func (b *Storage) InsertPubSubNodeAffiliation(affiliatiaon *pubsubmodel.Affiliation, host, name string) error {
	// TODO(ortuman): implement me!
	return errors.New("unimplemented method")
}

func (b *Storage) GetPubSubNodeAffiliation(host, name string) ([]pubsubmodel.Affiliation, error) {
	// TODO(ortuman): implement me!
	return nil, errors.New("unimplemented method")
}

func (b *Storage) pubSubStorageKey(host, name string) []byte {
	return []byte("pubsub:" + host + ":" + name)
}
