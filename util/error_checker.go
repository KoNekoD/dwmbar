package util

import (
	"fmt"
	"go.uber.org/zap"
)

type ErrorChecker struct {
	logger *zap.Logger
}

func (c *ErrorChecker) ErrorFound(err error) bool {
	if err == nil {
		return false
	}

	c.logger.Error(
		fmt.Sprintf("ErrorChecker error: %s", err.Error()),
		zap.NamedError("Error", err),
	)

	return true
}

func NewErrorChecker(logger *zap.Logger) *ErrorChecker {
	return &ErrorChecker{
		logger: logger,
	}
}
