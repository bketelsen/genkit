##GenKit
This is a proof of concept go generator that generates everything but the `main` package 
to create a service from one of your types.

Given this type:

```
type User struct {
	ID string
	Name string
}
```
to create a go-kit HTTP/JSON service, you simply need to annotate the type
with a comment like this:

```
// @service
type User struct {
	ID string
	Name string
}
```

and add `//go:generate genkit $GOFILE` as the first line of the file containing the type.

GenKit will generate files based on the example services in GoKit's repository.

## Installing

`go get github.com/bketelsen/genkit`

## LICENSE

MIT

based on [gokit](https://gokit.io) which is MIT licensed by Peter Bourgon

