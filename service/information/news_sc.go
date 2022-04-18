package information

import (
	"demo-gin-api/model"
	"demo-gin-api/model/dao"
	"demo-gin-api/model/vo"
	"demo-gin-api/util"
	"fmt"
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

func AddNews(data vo.NewsPostVO) (vo.NewsVO, error) {

	util.AddNewsMutex.Lock()

	news := newInsertNews(data)
	newsCreateId, err := dao.InsertNews(news)

	if err != nil {
		util.AddNewsMutex.Unlock()
		return vo.NewsVO{}, err
	}

	util.AddNewsMutex.Unlock()
	return vo.NewsVO{Id: newsCreateId}, nil
}

func EditNews(id int64, data vo.NewsPutVO) (string, error) {

	if editNews := dao.GetNewsById(id); editNews.Id == model.Zero_value {
		return model.Empty_string, fmt.Errorf("%s\n", model.Id_not_found_string)
	}

	news := newUpdateNews(data)

	if err := dao.UpdateNews(id, news); err != nil {
		return model.Empty_string, err
	}

	return model.Ok_string, nil
}

func DeleteNews(id int64) (string, error) {

	if deleteNews := dao.GetNewsById(id); deleteNews.Id == model.Zero_value {
		return model.Empty_string, fmt.Errorf("%s\n", model.Id_not_found_string)
	}

	if err := dao.DeleteNews(id); err != nil {
		return model.Empty_string, err
	}

	return model.Ok_string, nil
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

func newInsertNews(data vo.NewsPostVO) dao.News {

	var news dao.News

	news.Title = data.Title

	if data.Content != model.Empty_string {
		news.Content = data.Content
	}

	if data.Belong != model.Empty_string {
		news.Belong = data.Belong
	}

	if data.Url != model.Empty_string {
		news.Url = data.Url
	}

	if date, err := util.GetLocationDate(data.Date); err == nil {
		news.Date = date
	}

	return news
}

func newUpdateNews(data vo.NewsPutVO) dao.News {

	var news dao.News

	news.Title = data.Title

	if data.Content != model.Empty_string {
		news.Content = data.Content
	}

	if data.Belong != model.Empty_string {
		news.Belong = data.Belong
	}

	if data.Url != model.Empty_string {
		news.Url = data.Url
	}

	if date, err := util.GetLocationDate(data.Date); err == nil {
		news.Date = date
	}

	return news
}
