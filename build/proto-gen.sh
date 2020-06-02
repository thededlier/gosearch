protoc -I $GOPATH/src --go_out=$GOPATH/src $GOPATH/src/gosearch/internal/proto/domain/search.proto

protoc -I $GOPATH/src --go_out=$GOPATH/src $GOPATH/src/gosearch/internal/proto/service/search-service.proto