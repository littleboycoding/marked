build:
	cd ./web && npm run build
	go build -o ./dist/marked ./main.go
