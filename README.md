# mockerson
Go project to create easy mock http using json file

Why mockerson because I need a tool to define mocks (mocker) and use json to define i/o.
So removing the J we have Mockerson. XD


# How this work

Use listenandserve and read port and json file from flags
The json file should have paths, method and static responses
Also in json file you can put a file parameter and the code is going to write the file content as output

## How to build
```sh
go build -v -o bin/mockerson mockerson.go
```

## How to run
```sh
./bin/mockerson --port=8080 --json=defaults.json
```
Port and json use by default 8080 and defaults.json so you can run 
```sh
./bin/mockerson
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
        "method": "POST",
        "file": "tyrion.txt",
        "code": 200,
    }...
]
```
# Finally

Probably you could find better solutions or better implementations, probably smaller too.

But finally i had a great time doing this, so I do not care much if you do not like it.

I would like you to leave a comment, with suggestions for improvements or questions at djaque@gmail.com

Feel free to copy, fork, and use as you want, as long as you do not blame me for your problems.
