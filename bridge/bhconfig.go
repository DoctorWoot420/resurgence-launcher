package bridge

import (
	"strconv"

	"github.com/DoctorWoot420/resurgence-launcher/bhconfig"
	"github.com/DoctorWoot420/resurgence-launcher/log"
	"github.com/therecipe/qt/core"
)

// bhconfigBridge is the connection between QML and the bhconfig model.
type bhconfigBridge struct {
	core.QObject

	// Dependencies.
	bhconfigService bhconfig.Service
	logger          log.Logger

	// Properties.
	_ bool `property:"loading"`
	_ bool `property:"error"`

	// Models.
	bhconfigModel *core.QAbstractListModel `property:"items"`

	// Slots.
	_ func(gameIndex int) `slot:"getMaphackTextFromParams"`
}

// Connect will connect the QML signals to functions in Go.
func (b *bhconfigBridge) Connect() {
	b.ConnectGetMaphackTextFromParams(b.getMaphackTextFromParams)
}

func (b *bhconfigBridge) getMaphackTextFromParams(gameIndex int) {
	b.logger.Debug("bridge/bhconfig.go start getMaphackTextFromParams with gameIndex value: " + strconv.Itoa(gameIndex))

	// Tell the GUI that we're fetching data.
	b.SetLoading(true)

	err := b.bhconfigService.SetMaphackTextData(gameIndex)

	if err != nil {
		b.logger.Error(err)
		return
	}

	// Tell the UI that we're done fetching data.
	b.SetLoading(false)

	return
}

// NewBhconfig sets up a bhconfig bridge with all dependencies.
func NewBhconfig(bs bhconfig.Service, bm *bhconfig.Model, logger log.Logger) *bhconfigBridge {
	l := NewBhconfigBridge(nil)

	// Setup dependencies.
	l.bhconfigService = bs
	l.logger = logger

	// Setup model.
	l.SetItems(bm)

	// Set initial state.
	l.SetLoading(false)
	l.SetError(false)

	return l
}
