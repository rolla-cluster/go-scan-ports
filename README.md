# go-scan-ports
A quick and dirty concurrent Golang-based port scanner, this will scan ports 1 through 1024

## Usage:
Requires 1 command line argument of URL (FQDN or IP)
eg:
```
go run main.go <URL>
```

Example response:
```
Initializing port scanner version:  1.0.1
--- Rolla-cluster ---
22 open
80 open
```
