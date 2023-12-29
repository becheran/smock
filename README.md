# smock

[![Pipeline Status](https://github.com/becheran/smock/actions/workflows/go.yml/badge.svg)](https://github.com/becheran/smock/actions/workflows/go.yml)
[![Doc][go-doc-image]][go-doc-url]
[![Go Report Card][go-report-image]][go-report-url]
[![PRs Welcome][pr-welcome-image]][pr-welcome-url]
[![License][license-image]][license-url]

[license-url]: https://github.com/becheran/smock/blob/main/LICENSE
[license-image]: https://img.shields.io/badge/License-MIT-brightgreen.svg
[go-report-image]: https://goreportcard.com/badge/github.com/becheran/smock
[go-report-url]: https://goreportcard.com/report/github.com/becheran/smock
[pr-welcome-image]: https://img.shields.io/badge/PRs-welcome-brightgreen.svg
[pr-welcome-url]: https://github.com/becheran/smock/blob/main/CONTRIBUTING.md
[go-doc-image]: https://godoc.org/github.com/becheran/smock?status.svg
[go-doc-url]: https://godoc.org/github.com/becheran/smock

Simple and fast mock generator for golang

![Logo](./docs/logo.png)

## Features

Mocking interfaces for unit tests is a common task in *go* which can be done manually or using a generator which would automates the process of repeatably writing `structs` which fullfil the interfaces for testing.

There are at least two other popular mock generators that exist for *go* right now. The first is *mockgen* which is part of the [mock](https://github.com/golang/mock) module which was maintained by the *golang* team, but recently moved to a [*uber* as new maintainer](https://github.com/uber-go/mock). The other is [mockery](https://github.com/vektra/mockery) which uses the [testify mock](https://pkg.go.dev/github.com/stretchr/testify/mock) interfaces.

So why "yet another mock generator", or in other words "what does *smock* offer that *mockgen* or *mockery* doesn't?"

The intention of *smock* is to simplify the process of manually mocking interfaces. The tool focused on the following features:

- Can be used as a [library of a project](#library-tool) or as a [standalone tool](#standalone-tool) using `go generate`
- No boilerplate code such as `gomock.Controller` required to use the mock objects
- Keep type information of input and return parameter when using the test doubles
- Clear and intuitive interface for mocked objects. No unnecessary info provided. For example there is no `Return` expression for mocked functions that do not return a value.
- Fast parsing and generation
- No complex builtin assertion capabilities. Though, allow them to be added if needed for specific tests

## Setup

### Standalone Tool

Install latest version:

``` sh
go install github.com/becheran/smock@latest
```

Annotate `interface` which shall be mocked:

``` go
//go:generate smock
type MockMeIfYouCan interface {
 Foo(bar int, baz string) (res int, err error)
}
```

Run the go generate command in the module root directory next to the `go.mod` file to generate mocks for all annotated interfaces.

### Library Tool

Using *smock* as an installed tool which is the same for all other mocking frameworks has the drawback that mock generation will fail if the tool is not installed on a developer PC as a prerequisite.

Instead of using *smock* as a cli tool it is also possible to add *smock* as a library dependency to a project and still be able to run it via `go generate`.

Add smock as a dependency to your project:

``` sh
go get github.com/becheran/smock
```

Create a new main method which will be used to generate mocks. A recommendation is to put it in the `internal` directory to not expose it to the outside. For example `internal/cmd/smock/main.go`. Add the `go:generate` header to allow this method to be run from the `go generate` command. See the [documentation](https://pkg.go.dev/github.com/becheran/smock/smock) for how the mock generation can be configured:

``` go
package main

import "github.com/becheran/smock/smock"

//go:generate go run ./
func main() {
    smock.GenerateMocks()
}
```

## Generate Mocks

Once *smock* is [setup](#setup) the mock objects can be generated from the module root path:

``` sh
go generate ./...
```

All generated mocks appear in the directory `mocks` next to the corresponding module root path. The import name for the generated mocks will be `<PackageNameOfInterface>_mock`.

A good idea might be to ignore all generated mocks. This can be achieved for example by adding the  following line to your `.gitignore` file:

``` txt
*/**/*_mock
```

## Use Mocked Objects

The mocked interface can be used in unit tests. They have an additional `WHEN` function to set behaviors for each exposed function of the interface. The mock can either `Do` something or `Return` fixed values when a function is called.

The mocks can act like all types of mock objects [described by martin fowler](https://martinfowler.com/articles/mocksArentStubs.html).

### Dummy

Directly pass mock to consumer:

``` go
func TestMockMeIfYouCan(t *testing.T) {
 Consumer(foo_mock.NewMockMockMeIfYouCan(t))
}
```

### Stub

Return fixed answers:

``` go
func TestMockMeIfYouCan(t *testing.T) {
 mock := foo_mock.NewMockMockMeIfYouCan(t)
 mock.WHEN().Foo().Return(42, nil)
 Consumer(mock)
}
```

### Spy

Assert arguments when being called:

``` go
func TestMockMeIfYouCan(t *testing.T) {
 mock := gomod_test_mock.NewMockMockMeIfYouCan(t)
 mock.WHEN().Foo().Expect(match.Eq(42), nil, match.Not(match.Eq("invalid")))
 Consumer(mock)
}
```

### Mock or Fake

Do and return arbitrary stuff when being called:

``` go
func TestMockMeIfYouCan(t *testing.T) {
 mock := gomod_test_mock.NewMockMockMeIfYouCan(t)
 ctr := 0
 mock.WHEN().Foo().Do(func(bar int, baz string) (res int, err error) {
  ctr++
  return ctr, nil
 }).Times(2)
 Consumer(mock)
}
```
