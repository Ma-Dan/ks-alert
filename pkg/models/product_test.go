package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
	"fmt"
)

func TestCreateProduct(t *testing.T) {
	Convey("test product", t, func() {
		Convey("test create product", func() {
			var product = Product{
				MonitorCenterHost: "localhost",
				MonitorCenterPort: 8080,
				ProductName:       "kubesphere",
				HomePage:          "https://www.kubesphere.io/",
				Email:             "xxxx@yunify.com",
				Address:           "xxxxxxxxxxx",
				Contacts:          "Ray",
				Description:       "",
				Phone:             "400-8576-886",
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			}

			prod, err := CreateProduct(&product)
			So(err, ShouldBeNil)
			fmt.Println(prod)
		})
	})

	Convey("test get product1", t, func() {
		Convey("test get product", func() {
			product, err := GetProduct(&Product{
				ProductName: "kubesphere",
				//ProductID:   "product-9pjmqr3n4m7mmm",
			})
			So(err, ShouldBeNil)

			fmt.Println(product)
		})
	})
}

func TestGetProduct(t *testing.T) {
	Convey("test get product1", t, func() {
		Convey("test get product", func() {
			product, err := GetProduct(&Product{
				//ProductName: "kubesphere",
				ProductID: "product-9pjmqr3n4m7mmm",
			})
			So(err, ShouldBeNil)

			fmt.Println(product)
		})
	})
}

func TestDeleteProduct(t *testing.T) {
	Convey("test delete product", t, func() {
		Convey("test delete product", func() {
			err := DeleteProduct(&Product{
				ProductName: "kubesphere",
				//ProductID: "product-9pjmqr3n4m7mmm",
			})

			So(err, ShouldBeNil)
		})
	})
}



func TestUpdateProduct(t *testing.T) {
	Convey("test delete product", t, func() {
		Convey("test delete product", func() {
			err := UpdateProduct(&Product{
				MonitorCenterHost: "127.0.0.1",
				MonitorCenterPort: 80,
				ProductName:       "kubesphere",
				HomePage:          "https://www.kubesphere.io/",
				Email:             "xxxx@yunify.com",
				Address:           "xxxxxxxxxxx",
				Contacts:          "xxxxxxxxxxxx",
				Description:       "xxxxxxxxxx",
				Phone:             "xxxxxxxxxxxx",
				CreatedAt:         time.Now(),
				UpdatedAt:         time.Now(),
			})

			So(err, ShouldBeNil)
		})
	})
}