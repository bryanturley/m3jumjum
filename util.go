// Copyright 2021  Bryan Turley

package main

import (
	"fmt"
)

// return a offset into a LoRom for the bank
func BankOffset(bank uint8) (offset uint32) {
	if bank < 0x80 {
		panic(fmt.Sprintf("bank %d is < 0x80\n", bank))
	}
	if bank > 0xFD {
		panic(fmt.Sprintf("bank %d is > 0xFD\n", bank))
	}

	offset = (uint32(bank) - 0x80) * 0x8000
	if offset > (1 << 24) {
		panic(fmt.Sprintf("%d is out of 24bit addr space", offset))
	}

	return
}

// return a slice from a LoRom data that covers only this bank.
func BankSlice(b []byte, bank uint8) (offset uint32, data []byte) {
	offset = BankOffset(bank)
	if offset > uint32(len(b)) {
		panic(fmt.Sprintf("offset %d is > len(b) %d\n", offset, len(b)))
	}
	if offset + 0x8000 > uint32(len(b)) {
		panic(fmt.Sprintf("full bank at offset %d is > len(b) %d\n", offset + 0x8000, len(b)))
	}

	data = b[offset:offset + 0x8000]

	return
}
