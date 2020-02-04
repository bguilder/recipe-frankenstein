deploy-lambda:
	cd ./frank_server/cmd/lambda; GOOS=linux GOARCH=amd64 go build main.go; zip function.zip main; aws lambda update-function-code --function-name test --zip-file fileb://function.zip

deploy-s3:
	cd ./web_app; npm run build; aws s3 sync ./dist s3://recipefrankenstein.com

server: 
	cd ./frank_server; go run ./cmd/api/main.go
	
app:
	npm run --prefix ./web_app serve