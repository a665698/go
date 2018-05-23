package model

type Runtime struct {
	Id         int64
	Time       int
	DistrictId int64
	CreateTime int64 `xrom:"created"`
}

func NewRuntime(time int, districtId int64) *Runtime {
	return &Runtime{
		Time:       time,
		DistrictId: districtId,
	}
}

func (r *Runtime) Insert() (int64, error) {
	return engine.InsertOne(r)
}
