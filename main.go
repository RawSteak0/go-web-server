package main

import (
       "fmt"
       "github.com/labstack/echo"
//       "github.com/spf13/cobra"
       "net/http"
       "os"
//       "path/filepath"
//       "mime"
       "strings"
)


func main() {

     var frontOrigin string = "front_root"
     var defaultURL = "/index.html"

     inet := echo.New()

     inet.GET("/*", func (ctx echo.Context) error {
          req := ctx.Request()
          var url string = req.URL.Path
          if(url == "/"){
                 url = defaultURL
          }
          url = strings.Replace(url, "../", "" , -1)
          _, err := os.ReadFile(frontOrigin + url)
          if(err != nil){
                 fmt.Println("GET returned error: " + err.Error())
                 return ctx.String(http.StatusNotFound, "error: cannot retrieve file " + url)
          }
          return ctx.File(frontOrigin + url)
          //return ctx.HTML(http.StatusOK, "HTTP/1.0 200 OK\nServer: go\n, Content-type: " + mimeType + "\n\n" + dataStr)
     })

fmt.Println("server setup")

inet.Start(":8080")


}