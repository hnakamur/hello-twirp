package hellotwirp

//go:generate retool do protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./hello.proto
