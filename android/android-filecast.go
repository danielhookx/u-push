package android

import (
	"github.com/u-push"
)

type AndroidFilecast struct {
	*u_push.AndroidNotification
}

func NewAndroidFilecast(appkey, appMasterSecret string) AndroidFilecast {
	var t AndroidFilecast
	t.AndroidNotification = u_push.NewAndroidNotification()
	t.SetAppMasterSecret(appMasterSecret)
	t.SetPredefinedKeyValue("appkey", appkey)
	t.SetPredefinedKeyValue("type", "filecast")
	return t
}

func (t *AndroidFilecast) SetFileId(fileId string) {
	t.SetPredefinedKeyValue("file_id", fileId)
}
