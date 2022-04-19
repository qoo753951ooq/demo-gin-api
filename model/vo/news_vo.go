package vo

type NewsVO struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Belong  string `json:"belong"`
	Url     string `json:"url"`
	Date    string `json:"date"`
}

type NewsPostVO struct {
	Title   string `form:"title" default:"主題" binding:"required"`
	Content string `form:"content" default:"內容"`
	Belong  string `form:"belong" default:"機關"`
	Url     string `form:"url" default:"影片連結"`
	Date    string `form:"date" default:"發佈日期(yyyy-MM-dd)" binding:"required"`
}

type NewsPutVO struct {
	Title   string `json:"title" example:"主題" binding:"required"`
	Content string `json:"content" example:"內容"`
	Belong  string `json:"belong" example:"機關"`
	Url     string `json:"url" example:"影片連結"`
	Date    string `json:"date" example:"發佈日期(yyyy-MM-dd)" binding:"required"`
}
