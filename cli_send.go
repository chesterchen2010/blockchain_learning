package main

import (
	"fmt"
	"log"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !ValidateAddress(from) {
		log.Panic("ERROR: sender Address is not valid")
	}
	if !ValidateAddress(to) {
		log.Panic("ERROR: recipient Address is not valid")
	}
	bc := NewBlockchain(nodeID)
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	// tx := NewUTXOTransaction(from, to, amount, &UTXOSet)
	// // bc.MineBlock([]*Transaction{tx})
	// cbTx := NewCoinbaseTX(from, "")
	// txs := []*Transaction{cbTx, tx}

	// newBlock := bc.MineBlock(txs)
	// UTXOSet.Update(newBlock)

	wallets, err := NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := NewCoinbaseTX(from, "")
		txs := []*Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		sendTx(knownNodes[0], tx)
	}

	fmt.Println("Success send!")
}
