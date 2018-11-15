COMMAND_NAME=donezo
BIN_OUTPUT=dist/${COMMAND_NAME}

build:
	go build -o ${BIN_OUTPUT}

install: build
	cp ${BIN_OUTPUT} ${HOME}/go/bin/${COMMAND_NAME}

