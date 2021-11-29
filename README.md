# Recipe Frankenstein
Find the most common ingredients used in a food dish across many recipes.

## Prerequisites
- Go v1.17.3
- Docker

## Running locally
### Both Services
```
make start
```
### Only Backend
In `./frank_server` start the containers:
```
docker-compose up -d
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
```
npm run --prefix ./web_app serve
```