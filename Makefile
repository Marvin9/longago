.PHONY: check diff clean

clean:
	rm -rf uploads

check:
	mkdir -p uploads
	go run main.go
	make clean

diff:
	diff fixtures/100000.csv uploads/100000.csv