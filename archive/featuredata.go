package archive

type FeatureData struct {
	Unk1     uint32 // seems to always be 4
	Features []Feature
}

func (fd *FeatureData) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	dec.moonMark()

	fd.Unk1 = dec.uint32()
	fd.Features = make([]Feature, dec.uint32())
	for i := range fd.Features {
		fd.Features[i].Decode(dec)
	}

	dec.end()
}

type Feature struct {
	Unk1  uint32
	Unk2  uint8
	Unk3  bool
	Unk4  uint16
	Unk5  uint32
	Unk6  uint32
	Unk7  uint32
	Unk8  uint32
	Unk9  uint32
	Unk10 bool
	Unk11 bool
	Unk12 uint32
	Unk13 uint16
	Unk14 uint16
	Unk15 uint32
	Unk16 uint32
	Unk17 uint32
	Unk18 uint32
	Unk19 uint32
	Unk20 uint32
	Unk21 uint32
}

func (f *Feature) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	f.Unk1 = dec.expect32(62)
	f.Unk2 = dec.uint8()
	f.Unk3 = dec.bool()
	f.Unk4 = dec.uint16()
	f.Unk5 = dec.uint32()
	f.Unk6 = dec.uint32()
	f.Unk7 = dec.expect32(0)
	f.Unk8 = dec.expect32(0)
	f.Unk9 = dec.expect32(0)
	f.Unk10 = dec.bool()
	f.Unk11 = dec.bool()
	f.Unk12 = dec.expect32(0)
	f.Unk13 = dec.expect16(0)
	f.Unk14 = dec.expect16(16368)
	f.Unk15 = dec.uint32()
	f.Unk16 = dec.uint32()
	f.Unk17 = dec.uint32()
	f.Unk18 = dec.uint32()
	f.Unk19 = dec.uint32()
	f.Unk20 = dec.uint32()
	f.Unk21 = dec.uint32()
}
