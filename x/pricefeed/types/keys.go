package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "pricefeed"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_pricefeed"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	// CurrentPricePrefix prefix for the current price of an asset
	CurrentPricePrefix = []byte{0x00}

	// RawPricesNamespace is the bytes prefix for RawPrice objects state
	RawPricesNamespace = []byte{0x01}
	// RawPricesObjectsPrefix is the KV prefix in which RawPrices objects are stored.
	RawPricesObjectsPrefix = append(RawPricesNamespace, 0x00)

	// Snapshot prefix for the median oracle price at a specific point in time
	PriceSnapshotPrefix = []byte{0x03}
)

// CurrentPriceKey returns the prefix for the current price
func CurrentPriceKey(pairID string) []byte {
	return append(CurrentPricePrefix, []byte(pairID)...)
}

func PriceSnapshotKey(pairId string, blockHeight int64) []byte {
	return append(
		PriceSnapshotPrefix,
		append(
			[]byte(pairId),
			sdk.Uint64ToBigEndian(uint64(blockHeight))...,
		)...,
	)
}

// lengthPrefixWithByte returns the input bytes prefixes with one byte containing its length.
// It panics if the input is greater than 255 in length.
func lengthPrefixWithByte(bz []byte) []byte {
	length := len(bz)

	if length > 255 {
		panic("cannot length prefix more than 255 bytes with single byte")
	}

	return append([]byte{byte(length)}, bz...)
}
