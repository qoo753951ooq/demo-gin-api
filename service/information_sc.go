package service

import (
	"demo-gin-api/model/vo"
	"demo-gin-api/service/information"
)

func GetNewsList(starttime, endtime, top string) []*vo.NewsVO {
	news := information.GetNewsList(starttime, endtime, top)
	return news
}

func GetNews(id int64) vo.NewsVO {
	n := information.GetNews(id)
	return n
}

func PostNews(body vo.NewsPostVO) (vo.NewsVO, error) {
	news, err := information.AddNews(body)
	return news, err
}

func PutNews(id int64, body vo.NewsPutVO) (string, error) {
	result, err := information.EditNews(id, body)
	return result, err
}
