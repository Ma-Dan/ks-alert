package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateProduct(t *testing.T) {
	Convey("test database", t, func() {
		Convey("test database insert0", func() {

			enterprise, err := GetEnterprise(&Enterprise{EnterpriseName: "北京优帆科技有限公司武汉分公司"})

			So(err, ShouldBeNil)

			var product = Product{
				EnterpriseID:      enterprise.EnterpriseID,
				MonitorCenterHost: "localhost",
				MonitorCenterPort: 8080,
				ProductName:       "kubesphere",
				HomePage:          "https://www.kubesphere.io/",
				Email:             "ray@yunify.com",
				Contacts:          "Ray",
				Description:       "应用平台研发部",
				Phone:             "400-8576-886",
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			}
			e := CreateProduct(&product).Error
			So(e, ShouldBeNil)
		})
	})
}
