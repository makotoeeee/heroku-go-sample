init:
	heroku create
	git push heroku master

clean:
	rm -rf .git

deploy:
	git push heroku master

local:
	go run main.go

build:
	go build
