package register

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBatch(t *testing.T) {

	const (
		dbURL  = "sqlite:test_batch.db"
		dbFile = "test_batch.db"
	)

	defer os.Remove(dbFile)

	Convey("AlphaOmega", t, func() {

		r, err := New(dbURL)
		So(err, ShouldBeNil)

		exists, err := r.BatchExists(42)
		So(err, ShouldBeNil)
		So(exists, ShouldBeFalse)

		err = r.AddBatch(42)
		So(err, ShouldBeNil)

		exists2, err := r.BatchExists(42)
		So(err, ShouldBeNil)
		So(exists2, ShouldBeTrue)

		err = r.AddBatch(43)
		So(err, ShouldBeNil)

		batches, err := r.GetBatches()
		So(err, ShouldBeNil)
		So(len(batches), ShouldEqual, 2)
		So(batches[0].Batch, ShouldEqual, 43)
		So(batches[1].Batch, ShouldEqual, 42)

	})

}
