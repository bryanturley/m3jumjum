// Copyright 2021  Bryan Turley


package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)


const (
	SM3_ROM_MD5 = "0x21f3e98df4780ee1c667b84e57d88675"
	SM3_ROM_NAME = "Super Metroid (JU) [!].smc"
)


func main() {
	b, err := ioutil.ReadFile(SM3_ROM_NAME)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	// check md5 in background just in case
	md5Check := make(chan bool)
	go func() {
		sum := md5.Sum(b)
		// too lazy to change the hash string into a byte array..
		s := "0x"
		for _, v := range sum {
			s += fmt.Sprintf("%02x", v)
		}
		if s != SM3_ROM_MD5 {
			fmt.Printf("\n'%s' !=\n'%s' rom md5 mismatch!\n", s, SM3_ROM_MD5)
			md5Check <- false
			panic("bad input rom file")
		} else {
			fmt.Println("md5 all good")
		}
		md5Check <- true
	} ()
	defer func() {
		<- md5Check
	} ()


	
}
