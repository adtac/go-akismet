# go-akismet

[![GoDoc](https://godoc.org/github.com/adtac/go-akismet/github?status.svg)](https://godoc.org/github.com/adtac/go-akismet)

go-akismet is a Go client library for accessing the [Akismet API](https://akismet.com/development/api/) (v1.1).

### Usage

```go
import "github.com/adtac/go-akismet/akismet"
```

Here's an example if you want to check whether a particular comment is spam or not using the `akismet.Check` method:

```go
akismetKey := "abcdef012345"
isSpam, err := akismet.Check(akismet.Comment{
	Blog: "https://example.com",                 // required
	UserIP: "8.8.8.8",                           // required
	UserAgent: "...",                            // required
	CommentType: "comment",
	CommentAuthor: "Billie Joe",
	CommentAuthorEmail: "billie@example.com",
	CommentContent: "Something's on my mind",
	CommentDate: time.Now(),
}, akismetKey)

if err != nil {
	// There was some issue with the API request. Most probable cause is
	// missing required fields.
}
```

You can also submit false positives (comments that were wrongly marked as spam)
with the `akismet.SubmitHam` method. Or you can submit false negatives (comments
that should be marked as spam, but weren't) with the `akismet.SubmitSpam` method.
Both methods have the same method signature as the `akismet.Check` function: an
`akismet.Comment` structure as the first argument followed by your API key.

### Development

#### Contributing

Patches welcome! This library was primarily developed for use in [Commento](https://github.com/adtac/commento), so priority of features and bug fixes from me will probably follow requirements there.

#### Testing

To run tests, you first need to set up your Akismet API key in the `AKISMET_KEY` environment variable. Following this, run `go test` to automatically execute all tests. Note that this will make real HTTP requests to the Akismet API server.

```bash
$ export AKISMET_KEY=abcdef012345
$ go test
```

### License

```
Copyright 2018 Adhityaa Chandrasekar

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
