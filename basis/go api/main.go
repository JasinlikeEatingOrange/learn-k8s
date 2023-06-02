package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main()  {
   r:=gin.Default()
   r.Handle("GET","/", func(context *gin.Context) {
	    context.String(200,"this is myweb 0.7")
   })
	r.Handle("GET","/user", func(context *gin.Context) {
		context.String(200,"this is user 0.7")
	})

   err:=r.Run(":80")
   if err!=nil{
   	log.Fatal(err)
   }
}
