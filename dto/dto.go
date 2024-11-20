package dto

import (
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"github.com/yanghp/rule-client/pkg"
	"hash/fnv"
	"strconv"
	"time"
)

// Payload 与规则引擎的有强相关
type Payload struct {
	// 基础属性
	Os          string  `json:"os" schema:"os"`
	Channel     string  `json:"channel" schema:"channel"`
	PackageName string  `json:"package_name" schema:"package_name"`
	VersionCode uint32  `json:"version_code" schema:"version_code"`
	Imei        string  `json:"imei" schema:"imei"`
	Idfa        string  `json:"idfa" schema:"idfa"`
	Gaid        string  `json:"gaid" schema:"gaid"`
	Mac         string  `json:"mac" schema:"mac"`
	AndroidId   string  `json:"android_id" schema:"android_id"`
	UserId      string  `json:"user_id" schema:"user_id"`
	Sdkver      float32 `json:"sdkver" schema:"sdkver"`             //新增字段
	Country     string  `schema:"country" json:"country,omitempty"` //国家
	Ip          string  `json:"ip" schema:"ip"`
	Brand       string  `schema:"brand" json:"brand,omitempty"`
	Language    string  `schema:"language" json:"language,omitempty"`
	//广告相关字段
	AppName           string `json:"app_name" schema:"app_name"`
	AppId             string `json:"app_id" schema:"app_id"` //appid
	MediationPlatform int32  `json:"mediation_platform" schema:"mediation_platform"`
	AdType            int32  `json:"ad_type" schema:"ad_type"` //广告类型
	// 聚合平台的广告位 id
	PositionId string `json:"position_id" schema:"position_id"`
	// 我们自己的广告位 id
	Pid91 string `json:"pid91" schema:"pid91"`
	// 是否新用户
	IsNewUser bool    `json:"is_new_user" schema:"is_new_user"`
	ReqEcpm   float32 `json:"req_ecpm" schema:"req_ecpm"` // 素材请求的ecpm
	// router 相关
	RulePath     string `json:"-" schema:"-"` // 服务端使用的 rule path
	AppPath      string `json:"-" schema:"-"`
	RegisterTime string `json:"register_time" schema:"-"` // 用户注册时间
	// 广告数据
	UserAdData     UserAdData     `json:"user_ad_data,omitempty" schema:"user_ad_data"`           // 用户广告
	UserAdTypeData UserAdTypeData `json:"user_ad_type_data,omitempty" schema:"user_ad_type_data"` // 用户广告类型
	PositionAdData PositionAdData `json:"position_ad_data,omitempty" schema:"position_ad_data"`   // 聚合广告 id 维度
	UnionAdData    UnionAdData    `json:"union_ad_data,omitempty" schema:"union_ad_data"`         // 联盟广告 id
	AdTypeData     AdTypeData     `json:"ad_type_data,omitempty" schema:"ad_type_data"`           // 广告类型维度

	Q         map[string][]string    `json:"-" schema:"-"`
	B         map[string]interface{} `json:"-" schema:"-"`
	Context   context.Context        `json:"-" schema:"-"`
	RandomNum int32                  `json:"random_num" schema:"-"`
	GrowId    int64                  `json:"grow_id" schema:"-"` // 自定义用户增长id
}

type UserAdData struct {
	DayClick       uint32  `json:"day_click,omitempty" schema:"day_click"` //当天点击数
	DayExposure    uint32  `json:"day_exposure,omitempty" schema:"day_exposure"`
	DayIncome      float64 `json:"day_income,omitempty" schema:"day_income"`
	DayEcpm        float64 `json:"day_ecpm,omitempty" schema:"day_ecpm"`
	DayClickRate   float64 `json:"day_click_rate,omitempty" schema:"day_click_rate"`
	TotalClick     uint32  `json:"total_click,omitempty" schema:"total_click"`
	TotalExposure  uint32  `json:"total_exposure,omitempty" schema:"total_exposure"`
	TotalIncome    float64 `json:"total_income,omitempty" schema:"total_income"`
	TotalEcpm      float64 `json:"total_ecpm,omitempty" schema:"total_ecpm"`
	TotalClickRate float64 `json:"total_click_rate,omitempty" schema:"total_click_rate"`
}

