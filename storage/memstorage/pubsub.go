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
