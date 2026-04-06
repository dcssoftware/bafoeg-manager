package httpparams

import (
	"errors"
	"fmt"
	"strconv"

	safecast "github.com/ccoveille/go-safecast/v2"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
)

func GetParamsPageUint(queries map[string]string) (uint, error) {
	dataString, ok := queries["page"]
	if !ok {
		return 1, nil
	}

	resultInt, resultIntErr := strconv.ParseUint(dataString, 10, 32)
	if resultIntErr != nil {
		return 0, resultIntErr
	}

	result, resultErr := safecast.Convert[uint](resultInt)
	return result, resultErr
}

func GetParamsBoolean(queries map[string]string, identifier string, defaultOrRequired *bool) (bool, error) {
	dataString, ok := queries[identifier]
	if !ok && defaultOrRequired == nil {
		logger.Debug("", fmt.Sprintf("get param '%s' is not defined in url", identifier), "")
		return false, errors.New("param is not defined")
	} else if !ok {
		return *defaultOrRequired, nil
	}

	resultBool, resultBoolErr := strconv.ParseBool(dataString)
	if resultBoolErr != nil {
		return false, resultBoolErr
	}

	return resultBool, nil
}
