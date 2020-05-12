package faulttolerance

import (
	"context"
	"encoding/json"
	"errors"
	"mosn.io/api"
	v2 "mosn.io/mosn/pkg/config/v2"
	"mosn.io/mosn/pkg/filter/stream/faulttolerance/config"
	"strconv"
)

func init() {
	api.RegisterStream(v2.FaultTolerance, CreateSendFilterFactory)
}

type SendFilterFactory struct {
	config *v2.FaultToleranceFilterConfig
}

func (f *SendFilterFactory) CreateFilterChain(context context.Context, callbacks api.StreamFilterChainFactoryCallbacks) {
	filter := NewSendFilter(f.config)
	callbacks.AddStreamSenderFilter(filter)
}

func CreateSendFilterFactory(conf map[string]interface{}) (api.StreamFilterChainFactory, error) {
	if filterConfig, err := parseConfig(conf); err != nil {
		return nil, err
	} else {
		return &SendFilterFactory{
			config: filterConfig,
		}, nil
	}
}

func parseConfig(cfg map[string]interface{}) (*v2.FaultToleranceFilterConfig, error) {
	ruleJson := &config.FaultToleranceRuleJson{}
	data, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, ruleJson); err != nil {
		return nil, err
	}

	fillDefaultValue(ruleJson)

	if !isLegal(ruleJson) {
		return nil, errors.New("config is illegal")
	}

	exceptionTypes := map[uint32]bool{}
	for _, exceptionType := range ruleJson.ExceptionTypes {
		if code, err := strconv.ParseUint(exceptionType, 10, 32); err == nil {
			exceptionTypes[uint32(code)] = true
		}
	}
	filterConfig := &v2.FaultToleranceFilterConfig{
		Enabled:               ruleJson.Enabled,
		ExceptionTypes:        exceptionTypes,
		TimeWindow:            ruleJson.TimeWindow,
		LeastWindowCount:      ruleJson.LeastWindowCount,
		ExceptionRateMultiple: ruleJson.ExceptionRateMultiple,
		MaxIpCount:            ruleJson.MaxIpCount,
		MaxIpRatio:            ruleJson.MaxIpRatio,
		RecoverTime:           ruleJson.RecoverTime,
	}
	return filterConfig, nil
}

func isLegal(ruleJson *config.FaultToleranceRuleJson) bool {
	if ruleJson.TimeWindow < 10 {
		return false
	}
	if ruleJson.LeastWindowCount <= 0 {
		return false
	}
	if ruleJson.MaxIpCount < 0 {
		return false
	}
	if ruleJson.ExceptionRateMultiple <= 1 {
		return false
	}
	return true
}

func fillDefaultValue(ruleJson *config.FaultToleranceRuleJson) {
	if ruleJson.TimeWindow == 0 {
		ruleJson.TimeWindow = 10000
	}
	if ruleJson.ExceptionRateMultiple == 0 {
		ruleJson.ExceptionRateMultiple = 5
	}
	if ruleJson.ExceptionTypes == nil {
		ruleJson.ExceptionTypes = []string{"502", "503", "504"}
	}
	if ruleJson.MaxIpCount == 0 {
		ruleJson.MaxIpCount = 1
	}
	if ruleJson.RecoverTime == 0 {
		ruleJson.RecoverTime = 15 * 60000
	}
	if ruleJson.TaskSize == 0 {
		ruleJson.TaskSize = 20
	}
}