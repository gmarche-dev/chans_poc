package main

import (
	"chans_poc/service/routines_service"
	"fmt"
	"time"
)

func main() {
	routines := routinesservice.NewService()

	// add routine 1
	routines.AddRoutineEvent(
		1,
		routinesservice.AddFirstEvent(&routinesservice.FirstEvent{Name: fmt.Sprintf("first event id %d", 1)}),
		routinesservice.UpdateExpiredDate(2*time.Minute),
	)

	// add routine 2
	routines.AddRoutineEvent(
		2,
		routinesservice.AddThirdEvent(&routinesservice.ThirdEvent{Name: fmt.Sprintf("third event id %d", 2)}),
		routinesservice.UpdateExpiredDate(44*time.Second),
	)

	// update routine 1 value
	routines.AddRoutineEvent(
		1,
		routinesservice.AddThirdEvent(&routinesservice.ThirdEvent{Name: fmt.Sprintf("third event id %d", 1)}),
	)

	// update routine 1 expiration date
	routines.AddRoutineEvent(
		1,
		routinesservice.UpdateExpiredDate(0*time.Second),
	)

	// add routine 3
	routines.AddRoutineEvent(
		3,
		routinesservice.AddFirstEvent(&routinesservice.FirstEvent{Name: fmt.Sprintf("first event id %d", 3)}),
		routinesservice.UpdateExpiredDate(2*time.Minute),
	)

	// update routine 3 second value
	routines.AddRoutineEvent(
		3,
		routinesservice.AddSecondEvent(&routinesservice.SecondEvent{Name: fmt.Sprintf("second event id %d", 3)}),
		routinesservice.UpdateExpiredDate(2*time.Minute),
	)

	// update routine 3 third value
	routines.AddRoutineEvent(
		3,
		routinesservice.AddThirdEvent(&routinesservice.ThirdEvent{Name: fmt.Sprintf("third event id %d", 3)}),
	)
	time.Sleep(1 * time.Nanosecond)
	r3 := routines.GetRoutine(3)
	fmt.Println(fmt.Sprintf("check routine 3 is deleted %v \n -----------", r3))
	time.Sleep(6 * time.Second)

	r1 := routines.GetRoutine(1)
	fmt.Println(fmt.Sprintf("check routine 1 is deleted %v \n -----------", r1))

	time.Sleep(45 * time.Second)
	r2 := routines.GetRoutine(2)
	fmt.Println(fmt.Sprintf("check routine 2 is deleted %v \n -----------", r2))
}
