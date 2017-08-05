# tint

A text linter based on things like proselint

## Inspiration

While I like prose linters, they often don't do what I want and are SLOW.
This is a personal experiment in writing a fast text linter in Go.

# Some Optimization Work

In order to keep tint as fast as possible, the regular expressions used are
kept as simple as possible (since we have a bunch of them). Here are some
optimizations used:

* A separate "pre-processed" copy of the file text is kept for actual
  searching (the original text is kept for reporting). The only current
  processing is changing the entire file to lower case so that regexps
  don't need to be case-insensitive
* Instead of adding word boundary checking to every regexp that need to
  match entire words, we do post-find boundary checking. See
  `existence_checks.go` for an example.

## Repo layout

Currently the project is small enough that the command line program and all
code are right here in a single `main` package.

`./scripts` contains build/development scripts. `dmk` is used as the build
tool. See `pipeline.yaml` for details.

`./res` contains resources used for testing.

`./helper` contains wrappers around `tint` that you might find interesting.
For instance, if you have Python 3 with colorclass installed (you can install
colorclass with pip: `pip3 install --user colorclass`), then you can put tintc
in your path and the default output will be colored.

## Hacking

This project is so small, I'm using it to play with build tools, so...

* Build with [dmk](https://github.com/CraigKelly/dmk)
* If you need to clean, use the `./script/clean`
* Using the new [go dep](https://github.com/golang/dep) for dependency management (though there are very few deps)

