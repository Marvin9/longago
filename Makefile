PRETTY:="----------------"

.PHONY: clean
clean:
	@echo $(PRETTY)
	@echo "\nCLEANING PREVIOUS STORAGE DIRECTORY & BINARY\n"
	rm -rf tmp
	rm -rf atlan-collect
	@echo "\n"
	@echo $(PRETTY)

.PHONY: storage
storage:
	@echo $(PRETTY)
	sh ./scripts/storage.sh
	@echo $(PRETTY)

.PHONY: check
check: storage
	go run main.go

.PHONY: test
test: clean
	sh ./scripts/test.sh
	make clean	

.PHONY: docker-compose-up
docker-compose-up: storage
	sudo docker-compose up -d

.PHONY: build
build:
	go build

.PHONY: push
push: clean
	git add .
	git commit -m "$(commit)"
	git config --global credential.helper cache
	git push origin master

# test API manually
# start server by running
# make docker-compose-up
# or 
# make check
.PHONY: start-request
start-request:
	curl -F file=@"./fixtures/100000.csv" localhost:8000/p/start | json_pp

.PHONY: pause-request
pause-request:
	curl -H "content-type: application/json" -X POST -d '{"instance_id": "${instance_id}"}' localhost:8000/p/pause | json_pp

.PHONY: resume-request
resume-request:
	curl -F file=@"./fixtures/100000.csv" -F instance_id="${instance_id}" localhost:8000/p/resume | json_pp

.PHONY: stop-request
stop-request:
	curl -H "content-type: application/json" -X POST -d '{"instance_id": "${instance_id}"}' localhost:8000/p/stop | json_pp