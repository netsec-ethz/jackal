package storage

import (
	pubsubmodel "github.com/ortuman/jackal/model/pubsub"
)

type pubSubStorage interface {
	InsertOrUpdatePubSubNode(node *pubsubmodel.Node) error
	GetPubSubNode(host, name string) (*pubsubmodel.Node, error)

	InsertOrUpdatePubSubNodeItem(item *pubsubmodel.Item, host, name string, maxNodeItems int) error
	GetPubSubNodeItems(host, name string) ([]pubsubmodel.Item, error)

	InsertOrUpdatePubSubNodeAffiliation(affiliation *pubsubmodel.Affiliation, host, name string) error
	GetPubSubNodeAffiliations(host, name string) ([]pubsubmodel.Affiliation, error)
}

func InsertOrUpdatePubSubNode(node *pubsubmodel.Node) error {
	return inst.InsertOrUpdatePubSubNode(node)
}

func GetPubSubNode(host, name string) (*pubsubmodel.Node, error) {
	return inst.GetPubSubNode(host, name)
}

func InsertOrUpdatePubSubNodeItem(item *pubsubmodel.Item, host, name string, maxNodeItems int) error {
	return inst.InsertOrUpdatePubSubNodeItem(item, host, name, maxNodeItems)
}

func GetPubSubNodeItems(host, name string) ([]pubsubmodel.Item, error) {
	return inst.GetPubSubNodeItems(host, name)
}

func InsertOrUpdatePubSubNodeAffiliation(affiliatiaon *pubsubmodel.Affiliation, host, name string) error {
	return inst.InsertOrUpdatePubSubNodeAffiliation(affiliatiaon, host, name)
}

func GetPubSubNodeAffiliations(host, name string) ([]pubsubmodel.Affiliation, error) {
	return inst.GetPubSubNodeAffiliations(host, name)
}
