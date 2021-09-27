package models

type BookManagement struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Author       string `json:"author"`
	Prices       int    `json:"prices"`
	Available    string `json:"available"`
	PageQuality  string `json:"pagequality"`
	LaunchedYear string `json:"launchedyear"`
	Isbn         string `json:"isbn"`
	Stock        int    `json:"stock"`
}
