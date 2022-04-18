package dao

import (
	db "demo-gin-api/database"
	"fmt"
	"time"
)

type News struct {
	Id         int64 `gorm:"<-:create"`
	Title      string
	Content    string
	Belong     string
	Url        string
	Date       time.Time
	Created_at time.Time `gorm:"autoCreateTime;<-:create"`
	Updated_at time.Time `gorm:"autoUpdateTime;<-:update"`
}

func GetNews() []*News {
	var news []*News

	if mariaDB, err := db.GetMariaDB(); err == nil {
		mariaDB.Debug().Find(&news)
	}

	return news
}

func GetNewsByDate(starttime, endtime string) []*News {
	var news []*News

	if mariaDB, err := db.GetMariaDB(); err == nil {
		mariaDB.Debug().
			Where(`date >= ? AND date < ?`, starttime, endtime).
			Order("date asc").
			Find(&news)
	}

	return news
}

func GetNewsById(id int64) News {
	var news News

	if mariaDB, err := db.GetMariaDB(); err == nil {
		mariaDB.Debug().First(&news, id)
	}

	return news
}

func InsertNews(data News) (int64, error) {

	mariaDB, err := db.GetMariaDB()

	if err != nil {
		return data.Id, err
	}

	result := mariaDB.Debug().Create(&data)
	fmt.Printf("Insert %s RowsAffected: %d\n", "news", result.RowsAffected)
	return data.Id, result.Error
}
