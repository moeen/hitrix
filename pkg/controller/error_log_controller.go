package controller

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/coretrix/hitrix/pkg/response"
	"github.com/coretrix/hitrix/service"
	errorlogger "github.com/coretrix/hitrix/service/component/error_logger"
)

type ErrorLogController struct {
}

func (controller *ErrorLogController) GetErrors(c *gin.Context) {
	ormService := service.DI().OrmEngineForContext(c.Request.Context())

	type errorRow struct {
		File    string
		Line    int
		AppName string
		Request string
		Message string
		Stack   string
		Counter int
		Time    string
	}

	data := ormService.GetRedis().HGetAll(errorlogger.GroupError)

	errorsList := map[string]*errorRow{}

	for key, value := range data {
		// TODO: fix this hack
		if len(value) == 0 {
			continue
		}

		splitKeys := strings.Split(key, ":")

		if _, ok := errorsList[splitKeys[0]]; !ok {
			errorsList[splitKeys[0]] = &errorRow{}
		}

		if len(splitKeys) == 1 {
			errorMessage := &errorlogger.ErrorMessage{}
			err := json.Unmarshal([]byte(value), errorMessage)
			if err != nil {
				panic(err)
			}
			errorsList[splitKeys[0]].Request = string(errorMessage.Request)
			errorsList[splitKeys[0]].Stack = string(errorMessage.Stack)
			errorsList[splitKeys[0]].File = errorMessage.File
			errorsList[splitKeys[0]].Message = errorMessage.Message
			errorsList[splitKeys[0]].Line = errorMessage.Line
			errorsList[splitKeys[0]].AppName = errorMessage.AppName
		} else if len(splitKeys) == 2 {
			if splitKeys[1] == "time" {
				i, _ := strconv.ParseInt(value, 10, 64)
				errorsList[splitKeys[0]].Time = time.Unix(i, 0).String()
			} else if splitKeys[1] == "counter" {
				counter, _ := strconv.Atoi(value)
				errorsList[splitKeys[0]].Counter = counter
			}
		}
	}

	response.SuccessResponse(c, errorsList)
}

func (controller *ErrorLogController) DeleteError(c *gin.Context) {
	ormService := service.DI().OrmEngineForContext(c.Request.Context())

	id := c.Param("id")
	if len(id) <= 0 {
		response.ErrorResponseGlobal(c, "missing id", nil)
		return
	}
	ormService.GetRedis().HDel(errorlogger.GroupError, id)
	ormService.GetRedis().HDel(errorlogger.GroupError, id+":time")
	ormService.GetRedis().HDel(errorlogger.GroupError, id+":counter")

	response.SuccessResponse(c, nil)
}

func (controller *ErrorLogController) DeleteAllErrors(c *gin.Context) {
	ormService := service.DI().OrmEngineForContext(c.Request.Context())
	ormService.GetRedis().Del(errorlogger.GroupError)

	response.SuccessResponse(c, nil)
}

func (controller *ErrorLogController) Panic(_ *gin.Context) {
	panic("Forced Panic")
}
