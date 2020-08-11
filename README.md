# Portobello
A marine ports database.

## Running in Docker
```
docker-compose up
```

## Running manually
Execute the command below in both `PortService` & `PortClient` folders:
```
set -a && . .env && go run main.go
```

## Running tests
```
run/test
```

## API
[/ports/PORT_ID](http://localhost/ports/UAODS)

Returns port data for a port stored under the PORT_ID key in the `ports.json` file.

[/import](http://localhost/import)

Initiates data import. When visiting repeatedly while import is running, it reports the current import progress. You may need to hit `F5` repeatedly & quickly in order to see that because Go is ⚡fast.
