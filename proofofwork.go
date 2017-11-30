package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var maxNonce = math.MaxInt64

const targetBits = 16

//工作量证明结构体
type ProofOfWork struct {
	block  *Block
	target *big.Int //目标难度
}

/*工作量证明*/

//函数，新建一个工作量证明
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

//接收者为pow的方法，准备pow所需数据
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			// pow.block.Data,
			pow.block.HashTransactions(),
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	// data := bytes.Join([][]byte{pow.block.PrevBlockHash, pow.block.Data, IntToHex(pow.block.Timestamp), IntToHex(int64(targetBits)), IntToHex(int64(nonce))}, []byte{})
	return data
}

//接收者为pow的方法，运行工作量证明，返回随机数和哈希字符串(byte数组)
func (pow *ProofOfWork) RunPow() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	// fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	fmt.Printf("Mining a new block")

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:]) //SetBytes:将32字节表示的大整数转换成大整形数并赋给方法接收者

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

//验证pow
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
