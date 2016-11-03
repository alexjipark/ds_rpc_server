package main

import (
	"fmt"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/alexjipark/ds_rpc_server/protobuf"
	"strconv"
	"encoding/csv"
)

func main() {
	fmt.Println("Started DataStreet RPC Server..")
	channel := make(chan *ProtobufTest.TestMessage)
	go func() {
		for {
			message := <- channel
			writeValuesToFile(message)

		}
	}()

	//listen to the TCP Port
	listener, err := net.Listen("tcp", "127.0.0.1:2110")
	checkError(err)

	for {
		if conn, err := listener.Accept();  err == nil {
			go handleProtoClient(conn, channel)
		} else {
			continue
		}
	}
}

func handleProtoClient (conn net.Conn, c chan *ProtobufTest.TestMessage) {
	fmt.Println("Connection established..")

	defer conn.Close()

	//Create a data buffer of type byte slice with capacity of 4096
	data := make([]byte, 4096)
	//Read the data waiting on the connection and put it in the data buffer
	n, err := conn.Read(data)
	checkError(err)
	fmt.Println("Decoding Protobuf Message")

	//Create an struct pointer of type ProtobufTest.TestMessage Struct
	protodata := new(ProtobufTest.TestMessage)
	//Convert all the data retrieved into the ProtobufTest.TestMessage struct type
	err = proto.Unmarshal(data[0:n], protodata)
	checkError (err)

	//Push the protobuf message into a channel
	c <- protodata
}

func writeValuesToFile (data_received *ProtobufTest.TestMessage) {
	fmt.Println(data_received)

	//Retreive client information from the protobuf message
	ClientName := data_received.GetClientName()
	ClientDesc := data_received.GetDescription()
	ClientID := strconv.Itoa (int(data_received.GetClientId()))

	items := data_received.GetMessageitems()
	//Open file for writes, if the file does not exist then create it
	file,err := os.OpenFile("CSVValues.csv", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	checkError (err)
	//make sure the file gets closed once the function exists
	defer file.Close()

	writer := csv.NewWriter(file)
	for _, item := range items {
		record := []string {ClientID, ClientName, ClientDesc, strconv.Itoa(int(item.GetId())), item.GetItemName()}
		writer.Write(record)
		fmt.Println(record)
	}
	//flush data to the CSV file
	writer.Flush()
	fmt.Println("Finished writing value to CSV file")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error:%s", err.Error())
		os.Exit(1)
	}
}