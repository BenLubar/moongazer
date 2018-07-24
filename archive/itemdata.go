package archive

type ItemData struct {
	Unk1  uint32
	Items []Item
}

func (id *ItemData) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	dec.moonMark()

	id.Unk1 = dec.expect32(2)
	id.Items = make([]Item, dec.uint32())
	for i := range id.Items {
		id.Items[i].Decode(dec)
	}

	dec.end()
}

type Item struct {
	Unk1  uint32
	Unk2  uint32
	Unk3  uint32
	Unk4  uint16
	Unk5  bool
	Unk6  uint32
	Unk7  [3]byte
	Unk8  uint32
	Unk9  uint32
	Unk10 [20]byte
	Unk11 uint32
}

func (i *Item) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	i.Unk1 = dec.expect32(50)
	i.Unk2 = dec.uint32()
	i.Unk3 = dec.uint32()
	i.Unk4 = dec.uint16()
	i.Unk5 = dec.bool()
	i.Unk6 = dec.uint32()
	dec.expect0(i.Unk7[:])
	i.Unk8 = dec.uint32()
	i.Unk9 = dec.uint32()
	dec.expect0(i.Unk10[:])
	i.Unk11 = dec.uint32()
}
