package alertmanager

//import (
//	"testing"
//	"time"
//	"github.com/prometheus/common/model"
//	"github.com/prometheus/alertmanager/types"
//	"github.com/prometheus/alertmanager/api/v1"
//	"net/http"
//	"bytes"
//	"net/http/httptest"
//	"io/ioutil"
//	"github.com/stretchr/testify/require"
//	"fmt"
//	"encoding/json"
//	"errors"
//)
//
//func TestAddAlerts(t *testing.T) {
//	now := func(offset int) time.Time {
//		return time.Now().Add(time.Duration(offset) * time.Second)
//	}
//
//	for i, tc := range []struct {
//		start, end time.Time
//		err        bool
//		code       int
//	}{
//		{time.Time{}, time.Time{}, false, 200},
//		{now(0), time.Time{}, false, 200},
//		{time.Time{}, now(-1), false, 200},
//		{time.Time{}, now(0), false, 200},
//		{time.Time{}, now(1), false, 200},
//		{now(-2), now(-1), false, 200},
//		{now(1), now(2), false, 200},
//		{now(1), now(0), false, 400},
//		{now(0), time.Time{}, true, 500},
//	} {
//		alerts := []model.Alert{{
//			StartsAt:    tc.start,
//			EndsAt:      tc.end,
//			Labels:      model.LabelSet{"label1": "test1"},
//			Annotations: model.LabelSet{"annotation1": "some text"},
//		}}
//		b, err := json.Marshal(&alerts)
//		if err != nil {
//			t.Errorf("Unexpected error %v", err)
//		}
//
//		alertsProvider := newFakeAlerts([]*types.Alert{}, tc.err)
//		api := v1.New(alertsProvider, nil, newGetAlertStatus(alertsProvider), nil, nil)
//
//		r, err := http.NewRequest("POST", "/api/v1/alerts", bytes.NewReader(b))
//		w := httptest.NewRecorder()
//		if err != nil {
//			t.Errorf("Unexpected error %v", err)
//		}
//
//		api.addAlerts(w, r)
//		res := w.Result()
//		body, _ := ioutil.ReadAll(res.Body)
//
//		require.Equal(t, tc.code, w.Code, fmt.Sprintf("test case: %d, StartsAt %v, EndsAt %v, Response: %s", i, tc.start, tc.end, string(body)))
//	}
//}
//
//
//func newFakeAlerts(alerts []*types.Alert, withErr bool) * {
//	fps := make(map[model.Fingerprint]int)
//	for i, a := range alerts {
//		fps[a.Fingerprint()] = i
//	}
//	f := &fakeAlerts{
//		alerts: alerts,
//		fps:    fps,
//	}
//	if withErr {
//		f.err = errors.New("Error occured")
//	}
//	return f
//}
