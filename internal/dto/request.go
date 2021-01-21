package dto

/**
 * dsl run request
 * example: {"flow":"conditional","features":{"feature_1":5,"feature_2":3,"feature_3":true}}
 */
type DslRunRequest struct {
	Flow     string                 `json:"flow"`
	Features map[string]interface{} `json:"features"`
}
