/*
 * Copyright (c) 2019 Miguel Ángel Ortuño.
 * See the LICENSE file for more information.
 */

package badgerdb

import (
	"context"
	"strings"

	"github.com/dgraph-io/badger"
	pubsubmodel "github.com/ortuman/jackal/model/pubsub"
	"github.com/ortuman/jackal/model/serializer"
)

type badgerDBPubSub struct {
	*badgerDBStorage
}

func newPubSub(db *badger.DB) *badgerDBPubSub {
	return &badgerDBPubSub{badgerDBStorage: newStorage(db)}
}

func (b *badgerDBPubSub) FetchHosts(_ context.Context) ([]string, error) {
	var hosts []string

	err := b.db.View(func(txn *badger.Txn) error {
		return b.forEachKey([]byte("pubSubNodes:"), func(k []byte) error {
			key := string(k)
			keySplits := strings.Split(key, ":")
			if len(keySplits) != 3 {
				return nil
			}
			host := keySplits[1]

			var isPresent bool
			for _, h := range hosts {
				if h == host {
					isPresent = true
					break
				}
			}
			if isPresent {
				return nil // nothing to do here
			}
			hosts = append(hosts, host)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func (b *badgerDBPubSub) UpsertNode(_ context.Context, node *pubsubmodel.Node) error {
	return b.db.Update(func(tx *badger.Txn) error {
		return b.upsert(node, pubSubNodesKey(node.Host, node.Name), tx)
	})
}

func (b *badgerDBPubSub) FetchNode(_ context.Context, host, name string) (*pubsubmodel.Node, error) {
	var node pubsubmodel.Node
	err := b.db.View(func(txn *badger.Txn) error {
		return b.fetch(&node, pubSubNodesKey(host, name), txn)
	})
	switch err {
	case nil:
		return &node, nil
	case errEntityNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

func (b *badgerDBPubSub) FetchNodes(_ context.Context, host string) ([]pubsubmodel.Node, error) {
	var nodes []pubsubmodel.Node

	err := b.db.View(func(txn *badger.Txn) error {
		return b.forEachKey([]byte("pubSubNodes:"+host), func(k []byte) error {
			bs, err := b.getVal(k, txn)
			if err != nil {
				return err
			}
			var node pubsubmodel.Node
			if err := serializer.Deserialize(bs, &node); err != nil {
				return err
			}
			nodes = append(nodes, node)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (b *badgerDBPubSub) FetchSubscribedNodes(_ context.Context, jid string) ([]pubsubmodel.Node, error) {
	var nodes []pubsubmodel.Node

	err := b.db.View(func(txn *badger.Txn) error {
		return b.forEachKey([]byte("pubSubSubscriptions:"), func(k []byte) error {
			bs, err := b.getVal(k, txn)
			if err != nil {
				return err
			}
			keySplits := strings.Split(string(k), ":")
			if len(keySplits) != 3 {
				return nil // wrong key format
			}
			host := keySplits[1]
			name := keySplits[2]

			var subs []pubsubmodel.Subscription
			if err := serializer.DeserializeSlice(bs, &subs); err != nil {
				return err
			}
			for _, sub := range subs {
				if sub.JID != jid || sub.Subscription != pubsubmodel.Subscribed {
					continue
				}
				// fetch pubsub node
				var node pubsubmodel.Node

				b, err := b.getVal(pubSubNodesKey(host, name), txn)
				if err != nil {
					return err
				}
				if b == nil {
					continue
				}
				if err := serializer.Deserialize(b, &node); err != nil {
					return err
				}
				nodes = append(nodes, node)
				break
			}
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (b *badgerDBPubSub) DeleteNode(_ context.Context, host, name string) error {
	return b.db.Update(func(tx *badger.Txn) error {
		if err := b.delete(pubSubNodesKey(host, name), tx); err != nil {
			return err
		}
		if err := b.delete(pubSubItemsKey(host, name), tx); err != nil {
			return err
		}
		return b.delete(pubSubAffiliationsKey(host, name), tx)
	})
}

func (b *badgerDBPubSub) UpsertNodeItem(_ context.Context, item *pubsubmodel.Item, host, name string, maxNodeItems int) error {
	return b.db.Update(func(tx *badger.Txn) error {
		var items []pubsubmodel.Item
		if err := b.fetchSlice(&items, pubSubItemsKey(host, name), tx); err != nil {
			return err
		}
		var updated bool
		for i, itm := range items {
			if itm.ID == item.ID {
				items[i] = *item
				updated = true
				break
			}
		}
		if !updated {
			items = append(items, *item)
		}
		if len(items) > maxNodeItems {
			items = items[len(items)-maxNodeItems:] // remove oldest elements
		}
		return b.upsertSlice(&items, pubSubItemsKey(host, name), tx)
	})
}

func (b *badgerDBPubSub) FetchNodeItems(_ context.Context, host, name string) ([]pubsubmodel.Item, error) {
	var items []pubsubmodel.Item
	err := b.db.View(func(txn *badger.Txn) error {
		return b.fetchSlice(&items, pubSubItemsKey(host, name), txn)
	})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (b *badgerDBPubSub) FetchNodeItemsWithIDs(_ context.Context, host, name string, identifiers []string) ([]pubsubmodel.Item, error) {
	var items []pubsubmodel.Item
	err := b.db.View(func(txn *badger.Txn) error {
		return b.fetchSlice(&items, pubSubItemsKey(host, name), txn)
	})
	if err != nil {
		return nil, err
	}
	identifiersSet := make(map[string]struct{})
	for _, id := range identifiers {
		identifiersSet[id] = struct{}{}
	}
	var filteredItems []pubsubmodel.Item
	for _, itm := range items {
		if _, ok := identifiersSet[itm.ID]; ok {
			filteredItems = append(filteredItems, itm)
		}
	}
	return filteredItems, nil
}

func (b *badgerDBPubSub) FetchNodeLastItem(_ context.Context, host, name string) (*pubsubmodel.Item, error) {
	var items []pubsubmodel.Item
	err := b.db.View(func(txn *badger.Txn) error {
		return b.fetchSlice(&items, pubSubItemsKey(host, name), txn)
	})
	if err != nil {
		return nil, err
	}
	return &items[len(items)-1], nil
}

func (b *badgerDBPubSub) UpsertNodeAffiliation(_ context.Context, affiliation *pubsubmodel.Affiliation, host, name string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		var affiliations []pubsubmodel.Affiliation
		if err := b.fetchSlice(&affiliations, pubSubAffiliationsKey(host, name), txn); err != nil {
			return err
		}
		var updated bool
		for i, aff := range affiliations {
			if aff.JID == affiliation.JID {
				affiliations[i] = *affiliation
				updated = true
				break
			}
		}
		if !updated {
			affiliations = append(affiliations, *affiliation)
		}
		return b.upsertSlice(&affiliations, pubSubAffiliationsKey(host, name), txn)
	})
}

func (b *badgerDBPubSub) FetchNodeAffiliation(ctx context.Context, host, name, jid string) (*pubsubmodel.Affiliation, error) {
	affiliations, err := b.FetchNodeAffiliations(ctx, host, name)
	if err != nil {
		return nil, err
	}
	for _, aff := range affiliations {
		if aff.JID == jid {
			return &aff, nil
		}
	}
	return nil, nil
}

func (b *badgerDBPubSub) FetchNodeAffiliations(_ context.Context, host, name string) ([]pubsubmodel.Affiliation, error) {
	var affiliations []pubsubmodel.Affiliation
	err := b.db.View(func(txn *badger.Txn) error {
		return b.fetchSlice(&affiliations, pubSubAffiliationsKey(host, name), txn)
	})
	if err != nil {
		return nil, err
	}
	return affiliations, nil
}

func (b *badgerDBPubSub) DeleteNodeAffiliation(_ context.Context, jid, host, name string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		var affiliations []pubsubmodel.Affiliation
		if err := b.fetchSlice(&affiliations, pubSubAffiliationsKey(host, name), txn); err != nil {
			return err
		}
		var deleted bool
		for i, aff := range affiliations {
			if aff.JID == jid {
				affiliations = append(affiliations[:i], affiliations[i+1:]...)
				deleted = true
				break
			}
		}
		if !deleted {
			return nil
		}
		return b.upsertSlice(&affiliations, pubSubAffiliationsKey(host, name), txn)
	})
}

func (b *badgerDBPubSub) UpsertNodeSubscription(_ context.Context, subscription *pubsubmodel.Subscription, host, name string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		var subscriptions []pubsubmodel.Subscription
		if err := b.fetchSlice(&subscriptions, pubSubSubscriptionsKey(host, name), txn); err != nil {
			return err
		}
		var updated bool
		for i, sub := range subscriptions {
			if sub.JID == subscription.JID {
				subscriptions[i] = *subscription
				updated = true
				break
			}
		}
		if !updated {
			subscriptions = append(subscriptions, *subscription)
		}
		return b.upsertSlice(&subscriptions, pubSubSubscriptionsKey(host, name), txn)
	})
}

func (b *badgerDBPubSub) FetchNodeSubscriptions(_ context.Context, host, name string) ([]pubsubmodel.Subscription, error) {
	var subscriptions []pubsubmodel.Subscription
	err := b.db.View(func(txn *badger.Txn) error {
		return b.fetchSlice(&subscriptions, pubSubSubscriptionsKey(host, name), txn)
	})
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (b *badgerDBPubSub) DeleteNodeSubscription(_ context.Context, jid, host, name string) error {
	return b.db.Update(func(txn *badger.Txn) error {
		var subscriptions []pubsubmodel.Subscription
		if err := b.fetchSlice(&subscriptions, pubSubSubscriptionsKey(host, name), txn); err != nil {
			return err
		}
		var deleted bool
		for i, sub := range subscriptions {
			if sub.JID == jid {
				subscriptions = append(subscriptions[:i], subscriptions[i+1:]...)
				deleted = true
				break
			}
		}
		if !deleted {
			return nil
		}
		return b.upsertSlice(&subscriptions, pubSubSubscriptionsKey(host, name), txn)
	})
}

func pubSubNodesKey(host, name string) []byte {
	return []byte("pubSubNodes:" + host + ":" + name)
}

func pubSubAffiliationsKey(host, name string) []byte {
	return []byte("pubSubAffiliations:" + host + ":" + name)
}

func pubSubSubscriptionsKey(host, name string) []byte {
	return []byte("pubSubSubscriptions:" + host + ":" + name)
}

func pubSubItemsKey(host, name string) []byte {
	return []byte("pubSubItems:" + host + ":" + name)
}
