package savedgame

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"strings"
)

type Summary struct {
	Path string
	Size int64
	SummaryDetails
}

func (s *Summary) ByteOrder() binary.ByteOrder {
	if s.LittleEndian {
		return binary.LittleEndian
	}
	return binary.BigEndian
}

type SummaryDetails struct {
	// Serialization format
	LittleEndian  bool `json:"_littleEndian"`
	FormatVersion int  `json:"_version"`

	// Game setup
	GameVersion int    `json:"gp.version"`
	Seed        []byte `json:"gp.seed,base64"`
	Name        string `json:"name"`
	Plan        string `json:"plan"`

	// ???
	MoonTransits [256]int `json:"gp.moonTransits"`

	// Game state
	GameOver    bool `json:"gp.gameOver"`
	HasEnded    bool `json:"gp.hasEnded"`
	HasStarted  bool `json:"gp.hasStarted"`
	PlayerAlive bool `json:"gp.playerAlive"`

	// Counters
	BlocksMined int `json:"gp.blocksMined"`
	Deaths      int `json:"gp.deaths"`
	MobsKilled  int `json:"gp.mobsKilled"`

	// Moonstones (?)
	MesaMoonstone bool `json:"gp.mesaMoonstone"`
	PaleMoonstone bool `json:"gp.paleMoonstone"`
	PyreMoonstone bool `json:"gp.pyreMoonstone"`
	VeilMoonstone bool `json:"gp.veilMoonstone"`

	// Time
	LastSaved          int64 `json:"lastSaved,string"`
	TotalSecondsPlayed int   `json:"gp.totalSecondsPlayed"`
	WorldTime          int   `json:"gp.worldTime"`

	// Regions
	LastVisitedSafeRegion int              `json:"gp.lastVisitedSafeRegion"`
	LastRegion            int              `json:"lastRegion"`
	LastRegionName        string           `json:"lastRegionName"`
	NumRegions            FloatInt         `json:"numRegions"`
	Regions               []*SummaryRegion `json:"regions"`
}

type SummaryRegion struct {
	// Binary data
	Attributes    []byte `json:"_attributes,base64"`
	BlockDatabase []byte `json:"_block_database,base64"`
	Checkpoints   []byte `json:"_checkpoints,base64"`
	DoorData      []byte `json:"_door_data,base64"`
	FeatureData   []byte `json:"_feature_data,base64"`
	ItemData      []byte `json:"_item_data,base64"`
	NPCData       []byte `json:"_npc_data,base64"`
	Tags          []byte `json:"_tags,base64"`
	WorldObjects  []byte `json:"_wo_data,base64"`

	// ???
	Discovered      bool       `json:"discovered"`
	H               FloatInt   `json:"h"`
	HasMoonstone    bool       `json:"hasMoonstone"`
	Hidden          bool       `json:"hidden"`
	I               FloatInt   `json:"i"`
	MoonType        int        `json:"moonType"`
	Name            string     `json:"name"`
	RespawnOutside  bool       `json:"respawnOutside"`
	Seed            FloatInt   `json:"seed"`
	SpentSpawnCosts [6]float64 `json:"spent_spawn_costs"`
	Type            string     `json:"type"`
	Visited         bool       `json:"visited"`
	W               FloatInt   `json:"w"`
}

func GetSummary(path string) (*Summary, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return getSummary(fi, filepath.Dir(path))
}

func getSummary(fi os.FileInfo, savesDir string) (*Summary, error) {
	if fi.IsDir() || (!strings.HasSuffix(fi.Name(), ".sav") && !strings.HasSuffix(fi.Name(), ".bak")) {
		return nil, nil
	}

	fullName := filepath.Join(savesDir, fi.Name())

	var details SummaryDetails
	err := readSaveJSON(fullName, "save.json", &details)

	return &Summary{
		Path:           fullName,
		Size:           fi.Size(),
		SummaryDetails: details,
	}, err
}
