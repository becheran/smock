# smock

[![Pipeline Status](https://github.com/becheran/smock/actions/workflows/go.yml/badge.svg)](https://github.com/becheran/smock/actions/workflows/go.yml)
[![Go Report Card][go-report-image]][go-report-url]
[![PRs Welcome][pr-welcome-image]][pr-welcome-url]
[![License][license-image]][license-url]

[license-url]: https://github.com/becheran/smock/blob/main/LICENSE
[license-image]: https://img.shields.io/badge/License-MIT-brightgreen.svg
[go-report-image]: https://goreportcard.com/badge/github.com/becheran/smock
[go-report-url]: https://goreportcard.com/report/github.com/becheran/smock
[pr-welcome-image]: https://img.shields.io/badge/PRs-welcome-brightgreen.svg
[pr-welcome-url]: https://github.com/becheran/smock/blob/main/CONTRIBUTING.md

Simple and fast mock generator for golang

![Logo](./docs/logo.png)

## Features

Mocking interfaces for unit tests is a common task in *go* which can be done manually or using a generator which would automates the process of repeatably writing `structs` which fullfil the interfaces for testing.

There are at least two other popular mock generators that exist for *go* right now. The first is *mockgen* which is part of the [mock](https://github.com/golang/mock) module maintained by the *golang* team. The other is [mockery](https://github.com/vektra/mockery) which uses the [testify mock](https://pkg.go.dev/github.com/stretchr/testify/mock) interfaces.

So why "yet another mock generator"? Or in other words "what does *smock* offer that *mockgen* or *mockery* doesn't?"

The intention of *smock* is to simplify the process of manually mocking interfaces. The tool focused on the following features:

- No additional libraries required
- Work with `go:generate` annotation on interfaces without any configuration
- Keep type information when writing `Do` or `Return` blocks of mock object
- Fast parsing and generation
- No complex builtin assertion capabilities. Though, allow them to be added if needed for specific tests

## Getting Started

Install latest version:

``` sh
go install github.com/becheran/smock
```

Annotate `interface` which shall be mocked:

``` go
//go:generate smock
type MockMeIfYouCan interface {
 Foo(bar int, baz string) (res int, err error)
}
```

Run the go generate command in the module root directory next to the `go.mod` file to generate mocks for all annotated interfaces:

``` sh
go generate ./...
```

All generated mocks will appear in the folder `mocks` of the module root. The import name for the generated mocks will be `<PackageNameOfInterface>_mock`.

The mocked interface can be used in unit tests. They have an additional `WHEN` function to set behaviors for each exposed function of the interface. The mock can either `Do` something or `Return` fixed values when a function is called.

The mocks can act like all types of mock objects as [martin fowler described once](https://martinfowler.com/articles/mocksArentStubs.html).

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
 Consumer()
}
```

### Spy

Assert arguments when being called:

``` go
func TestMockMeIfYouCan(t *testing.T) {
 mock := gomod_test_mock.NewMockMockMeIfYouCan(t)
 mock.WHEN().Foo().Do(func(bar int, baz string) (res int, err error) {
  if bar != 42 {
   t.Fatal("bar must be 42")
  }
  return
 })
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
  if ctr > 2 {
   t.Fatal("shall only be called twice")
  }
  return ctr, nil
 })
 Consumer(mock)
}
```
