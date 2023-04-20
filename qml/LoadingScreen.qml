import QtQuick 2.12

Rectangle {
    id: prereqsBox
    anchors.fill: parent
    color: "#030202"

    Item {
        width: 400
        height: 400
        anchors.centerIn: parent

        Column {
            width: parent.width

            // Logtypes always showing.
            Item {
                width: parent.width
                height: 240

                Item {
                    width: 276
                    height: 162
                    anchors.horizontalCenter: parent.horizontalCenter
                    anchors.top: parent.top
                    anchors.topMargin: 109
                    Image { source: "assets/resurgence-logo.png"; anchors.fill: parent; fillMode: Image.PreserveAspectFit; }
                }
            }

            // While loading.
            Column {
                width: parent.width
                visible: (!settings.prerequisitesLoaded && !settings.prerequisitesError)

                Item { 
                    width: parent.width
                    height: 60

                    Title {
                        anchors.horizontalCenter: parent.horizontalCenter
                        anchors.bottom: parent.bottom
                        font.pixelSize: 16
                        color: "#736c6a"
                        text: "CONNECTING TO resurgence API..."
                    }
                }

                Item {
                    width: parent.width
                    height: 60

                    CircularProgress {
                        anchors.horizontalCenter: parent.horizontalCenter
                        anchors.verticalCenter: parent.verticalCenter
                        size: 30
                        visible: true
                    }
                }
            }
            
            // Shows when there's been an error.
            Column {
                width: parent.width
                visible: settings.prerequisitesError

                Item {
                    width: parent.width
                    height: 50

                    Title {
                        anchors.horizontalCenter: parent.horizontalCenter
                        anchors.verticalCenter: parent.verticalCenter
                        font.pixelSize: 14
                        color: "#736c6a"
                        text: "Something went wrong when connecting to the Resurgence API"
                    }
                }

                Item {
                    width: parent.width
                    height: 50

                    PlainButton {
                        width: 100
                        height: 50
                        anchors.horizontalCenter: parent.horizontalCenter
                        anchors.verticalCenter: parent.verticalCenter
                        label: "TRY AGAIN"

                        onClicked: settings.getPrerequisites()
                    }
                }
            }
        }
    }

    OpacityAnimator {
        target: prereqsBox;
        from: 1;
        to: 0;
        duration: 900
        running: (settings.prerequisitesLoaded && !settings.prerequisitesError)
    }

    Item {
        width: 115
        height: 40
        anchors.bottom: parent.bottom
        anchors.right: parent.right
        anchors.rightMargin: 8

        Title {
            text: settings.buildVersion
            font.pixelSize: 10
            anchors.centerIn: parent
        }
    }
}
