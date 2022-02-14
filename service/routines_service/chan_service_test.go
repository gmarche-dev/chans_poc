package routinesservice_test

import (
	routinesservice "chans_poc/service/routines_service"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type test struct {
	id      int
	expired time.Duration
}

func Test_Service(t *testing.T) {
	// list of events with expired value
	events := []test{
		test{1, 5 * time.Second},
		test{2, 5 * time.Second},
		test{3, 5 * time.Second},
		test{4, 5 * time.Second},
		test{5, 5 * time.Second},
		test{6, 5 * time.Second},
	}
	service := routinesservice.NewService()

	for _, event := range events {
		service.AddRoutineEvent(
			event.id,
			routinesservice.AddFirstEvent(&routinesservice.FirstEvent{Name: fmt.Sprintf("first event id %d", event.id)}),
			routinesservice.UpdateExpiredDate(5*time.Second))
	}

	for _, event := range events {
		if event.id%2 == 0 {
			service.AddRoutineEvent(
				event.id,
				routinesservice.AddSecondEvent(&routinesservice.SecondEvent{Name: fmt.Sprintf("second event id %d", event.id)}),
				routinesservice.AddThirdEvent(&routinesservice.ThirdEvent{Name: fmt.Sprintf("third event id %d", event.id)}))
		} else {
			service.AddRoutineEvent(
				event.id,
				routinesservice.UpdateExpiredDate(10*time.Second))
		}
	}

	time.Sleep(1 * time.Nanosecond)
	for _, event := range events {
		if event.id%2 == 0 {
			assert.Nil(t, service.GetRoutine(event.id))
		} else {
			assert.NotNil(t, service.GetRoutine(event.id))
		}
	}

	time.Sleep(1 * time.Nanosecond)
	time.Sleep(10 * time.Second)
	for _, event := range events {
		if event.id%2 != 0 {
			assert.Nil(t, service.GetRoutine(event.id))
		}
	}

	service.AddRoutineEvent(
		7,
		routinesservice.AddFirstEvent(&routinesservice.FirstEvent{Name: fmt.Sprintf("first event id %d", 7)}),
		routinesservice.UpdateExpiredDate(20*time.Second))

	time.Sleep(5 * time.Second)

	service.CloseRoutine(7)
	time.Sleep(1 * time.Nanosecond)
	assert.Nil(t, service.GetRoutine(7))
}
