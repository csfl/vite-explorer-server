package main

import (
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/ledger"
	"log"
	"math/big"
	"time"
	"math/rand"
	"github.com/vitelabs/go-vite/common/types"
)

var snapshotblockchain = access.GetSnapshotChainAccess()

func writeGenesisSnapshotBlock () {
	//var hash = []byte("000000000000000000")
	//var prevHash = []byte("000000000000000000")
	//block := createSnapshotBlock(hash, prevHash, big.NewInt(1))
	block := ledger.GetGenesisSnapshot()
	err := snapshotblockchain.WriteBlock(block)
	if err != nil {
		log.Fatal(err)
	}
}

func writeSnapshotChain()  {
	preBlock, glbErr := snapshotblockchain.GetLatestBlock()
	if glbErr != nil {
		log.Fatal(glbErr)
	}
	var height = &big.Int{}
	height = height.Add(preBlock.Height, big.NewInt(1))
	block := createSnapshotBlock(createHash(), preBlock.Hash, height)
	err := snapshotblockchain.WriteBlock(block)
	if err != nil {
		log.Fatal(err)
	}
}

func createSnapshotBlock (hash []byte, prevHash []byte, height *big.Int) *ledger.SnapshotBlock{
	snapshotBLock := &ledger.SnapshotBlock{
		Hash: hash,
		PrevHash: prevHash,
		Height: height,
		Producer: createSnapshotBlockProducer(),
		Snapshot: createSnapshot(),
		Signature: createAccountBlockSignature(),
		Timestamp: uint64(time.Now().Unix()),
	}
	return snapshotBLock
}

func createSnapshot () map[string] []byte{
	accountList := getAccountAddressList()
	//fmt.Printf("snapshot count is:%d\n", len(accountList))
	if accountList == nil {
		return nil
	}

	var snapshot map[string] []byte
	snapshot = make(map[string] []byte)

	for _, address := range accountList {
		//fmt.Println("snapshot[", data.String(), "]", data.Bytes())
		accoutblock, err := accountChainAccess.GetLatestBlockByAccountAddress(address)
		if err != nil {
			return nil
		}
		snapshot[address.String()] = accoutblock.Hash
	}
	return snapshot
}


func createHash () []byte {
	var letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 20)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func createSnapshotBlockProducer () []byte {
	accountAddressList := getAccountAddressList()
	if accountAddressList == nil {
		return []byte("00000000000000000001")
	}
	return accountAddressList[rand.Intn(len(accountAddressList))].Bytes()
}

func getAccountAddressList () []*types.Address {
	accountList, err := accountChainAccess.GetAccountList()
	if err != nil {
		log.Println("GetAccountList error.")
		return nil
	}
	return accountList
}

//func getSnapshotChainTest () {
//	snapshotblockchain, gbErr := snapshotblockchain.GetBlockList(0,1,200)
//	if gbErr !=nil {
//		log.Fatal(gbErr)
//	}
//	//fmt.Println("Length of the snapshotblockchain: ", len(snapshotblockchain))
//	for _, block := range snapshotblockchain {
//		fmt.Printf("Hash:%s,\nPrevHash:%s,\nProducer:%s,\nHeight:%s,\nTimestamp:%d,\n\n",
//			hex.EncodeToString(block.Hash), hex.EncodeToString(block.PrevHash),
//			hex.EncodeToString(block.Producer), block.Height, block.Timestamp)
//	}
//}
