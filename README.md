# Pier

A "clientless" implementation of [HTTPie][1] embedded in Go.

Pier converts any implementation of [http.Handler][2] to a CLI tool with the
convenient argument syntax and features of HTTPie.

## Motivation

Go's compile and execute cycle is fast enough that you can practice
"REPL-driven development" using Bash. However, this practice requires authoring
custom CLI tools, which can be tedious. Furthermore, if you're developing an
HTTP service, the work of building a CLI tool feels redundant with all the
work you're doing to build a nice HTTP API. With Pier, no server or
file-watcher is required during iteration.

## Usage

### Install

```bash
go get github.com/nojima/httpie-go
# Use forked version until httpie-go PR #70 lands.
go mod edit '-replace=github.com/nojima/httpie-go=github.com/deref/httpie-go@v0.7.1-0.20210924013123-066db544cf77' 
```

### Code

```go
import (
  "github.com/deref/pier"
  "example.com/your/app"
)

func main() {
  pier.Main(app.Handler)
}
```

Then simply invoke with HTTPie syntax.

There is an example app in this repository. Try this:

```bash
go run ./example /echo message=hello
```

See [https://httpie.io/docs](https://httpie.io/docs) for more detailed documentation.

## Status

Stable enough for development use (which is the intended use anyway!)

Pier is implemented in terms of [a fork][3] of [nojima/httppie-go][4] and
inherits its behavior completely. See the httpie-go readme for details.

## Trivia

Where does the name come from?

In the past the author has referred to CLI tools in this style that bypass the
client/server as operating in "peer" mode. HTTPie + Peer = HTTPier. Shorter is
better and "Pier" seemed to be mostly disused on Github, so here we are.

[1]: https://httpie.io/
[2]: https://golang.org/pkg/net/http/#Handler
[3]: https://github.com/deref/httpie-go
[4]: https://github.com/nojima/httpie-go
