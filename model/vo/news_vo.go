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
	Title   string `form:"title" binding:"required"`
	Content string `form:"content"`
	Belong  string `form:"belong"`
	Url     string `form:"url"`
	Date    string `form:"date" binding:"required"`
}

type NewsPutVO struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content"`
	Belong  string `json:"belong"`
	Url     string `json:"url"`
	Date    string `json:"date" binding:"required"`
}