// UserAdTypeData 用户广告类型
type UserAdTypeData struct {
	DayClick       uint32  `json:"day_click,omitempty" schema:"day_click"` //当天点击数
	DayExposure    uint32  `json:"day_exposure,omitempty" schema:"day_exposure"`
	DayIncome      float64 `json:"day_income,omitempty" schema:"day_income"`
	DayEcpm        float64 `json:"day_ecpm,omitempty" schema:"day_ecpm"`
	DayClickRate   float64 `json:"day_click_rate,omitempty" schema:"day_click_rate"`
	TotalClick     uint32  `json:"total_click,omitempty" schema:"total_click"`
	TotalExposure  uint32  `json:"total_exposure,omitempty" schema:"total_exposure"`
	TotalIncome    float64 `json:"total_income,omitempty" schema:"total_income"`
	TotalEcpm      float64 `json:"total_ecpm,omitempty" schema:"total_ecpm"`
	TotalClickRate float64 `json:"total_click_rate,omitempty" schema:"total_click_rate"`
}

type PositionAdData struct {
	DayClick       uint32  `json:"day_click,omitempty" schema:"day_click"` //当天点击数
	DayExposure    uint32  `json:"day_exposure,omitempty" schema:"day_exposure"`
	DayIncome      float64 `json:"day_income,omitempty" schema:"day_income"`
	DayEcpm        float64 `json:"day_ecpm,omitempty" schema:"day_ecpm"`
	DayClickRate   float64 `json:"day_click_rate,omitempty" schema:"day_click_rate"`
	TotalClick     uint32  `json:"total_click,omitempty" schema:"total_click"`
	TotalExposure  uint32  `json:"total_exposure,omitempty" schema:"total_exposure"`
	TotalIncome    float64 `json:"total_income,omitempty" schema:"total_income"`
	TotalEcpm      float64 `json:"total_ecpm,omitempty" schema:"total_ecpm"`
	TotalClickRate float64 `json:"total_click_rate,omitempty" schema:"total_click_rate"`
}

type AdTypeData struct {
	DayClick       uint32  `json:"day_click,omitempty" schema:"day_click"` //当天点击数
	DayExposure    uint32  `json:"day_exposure,omitempty" schema:"day_exposure"`
	DayIncome      float64 `json:"day_income,omitempty" schema:"day_income"`
	DayEcpm        float64 `json:"day_ecpm,omitempty" schema:"day_ecpm"`
	DayClickRate   float64 `json:"day_click_rate,omitempty" schema:"day_click_rate"`
	TotalClick     uint32  `json:"total_click,omitempty" schema:"total_click"`
	TotalExposure  uint32  `json:"total_exposure,omitempty" schema:"total_exposure"`
	TotalIncome    float64 `json:"total_income,omitempty" schema:"total_income"`
	TotalEcpm      float64 `json:"total_ecpm,omitempty" schema:"total_ecpm"`
	TotalClickRate float64 `json:"total_click_rate,omitempty" schema:"total_click_rate"`
}

type UnionAdData struct {
	DayClick       uint32  `json:"day_click,omitempty" schema:"day_click"` //当天点击数
	DayExposure    uint32  `json:"day_exposure,omitempty" schema:"day_exposure"`
	DayIncome      float64 `json:"day_income,omitempty" schema:"day_income"`
	DayEcpm        float64 `json:"day_ecpm,omitempty" schema:"day_ecpm"`
	DayClickRate   float64 `json:"day_click_rate,omitempty" schema:"day_click_rate"`
	TotalClick     uint32  `json:"total_click,omitempty" schema:"total_click"`
	TotalExposure  uint32  `json:"total_exposure,omitempty" schema:"total_exposure"`
	TotalIncome    float64 `json:"total_income,omitempty" schema:"total_income"`
	TotalEcpm      float64 `json:"total_ecpm,omitempty" schema:"total_ecpm"`
	TotalClickRate float64 `json:"total_click_rate,omitempty" schema:"total_click_rate"`
}

func string2UInt32(p string) uint32 {
	n, e := strconv.Atoi(p)
	if e != nil {
		return 0
	}
	return uint32(n)
}

