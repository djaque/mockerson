# mockerson
Go project to create easy mock http using json file

Use listenandserve but read port and json file from args
The json file should have paths ,method and static responses

## How to build
```sh
go build -v -o bin/mockerson mockerson.go
```

## How to run
```sh
./bin/mockerson --port=8080 --json=defaults.json
```
You can change the port and json file but the file should have this format

```json
[
    {
        "path": "/TywinLannister",
        "method": "GET",
        "body": "A lion doesn't concern himself with the opinions of a sheep",
        "code": 200
    },
    {
        "path": "/TyronLannister",
        "method": "GET",
        ...
    }
]
```