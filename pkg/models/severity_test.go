package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCreateSeverities(t *testing.T) {
	Convey("create severities", t, func() {

		productID := "product-4llxr47k7q82wz"

		err := CreateSeverities(&[]Severity{
			Severity{
				ProductID:      productID,
				CreatedBy:      "carmanzhang",
				SeverityEn:     "critical",
				SeverityCh:     "严重",
				SeverityDegree: 3,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},

			Severity{
				ProductID:      productID,
				CreatedBy:      "carmanzhang",
				SeverityEn:     "severity",
				SeverityCh:     "较严重",
				SeverityDegree: 2,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},

			Severity{
				ProductID:      productID,
				CreatedBy:      "carmanzhang",
				SeverityEn:     "warning",
				SeverityCh:     "警告",
				SeverityDegree: 1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
		})

		So(err, ShouldBeNil)

	})
}
