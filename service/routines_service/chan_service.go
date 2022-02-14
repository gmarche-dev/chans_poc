package routinesservice

import (
	"fmt"
	"sync"
	"time"
)

type Service struct {
	routines map[int]*goRoutine
	mux      *sync.RWMutex
}

func NewService() *Service {
	mux := &sync.RWMutex{}
	routines := map[int]*goRoutine{}
	return &Service{
		routines: routines,
		mux:      mux,
	}
}

func (s *Service) AddRoutineEvent(id int, re ...RoutineEvent) {
	s.mux.Lock()
	defer s.mux.Unlock()
	if r, exists := s.routines[id]; !exists {
		r = newRoutine(id, re...)
		s.routines[r.id] = r
		s.startRoutine(r)
	} else {
		r.UpdateRoutine(re...)
	}
}

func (s *Service) GetRoutine(id int) *goRoutine {
	s.mux.RLock()
	defer s.mux.RUnlock()
	fmt.Println(fmt.Sprintf("get routine value %d", id))
	if r, exists := s.routines[id]; exists {
		return r
	}

	return nil
}

func (s *Service) CloseRoutine(id int) {
	if r, exists := s.routines[id]; exists {
		r.CloseRoutine()
	}
}

func (s *Service) removeRoutine(r *goRoutine) {
	fmt.Println(fmt.Sprintf("do remove event %d", r.id))
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.routines, r.id)
}

func (s *Service) startRoutine(r *goRoutine) {
	go func() {
		select {
		case q := <-r.quit:
			if q {
				s.removeRoutine(r)
				fmt.Println(fmt.Sprintf("force quit goroutine %d", r.id))
				return
			}
		case <-time.After(r.expired):
			s.removeRoutine(r)
			fmt.Println(fmt.Sprintf("close routine %d expired", r.id))
			return
		}
	}()
}
