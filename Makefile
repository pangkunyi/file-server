export GOPATH=$(shell pwd)
gc=GODEBUG='gctrace=1'

install:
	@go install file-server
run: stop
	@mkdir -p ~/.logs/file-server/logs/
	@${gc} nohup ./bin/file-server 2>&1 > ~/.logs/file-server/logs/file-server.log &
test: stop
	@${gc} ./bin/file-server
stop:
	@pkill file-server || echo "no file-server process"