func (p *Payload) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

func (p Payload) Now() time.Time {
	return time.Now().UTC()
}

func (p Payload) AB(a, b int) bool {
	if p.GrowId == 0 {
		return false
	}
	return pkg.AB(int(p.GrowId), a, b)
}

func (p Payload) Random() int32 {
	h := fnv.New64()
	h.Write([]byte(p.AndroidId))
	hashValue := h.Sum64()
	rest := hashValue % 100
	return int32(rest)
}

func (p Payload) Date(s string) time.Time {
	date, err := time.ParseInLocation("2006-01-02", s, time.Local)
	if err != nil {
		panic(err)
	}
	return date
}

func (p Payload) DaysAgo(s string) int {
	if s == "" {
		return 0
	}
	return int(time.Since(p.DateTime(s)).Hours() / 24)
}

func (p Payload) HoursAgo(s string) int {
	if s == "" {
		return 0
	}
	return int(time.Now().UTC().Sub(p.DateTime(s)).Hours())
}

func (p Payload) MinutesAgo(s string) int {
	if s == "" {
		return 0
	}
	return int(time.Now().UTC().Sub(p.DateTime(s)).Minutes())
}

func (p Payload) DateTime(s string) time.Time {
	date, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	if err != nil {
		panic(err)
	}
	return date
}

func (p Payload) IsBefore(s string) bool {
	var (
		t   time.Time
		err error
	)
	if len(s) == 10 {
		t, err = time.ParseInLocation("2006-01-02", s, time.Local)
	} else {
		t, err = time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	}
	if err != nil {
		panic(err)
	}
	return time.Now().UTC().Before(t)
}

func (p Payload) IsAfter(s string) bool {
	var (
		t   time.Time
		err error
	)
	if len(s) == 10 {
		t, err = time.ParseInLocation("2006-01-02", s, time.Local)
	} else {
		t, err = time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	}
	if err != nil {
		panic(err)
	}
	return time.Now().UTC().After(t)
}

func (p Payload) IsBetween(begin string, end string) bool {
	return p.IsAfter(begin) && p.IsBefore(end)
}

func (p Payload) IsWeekday(day int) bool {
	return time.Now().UTC().Weekday() == time.Weekday(day)
}

func (p Payload) IsWeekend() bool {
	if weekday := time.Now().UTC().Weekday(); weekday == 0 || weekday == 6 {
		return true
	}
	return false
}

func (p Payload) IsToday(s string) bool {
	return time.Now().UTC().Format("2006-01-02") == s
}

func (p Payload) IsHourRange(begin int, end int) bool {
	now := time.Now().UTC().Hour()
	return now >= begin && now <= end
}

func (p Payload) RegisterBefore(date string) bool {
	var (
		timeVal time.Time
		err     error
	)
	if len(date) == 10 {
		timeVal, err = time.Parse("2006-01-02", date)
	} else {
		timeVal, err = time.Parse("2006-01-02 15:04:05", date)
	}
	if err != nil {
		return false
	}
	registerTime, e := time.Parse("2006-01-02 15:04:05", p.RegisterTime)
	if e != nil {
		return false
	}
	return registerTime.Before(timeVal)
}

func (p Payload) RegisterAfter(date string) bool {
	var (
		timeVal time.Time
		err     error
	)
	if len(date) == 10 {
		timeVal, err = time.Parse("2006-01-02", date)
	} else {
		timeVal, err = time.Parse("2006-01-02 15:04:05", date)
	}
	if err != nil {
		return false
	}
	registerTime, e := time.Parse("2006-01-02 15:04:05", p.RegisterTime)
	if e != nil {
		return false
	}
	return registerTime.After(timeVal)
}

func (p Payload) RegisterBetween(begin, end string) bool {
	return p.RegisterAfter(begin) && p.RegisterBefore(end)
}

func (p Payload) ToString(str interface{}) string {
	return cast.ToString(str)
}

func (p Payload) ToInt(int interface{}) int {
	return cast.ToInt(int)
}

type Data map[string]interface{}

type Response struct {
	Code    uint `json:"code"`
	Message uint `json:"message"`
	Data    Data `json:"event"`
}

func (p Response) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}
