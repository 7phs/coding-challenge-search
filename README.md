## Requirements

1. Go 1.11+

2. Libraries which used in a project will downloaded by ```go mod vendor`` command

## Build

### Manual

A developer should load all dependencies library before building a binary, run command:
```bash
GO111MODULE=on go mode vendor
``` 

Then building an executable:
```bash
go build -o ./search-service
``` 

### Using makefile

Run a command
```bash
make build
```

### Using docker

Run a command
```bash
make image
```

## Testing

### Manual

A developer should load all dependencies library before testing, run command:
```bash
GO111MODULE=on go mod vendor
``` 

Then building an executable:
```bash
go test ./...
``` 

### Using makefile

Run a command
```bash
make testing
```

## Run

### Quick

Run a service using a makefile command
```bash
make run
```

Or using built docker image

```bash
make image-run
```

### Manual

Environment variable:

- **ADDR** - address which a service listening. Default: ':8080'. Example: ':9090;

- **CORS** - switching on response headers supporting CORS. Default: 'false'. Supported: 'true' or 'false'

- **STAGE** - a name of running staging. Default: 'development'. Supported: 'production', 'development', 'testing'

- **LOG_LEVEL** - a level of logging messages. Default: 'info'. Supported: 'debug', 'info', 'warning', 'error'

- **DB_URL** - a path of a SQLite 3 db file. Default: './fatlama.sqlite3'.

- **KEYWORDS_LIMIT** - a maximum length of a keyword line. Default: 4KB.

Command to run:
```bash
STAGE=production ./search-service run 
```

# Implementation description

Implementation of a search services based on a trade-off: source records will never changing after services started.

A primary approaches - in-memory indexes.

No one components are using synchronization, because data and indexes will not changing during service working.  

## High-level architecture

A general architecture of an implemented search engine (**db** directory) included following layers:

1. **Cache**

    A caching layer stores a result of each request even an error.
    It helps prevent repeatedly execution heavy operations.

    A hash of request's parameters is a key for a cache.

2. **Indexes**

    The layer contains a several indexes of source data. 
    It is reindexing source data on initialization and using prepared indexes to quick search a requests's result.
    
    This layer is a good candidate to scaling horizontally.  

3. **Storage**

    The permanent or source storage of data.
    It is using only on initialization to load all data into memory.

## Indexes

A requirements of a search engine will probably growing using new index types. 

The current implementation is 

## Tiles indexes

    

# Trade-off

1. **Permanent data sources**
    
    

2. **Without permanent storage of indexes.** 

    Each time 
  
2. **Never cleaning cache**
