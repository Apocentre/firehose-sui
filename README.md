# firehose-sui
Firehose on Sui Blockchain

## Pre-requisites
You would need to download the `sui-sf-indexer` source code and build it locally.

```bash
git clone https://github.com/apocentre/sui-sf-indexer
cd sui-sf-indexer
cargo build --release
```
Make sure the `target/release/sui-sf-indexer` binary is moved to a folder that is in the PATH so it can be executed.

> Binaries will be available to download soon so you don't have to build from source code.


## Build Protobuf models

```bash
./pb/generate.sh
```

## Run

Before we run we need to create the `firesui` binary.

```bash
go install ./cmd/firesui
```

Then we just run it by pointing to the correct config file.

```bash
firesui -c ./devel/mainnet/indexer.yaml start
```

## License

[Apache 2.0](LICENSE)
