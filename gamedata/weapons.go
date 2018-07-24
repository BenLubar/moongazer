package gamedata

const weaponsFilename = "weapons-csv.json"

type weapons struct {
	Data []*Weapon `json:"data"`
}

type Weapon struct {
	SoundEffectType string   `json:"SFX_TYPE"`
	PDT             string   `json:"PDT"`
	Aspects         []string `json:"ASPECTS"`
	Description     string   `json:"DESCRIPTION"`
	GlowEffect      bool     `json:"glow_effect"`
	Hint            string   `json:"HINT"` // BASIC or NONE
	Gather          struct {
		Harvest *FloatInt `json:"HRV,omitempty"` // guessed
		Dig     *FloatInt `json:"DIG,omitempty"`
		Smash   *FloatInt `json:"SMS,omitempty"`
		Chop    *FloatInt `json:"CHP,omitempty"`
		Net     *FloatInt `json:"NET,omitempty"`
		Mine    *FloatInt `json:"ORE,omitempty"`
		Sweep   *FloatInt `json:"FIN,omitempty"`
	} `json:"MIN"`
	Joke       string   `json:"JOKE"`
	Sprite     string   `json:"SPR"`
	Flaming    bool     `json:"flaming"`
	Durability FloatInt `json:"DUR"`
	Name       string   `json:"NAM"`
	Damage     struct {
		Bash  *FloatInt `json:"BSH,omitempty"`
		Slash *FloatInt `json:"SLA,omitempty"`
		Poke  *FloatInt `json:"POK,omitempty"`
	} `json:"DAM"`
	Resistance struct {
		Fire *FloatInt `json:"FIRE,omitempty"`
		Rust *FloatInt `json:"RUST,omitempty"`
		// deprecated?
		FireOld *FloatInt `json:"burn,omitempty"`
		RustOld *FloatInt `json:"rust,omitempty"`
	} `json:"RES"`
	Style               string   `json:"STY"` // MEL, JAB, or BIG
	MetalEffect         bool     `json:"metal_effect"`
	Speed               FloatInt `json:"SPD"`
	ID                  string   `json:"ID"`
	Tags                []string `json:"TAGS"`
	SoundEffectMaterial string   `json:"SFX_MATERIAL"`
}
