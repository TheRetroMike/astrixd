package blockwindowheapslicestore

import (
	"github.com/astrix-network/astrixd/domain/consensus/model"
	"github.com/astrix-network/astrixd/domain/consensus/model/externalapi"
	"github.com/astrix-network/astrixd/domain/consensus/utils/lrucachehashandwindowsizetoblockghostdagdatahashpairs"
	"github.com/astrix-network/astrixd/infrastructure/db/database"
	"github.com/astrix-network/astrixd/util/staging"
	"github.com/pkg/errors"
)

type blockWindowHeapSliceStore struct {
	shardID model.StagingShardID
	cache   *lrucachehashandwindowsizetoblockghostdagdatahashpairs.LRUCache
}

// New instantiates a new WindowHeapSliceStore
func New(cacheSize int, preallocate bool) model.WindowHeapSliceStore {
	return &blockWindowHeapSliceStore{
		shardID: staging.GenerateShardingID(),
		cache:   lrucachehashandwindowsizetoblockghostdagdatahashpairs.New(cacheSize, preallocate),
	}
}

// Stage stages the given blockStatus for the given blockHash
func (bss *blockWindowHeapSliceStore) Stage(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash, windowSize int, heapSlice []*externalapi.BlockGHOSTDAGDataHashPair) {
	stagingShard := bss.stagingShard(stagingArea)
	stagingShard.toAdd[newShardKey(blockHash, windowSize)] = heapSlice
}

func (bss *blockWindowHeapSliceStore) IsStaged(stagingArea *model.StagingArea) bool {
	return bss.stagingShard(stagingArea).isStaged()
}

func (bss *blockWindowHeapSliceStore) Get(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash, windowSize int) ([]*externalapi.BlockGHOSTDAGDataHashPair, error) {
	stagingShard := bss.stagingShard(stagingArea)

	if heapSlice, ok := stagingShard.toAdd[newShardKey(blockHash, windowSize)]; ok {
		return heapSlice, nil
	}

	if heapSlice, ok := bss.cache.Get(blockHash, windowSize); ok {
		return heapSlice, nil
	}

	return nil, errors.Wrap(database.ErrNotFound, "Window heap slice not found")
}
