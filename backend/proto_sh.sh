cd pb && rm chord.pb.go
cd .. && protoc --proto_path=pb pb/*.proto --go_out=plugins=grpc:pb
cd pb && mv chord_backend/pb/chord.pb.go . && rm -R chord_backend
