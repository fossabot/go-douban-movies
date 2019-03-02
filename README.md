# go-douban-movies

[![Build Status](https://travis-ci.org/linehk/go-douban-movies.svg?branch=master)](https://travis-ci.org/linehk/go-douban-movies)
[![Go Report Card](https://goreportcard.com/badge/github.com/linehk/go-douban-movies)](https://goreportcard.com/report/github.com/linehk/go-douban-movies)

English | [简体中文](./README-zh_CN.md "简体中文")

go-douban-movies returns the content of [Douban Movie Top 250](https://movie.douban.com/top250 "Douban Movie Top 250") with indented JSON format by the crawler.

## Installation

```bash
git clone https://github.com/linehk/go-douban-movies.git
```

Or:

```bash
go get -u github.com/linehk/go-douban-movies
```

Then, build it:

```bash
go build
```

## Usages

1. Get all:`GET: http://localhost:8888/api/v1/movies`
2. Get partially：`GET: http://localhost:8888/api/v1/movies?s=100&e=200`

## License

[MIT License](./LICENSE "MIT License")