package storage

import "encoding/json"

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
	MaphackItemNameOption  string   `json:"maphack_item_name_option"`
	MaphackFilterBlocks    []string `json:"maphack_filter_blocks"`
}

// UnmarshalJSON accepts both launcher keys (maphack_*) and generator payload keys
// (item_name_option, rune_design, etc.) so a hand-edited config still loads.
func (g *Game) UnmarshalJSON(data []byte) error {
	type rawGame Game
	aux := struct {
		rawGame
		ItemNameOption  string   `json:"item_name_option"`
		RuneDesign      string   `json:"rune_design"`
		FilterBlocks    []string `json:"filter_blocks"`
		DefaultGs       string   `json:"default_gs"`
		DefaultGameName string   `json:"default_game_name"`
		DefaultPassword string   `json:"default_password"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*g = Game(aux.rawGame)
	if g.MaphackItemNameOption == "" && aux.ItemNameOption != "" {
		g.MaphackItemNameOption = aux.ItemNameOption
	}
	if g.MaphackRuneDesign == "" && aux.RuneDesign != "" {
		g.MaphackRuneDesign = aux.RuneDesign
	}
	if len(g.MaphackFilterBlocks) == 0 && len(aux.FilterBlocks) > 0 {
		g.MaphackFilterBlocks = aux.FilterBlocks
	}
	if g.MaphackDefaultGs == "" && aux.DefaultGs != "" {
		g.MaphackDefaultGs = aux.DefaultGs
	}
	if g.MaphackDefaultGameName == "" && aux.DefaultGameName != "" {
		g.MaphackDefaultGameName = aux.DefaultGameName
	}
	if g.MaphackDefaultPassword == "" && aux.DefaultPassword != "" {
		g.MaphackDefaultPassword = aux.DefaultPassword
	}
	return nil
}
