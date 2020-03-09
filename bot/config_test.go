package bot

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {
	Convey("Config", t, func() {
		Convey("It should return a instance of bot.Config", func() {
			expectedCfg := Config{}

			resultCfg := GetConfig("../settings.json")

			So(expectedCfg, ShouldHaveSameTypeAs, resultCfg)
		})

		Convey("It should panic if any error happened", func() {
			So(func() { GetConfig("SOME NOT EXISTING FILE") }, ShouldPanic)
		})
	})
}
