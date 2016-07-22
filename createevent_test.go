package calendar

import (
	"github.com/dmashuda/calendar/storage/memory"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateEvent(t *testing.T) {
	repo := memory.NewEventRepo()

	Convey("Given Blank Request Should Reject", t, func() {
		req := CreateEventRequest{}
		resp := CreateEventResponse{}
		com := NewCreateEvent(repo, req, &resp)
		err := com.Execute()
		So(err, ShouldNotEqual, nil)

	})

	Convey("Given End Before Start time Should Reject", t, func() {
		location, loc_error := time.LoadLocation("America/New_York")
		So(loc_error, ShouldEqual, nil)
		start := time.Date(2016, 2, 1, 1, 1, 1, 0, location)
		end := start.Add(-4 * time.Hour)
		req := CreateEventRequest{
			Start: start,
			End:   end,
			Title: "some title",
			Color: "RED",
		}
		resp := CreateEventResponse{}
		com := NewCreateEvent(repo, req, &resp)
		err := com.Execute()
		So(err, ShouldNotEqual, nil)
	})

	Convey("Given Start Before End time Should Accept", t, func() {
		location, loc_error := time.LoadLocation("America/New_York")
		So(loc_error, ShouldEqual, nil)
		start := time.Date(2016, 2, 1, 1, 1, 1, 0, location)
		end := start.Add(4 * time.Hour)
		req := CreateEventRequest{
			Start: start,
			End:   end,
			Title: "some title",
			Color: "RED",
		}
		resp := CreateEventResponse{}
		com := NewCreateEvent(repo, req, &resp)
		err := com.Execute()
		So(err, ShouldEqual, nil)

	})

}
