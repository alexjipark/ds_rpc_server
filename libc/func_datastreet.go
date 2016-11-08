package libc

import (
	"github.com/alexjipark/datastreet/test"
	"fmt"
	"encoding/hex"
	"github.com/tendermint/go-rpc/client"
	"github.com/alexjipark/datastreet/types"
	"github.com/tendermint/go-wire"

	. "github.com/tendermint/go-common"
	tdtypes "github.com/tendermint/tendermint/types"
	"github.com/tendermint/go-rpc/types"
	"github.com/gorilla/websocket"
	"bytes"
	"encoding/json"
	"time"
	"strings"
)

// 필수 API
// 1. 가상화폐 전송 trasaction
// 2. 특정 계정의 화폐량 balance 체크
// 3. 데이터 소유권 전송 transaction
// 4. 특정 계정의 소유 데이터 리스트 List

// Done..
func GenerateAddress (secret string) (address string) {

	genAccount := test.PrivateAccountFromSecret(secret)
	fmt.Println("Generated Account : " , genAccount)
	hexaddr := hex.EncodeToString(genAccount.Account.PubKey.Address())
	fmt.Println("Account's Address : ", hexaddr)
	return hexaddr
}

// Done..
func TransferCurrency (server_addr string, secret string, src_addr string, dst_addr string, amount int64, sequence int ) (bool, error) {

	ws := rpcclient.NewWSClient(server_addr, "/websocket")
	chainID := "chain-AMUKE0"

	_, err := ws.Start()
	if err != nil {
		fmt.Printf("Error in Starting a client for the Server:[%s]\n", server_addr)
		return false, nil
	}
	defer ws.Close()

	hex_src_addr,err := hex.DecodeString(src_addr)
	if err != nil {
		fmt.Printf("Error in Decoding str[%s]\n", src_addr)
		return false, nil
	}
	hex_dst_addr, err := hex.DecodeString(dst_addr)
	if err != nil {
		fmt.Printf("Error in Decoding str[%s]\n", dst_addr)
		return false, nil
	}
	fmt.Printf("SRC Addr[%X], DST Addr[%X]\n", hex_src_addr, hex_dst_addr)

	// Get the root account
	root := test.PrivateAccountFromSecret(secret)
	// Check Account
	fmt.Printf("Private Key : %X\n", root.PrivKey)
	fmt.Printf("Public Byte : %X\n", root.Account.PubKey.Bytes())
	fmt.Printf("Public Addr : %X\n", root.Account.PubKey.Address())

	// Transfer Coins to 'dst'
	tx := &types.SendTx{
		Inputs: []types.TxInput {
			types.TxInput{
				Address: root.Account.PubKey.Address(),
				PubKey: root.Account.PubKey,
				Coins: types.Coins {{"USD", amount}},
				Sequence: sequence+1,
			},
		},
		Outputs: []types.TxOutput {
			types.TxOutput{
				Address: hex_dst_addr,
				Coins: types.Coins{{"USD", amount}},
			},
		},
	}

	//Sign Request
	signBytes := tx.SignBytes(chainID)
	sig := root.PrivKey.Sign(signBytes)
	tx.Inputs[0].Signature = sig
	fmt.Println("tx: ", tx)

	txBytes := wire.BinaryBytes(struct{types.Tx}{tx})

	//Subscribe an event for Tx
	eid := tdtypes.EventStringTx(tdtypes.Tx(txBytes))
	if err = ws.Subscribe(eid); err != nil {
		fmt.Printf("Error in subscribing EventStringTx[%s]\n", eid)
		fmt.Println(err.Error())
	}
	defer func() {
		ws.Unsubscribe(eid)
	}()

	//Write Request
	request := rpctypes.NewRPCRequest("fakeid", "broadcast_tx_commit", Arr(txBytes))
	fmt.Println("Request: ", request)

	reqBytes := wire.JSONBytes(request)
	err = ws.WriteMessage(websocket.TextMessage, reqBytes)
	if err != nil {
		fmt.Println("Error in Writing request through websocket.. :", err.Error())
		return false, nil
	}

	//Wait For an Event
	test.WaitForEvent(ws, eid, true, func(){}, func(eid string, b interface{}) error {
		evt, ok := b.(tdtypes.EventDataTx)
		if !ok {
			fmt.Println("Got Wrong Event Type..", b)
		} else {
			if bytes.Compare(evt.Tx, txBytes) != 0 {
				fmt.Println("got returned Event with diffferent Tx")
			}
			fmt.Println("Event Code :", evt.Code)
			fmt.Println("Event Error:", evt.Error)
			fmt.Println("Event Log:", evt.Log)
			fmt.Println("Event Tx Hash : ", evt.Tx.Hash())
		}
		return nil
	})

	return true, nil
}

