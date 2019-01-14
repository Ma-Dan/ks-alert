package models

import (
	"testing"
	"time"
)

func TestCreateEnterprise(t *testing.T) {
	Convey("test database", t, func() {
		Convey("test enterprise create", func() {
			var enterprise = Enterprise{
				EnterpriseName: "qingcloud",
				HomePage:       "https://www.qingcloud.com/",
				Address:        "北京优帆科技有限公司",
				Email:          "yunify@yunify.com",
				Contacts:       "Richard",
				Description:    "云计算公司",
				Phone:          "400-8576-886",
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			}

			_, err := CreateEnterprise(&enterprise)
			So(err, ShouldBeNil)
		})
	})
}
