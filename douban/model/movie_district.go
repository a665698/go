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

func (md *MovieDistrict) Del() (int64, error) {
	return engine.Delete(md)
}

func DelMovieDistrictByMovieId(id int64) (int64, error) {
	md := new(MovieDistrict)
	md.MovieId = id
	return md.Del()
}
