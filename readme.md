# util for use gin web frame

## Usage

```go 
package main

import (
	"log"

	"github.com/cncsmonster/gin-util/responsewriter"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		bsw := responsewriter.NewBufferedResponseWriter(ctx.Writer)
		ctx.Writer = bsw
		ctx.Next()
		ctx.Abort()
		ctx.Writer = bsw.Base()
		data := bsw.Data()
		log.Println("data:", string(data))
		ctx.Writer = bsw.Base()
	})
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello world")
	})
}

```

## License

THIS PROJECT IS LICENSED UNDER THE MIT LICENSE, WHICH CAN BE FOUND IN THE [LICENSE](./LICENSE) DOCUMENT.