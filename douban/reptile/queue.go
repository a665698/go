package reptile

import (
	"douban/model"
	"sync"
)

type queue struct {
	sync.Mutex
	movie []*model.Movie
}

func NewMovieQueue() *queue {
	return &queue{
		movie: make([]*model.Movie, 0),
	}
}

func (q *queue) Put(movie ...*model.Movie) {
	q.Lock()
	defer q.Unlock()
	q.movie = append(q.movie, movie...)
}

func (q *queue) Len() int {
	q.Lock()
	defer q.Unlock()
	return len(q.movie)
}

func (q *queue) Poll() *model.Movie {
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

func MoviePut(movie ...*model.Movie) {
	mainMovieQueue.Put(movie...)
}

func MovieLen() int {
	return mainMovieQueue.Len()
}

func MoviePoll() *model.Movie {
	return mainMovieQueue.Poll()
}

var mainMovieQueue *queue
