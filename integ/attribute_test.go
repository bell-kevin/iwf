package integ

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/indeedeng/iwf/gen/iwfidl"
	"github.com/indeedeng/iwf/integ/workflow/attribute"
	"github.com/indeedeng/iwf/service"
	"github.com/stretchr/testify/assert"
)

func TestAttributeWorkflowTemporal(t *testing.T) {
	doTestAttributeWorkflow(t, service.BackendTypeTemporal)
}

func TestAttributeWorkflowCadence(t *testing.T) {
	doTestAttributeWorkflow(t, service.BackendTypeCadence)
}

func doTestAttributeWorkflow(t *testing.T, backendType service.BackendType) {
	wfHandler := attribute.NewHandler()
	closeFunc1 := startWorkflowWorker(wfHandler)
	defer closeFunc1()

	closeFunc2 := startIwfService(backendType)
	defer closeFunc2()

	// start a workflow
	apiClient := iwfidl.NewAPIClient(&iwfidl.Configuration{
		Servers: []iwfidl.ServerConfiguration{
			{
				URL: "http://localhost:" + testIwfServerPort,
			},
		},
	})
	wfId := attribute.WorkflowType + strconv.Itoa(int(time.Now().Unix()))
	reqStart := apiClient.DefaultApi.ApiV1WorkflowStartPost(context.Background())
	_, httpResp, err := reqStart.WorkflowStartRequest(iwfidl.WorkflowStartRequest{
		WorkflowId:             wfId,
		IwfWorkflowType:        attribute.WorkflowType,
		WorkflowTimeoutSeconds: 10,
		IwfWorkerUrl:           "http://localhost:" + testWorkflowServerPort,
		StartStateId:           attribute.State1,
	}).Execute()
	if err != nil {
		log.Fatalf("Fail to invoke start api %v", err)
	}
	if httpResp.StatusCode != http.StatusOK {
		log.Fatalf("Status not success" + httpResp.Status)
	}

	reqWait := apiClient.DefaultApi.ApiV1WorkflowGetWithWaitPost(context.Background())
	wfResponse, httpResp, err := reqWait.WorkflowGetRequest(iwfidl.WorkflowGetRequest{
		WorkflowId: wfId,
	}).Execute()
	if err != nil {
		log.Fatalf("Fail to invoke start api %v", err)
	}
	if httpResp.StatusCode != http.StatusOK {
		log.Fatalf("Fail to get workflow" + httpResp.Status)
	}

	reqQry := apiClient.DefaultApi.ApiV1WorkflowQueryattributesGetPost(context.Background())
	queryResult1, httpResp2, err := reqQry.WorkflowGetQueryAttributesRequest(iwfidl.WorkflowGetQueryAttributesRequest{
		WorkflowId: wfId,
		AttributeKeys: []string{
			attribute.TestQueryAttributeKey,
		},
	}).Execute()

	if err != nil || httpResp2.StatusCode != 200 {
		log.Fatalf("Fail to invoke query workflow for sigle attr %v %v", err, httpResp2)
	}

	queryResult2, httpResp2, err := reqQry.WorkflowGetQueryAttributesRequest(iwfidl.WorkflowGetQueryAttributesRequest{
		WorkflowId: wfId,
	}).Execute()

	if err != nil || httpResp2.StatusCode != 200 {
		log.Fatalf("Fail to invoke query workflow for sigle attr %v %v", err, httpResp2)
	}

	reqSearch := apiClient.DefaultApi.ApiV1WorkflowSearchattributesGetPost(context.Background())
	searchResult1, httpSearchResponse1, err := reqSearch.WorkflowGetSearchAttributesRequest(iwfidl.WorkflowGetSearchAttributesRequest{
		WorkflowId:    wfId,
		WorkflowRunId: &wfResponse.WorkflowRunId,
		AttributeKeys: []iwfidl.SearchAttributeKeyAndType{
			{
				Key:       iwfidl.PtrString(attribute.TestSearchAttributeKeywordKey),
				ValueType: iwfidl.PtrString(service.SearchAttributeValueTypeKeyword),
			},
		},
	}).Execute()

	if err != nil || httpSearchResponse1.StatusCode != 200 {
		log.Fatalf("Fail to invoke query workflow for sigle attr %v %v", err, httpSearchResponse1)
	}

	searchResult2, httpSearchResponse2, err := reqSearch.WorkflowGetSearchAttributesRequest(iwfidl.WorkflowGetSearchAttributesRequest{
		WorkflowId:    wfId,
		WorkflowRunId: &wfResponse.WorkflowRunId,
		AttributeKeys: []iwfidl.SearchAttributeKeyAndType{
			{
				Key:       iwfidl.PtrString(attribute.TestSearchAttributeIntKey),
				ValueType: iwfidl.PtrString(service.SearchAttributeValueTypeInt),
			},
		},
	}).Execute()

	if err != nil || httpSearchResponse2.StatusCode != 200 {
		log.Fatalf("Fail to invoke query workflow for sigle attr %v %v", err, httpSearchResponse2)
	}

	// assertion
	history, data := wfHandler.GetTestResult()
	assertions := assert.New(t)
	assertions.Equalf(map[string]int64{
		"S1_start":  1,
		"S1_decide": 1,
		"S2_start":  1,
		"S2_decide": 1,
	}, history, "attribute test fail, %v", history)

	if attribute.EnableTestingSearchAttribute {
		assertions.Equal(map[string]interface{}{
			"S1_decide_intSaFounds": 1,
			"S1_decide_kwSaFounds":  1,
			"S2_decide_intSaFounds": 1,
			"S2_decide_kwSaFounds":  1,
			"S2_start_intSaFounds":  1,
			"S2_start_kwSaFounds":   1,

			"S1_decide_localAttFound": true,
			"S1_decide_queryAttFound": true,
			"S2_decide_queryAttFound": true,
			"S2_start_queryAttFound":  true,
		}, data)
	} else {
		// map[S1_decide_intSaFounds:0 S1_decide_kwSaFounds:0 S1_decide_localAttFound:false
		//S1_decide_queryAttFound:true S2_decide_intSaFounds:0 S2_decide_kwSaFounds:0
		//S2_decide_queryAttFound:false S2_start_intSaFounds:0 S2_start_kwSaFounds:0 S2_start_queryAttFound:false]
		assertions.Equal(map[string]interface{}{
			"S1_decide_intSaFounds": 0,
			"S1_decide_kwSaFounds":  0,
			"S2_decide_intSaFounds": 0,
			"S2_decide_kwSaFounds":  0,
			"S2_start_intSaFounds":  0,
			"S2_start_kwSaFounds":   0,

			"S1_decide_localAttFound": true,
			"S1_decide_queryAttFound": true,
			"S2_decide_queryAttFound": true,
			"S2_start_queryAttFound":  true,
		}, data)
	}

	expected := []iwfidl.KeyValue{
		{
			Key:   iwfidl.PtrString(attribute.TestQueryAttributeKey),
			Value: &attribute.TestQueryVal2,
		},
	}
	assertions.Equal(expected, queryResult2.GetQueryAttributes())
	assertions.Equal(expected, queryResult1.GetQueryAttributes())

	expectedSearchAttributeInt := iwfidl.SearchAttribute{
		Key:          iwfidl.PtrString(attribute.TestSearchAttributeIntKey),
		ValueType:    iwfidl.PtrString(service.SearchAttributeValueTypeInt),
		IntegerValue: iwfidl.PtrInt64(attribute.TestSearchAttributeIntValue2),
	}

	expectedSearchAttributeKeyword := iwfidl.SearchAttribute{
		Key:         iwfidl.PtrString(attribute.TestSearchAttributeKeywordKey),
		ValueType:   iwfidl.PtrString(service.SearchAttributeValueTypeKeyword),
		StringValue: iwfidl.PtrString(attribute.TestSearchAttributeKeywordValue2),
	}

	assertions.Equal([]iwfidl.SearchAttribute{expectedSearchAttributeKeyword}, searchResult1.GetSearchAttributes())
	assertions.Equal([]iwfidl.SearchAttribute{expectedSearchAttributeInt}, searchResult2.GetSearchAttributes())
}
