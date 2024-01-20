package elastic_nethttp

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"keyword-generator/src/adaptor/dto/response"
	"keyword-generator/src/config"
	"log"
	"net/http"
	"strings"
)

func GetDataElastic(ctx context.Context, keyword string) ([]response.ContentDetail, error) {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	var (
		r      map[string]interface{}
		result []response.ContentDetail
	)
	//client := &http.Client{}
	jsonreq := fmt.Sprintf("{\"query\":{  \"match\":    {\"tags\":\"%s\"}  },  \"collapse\":{    \"field\":\"id\"  }}", keyword)
	payload := strings.NewReader(jsonreq)

	req, _ := http.NewRequest("GET", config.Config.GetString("ELASTIC_SEARCH_HOST")+"/short_idx/_search", payload)
	auth := fmt.Sprintf("%s:%s", config.Config.GetString("ELASTIC_SEARCH_USER"), config.Config.GetString("ELASTIC_SEARCH_PASS"))

	sEnc := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+sEnc)
	//client = apmhttp.WrapClient(client)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	if err2 := json.NewDecoder(res.Body).Decode(&r); err2 != nil {
		log.Fatalf("Error parsing the response body: %s", err2)
	}
	if _, ok := r["error"]; ok {
		log.Println("Error Elastic : ", r["reason"])
		return nil, nil
	}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var resp response.ContentDetail
		source := hit.(map[string]interface{})["_source"].(interface{})
		resp.Service = source.(map[string]interface{})["service"].(string)
		resp.ContentType = source.(map[string]interface{})["content_type"].(string)
		resp.ContentId = int(source.(map[string]interface{})["id"].(float64))
		result = append(result, resp)
	}
	return result, nil
}
