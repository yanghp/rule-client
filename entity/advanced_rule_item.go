package entity

import (
	"fmt"
	"github.com/hashicorp/go-multierror"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
	"github.com/knadh/koanf"
	"github.com/pkg/errors"
	"github.com/xeipuuv/gojsonschema"
	"github.com/yanghp/rule-client/dto"
	"github.com/yanghp/rule-client/msg"
)

type AdvancedRuleItem struct {
	iff     string
	then    dto.Data
	child   Ruler
	program *vm.Program
}

func (ar *AdvancedRuleItem) ValidateWithSchema(schema gojsonschema.JSONLoader) error {
	if ar.then == nil && ar.child != nil {
		return ar.child.ValidateWithSchema(schema)
	}
	document := gojsonschema.NewGoLoader(ar.then)
	result, err := gojsonschema.Validate(schema, document)
	if err != nil {
		return errors.Wrap(err, "fails to validate with json schema")
	}
	if !result.Valid() {
		var err multierror.Error
		for i := range result.Errors() {
			if result.Errors()[i] != nil {
				err.Errors = append(err.Errors, fmt.Errorf(result.Errors()[i].String()))
			}
		}
		if err.Len() > 0 {
			return &err
		}
		return nil
	}
	return nil
}

func (ar *AdvancedRuleItem) Unmarshal(reader *koanf.Koanf) error {
	ar.iff = reader.MustString("if")
	if len(ar.iff) == 0 {
		return errors.New("if condition not found in advanced rule")
	}
	err := reader.Unmarshal("then", &ar.then)
	if err != nil {
		return err
	}
	if ar.then == nil && reader.Exists("child") {
		style := reader.MustString("child.style")
		if style == "" {
			return errors.New("missing child style")
		}
		item, err := NewRuler(style)
		if err != nil {
			return err
		}
		err = item.Unmarshal(reader.Cut("child"))
		if err != nil {
			return err
		}
		ar.child = item
	}
	return nil
}

func (ar *AdvancedRuleItem) Compile() error {
	var err error
	ar.then = convert(ar.then)
	ar.program, err = expr.Compile(ar.iff, expr.Env(&dto.Payload{}))
	if err != nil {
		return err
	}
	if ar.program == nil {
		return fmt.Errorf("invalid expression: %s", ar.iff)
	}
	if ar.child != nil {
		if err = ar.child.Compile(); err != nil {
			return err
		}
	}
	return err
}

func (br *AdvancedRuleItem) Calculate(payload *dto.Payload) (dto.Data, error) {
	output, err := expr.Run(br.program, payload)
	if err != nil {
		return nil, errors.Wrap(err, msg.ErrorRules)
	}
	if i, ok := output.(int); ok && i == 0 {
		return nil, nil
	}
	if b, ok := output.(bool); ok && !b {
		return nil, nil
	}
	if br.then != nil {
		return br.then, nil
	}
	if br.child != nil {
		return br.child.Calculate(payload)
	}
	return nil, nil
}
