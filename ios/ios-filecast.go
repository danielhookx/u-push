package ios

import (
	"github.com/u-push"
)

type IOSFilecast struct {
	*u_push.IOSNotification
}

func NewIOSFilecast(appkey, appMasterSecret string) IOSFilecast {
	var t IOSFilecast
	t.IOSNotification = u_push.NewIOSNotification()
	t.SetAppMasterSecret(appMasterSecret)
	t.SetPredefinedKeyValue("appkey", appkey)
	t.SetPredefinedKeyValue("type", "filecast")
	return t
}

func (t *IOSFilecast) SetFileId(fileId string) {
	t.SetPredefinedKeyValue("file_id", fileId)
}
