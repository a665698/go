package reptile

import (
	"fmt"
	"testing"
)

func TestMovieInfo_SummaryHandle(t *testing.T) {
	movieInfo := NewMovieInfo()
	movieInfo.MovieId = 26992279
	movieInfo.Summary = `
                                    　　Synopsis
                                    　　There     is a man who is timid and shy but has a legendary package! Chi-soo is a repeater who gets bullied by his friends. However, he had a secret about his body that rumors wouldn't die down about...
                                    　　The secret was that he was the possessor of an outstanding package. Rumors about Chi-soo spread and all the women begin to desire him.           `
	movieInfo.SummaryHandle(12131)
	fmt.Println(movieInfo.Summary)
}

func TestDelRepeatMovie(t *testing.T) {
	DelRepeatMovie()
}
