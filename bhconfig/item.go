package bhconfig

import "github.com/therecipe/qt/core"

// Item represents a maphack config item
type Item struct {
	core.QObject

	ConfigText []string
}
