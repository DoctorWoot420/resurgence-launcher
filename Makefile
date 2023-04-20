# Use run if there's only changes to the QML layer.
run:
	./deploy/darwin/resurgence-launcher.app/Contents/MacOS/resurgence-launcher

# Rebuild the app if you made changes to the Go layer.
build:
	qtdeploy build