package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateEnterprise(t *testing.T) {
	Convey("test database", t, func() {
		Convey("test database insert0", func() {
			var enterprise = Enterprise{
				EnterpriseName: "北京优帆科技有限公司武汉分公司",
				HomePage:       "https://www.qingcloud.com/",
				Address:        "北京优帆科技有限公司",
				Email:          "yunify@yunify.com",
				Contacts:       "Richard",
				Description:    "云计算公司",
				Phone:          "400-8576-886",
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}

			err := CreateEnterprise(&enterprise)

			So(err, ShouldBeNil)
		})
	})
}
