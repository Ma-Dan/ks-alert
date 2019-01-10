package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateEnterprise(t *testing.T) {
	Convey("test database", t, func() {
		Convey("test enterprise update", func() {
			var enterprise = Enterprise{
				EnterpriseName: "qingcloud",
				HomePage:       "https://www.qingcloud.com/",
				Address:        "北京优帆科技有限公司",
				Email:          "yunify@yunify.com",
				Contacts:       "xxxxxxxx",
				Description:    "云计算公司",
				Phone:          "400-8576-886",
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}

			err := UpdateEnterprise(&enterprise)
			So(err, ShouldBeNil)
		})
	})
}
