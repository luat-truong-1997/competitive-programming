/**
* Problem: https://leetcode.com/problems/my-calendar-i
* Easy implement checking collision
 */

package main

type Event struct {
	start int
	end   int
}

func (e *Event) checkCollision(e2 *Event) bool {
	if e.start <= e2.start && e.end > e2.start {
		return true
	}
	if e2.start <= e.start && e2.end > e.start {
		return true
	}
	return false
}

type MyCalendar struct {
	bookedEvent []*Event
}

func Constructor() MyCalendar {
	return MyCalendar{
		bookedEvent: make([]*Event, 0),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	event := &Event{start: start, end: end}
	for _, e := range this.bookedEvent {
		if event.checkCollision(e) {
			return false
		}
	}
	this.bookedEvent = append(this.bookedEvent, event)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
