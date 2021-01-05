package tests

type book struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Publisher   string   `json:"publisher"`
	Tags        []string `json:"tags"`
}

type library struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Books       []book `json:"books"`
}

type libraryList struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type libraries struct {
	Libraries []libraryList `json:"libraries"`
}
