package archive

type Checkpoints struct {
	Version     uint32
	Checkpoints []Checkpoint
}

func (c *Checkpoints) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	dec.moonMark()

	c.Version = dec.expect32(1)
	c.Checkpoints = make([]Checkpoint, dec.uint32())
	for i := range c.Checkpoints {
		c.Checkpoints[i].Decode(dec)
	}

	dec.end()
}

type Checkpoint struct {
	Unk2    uint32 // index?
	Unk3    bool   // unlocked?
	Name    string
	Unknown []byte
}

func (c *Checkpoint) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	c.Unknown = dec.expectSize(133)
	c.Unk2 = dec.uint32()
	c.Unk3 = dec.bool()
	c.Name = dec.staticString(128)
}
