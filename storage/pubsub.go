package storage

import (
	pubsubmodel "github.com/ortuman/jackal/model/pubsub"
)

type pubSubStorage interface {
	UpsertPubSubNode(node *pubsubmodel.Node) error
	GetPubSubNode(host, name string) (*pubsubmodel.Node, error)

	UpsertPubSubNodeItem(item *pubsubmodel.Item, host, name string, maxNodeItems int) error
	GetPubSubNodeItems(host, name string) ([]pubsubmodel.Item, error)

	UpsertPubSubNodeAffiliation(affiliation *pubsubmodel.Affiliation, host, name string) error
	GetPubSubNodeAffiliations(host, name string) ([]pubsubmodel.Affiliation, error)
}

func UpsertPubSubNode(node *pubsubmodel.Node) error {
	return inst.UpsertPubSubNode(node)
}

func GetPubSubNode(host, name string) (*pubsubmodel.Node, error) {
	return inst.GetPubSubNode(host, name)
}

func UpsertPubSubNodeItem(item *pubsubmodel.Item, host, name string, maxNodeItems int) error {
	return inst.UpsertPubSubNodeItem(item, host, name, maxNodeItems)
}

func GetPubSubNodeItems(host, name string) ([]pubsubmodel.Item, error) {
	return inst.GetPubSubNodeItems(host, name)
}

func UpsertPubSubNodeAffiliation(affiliatiaon *pubsubmodel.Affiliation, host, name string) error {
	return inst.UpsertPubSubNodeAffiliation(affiliatiaon, host, name)
}

func GetPubSubNodeAffiliations(host, name string) ([]pubsubmodel.Affiliation, error) {
	return inst.GetPubSubNodeAffiliations(host, name)
}
