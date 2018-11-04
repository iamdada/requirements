package requirements

import (
	"fmt"
)

type Rule struct {
	Name           string
	Variable       string
	Value          int32
	Comparison     string
	Interests      []string
	Decline_reason string
	Passed         bool
}

type DecisionModuleResult struct {
	Passed bool
}

type DecisionInput struct {
	InputVariable int32
}

// {
// 	Name:           "minimum age requirement",
// 	Variable:       "age",
// 	Comparison:     ">",
// 	Value:          17,
// 	Interests:      []string{"art", "coding", "music", "travel"},
// 	Decline_reason: "Failed Age Requirement",
// 	Passed:         true,
// },
func buildScript(rule Rule) func(DecisionInput) bool {
	if rule.Comparison == ">" {
		return func(state DecisionInput) bool {
			if state.InputVariable > rule.Value {
				return true
			} else {
				return false
			}
		}
	} else {
		return func(state DecisionInput) bool { return false }
	}
}

func GenerateRule(rules []Rule) []func(DecisionInput) bool {
	ruleFunc := make([]func(DecisionInput) bool, len(rules))
	for index, rule := range rules {
		ruleFunc[index] = buildScript(rule)
	}
	return ruleFunc
}

func RunRule(rulesFuncs []func(DecisionInput) bool, state DecisionInput) DecisionModuleResult {
	passed := false
	for _, ruleFunc := range rulesFuncs {
		fmt.Println(ruleFunc(state))
		passed = passed || ruleFunc(state)
	}
	result := DecisionModuleResult{Passed: passed}
	return result
}
