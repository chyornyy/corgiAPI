package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type IndexPageData struct {
	PageTitle string
}

type Corgi struct {
	ID          string `json:"id"`
	Author      string `json:"author"`
	Name        string `json:"name"`
	Breed       string `json:"breed"`
	Color       string `json:"color"`
	Age         int    `json:"age"`
	Kennel      string `json:"kennel"`
	Description string `json:"description"`
	Photo       string `json:"photo_url"`
	Likes       int    `json:"total_likes"`
	//	Comments     []Comment `json:"comments"`
	DateCreated  string `json:"date_created"`
	DateModified string `json:"date_modified"`
}

type Comment struct {
	User      string `json:"user"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

// corgis slice to seed record cori data.
var corgis = []Corgi{
	{ID: "0", Author: "admin", Name: "Alex", Breed: "Pembroke Welsh Corgi", Color: "tricolor", Age: 5, Description: "", Photo: "https://storage.yandexcloud.net/corgiapi-bucket/corgi_photos/corgi_alex_1.HEIC", Likes: 0, DateCreated: "", DateModified: ""},
	{ID: "1", Author: "admin", Name: "Alex", Breed: "Pembroke Welsh Corgi", Color: "tricolor", Age: 5, Description: "", Photo: "https://storage.yandexcloud.net/corgiapi-bucket/corgi_photos/corgi_alex_2.HEIC", Likes: 0, DateCreated: "", DateModified: ""},
	{ID: "2", Author: "admin", Name: "Alex", Breed: "Pembroke Welsh Corgi", Color: "tricolor", Age: 5, Description: "", Photo: "https://storage.yandexcloud.net/corgiapi-bucket/corgi_photos/corgi_alex_3.HEIC", Likes: 0, DateCreated: "", DateModified: ""},
	{ID: "3", Author: "admin", Name: "Alex", Breed: "Pembroke Welsh Corgi", Color: "tricolor", Age: 5, Description: "", Photo: "https://storage.yandexcloud.net/corgiapi-bucket/corgi_photos/corgi_alex_4.HEIC", Likes: 0, DateCreated: "", DateModified: ""},
	{ID: "4", Author: "admin", Name: "Alex", Breed: "Pembroke Welsh Corgi", Color: "tricolor", Age: 5, Description: "", Photo: "https://storage.yandexcloud.net/corgiapi-bucket/corgi_photos/corgi_alex_5.HEIC", Likes: 0, DateCreated: "", DateModified: ""},
	{ID: "5", Author: "admin", Name: "Alex", Breed: "Pembroke Welsh Corgi", Color: "tricolor", Age: 5, Description: "", Photo: "https://storage.yandexcloud.net/corgiapi-bucket/corgi_photos/corgi_alex_6.HEIC", Likes: 0, DateCreated: "", DateModified: ""},
	{ID: "6", Author: "admin", Name: "Alex", Breed: "Pembroke Welsh Corgi", Color: "tricolor", Age: 5, Description: "", Photo: "https://storage.yandexcloud.net/corgiapi-bucket/corgi_photos/corgi_alex_7.HEIC", Likes: 0, DateCreated: "", DateModified: ""},
	{ID: "7", Author: "admin", Name: "Alex", Breed: "Pembroke Welsh Corgi", Color: "tricolor", Age: 5, Description: "", Photo: "https://storage.yandexcloud.net/corgiapi-bucket/corgi_photos/corgi_alex_8.HEIC", Likes: 0, DateCreated: "", DateModified: ""},
	{ID: "8", Author: "admin", Name: "Corgi Viking", Breed: "Pembroke Welsh Corgi", Color: "tricolor", Age: 3, Description: "", Photo: "https://storage.yandexcloud.net/corgiapi-bucket/corgi_photos/corgi_midjourney_viking_1.PNG", Likes: 0, DateCreated: "", DateModified: ""},
	{ID: "9", Author: "admin", Name: "Corgi Warrior", Breed: "Pembroke Welsh Corgi", Color: "bicolor", Age: 3, Description: "", Photo: "https://storage.yandexcloud.net/corgiapi-bucket/corgi_photos/corgi_midjourney_warrior_1.PNG", Likes: 0, DateCreated: "", DateModified: ""},
	{ID: "10", Author: "admin", Name: "Corgi Warrior", Breed: "Pembroke Welsh Corgi", Color: "bicolor", Age: 3, Description: "", Photo: "https://storage.yandexcloud.net/corgiapi-bucket/corgi_photos/corgi_midjourney_warrior_2.PNG", Likes: 0, DateCreated: "", DateModified: ""},
	{ID: "11", Author: "admin", Name: "Corgi Warrior", Breed: "Pembroke Welsh Corgi", Color: "bicolor", Age: 3, Description: "", Photo: "https://storage.yandexcloud.net/corgiapi-bucket/corgi_photos/corgi_midjourney_warrior_3.PNG", Likes: 0, DateCreated: "", DateModified: ""},
	{ID: "12", Author: "", Name: "Cardigan", Breed: "Cardigan Welsh Corgi", Color: "tricolor", Age: 1, Description: "", Photo: "", Likes: 0, DateCreated: "", DateModified: ""},
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
	router.GET("/corgis/random", getCorgiRandom)
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
	var newCorgi Corgi

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
	// ! Rewrite in binary search
	for _, a := range corgis {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "corgi with this ID not found"})
}

// Generates random ID in range len(corgis)
// and returns a .json with this ID
func getCorgiRandom(c *gin.Context) {
	random_id := rand.Intn(len(corgis))
	id := strconv.Itoa(random_id)

	for _, a := range corgis {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "corgi with this ID not found"})
}
