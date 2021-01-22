package register

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubscription(t *testing.T) {

	const (
		dbURL  = "sqlite:test_subscription.db"
		dbFile = "test_subscription.db"
	)

	defer os.Remove(dbFile)

	Convey("AlphaOmega", t, func() {

		r, err := New(dbURL)
		So(err, ShouldBeNil)

		const (
			name  = "My Name"
			email = "my.email@example.com"
		)

		exists, err := r.EmailExists(email)
		So(err, ShouldBeNil)
		So(exists, ShouldBeFalse)

		s, err := r.Subscribe(name, email)
		So(err, ShouldBeNil)
		So(s, ShouldNotBeNil)

		exists2, err := r.EmailExists(email)
		So(err, ShouldBeNil)
		So(exists2, ShouldBeTrue)

		So(s.Email, ShouldEqual, email)
		So(s.Name, ShouldEqual, name)
		So(s.UUID, ShouldNotBeBlank)
		So(s.Confirmed, ShouldBeFalse)

		_, err = r.ConfirmSubscription("random string")
		So(err, ShouldNotBeNil)

		s2, err := r.ConfirmSubscription(s.UUID)
		So(err, ShouldBeNil)
		So(s2, ShouldNotBeNil)
		So(s2.Confirmed, ShouldBeTrue)

		allS, err := r.GetAllSubscriptions()
		So(err, ShouldBeNil)
		So(allS, ShouldHaveLength, 1)

		err = r.UnSubscribe(email)
		So(err, ShouldBeNil)

		allS2, err := r.GetAllSubscriptions()
		So(err, ShouldBeNil)
		So(allS2, ShouldHaveLength, 0)

	})
}
