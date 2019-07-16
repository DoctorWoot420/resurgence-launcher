import QtQuick 2.12
import QtQuick.Layouts 1.3          // RowLayout

Rectangle {
    color: "#09030a"

    RowLayout {
        id: settingsLayout
        anchors.fill: parent
        spacing: 8

        // Left column.
        Item {
            Layout.fillWidth: true
            Layout.minimumWidth: 300
            Layout.preferredWidth: 300
            Layout.maximumWidth: 300
            Layout.fillHeight: true

            SText {
                text: "GAMES"
                anchors.top: parent.top
                anchors.left: parent.left
                anchors.topMargin: 20
                font.pixelSize: 15
                font.bold: true
                leftPadding: 15
            }
            
            ListView {
                width: parent.width - 15;
                height: 500
                anchors.top: parent.top
                anchors.right: parent.right
                anchors.topMargin: 50

                model: GameModel {}
                delegate: SettingsDelegate{}

                onCurrentItemChanged: {
                    gameSettings.game = model.get(this.currentIndex)
                }
            }
        }
        
        // Right column.
        Item {
            Layout.fillWidth: true
            Layout.fillHeight: true
            
            SText {
                text: "SETTINGS"
                anchors.top: parent.top
                anchors.left: parent.left
                anchors.topMargin: 20
                font.pixelSize: 15
                font.bold: true
                leftPadding: 20
            }

            GameSettings {
                id: gameSettings
                anchors.left: parent.left
                anchors.top: parent.top
                anchors.topMargin: 49
                anchors.horizontalCenter: parent.horizontalCenter
            }
        }
    }

    // Close button.
    Item {
        width: 30
        height: 30
        anchors.top: parent.top
        anchors.right: parent.right
        anchors.rightMargin: 20

        Image {
            id: closeSettings
            fillMode: Image.PreserveAspectFit
            anchors.verticalCenter: parent.verticalCenter
            anchors.left: parent.left
            width: 30
            height: 30
            source: "assets/svg/close.svg"
        }

        MouseArea {
            anchors.fill: parent
            cursorShape: Qt.PointingHandCursor
            onClicked: stack.pop()
        }
    }
}