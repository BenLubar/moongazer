package archive

import (
	"encoding/binary"
	"math"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

const moonMark uint32 = 0xF00DF00D

type Decoder struct {
	data    []byte
	version int
	order   binary.ByteOrder
}

func NewDecoder(data []byte, version int, order binary.ByteOrder) *Decoder {
	return &Decoder{
		data:    data,
		version: version,
		order:   order,
	}
}

func (dec *Decoder) supportMaxVersion(max int) {
	if dec.version > max {
		panic("unsupported game version")
	}
}

func (dec *Decoder) expectBytes(min int) {
	if len(dec.data) < min {
		spew.Dump(dec.data, min)
		panic("unexpected end of file")
	}
}

func (dec *Decoder) moonMark() {
	if dec.uint32() != moonMark {
		panic("invalid moon mark")
	}
}

func (dec *Decoder) bool() bool {
	switch n := dec.uint8(); n {
	case 0:
		return false
	case 1:
		return true
	default:
		spew.Dump(n)
		panic("invalid boolean")
	}
}

func (dec *Decoder) uint8() uint8 {
	dec.expectBytes(1)
	n := dec.data[0]
	dec.data = dec.data[1:]
	return n
}

func (dec *Decoder) uint16() uint16 {
	dec.expectBytes(2)
	n := dec.order.Uint16(dec.data)
	dec.data = dec.data[2:]
	return n
}

func (dec *Decoder) uint32() uint32 {
	dec.expectBytes(4)
	n := dec.order.Uint32(dec.data)
	dec.data = dec.data[4:]
	return n
}

func (dec *Decoder) int32() int32 {
	return int32(dec.uint32())
}

func (dec *Decoder) int32n1() int32 {
	n := dec.int32()
	if n < -1 {
		spew.Dump(n)
		panic("!int32n1")
	}
	return n
}

func (dec *Decoder) uint64() uint64 {
	dec.expectBytes(8)
	n := dec.order.Uint64(dec.data)
	dec.data = dec.data[8:]
	return n
}

func (dec *Decoder) expect32(expected uint32) uint32 {
	actual := dec.uint32()
	if actual != expected {
		spew.Dump(dec.data, actual, expected)
		panic("!expect32")
	}

	return actual
}

func (dec *Decoder) expect16(expected uint16) uint16 {
	actual := dec.uint16()
	if actual != expected {
		spew.Dump(dec.data, actual, expected)
		panic("!expect16")
	}

	return actual
}

func (dec *Decoder) expect0(b []byte) {
	dec.expectBytes(len(b))
	dec.data = dec.data[copy(b, dec.data):]
	for _, n := range b {
		if n != 0 {
			spew.Dump(b, dec.data)
			panic("nonzero byte")
		}
	}
}

func (dec *Decoder) expectSize(size uint32) []byte {
	length := dec.uint32()
	if size > length {
		spew.Dump(length, dec.data)
		panic("length too short")
	}
	dec.expectBytes(int(length))
	var extra []byte
	if length != size {
		extra = make([]byte, length-size)
		copy(extra, dec.data[size:])
		copy(dec.data[len(extra):], dec.data)
		dec.data = dec.data[len(extra):]
	}
	return extra
}

func (dec *Decoder) float32() float32 {
	n := dec.uint32()
	f := math.Float32frombits(n)
	if math.IsNaN(float64(f)) {
		panic("NaN")
	}
	return f
}

func (dec *Decoder) float64() float64 {
	n := dec.uint64()
	f := math.Float64frombits(n)
	if math.IsNaN(f) {
		panic("NaN")
	}
	return f
}

func (dec *Decoder) string() string {
	l := dec.uint32()
	dec.expectBytes(int(l))
	s := string(dec.data[:l])
	dec.data = dec.data[l:]
	return s
}

func (dec *Decoder) staticString(length int) string {
	dec.expectBytes(length)
	s := string(dec.data[:length])
	dec.data = dec.data[length:]

	end := strings.IndexByte(s, 0)
	if end == -1 {
		spew.Dump([]byte(s))
		panic("unterminated string")
	}

	for i := end; i < length; i++ {
		if s[i] != 0 {
			spew.Dump([]byte(s))
			panic("incorrectly terminated string")
		}
	}

	return s[:end]
}

func (dec *Decoder) end() {
	if len(dec.data) != 0 {
		spew.Dump(dec.data)
		panic("expected end of file")
	}
}
