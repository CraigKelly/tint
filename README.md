# tint

A text linter based on excellent projects like proselint, alex, and write-
good. In fact, we have stolen from them liberally.

## Warning

If you go looking at the source code, be warned that we check for lots of
profane and insensitive words.

## Inspiration

While I like prose linters, they often don't do what I want and are SLOW.
This is a personal experiment in writing a fast text linter in Go.

## Some Optimization Work

In order to keep tint as fast as possible, the regular expressions used are as
simple as possible (but no simpler). Here are some optimizations used:

* We keep two copies of a linted file: the original text and a processed
  version. The processed version allows for must faster searching. Warning
  messages quote from the original text. In order to prevent the use of *slow*
  case insensitive regex's, we lowercase the entire file. We also use the
  opportunity to replace string that we won't check with spaces (currently only
  URL's).
* Instead of adding word boundary checking to every regexp that need to
  match entire words, we do post-find boundary checking. See
  `existence_checks.go` for an example.
* Passive voice detection could be accomplished with a single super-complicated
  regex, but instead we use many single word regex's. The match routine for
  passive voice then perform a quick "look behind" to check for the other word
  common in the construction.

## Repo layout

Currently the project is small enough that the command line program and all
code are right here in a single `main` package.

`./scripts` contains build/development scripts. `make` is used as the build
tool. Running `make install` should be sufficient to test, build, and install.
See `./Makefile` for details.

`./res` contains resources used for testing.

`./helper` contains wrappers around `tint` that you might find interesting.
`make install` will create a symbolic link to tintc in `$GOPATH/bin`. However,
you'll need Python 3 and the Python library `colorclass` installed to use it.
(you can install colorclass with pip: `pip3 install --user colorclass`).

## Hacking

This project uses the new [go dep](https://github.com/golang/dep) for
dependency management (though there are few deps).  Use `make update` to update
dependencies.

`make test` runs `./script/test` which runs `go test` but also includes some
manual tests for common errors that have cropped up in the code base.

Running `make lint` will run `go fmt`, `go vet`, `golint`, and `goconst`. `make
lint-install` will install all the lint tools that you need.
