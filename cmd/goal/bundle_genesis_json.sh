#!/usr/bin/env bash

THISDIR=$(dirname $0)

cat <<EOM | gofmt > $THISDIR/bundledGenesisInject.go
// Code generated by bundle_genesis_json.sh. DO NOT EDIT.
package main
var genesisMainnet []byte
var genesisTestnet []byte
var genesisBetanet []byte
var genesisDevnet []byte
func init() {
        genesisMainnet = []byte(\`$(cat $THISDIR/../../installer/genesis/mainnet/genesis.json)\`)
        genesisTestnet = []byte(\`$(cat $THISDIR/../../installer/genesis/testnet/genesis.json)\`)
        genesisBetanet = []byte(\`$(cat $THISDIR/../../installer/genesis/betanet/genesis.json)\`)
        genesisDevnet = []byte(\`$(cat $THISDIR/../../installer/genesis/devnet/genesis.json)\`)
}
EOM