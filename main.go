//go:generate goversioninfo -icon=icon.ico -manifest=resurgence-launcher.exe.manifest -64

package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"

	"github.com/DoctorWoot420/resurgence-launcher/bhconfig"
	"github.com/DoctorWoot420/resurgence-launcher/bridge"
	bhconfigClient "github.com/DoctorWoot420/resurgence-launcher/clients/bhconfig"
	ladderClient "github.com/DoctorWoot420/resurgence-launcher/clients/ladder"
	"github.com/DoctorWoot420/resurgence-launcher/clients/resurgence"
	"github.com/DoctorWoot420/resurgence-launcher/config"
	"github.com/DoctorWoot420/resurgence-launcher/d2"
	"github.com/DoctorWoot420/resurgence-launcher/ladder"
	"github.com/DoctorWoot420/resurgence-launcher/log"
	"github.com/DoctorWoot420/resurgence-launcher/news"
	"github.com/DoctorWoot420/resurgence-launcher/storage"
	"github.com/nokka/goqmlframeless"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

func main() {
	// Environment variables set when building.
	var (
		debugMode    = envBool("DEBUG_MODE", true)
		environment  = envString("ENVIRONMENT", "production")
		buildVersion = envString("BUILD_VERSION", "v1.2.0")
	)

	// Set app context.
	core.QCoreApplication_SetApplicationName("Resurgence Launcher")
	core.QCoreApplication_SetOrganizationName("SlashdiabloResurgence")
	core.QCoreApplication_SetOrganizationDomain("https://resurgence.slashgaming.net/")
	core.QCoreApplication_SetApplicationVersion("1.2.0")

	// Enable high DPI scaling
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	core.QCoreApplication_SetAttribute(core.Qt__AA_UseHighDpiPixmaps, true)

	// Create base application.
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Define default window dimensions
	width := 1024
	height := 600

	// Get the primary screen
	screen := app.PrimaryScreen()
	screenGeometry := screen.Geometry()
	devicePixelRatio := screen.DevicePixelRatio()

	// Check if we're on a raw 1920x1080 screen (no scaling)
	if screenGeometry.Width() < 1280 {
		// Reduce window size for 1080p screens with high zoom
		scaleFactor := 1 / devicePixelRatio * 1.69
		width = int(float64(width) * scaleFactor)
		height = int(float64(height) * scaleFactor)
	}

	fw := goqmlframeless.NewWindow(goqmlframeless.Options{
		Width:  width,
		Height: height,
		Alpha:  1.0,
		Color:  goqmlframeless.RGB{R: 3, G: 2, B: 2},
	})

	// QML Widget that will be used to draw on.
	qmlWidget := quick.NewQQuickWidget(nil)
	qmlWidget.SetResizeMode(quick.QQuickWidget__SizeRootObjectToView)

	// Set margins to prevent unnecessary borders
	qmlWidget.SetContentsMargins(0, 0, 0, 0)
	fw.Layout.SetContentsMargins(0, 0, 0, 0)

	// Add QML widget to layout.
	fw.Layout.AddWidget(qmlWidget, 0, 0)

	configPath, err := getConfigPath()
	if err != nil {
		os.Exit(0)
	}

	// Data directory is a requirement for the app.
	os.MkdirAll(configPath, storage.Permissions)

	// Setup file logger.
	logger := log.NewLogger(configPath)

	// Enable debugger if it was enabled through the env variable.
	if debugMode {
		enableDebugger(logger)
	}

	// Setup local storage.
	store := storage.NewStore(configPath)
	if err := store.Load(); err != nil {
		logger.Error(errors.New("unable to load config"))
		os.Exit(0)
	}

	conf, err := store.Read()
	if err != nil {
		logger.Error(errors.New("unable to read config"))
		os.Exit(0)
	}

	// Models.
	lm := ladder.NewTopLadderModel(nil)
	gm := config.NewGameModel(nil)
	nm := news.NewModel(nil)
	fm := d2.NewFileModel(nil)
	bm := bhconfig.NewModel(nil)

	// Setup clients.
	sc := resurgence.NewClient()
	lc := ladderClient.NewClient()
	bc := bhconfigClient.NewClient(logger)

	// Setup services.
	cs := config.NewService(sc, store, gm, logger)
	d2s := d2.NewService(sc, cs, logger, fm)
	ls := ladder.NewService(lc, lm)
	ns := news.NewService(sc, nm)
	bs := bhconfig.NewService(bc, bm, cs, logger)

	// Populate the game model with the game config
	// before passing it to the config bridge.
	populateGameModel(conf, gm, logger)

	// Setup QML bridges with all dependencies.
	diabloBridge := bridge.NewDiablo(d2s, fm, conf.LaunchDelay, logger)
	configBridge := bridge.NewConfig(cs, gm, configPath, logger)
	ladderBridge := bridge.NewLadder(ls, lm, logger)
	newsBridge := bridge.NewNews(ns, nm, logger)
	bhconfigBridge := bridge.NewBhconfig(bs, bm, logger)

	// Add bridges to QML.
	qmlWidget.RootContext().SetContextProperty("diablo", diabloBridge)
	diabloBridge.Connect()

	qmlWidget.RootContext().SetContextProperty("settings", configBridge)
	configBridge.Connect()

	qmlWidget.RootContext().SetContextProperty("ladder", ladderBridge)
	ladderBridge.Connect()

	qmlWidget.RootContext().SetContextProperty("news", newsBridge)
	newsBridge.Connect()

	qmlWidget.RootContext().SetContextProperty("bhconfig", bhconfigBridge)
	bhconfigBridge.Connect()

	// Set build version on the bridge to inform the gui.
	configBridge.SetBuildVersion(buildVersion)

	// Make sure the window is allowed to minimize.
	goqmlframeless.AllowMinimize(fw.WinId())

	// Set the source for our drawable widget to our qml entrypoint.
	if environment == "production" {
		qmlWidget.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	} else {
		// Allows for reloading QML without rebuilding, useful while developing.
		qmlWidget.SetSource(core.NewQUrl3("qml/main.qml", 0))
	}

	fw.Show()
	app.Exec()
}

