package pbsui

import (
	"time"

	"github.com/streamingfast/bstream"
	firecore "github.com/streamingfast/firehose-core"
)

var _ firecore.Block = (*CheckpointData)(nil)

func (b *CheckpointData) GetFirehoseBlockID() string {
	return b.Checkpoint.Digest
}

func (b *CheckpointData) GetFirehoseBlockNumber() uint64 {
	return b.Checkpoint.SequenceNumber
}

func (b *CheckpointData) GetFirehoseBlockParentID() string {
	if b.Checkpoint.PreviousDigest == nil {
		return ""
	}

	return *b.Checkpoint.PreviousDigest
}

func (b *CheckpointData) GetFirehoseBlockParentNumber() uint64 {
	if b.Checkpoint.SequenceNumber == 0 {
		return 0
	}

	return b.Checkpoint.SequenceNumber - 1
}

func (b *CheckpointData) GetFirehoseBlockTime() time.Time {
	return time.Unix(0, int64(b.Checkpoint.TimestampMs)).UTC()
}

func (b *CheckpointData) GetFirehoseBlockLIBNum() uint64 {
	number := b.Checkpoint.SequenceNumber

	if number <= bstream.GetProtocolFirstStreamableBlock {
		return number
	}

	// Since there is no forks blocks on Sui, I'm pretty sure that last irreversible block number
	// is the block's number itself. However I'm not sure overall how the Firehose stack would react
	// to LIBNum == Num so to play safe for now, previous block of current is irreversible.
	return number - 1
}

func (b *CheckpointData) AsRef() bstream.BlockRef {
	return bstream.NewBlockRef(b.GetFirehoseBlockID(), b.Checkpoint.SequenceNumber)
}

