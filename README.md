go-geos
=========

It is a Go binding for [GEOS](http://trac.osgeo.org/geos/)

[![CI Status](https://travis-ci.org/warrenwyf/go-geos.png?branch=master)](https://travis-ci.org/warrenwyf/go-geos)


Install
-------

# Requirements

 * GEOS Library installed
 * Get source code with command `go get github.com/warrenwyf/go-geos`
 * (Optional) Change `CFLAGS` and `LDFLAGS` in source code to your library path if necessary


## Install GEOS library on Mac

```bash
$ brew install geos
```


## Install GEOS library on Ubuntu

```bash
$ sudo apt-add-repository -y ppa:ubuntugis/ubuntugis-unstable
$ sudo apt-get update
$ sudo apt-get install libgeos-dev libgeos-3.4.2
```


Quick Start
-----------

```go
package main

import (
	"github.com/warrenwyf/go-geos/geos"
)

func main() {
	geom := geos.CreatePoint(0, 0)
	wkt := geom.Buffer(10).ToWKT()
}
```