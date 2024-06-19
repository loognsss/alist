package halalcloud

import (
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/op"
)

type Addition struct {
	// Usually one of two
	driver.RootID
	// define other
	RefreshToken         string `json:"refresh_token" required:"true" help:"login type is refresh_token,this is required"`
	UploadThread         string `json:"upload_thread" default:"3" help:"1<=thread<=32"`
	CustomUploadPartSize string `json:"custom_upload_part_size" default:"0" help:"0 for auto"`
}

var config = driver.Config{
	Name:              "HalalCloud",
	LocalSort:         false,
	OnlyLocal:         true,
	OnlyProxy:         true,
	NoCache:           false,
	NoUpload:          false,
	NeedMs:            false,
	DefaultRoot:       "/",
	CheckStatus:       false,
	Alert:             "",
	NoOverwriteUpload: false,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &HalalCloud{}
	})
}
