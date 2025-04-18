package internal

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"

	"github.com/Paschalolo/reddit-recipie-aggregator/rss-parser/pkg"
	"github.com/gin-gonic/gin"
)

func GetFeedEntries(url string) (*[]pkg.Entry, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(resp.StatusCode)
	defer resp.Body.Close()
	byteValue, _ := io.ReadAll(resp.Body)
	// log.Println("Raw XML:\n", string(byteValue))
	var feed pkg.Feed
	if err := xml.Unmarshal(byteValue, &feed); err != nil {
		log.Println("could not unmarshal ")
		log.Println(err)
		return nil, err
	}
	log.Println(feed.Entries)
	return &feed.Entries, nil
}

func ParseHandler(c *gin.Context) {
	var request pkg.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	entries, err := GetFeedEntries(request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while parsing rss feed ",
		})
		return
	}

	c.JSON(http.StatusOK, entries)
}
