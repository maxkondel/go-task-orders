package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestApp(t *testing.T) {
	Convey("App should be started ", t, func() {

	})
}

func TestMail(m *testing.M) {

	m.Run()

}
