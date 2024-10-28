package main 

import (
  "context"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type URL struct {
  Url string `bson:"url"`
  ShortUrl string `bson:"shortUrl"`
}

var client *mongo.Client
var urlColl *mongo.Collection



func main()  {
   r := gin.Default()

  if err := godotenv.Load(); err != nil {
    log.Fatal("Failed to load .env")
  }

   uri := os.Getenv("MONGODB_URI")
   if uri == "" {
	log.Fatal("MONGODB URI is empty")
   }

   client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
   if err != nil {
	panic(err)
   }

   defer func ()  {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
   }()


   urlColl = client.Database("url-shortner").Collection("urls")


   // load static files
   r.StaticFile("/favicon.ico", "./static/favicon.ico")
   r.LoadHTMLFiles("static/form.html")

   r.GET("/", func(c *gin.Context) {
	c.HTML(200, "form.html",nil)
   })

   r.GET("/:shortUrl", getUrl)
   r.POST("/submit", handleForm)

  r.Run(":4000")
}


func handleForm(c *gin.Context) {
	var newUrl URL
	newUrl.Url = c.PostForm("url")
	newUrl.ShortUrl = generateRandomString(6)

	// Attempt to insert the new URL document
	_, err := urlColl.InsertOne(context.TODO(), newUrl)
	if err != nil {
		// Check if the error is due to a duplicate `ShortUrl`
		if mongo.IsDuplicateKeyError(err) {
			// Find the existing document by URL
			var existingUrl URL
			findErr := urlColl.FindOne(context.TODO(), bson.M{"url": newUrl.Url}).Decode(&existingUrl)
			if findErr != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve existing URL"})
				return
			}

			// Return the existing short URL if it already exists
			c.JSON(http.StatusOK, gin.H{"shortUrl": "http://localhost:4000/" + existingUrl.ShortUrl})
			return
		}

		// For other errors, return an internal server error
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL"})
		return
	}

	// If no error, return the new short URL
	c.JSON(http.StatusOK, gin.H{"shortUrl": "http://localhost:4000/" + newUrl.ShortUrl})
}


func getUrl(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	if shortUrl == "" {
		c.IndentedJSON(400, gin.H{"error": "shortUrl is required in the params"})
	}

	var url URL
	err := urlColl.FindOne(context.TODO(), bson.M{"shortUrl": shortUrl}).Decode(&url)
	if err != nil {
		c.IndentedJSON(404, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.Url)
}

func generateRandomString(n int) string {
	char := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZZ")
	b := make([]rune, n)
	for i := range b{
		b[i] = char[rand.Intn(len(char))]
	}

	return string(b)
}