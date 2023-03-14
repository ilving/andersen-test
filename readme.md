# Test.

## Conditions

### Prime Number Tester

Build API to check if a list of given numbers are primes.
API should only have a single POST endpoint to accept the request.
The request body must be a slice of integers, otherwise, return an error.
If the request is valid, return a slice of booleans either the number was prime
or not.
#### Examples
Success Example:
POST / 
```json
[2,3,4,5]
```
```json
[true, true, false, true]
```
Error Example:
POST /
```json
[2, 3, false, "nan"]
```
```json
{
  "Error": "error message"
}
```

## How to build and run

```bash
docker-compose up -d
```
OR build 
```bash
go build -a -installsuffix cgo -o main ./cmd/app/main.go
```

You can set web server port by set APP_PORT env variable. If it's not set server runs on default port 8081

Example request by curl

```bash
curl -XPOST http://localhost:8081 -d "[1,2,3,4,5,6,7]"
```

## Description

I use additional packages - fasthttp and fashttp-router (see go.mod). It's a fastest golang webserver with fastest router.

For find prime numbers I use math/big.Int{}.ProbablyPrime function, this function uses two most popular and  
fastest algos for check is number is prime
([Miller-Rabin tests](https://en.wikipedia.org/wiki/Miller%E2%80%93Rabin_primality_test) and [Baillie-PSW test](https://en.wikipedia.org/wiki/Baillie%E2%80%93PSW_primality_test)).
For count of rounds I use log2 of number (As I found it's enough for be sure that this tests work fine)

I didn't divide runtime to different goroutines because this algos utilize CPU, and that's why we can lose more time for switching context.



