package main

import "github.com/tendermint/tendermint/types"

func main() {
	vs := new(types.ValidatorSet)
	vs.FromBytes([]byte{0x01, 0x01, 0x30, 0x01, 0x06, 0x30, 0x30, 0x30, 0x30, 0x30,0x30})
}