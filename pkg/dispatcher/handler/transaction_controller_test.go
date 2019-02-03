package handler

import (
	"fmt"
	"kubesphere.io/ks-alert/pkg/models"
	"kubesphere.io/ks-alert/pkg/utils/dbutil"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCallReflect(t *testing.T) {
	Convey("test reflact call", t, func() {

		db, _ := dbutil.DBClient()

		tx := db.Begin()

		values, err := CallReflect(models.ReceiverGroup{}, "Create", tx, "hi jack")
		fmt.Println(values, err)

	})
}
