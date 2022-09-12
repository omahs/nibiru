package collections

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NibiruChain/nibiru/collections/keys"
)

func NewMap[K keys.Key, V any, PV interface {
	*V
	Object
}](cdc codec.BinaryCodec, sk sdk.StoreKey, prefix uint8) Map[K, V, PV] {
	return Map[K, V, PV]{
		cdc:    cdc,
		sk:     sk,
		prefix: []byte{prefix},
	}
}

// Map defines a collection which simply does mappings between primary keys and objects.
type Map[K keys.Key, V any, PV interface {
	*V
	Object
}] struct {
	cdc    codec.BinaryCodec
	sk     sdk.StoreKey
	prefix []byte
	_      K
	_      V
}

func (m Map[K, V, PV]) getStore(ctx sdk.Context) sdk.KVStore {
	return prefix.NewStore(ctx.KVStore(m.sk), m.prefix)
}

func (m Map[K, V, PV]) Insert(ctx sdk.Context, key K, object V) {
	store := m.getStore(ctx)
	store.Set(key.PrimaryKey(), m.cdc.MustMarshal(PV(&object)))
}

func (m Map[K, V, PV]) Get(ctx sdk.Context, key K) (V, error) {
	store := m.getStore(ctx)
	pk := key.PrimaryKey()
	bytes := store.Get(pk)
	if bytes == nil {
		var x V
		return x, ErrNotFound
	}

	x := new(V)
	m.cdc.MustUnmarshal(bytes, PV(x))
	return *x, nil
}

func (m Map[K, V, PV]) GetOr(ctx sdk.Context, key K, def V) V {
	got, err := m.Get(ctx, key)
	if err != nil {
		return def
	}

	return got
}

func (m Map[K, V, PV]) Delete(ctx sdk.Context, key K) error {
	store := m.getStore(ctx)
	pk := key.PrimaryKey()
	if !store.Has(pk) {
		return ErrNotFound
	}

	store.Delete(pk)
	return nil
}

func (m Map[K, V, PV]) Iterate(ctx sdk.Context, start keys.Bound[K], end keys.Bound[K], order keys.Order) MapIterator[K, V, PV] {
	store := m.getStore(ctx)
	return newMapIterator[K, V, PV](m.cdc, store, start, end, order)
}

func (m Map[K, V, PV]) GetAll(ctx sdk.Context) []V {
	iter := m.Iterate(ctx, keys.None[K](), keys.None[K](), keys.OrderAscending)
	defer iter.Close()

	var list []V
	for ; iter.Valid(); iter.Next() {
		list = append(list, iter.Value())
	}

	return list
}

func newMapIterator[K keys.Key, V any, PV interface {
	*V
	Object
}](cdc codec.BinaryCodec, store sdk.KVStore, start, end keys.Bound[K], order keys.Order) MapIterator[K, V, PV] {
	startBytes := start.Bytes()
	endBytes := end.Bytes()
	switch order {
	case keys.OrderAscending:
		return MapIterator[K, V, PV]{
			cdc:  cdc,
			iter: store.Iterator(startBytes, endBytes),
		}
	case keys.OrderDescending:
		return MapIterator[K, V, PV]{
			cdc:  cdc,
			iter: store.ReverseIterator(startBytes, endBytes),
		}
	default:
		panic(fmt.Errorf("unrecognized order"))
	}
}

type MapIterator[K keys.Key, V any, PV interface {
	*V
	Object
}] struct {
	cdc  codec.BinaryCodec
	iter sdk.Iterator
}

func (i MapIterator[K, V, PV]) Close() {
	_ = i.iter.Close()
}

func (i MapIterator[K, V, PV]) Next() {
	i.iter.Next()
}

func (i MapIterator[K, V, PV]) Valid() bool {
	return i.iter.Valid()
}

func (i MapIterator[K, V, PV]) Value() V {
	x := PV(new(V))
	i.cdc.MustUnmarshal(i.iter.Value(), x)
	return *x
}

func (i MapIterator[K, V, PV]) Key() K {
	var k K
	return k.FromPrimaryKeyBytes(i.iter.Key()).(K) // TODO implement this better
}

func (i MapIterator[K, V, PV]) Values() []V {
	var values []V
	for ; i.iter.Valid(); i.iter.Next() {
		values = append(values, i.Value())
	}
	return values
}

func (i MapIterator[K, V, PV]) Keys() []K {
	var keys []K
	for ; i.iter.Valid(); i.iter.Next() {
		keys = append(keys, i.Key())
	}
	return keys
}

type KeyValue[K keys.Key, V any, PV interface {
	*V
	Object
}] struct {
	Key   K
	Value V
}

func (i MapIterator[K, V, PV]) All() []KeyValue[K, V, PV] {
	var kvs []KeyValue[K, V, PV]
	for ; i.iter.Valid(); i.iter.Next() {
		kvs = append(kvs, KeyValue[K, V, PV]{
			Key:   i.Key(),
			Value: i.Value(),
		})
	}

	return kvs
}