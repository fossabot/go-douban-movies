# go-douban-movies

[![Build Status](https://travis-ci.org/linehk/go-douban-movies.svg?branch=master)](https://travis-ci.org/linehk/go-douban-movies)
[![codecov](https://codecov.io/gh/linehk/go-douban-movies/branch/master/graph/badge.svg)](https://codecov.io/gh/linehk/go-douban-movies)
[![Go Report Card](https://goreportcard.com/badge/github.com/linehk/go-douban-movies)](https://goreportcard.com/report/github.com/linehk/go-douban-movies)

[English](./README-en.md "English") | 简体中文

go-douban-movies 通过爬虫将 [豆瓣电影 Top 250](https://movie.douban.com/top250 "豆瓣电影 Top 250") 的内容通过带缩进的 JSON 格式返回给客户端。

## 安装

```bash
git clone https://github.com/linehk/go-douban-movies.git
```

然后进行编译：

```bash
go build -o go-douban-movies
```

再运行：

```bash
./go-douban-movies
```

## 使用

1. 获取全部：GET `http://localhost:8888/api/v1/movies`
2. 获取部分：GET `http://localhost:8888/api/v1/movies?s=100&e=200`

## 开源许可证

[MIT License](./LICENSE "MIT License")