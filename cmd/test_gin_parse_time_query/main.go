package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Queries struct {
	From time.Time `form:"from" time_format:"unix"`
	To   time.Time `form:"to"  time_format:"unix"`
}

func main() {
	r := gin.New()

	r.GET("/test", func(c *gin.Context) {
		q := new(Queries)
		if err := c.ShouldBindQuery(q); err != nil {
			panic(err)
		}

		log.Printf("%+v\n", q)
	})

	r.Run(":9090")
}
