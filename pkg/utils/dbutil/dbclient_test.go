package dbutil

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDBClient(t *testing.T) {
	Convey("test database", t, func() {

		Convey("test database connection", func() {
			_, err := DBClient()
			So(err, ShouldBeNil)
		})

		//Convey("test database insert0", func() {
		//	dbClient, err := DBClient()
		//	So(err, ShouldBeNil)
		//	var enterprise = models.Enterprise{
		//		EnterpriseID:   idutil.GetUuid36(""),
		//		EnterpriseName: "北京优帆科技有限公司武汉分公司",
		//		HomePage:       "https://www.qingcloud.com/",
		//		Address:        "北京优帆科技有限公司",
		//		Email:          "yunify@yunify.com",
		//		Contacts:       "Richard",
		//		Description:    "云计算公司",
		//		Phone:          "400-8576-886",
		//		CreatedAt:      time.Now(),
		//		UpdatedAt:      time.Now(),
		//	}
		//	err = dbClient.Create(enterprise).Error
		//	So(err, ShouldBeNil)
		//})
		//
		//Convey("test database select1", func() {
		//	dbClient, err := DBClient()
		//	So(err, ShouldBeNil)
		//	var count int
		//	dbClient.Model(&models.Enterprise{}).Count(&count)
		//	//So(count, ShouldEqual, 2)
		//
		//	var enterprises []models.Enterprise
		//	err = db.Model(&models.Enterprise{}).Where(&models.Enterprise{Contacts: "Richard"}).Find(&enterprises).Error
		//	So(err, ShouldBeNil)
		//	for i, _ := range enterprises {
		//		//fmt.Println(enterprises[i])
		//		fmt.Println(enterprises[i].UpdatedAt.Unix())
		//		fmt.Println(enterprises[i].CreatedAt.Unix())
		//	}
		//})
		//
		//Convey("test database insert2", func() {
		//	dbClient, err := DBClient()
		//	So(err, ShouldBeNil)
		//	var ent models.Enterprise
		//	dbClient.Model(&models.Enterprise{}).Where(&models.Enterprise{EnterpriseName:"北京优帆科技有限公司武汉分公司"}).First(&ent)
		//	fmt.Println(ent.EnterpriseID)
		//	var product = models.Product{
		//		ProductID:         idutil.GetUuid36(""),
		//		EnterpriseID:      ent.EnterpriseID,
		//		MonitorCenterHost: "localhost",
		//		MonitorCenterPort: 8080,
		//		ProductName:       "kubesphere",
		//		HomePage:          "https://www.kubesphere.io/",
		//		Email:             "ray@yunify.com",
		//		Contacts:          "Ray",
		//		Description:       "应用平台研发部",
		//		Phone:             "400-8576-886",
		//		CreatedAt:         time.Now(),
		//		UpdatedAt:         time.Now(),
		//	}
		//	err = dbClient.Create(product).Error
		//	So(err, ShouldBeNil)
		//})

	})
}

//func TestDBClientInsert(t *testing.T) {
//	Convey("test database", t, func() {
//
//		Convey("test database insert3", func() {
//			dbClient, err := DBClient()
//			So(err, ShouldBeNil)
//
//			var product = models.Product{
//				ProductID:         idutil.GetUuid36(""),
//				EnterpriseID:      idutil.GetUuid36(""),
//				MonitorCenterHost: "localhost",
//				MonitorCenterPort: 8080,
//				ProductName:       "kubesphere",
//				HomePage:          "https://www.kubesphere.io/",
//				Email:             "ray@yunify.com",
//				Contacts:          "Ray",
//				Description:       "应用平台研发部",
//				Phone:             "400-8576-886",
//				CreatedAt:         time.Now(),
//				UpdatedAt:         time.Now(),
//			}
//			err = dbClient.Create(product).Error
//			So(err, ShouldBeNil)
//		})
//	})
//}
