package util

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadDotEnv(t *testing.T) {
	Convey("LoadDotEnv", t, func() {
		os.Setenv("TEST_ENV_VAR", "true")

		LoadDotEnv("../.env.testing")

		Convey("os environment variables should take priority", func() {
			result, err := GetEnvStr("TEST_ENV_VAR", "false")
			So(result, ShouldEqual, "true")
			So(err, ShouldBeNil)
		})

		Convey("it should use the loaded variables if no environment variable is set", func() {
			result, err := GetEnvStr("ENV_VAR_TESTING", "true")
			So(result, ShouldEqual, "false")
			So(err, ShouldBeNil)
		})

		Convey("it should return the fallback if the variable is not found in the environment or the loaded variables", func() {
			result, err := GetEnvStr("ENV_VAR_TESTING_2", "false")
			So(result, ShouldEqual, "false")
			So(err, ShouldEqual, errEnvVarEmpty)
		})
	})
}

func TestGetEnvStr(t *testing.T) {
	Convey("GetEnvStr", t, func() {

		os.Setenv("ENV_TEST", "real value")

		Convey("it should return the value of the env variable", func() {
			result, err := GetEnvStr("ENV_TEST", "fallback")
			So(result, ShouldEqual, "real value")
			So(err, ShouldBeNil)
		})

		Convey("it return an error and the fallback value, if the env var does not exist", func() {
			result, err := GetEnvStr("MISSING_123ENV123_VAR", "fallback")
			So(result, ShouldEqual, "fallback")
			So(err, ShouldEqual, errEnvVarEmpty)
		})
	})
}

func TestGetEnvBool(t *testing.T) {
	Convey("GetEnvBool", t, func() {

		os.Setenv("ENV_TEST", "true")
		os.Setenv("ENV_FAIL_BOOL", "ME NOT YAYA")

		Convey("it should return the value of the env variable", func() {
			result, err := GetEnvBool("ENV_TEST", false)
			So(result, ShouldEqual, true)
			So(err, ShouldBeNil)
		})

		Convey("it return an error and the fallback value, if the env var does not exist", func() {
			result, err := GetEnvBool("MISSING_123ENV123_VAR", true)
			So(result, ShouldEqual, true)
			So(err, ShouldEqual, errEnvVarEmpty)
		})

		Convey("it should return an error and the fallback, if the value could not be parsed to an boolean", func() {
			result, err := GetEnvBool("ENV_FAIL_BOOL", false)
			So(result, ShouldEqual, false)
			So(err, ShouldNotBeNil)
		})
	})
}
