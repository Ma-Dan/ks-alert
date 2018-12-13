package models

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNonething(t *testing.T) {
	Convey("test database", t, func() {
		Convey("test database insert0", func() {
			fmt.Print("test")
		})
	})
}
