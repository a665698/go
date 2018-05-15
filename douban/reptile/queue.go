package reptile

import "sync"

type queue struct {
	sync.Mutex
	movie []*Movie
}

func New() *queue {
	return &queue{
		movie: make([]*Movie, 0),
	}
}

func (q *queue) Put(movie *Movie) {
	q.Lock()
	defer q.Unlock()
	q.movie = append(q.movie, movie)
}

func (q *queue) Len() int {
	q.Lock()
	defer q.Unlock()
	return len(q.movie)
}

func (q *queue) Poll() *Movie {
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
