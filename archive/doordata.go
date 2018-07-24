package archive

type DoorData struct {
	Unk1  uint32
	Doors []Door
}

func (dd *DoorData) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	dec.moonMark()

	dd.Unk1 = dec.expect32(3)
	dd.Doors = make([]Door, dec.uint32())
	for i := range dd.Doors {
		dd.Doors[i].Decode(dec)
	}

	dec.end()
}

type Door struct {
	Unk1  uint32
	Unk2  uint32
	Unk3  uint32
	Unk4  uint32
	Unk5  uint64
	Unk6  uint32
	Unk7  uint32
	Unk8  uint32
	Unk9  uint32
	Unk10 uint32
	Unk11 [131]byte
	Unk12 bool
	Unk13 uint32
}

func (d *Door) Decode(dec *Decoder) {
	d.Unk1 = dec.expect32(176)
	d.Unk2 = dec.uint32()
	d.Unk3 = dec.uint32()
	d.Unk4 = dec.expect32(0)
	d.Unk5 = dec.uint64()
	d.Unk6 = dec.expect32(0)
	d.Unk7 = dec.uint32()
	d.Unk8 = dec.expect32(0)
	d.Unk9 = dec.expect32(0)
	d.Unk10 = dec.uint32()
	dec.expect0(d.Unk11[:])
	d.Unk12 = dec.bool()
	d.Unk13 = dec.uint32()
}
