package calculate

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"keyword-generator/src/adaptor/dto/response"
	"keyword-generator/src/application/shared"
	"keyword-generator/src/config"
	"keyword-generator/src/infrastructure/shared/elastic_nethttp"
	"keyword-generator/src/infrastructure/shared/sentry"
	"log"
	"strings"
	"time"
)

func (u *useCase) CalculateAll(ctx context.Context) shared.StatusMessage {
	ucSpan := sentry.SentryConfig.UsecaseSpan.StartChild("calculate Usecase")
	defer ucSpan.Finish()
	listTitle, err := u.mongo.MongoCalculate(ctx)
	if err != nil {
		sentrygo.CaptureException(err)
		log.Printf("Failed Get Result From Mongo :%v", err)
		return shared.ERROR_MONGO
	}
	fmt.Println("Get Result From Mongo with len = ", len(listTitle))
	var arrayContent []response.ContentDetail
	var listContent []response.ContentDetail
	var titleFormated []string
	ttl := time.Duration(config.Config.GetInt("REDIS_TTL")) * 24 * time.Hour
	for _, v := range listTitle {
		//set to redis every user have keyword
		var userStr string
		switch user := v.UserID.(type) {
		case float64:
			userStr = fmt.Sprintf("%.0f", user)
		case string:
			userStr = user
		}
		titleFormated = nil
		for _, vTitle := range v.Title {
			if vTitle == "" {
				break
			}
			title := strings.ToLower(vTitle)
			titlestr := strings.ReplaceAll(title, ":", "")
			keyword := strings.ReplaceAll(titlestr, " ", "_")
			titleFormatted := strings.ReplaceAll(vTitle, "\n", "")
			titleForElastic := strings.ReplaceAll(titleFormatted, "\"", "")

			titleFormated = append(titleFormated, keyword)
			listContent, err = elastic_nethttp.GetDataElastic(ctx, titleForElastic)
			if err != nil {
				log.Printf("Failed Get Data From Elastic :%v", err)
				continue
			}

			var listContentHot []response.ContentDetail
			var listContentVideo []response.ContentDetail
			for _, vContent := range listContent {
				k := fmt.Sprintf("shorts:master:%s:content_type:%s:%d", vContent.Service, vContent.ContentType, vContent.ContentId)
				_, err = u.redis.Get(ctx, k)
				if err != nil {
					continue
				}

				//set to redis every keyword have list content group by service
				switch vContent.Service {
				case "hot":
					listContentHot = append(listContentHot, vContent)
				case "video":
					listContentVideo = append(listContentVideo, vContent)
				}
				arrayContent = append(arrayContent, vContent)
			}
			key := fmt.Sprintf("shorts:user:%s:keyword_generated", userStr)
			existingDataTittle, err1 := u.redis.Get(ctx, key)
			if err1 != nil {
				if err1 == redis.Nil {
					err = u.redis.Set(ctx, key, titleFormated, ttl)
					if err != nil {
						sentrygo.CaptureException(err)
						log.Printf("Failed Set User Keyword :%v", err)
					}
				}
			} else {
				var arrExistingDataTittle []string
				_ = json.Unmarshal([]byte(existingDataTittle), &arrExistingDataTittle)
				for _, vUserKW := range arrExistingDataTittle {
					titleFormated = append(titleFormated, vUserKW)
				}
				seen := make(map[string]bool)
				var uniqueTitleFormatted []string
				for _, val := range titleFormated {
					if _, ok := seen[val]; !ok {
						seen[val] = true
						uniqueTitleFormatted = append(uniqueTitleFormatted, val)
					}
				}
				err = u.redis.Set(ctx, key, uniqueTitleFormatted, ttl)
				if err != nil {
					sentrygo.CaptureException(err)
					log.Printf("Failed Set User Keyword :%v", err)
				}
			}

			//set to redis every keyword have list content group by service hot
			keyContentHot := fmt.Sprintf("shorts:lineup:hot:keyword:%s", keyword)
			if listContentHot != nil {
				err = u.redis.Set(ctx, keyContentHot, listContentHot, ttl)
				if err != nil {
					sentrygo.CaptureException(err)
					log.Printf("Failed Set Keyword HOT :%v", err)
				}
			}

			//set to redis every keyword have list content group by service video
			keyContentVideo := fmt.Sprintf("shorts:lineup:video:keyword:%s", keyword)
			if listContentVideo != nil {
				err = u.redis.Set(ctx, keyContentVideo, listContentVideo, ttl)
				if err != nil {
					sentrygo.CaptureException(err)
					log.Printf("Failed Set Keyword VIDEO :%v", err)
				}
			}

			//set to redis every keyword have list content group by fyp
			keyContent := fmt.Sprintf("shorts:lineup:fyp:keyword:%s", keyword)
			if listContent != nil {
				err = u.redis.Set(ctx, keyContent, listContent, ttl)
				if err != nil {
					sentrygo.CaptureException(err)
					log.Printf("Failed Set Keyword FYP :%v", err)
				}
			}
		}
	}
	fmt.Println("array content = ", len(arrayContent))

	return shared.SUCCESS
}