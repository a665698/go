package model

type Summary struct {
	Id         int64
	MovieId    int64
	Text       string
	CreateTime int64 `xorm:"created"`
}

func NewSummary(movieId int64, text string) *Summary {
	return &Summary{
		MovieId: movieId,
		Text:    text,
	}
}

func (s *Summary) Insert() (int64, error) {
	return engine.InsertOne(s)
}
