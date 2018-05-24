package model

type Runtime struct {
	Id         int64
	Time       int
	DistrictId int64
	MovieId    int64
	CreateTime int64 `xorm:"created"`
}

func NewRuntime(time int, districtId, movieId int64) *Runtime {
	return &Runtime{
		Time:       time,
		DistrictId: districtId,
		MovieId:    movieId,
	}
}

func (r *Runtime) Insert() (int64, error) {
	return engine.InsertOne(r)
}
