# firehose-sui
Firehose on Sui Blockchain

## Pre-requisites
You would need to download the `sf-sui-extractor` source code and build it locally.

```bash
git clone https://github.com/apocentre/sf-sui-extractor
cd sf-sui-extractor
cargo build --release
```
Make sure the `target/release/sf-sui-extractor` binary is moved to a folder that is in the PATH so it can be executed.

> Binaries will be available to download soon so you don't have to build from source code.


## Sui Config

1. `mkdir ~/.sf_sui && mkdir ~/.sf_sui/sui_config`
2. `cp devel/mainnet/config/* ~/.sf_sui/sui_config`

We would need to update the following fields to point to the paths of our choice:

1. `db-path`: Set this to the location where the SUI node will store it's database i.e. the blockchain data
2. `genesis-file-location`: This should be path to the `~/.sf_sui/sui_config/genesis.blob` but since we can't use ~ in yaml files we would need to replace it with `/Users/<user>` on Mac or `/home/<user>` on linux.

We should also update `--sui-node-config` in `indexer.yaml` to point to `/Users/<user>/.sf_sui/sui_config/full_node.yaml` (or `/home/<user>/.sf_sui/sui_config/full_node.yaml` on linux)

## Build Protobuf models

```bash
./pb/generate.sh
```

## Run

Before we run we need to create the `firesui` binary.

```bash
go install ./cmd/firesui
```

Run `reader-node` and `merger`. This instance will read data from the stdin and store (and merge) the results in files on the disk.

```bash
firesui -c ./devel/mainnet/indexer.yaml start
```

Finally, if we want to stream the above saved files to the upstream substreams we would need to run the `firehose` and `relayer`.

```bash
firesui -c ./devel/mainnet/firehose.yaml start
```


## License

[Apache 2.0](LICENSE)
