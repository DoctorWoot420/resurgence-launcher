import QtQuick 2.12
import QtQuick.Layouts 1.3

Item {
    id: topbar
    property string activeMenuItem: "launch"
    property bool settingsHovered: false
    property var menuSources: { 
        "launch": "LauncherView.qml",
        "armory": "ArmoryView.qml",
        "discord": "CommunityView.qml"
    }

    // Main menu.
    Item {
        width: 330
        height: parent.height
        anchors.left: parent.left
        anchors.leftMargin: 20
        
        RowLayout {
            id: mainMenu
            anchors.fill: parent
            Layout.alignment: Qt.AlignHCenter | Qt.AlignVCenter
            spacing: 6

            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 100
                
                MenuItem {
                    text: "LAUNCH"
                    active: (activeMenuItem == "launch")

                    onClicked: function() {
                        activeMenuItem = "launch"
                        contentLoader.source = menuSources.launch
                    }
                }
            }

            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 130
                
                MenuItem {
                    width: 110
                    text: "ARMORY"
                    active: (activeMenuItem == "armory")

                    onClicked: function() {
                        Qt.openUrlExternally("https://rustyshackleford1888.github.io/")
                    }

                    Image {
                        id: linkoutIcon
                        fillMode: Image.Pad
                        anchors.top: parent.top
                        anchors.right: parent.right
                        width: 16
                        height: 16
                        source: "assets/icons/out.png"
                        opacity: 0.2
                    }
                }
            }

            Item {
                Layout.alignment: Qt.AlignRight | Qt.AlignVCenter
                height: parent.height
                width: 130
                
                MenuItem {
                    width: 110
                    text: "DISCORD"
                    active: (activeMenuItem == "discord")

                    onClicked: function() {
                        Qt.openUrlExternally("https://discord.gg/hsagMy9d")
                    }

                    Image {
                        id: linkoutIcon
                        fillMode: Image.Pad
                        anchors.top: parent.top
                        anchors.right: parent.right
                        width: 16
                        height: 16
                        source: "assets/icons/out.png"
                        opacity: 0.2
                    }
                }
            }
        }
    }

    // Settings.
    Item {
        width: 120; height: parent.height
        anchors.right: parent.right
        anchors.rightMargin: 20

        Item {
            width: 120
            height: 20
            anchors.verticalCenter: parent.verticalCenter
            anchors.horizontalCenter: parent.horizontalCenter

            Image {
                id: optionsIcon
                fillMode: Image.Pad
                anchors.verticalCenter: parent.verticalCenter
                anchors.right: parent.right
                width: 16
                height: 16
                source: "assets/icons/settings.png"
                opacity: settingsHovered ? 1.0 : 0.5

                MouseArea {
                    anchors.fill: parent
                    hoverEnabled: true
                    cursorShape: Qt.PointingHandCursor
                    onClicked: settingsPopup.open()
                    onEntered: {
                        settingsHovered = true
                    }
                    onExited: {
                        settingsHovered = false
                    }
                }
            }
        }
    }

    Separator{
        color: "#161616"
    }
}
