package storage

import (
	pubsubmodel "github.com/ortuman/jackal/model/pubsub"
)

type pubSubStorage interface {
	InsertOrUpdatePubSubNode(node *pubsubmodel.Node) error
	GetPubSubNode(host, name string) (*pubsubmodel.Node, error)

	InsertPubSubNodeItem(item *pubsubmodel.Item, host, name string, maxNodeItems int) error
	GetPubSubNodeItems(host, name string) ([]pubsubmodel.Item, error)

	InsertPubSubNodeAffiliation(affiliation *pubsubmodel.Affiliation, host, name string) error
	GetPubSubNodeAffiliation(host, name string) ([]pubsubmodel.Affiliation, error)
}

func InsertOrUpdatePubSubNode(node *pubsubmodel.Node) error {
	return inst.InsertOrUpdatePubSubNode(node)
}

func GetPubSubNode(host, name string) (*pubsubmodel.Node, error) {
	return inst.GetPubSubNode(host, name)
}

func InsertNodeItem(item *pubsubmodel.Item, host, name string, maxNodeItems int) error {
	return inst.InsertPubSubNodeItem(item, host, name, maxNodeItems)
}

func GetNodeItems(host, name string) ([]pubsubmodel.Item, error) {
	return inst.GetPubSubNodeItems(host, name)
}

func InsertNodeAffiliation(affiliatiaon *pubsubmodel.Affiliation, host, name string) error {
	return inst.InsertPubSubNodeAffiliation(affiliatiaon, host, name)
}

func GetNodeAffiliation(host, name string) ([]pubsubmodel.Affiliation, error) {
	return inst.GetPubSubNodeAffiliation(host, name)
}
