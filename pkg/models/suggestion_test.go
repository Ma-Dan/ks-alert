package models

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestGetSuggestion(t *testing.T) {
	Convey("get suggestion", t, func() {
		x, err := GetSuggestion(&Suggestion{
			AlertConfigID: "111111111",
			AlertRuleID:   "222222222",
			ResourceID:    "333333333",
		})
		fmt.Println(x)
		fmt.Println(err)
	})
}

func TestUpdateSuggestion(t *testing.T) {
	Convey("update suggestion", t, func() {
		x, err := UpdateSuggestion(&Suggestion{
			AlertConfigID: "111111111",
			AlertRuleID:   "222222222",
			ResourceID:    "333333333",
			Message:       "messgaeeeeeeeeeeeeeeeeeeeeee",
			UpdatedAt:     time.Now(),
			CreatedAt:     time.Now(),
		})
		fmt.Println(x)
		fmt.Println(err)
	})
}
