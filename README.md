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

Implementation of search services based on a trade-off: source records will never changing after services started.

A primary approaches - in-memory indexes.

No one components are using synchronization, because data and indexes will not be changing during service work.  

## High-level architecture

A general architecture of an implemented search engine (**db** directory) included following layers:

1. **Cache**

    A caching layer stores a result of each request even an error.
    It helps repeatedly to prevent the execution of heavy operations.

    A hash of request's parameters is a key for a cache.

2. **Indexes**

    The layer contains several indexes of source data. 
    It is reindexing source data on initialization and using prepared indexes to quickly search a requests' result.
    
    This layer is an excellent candidate for scaling horizontally.  

3. **Storage**

    The permanent or source storage of data.
    It is using only on initialization to load all data into memory.

## Indexes

Requirements of a search engine will probably be growing the number of the new index types. 

The current implementation allows to growing the number of indexes just changing
a line of code when configures the index list.

The execution of the massive operation is an execution in different goroutines, and then results join and post-processed.

There is just two kind of indexes implementing:

1. **Words indexes** - stored a list of lemmatized keywords and linked records. Using to select a records
   by search terms keywords.

2. **Tiles indexes** - hierarchical geo-oriented tree, stored a sorted records list ordered
   by precalculated distances by the center of each edge tile. 

## Tiles indexes

A primary challenge of this challenge (tautology) is reducing distance calculation, in my opinion.

A brute-force solution is a calculating distance between a user's location and each item's location.

Another way is a point of view on a problem from the side of use cases:

1. A user will find the nearest point of his/her a home, and a home has a square or a "nearest" place, but not just a point.     

2. A user will find inside a strict area.

An idea of tiles indexes based on mapping services principles and items: different scaled tiles.

Searching on big scaling tiles on the first stage and then dive in into smaller scaling tiles
as children of the higher level tiles.


# Trade-off

1. **Permanent data sources**
    
    An implementation doesn't support real-time reindexing of changed data, but possible implementing
    inserting of the new records.
    
    A primary goal of the challenge is solving searching performance bottleneck but
    not working with real-time changed  

2. **Without permanent storage of indexes.** 

    Source data will reindex each time on initializing and spending a little bit of time.
    
    A result of reindexing (data of the two indexes) will not be stored on a drive. As said before
    I was implementing a search engine than its permanent storage. 
  
3. **Never cleaning cache**

    Related point 1. Cache never invalidated because a data are never changing.    
