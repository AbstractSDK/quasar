package keeper

import (
	"github.com/abag/quasarnode/x/osmolpv/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// @desc Set the list of strategy names in store with prefix key string equals orion_strategies
// Should be called only once. If even if called it should hold unique valued list
// Called from init genesis.
func (k Keeper) SetStrategyNames(ctx sdk.Context, names []string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.StrategyKBP)))
	key := types.CreateStrategyKey()

	var strategies types.Strategies
	b := store.Get(key)
	if b == nil {
		strategies.Names = names
		b := k.cdc.MustMarshal(&strategies)
		store.Set(key, b)
	} else {
		// Make sure only unique names are present
		k.cdc.MustUnmarshal(b, &strategies)
		strategies.Names = getUniqueNames(names, strategies.Names)
		b := k.cdc.MustMarshal(&strategies)
		store.Set(key, b)
	}
}

// @desc Get the list of strategy names from store with prefix key string equals orion_strategies
func (k Keeper) GetStrategyNames(ctx sdk.Context) (strategies types.Strategies, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.StrategyKBP)))
	key := types.CreateStrategyKey()

	b := store.Get(key)
	if b == nil {
		return strategies, false
	}
	k.cdc.MustUnmarshal(b, &strategies)
	return strategies, true
}

// @desc Set the list of sub strategy names in store with prefix key string equals "meissa", rigel" etc.
// Example list : meissa - meissa.7d, meissa.21d, meissa.1m
// Should be called only once. If even if called it should hold unique valued list
// Called from init genesis.
func (k Keeper) SetSubStrategyNames(ctx sdk.Context, sub string, names []string) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.StrategyKBP)))
	var key []byte
	if sub == types.MeissaStrategyName {
		key = types.CreateMeissaStrategyKey()
	} else if sub == types.RigelStrategyName {
		key = types.CreateRigelStrategyKey()
	}

	var strategies types.Strategies
	b := store.Get(key)
	if b == nil {
		strategies.Names = names
		b := k.cdc.MustMarshal(&strategies)
		store.Set(key, b)
	} else {
		// Make sure only unique names are present
		k.cdc.MustUnmarshal(b, &strategies)
		strategies.Names = getUniqueNames(names, strategies.Names)
		b := k.cdc.MustMarshal(&strategies)
		store.Set(key, b)
	}
}

// @desc Get the list of sub strategy names from store with prefix key string equals "meissa", rigel" etc.
// Example list : meissa - meissa.7d, meissa.21d, meissa.1m
func (k Keeper) GetSubStrategyNames(ctx sdk.Context, sub string) (strategies types.Strategies, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.StrategyKBP)))

	var key []byte
	if sub == types.MeissaStrategyName {
		key = types.CreateMeissaStrategyKey()
	} else if sub == types.RigelStrategyName {
		key = types.CreateRigelStrategyKey()
	}

	b := store.Get(key)
	if b == nil {
		return strategies, false
	}
	k.cdc.MustUnmarshal(b, &strategies)
	return strategies, true
}

// @desc Set the strategy current position
func (k Keeper) SetMeissaStrategyCurrPos(ctx sdk.Context, currPos types.CurrentPosition) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.MeissaStrategyKBP)))
	key := types.CreateMeissaPositionKey()
	// TODO - Validate this because types.CurrentPosition has pointer elements
	// inside as fields
	b := k.cdc.MustMarshal(&currPos)
	store.Set(key, b)
}

// @desc Get the strategy current position
func (k Keeper) GetMeissaStrategyCurrPos(ctx sdk.Context) (currPos types.CurrentPosition, found bool) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.MeissaStrategyKBP)))
	key := types.CreateMeissaPositionKey()
	b := store.Get(key)
	if b == nil {
		return currPos, false
	}
	// TODO - Validate this because types.CurrentPosition has pointer elements
	// inside as fields. In the unit test cases.

	k.cdc.MustUnmarshal(b, &currPos)
	return currPos, true
}

// @desc Set the strategy position at a given epochday
func (k Keeper) SetMeissaStrategyHistPos(ctx sdk.Context, currPos types.CurrentPosition, epochday uint64) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.MeissaStrategyKBP)))
	key := types.CreateMeissaEpochPositionKey(epochday)
	// TODO - Validate this because types.CurrentPosition has pointer elements
	// inside as fields. In the unit test cases.
	b := k.cdc.MustMarshal(&currPos)
	store.Set(key, b)
}

// @desc Get the strategy current position
func (k Keeper) GetMeissaStrategyHistPos(ctx sdk.Context, epochday uint64) (currPos types.CurrentPosition, found bool) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.MeissaStrategyKBP)))
	key := types.CreateMeissaEpochPositionKey(epochday)

	b := store.Get(key)
	if b == nil {
		return currPos, false
	}
	// TODO - Validate this because types.CurrentPosition has pointer elements
	// inside as fields

	k.cdc.MustUnmarshal(b, &currPos)
	return currPos, true
}

// @desc Set the strategy current position
func (k Keeper) SetMeissaStrategyTotalHistPos(ctx sdk.Context, histPos types.TotalHistPosition, epochday uint64) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.MeissaStrategyKBP)))
	key := types.CreateMeissaEpochPositionKey(epochday)
	// TODO - Validate this because types.CurrentPosition has pointer elements
	// inside as fields. In the unit test cases.
	b := k.cdc.MustMarshal(&histPos)
	store.Set(key, b)
}

// @desc Get the strategy current position
func (k Keeper) GetMeissaStrategyTotalHisPos(ctx sdk.Context, epochday uint64) (histPos types.TotalHistPosition, found bool) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(string(types.MeissaStrategyKBP)))
	key := types.CreateMeissaEpochPositionKey(epochday)

	b := store.Get(key)
	if b == nil {
		return histPos, false
	}
	// TODO - Validate this because types.CurrentPosition has pointer elements
	// inside as fields

	k.cdc.MustUnmarshal(b, &histPos)
	return histPos, true
}

func getUniqueNames(names, existingNames []string) []string {
	var uniqueNames []string
	uniqueNames = existingNames
	for _, name := range names {
		found := false
		for _, existingName := range existingNames {
			if name == existingName {
				found = true
				break
			}
		}
		if !found {
			uniqueNames = append(uniqueNames, name)
		}
	}
	return uniqueNames
}
