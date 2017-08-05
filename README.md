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

`./scripts` contains build/development scripts. `dmk` is used as the build
tool. See `pipeline.yaml` for details.

`./res` contains resources used for testing.

`./helper` contains wrappers around `tint` that you might find interesting.
For instance, if you have Python 3 with colorclass installed (you can install
colorclass with pip: `pip3 install --user colorclass`), then you can put tintc
in your path and the default output will be prettified in your terminal.

## Hacking

This project is so small, I'm using it to play with build tools, so...

* Build with [dmk](https://github.com/CraigKelly/dmk)
* If you need to clean, use the `./script/clean`
* Using the new [go dep](https://github.com/golang/dep) for dependency management (though there are few deps)
