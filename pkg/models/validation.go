package models

import (
	"fmt"
	"regexp"
)

func ValidateName(name string) error {
	illegalChars := `[!@#$%&*()^~+.=,\-/\\[\]{};:'"<>?]`
	if matched, _ := regexp.MatchString(illegalChars, name); matched {
		return fmt.Errorf("the %s contain invalid chars", name)
	}
	return nil
}
