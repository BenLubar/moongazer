package archive

type BlockDatabase struct {
	Unk1       uint32
	Foreground []string
	Background []string
}

func (bd *BlockDatabase) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	dec.moonMark()
	bd.Unk1 = dec.uint32()
	bd.Foreground = make([]string, dec.uint32())
	bd.Background = make([]string, dec.uint32())
	for i := range bd.Foreground {
		bd.Foreground[i] = dec.string()
	}
	for i := range bd.Background {
		bd.Background[i] = dec.string()
	}
	dec.end()
}
