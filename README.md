# go-geos
=========
It is a Go binding for [GEOS](http://trac.osgeo.org/geos/)


Install
-------

### Requirements

 * GEOS Library installed
 * Get source code with command `go get github.com/warrenwyf/go-geos`
 * Change `CFLAGS` and `LDFLAGS` in source code to your library path


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