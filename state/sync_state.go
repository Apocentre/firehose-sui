package syncstate

import (
	"context"

	firecore "github.com/streamingfast/firehose-core"
	"github.com/streamingfast/firehose-core/launcher"
	"go.uber.org/zap"
)

func Setup() {
	// UnsafeResolveReaderNodeStartBlock is a function that resolved the reader node start block num, by default it simply
	// returns the value of the 'reader-node-start-block-num'. However, the function may be overwritten in certain chains
	// to perform a more complex resolution logic.
	firecore.UnsafeResolveReaderNodeStartBlock = func(ctx context.Context, startBlockNum uint64, firstStreamableBlock uint64, runtime *launcher.Runtime, rootLog *zap.Logger) (uint64, error) {
		return 1000, nil
	}

	readerPlugin := launcher.AppRegistry["reader-node"];
	readerPlugin.OnBlockWritten(func(block *bstream.Block) error {

		return nil
	});
}
