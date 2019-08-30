package ios

import (
	"github.com/u-push"
)

type IOSBroadcast struct {
	*u_push.IOSNotification
}

func NewIOSBroadcast(appkey, appMasterSecret string) IOSBroadcast {
	var t IOSBroadcast
	t.IOSNotification = u_push.NewIOSNotification()
	t.SetAppMasterSecret(appMasterSecret)
	t.SetPredefinedKeyValue("appkey", appkey)
	t.SetPredefinedKeyValue("type", "broadcast")
	return t
}
