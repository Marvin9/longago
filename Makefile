.PHONY: check diff clean

clean:
	rm -rf uploads

check:
	mkdir uploads
	go run main.go
	make diff
	make clean

diff:
	diff fixtures/100000.csv uploads/100000.csv