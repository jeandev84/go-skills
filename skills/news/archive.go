package news


/*
import (
	"encoding/json"
)
*/

// хранить в памяти

// писать type с большими буквами (Topic) это type будет "public"
// а с маленькими буквами (topic) это "private"

type Topic struct {

	Type    string `json:"title"`
	Author  string `json:"author"`
	URL     string `json:"url"`
}  



// хранить Topic в памяте
type Archive map[string][]Topic



func New() map[string][]Topic {

   return make(map[string][]Topic, 0)
}



// добавить категорию в Архив
func (a *Archive) collect(category string) {

    sources := getSources(category)
    topics := getTopics(sources)


	a[category] = topics
}


// Получить категорию из Архива
func (a *Archive) result(category string) string []Topic{
	return a[category]
}