.PHONY: check diff clean build push

clean:
	rm -rf tmp
	rm -rf atlan-collect

check:
	mkdir -p ./tmp/uploads
	go run main.go

diff:
	diff fixtures/100000.csv ./tmp/uploads/100000.csv

build:
	go build

push: clean
	git add .
	git commit -m "$(commit)"
	git config --global credential.helper cache
	git push origin master

start-request:
	curl -F file=@"./fixtures/100000.csv" localhost:8000/p/start | json_pp

pause-request:
	curl -H "content-type: application/json" -X POST -d '{"instance_id": "${instance_id}"}' localhost:8000/p/pause | json_pp

resume-request:
	curl -F file=@"./fixtures/100000.csv" -F instance_id="${instance_id}" localhost:8000/p/resume | json_pp

stop-request:
	curl -H "content-type: application/json" -X POST -d '{"instance_id": "${instance_id}"}' localhost:8000/p/stop | json_pp