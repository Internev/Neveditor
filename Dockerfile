# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/internev/tesis

WORKDIR /go/src/github.com/internev/tesis/server

# Build the server command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/dchest/uniuri
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/websocket
RUN go install github.com/internev/tesis/server

# Run the server command by default when the container starts.
ENTRYPOINT /go/bin/server

# Document that the service listens on port 8000.
EXPOSE 8000
