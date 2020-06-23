package plog

import (
	"fmt"
)

var (
	InfoSeverityLevel    = severityName[infoLog]
	WarningSeverityLevel = severityName[warningLog]
	ErrorSeverityLevel   = severityName[errorLog]
	FatalSeverityLevel   = severityName[fatalLog]
)

// Hook to be fired when logging to a severity
type Hook interface {
	// SeverityLevel returns "INFO", "WARNING", "ERROR" or "FATAL"
	// Hook will be fired in all the cases when severity is greater than
	// or equal to the severity level
	SeverityLevel() string

	// Fire implements the actual hook task that needs to be triggered
	Fire(s string, args ...interface{}) error
}

func GetSeverityNames() []string {
	var severityNames []string
	for _, s := range severityName {
		severityNames = append(severityNames, s)
	}

	return severityNames
}

func IsSeverityLevelSupported(severityLevel string) error {
	severityNames := GetSeverityNames()
	for _, s := range severityNames {
		if s == severityLevel {
			return nil
		}
	}

	return fmt.Errorf(
		"not supported severity level: %s, supported severity levels are: %v",
		severityLevel,
		severityNames,
	)
}

type Hooks map[string][]Hook

func (hooks Hooks) Add(hook Hook) error {
	severityLevel := hook.SeverityLevel()
	err := IsSeverityLevelSupported(severityLevel)
	if err != nil {
		return err
	}
	hooks[severityLevel] = append(hooks[severityLevel], hook)

	return nil
}

func (hooks Hooks) Fire(s severity, args ...interface{}) error {
	for severityLevel, severityHooks := range hooks {
		level, ok := severityByName(severityLevel)
		if !ok {
			return fmt.Errorf(
				"error getting severity name for: %s", severityLevel,
			)
		}

		if s < level {
			continue
		}

		for _, hook := range severityHooks {
			if err := hook.Fire(severityName[s], args); err != nil {
				return err
			}
		}
	}

	return nil
}
