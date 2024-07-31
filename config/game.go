package config

import (
	"github.com/therecipe/qt/core"
)

const (
	// ModVersionNone is used to determine that no mod version has been chosen for a game.
	ModVersionNone = "none"
)

// Game represents a diablo installation in the configuration.
type Game struct {
	core.QObject

	ID                     string   `json:"id"`
	Location               string   `json:"location"`
	Instances              int      `json:"instances"`
	OverrideBHCfg          bool     `json:"override_bh_cfg"`
	Flags                  []string `json:"flags"`
	HDVersion              string   `json:"hd_version"`
	MaphackVersion         string   `json:"maphack_version"`
	MaphackDefaultGs       string   `json:"maphack_default_gs"`
	MaphackDefaultGameName string   `json:"maphack_default_game_name"`
	MaphackDefaultPassword string   `json:"maphack_default_password"`
	MaphackRuneDesign      string   `json:"maphack_rune_design"`
	MaphackFilterBlocks    []string `json:"maphack_filter_blocks"`
}

// GameMods represents the mods available for a Diablo II game.
type GameMods struct {
	HD      []string `json:"hd"`
	Maphack []string `json:"maphack"`
}

// MaphackOptions represents the options available to pass into maphack configuration
type MaphackOptions struct {
	GameServers  []string `json:"game_servers"`
	RuneDesigns  []string `json:"rune_designs"`
	FilterBlocks []string `json:"filter_blocks"`
}
