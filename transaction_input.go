package main

import "bytes"

//交易输入
type TXInput struct {
	Txid []byte //// ID of such transaction
	Vout int    // an index of an output in the transaction
	// ScriptSig string
	Signature []byte
	PubKey    []byte
}

//判断某一哈希是否为交易输入的公钥的哈希
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
