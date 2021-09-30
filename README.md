# This

<p align="center">
  <a href="https://pkg.go.dev/github.com/zeebo/this"><img src="https://img.shields.io/badge/doc-reference-007d9b?logo=go&style=flat-square" alt="go.dev" /></a>
  <a href="https://goreportcard.com/report/github.com/zeebo/this"><img src="https://goreportcard.com/badge/github.com/zeebo/this?style=flat-square" alt="Go Report Card" /></a>
  <a href="https://sourcegraph.com/github.com/zeebo/this?badge"><img src="https://sourcegraph.com/github.com/zeebo/this/-/badge.svg?style=flat-square" alt="SourceGraph" /></a>
  <img src="https://img.shields.io/badge/license-MIT-blue?style=flat-square" />
</p>

This is a package to report the calling function's name as cheaply as possible (~20ns with no allocations). It only provides speedups on go1.12+, amd64 and 386 architectures.
