frsh: lex.go main.go parse.y vars.go
	goyacc parse.y
	go build -o frsh

clean:
	@go clean
	@rm -f y.go y.output

.PHONY: clean
