package controller

import (
	"fmt"
	"net/http"
	"time"
	"urlshortner/constant"
	"urlshortner/database"
	"urlshortner/helper"
	"urlshortner/types"

	"github.com/gin-gonic/gin"
)

func ShortTheUrl(c *gin.Context) {
	// `var shortUrlBody types.ShortUrlBody` is declaring a variable named `shortUrlBody` of type
	// `types.ShortUrlBody`. This variable will be used to store the data received in the JSON format from
	// the request body when the `ShortTheUrl` function is called.
	var shortUrlBody types.ShortUrlBody
	err := c.BindJSON(&shortUrlBody) //is used to parse the JSON data containing the long URL from the request body and bind it to the shortUrlBody variable, which is of type types.ShortUrlBody.
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": constant.BindError})
		return
	}

	code := helper.GenRandomString(6) //This line generates a random string of length 6 characters using the GenRandomString function from the helper package.

	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

	if record.UrlCode != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "this code is already in use"})
		return
	}
// When a long URL is shortened, it needs a compact, unique identifier that users can use to access the original URL. This identifier is typically generated as a random string of characters.
	var url types.UrlDb

	url.CreatedAt = time.Now().Unix()
	url.ExpiredAt = time.Now().Unix()
	url.UrlCode = code
	url.LongUrl = shortUrlBody.LongUrl
	url.ShortUrl = constant.BaseUrl + code

	resp, err := database.Mgr.Insert(url, constant.UrlCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
		return
	}
	// gin.H{} is essentially shorthand provided by the Gin framework to create a map[string]interface{} 
	c.JSON(http.StatusOK, gin.H{"error": false, "data": resp, "short_url": url.ShortUrl})  
}

func RedirectURL(c *gin.Context) {
	code := c.Param("code")

	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)

	if record.UrlCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "there is no url found"})
		return
	}
	fmt.Println(record.LongUrl)

	c.Redirect(http.StatusPermanentRedirect, record.LongUrl)
}


// Overall:
// ShortTheUrl: This function handles the creation of a short URL. It first parses the JSON body from the request to obtain the long URL. 
// Then, it generates a random code using the helper function GenRandomString. It checks if the generated code is already in use in the database. 
// If not, it creates a new record in the database with the generated code, the long URL, and other metadata. Finally, it returns the short URL in the response.

// RedirectURL: This function is responsible for redirecting users from a short URL to its corresponding long URL. 
// It extracts the code from the URL path parameter. Then, it retrieves the record associated with that code from the database. 
// If a record is found, it performs a redirect to the corresponding long URL using HTTP status code 301 (Permanent Redirect).