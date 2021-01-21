package datasource

import (
	"github.com/skyhackvip/risk_engine/configs"
)

type Features map[string]interface{}

var result = make(map[string]interface{})

//get depend feature.
//can be get from db or http curl
func GetFeatures(features []string) map[string]interface{} {
	for _, feature := range features {
		if _, ok := result[feature]; !ok {
			result[feature] = configs.Mock[feature]
		}
	}
	return result
}

//mock data no need to lock
func SetFeature(feature_name string, value interface{}) {
	result[feature_name] = value
}
