package main

import (
	"fmt"
	"log"
)
//The IDstore contains a IDMap, which maps string to an array of bytes. This is used to check whether the id of the
//current input has already been retrieved previously.
type IDStore struct {
	IDMap map[string][]byte
}
//constructor function
func NewStore() *IDStore{
	var g IDStore
	g.IDMap = make(map[string][]byte)
	return &g
}
//input an ID of a request, and a response
func (bs *IDStore) PutBlock(ID string,info []byte) {
	//if the id has not been stored previously in the IDMap, store it with value being the response
	if _, ok := bs.IDMap[ID]; !ok {
		bs.IDMap[ID] = info
		fmt.Println("-------Stored an ID!-------")
		return
	}
	//otherwise, we have seen this ID, and thus we don't need to store it again.
	log.Println("*******It is a stored ID! Skip storing it!*******")
}

