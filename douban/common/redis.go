package common

import (
	"github.com/garyburd/redigo/redis"
)

var redisDb *redis.Pool

const (
	MaxIdle       = 5
	MaxActive     = 10
	DefaultDb     = 0
	MovieIdList   = "movie_id"
	MovieInfoList = "movie_list"
	IpPool        = "ip_pool"
)

func init() {
	redisDb = &redis.Pool{
		MaxIdle:   MaxIdle,
		MaxActive: MaxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				return nil, err
			}
			c.Do("select", DefaultDb)
			return c, nil
		},
	}
	rs := redisDb.Get()
	_, err := rs.Do("ping")
	if err != nil {
		panic(err)
	}
	rs.Close()
}

// 添加处理完的movie_id
func AddMovieId(id int64) (int, error) {
	rs := redisDb.Get()
	defer rs.Close()
	return redis.Int(rs.Do("sadd", MovieIdList, id))
}

// movie_id是否已处理
func IsExistMovieId(id int64) (int, error) {
	rs := redisDb.Get()
	defer rs.Close()
	return redis.Int(rs.Do("sismember", MovieIdList, id))
}

// 添加movie详情
func AddMovieInfo(info []byte) (int, error) {
	rs := redisDb.Get()
	defer rs.Close()
	return redis.Int(rs.Do("lpush", MovieInfoList, info))
}

// 获取一条movie详情并删除
func GetMovieInfo() ([]byte, error) {
	rs := redisDb.Get()
	defer rs.Close()
	//return redis.Bytes(rs.Do("LPOP", MovieInfoList))
	return redis.Bytes(rs.Do("lindex", "movie_list_2", 0))
}

// 添加ip到代理池中
func AddIpPool(ip string) (int, error) {
	rs := redisDb.Get()
	defer rs.Close()
	return redis.Int(rs.Do("sadd", IpPool, ip))
}

// 从代理池中随机获取一条记录
func GetIp() (string, error) {
	rs := redisDb.Get()
	defer rs.Close()
	return redis.String(rs.Do("SRANDMEMBER", IpPool))
}

// 删除代理池中的IP
func DelIp(ip string) (int, error) {
	rs := redisDb.Get()
	defer rs.Close()
	return redis.Int(rs.Do("SREM", IpPool, ip))
}
