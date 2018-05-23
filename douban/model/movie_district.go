package model

type MovieDistrict struct {
	Id         int64
	MovieId    int64
	DistrictId int64
	CreateTime int64 `xorm:"created"`
}

func NewMovieDistrict(movieId, districtId int64) *MovieDistrict {
	return &MovieDistrict{
		MovieId:    movieId,
		DistrictId: districtId,
	}
}

func (md *MovieDistrict) Insert() (int64, error) {
	return engine.InsertOne(md)
}