// getConfigPath returns the target specific application data directory.
func getConfigPath() (string, error) {
	locations := core.QStandardPaths_StandardLocations(
		core.QStandardPaths__AppLocalDataLocation,
	)
	if len(locations) == 0 {
		return "", errors.New("failed to locate application data directory")
	}

	// Grab the first available location.
	return locations[0], nil
}

func populateGameModel(conf *storage.Config, gm *config.GameModel, logger log.Logger) {
	for _, game := range conf.Games {
		g := config.NewGame(nil)
		g.ID = game.ID
		g.Location = game.Location
		g.Instances = game.Instances
		g.OverrideBHCfg = game.OverrideBHCfg
		g.Flags = game.Flags
		g.HDVersion = game.HDVersion
		g.MaphackVersion = game.MaphackVersion
		g.MaphackDefaultGs = game.MaphackDefaultGs
		g.MaphackDefaultGameName = game.MaphackDefaultGameName
		g.MaphackDefaultPassword = game.MaphackDefaultPassword
		g.MaphackRuneDesign = game.MaphackRuneDesign
		g.MaphackFilterBlocks = game.MaphackFilterBlocks

		gm.AddGame(g)
	}
}

// enableDebugger will capture stdout and stderr output.
func enableDebugger(logger log.Logger) {
	r, w, err := os.Pipe()
	if err != nil {
		os.Exit(0)
	}

	os.Stdout = w
	os.Stderr = w

	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := scanner.Text()
			logger.Debug(line)
		}
	}()
}

// envString extracts a string from os environment.
func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

// envBool extracts a bool from os environment.
func envBool(env string, fallback bool) bool {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}

	v, err := strconv.ParseBool(e)
	if err != nil {
		return fallback
	}

	return v
}
