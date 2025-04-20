package pkg

import "encoding/xml"

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
}

type Request struct {
	URL string `json:"url"`
}

type Entry struct {
	Link struct {
		Href string `xml:"href,attr" json:"href" bson:"href"`
	} `xml:"link" json:"link" bson:"link"`

	Title string `xml:"title" json:"title" bson:"title"`
}
