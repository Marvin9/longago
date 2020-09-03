.PHONY: dev diff

dev:
	go run main.go

diff:
	diff fixtures/100000.csv uploads/100000.csv