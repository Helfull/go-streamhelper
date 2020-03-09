package bot

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConfig(t *testing.T) {
	Convey("Config", t, func() {
		Convey("It should return a instance of bot.Config", func() {
			expectedCfg := Config{}

			resultCfg := GetConfig("../settings.json")

			So(expectedCfg, ShouldHaveSameTypeAs, resultCfg)
		})

		Convey("It should panic if any error happend", func() {
			So(func() { GetConfig("SOME NOT EXISTING FILE") }, ShouldPanic)
		})
	})
}
