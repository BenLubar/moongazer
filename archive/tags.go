package archive

type Tags struct {
	Version uint32
	Tags    []Tag
}

func (t *Tags) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	dec.moonMark()

	t.Version = dec.expect32(1)
	t.Tags = make([]Tag, dec.uint32())
	for i := range t.Tags {
		t.Tags[i].Decode(dec)
	}

	dec.end()
}

type TagType uint16

const (
	TagMining TagType = 1
	TagItem   TagType = 3
)

type Tag struct {
	Type         TagType
	Unk2         uint16
	MiningType   string
	Unk4         uint32
	Unk5         uint32
	Unk6         uint32
	Unk7         uint32
	Unk8         uint32
	Unk9         uint32
	MiningAmount uint32
	Unk11        string
	Unk12        uint32
	Unk13        uint32
	Unk14        uint32
	Unk15        uint32
	Unk16        uint32
	Unk17        [8]byte
	Item         string
	Unk19        uint32
	Unk20        uint32
	Unknown      []byte
}

func (t *Tag) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	t.Unknown = dec.expectSize(212)
	t.Type = TagType(dec.uint16())
	t.Unk2 = dec.uint16()
	t.MiningType = dec.staticString(8)
	t.Unk4 = dec.uint32()
	t.Unk5 = dec.uint32()
	t.Unk6 = dec.uint32()
	t.Unk7 = dec.uint32()
	t.Unk8 = dec.uint32()
	t.Unk9 = dec.uint32()
	t.MiningAmount = dec.uint32()
	t.Unk11 = dec.staticString(8)
	t.Unk12 = dec.uint32()
	t.Unk13 = dec.uint32()
	t.Unk14 = dec.uint32()
	t.Unk15 = dec.uint32()
	t.Unk16 = dec.uint32()
	dec.expect0(t.Unk17[:])
	t.Item = dec.staticString(128)
	t.Unk19 = dec.uint32()
	t.Unk20 = dec.uint32()
}
