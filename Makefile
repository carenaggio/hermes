BINARY_NAME=hermes
 
all: ${BINARY_NAME} test
 
${BINARY_NAME}:
	go build -o ${BINARY_NAME} *.go
 
run: ${BINARY_NAME}
	PORT=8081 HERMES_CONFIG=./etc/config.json ./${BINARY_NAME}

image:
	podman build -t ghcr.io/carenaggio/hermes .

image-push: image
	podman push ghcr.io/carenaggio/hermes

clean:
	go clean
	rm -f ${BINARY_NAME}
