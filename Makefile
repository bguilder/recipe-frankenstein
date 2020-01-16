server: 
	cd ./frank_server; go run ./cmd/api/main.go
	
app:
	npm run --prefix ./web_app serve