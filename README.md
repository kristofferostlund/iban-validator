# iban-validator

## Assignment

In the programming language of your choice, create a webserver that responds to a REST API call.
You should implement one endpoint, which validates if an [IBAN][1] number is valid or not.

Don't worry about making a perfect solution to this problem: no need to solve every corner case. We estimate that this should take you about 4 to 8 hours of programming time.

[1]: https://en.wikipedia.org/wiki/International_Bank_Account_Number

## Usage

### Server

The web server can be started using either `go run cli/cli.go -port=5000 -address=localhost` or by building a binary and executing it like so:

````bash
go build -o webserver cli/cli.go \
  && chmod +x ./webserver \
  && ./webserver -port=5000 -address=localhost
```

Then to make a request to the server, one could curl it from another terminal:

```bash
curl -X POST localhost:5000/iban/validate \
  -H 'Content-Type: application/json' \
  --data '{"iban": "GB82 WEST 1234 5698 7654 32"}'
# {"isValid":true,"message":"IBAN is valid"}
```

### Tests

To run the tests, from root run `go tests -v ./...`
