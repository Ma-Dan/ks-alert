package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCreateRegisteredEnterpriseItem(t *testing.T) {
	Convey("Test CreateRegisteredEnterprise Item", t, func() {

		enterpriseID := "enterprise-x2y6vmrk7q82wz"

		err := CreateRegisteredEnterpriseItem(&RegisteredEnterprise{
			EnterpriseID: enterpriseID,
		})

		So(err, ShouldBeNil)
	})
}

func TestCreateRegisteredProductItem(t *testing.T) {
	Convey("Test CreateRegisteredProduct Item", t, func() {

		enterpriseID := "enterprise-x2y6vmrk7q82wz"
		productID := "product-4llxr47k7q82wz"

		err := CreateRegisteredProductItem(&RegisteredProduct{
			EnterpriseID: enterpriseID,
			ProductID:    productID,
		})

		So(err, ShouldBeNil)
	})
}
