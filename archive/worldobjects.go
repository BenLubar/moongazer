package archive

type WorldObjects struct {
	Version uint32
	Objects []WorldObject
}

func (wo *WorldObjects) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	dec.moonMark()

	wo.Version = dec.expect32(3)
	wo.Objects = make([]WorldObject, dec.uint32())

	for i := range wo.Objects {
		wo.Objects[i].Decode(dec)
	}

	dec.end()
}

type WorldObject struct {
	Unk1         uint32
	Unk2         uint32
	Unk3         uint32
	Unk4         uint32
	Unk5         uint8
	Unk6         [11]byte
	Unk7         uint32
	OverrideName bool
	Name         string
	Unk8         bool
	Unk9         int32
	Unk10        int32
	Unk11        int32
	Unk12        int32
	Unk13        int32
	Unk14        int32
	Unk15        int32
	Unk16        int32
	Unk17        int32
	Unk18        int32
	Unk19        int32
	Unk20        int32
	Unk21        int32
	Unk22        int32
	Unk23        int32
	Unk24        int32
	Unk25        uint8
	Unknown      []byte
}

func (w *WorldObject) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	w.Unknown = dec.expectSize(227)
	w.Unk1 = dec.uint32()
	w.Unk2 = dec.uint32()
	w.Unk3 = dec.expect32(0)
	w.Unk4 = dec.uint32()
	w.Unk5 = dec.uint8()
	dec.expect0(w.Unk6[:])
	w.Unk7 = dec.uint32()
	w.OverrideName = dec.bool()
	w.Name = dec.staticString(128)
	w.Unk8 = dec.bool()
	w.Unk9 = dec.int32n1()
	w.Unk10 = dec.int32n1()
	w.Unk11 = dec.int32n1()
	w.Unk12 = dec.int32n1()
	w.Unk13 = dec.int32n1()
	w.Unk14 = dec.int32n1()
	w.Unk15 = dec.int32n1()
	w.Unk16 = dec.int32n1()
	w.Unk17 = dec.int32n1()
	w.Unk18 = dec.int32n1()
	w.Unk19 = dec.int32n1()
	w.Unk20 = dec.int32n1()
	w.Unk21 = dec.int32n1()
	w.Unk22 = dec.int32n1()
	w.Unk23 = dec.int32n1()
	w.Unk24 = dec.int32n1()
	w.Unk25 = dec.uint8()
}
