//libdatastreet.go
package main

import "C"
import (
	"github.com/alexjipark/ds_rpc_server/libc"
	"fmt"
)

//export genDataStAddr
func genDataStAddr(secret *C.char) *C.char{
	genAddr := libc.GenerateAddress(C.GoString(secret))
	return C.CString(genAddr)
}

//export transferCoin
func transferCoin (server_addr *C.char, secret *C.char, src_addr *C.char, dst_addr *C.char, amount int64, sequence int) bool {
	retbool, err := libc.TransferCurrency(C.GoString(server_addr), C.GoString(secret), C.GoString(src_addr), C.GoString(dst_addr), amount, sequence)
	fmt.Println(retbool, err)

	return retbool
}

//export checkBalance
func checkBalance (server_addr *C.char, src_addr *C.char) int64 {
	amount, _ := libc.CheckBalance(C.GoString(server_addr), C.GoString(src_addr))
	//fmt.Printf("[%s]'s Amount:%v .. with Error(%s)", C.GoString(src_addr), amount, err.Error())

	return amount
}

//export transferDataOwnership
func transferDataOwnership(server_addr *C.char, secret *C.char, src_addr *C.char, data_hash *C.char, dst_addr *C.char, sequence int) bool {
	//func TransferDataOwnership (server_addr string, secret string, src_addr string, data_hash string, dst_addr string, sequence int) (bool, error)
	retbool, err := libc.TransferDataOwnership(C.GoString(server_addr),C.GoString(secret),C.GoString(src_addr),C.GoString(data_hash),C.GoString(dst_addr), sequence)
	fmt.Println( retbool, err)

	return retbool
}

func main() {

}