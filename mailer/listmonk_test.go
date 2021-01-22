package mailer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestListmonk(t *testing.T) {

	Convey("SendMail", t, func() {

		err := SendMail("http://localhost:9000", "a6458b69-6214-494c-be34-7715475b24b5", "test", "test")

		So(err, ShouldBeNil)

	})

}