// Done..
func CheckBalance (server_addr string, addr string) (amount int64, err error) {
	ws := rpcclient.NewWSClient(server_addr, "/websocket")
	chainID := "chain-AMUKE0"
	amount = 0

	_,err = ws.Start()
	if err != nil {
		fmt.Printf("Error in Starting a client for the Server:[%s]\n", server_addr)
		return 0, nil
	}
	defer ws.Close()

	hex_addr_bytes, err := hex.DecodeString(addr)
	if err != nil {
		fmt.Println ("Error in Decoding Address..", err.Error())
		return 0, nil
	}

	queryBytes := make ([]byte, 1+ wire.ByteSliceSize(hex_addr_bytes))
	buf := queryBytes
	buf[0] = 0x01 //Get TypeByte
	buf = buf[1:]
	wire.PutByteSlice(buf, hex_addr_bytes)

	reqQuery := rpctypes.NewRPCRequest(chainID, "tmsp_query", Arr(queryBytes))
	reqQueryBytes := wire.JSONBytes(reqQuery)
	err = ws.WriteMessage(websocket.TextMessage, reqQueryBytes)
	if err != nil {
		fmt.Println("Error in Writing Web Socket Message..", err.Error())
		return 0,nil
	}

	// go routine to wait for websocket msg..
	goodCh := make(chan interface{})
	errCh  := make(chan error)
	// Read Responses
	go func() {
		LOOP:
		for{
			select {
			case res,ok := <-ws.ResultsCh:
				if !ok {
					fmt.Println("Not ok from rpcclient..")
					break LOOP
				}

				// Check the result
				var result []interface{}
				err := json.Unmarshal([]byte(string(res)), &result)
				if err != nil {
					fmt.Println("Error in Unmarshalling with ", err.Error())
					break
				}

				resData := result[1].(map[string]interface{})["result"].(map[string]interface{})["Data"]
				hexBytes, err := hex.DecodeString(resData.(string))
				if err != nil {
					fmt.Printf ("Error in Decoding [%s]\n", resData.(string))
					break LOOP
				}
				fmt.Println("Decoded HexBytes :", hexBytes)

				goodCh <- hexBytes
				break LOOP
			case err := <- ws.ErrorsCh:
				errCh <- err
				break LOOP
			case <-ws.Quit:
				break LOOP
			}
		}
	}()

	// wait for an event or timeout
	timeout := time.NewTimer(10 * time.Second)
	select {
	case <- timeout.C:
		ws.Stop()
		fmt.Println("timeout..")
	case eventData := <-goodCh:
		var acc *types.Account
		err = wire.ReadBinaryBytes(eventData.([]byte), &acc)

		// Result
		fmt.Printf("Account : %X\n", acc.PubKey)
		fmt.Printf("Balance : %v\n", acc.Balance)
		fmt.Printf("Sequence : %v\n", acc.Sequence)
		for _, pcoin := range acc.Balance {
			if strings.Compare(pcoin.Denom, "USD") == 0 {
				amount = pcoin.Amount
				break
			}
		}
	case err := <-errCh:
		panic(err)	// show the stack trace..
	}

	return amount, nil
}

func TransferDataOwnership (server_addr string, secret string, src_addr string, data_hash string, dst_addr string, sequence int) (bool, error) {

	ws := rpcclient.NewWSClient(server_addr, "/websocket")
	chainID := "chain-AMUKE0"

	_, err := ws.Start()
	if err != nil {
		fmt.Printf("Error in Starting a client for the Server:[%s]\n", server_addr)
		return false, nil
	}

	defer ws.Close()

	hex_src_addr, err := hex.DecodeString(src_addr)
	if err != nil {
		fmt.Printf("Error in Decoding str[%s]\n", src_addr)
		return false, nil
	}

	hex_dst_addr, err := hex.DecodeString(dst_addr)
	if err != nil {
		fmt.Printf("Error in Decoding str[%s]\n", dst_addr)
		return false, nil
	}
	fmt.Printf("SRC Addr[%X], DST Addr[%X]\n", hex_src_addr, hex_dst_addr)

	// Get the Source Account
	root := test.PrivateAccountFromSecret(secret)
	//Check Account
	fmt.Printf("Private Key : %X\n", root.PrivKey)
	fmt.Printf("Public Byte : %X\n", root.Account.PubKey.Bytes())
	fmt.Printf("Public Addr : %X\n", root.Account.PubKey.Address())

	// instead of coin type, put data_hash in place..
	tx := &types.SendTx {
		Inputs: []types.TxInput {
			types.TxInput {
				Address: hex_src_addr,	// root. Account.PubKey.Address()
				PubKey: root.Account.PubKey,
				Coins: types.Coins {{data_hash, 1}},
				Sequence: sequence+1,
			},
		},
		Outputs: []types.TxOutput {
			types.TxOutput{
				Address: hex_dst_addr,
				Coins: types.Coins{{data_hash, 1}},
			},
		},
	}

	//Sign Request.. Bytes to be signed..
	signBytes := tx.SignBytes(chainID)
	sig := root.PrivKey.Sign(signBytes) // made Signature from signBytes
	tx.Inputs[0].Signature = sig
	fmt.Println("tx: ", tx)

	txBytes := wire.BinaryBytes(struct{types.Tx}{tx})

	//Subscribe an event for Tx
	eid := tdtypes.EventStringTx(tdtypes.Tx(txBytes))
	if err = ws.Subscribe(eid); err != nil {
		fmt.Printf("Error in subscribing EventStringTx[%s]\n", eid)
		fmt.Println(err.Error())
	}
	defer ws.Unsubscribe(eid)

	//Write Request
	request := rpctypes.NewRPCRequest("fakeid", "broadcast_tx_commit", Arr(txBytes))
	reqBytes := wire.JSONBytes(request)
	err = ws.WriteMessage(websocket.TextMessage, reqBytes)

	if err != nil {
		fmt.Println("Error in Writing request through websocket.. :", err.Error())
		return false, nil
	}

	// Wait For an Event
	test.WaitForEvent(ws, eid, true, func(){}, func(eid string, b interface{}) error {
		evt, ok := b.(tdtypes.EventDataTx)
		if !ok {
			fmt.Println("Got Wrong Event Type..", b)
		} else {
			if bytes.Compare(evt.Tx, txBytes) != 0 {
				fmt.Println("got returned Event with diffferent Tx")
			}
			fmt.Println("Event Code :", evt.Code)
			fmt.Println("Event Error:", evt.Error)
			fmt.Println("Event Log:", evt.Log)
			fmt.Println("Event Tx Hash : ", evt.Tx.Hash())
		}
		return nil
	})

	return true, nil
}

func getListofDataOwnership (addr string) []string {
	return nil
}