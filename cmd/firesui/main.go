package main

import (
	"github.com/apocentre/firehose-sui/codec"
	pbsui "github.com/apocentre/firehose-sui/pb/sf/sui/type/v1"
	firecore "github.com/streamingfast/firehose-core"
	fhCmd "github.com/streamingfast/firehose-core/cmd"
)

func main() {
	// syncstate.SetupReaderStartBlock()
	
	fhCmd.Main(&firecore.Chain[*pbsui.CheckpointData]{
		ShortName:            "sui",
		LongName:             "Firehose SUI",
		ExecutableName:       "sui-sf-indexer",
		FullyQualifiedModule: "github.com/apocentre/firehose-sui",
		Version:              version,

		BlockFactory:         func() firecore.Block { return new(pbsui.CheckpointData) },
		ConsoleReaderFactory: codec.NewConsoleReader,

		Tools: &firecore.ToolsConfig[*pbsui.CheckpointData]{},
	})
}

// Version value, injected via go build `ldflags` at build time, **must** not be removed or inlined
var version = "dev"
