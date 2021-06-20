# Pier

A "clientless" implementation of [HTTPie][1] embedded in Go.

## Motivation

Go's compile and execute cycle is fast enough that you can practice
"REPL-driven development" using Bash. However, this practice requires authoring
custom CLI tools, which can be tedious. Furthermore, if you're developing an
HTTP service, the work of building a CLI tool feels redundant with all the
work you're doing to build a nice HTTP API.

Pier converts any implementation of [http.Handler][2] to a CLI tool with
the convenient argument syntax and features of [HTTPie][1].

## Usage

```go
import (
  "github.com/deref/pier"
  "example.com/your/app"
)

func main() {
  pier.Main(app.Handler)
}
```

Then invoke with HTTPie syntax.

```bash
go run example /echo message=hello
```

See [https://httpie.io/docs](https://httpie.io/docs) for arguments syntax.

## Status

Alpha! Only a small part of HTTPie's usage is supported.

## Trivia

Where does the name come from?

In the past the author has referred to CLI tools in this style that bypass the
client/server as operating in "peer" mode. HTTPie + Peer = HTTPier. Shorter is
better and "Pier" seemed to be mostly disused on Github, so here we are.

[1]: https://httpie.io/
[2]: https://golang.org/pkg/net/http/#Handler

```

```
