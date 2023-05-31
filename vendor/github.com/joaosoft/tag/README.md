Tag
================

[![Build Status](https://travis-ci.org/joaosoft/validator.svg?branch=master)](https://travis-ci.org/joaosoft/validator) | [![codecov](https://codecov.io/gh/joaosoft/validator/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/validator) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/validator)](https://goreportcard.com/report/github.com/joaosoft/validator) | [![GoDoc](https://godoc.org/github.com/joaosoft/validator?status.svg)](https://godoc.org/github.com/joaosoft/validator)

A simple struct loader by tags (exported fields only).

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## With support for tags
* "-"

## Dependecy Management
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`


>### Go
```
go get github.com/joaosoft/tag
```

## Usage 
This examples are available in the project at [validator/examples](https://github.com/joaosoft/validator/tree/master/examples)

### Code
```go
import "github.com/joaosoft/tag"

type (
	Person struct {
		Id      int    `person:"id"`
		Name    string `person:"name"`
		Age     int    `person:"age"`
		Book    string `book:"book"`
		Address Address
	}
	Address struct {
		Street     string `address:"street"`
		Number     int    `address:"number"`
		PostalCode string `address:"-"`
	}
)

func main() {
	person := Person{
		Id:   1,
		Name: "joao",
		Age:  30,
		Book: "technology for dummies",
		Address: Address{
			Street:     "apple street",
			Number:     1,
			PostalCode: "1234-567",
		},
	}

	values := tag.Load(person, "person", "book")
	fmt.Println("Person:")
	for key, value := range values {
		fmt.Printf("Key: %s, Value: %+v\n", key, value)
	}

	values = tag.Load(person, "address")
	fmt.Println("\nAddress:")
	for key, value := range values {
		fmt.Printf("Key: %s, Value: %+v\n", key, value)
	}
}
```

> ##### Response:
```
Person:
Key: id, Value: 1
Key: name, Value: joao
Key: age, Value: 30
Key: book, Value: technology for dummies

Address:
Key: street, Value: apple street
Key: number, Value: 1
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
