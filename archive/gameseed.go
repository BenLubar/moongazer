package archive

import "github.com/davecgh/go-spew/spew"

type GameSeed struct {
}

func (gs *GameSeed) Decode(dec *Decoder) {
	dec.supportMaxVersion(1)

	// TODO

	spew.Dump(dec.data)
	//dec.end()
}
