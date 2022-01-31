# Recipe Frankenstein 🍜
Find the missing ingredient for your next dish.

Given a meal, Recipe Frankenstein searches through multiple recipes to find what ingredients are commonly used.

## Prerequisites

- Docker
- Go v1.17.3
- Node v16.13.0
- Make

## Running locally
### Run both server and web app
```
$ make start
```
### Server
In `./frank_server` start the containers:
```
$ docker-compose up -d
```
Debugging in VS code, add the following `.vscode/launch.json` file:
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Run Server",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "/home/your_user/dev/projects/recipe-frankenstein/frank_server/cmd/api/main.go"
        }
    ]
}
```
### Web App
Inside of the `web_app` directory:
```
$ npm install
$ npm run --prefix ./web_app serve
```

### Troubleshooting:
- The table must exist in local dynamodb. This can be created by running the cache tests - `go test cache/dynamo`.

## Deployment:
1. Set credentials in .aws directory. 
2. Deploy the server to AWS Lambda with `make deploy-lambda`
3. Deploy the web app to AWS S3 with `make deploy-s3`
