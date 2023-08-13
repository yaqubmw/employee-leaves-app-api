package helper

import (
	"strings"
)

// MatchKeyword checks if any of the given keywords is found in the text.
func MatchKeyword(leaveTypeName string) string {

	annualMatcher := []string{"annual", "tahunan"}
	maternityMatcher := []string{"maternity", "hamil", "mengandung"}
	marriageMatcher := []string{"marriage", "menikah", "nikah"}
	menstrualMatcher := []string{"menstrual", "haid", "period", "mens"}
	paternityMatcher := []string{"paternity", "istri"}

	for _, keyword := range annualMatcher {
		if strings.Contains(strings.ToLower(leaveTypeName), strings.ToLower(keyword)) {
			return "annual"
		}
	}

	for _, keyword := range maternityMatcher {
		if strings.Contains(strings.ToLower(leaveTypeName), strings.ToLower(keyword)) {
			return "maternity"
		}
	}

	for _, keyword := range marriageMatcher {
		if strings.Contains(strings.ToLower(leaveTypeName), strings.ToLower(keyword)) {
			return "marriage"
		}
	}

	for _, keyword := range menstrualMatcher {
		if strings.Contains(strings.ToLower(leaveTypeName), strings.ToLower(keyword)) {
			return "menstrual"
		}
	}

	for _, keyword := range paternityMatcher {
		if strings.Contains(strings.ToLower(leaveTypeName), strings.ToLower(keyword)) {
			return "paternity"
		}
	}
	return ""

}
