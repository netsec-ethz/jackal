package memstorage

import (
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
