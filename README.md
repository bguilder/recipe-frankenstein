# Recipe Frankenstein
Find the missing ingredient for your next dish. 🍜

Given a meal, Recipe Frankenstein searches through multiple recipes to find what ingredients are commonly used and what ingredients you might be missing.

## Prerequisites
- Docker
- Go v1.17.3
- Node v16.13.0

## Running locally
### Both Services
```
$ make start
```
### Only Backend
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
### Only Frontend
Inside of the `web_app` directory:
```
$ npm install
$ npm run --prefix ./web_app serve
```

### Troubleshooting:
- The table must exist in local dynamodb. This can be created by running the cache tests - `go test cache/dynamo`.
