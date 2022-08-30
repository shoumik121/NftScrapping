package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/2468f3c54a7b498284b55d91676c0913")

	if err != nil {
		log.Fatal("Failed to connect to Eth Node", err)
	}

	cntxt := context.Background()

	txn, pending, _ := conn.TransactionByHash(cntxt, common.HexToHash("0xd0d904e5bdae559c630889817f88498eb846161410cc349bd240dfa86191926f"))

	if pending {
		fmt.Println("Transaction is pending", txn)
	} else {
		fmt.Println("Transaction is not pending", txn)
	}

}
