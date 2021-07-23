## sifnoded query ibc-transfer params

Query the current ibc-transfer parameters

### Synopsis

Query the current ibc-transfer parameters

```
sifnoded query ibc-transfer params [flags]
```

### Examples

```
<appd> query ibc-transfer params
```

### Options

```
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for params
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID
```

### SEE ALSO

* [sifnoded query ibc-transfer](sifnoded_query_ibc-transfer.md)	 - IBC fungible token transfer query subcommands

###### Auto generated by spf13/cobra on 2-Jul-2021