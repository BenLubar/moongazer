package archive

type Attributes struct {
	Unk1  uint32
	Unk2  uint32
	Unk3  uint32
	Unk4  bool
	Unk5  uint32
	Unk6  uint8
	Unk7  [7]byte
	Unk8  bool
	Unk9  uint32
	Unk10 uint16
	Unk11 uint16
	Unk12 uint32
	Unk13 uint32
	Unk14 uint32
	Unk15 uint32
	Unk16 uint32
	Unk17 uint32
	Unk18 uint32
	Unk19 uint16
	Unk20 uint32
	Unk21 uint32
	Unk22 uint32
	Unk23 uint16
	Unk24 uint16
	Unk25 uint32
	Unk26 uint32
	Unk27 uint32
	Unk28 uint32
	Unk29 uint32
	Unk30 uint32
	Unk31 uint16
	Unk32 uint16
	Unk33 uint16
	Unk34 uint16
	Unk35 uint16
	Unk36 uint16
	Unk37 uint16
	Unk38 uint16
	Unk39 uint16
	Unk40 uint16
	Unk41 uint16
	Unk42 uint16
	Unk43 uint32
	Unk44 uint32
	Unk45 uint32
	Unk46 uint32
	Unk47 uint32
	Unk48 uint32
	Unk49 uint32
	Unk50 uint32
	Unk51 uint32
	Unk52 uint32
	Unk53 uint32
	Unk54 [6]byte
	Unk55 uint16
	Unk56 uint8
	Unk57 uint32
}

func (a *Attributes) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	dec.moonMark()

	a.Unk1 = dec.expect32(11)
	a.Unk2 = dec.expect32(177)
	a.Unk3 = dec.uint32()
	a.Unk4 = dec.bool()
	a.Unk5 = dec.uint32()
	a.Unk6 = dec.uint8()
	dec.expect0(a.Unk7[:])
	a.Unk8 = dec.bool()
	a.Unk9 = dec.uint32()
	a.Unk10 = dec.uint16()
	a.Unk11 = dec.uint16()
	a.Unk12 = dec.uint32()
	a.Unk13 = dec.uint32()
	a.Unk14 = dec.expect32(210)
	a.Unk15 = dec.expect32(210)
	a.Unk16 = dec.expect32(261)
	a.Unk17 = dec.expect32(210)
	a.Unk18 = dec.expect32(0)
	a.Unk19 = dec.expect16(16368)
	a.Unk20 = dec.uint32()
	a.Unk21 = dec.uint32()
	a.Unk22 = dec.uint32()
	a.Unk23 = dec.expect16(210)
	a.Unk24 = dec.uint16()
	a.Unk25 = dec.uint32()
	a.Unk26 = dec.uint32()
	a.Unk27 = dec.uint32()
	a.Unk28 = dec.uint32()
	a.Unk29 = dec.uint32()
	a.Unk30 = dec.uint32()
	a.Unk31 = dec.expect16(0)
	a.Unk32 = dec.uint16()
	a.Unk33 = dec.expect16(0)
	a.Unk34 = dec.uint16()
	a.Unk35 = dec.uint16()
	a.Unk36 = dec.uint16()
	a.Unk37 = dec.expect16(0)
	a.Unk38 = dec.uint16()
	a.Unk39 = dec.expect16(0)
	a.Unk40 = dec.uint16()
	a.Unk41 = dec.uint16()
	a.Unk42 = dec.uint16()
	a.Unk43 = dec.uint32()
	a.Unk44 = dec.uint32()
	a.Unk45 = dec.uint32()
	a.Unk46 = dec.uint32()
	a.Unk47 = dec.uint32()
	a.Unk48 = dec.uint32()
	a.Unk49 = dec.uint32()
	a.Unk50 = dec.uint32()
	a.Unk51 = dec.uint32()
	a.Unk52 = dec.uint32()
	a.Unk53 = dec.expect32(1)
	dec.expect0(a.Unk54[:])
	a.Unk55 = dec.uint16()
	a.Unk56 = dec.uint8()
	a.Unk57 = dec.uint32()

	dec.end()
}
