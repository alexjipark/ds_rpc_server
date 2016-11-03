package main

import (
	"github.com/alexjipark/ds_rpc_server/libc"
	"fmt"
	"github.com/alexjipark/datastreet/test"
	"encoding/hex"
)

func main() {
	// Test - Generating Address..
	addr := libc.GenerateAddress("test1")
	fmt.Printf("Generated Address : %s\n", addr)

	src_account := test.PrivateAccountFromSecret("test")
//	dst_account := test.PrivateAccountFromSecret("test1")
	hex_src_account := hex.EncodeToString(src_account.Account.PubKey.Address())
//	hex_dst_account := hex.EncodeToString(dst_account.Account.PubKey.Address())
/*
	//transferCurrency (server_addr string, secret string, src_addr string, dst_addr string, amount int64 ) (bool, error)
	ret, err := libc.TransferCurrency("35.160.145.128:46657", "test", hex_src_account, hex_dst_account, 10, 15)

	if err != nil {
		if ret == false {
			fmt.Println("Transfer Currency Done..")
		}
	}
*/
	//func CheckBalance (server_addr string, addr string) (amount int, err error)
	retAmount, _ := libc.CheckBalance("35.160.145.128:46657", hex_src_account)

	fmt.Println("Amount :", retAmount)
}
