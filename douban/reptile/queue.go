package reptile

import (
	"sync"
	"time"
)

type queue struct {
	sync.Mutex
	movie []*MovieInfo
}

func NewMovieQueue() *queue {
	return &queue{
		movie: make([]*MovieInfo, 0),
	}
}

// 加入队列
func (q *queue) Put(movie ...*MovieInfo) {
	q.Lock()
	defer q.Unlock()
	q.movie = append(q.movie, movie...)
}

// 队列长度
func (q *queue) Len() int {
	q.Lock()
	defer q.Unlock()
	return len(q.movie)
}

// 从队列顶部获取一条记录并删除
func (q *queue) Poll() *MovieInfo {
	q.Lock()
	defer q.Unlock()
	if len(q.movie) <= 0 {
		return nil
	}
	movie := q.movie[0]
	q.movie[0] = nil
	q.movie = q.movie[1:]
	return movie
}

func init() {
	mainMovieQueue = NewMovieQueue()
}

// 添加队列
func MoviePut(movie ...*MovieInfo) {
	mainMovieQueue.Put(movie...)
}

func MovieLen() int {
	return mainMovieQueue.Len()
}

// 从队列顶部获取一条记录并删除
func MoviePoll() *MovieInfo {
	for {
		m := mainMovieQueue.Poll()
		if m != nil {
			return m
		}
		time.Sleep(time.Second)
	}
}

var mainMovieQueue *queue
