# basex-service
Simple golang service that returns the basex conversion from a base10 number


## Example Usage
```
go build -o basex-service
./basex-service

# convert 32 to base 19
curl -X POST http://localhost:1323/convert \
     -H 'Content-Type: application/json' \
     -d '{"base":19,"base10num":32}'
```