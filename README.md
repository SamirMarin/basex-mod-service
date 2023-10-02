# basex-service
Simple golang service that converts a base10 number into a baseX number. 


## Example Usage
```
go build -o basex-service
./basex-service

# convert 32 to base 19
curl -X POST http://localhost:1323/convert \
     -H 'Content-Type: application/json' \
     -d '{"base":19,"base10num":32}'
```


