run: build
	@./bin/app

build:
	@go build -o bin/app .

templ:
	@templ generates

git:
	@git add .
	@git commit -m "update"
	@git push