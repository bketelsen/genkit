##GenKit
This is a proof of concept go generator that generates everything but the `main` package 
to create a service from one of your types.

It's slightly opinionated, in that I needed a service that publishes to statsd, so I used the 
statsd metrics package.  Pull requests to make this optional/switchable are welcome.

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

GenKit will generate files based on the example services in GoKit's repository.  All that's missing is a main.go in a parent directory.

Simplified example main :
```
package main
import (
	... some packages ...
	github.com/you/yourapp/user
)
func main() {
	ctx := context.Background()
	http.Handle("/users/", http.StripPrefix("/users", user.GetMux(ctx)))
	listen := fmt.Sprintf(":%d", port)
	err:= http.ListenAndServe(listen, nil)
	...
}
```

## Installing

`go get github.com/bketelsen/genkit`

## LICENSE

MIT

based on [gokit](https://gokit.io) which is MIT licensed by Peter Bourgon

