# go-explore
Go Explore is a CLI driven Tezos block explorer currently in development. 

# Install
```
git clone https://github.com/DefinitelyNotAGoat/go-explore.git
cd go-explore 
go build
```
# Usage
## Help
```
./go-explore sync --help
sync is the main driver for syncing the explorer into mongodb

Usage:
  Explorer sync [flags]

Flags:
  -d, --database string   databse name (default "tezos")
  -h, --help              help for sync
  -n, --mongo string      mondo dial address (default "mongodb://localhost:27017")
  -u, --node string       address to the node to query (default http://127.0.0.1:8732)(e.g. https://mainnet-node.tzscan.io:443) (default "http://127.0.0.1:8732")
  -r, --refresh int       refresh interval in seconds for sync (default 10)
```
## Example
```
./go-explore sync 
```
