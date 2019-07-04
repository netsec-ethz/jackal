package pgsql

import (
	"errors"

	pubsubmodel "github.com/ortuman/jackal/model/pubsub"
)

func (s *Storage) InsertOrUpdatePubSubNode(node *pubsubmodel.Node) error {
	return errors.New("unimplemented method")
}

func (s *Storage) GetPubSubNode(host, name string) (*pubsubmodel.Node, error) {
	return nil, errors.New("unimplemented method")
}
