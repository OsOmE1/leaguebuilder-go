package datadragon

import (
	"encoding/json"
	"fmt"
	"github.com/cespare/xxhash"
	"regexp"
	"strings"
)

func JsonMapToInterface(m map[string]interface{}, i interface{}) error {
	d, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(d, i)
}

func GetGroups(regEx, s string) (groups []string) {
	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(s)

	for i, _ := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			groups = append(groups, match[i])
		}
	}
	return groups
}

func RemoveDuplicateStr(strSlice []string) (slice []string) {
	allKeys := make(map[string]bool)
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			slice = append(slice, item)
		}
	}
	return slice
}

func GetBinHash(s string) string {
	h := 0x811c9dc5
	for _, b := range strings.ToLower(s) {
		h = ((h ^ int(b)) * 0x01000193) % 0x100000000
	}
	return fmt.Sprintf("%x", h)
}

func GetXXH64(s string) string {
	h := xxhash.Sum64String(strings.ToLower(s))
	h = h & ((1 << 40) - 1)
	return fmt.Sprintf("%x", h)
}

func GetValueFromBinMap(key string, m map[string]interface{}) interface{} {
	if val, ok := m[key]; ok {
		return val
	}
	if val, ok := m[fmt.Sprintf("{%s}", GetBinHash(key))]; ok {
		return val
	}
	return nil
}
