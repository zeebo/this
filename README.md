# This

<p align="center">
  <a href="https://pkg.go.dev/github.com/zeebo/this"><img src="https://img.shields.io/badge/doc-reference-007d9b?logo=go&style=flat-square" alt="go.dev" /></a>
  <a href="https://goreportcard.com/report/github.com/zeebo/this"><img src="https://goreportcard.com/badge/github.com/zeebo/this?style=flat-square" alt="Go Report Card" /></a>
  <a href="https://sourcegraph.com/github.com/zeebo/this?badge"><img src="https://sourcegraph.com/github.com/zeebo/this/-/badge.svg?style=flat-square" alt="SourceGraph" /></a>
  <img src="https://img.shields.io/badge/license-MIT-blue?style=flat-square" />
</p>

This package used to report the calling function's name as cheaply as possible, but has rotted over time due to depending on internal unstable details. It used to have assembly implementations to get the caller pc very cheaply, but I got spooked by my handling of pointers inside of some assembly routines (it used to assume all the pointers were to readonly data, but that may not be the case depending on some internal details). So now it just has some wrappers around runtime.Callers and runtime.FuncForPC and provides no special speedups.
