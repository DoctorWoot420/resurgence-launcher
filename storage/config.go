package storage

// DefaultLaunchDelay is used if a launch delay hasn't been set by a user.
const DefaultLaunchDelay = 1000

// Config is the configuration required to run the app.
type Config struct {
	Games       []Game `json:"games"`
	LaunchDelay int    `json:"launch_delay"`
}

// Game represents a game setup by the user.
type Game struct {
	OverrideBHCfg          bool     `json:"override_bh_cfg"`
	ID                     string   `json:"id"`
	Location               string   `json:"location"`
	Instances              int      `json:"instances"`
	Flags                  []string `json:"flags"`
	HDVersion              string   `json:"hd_version"`
	MaphackVersion         string   `json:"maphack_version"`
	MaphackDefaultGs       string   `json:"maphack_default_gs"`
	MaphackDefaultGameName string   `json:"maphack_default_game_name"`
	MaphackDefaultPassword string   `json:"maphack_default_password"`
	MaphackRuneDesign      string   `json:"maphack_rune_design"`
	MaphackFilterBlocks    []string `json:"maphack_filter_blocks"`
}
