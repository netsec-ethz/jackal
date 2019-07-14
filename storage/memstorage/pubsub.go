package memstorage

import (
	"errors"

	pubsubmodel "github.com/ortuman/jackal/model/pubsub"
)

func (m *Storage) InsertOrUpdatePubSubNode(node *pubsubmodel.Node) error {
	return m.inWriteLock(func() error {
		m.pubSubNodes[node.Host+"-"+node.Name] = node
		return nil
	})
}

func (m *Storage) GetPubSubNode(host, name string) (node *pubsubmodel.Node, err error) {
	err = m.inReadLock(func() error {
		node = m.pubSubNodes[host+"-"+name]
		return nil
	})
	return
}

func (m *Storage) InsertOrUpdatePubSubNodeItem(item *pubsubmodel.Item, host, name string, maxNodeItems int) error {
	// TODO(ortuman): implement me!
	return errors.New("unimplemented method")
}

func (m *Storage) GetPubSubNodeItems(host, name string) ([]pubsubmodel.Item, error) {
	// TODO(ortuman): implement me!
	return nil, errors.New("unimplemented method")
}

func (m *Storage) InsertOrUpdatePubSubNodeAffiliation(affiliation *pubsubmodel.Affiliation, host, name string) error {
	// TODO(ortuman): implement me!
	return errors.New("unimplemented method")
}

func (m *Storage) GetPubSubNodeAffiliations(host, name string) ([]pubsubmodel.Affiliation, error) {
	// TODO(ortuman): implement me!
	return nil, errors.New("unimplemented method")
}
