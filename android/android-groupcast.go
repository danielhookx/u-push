package android

import (
	"encoding/json"
	"github.com/u-push"
)

type AndroidGroupcast struct {
	*u_push.AndroidNotification
}

func NewAndroidGroupcast(appkey, appMasterSecret string) AndroidGroupcast {
	var t AndroidGroupcast
	t.AndroidNotification = u_push.NewAndroidNotification()
	t.SetAppMasterSecret(appMasterSecret)
	t.SetPredefinedKeyValue("appkey", appkey)
	t.SetPredefinedKeyValue("type", "groupcast")
	return t
}

func (t *AndroidGroupcast) SetFilter(filter json.RawMessage) {
	t.SetPredefinedKeyValue("filter", filter)
}
