# Go HTTP Unit Tests

This project demonstrates basic tests on server endpoints created using Gorilla MUX.

It also demonstrates the use of build tags to exclude tests.

Last updated: 20 December 2018

## Description

### Endpoint tests

The basic premise of the tests is to create an artificial request and response for the endpoint.

The request simulates the incoming HTTP request.

The response from the server is captured by a "recorder" (`httptest.ResponseRecorder`) that is passed to the MUX Router's `ServeHTTP` method.

### Selective testing

Tests may be selectively run in two ways.

One is by specifying the test name. For example:

```bash
go test -run "Reply"
```

This will run tests containing the word "Reply" in the test functions e.g. `TestReply` and `TestReplyEndpoint`.

The other way is by specifying build tags.

To use this method, the build tag needs to be added to the test files. E.g. `// +build alltests restapi`

The command to run the tests is:

```bash
go test -tags restapi
```

The caveat to using build tags is that tests without an explicit tag will not run. In other words, the plain command `go test` will only run tests with **no** build tags - tests with build tags will not run. This raises the problem of being unable to run all the tests at one go.

The solution is to add a common build tag to all tests e.g. `// +build alltests`

Tests that should not be run with the others can then *not* include that tag.

Test with separate tags can be run together:

```bash
go test -tags "restapi pages"
```

## Dependencies

The following dependencies need to be installed.

```bash
# The router that is used to create the endpoints.
go get github.com/gorilla/mux
# The testing package to reduce test code verbosity.
go get github.com/stretchr/testify/assert
```

## Reference

### Endpoint tests

[https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http/](https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http/)

### Selective testing

[https://stackoverflow.com/questions/24030059/skip-some-tests-with-go-test](https://stackoverflow.com/questions/24030059/skip-some-tests-with-go-test)

Tags: #go #testing #http #buildtags
