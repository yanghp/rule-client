package dto

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
)

type Decoder struct {
	decoder *schema.Decoder
}

func NewDecoder() *Decoder {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	decoder.SetAliasTag("json")
	return &Decoder{decoder: decoder}
}

var decoder *schema.Decoder

func init() {
	decoder = schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
}

func (d *Decoder) Decode(payload *Payload, r *http.Request) error {
	if r.Method == "POST" {
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return errors.Wrapf(err, "cannot read body of http request")
		}
		err = json.Unmarshal(buf, payload)
		if err != nil {
			return errors.Wrap(err, "cannot json unmarshal")
		}
		err = json.Unmarshal(buf, &payload.B)
		if err != nil {
			return errors.Wrap(err, "cannot json unmarshal")
		}
		return nil
	}
	query := r.URL.Query()
	err := d.decoder.Decode(payload, query)
	if err != nil {
		return errors.Wrap(err, "fails to decode")
	}
	//上下兼容其它字段
	compatParamToPayLoad(payload, &query)
	// store extra queries here.
	payload.Q = query
	return nil
}

func DecodeFromRequest(payload *Payload, r *http.Request) error {

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.Wrapf(err, "cannot read body of http request")
	}
	defer r.Body.Close()
	err = json.Unmarshal(buf, payload)
	if err != nil {
		return errors.Wrap(err, "cannot json unmarshal")
	}
	err = json.Unmarshal(buf, &payload.B)
	if err != nil {
		return errors.Wrap(err, "cannot json unmarshal")
	}

	query := r.URL.Query()
	path := r.Header.Get("X-Config")
	payload.RulePath = path
	//上下兼容其它字段
	compatParamToPayLoad(payload, &query)
	// store extra queries here.
	payload.Q = query
	return nil
}

func compatParamToPayLoad(payload *Payload, query *url.Values) {
	//处理字段兼容问题
	if payload.PackageName == "" {
		payload.PackageName = query.Get("pkgname")
	}

	if payload.VersionCode == 0 {
		payload.VersionCode = string2UInt32(query.Get("pkgver"))
	}

	if payload.AndroidId == "" {
		payload.AndroidId = query.Get("dpid")
	}
}
