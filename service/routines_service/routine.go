package routinesservice

import (
	"fmt"
	"time"
)

type RoutineEvent func(*goRoutine)

type FirstEvent struct {
	Name string
}

type SecondEvent struct {
	Name string
}

type ThirdEvent struct {
	Name string
}

func AddFirstEvent(f *FirstEvent) RoutineEvent {
	return func(r *goRoutine) {
		r.First = f
	}
}

func AddSecondEvent(s *SecondEvent) RoutineEvent {
	return func(r *goRoutine) {
		r.Second = s
	}
}

func AddThirdEvent(t *ThirdEvent) RoutineEvent {
	return func(r *goRoutine) {
		r.Third = t
	}
}

func UpdateExpiredDate(t time.Duration) RoutineEvent {
	return func(r *goRoutine) {
		r.expired = t
	}
}

type goRoutine struct {
	id          int
	First       *FirstEvent
	Second      *SecondEvent
	Third       *ThirdEvent
	quit        chan bool
	expired     time.Duration
}

func newRoutine(id int, re ...RoutineEvent) *goRoutine {
	quit := make(chan bool, 1)
	r := &goRoutine{id: id, quit: quit}
	r.UpdateRoutine(re...)
	return r
}

func (g *goRoutine) CloseRoutine() {
	g.quit <- true
}

func (g *goRoutine) UpdateRoutine(re ...RoutineEvent) {
	for _, r := range re {
		r(g)
	}
	if g.isComplete() {
		fmt.Println(fmt.Sprintf("event %d is complete", g.id))
		g.CloseRoutine()
	}
}

func (g *goRoutine) isComplete() bool {
	return g.First != nil && g.Second != nil && g.Third != nil
}