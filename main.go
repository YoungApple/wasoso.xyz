package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var api = "https://www.googleapis.com/customsearch/v1"
var key = "AIzaSyATlOUp5Vl8cVBjVHd-EJw7_dPf80q0ahs"
var cx = "011433635156060699376%3Ah8ljdzqzt4q"

func Index(c *gin.Context) {
	//c.String(200, "Hello root.")
}

func SearchHandler(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.String(500, "parameter q=<query> is required.")
		return
	}

	if query == "sex" {
		c.Redirect(301, "/error")
		return
	}

	getUrl := fmt.Sprintf("%s?key=%s&cx=%s&q=%s&alt=json", api, key, cx, query)
	response, err := http.Get(getUrl)
	if err != nil {
		c.Redirect(500, "/error")
		return
	}
	data, err := ioutil.ReadAll(response.Body)
	c.Writer.Write(data)
}

func ErrorHandler(c *gin.Context) {
	c.String(500, "Haha haha error. Can't search this query.")
}

func main() {
	// command line flags
	port := 8888
	//flag.Uint("port", 8888, "port to serve on")
	//dir := flag.String("directory", "web", "directory of web files")
	//flag.Parse()

	router := gin.Default()
	router.StaticFS("/static", http.Dir("web"))

	router.GET("/", Index)
	router.GET("/search", SearchHandler)
	router.GET("/error", ErrorHandler)
	router.Run(fmt.Sprintf("localhost:%d", port))
}
