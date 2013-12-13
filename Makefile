export GOPATH=$(shell pwd)
install:
	@go install file-server
run:
	@pkill file-server || echo "no file-server process"
	@mkdir -p ~/.logs/file-server/logs/
	@nohup ./bin/file-server 2>&1 > ~/.logs/file-server/logs/file-server.log &
stop:
	@pkill file-server || echo "no file-server process"
test:
	@go install test
	@./bin/test
