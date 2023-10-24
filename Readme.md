# URL Shortener

A tiny CLI to understand the nuances of Golang while working with some third party libraries.

# Build

```
go build -o bin/URLShortener
```

## Run

```
go run . "https://link.com/abc" "store" # To generate and store a tiny url
go run . "tinyUrl" "fetch" # To fetch the destination URL
# Or
go build -o bin/URLShortener
./bin/URLShortener <command line flags>
```
