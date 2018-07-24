package moongazer_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"
	"unicode"

	"github.com/BenLubar/moongazer/archive"
	"github.com/BenLubar/moongazer/savedgame"
	"github.com/davecgh/go-spew/spew"
)

var arrayIndex = regexp.MustCompile(`\[[0-9]+\]|^[0-9_]+\.(?:bak|sav)`)
var unknownFields = make(map[string][]reflect.Value)

func TestSaveReader(t *testing.T) {
	baseDir, err := savedgame.FindBaseDir()
	if err != nil {
		t.Fatalf("failed to find base save directory: %v", err)
	}
	savesDir := filepath.Join(baseDir, "saves")
	liveSaves, err := filepath.Glob(filepath.Join(savesDir, "28786.sav"))
	if err != nil {
		t.Errorf("failed to list live saves: %v", err)
	}
	backupSaves, err := filepath.Glob(filepath.Join(savesDir, "*.bak_"))
	if err != nil {
		t.Errorf("failed to list backup saves: %v", err)
	}

	games := make([]*savedgame.Summary, 0, len(liveSaves)+len(backupSaves))
	for _, paths := range [][]string{liveSaves, backupSaves} {
		for _, path := range paths {
			g, err := savedgame.GetSummary(path)
			if g == nil || err != nil {
				t.Errorf("failed to read save file %q: %v", path, err)
				continue
			}
			games = append(games, g)
		}
	}

	for _, g := range games {
		name := filepath.Base(g.Path) + "/save.json"
		testDecode(t, g, name+"/seed", g.Seed, &archive.GameSeed{})

		for i, r := range g.Regions {
			if i != g.LastRegion {
				continue
			}
			rName := fmt.Sprintf("%s/regions[%d]", name, i)
			testDecode(t, g, rName+"/_attributes", r.Attributes, &archive.Attributes{})
			testDecode(t, g, rName+"/_block_database", r.BlockDatabase, &archive.BlockDatabase{})
			testDecode(t, g, rName+"/_checkpoints", r.Checkpoints, &archive.Checkpoints{})
			testDecode(t, g, rName+"/_door_data", r.DoorData, &archive.DoorData{})
			testDecode(t, g, rName+"/_feature_data", r.FeatureData, &archive.FeatureData{})
			testDecode(t, g, rName+"/_item_data", r.ItemData, &archive.ItemData{})
			testDecode(t, g, rName+"/_npc_data", r.NPCData, &archive.NPCData{})
			testDecode(t, g, rName+"/_tags", r.Tags, &archive.Tags{})
			testDecode(t, g, rName+"/_world_objects", r.WorldObjects, &archive.WorldObjects{})
		}
	}

	if t.Failed() {
		// don't show unknown fields if we already failed
		return
	}

	unknownFieldNames := make([]string, 0, len(unknownFields))
	for fn := range unknownFields {
		unknownFieldNames = append(unknownFieldNames, fn)
	}
	sort.Strings(unknownFieldNames)

	for _, fn := range unknownFieldNames {
		values := make(map[interface{}]int)
		for _, v := range unknownFields[fn] {
			values[v.Interface()]++
		}
		if len(values) <= 1 {
			continue
		}
		t.Errorf("Unknown field: %q", fn)
		type valueCount struct {
			value string
			count int
		}
		valueCounts := make([]valueCount, 0, len(values))
		for v, count := range values {
			valueCounts = append(valueCounts, valueCount{
				value: spew.Sdump(v),
				count: count,
			})
		}
		sort.Slice(valueCounts, func(i, j int) bool {
			if valueCounts[i].count == valueCounts[j].count {
				return valueCounts[i].value < valueCounts[j].value
			}
			return valueCounts[i].count > valueCounts[j].count
		})
		for _, vc := range valueCounts {
			t.Errorf(" %d \u00d7 %s", vc.count, vc.value)
		}
	}
}

func testDecode(t *testing.T, g *savedgame.Summary, name string, data []byte, structure interface{ Decode(*archive.Decoder) }) {
	dec := archive.NewDecoder(data, g.GameVersion, g.ByteOrder())
	structure.Decode(dec)
	checkStructure(t, name, reflect.ValueOf(structure).Elem())
	if testing.Verbose() {
		spew.Dump(structure)
	}
}

func checkStructure(t *testing.T, name string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Ptr:
		if !v.IsNil() {
			checkStructure(t, name+"*", v.Elem())
		}
	case reflect.Struct:
		for fi, fn := 0, v.NumField(); fi < fn; fi++ {
			ft := v.Type().Field(fi)
			fv := v.Field(fi)
			fn := name + "/" + ft.Name

			if ft.Type == reflect.TypeOf([]byte(nil)) {
				if fv.Len() > 0 {
					t.Errorf("%s contains unparsed data:\n%s", fn, spew.Sdump(fv.Interface()))
				}
				continue
			}

			checkStructure(t, fn, fv)
			if strings.HasPrefix(ft.Name, "Unk") {
				strippedName := arrayIndex.ReplaceAllString(fn, "[?]")
				unknownFields[strippedName] = append(unknownFields[strippedName], fv)
			}
		}
	case reflect.Array, reflect.Slice:
		for i, l := 0, v.Len(); i < l; i++ {
			checkStructure(t, fmt.Sprintf("%s[%d]", name, i), v.Index(i))
		}
	case reflect.String:
		nonPrint := strings.FieldsFunc(v.String(), unicode.IsPrint)
		if len(nonPrint) != 0 {
			t.Errorf("%s contains invalid string: %q", name, v.String())
		}
	case reflect.Bool:
		// do nothing
	case reflect.Uint8:
		// do nothing
	case reflect.Uint16:
		// do nothing (for now)
	case reflect.Int32:
		// do nothing (for now)
	case reflect.Uint32:
		if v.Uint()&0xffff == 0 && v.Uint()&0xffff0000 != 0 {
			t.Errorf("%s may be incorrectly defined as uint32: %d (%08x)", name, v.Uint(), v.Uint())
		}
	case reflect.Uint64:
		if v.Uint()&0xffff == 0 && v.Uint()&0xffffffffffff0000 != 0 {
			t.Errorf("%s may be incorrectly defined as uint64: %d (%08x)", name, v.Uint(), v.Uint())
		}
	default:
		panic(name + ": unknown kind: " + v.Kind().String())
	}
}
