package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IndexPageData struct {
	PageTitle string
}

type corgi struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Color string `json:"color"`
	Age   int    `json:"age"`
	// Photo
}

// albums slice to seed record album data.
var corgis = []corgi{
	{ID: "1", Name: "Pembroke", Breed: "Pembroke Welsh Corgi", Color: "bicolor", Age: 1},
	{ID: "2", Name: "Cardigan", Breed: "Cardigan Welsh Corgi", Color: "tricolor", Age: 1},
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("static/html/index.html")
	router.Static("/static", "./static/")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/corgis", getCorgis)
	router.GET("/corgis/:id", getCorgiByID)
	// router.GET("/corgis/random", getCorgiRandom)
	router.POST("/corgis", postCorgis)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// getCorgis responds with the list of all corgis as JSON.
func getCorgis(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, corgis)
}

// postCorgis adds an corgi from JSON received in the request body.
func postCorgis(c *gin.Context) {
	var newCorgi corgi

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newCorgi); err != nil {
		return
	}

	// Add the new album to the slice.
	corgis = append(corgis, newCorgi)
	c.IndentedJSON(http.StatusCreated, newCorgi)
}

// getCorgiByID locates the corgi whose ID value matches the id
// parameter sent by the client, then returns that corgi as a response.
func getCorgiByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of corgis, looking for
	// an corgi whose ID value matches the parameter.
	// Rewrite in binary search
	for _, a := range corgis {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "corgi with this ID not found"})
}

//func getCorgiRandom() {

// }
