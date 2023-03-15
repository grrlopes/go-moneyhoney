## Money Honey

App capable to manage the financial status as expenses and income. In
additional, it sends reports by email and generate charts.

## Dependencies:

```
Go => 1.9.3
make => 4.2.1
```

## ENV

```
MONGO_URI="mongodb+srv://user:password@cluster.mongodb.net"
SCHEMA=mastodonte
MODE=prod or debug
PORT=8080
JWTKEY=secret
```

## Run in Production
You could run using binary related to target OS.
https://github.com/grrlopes/go-moneyhoney/releases

## Run the project:
Developement mode could be executed as `make dev`

To run locally, use `make test` to run the tests.
