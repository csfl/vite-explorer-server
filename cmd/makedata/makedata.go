package main

import (
	"github.com/vitelabs/go-vite/vitedb"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func main() {
	writeGenesisBlocks()
	writeGenesisSnapshotBlock()
	//
	writeAccoutChain()
	//test()
	//for i:=0; i < 102; i ++{
	//	writeSnapshotChain()
	//	time.Sleep(10000)
	//}
}

//func testWrite()  {
//	iterKey := []byte("c.%01.%001")
//	db, _ := vitedb.GetLDBDataBase(vitedb.DB_BLOCK)
//	db.Leveldb.Put(iterKey, []byte("hello world"), nil)
//
//	iterKey2 := []byte("c.%01.%002")
//	db.Leveldb.Put(iterKey2, []byte("hello world2"), nil)
//}
func test () {
	iterKey := []byte("h.")
	db, _ := vitedb.GetLDBDataBase(vitedb.DB_BLOCK)
	//value, err := db.Leveldb.Get(iterKey, nil)
	//fmt.Println(err)
	//fmt.Println(string(value))
	fmt.Println(string(util.BytesPrefix(iterKey).Start))
	fmt.Println(string(util.BytesPrefix(iterKey).Limit))
	//fmt.Println()
	iter := db.Leveldb.NewIterator(util.BytesPrefix(iterKey), nil)
	//
	fmt.Println(string(iter.Key()))
	fmt.Println(iter.Value())
	for iter.Next() {
		fmt.Println(string(iter.Key()))
		fmt.Println(iter.Value())
	}
}


