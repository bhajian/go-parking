# go-parking

- export GOPATH=$(go env GOPATH)
- mkdirp ~/go/src/github.com/user/todo
- git clone git@github.com:bhajian/go-parking.git
- cd ~/go/src/github.com/user/todo
- go install ./cmd/parking-app-server/
- $GOPATH/bin/parking-app-server --scheme http

# APIs to call:

- curl http://127.0.0.1:51304/api/lot

- get
- post 
- put, etc

- please see swagger for more info about the API

