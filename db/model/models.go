package model

type Songs struct {
	Id           int64
	Title        string
	YoutubeTitle string `xorm:"unique"`
}

type Playlist struct {
	CreatedAt int64 `xorm:"created"`
	Id        int64
	Video1    int
	Video2    int
	Video3    int
	Video4    int
	Video5    int
	Video6    int
}
