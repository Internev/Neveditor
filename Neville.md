##Go specific instructions:##

Make sure you are using Go's recommended folder structure.
Following Go's setup instructions should ensure this but in short it's:
$GOPATH/src/github.com/~github username~/~repo~

Go dependency installation is still a bit of a mystery to me. I think you can just run "go install" from the server folder (or wherever main package is).

If that doesn't work, run 'go get' on:
* github.com/gorilla/mux
* github.com/dchest/uniuri
* github.com/gorilla/websocket

To start server run:
```
go run server/*.go
```
from root.
