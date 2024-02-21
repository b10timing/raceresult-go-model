package support

import "time"

type Chat struct {
	ID        int
	UserName  string
	UserID    int
	Timestamp time.Time
	Text      string
	Lang      string
}

type Forum struct {
	ID          int
	Parent      int
	UserID      int
	UserName    string
	Timestamp   time.Time
	Subject     string
	Text        string
	Replies     int
	LastMessage time.Time
	IPAddr      string
	ForumNo     int
	Lang        string
}

type KBItem struct {
	ID             int
	Title          string
	Text           string
	Tags           string
	Files          string
	CreatedBy      int
	Created        time.Time
	Owner          int
	LastUpdated    time.Time
	LastUpdatedBy  int
	Language       string
	ParentID       int
	SiblingID      int
	ChildOrder     float64
	Deleted        bool
	Private        int
	HistoryOfID    int
	ChangeCounter  int
	ViewCounter    int
	AutoSort       bool
	PageRank       float64
	PageRankFactor float64
	Invalid        bool
}

const (
	ArticlePrivate      int = 4
	ArticleHQUS         int = 3
	ArticleHQandDistris int = 2
	ArticlePublicLink   int = 1
	ArticlePublic       int = 0
)

type Translation struct {
	ID            int
	Language      string
	Invalid       bool
	LastUpdatedBy int
}

type Child struct {
	ID        int
	Title     string
	Language  string
	SiblingID int
	CreatedBy int
	Owner     int
	Invalid   bool
	Private   int
}

type Article struct {
	ID            int
	Title         string
	Text          string
	Tags          string
	Files         []string
	CreatedBy     int
	Created       time.Time
	Owner         int
	LastUpdated   time.Time
	LastUpdatedBy int
	Language      string
	ParentID      int
	SiblingID     int
	Private       int
	ChangeCounter int
	ViewCounter   int
	Path          [][]interface{}
	Translations  []Translation
	Children      []Child
}
