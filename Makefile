dev:
	@cd src && air server --port 8080

test:
	@go1.19.3 test ./tests/*

clean:
	@rm -rf ./src/tmp && rm -rf ./tmp
