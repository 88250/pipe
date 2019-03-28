pangu.go
========

[![](https://img.shields.io/travis/vinta/pangu.svg?style=flat-square)](https://travis-ci.org/vinta/pangu)
[![](https://img.shields.io/badge/made%20with-%e2%9d%a4-ff69b4.svg?style=flat-square)](https://vinta.ws/code/)

Paranoid text spacing for good readability, to automatically insert whitespace between CJK (Chinese, Japanese, Korean) and half-width characters (alphabetical letters, numerical digits and symbols).

- [pangu.go](https://github.com/vinta/pangu) (Go)
- [pangu.java](https://github.com/vinta/pangu.java) (Java)
- [pangu.js](https://github.com/vinta/pangu.js) (JavaScript)
- [pangu.py](https://github.com/vinta/pangu.py) (Python)
- [pangu.space](https://github.com/vinta/pangu.space) (Web API)

## Installation

To install the package, `pangu`, for using in your Go programs:

```console
$ go get -u github.com/vinta/pangu
```

To install the command-line tool, `pangu-axe`:

```console
$ go get -u github.com/vinta/pangu/pangu-axe
```

## Usage

### Package

```go
package main

import (
    "fmt"
    "github.com/vinta/pangu"
)

func main() {
    s := pangu.SpacingText("當你凝視著bug，bug也凝視著你")
    fmt.Println(s)
    // Output:
    // 當你凝視著 bug，bug 也凝視著你
}
```

### Command-line Interface

```console
$ pangu-axe text "與PM戰鬥的人，應當小心自己不要成為PM"
與 PM 戰鬥的人，應當小心自己不要成為 PM

$ pangu-axe file 銀河便車指南.txt
$ pangu-axe file 宇宙盡頭的餐廳.txt -o 宇宙盡頭的餐廳（好讀版）.txt
$ pangu-axe file 生命、宇宙及萬事萬物.txt 再見，謝謝你的魚.txt 基本無害.txt
```

## Documentation

- `pangu` on [GoDoc](https://godoc.org/github.com/vinta/pangu)
- `pangu-axe` on [GoDoc](https://godoc.org/github.com/vinta/pangu/pangu-axe)

Have a question? Ask it on the [GitHub issues](https://github.com/vinta/pangu/issues)!
