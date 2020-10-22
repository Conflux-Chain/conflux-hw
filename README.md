# Conflux Offline Signer

## Build
Run `go build` under the root directory to generate binary.

## Offline Operations
1. Create an account for the first time:
```
./conflux-hw
```

2. Sign a transaction to transfer to specified user:
```
./conflux-hw sign --epoch <EPOCH_NUM> --nonce <NONCE> --value <VALUE_IN_DRIP> --to <TO_ADDRESS>
```

## Online Operations
1. Query blockchain information for specified account:
```
./conflux-hw query --account <ACCOUNT_ADDRESS>
```

2. Send signed transaction to RPC server:
```
./conflux-hw send --raw <RAW_TX_HEX>
```