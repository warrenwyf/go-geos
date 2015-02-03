go-geos
=========

It is a Go binding for [GEOS](http://trac.osgeo.org/geos/)

[![CI Status](https://travis-ci.org/warrenwyf/go-geos.png?branch=master)](https://travis-ci.org/warrenwyf/go-geos)


Install
-------

### Requirements

 * GEOS Library installed
 * Get source code with command `go get github.com/warrenwyf/go-geos`
 * (Optional) Change `CFLAGS` and `LDFLAGS` in source code to your library path if necessary


#### GEOS library on Mac

```bash
$ brew install geos
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