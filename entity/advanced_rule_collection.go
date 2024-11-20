package entity

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/knadh/koanf"
	"github.com/xeipuuv/gojsonschema"
	"github.com/yanghp/rule-client/dto"
)

type AdvancedRuleCollection struct {
	style  string
	enrich bool
	items  []*AdvancedRuleItem
}

func NewAdvancedRule() *AdvancedRuleCollection {
	return &AdvancedRuleCollection{
		style: "advanced",
		items: nil,
	}
}

func (ar *AdvancedRuleCollection) ValidateWithSchema(schema gojsonschema.JSONLoader) error {
	var err multierror.Error
	for i := range ar.items {
		errors := ar.items[i].ValidateWithSchema(schema)
		if errors != nil {
			err.Errors = append(err.Errors, errors)
		}
	}
	if err.Len() > 0 {
		return &err
	}
	return nil
}

func (ar *AdvancedRuleCollection) ShouldEnrich() bool {
	return ar.enrich
}

func (ar *AdvancedRuleCollection) Unmarshal(reader *koanf.Koanf) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %s", r)
		}
	}()
	ar.style = reader.String("style")
	ar.enrich = reader.Bool("enrich")
	slc := reader.Slices("rule")
	for _, subReader := range slc {
		var item AdvancedRuleItem
		err := item.Unmarshal(subReader)
		if err != nil {
			return err
		}
		ar.items = append(ar.items, &item)
	}
	return nil
}

func (ar *AdvancedRuleCollection) Compile() error {
	var err error
	for i := range ar.items {
		err = ar.items[i].Compile()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ar *AdvancedRuleCollection) Calculate(payload *dto.Payload) (dto.Data, error) {
	var rest dto.Data
	var index any
	for i, item := range ar.items {
		data, err := item.Calculate(payload)
		if err != nil {
			return nil, err
		}
		if data != nil {
			rest = data
			if i == len(ar.items)-1 {
				index = "last"
			} else {
				index = i + 1
			}

			break
		}
	}
	// 复制避免 data race
	if rest != nil {
		newMap := make(map[string]interface{}, len(rest)+1)
		for k, v := range rest {
			newMap[k] = v
		}
		newMap["abRule"] = index
		return newMap, nil
	}
	return dto.Data{}, nil
}
