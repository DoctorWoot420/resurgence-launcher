import QtQuick 2.12
import QtQuick.Layouts 1.3		// ColumnLayout
import QtQuick.Controls 2.1     // TextField
import QtQuick.Dialogs 1.3      // FileDialog

Item {
    property var game: {}
    property bool depApplied: false
    property bool depError: false
    property int activeHDIndex: 0
    property int activeMaphackIndex: 0
    property int activeMaphackDefaultGsIndex: 0
    property string activeMaphackDefaultGameName: ""
    property string activeMaphackDefaultPassword: ""
    property int activeMaphackRuneDesignIndex: 0
    property int boxHeight: 58

    function setGame(current) {
        // Set current game instance to the view.
        game = current

        // Textfield needs to be set explicitly since it's read only.
        if(game.location != undefined) {
            d2pathInput.text = game.location
        }

        // Update initial states without triggering an animation.
        updateToggleBoxes(current)
        updateHDVersions(current)
        updateMaphackVersions(current)
        updateMaphackDefaultGs(current)
        updateMaphackDefaultGameName(current)
        updateMaphackDefaultPassword(current)
        updateMaphackRuneDesign(current)
        updateMaphackFilterBlocks(current)
    }

    function updateToggleBoxes(current) {
        if(current.flags != null) {
            windowModeFlag.active = current.flags.includes("-w")
            gfxFlag.active = current.flags.includes("-3dfx")
            skipFlag.active = current.flags.includes("-skiptobnet")
        } else {
            windowModeFlag.active = false
            gfxFlag.active = false
            skipFlag.active = false
        }
    }

    // updateHDVersions will set the correct index of the hd mod dropdown.
    function updateHDVersions(current) {
        if(settings.availableHDMods.length > 0) {
            // Find the correct index.
            for(var i = 0; i < settings.availableHDMods.length; i++) {
                if(settings.availableHDMods[i] == current.hd_version) {
                    activeHDIndex = i
                    hdVersion.currentIndex = i
                    return
                }
            }
        }

        // Default to first index in list.
        activeHDIndex = 0
        hdVersion.currentIndex = 0
    }

    // updateMaphackVersions will set the correct index of the maphack mod dropdown.
    function updateMaphackVersions(current) {            
        if(settings.availableMaphackMods.length > 0) {
            // Find the correct index.
            for(var i = 0; i < settings.availableMaphackMods.length; i++) {
                if(settings.availableMaphackMods[i] == current.maphack_version) {
                    activeMaphackIndex = i
                    maphackVersion.currentIndex = i
                    return
                }
            }
        }

        // Default to first index in list.
        activeMaphackIndex = 0
        maphackVersion.currentIndex = 0
    }

    // updateMaphackDefaultGs will set the correct index of the default gs dropdown.
    function updateMaphackDefaultGs(current) {
        if (settings.availableGameServers.length > 0) {
            // Find the correct index.
            for (var i = 0; i < settings.availableGameServers.length; i++) {
                if (settings.availableGameServers[i].toLowerCase() == current.maphack_default_gs.toLowerCase()) {
                    activeMaphackDefaultGsIndex = i;
                    maphackDefaultGs.currentIndex = i;
                    return;
                }
            }
        }

        // Default to first index in list.
        activeMaphackDefaultGsIndex = 0;
        maphackDefaultGs.currentIndex = 0;
    }

    // updateMaphackDefaultGameName will set the correct value of the default game name field.
    function updateMaphackDefaultGameName(current) {
        activeMaphackDefaultGameName = current.maphack_default_game_name !== undefined ? current.maphack_default_game_name : ""
    }

    // updateMaphackDefaultPassword will set the correct value of the default game password field.
    function updateMaphackDefaultPassword(current) {
        activeMaphackDefaultPassword = current.maphack_default_password !== undefined ? current.maphack_default_password : ""
    }

    // updateMaphackRuneDesign will set the correct index of the rune design
    function updateMaphackRuneDesign(current) {
        if(settings.availableRuneDesigns.length > 0) {
            // Find the correct index.
            for(var i = 0; i < settings.availableRuneDesigns.length; i++) {
                if(settings.availableRuneDesigns[i].toLowerCase() == current.maphack_rune_design.toLowerCase()) {
                    activeMaphackRuneDesignIndex = i
                    maphackRuneDesign.currentIndex = i
                    return
                }
            }
        }

        // Default to first index in list.
        activeMaphackRuneDesignIndex = 0
        maphackRuneDesign.currentIndex = 0
    }

    // updateMaphackFilterBlocks will set the correct values of the filter blocks
    function updateMaphackFilterBlocks(current) {
        if(current.maphack_filter_blocks != null) {
            blockLeveling.active = current.maphack_filter_blocks.includes("leveling")
            blockAmazon.active = current.maphack_filter_blocks.includes("amazon")
            blockAssassin.active = current.maphack_filter_blocks.includes("assassin")
            blockBarbarian.active = current.maphack_filter_blocks.includes("barbarian")
            blockDruid.active = current.maphack_filter_blocks.includes("druid")
            blockNecromancer.active = current.maphack_filter_blocks.includes("necromancer")
            blockPaladin.active = current.maphack_filter_blocks.includes("paladin")
            blockSorceress.active = current.maphack_filter_blocks.includes("sorceress")
        } else {
            blockLeveling.active = false
            blockAmazon.active = false
            blockAssassin.active = false
            blockBarbarian.active = false
            blockDruid.active = false
            blockNecromancer.active = false
            blockPaladin.active = false
            blockSorceress.active = false
        }
    }

    function makeFlagList() {
        var flags = []
        if(windowModeFlag.active) {
            flags.push("-w")
        }
        
        if(gfxFlag.active) {
            flags.push("-3dfx")
        }

        if(skipFlag.active) {
            flags.push("-skiptobnet")
        }

        if(nsFlag.active) {
            flags.push("-ns")
        }

        if(nofixaspectFlag.active) {
            flags.push("-nofixaspect")
        }

        if(directTxtFlag.active) {
            flags.push("-direct -txt")
        }

        return flags
    }

    function makeBlockList() {
        var maphackFilterBlocks = []
        if(blockLeveling.active) {
            maphackFilterBlocks.push("leveling")
        }
        if(blockAmazon.active) {
            maphackFilterBlocks.push("amazon")
        }
        if(blockAssassin.active) {
            maphackFilterBlocks.push("assassin")
        }
        if(blockBarbarian.active) {
            maphackFilterBlocks.push("barbarian")
        }
        if(blockDruid.active) {
            maphackFilterBlocks.push("druid")
        }
        if(blockNecromancer.active) {
            maphackFilterBlocks.push("necromancer")
        }
        if(blockPaladin.active) {
            maphackFilterBlocks.push("paladin")
        }
        if(blockSorceress.active) {
            maphackFilterBlocks.push("sorceress")
        }

        return maphackFilterBlocks
    }

    function updateGameModel() {
        if(game != undefined) {
            var body = {
                id: game.id,
                location: d2pathInput.text,
                instances: gameInstances.currentIndex,
                flags: makeFlagList(),
                hd_version: hdVersion.currentText,
                maphack_version: maphackVersion.currentText,
                maphack_default_gs: maphackDefaultGs.currentText,
                maphack_default_game_name: maphackDefaultGameName.text,
                maphack_default_password: maphackDefaultPassword.text,
                maphack_rune_design: maphackRuneDesign.currentText,
                maphack_filter_blocks: makeBlockList(),
            }
            
            settings.upsertGame(JSON.stringify(body))
        }
    }

    Rectangle {
        width: parent.width
        height: 450
        color: "transparent"
        border.width: 0

        Rectangle {
            id: flickableContainer
            width: parent.width
            height: parent.height
            color: "transparent"
            border.width: 0

            Flickable {
                id: currentGame
                width: parent.width
                height: 450
                contentWidth: parent.width
                contentHeight: contentColumn.height
                clip: true

                //anchors.top: parent.top
                //anchors.horizontalCenter: parent.horizontalCenter

                Column {
                    id: contentColumn
                    width: parent.width

                    ColumnLayout {
                        id: settingsLayout
                        width: (currentGame.width * 0.95)
                        spacing: 2

                        anchors.horizontalCenter: parent.horizontalCenter

                        // D2 Directory box.
                        Item {
                            id: fileDialogBox
                            Layout.preferredWidth: settingsLayout.width
                            Layout.preferredHeight: 85

                            Column {
                                anchors.top: parent.top
                                topPadding: 10
                                spacing: 5

                                Title {
                                    text: "SET DIABLO II DIRECTORY"
                                    font.pixelSize: 13
                                }
                            }

                            Row {
                                anchors.bottom: parent.bottom
                                anchors.bottomMargin: 15

                                TextField {
                                    id: d2pathInput
                                    width: fileDialogBox.width * 0.80; height: 35
                                    font.pixelSize: 11
                                    color: "#676767"
                                    readOnly: true
                                    text: (game != undefined ? game.location : "")

                                    background: Rectangle {
                                        color: "#131313"
                                    }
                                }

                                SButton {
                                    id: chooseD2Path
                                    label: "Open"
                                    borderRadius: 0
                                    borderColor: "#373737"
                                    width: fileDialogBox.width * 0.20; height: 35
                                    cursorShape: Qt.PointingHandCursor

                                    onClicked: d2PathDialog.open()
                                }

                                // File dialog.
                                FileDialog {
                                    id: d2PathDialog
                                    selectFolder: true
                                    folder: shortcuts.home
                                    
                                    onAccepted: {
                                        var path = d2PathDialog.fileUrl.toString()
                                        path = path.replace(/^(file:\/{2})/,"")
                                        d2pathInput.text = path
                                        
                                        // Update the game model.
                                        updateGameModel()
                                    }
                                }
                            }
                            
                            Separator{}
                        }

                        // Flags box.
                        Item {
                            Layout.preferredWidth: settingsLayout.width
                            Layout.preferredHeight: boxHeight

                            Row {
                                topPadding: 10

                                Column {
                                    id: parametersText
                                    width: 225
                                    
                                    Title {
                                        text: "LAUNCH PARAMETERS"
                                        font.pixelSize: 13
                                    }

                                    SText {
                                        text: "Set when the game launches"
                                        font.pixelSize: 11
                                        topPadding: 5
                                        color: "#676767"
                                    }
                                }

                                Column {
                                    width: (settingsLayout.width - parametersText.width)

                                    Row {
                                        spacing: 2
                                        leftPadding: 2

                                        ToggleButton {
                                            id: windowModeFlag
                                            label: "-w"
                                            width: 35
                                            height: 35
                                            onClicked: updateGameModel()
                                        }

                                        ToggleButton {
                                            id: gfxFlag
                                            label: "-3dfx"
                                            width: 35
                                            height: 35
                                            onClicked: updateGameModel()
                                        }

                                        ToggleButton {
                                            id: skipFlag
                                            label: "-skip"
                                            width: 35
                                            height: 35
                                            onClicked: updateGameModel()
                                        }

                                        ToggleButton {
                                            id: nsFlag
                                            label: "-ns"
                                            width: 35
                                            height: 35
                                            onClicked: updateGameModel()
                                        }

                                        ToggleButton {
                                            id: nofixaspectFlag
                                            label: "-nofixaspect"
                                            width: 70
                                            height: 35
                                            onClicked: updateGameModel()
                                        }

                                        ToggleButton {
                                            id: directTxtFlag
                                            label: "-direct -txt"
                                            width: 70
                                            height: 35
                                            onClicked: updateGameModel()
                                        }
                                    }
                                }
                            }
                            
                            Separator{}
                        }


                        // Game instances box.
                        Item {
                            Layout.preferredWidth: settingsLayout.width
                            Layout.preferredHeight: boxHeight

                            Row {
                                topPadding: 10

                                Column {
                                    width: (settingsLayout.width - instancesDropdown.width)
                                    
                                    Title {
                                        text: "INSTANCES TO LAUNCH"
                                        font.pixelSize: 13
                                    }

                                    SText {
                                        text: "Number of this specific install that will launch when playing the game"
                                        font.pixelSize: 11
                                        topPadding: 5
                                        color: "#676767"
                                    }
                                }
                                Column {
                                    id: instancesDropdown
                                    width: 60
                                    Dropdown{
                                        id: gameInstances
                                        currentIndex: ((game != undefined && game.instances != undefined) ? (game.instances) : 0)
                                        model: [ 0, 1, 2, 3, 4 ]
                                        height: 30
                                        width: 60

                                        onActivated: updateGameModel()
                                    }
                                }
                            }
                            
                            Separator{}
                        }

                        // Include HD box.
                        Item {
                            Layout.preferredWidth: settingsLayout.width
                            Layout.preferredHeight: boxHeight

                            Row {
                                topPadding: 10

                                Column {
                                    width: (settingsLayout.width - includeHD.width)
                                    Title {
                                        text: "HD MOD VERSION"
                                        font.pixelSize: 13
                                    }

                                    SText {
                                        text: "Select if you want any HD mod installed"
                                        font.pixelSize: 11
                                        topPadding: 5
                                        color: "#676767"
                                    }
                                }
                                Column {
                                    id: includeHD
                                    width: 90

                                    Dropdown{
                                        id: hdVersion
                                        currentIndex: activeHDIndex
                                        model: settings.availableHDMods
                                        height: 30
                                        width: 90

                                        onActivated: updateGameModel()
                                        
                                    }
                                }
                            }
                            
                            Separator{}
                        }

                        // Include maphack box.
                        Item {
                            Layout.preferredWidth: settingsLayout.width
                            Layout.preferredHeight: boxHeight

                            Row {
                                topPadding: 10

                                Column {
                                    width: (settingsLayout.width - includeMaphack.width)
                                    Title {
                                        text: "MAPHACK VERSION"
                                        font.pixelSize: 13
                                    }

                                    SText {
                                        text: "Select if you want any maphack installed"
                                        font.pixelSize: 9
                                        topPadding: 5
                                        color: "#676767"
                                    }
                                }
                                Column {
                                    id: includeMaphack
                                    width: 90

                                    Dropdown{
                                        id: maphackVersion
                                        currentIndex: activeMaphackIndex
                                        model: settings.availableMaphackMods
                                        height: 30
                                        width: 90

                                        onActivated: updateGameModel()
                                    }
                                } 
                            }
                            
                            Separator{}
                        }

                        // Default Game Server Dropdown
                        Item {
                            Layout.preferredWidth: settingsLayout.width
                            Layout.preferredHeight: boxHeight

                            Row {
                                topPadding: 10

                                Column {
                                    width: (settingsLayout.width - includeMaphackDefaultGs.width)
                                    Title {
                                        text: "DEFAULT GAME SERVER"
                                        font.pixelSize: 13
                                    }

                                    SText {
                                        text: "Select the game server to use by default when opening d2"
                                        font.pixelSize: 11
                                        topPadding: 5
                                        color: "#676767"
                                    }
                                }
                                Column {
                                    id: includeMaphackDefaultGs
                                    width: 140

                                    Dropdown{
                                        id: maphackDefaultGs
                                        currentIndex: activeMaphackDefaultGsIndex
                                        model: settings.availableGameServers
                                        height: 30
                                        width: 140

                                        onActivated: updateGameModel()
                                    }
                                } 
                            }
                            
                            Separator{}
                        }

                        // For the Default Game Name Textfield
                        Item {
                            Layout.preferredWidth: settingsLayout.width
                            Layout.preferredHeight: boxHeight

                            Row {
                                topPadding: 10

                                Column {
                                    width: (settingsLayout.width - includeMaphackDefaultGameName.width)
                                    Title {
                                        text: "DEFAULT GAME NAME"
                                        font.pixelSize: 13
                                    }

                                    SText {
                                        text: "Enter a default game when opening d2, blank for none"
                                        font.pixelSize: 11
                                        topPadding: 5
                                        color: "#676767"
                                    }
                                }
                                Column {
                                    id: includeMaphackDefaultGameName
                                    width: 140

                                    TextField {
                                        id: maphackDefaultGameName
                                        text: activeMaphackDefaultGameName
                                        height: 30
                                        width: 140
                                        font.pixelSize: 12
                                        background: Rectangle {
                                            color: "#333333"
                                        }
                                        color: "#FFFFFF"

                                        onEditingFinished: updateGameModel()
                                    }
                                } 
                            }
                            
                            Separator{}
                        }

                        // For the Default Password Textfield
                        Item {
                            Layout.preferredWidth: settingsLayout.width
                            Layout.preferredHeight: boxHeight

                            Row {
                                topPadding: 10

                                Column {
                                    width: (settingsLayout.width - includeMaphackDefaultPassword.width)
                                    Title {
                                        text: "DEFAULT GAME PASSWORD"
                                        font.pixelSize: 13
                                    }

                                    SText {
                                        text: "Enter a default password, blank for none"
                                        font.pixelSize: 11
                                        topPadding: 5
                                        color: "#676767"
                                    }
                                }
                                Column {
                                    id: includeMaphackDefaultPassword
                                    width: 140

                                    TextField {
                                        id: maphackDefaultPassword
                                        text: activeMaphackDefaultPassword
                                        height: 30
                                        width: 140
                                        font.pixelSize: 12
                                        background: Rectangle {
                                            color: "#333333"
                                        }
                                        color: "#FFFFFF"

                                        onEditingFinished: updateGameModel()
                                    }
                                } 
                            }
                            
                            Separator{}
                        }

                        // Rune Design Dropdown
                        Item {
                            Layout.preferredWidth: settingsLayout.width
                            Layout.preferredHeight: boxHeight

                            Row {
                                topPadding: 10

                                Column {
                                    width: (settingsLayout.width - includeMaphackRuneDesign.width)
                                    Title {
                                        text: "RUNE DESIGN"
                                        font.pixelSize: 13
                                    }

                                    SText {
                                        text: "Choose how you want runes to appear"
                                        font.pixelSize: 11
                                        topPadding: 5
                                        color: "#676767"
                                    }
                                }
                                Column {
                                    id: includeMaphackRuneDesign
                                    width: 140

                                    Dropdown{
                                        id: maphackRuneDesign
                                        currentIndex: activeMaphackRuneDesignIndex
                                        model: settings.availableRuneDesigns
                                        height: 30
                                        width: 140

                                        onActivated: updateGameModel()
                                    }
                                } 
                            }
                            
                            Separator{}
                        }

                        // Filter Blocks box
                        Item {
                            Layout.preferredWidth: settingsLayout.width
                            Layout.preferredHeight: 95

                            Row {
                                topPadding: 10

                                Column {
                                    id: maphackFilterBlocksText
                                    width: 235
                                    
                                    Title {
                                        text: "FILTER BLOCKS"
                                        font.pixelSize: 13
                                    }

                                    SText {
                                        text: "Enables groups of alch/craft/base lines"
                                        font.pixelSize: 11
                                        topPadding: 5
                                        color: "#676767"
                                    }
                                }

                                Column {
                                    width: (settingsLayout.width - maphackFilterBlocksText.width)

                                    Row {
                                        spacing: 2
                                        leftPadding: 2

                                        ToggleButton {
                                            id: blockLeveling
                                            label: "Leveling"
                                            width: 70
                                            height: 35
                                            onClicked: updateGameModel()
                                        }
                                        ToggleButton {
                                            id: blockAmazon
                                            label: "Amazon"
                                            width: 70
                                            height: 35
                                            onClicked: updateGameModel()
                                        }
                                        ToggleButton {
                                            id: blockAssassin
                                            label: "Assassin"
                                            width: 70
                                            height: 35
                                            onClicked: updateGameModel()
                                        }
                                        ToggleButton {
                                            id: blockBarbarian
                                            label: "Barbarian"
                                            width: 70
                                            height: 35
                                            onClicked: updateGameModel()
                                        }
                                    }
                                    Row {
                                        spacing: 2
                                        leftPadding: 2

                                        ToggleButton {
                                            id: blockDruid
                                            label: "Druid"
                                            width: 70
                                            height: 35
                                            onClicked: updateGameModel()
                                        }
                                        ToggleButton {
                                            id: blockNecromancer
                                            label: "Necro"
                                            width: 70
                                            height: 35
                                            onClicked: updateGameModel()
                                        }
                                        ToggleButton {
                                            id: blockPaladin
                                            label: "Paladin"
                                            width: 70
                                            height: 35
                                            onClicked: updateGameModel()
                                        }
                                        ToggleButton {
                                            id: blockSorceress
                                            label: "Sorceress"
                                            width: 70
                                            height: 35
                                            onClicked: updateGameModel()
                                        }

                                    }
                                }
                            }
                            
                            Separator{}
                        }

                        
                        // Apply custom maphack settings
                        Item {
                            Layout.preferredWidth: settingsLayout.width
                            Layout.preferredHeight: boxHeight

                            Row {
                                topPadding: 10

                                Column {
                                    width: (settingsLayout.width - applyCustomMaphack.width)
                                    Title {
                                        text: "APPLY CUSTOM MAPHACK SETTINGS"
                                        font.pixelSize: 13
                                    }

                                    SText {
                                        text: "Overwrite your current BH.cfg with the settings above"
                                        font.pixelSize: 11
                                        topPadding: 5
                                        color: "#676767"
                                    }
                                }
                                Column {
                                    id: applyCustomMaphack
                                    width: 60
                                    PlainButton {
                                        id: applyButton
                                        label: "APPLY"
                                        width: 60
                                        font.pixelSize: 10
                                        height: 40

                                        onClicked: {
                                            // Reset error.
                                            errored = false
                                            if(validateGames()) {
                                                var success = settings.persistGameModel()
                                                if(success) {
                                                    bhconfig.getMaphackTextFromParams(0)
                                                    //Temporary hack to pass in a game index of 0, will figure out a proper way to do this later
                                                    //bhconfig.getMaphackTextFromParams(current)
                                                }
                                            } else {
                                                // Show error.
                                                errored = true

                                                // Remove error after a timeout.
                                                errorTimer.restart()
                                            }
                                        }
                                    }
                                } 
                            }
                            
                            Separator{}
                        }

                        // DEP success message.
                        Rectangle {
                            visible: depApplied
                            width: parent.width
                            height: parent.height
                            color: "#00632e"
                            border.width: 1
                            border.color: "#000000"

                            SText {
                                text: "DEP fix successfully applied - don't forget to reboot!"
                                font.pixelSize: 11
                                anchors.centerIn: parent
                                color: "#ffffff"
                            }
                        }

                        // DEP error message.
                        Rectangle {
                            visible: depError
                            width: parent.width
                            height: parent.height
                            color: "#8f3131"
                            border.width: 1
                            border.color: "#000000"

                            SText {
                                text: "There was an error while applying DEP, please try again!"
                                font.pixelSize: 11
                                anchors.centerIn: parent
                                color: "#ffffff"
                            }
                        }
                    }
                }
            }
        }
    }

    Timer {
        id: depAppliedTimer
        interval: 3000; running: false; repeat: false
        onTriggered: depApplied = false
    }

    Timer {
        id: depErrorTimer
        interval: 3000; running: false; repeat: false
        onTriggered: depError = false
    }
}
