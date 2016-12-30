# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container's workspace.
ADD . $GOPATH/src/github.com/internev/tesis

WORKDIR $GOPATH/src/github.com/internev/tesis/server

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/dchest/uniuri
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/websocket
RUN go install github.com/internev/tesis/server

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/server

# Document that the service listens on port 8000.
EXPOSE 8000
