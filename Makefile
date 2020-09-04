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
