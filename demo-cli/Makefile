build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t laracom-demo-cli .

run:
#没有后台启动 随时run
	docker run laracom-demo-cli