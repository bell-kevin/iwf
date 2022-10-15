package api

import (
	"github.com/gin-gonic/gin"
)

const WorkflowStartApiPath = "/api/v1/workflow/start"
const WorkflowSignalApiPath = "/api/v1/workflow/signal"
const WorkflowGetQueryAttributesApiPath = "/api/v1/workflow/queryattributes/get"
const WorkflowGetApiPath = "/api/v1/workflow/get"
const WorkflowGetWithWaitApiPath = "/api/v1/workflow/getWithWait"
const WorkflowSearchApiPath = "/api/v1/workflow/search"
const WorkflowResetApiPath = "/api/v1/workflow/reset"

// NewService returns a new router.
func NewService(client UnifiedClient) *gin.Engine {
	router := gin.Default()

	handler := newHandler(client)

	router.GET("/", handler.index)
	router.POST(WorkflowStartApiPath, handler.apiV1WorkflowStartPost)
	router.POST(WorkflowSignalApiPath, handler.apiV1WorkflowSignalPost)
	router.POST(WorkflowGetQueryAttributesApiPath, handler.apiV1WorkflowQueryPost)
	router.POST(WorkflowGetApiPath, handler.apiV1WorkflowGetPost)
	router.POST(WorkflowGetWithWaitApiPath, handler.apiV1WorkflowGetWithWaitPost)
	router.POST(WorkflowSearchApiPath, handler.apiV1WorkflowSearchPost)
	router.POST(WorkflowResetApiPath, handler.apiV1WorkflowResetPost)

	return router
}
