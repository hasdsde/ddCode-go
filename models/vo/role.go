package vo

import "ddCode-server/models/dto"

func RolePolicyToSlice(roleName string, policies []dto.PolicyAddParam) [][]string {
	result := make([][]string, len(policies))
	for i := 0; i < len(policies); i++ {
		result[i] = []string{roleName, policies[i].Url, policies[i].Method}
	}
	return result
}
