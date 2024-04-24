package main

import (
	"github.com/apocentre/firehose-sui/codec"
	pbsui "github.com/apocentre/firehose-sui/pb/sf/sui/type/v1"
	firecore "github.com/streamingfast/firehose-core"
	fhCmd "github.com/streamingfast/firehose-core/cmd"
)

func main() {
	fhCmd.Main(&firecore.Chain[*pbsui.CheckpointData]{
		ShortName:            "firesui",
		LongName:             "SUI",
		ExecutableName:       "sui-sf-indexer",
		FullyQualifiedModule: "github.com/apocentre/firehose-sui",
		Version:              version,

		FirstStreamableBlock: 1,

		BlockFactory:         func() firecore.Block { return new(pbsui.CheckpointData) },
		ConsoleReaderFactory: codec.NewConsoleReader,

		Tools: &firecore.ToolsConfig[*pbsui.CheckpointData]{},
	})
}

// Version value, injected via go build `ldflags` at build time, **must** not be removed or inlined
var version = "dev"
