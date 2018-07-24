package archive

import "github.com/davecgh/go-spew/spew"

type NPCData struct {
	Version uint32
	NPCs    []NPC
}

func (nd *NPCData) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	dec.moonMark()

	nd.Version = dec.expect32(4)
	nd.NPCs = make([]NPC, dec.uint32())
	for i := range nd.NPCs {
		nd.NPCs[i].Decode(dec)
	}

	dec.end()
}

type NPC struct {
	Unk2         uint32
	Unk3         uint32
	Unk4         uint32
	Unk5         uint32
	Unk6         uint32
	Unk7         uint32
	Unk8         bool
	Unk9         bool
	Unk10        uint32
	Unk11        [8]byte
	Unk12        uint32
	OverrideName bool
	Name         string
	Unk15        bool
	Unk16        uint32
	Unk17        uint32
	Unk18        uint32
	Unk19        [25]byte
	Unknown      []byte
}

func (n *NPC) Decode(dec *Decoder) {
	n.Unknown = dec.expectSize(209)
	n.Unk2 = dec.uint32()
	if n.Unk2 != 8 && n.Unk2 != 16 && n.Unk2 != 32 {
		spew.Dump(dec.data, n)
		panic("unk2")
	}
	n.Unk3 = dec.uint32()
	n.Unk4 = dec.uint32()
	n.Unk5 = dec.uint32()
	n.Unk6 = dec.uint32()
	n.Unk7 = dec.uint32()
	n.Unk8 = dec.bool()
	n.Unk9 = dec.bool()
	n.Unk10 = dec.uint32()
	dec.expect0(n.Unk11[:])
	n.Unk12 = dec.uint32()
	n.OverrideName = dec.bool()
	n.Name = dec.staticString(128)
	n.Unk15 = dec.bool()
	n.Unk16 = dec.uint32()
	n.Unk17 = dec.uint32()
	n.Unk18 = dec.uint32()
	dec.expect0(n.Unk19[:])
}
