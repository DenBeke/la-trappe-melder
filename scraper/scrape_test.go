package scraper

import (
	"io/ioutil"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	latrappeURL  = "https://nl.latrappetrappist.com/nl/nl/producten/batchregister.html"
	batch        = "39"
	testPagefile = "batchregister.html"
)

func TestScraper(t *testing.T) {

	Convey("GetBatchVersion", t, func() {

		batch, err := GetBatchVersion(latrappeURL)
		So(err, ShouldBeNil)

		So(batch, ShouldEqual, "39")

	})

	Convey("getBatchFromPage", t, func() {

		pageBytes, err := ioutil.ReadFile(testPagefile)
		So(err, ShouldEqual, nil)
		page := string(pageBytes)

		b, err := getBatchFromPage(strings.NewReader(page))
		So(err, ShouldEqual, nil)
		So(b, ShouldEqual, batch)

	})

	Convey("getBatchFromH4Tag", t, func() {

		b, err := getBatchFromH4Tag("Batch 39")
		So(err, ShouldEqual, nil)
		So(b, ShouldEqual, batch)

		b, err = getBatchFromH4Tag(" Batch 39 ")
		So(err, ShouldEqual, nil)
		So(b, ShouldEqual, batch)

		_, err = getBatchFromH4Tag("batch 39")
		So(err, ShouldNotEqual, nil)

		_, err = getBatchFromH4Tag("39")
		So(err, ShouldNotEqual, nil)

		_, err = getBatchFromH4Tag("#39")
		So(err, ShouldNotEqual, nil)

		_, err = getBatchFromH4Tag("Batch NotANumber")
		So(err, ShouldNotEqual, nil)

		_, err = getBatchFromH4Tag("Batch Not A Number")
		So(err, ShouldNotEqual, nil)

	})

}
