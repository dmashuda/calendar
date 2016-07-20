package calendar

import (
	"github.com/dmashuda/calendar/storage/memory"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBasic(t *testing.T) {

	Convey("Given something", t, func() {
		req := CreateEventRequest{}
		com := NewCreateEvent(memory.EventRepo{}, &req)
		com.Execute()
		So(1, ShouldEqual, 1)
	})
}
