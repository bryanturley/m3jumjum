// Copyright 2021  Bryan Turley


package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)


const (
	M3_ROM_MD5 = "0x21f3e98df4780ee1c667b84e57d88675"
	M3_ROM_NAME = "Super Metroid (JU) [!].smc"

	M3_ROM_HDR_NAME_LOC = 0x7FC0
	M3_ROM_HDR_NAME = "Super Metroid        "

)

func main() {
	b, err := ioutil.ReadFile(M3_ROM_NAME)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("'%s' %d bytes\n", M3_ROM_NAME, len(b))
	if len(b) < 2 << 20 {
		fmt.Printf("'%s' is to small\n", M3_ROM_NAME)
		return
	}
	if uint32(len(b)) % 0x8000 != 0 {
		fmt.Printf("'%s' is not a multiple of 0x8000\n", M3_ROM_NAME)
		return
	}
	// check for rom header name
	hdrName := string(b[M3_ROM_HDR_NAME_LOC:M3_ROM_HDR_NAME_LOC + 21])
	fmt.Printf("header name: '%s'\n", hdrName)
	if hdrName != M3_ROM_HDR_NAME {
		fmt.Println("incorrect header name in rom")
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
		if s != M3_ROM_MD5 {
			fmt.Printf("\n'%s' !=\n'%s' rom md5 mismatch!\n", s, M3_ROM_MD5)
			md5Check <- false
			panic("bad input rom file")
		}
		md5Check <- true
	} ()
	defer func() {
		<- md5Check
	} ()

	g := AbsorbGame(b)

	if g == nil {}
}




func AbsorbGame(rom []byte) (g *Game) {
	g = new(Game)
	g.Rom = rom
	
	g.AbsorbBank83()

	return
}

