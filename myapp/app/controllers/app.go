package controllers

import (
	"github.com/revel/revel"
    "net/http"
	"io/ioutil"
	// "fmt"
	"encoding/json"
	"log"
)


type App struct {
	*revel.Controller
}

type Item struct {
	ItemName string `json:"itemName"`
	ItemPrice int `json:"itemPrice"`
	ItemUrl string `json:"itemUrl"`
	ImageUrls []string `json:"mediumImageUrls"`
}
type ApiResult struct {
	Count int  `json:"count"`
	Page int  `json:"page"`
	PageCount int  `json:"pageCount"`
    Items []Item `json:"Items"`
}



func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Search() revel.Result {
	
	return c.Render()
}

func (c App) SearchKeyword() revel.Result {
	args := make(map[string]string)
	
	keyword := c.Params.Route.Get("keyword")
	args["keyword"] = keyword

	rakutenApiUri := "https://app.rakuten.co.jp/services/api/IchibaItem/Search/20170706";
	rakutenApiQuery := make(map[string]string)

	rakutenApiQuery["keyword"] = keyword
	rakutenApiQuery["applicationId"] = "1064956328169495300"
	rakutenApiQuery["formatVersion"] = "2"
	// yahooApiQuery["sort"] = "-score"
	// yahooApiQuery["availability"] = "1"
	// https://app.rakuten.co.jp/services/api/IchibaItem/Search/20170706?format=json&keyword=%E6%A5%BD%E5%A4%A9&genreId=559887&shopCode=rakuten24&applicationId=1064956328169495300

	query := NewSortedQuery(rakutenApiQuery)

	resp, _ := http.Get(rakutenApiUri+"?"+query.String())
	defer resp.Body.Close()
  
	byteArray, _ := ioutil.ReadAll(resp.Body)
	// fmt.Printf( string(byteArray) )
	// result := json.NewDecoder(string(byteArray));
    // JSONデコード
    var apiresult ApiResult
    if err := json.Unmarshal(byteArray, &apiresult); err != nil {
        log.Fatal(err)
    }
    // デコードしたデータを表示
    // for _, g := range apiresult.Items{
    //     fmt.Printf("%d : %s\n", g.ItemName)
    // }
	
	return c.Render(keyword, apiresult)
}


