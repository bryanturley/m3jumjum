// Copyright 2021  Bryan Turley

package main


// fields in rom order. packing likely is not accurate though


type RoomStateHdr struct {
	LvlDataPtr   uint32
	Tileset      uint8
	Songset      uint8
	PlayInd      uint8
	FXPtr        uint16
	EnemySetPtr  uint16
	EnemyGfxPtr  uint16
	BGXYScroll   uint16
	RScrollPtr   uint16
	UnusedPtr    uint16
	MainAsm      uint16
	PLMSetPtr    uint16
	BGPtr        uint16
	SetupAsmPtr  uint16
}


type RoomHdr struct {
	Ind          uint8
	Area         uint8
	MapX         uint8
	MapY         uint8
	W            uint8
	H            uint8
	UpScroll     uint8
	DnScroll     uint8
	SpGfxBits    uint8
	DoorOutPtr   uint16

	// Event1 Data is Optional ??
	Ev1Ptr       uint16
	Ev1Val       uint8
	Ev1RoomPtr   uint16

	Std1Ptr      uint16
	Std          RoomStateHdr

	// Event1 Data is Optional ??
	Ev1          RoomStateHdr

	DoorPtr      [2]uint16
}

type Door struct {
	RoomID     uint16
	Bits       uint8
	Dir        uint8
	DoorX      uint8
	DoorY      uint8
	ScreenX    uint8
	ScreenY    uint8
	DistSpawn  uint16
	DoorAsm    uint16
}



// all the game data, that we process
type Game struct {
	Rom []byte

	Doors []Door
}


// Bank83 contains FX Data, Door Data, more...
func (g *Game) AbsorbBank83() {
	g.Doors = make([]Door, 0, 2048)
	
}
