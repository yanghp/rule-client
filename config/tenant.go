package config

import (
	"context"
)

var TenantKey = struct{}{}

type Tenant struct {
	Channel     string `json:"channel"`
	VersionCode uint32 `json:"version_code"`
	Os          string `json:"os"`
	UserId      string `json:"user_id"`
	Imei        string `json:"imei"`
	Idfa        string `json:"idfa"`
	Oaid        string `json:"oaid"`
	Mac         string `json:"mac"`
	AndroidId   string `json:"android_id"`
	PackageName string `json:"package_name"`
	Ip          string `json:"ip"`
	VersionName string `json:"version_name"` //新增版本号
	Orientation uint8  `json:"orientation"`  //新增横竖屏
	Sdkver      uint32 `json:"sdkver"`       //新增版本号
	AppId       string `json:"app_id" `      //媒体ID
	AdPlatform  uint8  `json:"ad_platform" ` //广告平台
	AdType      uint8  `json:"ad_type" `     //广告类型
	AdCode      string `json:"ad_code" `     //聚合的广告位
	CodeId      string `json:"code_id" `     //联盟的广告位
	DClick      uint8  `json:"d_click" `     //当天点击数
	DExposure   uint8  `json:"d_exposure" `  //当天展示数

	Context context.Context `json:"-"`
}

func GetTenant(ctx context.Context) *Tenant {
	if c, ok := ctx.Value(TenantKey).(*Tenant); ok {
		c.Context = ctx
		return c
	}
	return &Tenant{Context: context.Background()}
}

type DynamicConfigReader interface {
	Tenant(tenant *Tenant) (ConfigReader, error)
}
