gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:p

clean:
	rm pb/*.pb

run:
	go run main.go

