
## Prerequisites
[GoLang](https://golang.org/dl/) <br>

## Development setup
### Project setup
Set the go path in ~/.bash_profile
```
vi ~/.bash_profile
```

```
export GOPATH="<desired path>"
eg: export GOPATH="/Users/me/go"
```
Create the following folder structure from the go path
```
cd <go path>
eg:cd /Users/me/go
```

### Build and Run Locally
Go to the git repo and run dep to fetch vendor libraries
```
dep ensure
```
Build project
```
go build
```
Run project
```
./ic-indexer-service
```

### Elastic search & Kibana setup

    
```
Install elasticsearch-6.5.1
kibana-6.5.1-darwin-x86_64
```
Run elastic seatch
```
elasticsearch-6.5.1/bin/elasticsearch
```
Run kibana
```
kibana-6.5.1-darwin-x86_64/bin/kibana
```

Please refer `elastic_search.go` to view the mappings for icecream

### Developer Notes

## Swagger

The project is configured with Swagger to leverage the API documentation and also helpful to sync up with icecream-indexer-service from icecream-indexer-worker.

### JWT Authentication

This service is authenticated by JWT token. Please use the below login when making request to the system.

```
 username: zalora
 password: zalora
```


### GET - Authenticated

This service solely responsible for all the `GET` requests. 

### PUT, DELETE  - unAuthenticated

These methods are for sync up to elastic search during create or update or delete requests to maintain the data uniformity.

### TestCases:

Test cases are written and it has 100% of test coverage and code coverage. It includes both *unit test cases*.

## Environment based Config files

The project uses config files based on the following environments

    - dev
    - staging
    - uat
    
## Configuration

The configurations need to run the project 

    - server Port Number
    - elastic search configuration
        
