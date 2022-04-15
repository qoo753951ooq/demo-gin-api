package vo

type NewsVO struct {
	Id      int64  `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Belong  string `json:"belong"`
	Url     string `json:"url"`
}
