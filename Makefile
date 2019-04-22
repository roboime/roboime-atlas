setup:
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@cd src/protos/ && ./setup.sh
deps:
	@dep ensure
proto:
	mkdir build
	mkdir build/gen
	protoc --proto_path=./src --go_out=build/gen src/protos/messages_robocup_ssl_wrapper/messages_robocup_ssl_wrapper.proto
	protoc --proto_path=./src --go_out=build/gen src/protos/messages_robocup_ssl_refbox_log/messages_robocup_ssl_refbox_log.proto
	protoc --proto_path=./src --go_out=build/gen src/protos/messages_robocup_ssl_geometry/messages_robocup_ssl_geometry.proto
	protoc --proto_path=./src --go_out=build/gen src/protos/messages_robocup_ssl_detection/messages_robocup_ssl_detection.proto
	protoc --proto_path=./src --go_out=build/gen src/protos/referee/referee.proto
clean:
	@rm -rf build
