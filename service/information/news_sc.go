package information

import (
	"demo-gin-api/model"
	"demo-gin-api/model/dao"
	"demo-gin-api/model/vo"
	"demo-gin-api/util"
	"strconv"
)

func GetNewsList(starttime, endtime, top string) []*vo.NewsVO {

	news := make([]*vo.NewsVO, 0)

	news = getNewsList(starttime, endtime)

	if top == model.Empty_string {
		return news
	}

	if count, err := strconv.Atoi(top); err == nil && len(news) != 0 {
		lastCount := util.GetTopLastLenByCount(count, len(news))
		news = news[model.Zero_value:lastCount]
	}

	return news
}

func GetNews(id int64) vo.NewsVO {

	newsData := dao.GetNewsById(id)

	if newsData.Id == model.Zero_value {
		return vo.NewsVO{}
	}

	n := getNews(newsData)
	return n
}

func getNewsList(starttime, endtime string) []*vo.NewsVO {

	news := make([]*vo.NewsVO, 0)

	var dbDatas []*dao.News

	if starttime != model.Empty_string && endtime != model.Empty_string {
		dbDatas = dao.GetNewsByDate(starttime, endtime)
	} else {
		dbDatas = dao.GetNews()
	}

	for _, newsData := range dbDatas {
		n := getNews(*newsData)
		news = append(news, &n)
	}

	return news
}

func getNews(data dao.News) vo.NewsVO {

	var n vo.NewsVO

	n.Id = data.Id
	n.Title = data.Title
	n.Content = data.Content
	n.Belong = data.Belong
	n.Url = data.Url
	n.Date = data.Date.Format(util.DateFormat)

	return n
}
