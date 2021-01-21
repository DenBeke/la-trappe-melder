package scraper

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	latrappeURL = "https://www.kloosterwinkelonline.nl/la-trappe-trappist-oak-aged"
	pageTitle   = "Proef de stilte. La Trappe Quadrupel Oak Aged batch #39"
	batch       = "39"
)

func TestScraper(t *testing.T) {

	Convey("GetBatchVersion", t, func() {

		batch, err := GetBatchVersion(latrappeURL)
		So(err, ShouldBeNil)

		So(batch, ShouldEqual, "39")

	})

	Convey("getBatchVersionFromPageTitle", t, func() {

		batch, err := getBatchVersionFromPageTitle(pageTitle)
		So(err, ShouldBeNil)

		So(batch, ShouldEqual, "39")

	})
}
