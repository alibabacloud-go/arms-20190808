// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CallChainInfo struct {
	AdditionalInfo *string                           `json:"AdditionalInfo,omitempty" xml:"AdditionalInfo,omitempty"`
	AppName        *string                           `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType        *string                           `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Children       []*CallChainInfo                  `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	HaveSpan       *bool                             `json:"HaveSpan,omitempty" xml:"HaveSpan,omitempty"`
	LogMap         map[string]map[string]interface{} `json:"LogMap,omitempty" xml:"LogMap,omitempty"`
	LogTime        *int64                            `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	ParentSpanId   *string                           `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	Pid            *string                           `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId       *string                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResultCode     *string                           `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	Rpc            *string                           `json:"Rpc,omitempty" xml:"Rpc,omitempty"`
	RpcId          *string                           `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	RpcType        *int64                            `json:"RpcType,omitempty" xml:"RpcType,omitempty"`
	ServerIp       *string                           `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	Span           *int64                            `json:"Span,omitempty" xml:"Span,omitempty"`
	SpanId         *string                           `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	TagMap         map[string]*string                `json:"TagMap,omitempty" xml:"TagMap,omitempty"`
	TraceId        *string                           `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CallChainInfo) String() string {
	return tea.Prettify(s)
}

func (s CallChainInfo) GoString() string {
	return s.String()
}

func (s *CallChainInfo) SetAdditionalInfo(v string) *CallChainInfo {
	s.AdditionalInfo = &v
	return s
}

func (s *CallChainInfo) SetAppName(v string) *CallChainInfo {
	s.AppName = &v
	return s
}

func (s *CallChainInfo) SetAppType(v string) *CallChainInfo {
	s.AppType = &v
	return s
}

func (s *CallChainInfo) SetChildren(v []*CallChainInfo) *CallChainInfo {
	s.Children = v
	return s
}

func (s *CallChainInfo) SetHaveSpan(v bool) *CallChainInfo {
	s.HaveSpan = &v
	return s
}

func (s *CallChainInfo) SetLogMap(v map[string]map[string]interface{}) *CallChainInfo {
	s.LogMap = v
	return s
}

func (s *CallChainInfo) SetLogTime(v int64) *CallChainInfo {
	s.LogTime = &v
	return s
}

func (s *CallChainInfo) SetParentSpanId(v string) *CallChainInfo {
	s.ParentSpanId = &v
	return s
}

func (s *CallChainInfo) SetPid(v string) *CallChainInfo {
	s.Pid = &v
	return s
}

func (s *CallChainInfo) SetRegionId(v string) *CallChainInfo {
	s.RegionId = &v
	return s
}

func (s *CallChainInfo) SetResultCode(v string) *CallChainInfo {
	s.ResultCode = &v
	return s
}

func (s *CallChainInfo) SetRpc(v string) *CallChainInfo {
	s.Rpc = &v
	return s
}

func (s *CallChainInfo) SetRpcId(v string) *CallChainInfo {
	s.RpcId = &v
	return s
}

func (s *CallChainInfo) SetRpcType(v int64) *CallChainInfo {
	s.RpcType = &v
	return s
}

func (s *CallChainInfo) SetServerIp(v string) *CallChainInfo {
	s.ServerIp = &v
	return s
}

func (s *CallChainInfo) SetSpan(v int64) *CallChainInfo {
	s.Span = &v
	return s
}

func (s *CallChainInfo) SetSpanId(v string) *CallChainInfo {
	s.SpanId = &v
	return s
}

func (s *CallChainInfo) SetTagMap(v map[string]*string) *CallChainInfo {
	s.TagMap = v
	return s
}

func (s *CallChainInfo) SetTraceId(v string) *CallChainInfo {
	s.TraceId = &v
	return s
}

type AddAliClusterIdsToPrometheusGlobalViewRequest struct {
	ClusterIds          *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	GlobalViewClusterId *string `json:"GlobalViewClusterId,omitempty" xml:"GlobalViewClusterId,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddAliClusterIdsToPrometheusGlobalViewRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAliClusterIdsToPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) SetClusterIds(v string) *AddAliClusterIdsToPrometheusGlobalViewRequest {
	s.ClusterIds = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) SetGlobalViewClusterId(v string) *AddAliClusterIdsToPrometheusGlobalViewRequest {
	s.GlobalViewClusterId = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) SetGroupName(v string) *AddAliClusterIdsToPrometheusGlobalViewRequest {
	s.GroupName = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewRequest) SetRegionId(v string) *AddAliClusterIdsToPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

type AddAliClusterIdsToPrometheusGlobalViewResponseBody struct {
	Data      *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAliClusterIdsToPrometheusGlobalViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAliClusterIdsToPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBody) SetData(v *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) *AddAliClusterIdsToPrometheusGlobalViewResponseBody {
	s.Data = v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBody) SetRequestId(v string) *AddAliClusterIdsToPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

type AddAliClusterIdsToPrometheusGlobalViewResponseBodyData struct {
	Info    *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Msg     *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) SetInfo(v string) *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData {
	s.Info = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) SetMsg(v string) *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData {
	s.Msg = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) SetSuccess(v bool) *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData {
	s.Success = &v
	return s
}

type AddAliClusterIdsToPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddAliClusterIdsToPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAliClusterIdsToPrometheusGlobalViewResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAliClusterIdsToPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *AddAliClusterIdsToPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponse) SetStatusCode(v int32) *AddAliClusterIdsToPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponse) SetBody(v *AddAliClusterIdsToPrometheusGlobalViewResponseBody) *AddAliClusterIdsToPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

type AddGrafanaRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Integration *string `json:"Integration,omitempty" xml:"Integration,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddGrafanaRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGrafanaRequest) GoString() string {
	return s.String()
}

func (s *AddGrafanaRequest) SetClusterId(v string) *AddGrafanaRequest {
	s.ClusterId = &v
	return s
}

func (s *AddGrafanaRequest) SetIntegration(v string) *AddGrafanaRequest {
	s.Integration = &v
	return s
}

func (s *AddGrafanaRequest) SetRegionId(v string) *AddGrafanaRequest {
	s.RegionId = &v
	return s
}

type AddGrafanaResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddGrafanaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGrafanaResponseBody) GoString() string {
	return s.String()
}

func (s *AddGrafanaResponseBody) SetData(v string) *AddGrafanaResponseBody {
	s.Data = &v
	return s
}

func (s *AddGrafanaResponseBody) SetRequestId(v string) *AddGrafanaResponseBody {
	s.RequestId = &v
	return s
}

type AddGrafanaResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddGrafanaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddGrafanaResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGrafanaResponse) GoString() string {
	return s.String()
}

func (s *AddGrafanaResponse) SetHeaders(v map[string]*string) *AddGrafanaResponse {
	s.Headers = v
	return s
}

func (s *AddGrafanaResponse) SetStatusCode(v int32) *AddGrafanaResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGrafanaResponse) SetBody(v *AddGrafanaResponseBody) *AddGrafanaResponse {
	s.Body = v
	return s
}

type AddIntegrationRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Integration *string `json:"Integration,omitempty" xml:"Integration,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddIntegrationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIntegrationRequest) GoString() string {
	return s.String()
}

func (s *AddIntegrationRequest) SetClusterId(v string) *AddIntegrationRequest {
	s.ClusterId = &v
	return s
}

func (s *AddIntegrationRequest) SetIntegration(v string) *AddIntegrationRequest {
	s.Integration = &v
	return s
}

func (s *AddIntegrationRequest) SetRegionId(v string) *AddIntegrationRequest {
	s.RegionId = &v
	return s
}

type AddIntegrationResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIntegrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *AddIntegrationResponseBody) SetData(v string) *AddIntegrationResponseBody {
	s.Data = &v
	return s
}

func (s *AddIntegrationResponseBody) SetRequestId(v string) *AddIntegrationResponseBody {
	s.RequestId = &v
	return s
}

type AddIntegrationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddIntegrationResponse) String() string {
	return tea.Prettify(s)
}

func (s AddIntegrationResponse) GoString() string {
	return s.String()
}

func (s *AddIntegrationResponse) SetHeaders(v map[string]*string) *AddIntegrationResponse {
	s.Headers = v
	return s
}

func (s *AddIntegrationResponse) SetStatusCode(v int32) *AddIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIntegrationResponse) SetBody(v *AddIntegrationResponseBody) *AddIntegrationResponse {
	s.Body = v
	return s
}

type AddPrometheusGlobalViewRequest struct {
	Clusters  *string `json:"Clusters,omitempty" xml:"Clusters,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddPrometheusGlobalViewRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewRequest) SetClusters(v string) *AddPrometheusGlobalViewRequest {
	s.Clusters = &v
	return s
}

func (s *AddPrometheusGlobalViewRequest) SetGroupName(v string) *AddPrometheusGlobalViewRequest {
	s.GroupName = &v
	return s
}

func (s *AddPrometheusGlobalViewRequest) SetRegionId(v string) *AddPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

type AddPrometheusGlobalViewResponseBody struct {
	Data      *AddPrometheusGlobalViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPrometheusGlobalViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewResponseBody) SetData(v *AddPrometheusGlobalViewResponseBodyData) *AddPrometheusGlobalViewResponseBody {
	s.Data = v
	return s
}

func (s *AddPrometheusGlobalViewResponseBody) SetRequestId(v string) *AddPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

type AddPrometheusGlobalViewResponseBodyData struct {
	Info    *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Msg     *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddPrometheusGlobalViewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddPrometheusGlobalViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewResponseBodyData) SetInfo(v string) *AddPrometheusGlobalViewResponseBodyData {
	s.Info = &v
	return s
}

func (s *AddPrometheusGlobalViewResponseBodyData) SetMsg(v string) *AddPrometheusGlobalViewResponseBodyData {
	s.Msg = &v
	return s
}

func (s *AddPrometheusGlobalViewResponseBodyData) SetSuccess(v bool) *AddPrometheusGlobalViewResponseBodyData {
	s.Success = &v
	return s
}

type AddPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddPrometheusGlobalViewResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *AddPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *AddPrometheusGlobalViewResponse) SetStatusCode(v int32) *AddPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPrometheusGlobalViewResponse) SetBody(v *AddPrometheusGlobalViewResponseBody) *AddPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

type AddPrometheusGlobalViewByAliClusterIdsRequest struct {
	ClusterIds  *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddPrometheusGlobalViewByAliClusterIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPrometheusGlobalViewByAliClusterIdsRequest) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) SetClusterIds(v string) *AddPrometheusGlobalViewByAliClusterIdsRequest {
	s.ClusterIds = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) SetGroupName(v string) *AddPrometheusGlobalViewByAliClusterIdsRequest {
	s.GroupName = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) SetProductCode(v string) *AddPrometheusGlobalViewByAliClusterIdsRequest {
	s.ProductCode = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsRequest) SetRegionId(v string) *AddPrometheusGlobalViewByAliClusterIdsRequest {
	s.RegionId = &v
	return s
}

type AddPrometheusGlobalViewByAliClusterIdsResponseBody struct {
	Data      *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPrometheusGlobalViewByAliClusterIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddPrometheusGlobalViewByAliClusterIdsResponseBody) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBody) SetData(v *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) *AddPrometheusGlobalViewByAliClusterIdsResponseBody {
	s.Data = v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBody) SetRequestId(v string) *AddPrometheusGlobalViewByAliClusterIdsResponseBody {
	s.RequestId = &v
	return s
}

type AddPrometheusGlobalViewByAliClusterIdsResponseBodyData struct {
	Info    *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Msg     *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) SetInfo(v string) *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData {
	s.Info = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) SetMsg(v string) *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData {
	s.Msg = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) SetSuccess(v bool) *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData {
	s.Success = &v
	return s
}

type AddPrometheusGlobalViewByAliClusterIdsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddPrometheusGlobalViewByAliClusterIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddPrometheusGlobalViewByAliClusterIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPrometheusGlobalViewByAliClusterIdsResponse) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponse) SetHeaders(v map[string]*string) *AddPrometheusGlobalViewByAliClusterIdsResponse {
	s.Headers = v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponse) SetStatusCode(v int32) *AddPrometheusGlobalViewByAliClusterIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponse) SetBody(v *AddPrometheusGlobalViewByAliClusterIdsResponseBody) *AddPrometheusGlobalViewByAliClusterIdsResponse {
	s.Body = v
	return s
}

type AddPrometheusInstanceRequest struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddPrometheusInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPrometheusInstanceRequest) GoString() string {
	return s.String()
}

func (s *AddPrometheusInstanceRequest) SetName(v string) *AddPrometheusInstanceRequest {
	s.Name = &v
	return s
}

func (s *AddPrometheusInstanceRequest) SetRegionId(v string) *AddPrometheusInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *AddPrometheusInstanceRequest) SetType(v string) *AddPrometheusInstanceRequest {
	s.Type = &v
	return s
}

type AddPrometheusInstanceResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPrometheusInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddPrometheusInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AddPrometheusInstanceResponseBody) SetData(v string) *AddPrometheusInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *AddPrometheusInstanceResponseBody) SetRequestId(v string) *AddPrometheusInstanceResponseBody {
	s.RequestId = &v
	return s
}

type AddPrometheusInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddPrometheusInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddPrometheusInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPrometheusInstanceResponse) GoString() string {
	return s.String()
}

func (s *AddPrometheusInstanceResponse) SetHeaders(v map[string]*string) *AddPrometheusInstanceResponse {
	s.Headers = v
	return s
}

func (s *AddPrometheusInstanceResponse) SetStatusCode(v int32) *AddPrometheusInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPrometheusInstanceResponse) SetBody(v *AddPrometheusInstanceResponseBody) *AddPrometheusInstanceResponse {
	s.Body = v
	return s
}

type AddRecordingRuleRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleYaml  *string `json:"RuleYaml,omitempty" xml:"RuleYaml,omitempty"`
}

func (s AddRecordingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRecordingRuleRequest) GoString() string {
	return s.String()
}

func (s *AddRecordingRuleRequest) SetClusterId(v string) *AddRecordingRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *AddRecordingRuleRequest) SetRegionId(v string) *AddRecordingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *AddRecordingRuleRequest) SetRuleYaml(v string) *AddRecordingRuleRequest {
	s.RuleYaml = &v
	return s
}

type AddRecordingRuleResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddRecordingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRecordingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddRecordingRuleResponseBody) SetData(v string) *AddRecordingRuleResponseBody {
	s.Data = &v
	return s
}

func (s *AddRecordingRuleResponseBody) SetRequestId(v string) *AddRecordingRuleResponseBody {
	s.RequestId = &v
	return s
}

type AddRecordingRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddRecordingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRecordingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRecordingRuleResponse) GoString() string {
	return s.String()
}

func (s *AddRecordingRuleResponse) SetHeaders(v map[string]*string) *AddRecordingRuleResponse {
	s.Headers = v
	return s
}

func (s *AddRecordingRuleResponse) SetStatusCode(v int32) *AddRecordingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRecordingRuleResponse) SetBody(v *AddRecordingRuleResponseBody) *AddRecordingRuleResponse {
	s.Body = v
	return s
}

type AppendInstancesToPrometheusGlobalViewRequest struct {
	Clusters            *string `json:"Clusters,omitempty" xml:"Clusters,omitempty"`
	GlobalViewClusterId *string `json:"GlobalViewClusterId,omitempty" xml:"GlobalViewClusterId,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AppendInstancesToPrometheusGlobalViewRequest) String() string {
	return tea.Prettify(s)
}

func (s AppendInstancesToPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) SetClusters(v string) *AppendInstancesToPrometheusGlobalViewRequest {
	s.Clusters = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) SetGlobalViewClusterId(v string) *AppendInstancesToPrometheusGlobalViewRequest {
	s.GlobalViewClusterId = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) SetGroupName(v string) *AppendInstancesToPrometheusGlobalViewRequest {
	s.GroupName = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewRequest) SetRegionId(v string) *AppendInstancesToPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

type AppendInstancesToPrometheusGlobalViewResponseBody struct {
	Data      *AppendInstancesToPrometheusGlobalViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppendInstancesToPrometheusGlobalViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppendInstancesToPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBody) SetData(v *AppendInstancesToPrometheusGlobalViewResponseBodyData) *AppendInstancesToPrometheusGlobalViewResponseBody {
	s.Data = v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBody) SetRequestId(v string) *AppendInstancesToPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

type AppendInstancesToPrometheusGlobalViewResponseBodyData struct {
	Info    *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Msg     *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AppendInstancesToPrometheusGlobalViewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AppendInstancesToPrometheusGlobalViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBodyData) SetInfo(v string) *AppendInstancesToPrometheusGlobalViewResponseBodyData {
	s.Info = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBodyData) SetMsg(v string) *AppendInstancesToPrometheusGlobalViewResponseBodyData {
	s.Msg = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponseBodyData) SetSuccess(v bool) *AppendInstancesToPrometheusGlobalViewResponseBodyData {
	s.Success = &v
	return s
}

type AppendInstancesToPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AppendInstancesToPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppendInstancesToPrometheusGlobalViewResponse) String() string {
	return tea.Prettify(s)
}

func (s AppendInstancesToPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *AppendInstancesToPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *AppendInstancesToPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponse) SetStatusCode(v int32) *AppendInstancesToPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *AppendInstancesToPrometheusGlobalViewResponse) SetBody(v *AppendInstancesToPrometheusGlobalViewResponseBody) *AppendInstancesToPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

type ApplyScenarioRequest struct {
	AppId        *string                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Config       map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	Name         *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId     *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scenario     *string                `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	Sign         *string                `json:"Sign,omitempty" xml:"Sign,omitempty"`
	SnDump       *bool                  `json:"SnDump,omitempty" xml:"SnDump,omitempty"`
	SnForce      *bool                  `json:"SnForce,omitempty" xml:"SnForce,omitempty"`
	SnStat       *bool                  `json:"SnStat,omitempty" xml:"SnStat,omitempty"`
	SnTransfer   *bool                  `json:"SnTransfer,omitempty" xml:"SnTransfer,omitempty"`
	UpdateOption *bool                  `json:"UpdateOption,omitempty" xml:"UpdateOption,omitempty"`
}

func (s ApplyScenarioRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyScenarioRequest) GoString() string {
	return s.String()
}

func (s *ApplyScenarioRequest) SetAppId(v string) *ApplyScenarioRequest {
	s.AppId = &v
	return s
}

func (s *ApplyScenarioRequest) SetConfig(v map[string]interface{}) *ApplyScenarioRequest {
	s.Config = v
	return s
}

func (s *ApplyScenarioRequest) SetName(v string) *ApplyScenarioRequest {
	s.Name = &v
	return s
}

func (s *ApplyScenarioRequest) SetRegionId(v string) *ApplyScenarioRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyScenarioRequest) SetScenario(v string) *ApplyScenarioRequest {
	s.Scenario = &v
	return s
}

func (s *ApplyScenarioRequest) SetSign(v string) *ApplyScenarioRequest {
	s.Sign = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnDump(v bool) *ApplyScenarioRequest {
	s.SnDump = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnForce(v bool) *ApplyScenarioRequest {
	s.SnForce = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnStat(v bool) *ApplyScenarioRequest {
	s.SnStat = &v
	return s
}

func (s *ApplyScenarioRequest) SetSnTransfer(v bool) *ApplyScenarioRequest {
	s.SnTransfer = &v
	return s
}

func (s *ApplyScenarioRequest) SetUpdateOption(v bool) *ApplyScenarioRequest {
	s.UpdateOption = &v
	return s
}

type ApplyScenarioShrinkRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scenario     *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	Sign         *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	SnDump       *bool   `json:"SnDump,omitempty" xml:"SnDump,omitempty"`
	SnForce      *bool   `json:"SnForce,omitempty" xml:"SnForce,omitempty"`
	SnStat       *bool   `json:"SnStat,omitempty" xml:"SnStat,omitempty"`
	SnTransfer   *bool   `json:"SnTransfer,omitempty" xml:"SnTransfer,omitempty"`
	UpdateOption *bool   `json:"UpdateOption,omitempty" xml:"UpdateOption,omitempty"`
}

func (s ApplyScenarioShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyScenarioShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyScenarioShrinkRequest) SetAppId(v string) *ApplyScenarioShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetConfigShrink(v string) *ApplyScenarioShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetName(v string) *ApplyScenarioShrinkRequest {
	s.Name = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetRegionId(v string) *ApplyScenarioShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetScenario(v string) *ApplyScenarioShrinkRequest {
	s.Scenario = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSign(v string) *ApplyScenarioShrinkRequest {
	s.Sign = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnDump(v bool) *ApplyScenarioShrinkRequest {
	s.SnDump = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnForce(v bool) *ApplyScenarioShrinkRequest {
	s.SnForce = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnStat(v bool) *ApplyScenarioShrinkRequest {
	s.SnStat = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetSnTransfer(v bool) *ApplyScenarioShrinkRequest {
	s.SnTransfer = &v
	return s
}

func (s *ApplyScenarioShrinkRequest) SetUpdateOption(v bool) *ApplyScenarioShrinkRequest {
	s.UpdateOption = &v
	return s
}

type ApplyScenarioResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ApplyScenarioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyScenarioResponseBody) SetRequestId(v string) *ApplyScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyScenarioResponseBody) SetResult(v string) *ApplyScenarioResponseBody {
	s.Result = &v
	return s
}

type ApplyScenarioResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApplyScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyScenarioResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyScenarioResponse) GoString() string {
	return s.String()
}

func (s *ApplyScenarioResponse) SetHeaders(v map[string]*string) *ApplyScenarioResponse {
	s.Headers = v
	return s
}

func (s *ApplyScenarioResponse) SetStatusCode(v int32) *ApplyScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyScenarioResponse) SetBody(v *ApplyScenarioResponseBody) *ApplyScenarioResponse {
	s.Body = v
	return s
}

type CheckServiceStatusRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SvcCode  *string `json:"SvcCode,omitempty" xml:"SvcCode,omitempty"`
}

func (s CheckServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceStatusRequest) SetRegionId(v string) *CheckServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *CheckServiceStatusRequest) SetSvcCode(v string) *CheckServiceStatusRequest {
	s.SvcCode = &v
	return s
}

type CheckServiceStatusResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceStatusResponseBody) SetData(v string) *CheckServiceStatusResponseBody {
	s.Data = &v
	return s
}

func (s *CheckServiceStatusResponseBody) SetRequestId(v string) *CheckServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

type CheckServiceStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceStatusResponse) SetHeaders(v map[string]*string) *CheckServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceStatusResponse) SetStatusCode(v int32) *CheckServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceStatusResponse) SetBody(v *CheckServiceStatusResponseBody) *CheckServiceStatusResponse {
	s.Body = v
	return s
}

type ConfigAppRequest struct {
	AppIds   *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	Enable   *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigAppRequest) GoString() string {
	return s.String()
}

func (s *ConfigAppRequest) SetAppIds(v string) *ConfigAppRequest {
	s.AppIds = &v
	return s
}

func (s *ConfigAppRequest) SetEnable(v string) *ConfigAppRequest {
	s.Enable = &v
	return s
}

func (s *ConfigAppRequest) SetRegionId(v string) *ConfigAppRequest {
	s.RegionId = &v
	return s
}

type ConfigAppResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigAppResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigAppResponseBody) SetData(v string) *ConfigAppResponseBody {
	s.Data = &v
	return s
}

func (s *ConfigAppResponseBody) SetRequestId(v string) *ConfigAppResponseBody {
	s.RequestId = &v
	return s
}

type ConfigAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfigAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigAppResponse) GoString() string {
	return s.String()
}

func (s *ConfigAppResponse) SetHeaders(v map[string]*string) *ConfigAppResponse {
	s.Headers = v
	return s
}

func (s *ConfigAppResponse) SetStatusCode(v int32) *ConfigAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigAppResponse) SetBody(v *ConfigAppResponseBody) *ConfigAppResponse {
	s.Body = v
	return s
}

type CreateAlertContactRequest struct {
	ContactName         *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	DingRobotWebhookUrl *string `json:"DingRobotWebhookUrl,omitempty" xml:"DingRobotWebhookUrl,omitempty"`
	Email               *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PhoneNum            *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SystemNoc           *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
}

func (s CreateAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertContactRequest) SetContactName(v string) *CreateAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *CreateAlertContactRequest) SetDingRobotWebhookUrl(v string) *CreateAlertContactRequest {
	s.DingRobotWebhookUrl = &v
	return s
}

func (s *CreateAlertContactRequest) SetEmail(v string) *CreateAlertContactRequest {
	s.Email = &v
	return s
}

func (s *CreateAlertContactRequest) SetPhoneNum(v string) *CreateAlertContactRequest {
	s.PhoneNum = &v
	return s
}

func (s *CreateAlertContactRequest) SetRegionId(v string) *CreateAlertContactRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAlertContactRequest) SetSystemNoc(v bool) *CreateAlertContactRequest {
	s.SystemNoc = &v
	return s
}

type CreateAlertContactResponseBody struct {
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlertContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertContactResponseBody) SetContactId(v string) *CreateAlertContactResponseBody {
	s.ContactId = &v
	return s
}

func (s *CreateAlertContactResponseBody) SetRequestId(v string) *CreateAlertContactResponseBody {
	s.RequestId = &v
	return s
}

type CreateAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlertContactResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertContactResponse) SetHeaders(v map[string]*string) *CreateAlertContactResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertContactResponse) SetStatusCode(v int32) *CreateAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertContactResponse) SetBody(v *CreateAlertContactResponseBody) *CreateAlertContactResponse {
	s.Body = v
	return s
}

type CreateAlertContactGroupRequest struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	ContactIds       *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAlertContactGroupRequest) SetContactGroupName(v string) *CreateAlertContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *CreateAlertContactGroupRequest) SetContactIds(v string) *CreateAlertContactGroupRequest {
	s.ContactIds = &v
	return s
}

func (s *CreateAlertContactGroupRequest) SetRegionId(v string) *CreateAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

type CreateAlertContactGroupResponseBody struct {
	ContactGroupId *string `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlertContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertContactGroupResponseBody) SetContactGroupId(v string) *CreateAlertContactGroupResponseBody {
	s.ContactGroupId = &v
	return s
}

func (s *CreateAlertContactGroupResponseBody) SetRequestId(v string) *CreateAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlertContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertContactGroupResponse) SetHeaders(v map[string]*string) *CreateAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertContactGroupResponse) SetStatusCode(v int32) *CreateAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertContactGroupResponse) SetBody(v *CreateAlertContactGroupResponseBody) *CreateAlertContactGroupResponse {
	s.Body = v
	return s
}

type CreateDispatchRuleRequest struct {
	DispatchRule *string `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDispatchRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDispatchRuleRequest) SetDispatchRule(v string) *CreateDispatchRuleRequest {
	s.DispatchRule = &v
	return s
}

func (s *CreateDispatchRuleRequest) SetRegionId(v string) *CreateDispatchRuleRequest {
	s.RegionId = &v
	return s
}

type CreateDispatchRuleResponseBody struct {
	DispatchRuleId *int64  `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDispatchRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDispatchRuleResponseBody) SetDispatchRuleId(v int64) *CreateDispatchRuleResponseBody {
	s.DispatchRuleId = &v
	return s
}

func (s *CreateDispatchRuleResponseBody) SetRequestId(v string) *CreateDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateDispatchRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDispatchRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDispatchRuleResponse) SetHeaders(v map[string]*string) *CreateDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDispatchRuleResponse) SetStatusCode(v int32) *CreateDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDispatchRuleResponse) SetBody(v *CreateDispatchRuleResponseBody) *CreateDispatchRuleResponse {
	s.Body = v
	return s
}

type CreateIntegrationRequest struct {
	AutoRecover            *bool   `json:"AutoRecover,omitempty" xml:"AutoRecover,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IntegrationName        *string `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	IntegrationProductType *string `json:"IntegrationProductType,omitempty" xml:"IntegrationProductType,omitempty"`
	RecoverTime            *int64  `json:"RecoverTime,omitempty" xml:"RecoverTime,omitempty"`
}

func (s CreateIntegrationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegrationRequest) GoString() string {
	return s.String()
}

func (s *CreateIntegrationRequest) SetAutoRecover(v bool) *CreateIntegrationRequest {
	s.AutoRecover = &v
	return s
}

func (s *CreateIntegrationRequest) SetDescription(v string) *CreateIntegrationRequest {
	s.Description = &v
	return s
}

func (s *CreateIntegrationRequest) SetIntegrationName(v string) *CreateIntegrationRequest {
	s.IntegrationName = &v
	return s
}

func (s *CreateIntegrationRequest) SetIntegrationProductType(v string) *CreateIntegrationRequest {
	s.IntegrationProductType = &v
	return s
}

func (s *CreateIntegrationRequest) SetRecoverTime(v int64) *CreateIntegrationRequest {
	s.RecoverTime = &v
	return s
}

type CreateIntegrationResponseBody struct {
	Integration *CreateIntegrationResponseBodyIntegration `json:"Integration,omitempty" xml:"Integration,omitempty" type:"Struct"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIntegrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntegrationResponseBody) SetIntegration(v *CreateIntegrationResponseBodyIntegration) *CreateIntegrationResponseBody {
	s.Integration = v
	return s
}

func (s *CreateIntegrationResponseBody) SetRequestId(v string) *CreateIntegrationResponseBody {
	s.RequestId = &v
	return s
}

type CreateIntegrationResponseBodyIntegration struct {
	AutoRecover            *bool   `json:"AutoRecover,omitempty" xml:"AutoRecover,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IntegrationId          *int64  `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
	IntegrationName        *string `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	IntegrationProductType *string `json:"IntegrationProductType,omitempty" xml:"IntegrationProductType,omitempty"`
	RecoverTime            *int64  `json:"RecoverTime,omitempty" xml:"RecoverTime,omitempty"`
}

func (s CreateIntegrationResponseBodyIntegration) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegrationResponseBodyIntegration) GoString() string {
	return s.String()
}

func (s *CreateIntegrationResponseBodyIntegration) SetAutoRecover(v bool) *CreateIntegrationResponseBodyIntegration {
	s.AutoRecover = &v
	return s
}

func (s *CreateIntegrationResponseBodyIntegration) SetDescription(v string) *CreateIntegrationResponseBodyIntegration {
	s.Description = &v
	return s
}

func (s *CreateIntegrationResponseBodyIntegration) SetIntegrationId(v int64) *CreateIntegrationResponseBodyIntegration {
	s.IntegrationId = &v
	return s
}

func (s *CreateIntegrationResponseBodyIntegration) SetIntegrationName(v string) *CreateIntegrationResponseBodyIntegration {
	s.IntegrationName = &v
	return s
}

func (s *CreateIntegrationResponseBodyIntegration) SetIntegrationProductType(v string) *CreateIntegrationResponseBodyIntegration {
	s.IntegrationProductType = &v
	return s
}

func (s *CreateIntegrationResponseBodyIntegration) SetRecoverTime(v int64) *CreateIntegrationResponseBodyIntegration {
	s.RecoverTime = &v
	return s
}

type CreateIntegrationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIntegrationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIntegrationResponse) GoString() string {
	return s.String()
}

func (s *CreateIntegrationResponse) SetHeaders(v map[string]*string) *CreateIntegrationResponse {
	s.Headers = v
	return s
}

func (s *CreateIntegrationResponse) SetStatusCode(v int32) *CreateIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIntegrationResponse) SetBody(v *CreateIntegrationResponseBody) *CreateIntegrationResponse {
	s.Body = v
	return s
}

type CreateOrUpdateAlertRuleRequest struct {
	AlertCheckType        *string `json:"AlertCheckType,omitempty" xml:"AlertCheckType,omitempty"`
	AlertGroup            *int64  `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	AlertId               *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName             *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertRuleContent      *string `json:"AlertRuleContent,omitempty" xml:"AlertRuleContent,omitempty"`
	AlertStatus           *string `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	AlertType             *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	Annotations           *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	AutoAddNewApplication *bool   `json:"AutoAddNewApplication,omitempty" xml:"AutoAddNewApplication,omitempty"`
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Duration              *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Filters               *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	Labels                *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Level                 *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message               *string `json:"Message,omitempty" xml:"Message,omitempty"`
	MetricsKey            *string `json:"MetricsKey,omitempty" xml:"MetricsKey,omitempty"`
	MetricsType           *string `json:"MetricsType,omitempty" xml:"MetricsType,omitempty"`
	NotifyStrategy        *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	Pids                  *string `json:"Pids,omitempty" xml:"Pids,omitempty"`
	PromQL                *string `json:"PromQL,omitempty" xml:"PromQL,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateOrUpdateAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertCheckType(v string) *CreateOrUpdateAlertRuleRequest {
	s.AlertCheckType = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertGroup(v int64) *CreateOrUpdateAlertRuleRequest {
	s.AlertGroup = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertId(v int64) *CreateOrUpdateAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertName(v string) *CreateOrUpdateAlertRuleRequest {
	s.AlertName = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertRuleContent(v string) *CreateOrUpdateAlertRuleRequest {
	s.AlertRuleContent = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertStatus(v string) *CreateOrUpdateAlertRuleRequest {
	s.AlertStatus = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertType(v string) *CreateOrUpdateAlertRuleRequest {
	s.AlertType = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAnnotations(v string) *CreateOrUpdateAlertRuleRequest {
	s.Annotations = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAutoAddNewApplication(v bool) *CreateOrUpdateAlertRuleRequest {
	s.AutoAddNewApplication = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetClusterId(v string) *CreateOrUpdateAlertRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetDuration(v int64) *CreateOrUpdateAlertRuleRequest {
	s.Duration = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetFilters(v string) *CreateOrUpdateAlertRuleRequest {
	s.Filters = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetLabels(v string) *CreateOrUpdateAlertRuleRequest {
	s.Labels = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetLevel(v string) *CreateOrUpdateAlertRuleRequest {
	s.Level = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetMessage(v string) *CreateOrUpdateAlertRuleRequest {
	s.Message = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetMetricsKey(v string) *CreateOrUpdateAlertRuleRequest {
	s.MetricsKey = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetMetricsType(v string) *CreateOrUpdateAlertRuleRequest {
	s.MetricsType = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetNotifyStrategy(v string) *CreateOrUpdateAlertRuleRequest {
	s.NotifyStrategy = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetPids(v string) *CreateOrUpdateAlertRuleRequest {
	s.Pids = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetPromQL(v string) *CreateOrUpdateAlertRuleRequest {
	s.PromQL = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetRegionId(v string) *CreateOrUpdateAlertRuleRequest {
	s.RegionId = &v
	return s
}

type CreateOrUpdateAlertRuleResponseBody struct {
	AlertRule *CreateOrUpdateAlertRuleResponseBodyAlertRule `json:"AlertRule,omitempty" xml:"AlertRule,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBody) SetAlertRule(v *CreateOrUpdateAlertRuleResponseBodyAlertRule) *CreateOrUpdateAlertRuleResponseBody {
	s.AlertRule = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBody) SetRequestId(v string) *CreateOrUpdateAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrUpdateAlertRuleResponseBodyAlertRule struct {
	AlertCheckType        *string                                                       `json:"AlertCheckType,omitempty" xml:"AlertCheckType,omitempty"`
	AlertGroup            *int64                                                        `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	AlertId               *float32                                                      `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName             *string                                                       `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertRuleContent      *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent `json:"AlertRuleContent,omitempty" xml:"AlertRuleContent,omitempty" type:"Struct"`
	AlertStatus           *string                                                       `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	AlertType             *string                                                       `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	Annotations           []*CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations    `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	AutoAddNewApplication *bool                                                         `json:"AutoAddNewApplication,omitempty" xml:"AutoAddNewApplication,omitempty"`
	ClusterId             *string                                                       `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreatedTime           *int64                                                        `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Duration              *string                                                       `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Extend                *string                                                       `json:"Extend,omitempty" xml:"Extend,omitempty"`
	Filters               *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters          `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	Labels                []*CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels         `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Level                 *string                                                       `json:"Level,omitempty" xml:"Level,omitempty"`
	Message               *string                                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	MetricsType           *string                                                       `json:"MetricsType,omitempty" xml:"MetricsType,omitempty"`
	NotifyStrategy        *string                                                       `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	Pids                  []*string                                                     `json:"Pids,omitempty" xml:"Pids,omitempty" type:"Repeated"`
	PromQL                *string                                                       `json:"PromQL,omitempty" xml:"PromQL,omitempty"`
	RegionId              *string                                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UpdatedTime           *int64                                                        `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	UserId                *string                                                       `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRule) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRule) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertCheckType(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertCheckType = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertGroup(v int64) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertGroup = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertId(v float32) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertName(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertName = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertRuleContent(v *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertRuleContent = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertStatus(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertStatus = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAlertType(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AlertType = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAnnotations(v []*CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Annotations = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetAutoAddNewApplication(v bool) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.AutoAddNewApplication = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetClusterId(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.ClusterId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetCreatedTime(v int64) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.CreatedTime = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetDuration(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Duration = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetExtend(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Extend = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetFilters(v *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Filters = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetLabels(v []*CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Labels = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetLevel(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Level = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetMessage(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Message = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetMetricsType(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.MetricsType = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetNotifyStrategy(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.NotifyStrategy = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetPids(v []*string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.Pids = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetPromQL(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.PromQL = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetRegionId(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.RegionId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetUpdatedTime(v int64) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.UpdatedTime = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRule) SetUserId(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRule {
	s.UserId = &v
	return s
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent struct {
	AlertRuleItems []*CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems `json:"AlertRuleItems,omitempty" xml:"AlertRuleItems,omitempty" type:"Repeated"`
	Condition      *string                                                                       `json:"Condition,omitempty" xml:"Condition,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) SetAlertRuleItems(v []*CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent {
	s.AlertRuleItems = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent) SetCondition(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContent {
	s.Condition = &v
	return s
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems struct {
	Aggregate *string  `json:"Aggregate,omitempty" xml:"Aggregate,omitempty"`
	MetricKey *string  `json:"MetricKey,omitempty" xml:"MetricKey,omitempty"`
	N         *float32 `json:"N,omitempty" xml:"N,omitempty"`
	Operator  *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value     *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) SetAggregate(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems {
	s.Aggregate = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) SetMetricKey(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems {
	s.MetricKey = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) SetN(v float32) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems {
	s.N = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) SetOperator(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems {
	s.Operator = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems) SetValue(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAlertRuleContentAlertRuleItems {
	s.Value = &v
	return s
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) SetName(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations) SetValue(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleAnnotations {
	s.Value = &v
	return s
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters struct {
	CustomSLSFilters           []*CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters `json:"CustomSLSFilters,omitempty" xml:"CustomSLSFilters,omitempty" type:"Repeated"`
	CustomSLSGroupByDimensions []*string                                                              `json:"CustomSLSGroupByDimensions,omitempty" xml:"CustomSLSGroupByDimensions,omitempty" type:"Repeated"`
	CustomSLSWheres            []*string                                                              `json:"CustomSLSWheres,omitempty" xml:"CustomSLSWheres,omitempty" type:"Repeated"`
	DimFilters                 []*CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters       `json:"DimFilters,omitempty" xml:"DimFilters,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) SetCustomSLSFilters(v []*CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters {
	s.CustomSLSFilters = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) SetCustomSLSGroupByDimensions(v []*string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters {
	s.CustomSLSGroupByDimensions = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) SetCustomSLSWheres(v []*string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters {
	s.CustomSLSWheres = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters) SetDimFilters(v []*CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFilters {
	s.DimFilters = v
	return s
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Opt   *string `json:"Opt,omitempty" xml:"Opt,omitempty"`
	Show  *bool   `json:"Show,omitempty" xml:"Show,omitempty"`
	T     *string `json:"T,omitempty" xml:"T,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) SetKey(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters {
	s.Key = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) SetOpt(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters {
	s.Opt = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) SetShow(v bool) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters {
	s.Show = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) SetT(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters {
	s.T = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters) SetValue(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersCustomSLSFilters {
	s.Value = &v
	return s
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters struct {
	FilterKey    *string   `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterOpt    *string   `json:"FilterOpt,omitempty" xml:"FilterOpt,omitempty"`
	FilterValues []*string `json:"FilterValues,omitempty" xml:"FilterValues,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) SetFilterKey(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters {
	s.FilterKey = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) SetFilterOpt(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters {
	s.FilterOpt = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters) SetFilterValues(v []*string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleFiltersDimFilters {
	s.FilterValues = v
	return s
}

type CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) SetName(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels) SetValue(v string) *CreateOrUpdateAlertRuleResponseBodyAlertRuleLabels {
	s.Value = &v
	return s
}

type CreateOrUpdateAlertRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrUpdateAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrUpdateAlertRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleResponse) SetHeaders(v map[string]*string) *CreateOrUpdateAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateAlertRuleResponse) SetStatusCode(v int32) *CreateOrUpdateAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateAlertRuleResponse) SetBody(v *CreateOrUpdateAlertRuleResponseBody) *CreateOrUpdateAlertRuleResponse {
	s.Body = v
	return s
}

type CreateOrUpdateContactRequest struct {
	ContactId         *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName       *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Phone             *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	ReissueSendNotice *int64  `json:"ReissueSendNotice,omitempty" xml:"ReissueSendNotice,omitempty"`
}

func (s CreateOrUpdateContactRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateContactRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactRequest) SetContactId(v int64) *CreateOrUpdateContactRequest {
	s.ContactId = &v
	return s
}

func (s *CreateOrUpdateContactRequest) SetContactName(v string) *CreateOrUpdateContactRequest {
	s.ContactName = &v
	return s
}

func (s *CreateOrUpdateContactRequest) SetEmail(v string) *CreateOrUpdateContactRequest {
	s.Email = &v
	return s
}

func (s *CreateOrUpdateContactRequest) SetPhone(v string) *CreateOrUpdateContactRequest {
	s.Phone = &v
	return s
}

func (s *CreateOrUpdateContactRequest) SetReissueSendNotice(v int64) *CreateOrUpdateContactRequest {
	s.ReissueSendNotice = &v
	return s
}

type CreateOrUpdateContactResponseBody struct {
	AlertContact *CreateOrUpdateContactResponseBodyAlertContact `json:"AlertContact,omitempty" xml:"AlertContact,omitempty" type:"Struct"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactResponseBody) SetAlertContact(v *CreateOrUpdateContactResponseBodyAlertContact) *CreateOrUpdateContactResponseBody {
	s.AlertContact = v
	return s
}

func (s *CreateOrUpdateContactResponseBody) SetRequestId(v string) *CreateOrUpdateContactResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrUpdateContactResponseBodyAlertContact struct {
	ContactId         *float32 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName       *string  `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Email             *string  `json:"Email,omitempty" xml:"Email,omitempty"`
	IsVerify          *bool    `json:"IsVerify,omitempty" xml:"IsVerify,omitempty"`
	Phone             *string  `json:"Phone,omitempty" xml:"Phone,omitempty"`
	ReissueSendNotice *int64   `json:"ReissueSendNotice,omitempty" xml:"ReissueSendNotice,omitempty"`
}

func (s CreateOrUpdateContactResponseBodyAlertContact) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateContactResponseBodyAlertContact) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetContactId(v float32) *CreateOrUpdateContactResponseBodyAlertContact {
	s.ContactId = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetContactName(v string) *CreateOrUpdateContactResponseBodyAlertContact {
	s.ContactName = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetEmail(v string) *CreateOrUpdateContactResponseBodyAlertContact {
	s.Email = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetIsVerify(v bool) *CreateOrUpdateContactResponseBodyAlertContact {
	s.IsVerify = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetPhone(v string) *CreateOrUpdateContactResponseBodyAlertContact {
	s.Phone = &v
	return s
}

func (s *CreateOrUpdateContactResponseBodyAlertContact) SetReissueSendNotice(v int64) *CreateOrUpdateContactResponseBodyAlertContact {
	s.ReissueSendNotice = &v
	return s
}

type CreateOrUpdateContactResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrUpdateContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrUpdateContactResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateContactResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactResponse) SetHeaders(v map[string]*string) *CreateOrUpdateContactResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateContactResponse) SetStatusCode(v int32) *CreateOrUpdateContactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateContactResponse) SetBody(v *CreateOrUpdateContactResponseBody) *CreateOrUpdateContactResponse {
	s.Body = v
	return s
}

type CreateOrUpdateContactGroupRequest struct {
	ContactGroupId   *int64  `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	ContactIds       *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
}

func (s CreateOrUpdateContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateContactGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactGroupRequest) SetContactGroupId(v int64) *CreateOrUpdateContactGroupRequest {
	s.ContactGroupId = &v
	return s
}

func (s *CreateOrUpdateContactGroupRequest) SetContactGroupName(v string) *CreateOrUpdateContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *CreateOrUpdateContactGroupRequest) SetContactIds(v string) *CreateOrUpdateContactGroupRequest {
	s.ContactIds = &v
	return s
}

type CreateOrUpdateContactGroupResponseBody struct {
	AlertContactGroup *CreateOrUpdateContactGroupResponseBodyAlertContactGroup `json:"AlertContactGroup,omitempty" xml:"AlertContactGroup,omitempty" type:"Struct"`
	RequestId         *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactGroupResponseBody) SetAlertContactGroup(v *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) *CreateOrUpdateContactGroupResponseBody {
	s.AlertContactGroup = v
	return s
}

func (s *CreateOrUpdateContactGroupResponseBody) SetRequestId(v string) *CreateOrUpdateContactGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrUpdateContactGroupResponseBodyAlertContactGroup struct {
	ContactGroupId   *float32 `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	ContactGroupName *string  `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	ContactIds       *string  `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
}

func (s CreateOrUpdateContactGroupResponseBodyAlertContactGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateContactGroupResponseBodyAlertContactGroup) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) SetContactGroupId(v float32) *CreateOrUpdateContactGroupResponseBodyAlertContactGroup {
	s.ContactGroupId = &v
	return s
}

func (s *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) SetContactGroupName(v string) *CreateOrUpdateContactGroupResponseBodyAlertContactGroup {
	s.ContactGroupName = &v
	return s
}

func (s *CreateOrUpdateContactGroupResponseBodyAlertContactGroup) SetContactIds(v string) *CreateOrUpdateContactGroupResponseBodyAlertContactGroup {
	s.ContactIds = &v
	return s
}

type CreateOrUpdateContactGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrUpdateContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrUpdateContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateContactGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateContactGroupResponse) SetHeaders(v map[string]*string) *CreateOrUpdateContactGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateContactGroupResponse) SetStatusCode(v int32) *CreateOrUpdateContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateContactGroupResponse) SetBody(v *CreateOrUpdateContactGroupResponseBody) *CreateOrUpdateContactGroupResponse {
	s.Body = v
	return s
}

type CreateOrUpdateEventBridgeIntegrationRequest struct {
	AccessKey        *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AccessSecret     *string `json:"AccessSecret,omitempty" xml:"AccessSecret,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Endpoint         *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EventBusName     *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventBusRegionId *string `json:"EventBusRegionId,omitempty" xml:"EventBusRegionId,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Source           *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateOrUpdateEventBridgeIntegrationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateEventBridgeIntegrationRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetAccessKey(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.AccessKey = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetAccessSecret(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.AccessSecret = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetDescription(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.Description = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetEndpoint(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.Endpoint = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetEventBusName(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.EventBusName = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetEventBusRegionId(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.EventBusRegionId = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetId(v int64) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetName(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationRequest) SetSource(v string) *CreateOrUpdateEventBridgeIntegrationRequest {
	s.Source = &v
	return s
}

type CreateOrUpdateEventBridgeIntegrationResponseBody struct {
	EventBridgeIntegration *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration `json:"EventBridgeIntegration,omitempty" xml:"EventBridgeIntegration,omitempty" type:"Struct"`
	RequestId              *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateEventBridgeIntegrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateEventBridgeIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBody) SetEventBridgeIntegration(v *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) *CreateOrUpdateEventBridgeIntegrationResponseBody {
	s.EventBridgeIntegration = v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBody) SetRequestId(v string) *CreateOrUpdateEventBridgeIntegrationResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration struct {
	AccessKey        *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	AccessSecret     *string `json:"AccessSecret,omitempty" xml:"AccessSecret,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Endpoint         *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	EventBusName     *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
	EventBusRegionId *string `json:"EventBusRegionId,omitempty" xml:"EventBusRegionId,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Source           *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetAccessKey(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.AccessKey = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetAccessSecret(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.AccessSecret = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetDescription(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.Description = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetEndpoint(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.Endpoint = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetEventBusName(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.EventBusName = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetEventBusRegionId(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.EventBusRegionId = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetId(v int64) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetName(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration) SetSource(v string) *CreateOrUpdateEventBridgeIntegrationResponseBodyEventBridgeIntegration {
	s.Source = &v
	return s
}

type CreateOrUpdateEventBridgeIntegrationResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrUpdateEventBridgeIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrUpdateEventBridgeIntegrationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateEventBridgeIntegrationResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateEventBridgeIntegrationResponse) SetHeaders(v map[string]*string) *CreateOrUpdateEventBridgeIntegrationResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponse) SetStatusCode(v int32) *CreateOrUpdateEventBridgeIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateEventBridgeIntegrationResponse) SetBody(v *CreateOrUpdateEventBridgeIntegrationResponseBody) *CreateOrUpdateEventBridgeIntegrationResponse {
	s.Body = v
	return s
}

type CreateOrUpdateIMRobotRequest struct {
	CardTemplate   *string `json:"CardTemplate,omitempty" xml:"CardTemplate,omitempty"`
	DailyNoc       *bool   `json:"DailyNoc,omitempty" xml:"DailyNoc,omitempty"`
	DailyNocTime   *string `json:"DailyNocTime,omitempty" xml:"DailyNocTime,omitempty"`
	EnableOutgoing *bool   `json:"EnableOutgoing,omitempty" xml:"EnableOutgoing,omitempty"`
	RobotAddress   *string `json:"RobotAddress,omitempty" xml:"RobotAddress,omitempty"`
	RobotId        *int64  `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	RobotName      *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	Token          *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateOrUpdateIMRobotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateIMRobotRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateIMRobotRequest) SetCardTemplate(v string) *CreateOrUpdateIMRobotRequest {
	s.CardTemplate = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetDailyNoc(v bool) *CreateOrUpdateIMRobotRequest {
	s.DailyNoc = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetDailyNocTime(v string) *CreateOrUpdateIMRobotRequest {
	s.DailyNocTime = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetEnableOutgoing(v bool) *CreateOrUpdateIMRobotRequest {
	s.EnableOutgoing = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetRobotAddress(v string) *CreateOrUpdateIMRobotRequest {
	s.RobotAddress = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetRobotId(v int64) *CreateOrUpdateIMRobotRequest {
	s.RobotId = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetRobotName(v string) *CreateOrUpdateIMRobotRequest {
	s.RobotName = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetToken(v string) *CreateOrUpdateIMRobotRequest {
	s.Token = &v
	return s
}

func (s *CreateOrUpdateIMRobotRequest) SetType(v string) *CreateOrUpdateIMRobotRequest {
	s.Type = &v
	return s
}

type CreateOrUpdateIMRobotResponseBody struct {
	AlertRobot *CreateOrUpdateIMRobotResponseBodyAlertRobot `json:"AlertRobot,omitempty" xml:"AlertRobot,omitempty" type:"Struct"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateIMRobotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateIMRobotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateIMRobotResponseBody) SetAlertRobot(v *CreateOrUpdateIMRobotResponseBodyAlertRobot) *CreateOrUpdateIMRobotResponseBody {
	s.AlertRobot = v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBody) SetRequestId(v string) *CreateOrUpdateIMRobotResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrUpdateIMRobotResponseBodyAlertRobot struct {
	CardTemplate   *string  `json:"CardTemplate,omitempty" xml:"CardTemplate,omitempty"`
	DailyNoc       *bool    `json:"DailyNoc,omitempty" xml:"DailyNoc,omitempty"`
	DailyNocTime   *string  `json:"DailyNocTime,omitempty" xml:"DailyNocTime,omitempty"`
	EnableOutgoing *bool    `json:"EnableOutgoing,omitempty" xml:"EnableOutgoing,omitempty"`
	RobotAddress   *string  `json:"RobotAddress,omitempty" xml:"RobotAddress,omitempty"`
	RobotId        *float32 `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	RobotName      *string  `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	Token          *string  `json:"Token,omitempty" xml:"Token,omitempty"`
	Type           *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateOrUpdateIMRobotResponseBodyAlertRobot) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateIMRobotResponseBodyAlertRobot) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetCardTemplate(v string) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.CardTemplate = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetDailyNoc(v bool) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.DailyNoc = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetDailyNocTime(v string) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.DailyNocTime = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetEnableOutgoing(v bool) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.EnableOutgoing = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetRobotAddress(v string) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.RobotAddress = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetRobotId(v float32) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.RobotId = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetRobotName(v string) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.RobotName = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetToken(v string) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.Token = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponseBodyAlertRobot) SetType(v string) *CreateOrUpdateIMRobotResponseBodyAlertRobot {
	s.Type = &v
	return s
}

type CreateOrUpdateIMRobotResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrUpdateIMRobotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrUpdateIMRobotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateIMRobotResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateIMRobotResponse) SetHeaders(v map[string]*string) *CreateOrUpdateIMRobotResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateIMRobotResponse) SetStatusCode(v int32) *CreateOrUpdateIMRobotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponse) SetBody(v *CreateOrUpdateIMRobotResponseBody) *CreateOrUpdateIMRobotResponse {
	s.Body = v
	return s
}

type CreateOrUpdateNotificationPolicyRequest struct {
	EscalationPolicyId *int64  `json:"EscalationPolicyId,omitempty" xml:"EscalationPolicyId,omitempty"`
	GroupRule          *string `json:"GroupRule,omitempty" xml:"GroupRule,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IntegrationId      *int64  `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
	MatchingRules      *string `json:"MatchingRules,omitempty" xml:"MatchingRules,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NotifyRule         *string `json:"NotifyRule,omitempty" xml:"NotifyRule,omitempty"`
	NotifyTemplate     *string `json:"NotifyTemplate,omitempty" xml:"NotifyTemplate,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Repeat             *bool   `json:"Repeat,omitempty" xml:"Repeat,omitempty"`
	RepeatInterval     *int64  `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	SendRecoverMessage *bool   `json:"SendRecoverMessage,omitempty" xml:"SendRecoverMessage,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetEscalationPolicyId(v int64) *CreateOrUpdateNotificationPolicyRequest {
	s.EscalationPolicyId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetGroupRule(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.GroupRule = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetId(v int64) *CreateOrUpdateNotificationPolicyRequest {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetIntegrationId(v int64) *CreateOrUpdateNotificationPolicyRequest {
	s.IntegrationId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetMatchingRules(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.MatchingRules = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetName(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetNotifyRule(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.NotifyRule = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetNotifyTemplate(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.NotifyTemplate = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetRegionId(v string) *CreateOrUpdateNotificationPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetRepeat(v bool) *CreateOrUpdateNotificationPolicyRequest {
	s.Repeat = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetRepeatInterval(v int64) *CreateOrUpdateNotificationPolicyRequest {
	s.RepeatInterval = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyRequest) SetSendRecoverMessage(v bool) *CreateOrUpdateNotificationPolicyRequest {
	s.SendRecoverMessage = &v
	return s
}

type CreateOrUpdateNotificationPolicyResponseBody struct {
	NotificationPolicy *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy `json:"NotificationPolicy,omitempty" xml:"NotificationPolicy,omitempty" type:"Struct"`
	RequestId          *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBody) SetNotificationPolicy(v *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) *CreateOrUpdateNotificationPolicyResponseBody {
	s.NotificationPolicy = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBody) SetRequestId(v string) *CreateOrUpdateNotificationPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy struct {
	EscalationPolicyId *int64                                                                         `json:"EscalationPolicyId,omitempty" xml:"EscalationPolicyId,omitempty"`
	GroupRule          *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule       `json:"GroupRule,omitempty" xml:"GroupRule,omitempty" type:"Struct"`
	Id                 *int64                                                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	IntegrationId      *int64                                                                         `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
	MatchingRules      []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules `json:"MatchingRules,omitempty" xml:"MatchingRules,omitempty" type:"Repeated"`
	Name               *string                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	NotifyRule         *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule      `json:"NotifyRule,omitempty" xml:"NotifyRule,omitempty" type:"Struct"`
	NotifyTemplate     *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate  `json:"NotifyTemplate,omitempty" xml:"NotifyTemplate,omitempty" type:"Struct"`
	Repeat             *bool                                                                          `json:"Repeat,omitempty" xml:"Repeat,omitempty"`
	RepeatInterval     *int64                                                                         `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	SendRecoverMessage *bool                                                                          `json:"SendRecoverMessage,omitempty" xml:"SendRecoverMessage,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetEscalationPolicyId(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.EscalationPolicyId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetGroupRule(v *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.GroupRule = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetId(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetIntegrationId(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.IntegrationId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetMatchingRules(v []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.MatchingRules = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetName(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetNotifyRule(v *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.NotifyRule = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetNotifyTemplate(v *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.NotifyTemplate = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetRepeat(v bool) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.Repeat = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetRepeatInterval(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.RepeatInterval = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy) SetSendRecoverMessage(v bool) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicy {
	s.SendRecoverMessage = &v
	return s
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule struct {
	GroupInterval  *int64    `json:"GroupInterval,omitempty" xml:"GroupInterval,omitempty"`
	GroupWait      *int64    `json:"GroupWait,omitempty" xml:"GroupWait,omitempty"`
	GroupingFields []*string `json:"GroupingFields,omitempty" xml:"GroupingFields,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) SetGroupInterval(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule {
	s.GroupInterval = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) SetGroupWait(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule {
	s.GroupWait = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule) SetGroupingFields(v []*string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyGroupRule {
	s.GroupingFields = v
	return s
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules struct {
	MatchingConditions []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions `json:"MatchingConditions,omitempty" xml:"MatchingConditions,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules) SetMatchingConditions(v []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRules {
	s.MatchingConditions = v
	return s
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) SetKey(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions {
	s.Key = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) SetOperator(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions {
	s.Operator = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions) SetValue(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyMatchingRulesMatchingConditions {
	s.Value = &v
	return s
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule struct {
	NotifyChannels  []*string                                                                                `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	NotifyEndTime   *string                                                                                  `json:"NotifyEndTime,omitempty" xml:"NotifyEndTime,omitempty"`
	NotifyObjects   []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects `json:"NotifyObjects,omitempty" xml:"NotifyObjects,omitempty" type:"Repeated"`
	NotifyStartTime *string                                                                                  `json:"NotifyStartTime,omitempty" xml:"NotifyStartTime,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) SetNotifyChannels(v []*string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule {
	s.NotifyChannels = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) SetNotifyEndTime(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule {
	s.NotifyEndTime = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) SetNotifyObjects(v []*CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule {
	s.NotifyObjects = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule) SetNotifyStartTime(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRule {
	s.NotifyStartTime = &v
	return s
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects struct {
	NotifyObjectId   *int64  `json:"NotifyObjectId,omitempty" xml:"NotifyObjectId,omitempty"`
	NotifyObjectName *string `json:"NotifyObjectName,omitempty" xml:"NotifyObjectName,omitempty"`
	NotifyObjectType *string `json:"NotifyObjectType,omitempty" xml:"NotifyObjectType,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) SetNotifyObjectId(v int64) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects {
	s.NotifyObjectId = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) SetNotifyObjectName(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects {
	s.NotifyObjectName = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects) SetNotifyObjectType(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyRuleNotifyObjects {
	s.NotifyObjectType = &v
	return s
}

type CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate struct {
	EmailContent        *string `json:"EmailContent,omitempty" xml:"EmailContent,omitempty"`
	EmailRecoverContent *string `json:"EmailRecoverContent,omitempty" xml:"EmailRecoverContent,omitempty"`
	EmailRecoverTitle   *string `json:"EmailRecoverTitle,omitempty" xml:"EmailRecoverTitle,omitempty"`
	EmailTitle          *string `json:"EmailTitle,omitempty" xml:"EmailTitle,omitempty"`
	RobotContent        *string `json:"RobotContent,omitempty" xml:"RobotContent,omitempty"`
	SmsContent          *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	SmsRecoverContent   *string `json:"SmsRecoverContent,omitempty" xml:"SmsRecoverContent,omitempty"`
	TtsContent          *string `json:"TtsContent,omitempty" xml:"TtsContent,omitempty"`
	TtsRecoverContent   *string `json:"TtsRecoverContent,omitempty" xml:"TtsRecoverContent,omitempty"`
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetEmailContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.EmailContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetEmailRecoverContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.EmailRecoverContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetEmailRecoverTitle(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.EmailRecoverTitle = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetEmailTitle(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.EmailTitle = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetRobotContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.RobotContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetSmsContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.SmsContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetSmsRecoverContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.SmsRecoverContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetTtsContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.TtsContent = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate) SetTtsRecoverContent(v string) *CreateOrUpdateNotificationPolicyResponseBodyNotificationPolicyNotifyTemplate {
	s.TtsRecoverContent = &v
	return s
}

type CreateOrUpdateNotificationPolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrUpdateNotificationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrUpdateNotificationPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateNotificationPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateNotificationPolicyResponse) SetHeaders(v map[string]*string) *CreateOrUpdateNotificationPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponse) SetStatusCode(v int32) *CreateOrUpdateNotificationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateNotificationPolicyResponse) SetBody(v *CreateOrUpdateNotificationPolicyResponseBody) *CreateOrUpdateNotificationPolicyResponse {
	s.Body = v
	return s
}

type CreateOrUpdateSilencePolicyRequest struct {
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	MatchingRules *string `json:"MatchingRules,omitempty" xml:"MatchingRules,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateOrUpdateSilencePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateSilencePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSilencePolicyRequest) SetId(v int64) *CreateOrUpdateSilencePolicyRequest {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyRequest) SetMatchingRules(v string) *CreateOrUpdateSilencePolicyRequest {
	s.MatchingRules = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyRequest) SetName(v string) *CreateOrUpdateSilencePolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyRequest) SetRegionId(v string) *CreateOrUpdateSilencePolicyRequest {
	s.RegionId = &v
	return s
}

type CreateOrUpdateSilencePolicyResponseBody struct {
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SilencePolicy *CreateOrUpdateSilencePolicyResponseBodySilencePolicy `json:"SilencePolicy,omitempty" xml:"SilencePolicy,omitempty" type:"Struct"`
}

func (s CreateOrUpdateSilencePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateSilencePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSilencePolicyResponseBody) SetRequestId(v string) *CreateOrUpdateSilencePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBody) SetSilencePolicy(v *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) *CreateOrUpdateSilencePolicyResponseBody {
	s.SilencePolicy = v
	return s
}

type CreateOrUpdateSilencePolicyResponseBodySilencePolicy struct {
	Id            *int64                                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	MatchingRules []*CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules `json:"MatchingRules,omitempty" xml:"MatchingRules,omitempty" type:"Repeated"`
	Name          *string                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateOrUpdateSilencePolicyResponseBodySilencePolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateSilencePolicyResponseBodySilencePolicy) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) SetId(v int64) *CreateOrUpdateSilencePolicyResponseBodySilencePolicy {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) SetMatchingRules(v []*CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules) *CreateOrUpdateSilencePolicyResponseBodySilencePolicy {
	s.MatchingRules = v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) SetName(v string) *CreateOrUpdateSilencePolicyResponseBodySilencePolicy {
	s.Name = &v
	return s
}

type CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules struct {
	MatchingConditions []*CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions `json:"MatchingConditions,omitempty" xml:"MatchingConditions,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules) SetMatchingConditions(v []*CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules {
	s.MatchingConditions = v
	return s
}

type CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) SetKey(v string) *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions {
	s.Key = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) SetOperator(v string) *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions {
	s.Operator = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) SetValue(v string) *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions {
	s.Value = &v
	return s
}

type CreateOrUpdateSilencePolicyResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrUpdateSilencePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrUpdateSilencePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateSilencePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSilencePolicyResponse) SetHeaders(v map[string]*string) *CreateOrUpdateSilencePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponse) SetStatusCode(v int32) *CreateOrUpdateSilencePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponse) SetBody(v *CreateOrUpdateSilencePolicyResponseBody) *CreateOrUpdateSilencePolicyResponse {
	s.Body = v
	return s
}

type CreateOrUpdateWebhookContactRequest struct {
	BizHeaders  *string `json:"BizHeaders,omitempty" xml:"BizHeaders,omitempty"`
	BizParams   *string `json:"BizParams,omitempty" xml:"BizParams,omitempty"`
	Body        *string `json:"Body,omitempty" xml:"Body,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	RecoverBody *string `json:"RecoverBody,omitempty" xml:"RecoverBody,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
	WebhookId   *int64  `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
	WebhookName *string `json:"WebhookName,omitempty" xml:"WebhookName,omitempty"`
}

func (s CreateOrUpdateWebhookContactRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateWebhookContactRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateWebhookContactRequest) SetBizHeaders(v string) *CreateOrUpdateWebhookContactRequest {
	s.BizHeaders = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetBizParams(v string) *CreateOrUpdateWebhookContactRequest {
	s.BizParams = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetBody(v string) *CreateOrUpdateWebhookContactRequest {
	s.Body = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetMethod(v string) *CreateOrUpdateWebhookContactRequest {
	s.Method = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetRecoverBody(v string) *CreateOrUpdateWebhookContactRequest {
	s.RecoverBody = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetUrl(v string) *CreateOrUpdateWebhookContactRequest {
	s.Url = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetWebhookId(v int64) *CreateOrUpdateWebhookContactRequest {
	s.WebhookId = &v
	return s
}

func (s *CreateOrUpdateWebhookContactRequest) SetWebhookName(v string) *CreateOrUpdateWebhookContactRequest {
	s.WebhookName = &v
	return s
}

type CreateOrUpdateWebhookContactResponseBody struct {
	RequestId      *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebhookContact *CreateOrUpdateWebhookContactResponseBodyWebhookContact `json:"WebhookContact,omitempty" xml:"WebhookContact,omitempty" type:"Struct"`
}

func (s CreateOrUpdateWebhookContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateWebhookContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateWebhookContactResponseBody) SetRequestId(v string) *CreateOrUpdateWebhookContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBody) SetWebhookContact(v *CreateOrUpdateWebhookContactResponseBodyWebhookContact) *CreateOrUpdateWebhookContactResponseBody {
	s.WebhookContact = v
	return s
}

type CreateOrUpdateWebhookContactResponseBodyWebhookContact struct {
	Webhook     *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook `json:"Webhook,omitempty" xml:"Webhook,omitempty" type:"Struct"`
	WebhookId   *float32                                                       `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
	WebhookName *string                                                        `json:"WebhookName,omitempty" xml:"WebhookName,omitempty"`
}

func (s CreateOrUpdateWebhookContactResponseBodyWebhookContact) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateWebhookContactResponseBodyWebhookContact) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContact) SetWebhook(v *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) *CreateOrUpdateWebhookContactResponseBodyWebhookContact {
	s.Webhook = v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContact) SetWebhookId(v float32) *CreateOrUpdateWebhookContactResponseBodyWebhookContact {
	s.WebhookId = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContact) SetWebhookName(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContact {
	s.WebhookName = &v
	return s
}

type CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook struct {
	BizHeaders  *string `json:"BizHeaders,omitempty" xml:"BizHeaders,omitempty"`
	BizParams   *string `json:"BizParams,omitempty" xml:"BizParams,omitempty"`
	Body        *string `json:"Body,omitempty" xml:"Body,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	RecoverBody *string `json:"RecoverBody,omitempty" xml:"RecoverBody,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) SetBizHeaders(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	s.BizHeaders = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) SetBizParams(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	s.BizParams = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) SetBody(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	s.Body = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) SetMethod(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	s.Method = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) SetRecoverBody(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	s.RecoverBody = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook) SetUrl(v string) *CreateOrUpdateWebhookContactResponseBodyWebhookContactWebhook {
	s.Url = &v
	return s
}

type CreateOrUpdateWebhookContactResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOrUpdateWebhookContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOrUpdateWebhookContactResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrUpdateWebhookContactResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateWebhookContactResponse) SetHeaders(v map[string]*string) *CreateOrUpdateWebhookContactResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateWebhookContactResponse) SetStatusCode(v int32) *CreateOrUpdateWebhookContactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateWebhookContactResponse) SetBody(v *CreateOrUpdateWebhookContactResponseBody) *CreateOrUpdateWebhookContactResponse {
	s.Body = v
	return s
}

type CreatePrometheusAlertRuleRequest struct {
	AlertName      *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64  `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePrometheusAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleRequest) SetAlertName(v string) *CreatePrometheusAlertRuleRequest {
	s.AlertName = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetAnnotations(v string) *CreatePrometheusAlertRuleRequest {
	s.Annotations = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetClusterId(v string) *CreatePrometheusAlertRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetDispatchRuleId(v int64) *CreatePrometheusAlertRuleRequest {
	s.DispatchRuleId = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetDuration(v string) *CreatePrometheusAlertRuleRequest {
	s.Duration = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetExpression(v string) *CreatePrometheusAlertRuleRequest {
	s.Expression = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetLabels(v string) *CreatePrometheusAlertRuleRequest {
	s.Labels = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetMessage(v string) *CreatePrometheusAlertRuleRequest {
	s.Message = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetNotifyType(v string) *CreatePrometheusAlertRuleRequest {
	s.NotifyType = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetRegionId(v string) *CreatePrometheusAlertRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrometheusAlertRuleRequest) SetType(v string) *CreatePrometheusAlertRuleRequest {
	s.Type = &v
	return s
}

type CreatePrometheusAlertRuleResponseBody struct {
	PrometheusAlertRule *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule `json:"PrometheusAlertRule,omitempty" xml:"PrometheusAlertRule,omitempty" type:"Struct"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBody) SetPrometheusAlertRule(v *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) *CreatePrometheusAlertRuleResponseBody {
	s.PrometheusAlertRule = v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBody) SetRequestId(v string) *CreatePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule struct {
	AlertId        *int64                                                                 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string                                                                `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	ClusterId      *string                                                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64                                                                 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string                                                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string                                                                `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message        *string                                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string                                                                `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Status         *int32                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertId(v int64) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertName(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertName = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAnnotations(v []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Annotations = v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetClusterId(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.ClusterId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDispatchRuleId(v int64) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.DispatchRuleId = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDuration(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Duration = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetExpression(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Expression = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetLabels(v []*CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Labels = v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetMessage(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Message = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetNotifyType(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.NotifyType = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetStatus(v int32) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Status = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetType(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Type = &v
	return s
}

type CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetName(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetValue(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Value = &v
	return s
}

type CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetName(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetValue(v string) *CreatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Value = &v
	return s
}

type CreatePrometheusAlertRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePrometheusAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePrometheusAlertRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePrometheusAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *CreatePrometheusAlertRuleResponse) SetHeaders(v map[string]*string) *CreatePrometheusAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *CreatePrometheusAlertRuleResponse) SetStatusCode(v int32) *CreatePrometheusAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrometheusAlertRuleResponse) SetBody(v *CreatePrometheusAlertRuleResponseBody) *CreatePrometheusAlertRuleResponse {
	s.Body = v
	return s
}

type CreateRetcodeAppRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RetcodeAppName *string `json:"RetcodeAppName,omitempty" xml:"RetcodeAppName,omitempty"`
	RetcodeAppType *string `json:"RetcodeAppType,omitempty" xml:"RetcodeAppType,omitempty"`
}

func (s CreateRetcodeAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRetcodeAppRequest) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppRequest) SetRegionId(v string) *CreateRetcodeAppRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRetcodeAppRequest) SetRetcodeAppName(v string) *CreateRetcodeAppRequest {
	s.RetcodeAppName = &v
	return s
}

func (s *CreateRetcodeAppRequest) SetRetcodeAppType(v string) *CreateRetcodeAppRequest {
	s.RetcodeAppType = &v
	return s
}

type CreateRetcodeAppResponseBody struct {
	RequestId          *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetcodeAppDataBean *CreateRetcodeAppResponseBodyRetcodeAppDataBean `json:"RetcodeAppDataBean,omitempty" xml:"RetcodeAppDataBean,omitempty" type:"Struct"`
}

func (s CreateRetcodeAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRetcodeAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponseBody) SetRequestId(v string) *CreateRetcodeAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRetcodeAppResponseBody) SetRetcodeAppDataBean(v *CreateRetcodeAppResponseBodyRetcodeAppDataBean) *CreateRetcodeAppResponseBody {
	s.RetcodeAppDataBean = v
	return s
}

type CreateRetcodeAppResponseBodyRetcodeAppDataBean struct {
	AppId *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Pid   *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBean) String() string {
	return tea.Prettify(s)
}

func (s CreateRetcodeAppResponseBodyRetcodeAppDataBean) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) SetAppId(v int64) *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	s.AppId = &v
	return s
}

func (s *CreateRetcodeAppResponseBodyRetcodeAppDataBean) SetPid(v string) *CreateRetcodeAppResponseBodyRetcodeAppDataBean {
	s.Pid = &v
	return s
}

type CreateRetcodeAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRetcodeAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRetcodeAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRetcodeAppResponse) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppResponse) SetHeaders(v map[string]*string) *CreateRetcodeAppResponse {
	s.Headers = v
	return s
}

func (s *CreateRetcodeAppResponse) SetStatusCode(v int32) *CreateRetcodeAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRetcodeAppResponse) SetBody(v *CreateRetcodeAppResponseBody) *CreateRetcodeAppResponse {
	s.Body = v
	return s
}

type CreateSyntheticTaskRequest struct {
	CommonParam    *CreateSyntheticTaskRequestCommonParam    `json:"CommonParam,omitempty" xml:"CommonParam,omitempty" type:"Struct"`
	ExtendInterval *CreateSyntheticTaskRequestExtendInterval `json:"ExtendInterval,omitempty" xml:"ExtendInterval,omitempty" type:"Struct"`
	IntervalTime   *string                                   `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	IntervalType   *string                                   `json:"IntervalType,omitempty" xml:"IntervalType,omitempty"`
	IpType         *int64                                    `json:"IpType,omitempty" xml:"IpType,omitempty"`
	MonitorList    []*CreateSyntheticTaskRequestMonitorList  `json:"MonitorList,omitempty" xml:"MonitorList,omitempty" type:"Repeated"`
	Net            *CreateSyntheticTaskRequestNet            `json:"Net,omitempty" xml:"Net,omitempty" type:"Struct"`
	RegionId       *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskName       *string                                   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskType       *int64                                    `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	Url            *string                                   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateSyntheticTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSyntheticTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequest) SetCommonParam(v *CreateSyntheticTaskRequestCommonParam) *CreateSyntheticTaskRequest {
	s.CommonParam = v
	return s
}

func (s *CreateSyntheticTaskRequest) SetExtendInterval(v *CreateSyntheticTaskRequestExtendInterval) *CreateSyntheticTaskRequest {
	s.ExtendInterval = v
	return s
}

func (s *CreateSyntheticTaskRequest) SetIntervalTime(v string) *CreateSyntheticTaskRequest {
	s.IntervalTime = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetIntervalType(v string) *CreateSyntheticTaskRequest {
	s.IntervalType = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetIpType(v int64) *CreateSyntheticTaskRequest {
	s.IpType = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetMonitorList(v []*CreateSyntheticTaskRequestMonitorList) *CreateSyntheticTaskRequest {
	s.MonitorList = v
	return s
}

func (s *CreateSyntheticTaskRequest) SetNet(v *CreateSyntheticTaskRequestNet) *CreateSyntheticTaskRequest {
	s.Net = v
	return s
}

func (s *CreateSyntheticTaskRequest) SetRegionId(v string) *CreateSyntheticTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetTaskName(v string) *CreateSyntheticTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetTaskType(v int64) *CreateSyntheticTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateSyntheticTaskRequest) SetUrl(v string) *CreateSyntheticTaskRequest {
	s.Url = &v
	return s
}

type CreateSyntheticTaskRequestCommonParam struct {
	AlarmFlag          *string                                           `json:"AlarmFlag,omitempty" xml:"AlarmFlag,omitempty"`
	AlertList          []*CreateSyntheticTaskRequestCommonParamAlertList `json:"AlertList,omitempty" xml:"AlertList,omitempty" type:"Repeated"`
	AlertNotifierId    *string                                           `json:"AlertNotifierId,omitempty" xml:"AlertNotifierId,omitempty"`
	AlertPolicyId      *string                                           `json:"AlertPolicyId,omitempty" xml:"AlertPolicyId,omitempty"`
	MonitorSamples     *int64                                            `json:"MonitorSamples,omitempty" xml:"MonitorSamples,omitempty"`
	StartExecutionTime *int64                                            `json:"StartExecutionTime,omitempty" xml:"StartExecutionTime,omitempty"`
}

func (s CreateSyntheticTaskRequestCommonParam) String() string {
	return tea.Prettify(s)
}

func (s CreateSyntheticTaskRequestCommonParam) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestCommonParam) SetAlarmFlag(v string) *CreateSyntheticTaskRequestCommonParam {
	s.AlarmFlag = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParam) SetAlertList(v []*CreateSyntheticTaskRequestCommonParamAlertList) *CreateSyntheticTaskRequestCommonParam {
	s.AlertList = v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParam) SetAlertNotifierId(v string) *CreateSyntheticTaskRequestCommonParam {
	s.AlertNotifierId = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParam) SetAlertPolicyId(v string) *CreateSyntheticTaskRequestCommonParam {
	s.AlertPolicyId = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParam) SetMonitorSamples(v int64) *CreateSyntheticTaskRequestCommonParam {
	s.MonitorSamples = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParam) SetStartExecutionTime(v int64) *CreateSyntheticTaskRequestCommonParam {
	s.StartExecutionTime = &v
	return s
}

type CreateSyntheticTaskRequestCommonParamAlertList struct {
	IsCritical *int64  `json:"IsCritical,omitempty" xml:"IsCritical,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Symbols    *int64  `json:"Symbols,omitempty" xml:"Symbols,omitempty"`
}

func (s CreateSyntheticTaskRequestCommonParamAlertList) String() string {
	return tea.Prettify(s)
}

func (s CreateSyntheticTaskRequestCommonParamAlertList) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestCommonParamAlertList) SetIsCritical(v int64) *CreateSyntheticTaskRequestCommonParamAlertList {
	s.IsCritical = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParamAlertList) SetName(v string) *CreateSyntheticTaskRequestCommonParamAlertList {
	s.Name = &v
	return s
}

func (s *CreateSyntheticTaskRequestCommonParamAlertList) SetSymbols(v int64) *CreateSyntheticTaskRequestCommonParamAlertList {
	s.Symbols = &v
	return s
}

type CreateSyntheticTaskRequestExtendInterval struct {
	Days        []*int64 `json:"Days,omitempty" xml:"Days,omitempty" type:"Repeated"`
	EndHour     *int64   `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	EndMinute   *int64   `json:"EndMinute,omitempty" xml:"EndMinute,omitempty"`
	EndTime     *string  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartHour   *int64   `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
	StartMinute *int64   `json:"StartMinute,omitempty" xml:"StartMinute,omitempty"`
	StartTime   *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateSyntheticTaskRequestExtendInterval) String() string {
	return tea.Prettify(s)
}

func (s CreateSyntheticTaskRequestExtendInterval) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetDays(v []*int64) *CreateSyntheticTaskRequestExtendInterval {
	s.Days = v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetEndHour(v int64) *CreateSyntheticTaskRequestExtendInterval {
	s.EndHour = &v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetEndMinute(v int64) *CreateSyntheticTaskRequestExtendInterval {
	s.EndMinute = &v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetEndTime(v string) *CreateSyntheticTaskRequestExtendInterval {
	s.EndTime = &v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetStartHour(v int64) *CreateSyntheticTaskRequestExtendInterval {
	s.StartHour = &v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetStartMinute(v int64) *CreateSyntheticTaskRequestExtendInterval {
	s.StartMinute = &v
	return s
}

func (s *CreateSyntheticTaskRequestExtendInterval) SetStartTime(v string) *CreateSyntheticTaskRequestExtendInterval {
	s.StartTime = &v
	return s
}

type CreateSyntheticTaskRequestMonitorList struct {
	CityCode     *int64 `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	MonitorType  *int64 `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	NetServiceId *int64 `json:"NetServiceId,omitempty" xml:"NetServiceId,omitempty"`
}

func (s CreateSyntheticTaskRequestMonitorList) String() string {
	return tea.Prettify(s)
}

func (s CreateSyntheticTaskRequestMonitorList) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestMonitorList) SetCityCode(v int64) *CreateSyntheticTaskRequestMonitorList {
	s.CityCode = &v
	return s
}

func (s *CreateSyntheticTaskRequestMonitorList) SetMonitorType(v int64) *CreateSyntheticTaskRequestMonitorList {
	s.MonitorType = &v
	return s
}

func (s *CreateSyntheticTaskRequestMonitorList) SetNetServiceId(v int64) *CreateSyntheticTaskRequestMonitorList {
	s.NetServiceId = &v
	return s
}

type CreateSyntheticTaskRequestNet struct {
	NetDNSNs             *string `json:"NetDNSNs,omitempty" xml:"NetDNSNs,omitempty"`
	NetDNSQueryMethod    *int64  `json:"NetDNSQueryMethod,omitempty" xml:"NetDNSQueryMethod,omitempty"`
	NetDNSServer         *int64  `json:"NetDNSServer,omitempty" xml:"NetDNSServer,omitempty"`
	NetDNSSwitch         *int64  `json:"NetDNSSwitch,omitempty" xml:"NetDNSSwitch,omitempty"`
	NetDNSTimeout        *int64  `json:"NetDNSTimeout,omitempty" xml:"NetDNSTimeout,omitempty"`
	NetDigSwitch         *int64  `json:"NetDigSwitch,omitempty" xml:"NetDigSwitch,omitempty"`
	NetICMPActive        *int64  `json:"NetICMPActive,omitempty" xml:"NetICMPActive,omitempty"`
	NetICMPInterval      *int64  `json:"NetICMPInterval,omitempty" xml:"NetICMPInterval,omitempty"`
	NetICMPNum           *int64  `json:"NetICMPNum,omitempty" xml:"NetICMPNum,omitempty"`
	NetICMPSize          *int64  `json:"NetICMPSize,omitempty" xml:"NetICMPSize,omitempty"`
	NetICMPSwitch        *int64  `json:"NetICMPSwitch,omitempty" xml:"NetICMPSwitch,omitempty"`
	NetTraceRouteNum     *int64  `json:"NetTraceRouteNum,omitempty" xml:"NetTraceRouteNum,omitempty"`
	NetTraceRouteSwitch  *int64  `json:"NetTraceRouteSwitch,omitempty" xml:"NetTraceRouteSwitch,omitempty"`
	NetTraceRouteTimeout *int64  `json:"NetTraceRouteTimeout,omitempty" xml:"NetTraceRouteTimeout,omitempty"`
	WhiteList            *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s CreateSyntheticTaskRequestNet) String() string {
	return tea.Prettify(s)
}

func (s CreateSyntheticTaskRequestNet) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskRequestNet) SetNetDNSNs(v string) *CreateSyntheticTaskRequestNet {
	s.NetDNSNs = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetDNSQueryMethod(v int64) *CreateSyntheticTaskRequestNet {
	s.NetDNSQueryMethod = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetDNSServer(v int64) *CreateSyntheticTaskRequestNet {
	s.NetDNSServer = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetDNSSwitch(v int64) *CreateSyntheticTaskRequestNet {
	s.NetDNSSwitch = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetDNSTimeout(v int64) *CreateSyntheticTaskRequestNet {
	s.NetDNSTimeout = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetDigSwitch(v int64) *CreateSyntheticTaskRequestNet {
	s.NetDigSwitch = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetICMPActive(v int64) *CreateSyntheticTaskRequestNet {
	s.NetICMPActive = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetICMPInterval(v int64) *CreateSyntheticTaskRequestNet {
	s.NetICMPInterval = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetICMPNum(v int64) *CreateSyntheticTaskRequestNet {
	s.NetICMPNum = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetICMPSize(v int64) *CreateSyntheticTaskRequestNet {
	s.NetICMPSize = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetICMPSwitch(v int64) *CreateSyntheticTaskRequestNet {
	s.NetICMPSwitch = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetTraceRouteNum(v int64) *CreateSyntheticTaskRequestNet {
	s.NetTraceRouteNum = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetTraceRouteSwitch(v int64) *CreateSyntheticTaskRequestNet {
	s.NetTraceRouteSwitch = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetNetTraceRouteTimeout(v int64) *CreateSyntheticTaskRequestNet {
	s.NetTraceRouteTimeout = &v
	return s
}

func (s *CreateSyntheticTaskRequestNet) SetWhiteList(v string) *CreateSyntheticTaskRequestNet {
	s.WhiteList = &v
	return s
}

type CreateSyntheticTaskShrinkRequest struct {
	CommonParamShrink    *string `json:"CommonParam,omitempty" xml:"CommonParam,omitempty"`
	ExtendIntervalShrink *string `json:"ExtendInterval,omitempty" xml:"ExtendInterval,omitempty"`
	IntervalTime         *string `json:"IntervalTime,omitempty" xml:"IntervalTime,omitempty"`
	IntervalType         *string `json:"IntervalType,omitempty" xml:"IntervalType,omitempty"`
	IpType               *int64  `json:"IpType,omitempty" xml:"IpType,omitempty"`
	MonitorListShrink    *string `json:"MonitorList,omitempty" xml:"MonitorList,omitempty"`
	NetShrink            *string `json:"Net,omitempty" xml:"Net,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskType             *int64  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	Url                  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateSyntheticTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSyntheticTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskShrinkRequest) SetCommonParamShrink(v string) *CreateSyntheticTaskShrinkRequest {
	s.CommonParamShrink = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetExtendIntervalShrink(v string) *CreateSyntheticTaskShrinkRequest {
	s.ExtendIntervalShrink = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetIntervalTime(v string) *CreateSyntheticTaskShrinkRequest {
	s.IntervalTime = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetIntervalType(v string) *CreateSyntheticTaskShrinkRequest {
	s.IntervalType = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetIpType(v int64) *CreateSyntheticTaskShrinkRequest {
	s.IpType = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetMonitorListShrink(v string) *CreateSyntheticTaskShrinkRequest {
	s.MonitorListShrink = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetNetShrink(v string) *CreateSyntheticTaskShrinkRequest {
	s.NetShrink = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetRegionId(v string) *CreateSyntheticTaskShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetTaskName(v string) *CreateSyntheticTaskShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetTaskType(v int64) *CreateSyntheticTaskShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *CreateSyntheticTaskShrinkRequest) SetUrl(v string) *CreateSyntheticTaskShrinkRequest {
	s.Url = &v
	return s
}

type CreateSyntheticTaskResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateSyntheticTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                              `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSyntheticTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSyntheticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskResponseBody) SetCode(v string) *CreateSyntheticTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSyntheticTaskResponseBody) SetData(v *CreateSyntheticTaskResponseBodyData) *CreateSyntheticTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateSyntheticTaskResponseBody) SetMsg(v string) *CreateSyntheticTaskResponseBody {
	s.Msg = &v
	return s
}

func (s *CreateSyntheticTaskResponseBody) SetRequestId(v string) *CreateSyntheticTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateSyntheticTaskResponseBodyData struct {
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSyntheticTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSyntheticTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskResponseBodyData) SetTaskId(v int64) *CreateSyntheticTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type CreateSyntheticTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSyntheticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSyntheticTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSyntheticTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSyntheticTaskResponse) SetHeaders(v map[string]*string) *CreateSyntheticTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSyntheticTaskResponse) SetStatusCode(v int32) *CreateSyntheticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSyntheticTaskResponse) SetBody(v *CreateSyntheticTaskResponseBody) *CreateSyntheticTaskResponse {
	s.Body = v
	return s
}

type CreateWebhookRequest struct {
	Body        *string `json:"Body,omitempty" xml:"Body,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	HttpHeaders *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	HttpParams  *string `json:"HttpParams,omitempty" xml:"HttpParams,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	RecoverBody *string `json:"RecoverBody,omitempty" xml:"RecoverBody,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWebhookRequest) GoString() string {
	return s.String()
}

func (s *CreateWebhookRequest) SetBody(v string) *CreateWebhookRequest {
	s.Body = &v
	return s
}

func (s *CreateWebhookRequest) SetContactName(v string) *CreateWebhookRequest {
	s.ContactName = &v
	return s
}

func (s *CreateWebhookRequest) SetHttpHeaders(v string) *CreateWebhookRequest {
	s.HttpHeaders = &v
	return s
}

func (s *CreateWebhookRequest) SetHttpParams(v string) *CreateWebhookRequest {
	s.HttpParams = &v
	return s
}

func (s *CreateWebhookRequest) SetMethod(v string) *CreateWebhookRequest {
	s.Method = &v
	return s
}

func (s *CreateWebhookRequest) SetRecoverBody(v string) *CreateWebhookRequest {
	s.RecoverBody = &v
	return s
}

func (s *CreateWebhookRequest) SetRegionId(v string) *CreateWebhookRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWebhookRequest) SetUrl(v string) *CreateWebhookRequest {
	s.Url = &v
	return s
}

type CreateWebhookResponseBody struct {
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebhookResponseBody) SetContactId(v string) *CreateWebhookResponseBody {
	s.ContactId = &v
	return s
}

func (s *CreateWebhookResponseBody) SetRequestId(v string) *CreateWebhookResponseBody {
	s.RequestId = &v
	return s
}

type CreateWebhookResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWebhookResponse) GoString() string {
	return s.String()
}

func (s *CreateWebhookResponse) SetHeaders(v map[string]*string) *CreateWebhookResponse {
	s.Headers = v
	return s
}

func (s *CreateWebhookResponse) SetStatusCode(v int32) *CreateWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWebhookResponse) SetBody(v *CreateWebhookResponseBody) *CreateWebhookResponse {
	s.Body = v
	return s
}

type DelAuthTokenRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DelAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s DelAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *DelAuthTokenRequest) SetClusterId(v string) *DelAuthTokenRequest {
	s.ClusterId = &v
	return s
}

func (s *DelAuthTokenRequest) SetRegionId(v string) *DelAuthTokenRequest {
	s.RegionId = &v
	return s
}

type DelAuthTokenResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DelAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DelAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DelAuthTokenResponseBody) SetData(v string) *DelAuthTokenResponseBody {
	s.Data = &v
	return s
}

func (s *DelAuthTokenResponseBody) SetRequestId(v string) *DelAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

type DelAuthTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DelAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DelAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DelAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *DelAuthTokenResponse) SetHeaders(v map[string]*string) *DelAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *DelAuthTokenResponse) SetStatusCode(v int32) *DelAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DelAuthTokenResponse) SetBody(v *DelAuthTokenResponseBody) *DelAuthTokenResponse {
	s.Body = v
	return s
}

type DeleteAlertContactRequest struct {
	ContactId *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactRequest) SetContactId(v int64) *DeleteAlertContactRequest {
	s.ContactId = &v
	return s
}

func (s *DeleteAlertContactRequest) SetRegionId(v string) *DeleteAlertContactRequest {
	s.RegionId = &v
	return s
}

type DeleteAlertContactResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlertContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactResponseBody) SetIsSuccess(v bool) *DeleteAlertContactResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteAlertContactResponseBody) SetRequestId(v string) *DeleteAlertContactResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlertContactResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactResponse) SetHeaders(v map[string]*string) *DeleteAlertContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertContactResponse) SetStatusCode(v int32) *DeleteAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertContactResponse) SetBody(v *DeleteAlertContactResponseBody) *DeleteAlertContactResponse {
	s.Body = v
	return s
}

type DeleteAlertContactGroupRequest struct {
	ContactGroupId *int64  `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupRequest) SetContactGroupId(v int64) *DeleteAlertContactGroupRequest {
	s.ContactGroupId = &v
	return s
}

func (s *DeleteAlertContactGroupRequest) SetRegionId(v string) *DeleteAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

type DeleteAlertContactGroupResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlertContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupResponseBody) SetIsSuccess(v bool) *DeleteAlertContactGroupResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteAlertContactGroupResponseBody) SetRequestId(v string) *DeleteAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlertContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupResponse) SetHeaders(v map[string]*string) *DeleteAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertContactGroupResponse) SetStatusCode(v int32) *DeleteAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertContactGroupResponse) SetBody(v *DeleteAlertContactGroupResponseBody) *DeleteAlertContactGroupResponse {
	s.Body = v
	return s
}

type DeleteAlertRuleRequest struct {
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
}

func (s DeleteAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertRuleRequest) SetAlertId(v int64) *DeleteAlertRuleRequest {
	s.AlertId = &v
	return s
}

type DeleteAlertRuleResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlertRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertRuleResponseBody) SetIsSuccess(v bool) *DeleteAlertRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteAlertRuleResponseBody) SetRequestId(v string) *DeleteAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAlertRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlertRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertRuleResponse) SetHeaders(v map[string]*string) *DeleteAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertRuleResponse) SetStatusCode(v int32) *DeleteAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertRuleResponse) SetBody(v *DeleteAlertRuleResponseBody) *DeleteAlertRuleResponse {
	s.Body = v
	return s
}

type DeleteAlertRulesRequest struct {
	AlertIds *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertRulesRequest) SetAlertIds(v string) *DeleteAlertRulesRequest {
	s.AlertIds = &v
	return s
}

func (s *DeleteAlertRulesRequest) SetRegionId(v string) *DeleteAlertRulesRequest {
	s.RegionId = &v
	return s
}

type DeleteAlertRulesResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertRulesResponseBody) SetIsSuccess(v bool) *DeleteAlertRulesResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteAlertRulesResponseBody) SetRequestId(v string) *DeleteAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAlertRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlertRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertRulesResponse) SetHeaders(v map[string]*string) *DeleteAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertRulesResponse) SetStatusCode(v int32) *DeleteAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertRulesResponse) SetBody(v *DeleteAlertRulesResponseBody) *DeleteAlertRulesResponse {
	s.Body = v
	return s
}

type DeleteCmsExporterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCmsExporterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCmsExporterRequest) GoString() string {
	return s.String()
}

func (s *DeleteCmsExporterRequest) SetClusterId(v string) *DeleteCmsExporterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteCmsExporterRequest) SetRegionId(v string) *DeleteCmsExporterRequest {
	s.RegionId = &v
	return s
}

type DeleteCmsExporterResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCmsExporterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCmsExporterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCmsExporterResponseBody) SetData(v string) *DeleteCmsExporterResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCmsExporterResponseBody) SetRequestId(v string) *DeleteCmsExporterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCmsExporterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCmsExporterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCmsExporterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCmsExporterResponse) GoString() string {
	return s.String()
}

func (s *DeleteCmsExporterResponse) SetHeaders(v map[string]*string) *DeleteCmsExporterResponse {
	s.Headers = v
	return s
}

func (s *DeleteCmsExporterResponse) SetStatusCode(v int32) *DeleteCmsExporterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCmsExporterResponse) SetBody(v *DeleteCmsExporterResponseBody) *DeleteCmsExporterResponse {
	s.Body = v
	return s
}

type DeleteContactRequest struct {
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s DeleteContactRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactRequest) SetContactId(v int64) *DeleteContactRequest {
	s.ContactId = &v
	return s
}

type DeleteContactResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactResponseBody) SetIsSuccess(v bool) *DeleteContactResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteContactResponseBody) SetRequestId(v string) *DeleteContactResponseBody {
	s.RequestId = &v
	return s
}

type DeleteContactResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContactResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactResponse) SetHeaders(v map[string]*string) *DeleteContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactResponse) SetStatusCode(v int32) *DeleteContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactResponse) SetBody(v *DeleteContactResponseBody) *DeleteContactResponse {
	s.Body = v
	return s
}

type DeleteContactGroupRequest struct {
	ContactGroupId *int64 `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
}

func (s DeleteContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupRequest) SetContactGroupId(v int64) *DeleteContactGroupRequest {
	s.ContactGroupId = &v
	return s
}

type DeleteContactGroupResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupResponseBody) SetIsSuccess(v bool) *DeleteContactGroupResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteContactGroupResponseBody) SetRequestId(v string) *DeleteContactGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteContactGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContactGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactGroupResponse) SetHeaders(v map[string]*string) *DeleteContactGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactGroupResponse) SetStatusCode(v int32) *DeleteContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactGroupResponse) SetBody(v *DeleteContactGroupResponseBody) *DeleteContactGroupResponse {
	s.Body = v
	return s
}

type DeleteDispatchRuleRequest struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDispatchRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDispatchRuleRequest) SetId(v string) *DeleteDispatchRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteDispatchRuleRequest) SetRegionId(v string) *DeleteDispatchRuleRequest {
	s.RegionId = &v
	return s
}

type DeleteDispatchRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDispatchRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDispatchRuleResponseBody) SetRequestId(v string) *DeleteDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDispatchRuleResponseBody) SetSuccess(v bool) *DeleteDispatchRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteDispatchRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDispatchRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDispatchRuleResponse) SetHeaders(v map[string]*string) *DeleteDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDispatchRuleResponse) SetStatusCode(v int32) *DeleteDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDispatchRuleResponse) SetBody(v *DeleteDispatchRuleResponseBody) *DeleteDispatchRuleResponse {
	s.Body = v
	return s
}

type DeleteEventBridgeIntegrationRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteEventBridgeIntegrationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventBridgeIntegrationRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventBridgeIntegrationRequest) SetId(v int64) *DeleteEventBridgeIntegrationRequest {
	s.Id = &v
	return s
}

type DeleteEventBridgeIntegrationResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEventBridgeIntegrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventBridgeIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventBridgeIntegrationResponseBody) SetIsSuccess(v bool) *DeleteEventBridgeIntegrationResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteEventBridgeIntegrationResponseBody) SetRequestId(v string) *DeleteEventBridgeIntegrationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEventBridgeIntegrationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEventBridgeIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEventBridgeIntegrationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventBridgeIntegrationResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventBridgeIntegrationResponse) SetHeaders(v map[string]*string) *DeleteEventBridgeIntegrationResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventBridgeIntegrationResponse) SetStatusCode(v int32) *DeleteEventBridgeIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventBridgeIntegrationResponse) SetBody(v *DeleteEventBridgeIntegrationResponseBody) *DeleteEventBridgeIntegrationResponse {
	s.Body = v
	return s
}

type DeleteGrafanaResourceRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGrafanaResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGrafanaResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaResourceRequest) SetClusterId(v string) *DeleteGrafanaResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteGrafanaResourceRequest) SetClusterName(v string) *DeleteGrafanaResourceRequest {
	s.ClusterName = &v
	return s
}

func (s *DeleteGrafanaResourceRequest) SetRegionId(v string) *DeleteGrafanaResourceRequest {
	s.RegionId = &v
	return s
}

type DeleteGrafanaResourceResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGrafanaResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGrafanaResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaResourceResponseBody) SetData(v string) *DeleteGrafanaResourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGrafanaResourceResponseBody) SetRequestId(v string) *DeleteGrafanaResourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGrafanaResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGrafanaResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGrafanaResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGrafanaResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaResourceResponse) SetHeaders(v map[string]*string) *DeleteGrafanaResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGrafanaResourceResponse) SetStatusCode(v int32) *DeleteGrafanaResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGrafanaResourceResponse) SetBody(v *DeleteGrafanaResourceResponseBody) *DeleteGrafanaResourceResponse {
	s.Body = v
	return s
}

type DeleteIMRobotRequest struct {
	RobotId *int64 `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
}

func (s DeleteIMRobotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIMRobotRequest) GoString() string {
	return s.String()
}

func (s *DeleteIMRobotRequest) SetRobotId(v int64) *DeleteIMRobotRequest {
	s.RobotId = &v
	return s
}

type DeleteIMRobotResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIMRobotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIMRobotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIMRobotResponseBody) SetIsSuccess(v bool) *DeleteIMRobotResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteIMRobotResponseBody) SetRequestId(v string) *DeleteIMRobotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIMRobotResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIMRobotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIMRobotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIMRobotResponse) GoString() string {
	return s.String()
}

func (s *DeleteIMRobotResponse) SetHeaders(v map[string]*string) *DeleteIMRobotResponse {
	s.Headers = v
	return s
}

func (s *DeleteIMRobotResponse) SetStatusCode(v int32) *DeleteIMRobotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIMRobotResponse) SetBody(v *DeleteIMRobotResponseBody) *DeleteIMRobotResponse {
	s.Body = v
	return s
}

type DeleteIntegrationRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Integration *string `json:"Integration,omitempty" xml:"Integration,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIntegrationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntegrationRequest) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationRequest) SetClusterId(v string) *DeleteIntegrationRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteIntegrationRequest) SetIntegration(v string) *DeleteIntegrationRequest {
	s.Integration = &v
	return s
}

func (s *DeleteIntegrationRequest) SetRegionId(v string) *DeleteIntegrationRequest {
	s.RegionId = &v
	return s
}

type DeleteIntegrationResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIntegrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationResponseBody) SetData(v string) *DeleteIntegrationResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteIntegrationResponseBody) SetRequestId(v string) *DeleteIntegrationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIntegrationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIntegrationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntegrationResponse) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationResponse) SetHeaders(v map[string]*string) *DeleteIntegrationResponse {
	s.Headers = v
	return s
}

func (s *DeleteIntegrationResponse) SetStatusCode(v int32) *DeleteIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIntegrationResponse) SetBody(v *DeleteIntegrationResponseBody) *DeleteIntegrationResponse {
	s.Body = v
	return s
}

type DeleteIntegrationsRequest struct {
	IntegrationId *int64 `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
}

func (s DeleteIntegrationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntegrationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationsRequest) SetIntegrationId(v int64) *DeleteIntegrationsRequest {
	s.IntegrationId = &v
	return s
}

type DeleteIntegrationsResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIntegrationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntegrationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationsResponseBody) SetIsSuccess(v bool) *DeleteIntegrationsResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteIntegrationsResponseBody) SetRequestId(v string) *DeleteIntegrationsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIntegrationsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIntegrationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIntegrationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIntegrationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationsResponse) SetHeaders(v map[string]*string) *DeleteIntegrationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteIntegrationsResponse) SetStatusCode(v int32) *DeleteIntegrationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIntegrationsResponse) SetBody(v *DeleteIntegrationsResponseBody) *DeleteIntegrationsResponse {
	s.Body = v
	return s
}

type DeleteNotificationPolicyRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteNotificationPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNotificationPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteNotificationPolicyRequest) SetId(v int64) *DeleteNotificationPolicyRequest {
	s.Id = &v
	return s
}

type DeleteNotificationPolicyResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNotificationPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNotificationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNotificationPolicyResponseBody) SetIsSuccess(v bool) *DeleteNotificationPolicyResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteNotificationPolicyResponseBody) SetRequestId(v string) *DeleteNotificationPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNotificationPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteNotificationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNotificationPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNotificationPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteNotificationPolicyResponse) SetHeaders(v map[string]*string) *DeleteNotificationPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteNotificationPolicyResponse) SetStatusCode(v int32) *DeleteNotificationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNotificationPolicyResponse) SetBody(v *DeleteNotificationPolicyResponseBody) *DeleteNotificationPolicyResponse {
	s.Body = v
	return s
}

type DeletePrometheusAlertRuleRequest struct {
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
}

func (s DeletePrometheusAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *DeletePrometheusAlertRuleRequest) SetAlertId(v int64) *DeletePrometheusAlertRuleRequest {
	s.AlertId = &v
	return s
}

type DeletePrometheusAlertRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePrometheusAlertRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrometheusAlertRuleResponseBody) SetRequestId(v string) *DeletePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrometheusAlertRuleResponseBody) SetSuccess(v bool) *DeletePrometheusAlertRuleResponseBody {
	s.Success = &v
	return s
}

type DeletePrometheusAlertRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePrometheusAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePrometheusAlertRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePrometheusAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *DeletePrometheusAlertRuleResponse) SetHeaders(v map[string]*string) *DeletePrometheusAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *DeletePrometheusAlertRuleResponse) SetStatusCode(v int32) *DeletePrometheusAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrometheusAlertRuleResponse) SetBody(v *DeletePrometheusAlertRuleResponseBody) *DeletePrometheusAlertRuleResponse {
	s.Body = v
	return s
}

type DeletePrometheusGlobalViewRequest struct {
	GlobalViewClusterId *string `json:"GlobalViewClusterId,omitempty" xml:"GlobalViewClusterId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePrometheusGlobalViewRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *DeletePrometheusGlobalViewRequest) SetGlobalViewClusterId(v string) *DeletePrometheusGlobalViewRequest {
	s.GlobalViewClusterId = &v
	return s
}

func (s *DeletePrometheusGlobalViewRequest) SetRegionId(v string) *DeletePrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

type DeletePrometheusGlobalViewResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrometheusGlobalViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrometheusGlobalViewResponseBody) SetData(v string) *DeletePrometheusGlobalViewResponseBody {
	s.Data = &v
	return s
}

func (s *DeletePrometheusGlobalViewResponseBody) SetRequestId(v string) *DeletePrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

type DeletePrometheusGlobalViewResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePrometheusGlobalViewResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *DeletePrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *DeletePrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *DeletePrometheusGlobalViewResponse) SetStatusCode(v int32) *DeletePrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrometheusGlobalViewResponse) SetBody(v *DeletePrometheusGlobalViewResponseBody) *DeletePrometheusGlobalViewResponse {
	s.Body = v
	return s
}

type DeleteRetcodeAppRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRetcodeAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRetcodeAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteRetcodeAppRequest) SetAppId(v string) *DeleteRetcodeAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteRetcodeAppRequest) SetRegionId(v string) *DeleteRetcodeAppRequest {
	s.RegionId = &v
	return s
}

type DeleteRetcodeAppResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRetcodeAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRetcodeAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRetcodeAppResponseBody) SetData(v string) *DeleteRetcodeAppResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteRetcodeAppResponseBody) SetRequestId(v string) *DeleteRetcodeAppResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRetcodeAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRetcodeAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRetcodeAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRetcodeAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteRetcodeAppResponse) SetHeaders(v map[string]*string) *DeleteRetcodeAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteRetcodeAppResponse) SetStatusCode(v int32) *DeleteRetcodeAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRetcodeAppResponse) SetBody(v *DeleteRetcodeAppResponseBody) *DeleteRetcodeAppResponse {
	s.Body = v
	return s
}

type DeleteScenarioRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScenarioId *int64  `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
}

func (s DeleteScenarioRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScenarioRequest) GoString() string {
	return s.String()
}

func (s *DeleteScenarioRequest) SetRegionId(v string) *DeleteScenarioRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScenarioRequest) SetScenarioId(v int64) *DeleteScenarioRequest {
	s.ScenarioId = &v
	return s
}

type DeleteScenarioResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteScenarioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScenarioResponseBody) SetRequestId(v string) *DeleteScenarioResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScenarioResponseBody) SetResult(v bool) *DeleteScenarioResponseBody {
	s.Result = &v
	return s
}

type DeleteScenarioResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScenarioResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScenarioResponse) GoString() string {
	return s.String()
}

func (s *DeleteScenarioResponse) SetHeaders(v map[string]*string) *DeleteScenarioResponse {
	s.Headers = v
	return s
}

func (s *DeleteScenarioResponse) SetStatusCode(v int32) *DeleteScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScenarioResponse) SetBody(v *DeleteScenarioResponseBody) *DeleteScenarioResponse {
	s.Body = v
	return s
}

type DeleteSilencePolicyRequest struct {
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteSilencePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSilencePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSilencePolicyRequest) SetId(v int64) *DeleteSilencePolicyRequest {
	s.Id = &v
	return s
}

type DeleteSilencePolicyResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSilencePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSilencePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSilencePolicyResponseBody) SetIsSuccess(v bool) *DeleteSilencePolicyResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteSilencePolicyResponseBody) SetRequestId(v string) *DeleteSilencePolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSilencePolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSilencePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSilencePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSilencePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSilencePolicyResponse) SetHeaders(v map[string]*string) *DeleteSilencePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSilencePolicyResponse) SetStatusCode(v int32) *DeleteSilencePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSilencePolicyResponse) SetBody(v *DeleteSilencePolicyResponseBody) *DeleteSilencePolicyResponse {
	s.Body = v
	return s
}

type DeleteSourceMapRequest struct {
	FidList  []*string `json:"FidList,omitempty" xml:"FidList,omitempty" type:"Repeated"`
	Pid      *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteSourceMapRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSourceMapRequest) GoString() string {
	return s.String()
}

func (s *DeleteSourceMapRequest) SetFidList(v []*string) *DeleteSourceMapRequest {
	s.FidList = v
	return s
}

func (s *DeleteSourceMapRequest) SetPid(v string) *DeleteSourceMapRequest {
	s.Pid = &v
	return s
}

func (s *DeleteSourceMapRequest) SetRegionId(v string) *DeleteSourceMapRequest {
	s.RegionId = &v
	return s
}

type DeleteSourceMapShrinkRequest struct {
	FidListShrink *string `json:"FidList,omitempty" xml:"FidList,omitempty"`
	Pid           *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteSourceMapShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSourceMapShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteSourceMapShrinkRequest) SetFidListShrink(v string) *DeleteSourceMapShrinkRequest {
	s.FidListShrink = &v
	return s
}

func (s *DeleteSourceMapShrinkRequest) SetPid(v string) *DeleteSourceMapShrinkRequest {
	s.Pid = &v
	return s
}

func (s *DeleteSourceMapShrinkRequest) SetRegionId(v string) *DeleteSourceMapShrinkRequest {
	s.RegionId = &v
	return s
}

type DeleteSourceMapResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSourceMapResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSourceMapResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSourceMapResponseBody) SetData(v string) *DeleteSourceMapResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteSourceMapResponseBody) SetRequestId(v string) *DeleteSourceMapResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSourceMapResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSourceMapResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSourceMapResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSourceMapResponse) GoString() string {
	return s.String()
}

func (s *DeleteSourceMapResponse) SetHeaders(v map[string]*string) *DeleteSourceMapResponse {
	s.Headers = v
	return s
}

func (s *DeleteSourceMapResponse) SetStatusCode(v int32) *DeleteSourceMapResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSourceMapResponse) SetBody(v *DeleteSourceMapResponseBody) *DeleteSourceMapResponse {
	s.Body = v
	return s
}

type DeleteTraceAppRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Pid      *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteTraceAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTraceAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppRequest) SetAppId(v string) *DeleteTraceAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteTraceAppRequest) SetPid(v string) *DeleteTraceAppRequest {
	s.Pid = &v
	return s
}

func (s *DeleteTraceAppRequest) SetRegionId(v string) *DeleteTraceAppRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTraceAppRequest) SetType(v string) *DeleteTraceAppRequest {
	s.Type = &v
	return s
}

type DeleteTraceAppResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTraceAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTraceAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppResponseBody) SetData(v string) *DeleteTraceAppResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteTraceAppResponseBody) SetRequestId(v string) *DeleteTraceAppResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTraceAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTraceAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTraceAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTraceAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppResponse) SetHeaders(v map[string]*string) *DeleteTraceAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteTraceAppResponse) SetStatusCode(v int32) *DeleteTraceAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTraceAppResponse) SetBody(v *DeleteTraceAppResponseBody) *DeleteTraceAppResponse {
	s.Body = v
	return s
}

type DeleteWebhookContactRequest struct {
	WebhookId *int64 `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
}

func (s DeleteWebhookContactRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebhookContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebhookContactRequest) SetWebhookId(v int64) *DeleteWebhookContactRequest {
	s.WebhookId = &v
	return s
}

type DeleteWebhookContactResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebhookContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebhookContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebhookContactResponseBody) SetIsSuccess(v bool) *DeleteWebhookContactResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteWebhookContactResponseBody) SetRequestId(v string) *DeleteWebhookContactResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWebhookContactResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteWebhookContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWebhookContactResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWebhookContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebhookContactResponse) SetHeaders(v map[string]*string) *DeleteWebhookContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebhookContactResponse) SetStatusCode(v int32) *DeleteWebhookContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebhookContactResponse) SetBody(v *DeleteWebhookContactResponseBody) *DeleteWebhookContactResponse {
	s.Body = v
	return s
}

type DescribeContactGroupsRequest struct {
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	IsDetail         *bool   `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	Page             *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeContactGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupsRequest) SetContactGroupName(v string) *DescribeContactGroupsRequest {
	s.ContactGroupName = &v
	return s
}

func (s *DescribeContactGroupsRequest) SetIsDetail(v bool) *DescribeContactGroupsRequest {
	s.IsDetail = &v
	return s
}

func (s *DescribeContactGroupsRequest) SetPage(v int64) *DescribeContactGroupsRequest {
	s.Page = &v
	return s
}

func (s *DescribeContactGroupsRequest) SetSize(v int64) *DescribeContactGroupsRequest {
	s.Size = &v
	return s
}

type DescribeContactGroupsResponseBody struct {
	PageBean  *DescribeContactGroupsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContactGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupsResponseBody) SetPageBean(v *DescribeContactGroupsResponseBodyPageBean) *DescribeContactGroupsResponseBody {
	s.PageBean = v
	return s
}

func (s *DescribeContactGroupsResponseBody) SetRequestId(v string) *DescribeContactGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeContactGroupsResponseBodyPageBean struct {
	AlertContactGroups []*DescribeContactGroupsResponseBodyPageBeanAlertContactGroups `json:"AlertContactGroups,omitempty" xml:"AlertContactGroups,omitempty" type:"Repeated"`
	Page               *int64                                                         `json:"Page,omitempty" xml:"Page,omitempty"`
	Size               *int64                                                         `json:"Size,omitempty" xml:"Size,omitempty"`
	Total              *int64                                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeContactGroupsResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupsResponseBodyPageBean) SetAlertContactGroups(v []*DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) *DescribeContactGroupsResponseBodyPageBean {
	s.AlertContactGroups = v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBean) SetPage(v int64) *DescribeContactGroupsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBean) SetSize(v int64) *DescribeContactGroupsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBean) SetTotal(v int64) *DescribeContactGroupsResponseBodyPageBean {
	s.Total = &v
	return s
}

type DescribeContactGroupsResponseBodyPageBeanAlertContactGroups struct {
	ContactGroupId   *float32                                                               `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	ContactGroupName *string                                                                `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	Contacts         []*DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
}

func (s DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) SetContactGroupId(v float32) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups {
	s.ContactGroupId = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) SetContactGroupName(v string) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups {
	s.ContactGroupName = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) SetContacts(v []*DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups {
	s.Contacts = v
	return s
}

type DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts struct {
	ContactId   *float32 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string  `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Email       *string  `json:"Email,omitempty" xml:"Email,omitempty"`
	Phone       *string  `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) SetContactId(v float32) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts {
	s.ContactId = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) SetContactName(v string) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts {
	s.ContactName = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) SetEmail(v string) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts {
	s.Email = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) SetPhone(v string) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts {
	s.Phone = &v
	return s
}

type DescribeContactGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContactGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContactGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupsResponse) SetHeaders(v map[string]*string) *DescribeContactGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContactGroupsResponse) SetStatusCode(v int32) *DescribeContactGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContactGroupsResponse) SetBody(v *DescribeContactGroupsResponseBody) *DescribeContactGroupsResponse {
	s.Body = v
	return s
}

type DescribeContactsRequest struct {
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Page        *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContactsRequest) SetContactName(v string) *DescribeContactsRequest {
	s.ContactName = &v
	return s
}

func (s *DescribeContactsRequest) SetEmail(v string) *DescribeContactsRequest {
	s.Email = &v
	return s
}

func (s *DescribeContactsRequest) SetPage(v int64) *DescribeContactsRequest {
	s.Page = &v
	return s
}

func (s *DescribeContactsRequest) SetPhone(v string) *DescribeContactsRequest {
	s.Phone = &v
	return s
}

func (s *DescribeContactsRequest) SetSize(v int64) *DescribeContactsRequest {
	s.Size = &v
	return s
}

type DescribeContactsResponseBody struct {
	PageBean  *DescribeContactsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContactsResponseBody) SetPageBean(v *DescribeContactsResponseBodyPageBean) *DescribeContactsResponseBody {
	s.PageBean = v
	return s
}

func (s *DescribeContactsResponseBody) SetRequestId(v string) *DescribeContactsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeContactsResponseBodyPageBean struct {
	AlertContacts []*DescribeContactsResponseBodyPageBeanAlertContacts `json:"AlertContacts,omitempty" xml:"AlertContacts,omitempty" type:"Repeated"`
	Page          *int64                                               `json:"Page,omitempty" xml:"Page,omitempty"`
	Size          *int64                                               `json:"Size,omitempty" xml:"Size,omitempty"`
	Total         *int64                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeContactsResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *DescribeContactsResponseBodyPageBean) SetAlertContacts(v []*DescribeContactsResponseBodyPageBeanAlertContacts) *DescribeContactsResponseBodyPageBean {
	s.AlertContacts = v
	return s
}

func (s *DescribeContactsResponseBodyPageBean) SetPage(v int64) *DescribeContactsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBean) SetSize(v int64) *DescribeContactsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBean) SetTotal(v int64) *DescribeContactsResponseBodyPageBean {
	s.Total = &v
	return s
}

type DescribeContactsResponseBodyPageBeanAlertContacts struct {
	ContactId         *float32 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName       *string  `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Email             *string  `json:"Email,omitempty" xml:"Email,omitempty"`
	IsVerify          *bool    `json:"IsVerify,omitempty" xml:"IsVerify,omitempty"`
	Phone             *string  `json:"Phone,omitempty" xml:"Phone,omitempty"`
	ReissueSendNotice *int64   `json:"ReissueSendNotice,omitempty" xml:"ReissueSendNotice,omitempty"`
}

func (s DescribeContactsResponseBodyPageBeanAlertContacts) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactsResponseBodyPageBeanAlertContacts) GoString() string {
	return s.String()
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetContactId(v float32) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.ContactId = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetContactName(v string) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.ContactName = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetEmail(v string) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.Email = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetIsVerify(v bool) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.IsVerify = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetPhone(v string) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.Phone = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetReissueSendNotice(v int64) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.ReissueSendNotice = &v
	return s
}

type DescribeContactsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeContactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContactsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContactsResponse) SetHeaders(v map[string]*string) *DescribeContactsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContactsResponse) SetStatusCode(v int32) *DescribeContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContactsResponse) SetBody(v *DescribeContactsResponseBody) *DescribeContactsResponse {
	s.Body = v
	return s
}

type DescribeDispatchRuleRequest struct {
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDispatchRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleRequest) SetId(v string) *DescribeDispatchRuleRequest {
	s.Id = &v
	return s
}

func (s *DescribeDispatchRuleRequest) SetRegionId(v string) *DescribeDispatchRuleRequest {
	s.RegionId = &v
	return s
}

type DescribeDispatchRuleResponseBody struct {
	DispatchRule *DescribeDispatchRuleResponseBodyDispatchRule `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty" type:"Struct"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDispatchRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBody) SetDispatchRule(v *DescribeDispatchRuleResponseBodyDispatchRule) *DescribeDispatchRuleResponseBody {
	s.DispatchRule = v
	return s
}

func (s *DescribeDispatchRuleResponseBody) SetRequestId(v string) *DescribeDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRule struct {
	DispatchType             *string                                                               `json:"DispatchType,omitempty" xml:"DispatchType,omitempty"`
	GroupRules               []*DescribeDispatchRuleResponseBodyDispatchRuleGroupRules             `json:"GroupRules,omitempty" xml:"GroupRules,omitempty" type:"Repeated"`
	IsRecover                *bool                                                                 `json:"IsRecover,omitempty" xml:"IsRecover,omitempty"`
	LabelMatchExpressionGrid *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid `json:"LabelMatchExpressionGrid,omitempty" xml:"LabelMatchExpressionGrid,omitempty" type:"Struct"`
	Name                     *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	NotifyRules              []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules            `json:"NotifyRules,omitempty" xml:"NotifyRules,omitempty" type:"Repeated"`
	RuleId                   *int64                                                                `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	State                    *string                                                               `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRule) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetDispatchType(v string) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.DispatchType = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetGroupRules(v []*DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.GroupRules = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetIsRecover(v bool) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.IsRecover = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetLabelMatchExpressionGrid(v *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.LabelMatchExpressionGrid = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetName(v string) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.Name = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetNotifyRules(v []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.NotifyRules = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetRuleId(v int64) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.RuleId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetState(v string) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.State = &v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleGroupRules struct {
	GroupId        *int64    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupInterval  *int64    `json:"GroupInterval,omitempty" xml:"GroupInterval,omitempty"`
	GroupWaitTime  *int64    `json:"GroupWaitTime,omitempty" xml:"GroupWaitTime,omitempty"`
	GroupingFields []*string `json:"GroupingFields,omitempty" xml:"GroupingFields,omitempty" type:"Repeated"`
	RepeatInterval *int64    `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupId(v int64) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupInterval(v int64) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupInterval = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupWaitTime(v int64) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupWaitTime = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupingFields(v []*string) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupingFields = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetRepeatInterval(v int64) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.RepeatInterval = &v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid struct {
	LabelMatchExpressionGroups []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups `json:"LabelMatchExpressionGroups,omitempty" xml:"LabelMatchExpressionGroups,omitempty" type:"Repeated"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) SetLabelMatchExpressionGroups(v []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid {
	s.LabelMatchExpressionGroups = v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups struct {
	LabelMatchExpressions []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions `json:"LabelMatchExpressions,omitempty" xml:"LabelMatchExpressions,omitempty" type:"Repeated"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) SetLabelMatchExpressions(v []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups {
	s.LabelMatchExpressions = v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetKey(v string) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Key = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetOperator(v string) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Operator = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetValue(v string) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Value = &v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules struct {
	NotifyChannels []*string                                                               `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	NotifyObjects  []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects `json:"NotifyObjects,omitempty" xml:"NotifyObjects,omitempty" type:"Repeated"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) SetNotifyChannels(v []*string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules {
	s.NotifyChannels = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) SetNotifyObjects(v []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules {
	s.NotifyObjects = v
	return s
}

type DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NotifyObjectId *string `json:"NotifyObjectId,omitempty" xml:"NotifyObjectId,omitempty"`
	NotifyType     *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) SetName(v string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	s.Name = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) SetNotifyObjectId(v string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	s.NotifyObjectId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) SetNotifyType(v string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	s.NotifyType = &v
	return s
}

type DescribeDispatchRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDispatchRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponse) SetHeaders(v map[string]*string) *DescribeDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeDispatchRuleResponse) SetStatusCode(v int32) *DescribeDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDispatchRuleResponse) SetBody(v *DescribeDispatchRuleResponseBody) *DescribeDispatchRuleResponse {
	s.Body = v
	return s
}

type DescribeIMRobotsRequest struct {
	Page      *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	RobotName *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	Size      *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeIMRobotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIMRobotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIMRobotsRequest) SetPage(v int64) *DescribeIMRobotsRequest {
	s.Page = &v
	return s
}

func (s *DescribeIMRobotsRequest) SetRobotName(v string) *DescribeIMRobotsRequest {
	s.RobotName = &v
	return s
}

func (s *DescribeIMRobotsRequest) SetSize(v int64) *DescribeIMRobotsRequest {
	s.Size = &v
	return s
}

type DescribeIMRobotsResponseBody struct {
	PageBean  *DescribeIMRobotsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIMRobotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIMRobotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIMRobotsResponseBody) SetPageBean(v *DescribeIMRobotsResponseBodyPageBean) *DescribeIMRobotsResponseBody {
	s.PageBean = v
	return s
}

func (s *DescribeIMRobotsResponseBody) SetRequestId(v string) *DescribeIMRobotsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeIMRobotsResponseBodyPageBean struct {
	AlertIMRobots []*DescribeIMRobotsResponseBodyPageBeanAlertIMRobots `json:"AlertIMRobots,omitempty" xml:"AlertIMRobots,omitempty" type:"Repeated"`
	Page          *int64                                               `json:"Page,omitempty" xml:"Page,omitempty"`
	Size          *int64                                               `json:"Size,omitempty" xml:"Size,omitempty"`
	Total         *int64                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeIMRobotsResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s DescribeIMRobotsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *DescribeIMRobotsResponseBodyPageBean) SetAlertIMRobots(v []*DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) *DescribeIMRobotsResponseBodyPageBean {
	s.AlertIMRobots = v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBean) SetPage(v int64) *DescribeIMRobotsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBean) SetSize(v int64) *DescribeIMRobotsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBean) SetTotal(v int64) *DescribeIMRobotsResponseBodyPageBean {
	s.Total = &v
	return s
}

type DescribeIMRobotsResponseBodyPageBeanAlertIMRobots struct {
	DailyNoc     *bool    `json:"DailyNoc,omitempty" xml:"DailyNoc,omitempty"`
	DailyNocTime *string  `json:"DailyNocTime,omitempty" xml:"DailyNocTime,omitempty"`
	RobotAddr    *string  `json:"RobotAddr,omitempty" xml:"RobotAddr,omitempty"`
	RobotId      *float32 `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	RobotName    *string  `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	Type         *string  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) String() string {
	return tea.Prettify(s)
}

func (s DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) GoString() string {
	return s.String()
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetDailyNoc(v bool) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.DailyNoc = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetDailyNocTime(v string) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.DailyNocTime = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetRobotAddr(v string) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.RobotAddr = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetRobotId(v float32) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.RobotId = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetRobotName(v string) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.RobotName = &v
	return s
}

func (s *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots) SetType(v string) *DescribeIMRobotsResponseBodyPageBeanAlertIMRobots {
	s.Type = &v
	return s
}

type DescribeIMRobotsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeIMRobotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIMRobotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIMRobotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIMRobotsResponse) SetHeaders(v map[string]*string) *DescribeIMRobotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIMRobotsResponse) SetStatusCode(v int32) *DescribeIMRobotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIMRobotsResponse) SetBody(v *DescribeIMRobotsResponseBody) *DescribeIMRobotsResponse {
	s.Body = v
	return s
}

type DescribePrometheusAlertRuleRequest struct {
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
}

func (s DescribePrometheusAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleRequest) SetAlertId(v int64) *DescribePrometheusAlertRuleRequest {
	s.AlertId = &v
	return s
}

type DescribePrometheusAlertRuleResponseBody struct {
	PrometheusAlertRule *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule `json:"PrometheusAlertRule,omitempty" xml:"PrometheusAlertRule,omitempty" type:"Struct"`
	RequestId           *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBody) SetPrometheusAlertRule(v *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) *DescribePrometheusAlertRuleResponseBody {
	s.PrometheusAlertRule = v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBody) SetRequestId(v string) *DescribePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

type DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule struct {
	AlertId        *int64                                                                   `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string                                                                  `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	ClusterId      *string                                                                  `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64                                                                   `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string                                                                  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string                                                                  `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message        *string                                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string                                                                  `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Status         *int32                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) String() string {
	return tea.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertId(v int64) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertName(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertName = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAnnotations(v []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Annotations = v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetClusterId(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.ClusterId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDispatchRuleId(v int64) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.DispatchRuleId = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDuration(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Duration = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetExpression(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Expression = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetLabels(v []*DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Labels = v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetMessage(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Message = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetNotifyType(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.NotifyType = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetStatus(v int32) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Status = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetType(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Type = &v
	return s
}

type DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) String() string {
	return tea.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetName(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetValue(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Value = &v
	return s
}

type DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) String() string {
	return tea.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetName(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetValue(v string) *DescribePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Value = &v
	return s
}

type DescribePrometheusAlertRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePrometheusAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePrometheusAlertRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePrometheusAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribePrometheusAlertRuleResponse) SetHeaders(v map[string]*string) *DescribePrometheusAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribePrometheusAlertRuleResponse) SetStatusCode(v int32) *DescribePrometheusAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrometheusAlertRuleResponse) SetBody(v *DescribePrometheusAlertRuleResponseBody) *DescribePrometheusAlertRuleResponse {
	s.Body = v
	return s
}

type DescribeTraceLicenseKeyRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTraceLicenseKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceLicenseKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeTraceLicenseKeyRequest) SetRegionId(v string) *DescribeTraceLicenseKeyRequest {
	s.RegionId = &v
	return s
}

type DescribeTraceLicenseKeyResponseBody struct {
	LicenseKey *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTraceLicenseKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceLicenseKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTraceLicenseKeyResponseBody) SetLicenseKey(v string) *DescribeTraceLicenseKeyResponseBody {
	s.LicenseKey = &v
	return s
}

func (s *DescribeTraceLicenseKeyResponseBody) SetRequestId(v string) *DescribeTraceLicenseKeyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTraceLicenseKeyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTraceLicenseKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTraceLicenseKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTraceLicenseKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeTraceLicenseKeyResponse) SetHeaders(v map[string]*string) *DescribeTraceLicenseKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeTraceLicenseKeyResponse) SetStatusCode(v int32) *DescribeTraceLicenseKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTraceLicenseKeyResponse) SetBody(v *DescribeTraceLicenseKeyResponseBody) *DescribeTraceLicenseKeyResponse {
	s.Body = v
	return s
}

type DescribeWebhookContactsRequest struct {
	Page        *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	WebhookName *string `json:"WebhookName,omitempty" xml:"WebhookName,omitempty"`
}

func (s DescribeWebhookContactsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebhookContactsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebhookContactsRequest) SetPage(v int64) *DescribeWebhookContactsRequest {
	s.Page = &v
	return s
}

func (s *DescribeWebhookContactsRequest) SetSize(v int64) *DescribeWebhookContactsRequest {
	s.Size = &v
	return s
}

func (s *DescribeWebhookContactsRequest) SetWebhookName(v string) *DescribeWebhookContactsRequest {
	s.WebhookName = &v
	return s
}

type DescribeWebhookContactsResponseBody struct {
	PageBean  *DescribeWebhookContactsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebhookContactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebhookContactsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebhookContactsResponseBody) SetPageBean(v *DescribeWebhookContactsResponseBodyPageBean) *DescribeWebhookContactsResponseBody {
	s.PageBean = v
	return s
}

func (s *DescribeWebhookContactsResponseBody) SetRequestId(v string) *DescribeWebhookContactsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeWebhookContactsResponseBodyPageBean struct {
	Page            *int64                                                        `json:"Page,omitempty" xml:"Page,omitempty"`
	Size            *int64                                                        `json:"Size,omitempty" xml:"Size,omitempty"`
	Total           *int64                                                        `json:"Total,omitempty" xml:"Total,omitempty"`
	WebhookContacts []*DescribeWebhookContactsResponseBodyPageBeanWebhookContacts `json:"WebhookContacts,omitempty" xml:"WebhookContacts,omitempty" type:"Repeated"`
}

func (s DescribeWebhookContactsResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebhookContactsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *DescribeWebhookContactsResponseBodyPageBean) SetPage(v int64) *DescribeWebhookContactsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBean) SetSize(v int64) *DescribeWebhookContactsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBean) SetTotal(v int64) *DescribeWebhookContactsResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBean) SetWebhookContacts(v []*DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) *DescribeWebhookContactsResponseBodyPageBean {
	s.WebhookContacts = v
	return s
}

type DescribeWebhookContactsResponseBodyPageBeanWebhookContacts struct {
	Webhook     *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook `json:"Webhook,omitempty" xml:"Webhook,omitempty" type:"Struct"`
	WebhookId   *float32                                                           `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
	WebhookName *string                                                            `json:"WebhookName,omitempty" xml:"WebhookName,omitempty"`
}

func (s DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) GoString() string {
	return s.String()
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) SetWebhook(v *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts {
	s.Webhook = v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) SetWebhookId(v float32) *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts {
	s.WebhookId = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts) SetWebhookName(v string) *DescribeWebhookContactsResponseBodyPageBeanWebhookContacts {
	s.WebhookName = &v
	return s
}

type DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook struct {
	BizHeaders  map[string]interface{} `json:"BizHeaders,omitempty" xml:"BizHeaders,omitempty"`
	BizParams   map[string]interface{} `json:"BizParams,omitempty" xml:"BizParams,omitempty"`
	Body        *string                `json:"Body,omitempty" xml:"Body,omitempty"`
	Method      *string                `json:"Method,omitempty" xml:"Method,omitempty"`
	RecoverBody *string                `json:"RecoverBody,omitempty" xml:"RecoverBody,omitempty"`
	Url         *string                `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) GoString() string {
	return s.String()
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) SetBizHeaders(v map[string]interface{}) *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	s.BizHeaders = v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) SetBizParams(v map[string]interface{}) *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	s.BizParams = v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) SetBody(v string) *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	s.Body = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) SetMethod(v string) *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	s.Method = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) SetRecoverBody(v string) *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	s.RecoverBody = &v
	return s
}

func (s *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook) SetUrl(v string) *DescribeWebhookContactsResponseBodyPageBeanWebhookContactsWebhook {
	s.Url = &v
	return s
}

type DescribeWebhookContactsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeWebhookContactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeWebhookContactsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeWebhookContactsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebhookContactsResponse) SetHeaders(v map[string]*string) *DescribeWebhookContactsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebhookContactsResponse) SetStatusCode(v int32) *DescribeWebhookContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebhookContactsResponse) SetBody(v *DescribeWebhookContactsResponseBody) *DescribeWebhookContactsResponse {
	s.Body = v
	return s
}

type GetAgentDownloadUrlRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAgentDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlRequest) SetRegionId(v string) *GetAgentDownloadUrlRequest {
	s.RegionId = &v
	return s
}

type GetAgentDownloadUrlResponseBody struct {
	ArmsAgentDownloadUrl *string `json:"ArmsAgentDownloadUrl,omitempty" xml:"ArmsAgentDownloadUrl,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlResponseBody) SetArmsAgentDownloadUrl(v string) *GetAgentDownloadUrlResponseBody {
	s.ArmsAgentDownloadUrl = &v
	return s
}

func (s *GetAgentDownloadUrlResponseBody) SetRequestId(v string) *GetAgentDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetAgentDownloadUrlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAgentDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAgentDownloadUrlResponse) SetHeaders(v map[string]*string) *GetAgentDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAgentDownloadUrlResponse) SetStatusCode(v int32) *GetAgentDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentDownloadUrlResponse) SetBody(v *GetAgentDownloadUrlResponseBody) *GetAgentDownloadUrlResponse {
	s.Body = v
	return s
}

type GetAlertRulesRequest struct {
	AlertIds    *string `json:"AlertIds,omitempty" xml:"AlertIds,omitempty"`
	AlertNames  *string `json:"AlertNames,omitempty" xml:"AlertNames,omitempty"`
	AlertStatus *string `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	AlertType   *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	Page        *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Size        *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s GetAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *GetAlertRulesRequest) SetAlertIds(v string) *GetAlertRulesRequest {
	s.AlertIds = &v
	return s
}

func (s *GetAlertRulesRequest) SetAlertNames(v string) *GetAlertRulesRequest {
	s.AlertNames = &v
	return s
}

func (s *GetAlertRulesRequest) SetAlertStatus(v string) *GetAlertRulesRequest {
	s.AlertStatus = &v
	return s
}

func (s *GetAlertRulesRequest) SetAlertType(v string) *GetAlertRulesRequest {
	s.AlertType = &v
	return s
}

func (s *GetAlertRulesRequest) SetPage(v int64) *GetAlertRulesRequest {
	s.Page = &v
	return s
}

func (s *GetAlertRulesRequest) SetRegionId(v string) *GetAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *GetAlertRulesRequest) SetSize(v int64) *GetAlertRulesRequest {
	s.Size = &v
	return s
}

type GetAlertRulesResponseBody struct {
	PageBean  *GetAlertRulesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBody) SetPageBean(v *GetAlertRulesResponseBodyPageBean) *GetAlertRulesResponseBody {
	s.PageBean = v
	return s
}

func (s *GetAlertRulesResponseBody) SetRequestId(v string) *GetAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

type GetAlertRulesResponseBodyPageBean struct {
	AlertRules []*GetAlertRulesResponseBodyPageBeanAlertRules `json:"AlertRules,omitempty" xml:"AlertRules,omitempty" type:"Repeated"`
	Page       *int64                                         `json:"Page,omitempty" xml:"Page,omitempty"`
	Size       *int64                                         `json:"Size,omitempty" xml:"Size,omitempty"`
	Total      *int64                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAlertRulesResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBean) SetAlertRules(v []*GetAlertRulesResponseBodyPageBeanAlertRules) *GetAlertRulesResponseBodyPageBean {
	s.AlertRules = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBean) SetPage(v int64) *GetAlertRulesResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBean) SetSize(v int64) *GetAlertRulesResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBean) SetTotal(v int64) *GetAlertRulesResponseBodyPageBean {
	s.Total = &v
	return s
}

type GetAlertRulesResponseBodyPageBeanAlertRules struct {
	AlertCheckType        *string                                                      `json:"AlertCheckType,omitempty" xml:"AlertCheckType,omitempty"`
	AlertGroup            *int64                                                       `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	AlertId               *float32                                                     `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName             *string                                                      `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertRuleContent      *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent `json:"AlertRuleContent,omitempty" xml:"AlertRuleContent,omitempty" type:"Struct"`
	AlertStatus           *string                                                      `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	AlertType             *string                                                      `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	Annotations           []*GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations    `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	AutoAddNewApplication *bool                                                        `json:"AutoAddNewApplication,omitempty" xml:"AutoAddNewApplication,omitempty"`
	ClusterId             *string                                                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreatedTime           *int64                                                       `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Duration              *string                                                      `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Extend                *string                                                      `json:"Extend,omitempty" xml:"Extend,omitempty"`
	Filters               *GetAlertRulesResponseBodyPageBeanAlertRulesFilters          `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	Labels                []*GetAlertRulesResponseBodyPageBeanAlertRulesLabels         `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Level                 *string                                                      `json:"Level,omitempty" xml:"Level,omitempty"`
	Message               *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	MetricsType           *string                                                      `json:"MetricsType,omitempty" xml:"MetricsType,omitempty"`
	NotifyStrategy        *string                                                      `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	Pids                  []*string                                                    `json:"Pids,omitempty" xml:"Pids,omitempty" type:"Repeated"`
	PromQL                *string                                                      `json:"PromQL,omitempty" xml:"PromQL,omitempty"`
	RegionId              *string                                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UpdatedTime           *int64                                                       `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	UserId                *string                                                      `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAlertRulesResponseBodyPageBeanAlertRules) String() string {
	return tea.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRules) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertCheckType(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertCheckType = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertGroup(v int64) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertGroup = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertId(v float32) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertId = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertName(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertName = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertRuleContent(v *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertRuleContent = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertStatus(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertStatus = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAlertType(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertType = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAnnotations(v []*GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Annotations = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetAutoAddNewApplication(v bool) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.AutoAddNewApplication = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetClusterId(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.ClusterId = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetCreatedTime(v int64) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.CreatedTime = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetDuration(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Duration = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetExtend(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Extend = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetFilters(v *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Filters = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetLabels(v []*GetAlertRulesResponseBodyPageBeanAlertRulesLabels) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Labels = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetLevel(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Level = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetMessage(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Message = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetMetricsType(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.MetricsType = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetNotifyStrategy(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.NotifyStrategy = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetPids(v []*string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.Pids = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetPromQL(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.PromQL = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetRegionId(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.RegionId = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetUpdatedTime(v int64) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.UpdatedTime = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRules) SetUserId(v string) *GetAlertRulesResponseBodyPageBeanAlertRules {
	s.UserId = &v
	return s
}

type GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent struct {
	AlertRuleItems []*GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems `json:"AlertRuleItems,omitempty" xml:"AlertRuleItems,omitempty" type:"Repeated"`
	Condition      *string                                                                      `json:"Condition,omitempty" xml:"Condition,omitempty"`
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) String() string {
	return tea.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) SetAlertRuleItems(v []*GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent {
	s.AlertRuleItems = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent) SetCondition(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContent {
	s.Condition = &v
	return s
}

type GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems struct {
	Aggregate *string  `json:"Aggregate,omitempty" xml:"Aggregate,omitempty"`
	MetricKey *string  `json:"MetricKey,omitempty" xml:"MetricKey,omitempty"`
	N         *float32 `json:"N,omitempty" xml:"N,omitempty"`
	Operator  *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value     *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) String() string {
	return tea.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) SetAggregate(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems {
	s.Aggregate = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) SetMetricKey(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems {
	s.MetricKey = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) SetN(v float32) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems {
	s.N = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) SetOperator(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems {
	s.Operator = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems) SetValue(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAlertRuleContentAlertRuleItems {
	s.Value = &v
	return s
}

type GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) String() string {
	return tea.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) SetName(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations {
	s.Name = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations) SetValue(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesAnnotations {
	s.Value = &v
	return s
}

type GetAlertRulesResponseBodyPageBeanAlertRulesFilters struct {
	CustomSLSFilters           []*GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters `json:"CustomSLSFilters,omitempty" xml:"CustomSLSFilters,omitempty" type:"Repeated"`
	CustomSLSGroupByDimensions []*string                                                             `json:"CustomSLSGroupByDimensions,omitempty" xml:"CustomSLSGroupByDimensions,omitempty" type:"Repeated"`
	CustomSLSWheres            []*string                                                             `json:"CustomSLSWheres,omitempty" xml:"CustomSLSWheres,omitempty" type:"Repeated"`
	DimFilters                 []*GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters       `json:"DimFilters,omitempty" xml:"DimFilters,omitempty" type:"Repeated"`
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesFilters) String() string {
	return tea.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesFilters) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) SetCustomSLSFilters(v []*GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) *GetAlertRulesResponseBodyPageBeanAlertRulesFilters {
	s.CustomSLSFilters = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) SetCustomSLSGroupByDimensions(v []*string) *GetAlertRulesResponseBodyPageBeanAlertRulesFilters {
	s.CustomSLSGroupByDimensions = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) SetCustomSLSWheres(v []*string) *GetAlertRulesResponseBodyPageBeanAlertRulesFilters {
	s.CustomSLSWheres = v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFilters) SetDimFilters(v []*GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) *GetAlertRulesResponseBodyPageBeanAlertRulesFilters {
	s.DimFilters = v
	return s
}

type GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Opt   *string `json:"Opt,omitempty" xml:"Opt,omitempty"`
	Show  *bool   `json:"Show,omitempty" xml:"Show,omitempty"`
	T     *string `json:"T,omitempty" xml:"T,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) String() string {
	return tea.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) SetKey(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters {
	s.Key = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) SetOpt(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters {
	s.Opt = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) SetShow(v bool) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters {
	s.Show = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) SetT(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters {
	s.T = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters) SetValue(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersCustomSLSFilters {
	s.Value = &v
	return s
}

type GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters struct {
	FilterKey    *string   `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	FilterOpt    *string   `json:"FilterOpt,omitempty" xml:"FilterOpt,omitempty"`
	FilterValues []*string `json:"FilterValues,omitempty" xml:"FilterValues,omitempty" type:"Repeated"`
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) String() string {
	return tea.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) SetFilterKey(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters {
	s.FilterKey = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) SetFilterOpt(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters {
	s.FilterOpt = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters) SetFilterValues(v []*string) *GetAlertRulesResponseBodyPageBeanAlertRulesFiltersDimFilters {
	s.FilterValues = v
	return s
}

type GetAlertRulesResponseBodyPageBeanAlertRulesLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesLabels) String() string {
	return tea.Prettify(s)
}

func (s GetAlertRulesResponseBodyPageBeanAlertRulesLabels) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesLabels) SetName(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesLabels {
	s.Name = &v
	return s
}

func (s *GetAlertRulesResponseBodyPageBeanAlertRulesLabels) SetValue(v string) *GetAlertRulesResponseBodyPageBeanAlertRulesLabels {
	s.Value = &v
	return s
}

type GetAlertRulesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlertRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponse) SetHeaders(v map[string]*string) *GetAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *GetAlertRulesResponse) SetStatusCode(v int32) *GetAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlertRulesResponse) SetBody(v *GetAlertRulesResponseBody) *GetAlertRulesResponse {
	s.Body = v
	return s
}

type GetAppApiByPageRequest struct {
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IntervalMills *int32  `json:"IntervalMills,omitempty" xml:"IntervalMills,omitempty"`
	PId           *string `json:"PId,omitempty" xml:"PId,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAppApiByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppApiByPageRequest) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageRequest) SetCurrentPage(v int32) *GetAppApiByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetAppApiByPageRequest) SetEndTime(v int64) *GetAppApiByPageRequest {
	s.EndTime = &v
	return s
}

func (s *GetAppApiByPageRequest) SetIntervalMills(v int32) *GetAppApiByPageRequest {
	s.IntervalMills = &v
	return s
}

func (s *GetAppApiByPageRequest) SetPId(v string) *GetAppApiByPageRequest {
	s.PId = &v
	return s
}

func (s *GetAppApiByPageRequest) SetPageSize(v int32) *GetAppApiByPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetAppApiByPageRequest) SetRegionId(v string) *GetAppApiByPageRequest {
	s.RegionId = &v
	return s
}

func (s *GetAppApiByPageRequest) SetStartTime(v int64) *GetAppApiByPageRequest {
	s.StartTime = &v
	return s
}

type GetAppApiByPageResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAppApiByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAppApiByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppApiByPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageResponseBody) SetCode(v int32) *GetAppApiByPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetData(v *GetAppApiByPageResponseBodyData) *GetAppApiByPageResponseBody {
	s.Data = v
	return s
}

func (s *GetAppApiByPageResponseBody) SetMessage(v string) *GetAppApiByPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetRequestId(v string) *GetAppApiByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetSuccess(v bool) *GetAppApiByPageResponseBody {
	s.Success = &v
	return s
}

type GetAppApiByPageResponseBodyData struct {
	Items    []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	Page     *int32                   `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *string                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAppApiByPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppApiByPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageResponseBodyData) SetItems(v []map[string]interface{}) *GetAppApiByPageResponseBodyData {
	s.Items = v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetPage(v int32) *GetAppApiByPageResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetPageSize(v int32) *GetAppApiByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetTotal(v string) *GetAppApiByPageResponseBodyData {
	s.Total = &v
	return s
}

type GetAppApiByPageResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppApiByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppApiByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppApiByPageResponse) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageResponse) SetHeaders(v map[string]*string) *GetAppApiByPageResponse {
	s.Headers = v
	return s
}

func (s *GetAppApiByPageResponse) SetStatusCode(v int32) *GetAppApiByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppApiByPageResponse) SetBody(v *GetAppApiByPageResponseBody) *GetAppApiByPageResponse {
	s.Body = v
	return s
}

type GetAuthTokenRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAuthTokenRequest) SetClusterId(v string) *GetAuthTokenRequest {
	s.ClusterId = &v
	return s
}

func (s *GetAuthTokenRequest) SetRegionId(v string) *GetAuthTokenRequest {
	s.RegionId = &v
	return s
}

type GetAuthTokenResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBody) SetData(v string) *GetAuthTokenResponseBody {
	s.Data = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetRequestId(v string) *GetAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

type GetAuthTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponse) SetHeaders(v map[string]*string) *GetAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAuthTokenResponse) SetStatusCode(v int32) *GetAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthTokenResponse) SetBody(v *GetAuthTokenResponseBody) *GetAuthTokenResponse {
	s.Body = v
	return s
}

type GetClusterAllUrlRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetClusterAllUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterAllUrlRequest) GoString() string {
	return s.String()
}

func (s *GetClusterAllUrlRequest) SetClusterId(v string) *GetClusterAllUrlRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterAllUrlRequest) SetRegionId(v string) *GetClusterAllUrlRequest {
	s.RegionId = &v
	return s
}

type GetClusterAllUrlResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClusterAllUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterAllUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterAllUrlResponseBody) SetData(v string) *GetClusterAllUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetClusterAllUrlResponseBody) SetRequestId(v string) *GetClusterAllUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetClusterAllUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetClusterAllUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterAllUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterAllUrlResponse) GoString() string {
	return s.String()
}

func (s *GetClusterAllUrlResponse) SetHeaders(v map[string]*string) *GetClusterAllUrlResponse {
	s.Headers = v
	return s
}

func (s *GetClusterAllUrlResponse) SetStatusCode(v int32) *GetClusterAllUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterAllUrlResponse) SetBody(v *GetClusterAllUrlResponseBody) *GetClusterAllUrlResponse {
	s.Body = v
	return s
}

type GetExploreUrlRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetExploreUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExploreUrlRequest) GoString() string {
	return s.String()
}

func (s *GetExploreUrlRequest) SetClusterId(v string) *GetExploreUrlRequest {
	s.ClusterId = &v
	return s
}

func (s *GetExploreUrlRequest) SetExpression(v string) *GetExploreUrlRequest {
	s.Expression = &v
	return s
}

func (s *GetExploreUrlRequest) SetRegionId(v string) *GetExploreUrlRequest {
	s.RegionId = &v
	return s
}

func (s *GetExploreUrlRequest) SetType(v string) *GetExploreUrlRequest {
	s.Type = &v
	return s
}

type GetExploreUrlResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExploreUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExploreUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetExploreUrlResponseBody) SetData(v string) *GetExploreUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetExploreUrlResponseBody) SetRequestId(v string) *GetExploreUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetExploreUrlResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExploreUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExploreUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExploreUrlResponse) GoString() string {
	return s.String()
}

func (s *GetExploreUrlResponse) SetHeaders(v map[string]*string) *GetExploreUrlResponse {
	s.Headers = v
	return s
}

func (s *GetExploreUrlResponse) SetStatusCode(v int32) *GetExploreUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExploreUrlResponse) SetBody(v *GetExploreUrlResponseBody) *GetExploreUrlResponse {
	s.Body = v
	return s
}

type GetIntegrationStateRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Integration *string `json:"Integration,omitempty" xml:"Integration,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetIntegrationStateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIntegrationStateRequest) GoString() string {
	return s.String()
}

func (s *GetIntegrationStateRequest) SetClusterId(v string) *GetIntegrationStateRequest {
	s.ClusterId = &v
	return s
}

func (s *GetIntegrationStateRequest) SetIntegration(v string) *GetIntegrationStateRequest {
	s.Integration = &v
	return s
}

func (s *GetIntegrationStateRequest) SetRegionId(v string) *GetIntegrationStateRequest {
	s.RegionId = &v
	return s
}

type GetIntegrationStateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State     *bool   `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetIntegrationStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIntegrationStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntegrationStateResponseBody) SetRequestId(v string) *GetIntegrationStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntegrationStateResponseBody) SetState(v bool) *GetIntegrationStateResponseBody {
	s.State = &v
	return s
}

type GetIntegrationStateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIntegrationStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIntegrationStateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIntegrationStateResponse) GoString() string {
	return s.String()
}

func (s *GetIntegrationStateResponse) SetHeaders(v map[string]*string) *GetIntegrationStateResponse {
	s.Headers = v
	return s
}

func (s *GetIntegrationStateResponse) SetStatusCode(v int32) *GetIntegrationStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntegrationStateResponse) SetBody(v *GetIntegrationStateResponseBody) *GetIntegrationStateResponse {
	s.Body = v
	return s
}

type GetMultipleTraceRequest struct {
	RegionId *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TraceIDs []*string `json:"TraceIDs,omitempty" xml:"TraceIDs,omitempty" type:"Repeated"`
}

func (s GetMultipleTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceRequest) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceRequest) SetRegionId(v string) *GetMultipleTraceRequest {
	s.RegionId = &v
	return s
}

func (s *GetMultipleTraceRequest) SetTraceIDs(v []*string) *GetMultipleTraceRequest {
	s.TraceIDs = v
	return s
}

type GetMultipleTraceResponseBody struct {
	MultiCallChainInfos []*GetMultipleTraceResponseBodyMultiCallChainInfos `json:"MultiCallChainInfos,omitempty" xml:"MultiCallChainInfos,omitempty" type:"Repeated"`
	RequestId           *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMultipleTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBody) SetMultiCallChainInfos(v []*GetMultipleTraceResponseBodyMultiCallChainInfos) *GetMultipleTraceResponseBody {
	s.MultiCallChainInfos = v
	return s
}

func (s *GetMultipleTraceResponseBody) SetRequestId(v string) *GetMultipleTraceResponseBody {
	s.RequestId = &v
	return s
}

type GetMultipleTraceResponseBodyMultiCallChainInfos struct {
	Spans   []*GetMultipleTraceResponseBodyMultiCallChainInfosSpans `json:"Spans,omitempty" xml:"Spans,omitempty" type:"Repeated"`
	TraceID *string                                                 `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfos) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfos) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) SetSpans(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpans) *GetMultipleTraceResponseBodyMultiCallChainInfos {
	s.Spans = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfos) SetTraceID(v string) *GetMultipleTraceResponseBodyMultiCallChainInfos {
	s.TraceID = &v
	return s
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpans struct {
	Duration      *int64                                                              `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HaveStack     *bool                                                               `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	LogEventList  []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Repeated"`
	OperationName *string                                                             `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ParentSpanId  *string                                                             `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	ResultCode    *string                                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	RpcId         *string                                                             `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	RpcType       *int32                                                              `json:"RpcType,omitempty" xml:"RpcType,omitempty"`
	ServiceIp     *string                                                             `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string                                                             `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanId        *string                                                             `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	TagEntryList  []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	Timestamp     *int64                                                              `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string                                                             `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpans) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpans) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetDuration(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.Duration = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetHaveStack(v bool) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.HaveStack = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetLogEventList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.LogEventList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetOperationName(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.OperationName = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetParentSpanId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ParentSpanId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetResultCode(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ResultCode = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetRpcId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.RpcId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetRpcType(v int32) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.RpcType = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetServiceIp(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ServiceIp = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetServiceName(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.ServiceName = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetSpanId(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.SpanId = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTagEntryList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.TagEntryList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTimestamp(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.Timestamp = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpans) SetTraceID(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpans {
	s.TraceID = &v
	return s
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList struct {
	TagEntryList []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	Timestamp    *int64                                                                          `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) SetTagEntryList(v []*GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList {
	s.TagEntryList = v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList) SetTimestamp(v int64) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventList {
	s.Timestamp = &v
	return s
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) SetKey(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList {
	s.Key = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList) SetValue(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansLogEventListTagEntryList {
	s.Value = &v
	return s
}

type GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) SetKey(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList {
	s.Key = &v
	return s
}

func (s *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList) SetValue(v string) *GetMultipleTraceResponseBodyMultiCallChainInfosSpansTagEntryList {
	s.Value = &v
	return s
}

type GetMultipleTraceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetMultipleTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMultipleTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMultipleTraceResponse) GoString() string {
	return s.String()
}

func (s *GetMultipleTraceResponse) SetHeaders(v map[string]*string) *GetMultipleTraceResponse {
	s.Headers = v
	return s
}

func (s *GetMultipleTraceResponse) SetStatusCode(v int32) *GetMultipleTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultipleTraceResponse) SetBody(v *GetMultipleTraceResponseBody) *GetMultipleTraceResponse {
	s.Body = v
	return s
}

type GetOnCallSchedulesDetailRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetOnCallSchedulesDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOnCallSchedulesDetailRequest) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailRequest) SetEndTime(v string) *GetOnCallSchedulesDetailRequest {
	s.EndTime = &v
	return s
}

func (s *GetOnCallSchedulesDetailRequest) SetId(v int64) *GetOnCallSchedulesDetailRequest {
	s.Id = &v
	return s
}

func (s *GetOnCallSchedulesDetailRequest) SetStartTime(v string) *GetOnCallSchedulesDetailRequest {
	s.StartTime = &v
	return s
}

type GetOnCallSchedulesDetailResponseBody struct {
	Data      *GetOnCallSchedulesDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBody) SetData(v *GetOnCallSchedulesDetailResponseBodyData) *GetOnCallSchedulesDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBody) SetRequestId(v string) *GetOnCallSchedulesDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetOnCallSchedulesDetailResponseBodyData struct {
	AlertRobotId              *int64                                                               `json:"AlertRobotId,omitempty" xml:"AlertRobotId,omitempty"`
	Description               *string                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Id                        *int64                                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                      *string                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	RenderedFinnalEntries     []*GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries     `json:"RenderedFinnalEntries,omitempty" xml:"RenderedFinnalEntries,omitempty" type:"Repeated"`
	RenderedLayerEntries      [][]*GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries    `json:"RenderedLayerEntries,omitempty" xml:"RenderedLayerEntries,omitempty" type:"Repeated"`
	RenderedSubstitudeEntries []*GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries `json:"RenderedSubstitudeEntries,omitempty" xml:"RenderedSubstitudeEntries,omitempty" type:"Repeated"`
	ScheduleLayers            []*GetOnCallSchedulesDetailResponseBodyDataScheduleLayers            `json:"ScheduleLayers,omitempty" xml:"ScheduleLayers,omitempty" type:"Repeated"`
}

func (s GetOnCallSchedulesDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetAlertRobotId(v int64) *GetOnCallSchedulesDetailResponseBodyData {
	s.AlertRobotId = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetDescription(v string) *GetOnCallSchedulesDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetId(v int64) *GetOnCallSchedulesDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetName(v string) *GetOnCallSchedulesDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetRenderedFinnalEntries(v []*GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) *GetOnCallSchedulesDetailResponseBodyData {
	s.RenderedFinnalEntries = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetRenderedLayerEntries(v [][]*GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) *GetOnCallSchedulesDetailResponseBodyData {
	s.RenderedLayerEntries = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetRenderedSubstitudeEntries(v []*GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) *GetOnCallSchedulesDetailResponseBodyData {
	s.RenderedSubstitudeEntries = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetScheduleLayers(v []*GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) *GetOnCallSchedulesDetailResponseBodyData {
	s.ScheduleLayers = v
	return s
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries struct {
	End           *string                                                                     `json:"End,omitempty" xml:"End,omitempty"`
	SimpleContact *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact `json:"SimpleContact,omitempty" xml:"SimpleContact,omitempty" type:"Struct"`
	Start         *string                                                                     `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) String() string {
	return tea.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) SetEnd(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries {
	s.End = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) SetSimpleContact(v *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries {
	s.SimpleContact = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) SetStart(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries {
	s.Start = &v
	return s
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) String() string {
	return tea.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) SetId(v int64) *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact {
	s.Id = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) SetName(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact {
	s.Name = &v
	return s
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries struct {
	Start         *string                                                                    `json:"Start,omitempty" xml:"Start,omitempty"`
	End           *string                                                                    `json:"End,omitempty" xml:"End,omitempty"`
	SimpleContact *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact `json:"SimpleContact,omitempty" xml:"SimpleContact,omitempty" type:"Struct"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) String() string {
	return tea.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) SetStart(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries {
	s.Start = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) SetEnd(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries {
	s.End = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) SetSimpleContact(v *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries {
	s.SimpleContact = v
	return s
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) String() string {
	return tea.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) SetId(v int64) *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact {
	s.Id = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) SetName(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact {
	s.Name = &v
	return s
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries struct {
	End           *string                                                                         `json:"End,omitempty" xml:"End,omitempty"`
	SimpleContact *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact `json:"SimpleContact,omitempty" xml:"SimpleContact,omitempty" type:"Struct"`
	Start         *string                                                                         `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) String() string {
	return tea.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) SetEnd(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries {
	s.End = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) SetSimpleContact(v *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries {
	s.SimpleContact = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) SetStart(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries {
	s.Start = &v
	return s
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) String() string {
	return tea.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) SetId(v int64) *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact {
	s.Id = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) SetName(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact {
	s.Name = &v
	return s
}

type GetOnCallSchedulesDetailResponseBodyDataScheduleLayers struct {
	ContactIds   []*int64                                                              `json:"ContactIds,omitempty" xml:"ContactIds,omitempty" type:"Repeated"`
	Restrictions []*GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions `json:"Restrictions,omitempty" xml:"Restrictions,omitempty" type:"Repeated"`
	RotationType *string                                                               `json:"RotationType,omitempty" xml:"RotationType,omitempty"`
	ShiftLength  *int64                                                                `json:"ShiftLength,omitempty" xml:"ShiftLength,omitempty"`
	StartTime    *string                                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) String() string {
	return tea.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) SetContactIds(v []*int64) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers {
	s.ContactIds = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) SetRestrictions(v []*GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers {
	s.Restrictions = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) SetRotationType(v string) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers {
	s.RotationType = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) SetShiftLength(v int64) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers {
	s.ShiftLength = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) SetStartTime(v string) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers {
	s.StartTime = &v
	return s
}

type GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions struct {
	EndTimeOfDay    *string `json:"EndTimeOfDay,omitempty" xml:"EndTimeOfDay,omitempty"`
	RestrictionType *string `json:"RestrictionType,omitempty" xml:"RestrictionType,omitempty"`
	StartTimeOfDay  *string `json:"StartTimeOfDay,omitempty" xml:"StartTimeOfDay,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) String() string {
	return tea.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) SetEndTimeOfDay(v string) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions {
	s.EndTimeOfDay = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) SetRestrictionType(v string) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions {
	s.RestrictionType = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) SetStartTimeOfDay(v string) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions {
	s.StartTimeOfDay = &v
	return s
}

type GetOnCallSchedulesDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOnCallSchedulesDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOnCallSchedulesDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponse) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponse) SetHeaders(v map[string]*string) *GetOnCallSchedulesDetailResponse {
	s.Headers = v
	return s
}

func (s *GetOnCallSchedulesDetailResponse) SetStatusCode(v int32) *GetOnCallSchedulesDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponse) SetBody(v *GetOnCallSchedulesDetailResponseBody) *GetOnCallSchedulesDetailResponse {
	s.Body = v
	return s
}

type GetPrometheusApiTokenRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPrometheusApiTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrometheusApiTokenRequest) GoString() string {
	return s.String()
}

func (s *GetPrometheusApiTokenRequest) SetRegionId(v string) *GetPrometheusApiTokenRequest {
	s.RegionId = &v
	return s
}

type GetPrometheusApiTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetPrometheusApiTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrometheusApiTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrometheusApiTokenResponseBody) SetRequestId(v string) *GetPrometheusApiTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrometheusApiTokenResponseBody) SetToken(v string) *GetPrometheusApiTokenResponseBody {
	s.Token = &v
	return s
}

type GetPrometheusApiTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPrometheusApiTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPrometheusApiTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrometheusApiTokenResponse) GoString() string {
	return s.String()
}

func (s *GetPrometheusApiTokenResponse) SetHeaders(v map[string]*string) *GetPrometheusApiTokenResponse {
	s.Headers = v
	return s
}

func (s *GetPrometheusApiTokenResponse) SetStatusCode(v int32) *GetPrometheusApiTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrometheusApiTokenResponse) SetBody(v *GetPrometheusApiTokenResponseBody) *GetPrometheusApiTokenResponse {
	s.Body = v
	return s
}

type GetPrometheusGlobalViewRequest struct {
	GlobalViewClusterId *string `json:"GlobalViewClusterId,omitempty" xml:"GlobalViewClusterId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPrometheusGlobalViewRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *GetPrometheusGlobalViewRequest) SetGlobalViewClusterId(v string) *GetPrometheusGlobalViewRequest {
	s.GlobalViewClusterId = &v
	return s
}

func (s *GetPrometheusGlobalViewRequest) SetRegionId(v string) *GetPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

type GetPrometheusGlobalViewResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPrometheusGlobalViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrometheusGlobalViewResponseBody) SetData(v string) *GetPrometheusGlobalViewResponseBody {
	s.Data = &v
	return s
}

func (s *GetPrometheusGlobalViewResponseBody) SetRequestId(v string) *GetPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

type GetPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPrometheusGlobalViewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *GetPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *GetPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *GetPrometheusGlobalViewResponse) SetStatusCode(v int32) *GetPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrometheusGlobalViewResponse) SetBody(v *GetPrometheusGlobalViewResponseBody) *GetPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

type GetRecordingRuleRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetRecordingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRecordingRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRecordingRuleRequest) SetClusterId(v string) *GetRecordingRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *GetRecordingRuleRequest) SetRegionId(v string) *GetRecordingRuleRequest {
	s.RegionId = &v
	return s
}

type GetRecordingRuleResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRecordingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRecordingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRecordingRuleResponseBody) SetData(v string) *GetRecordingRuleResponseBody {
	s.Data = &v
	return s
}

func (s *GetRecordingRuleResponseBody) SetRequestId(v string) *GetRecordingRuleResponseBody {
	s.RequestId = &v
	return s
}

type GetRecordingRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRecordingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRecordingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRecordingRuleResponse) GoString() string {
	return s.String()
}

func (s *GetRecordingRuleResponse) SetHeaders(v map[string]*string) *GetRecordingRuleResponse {
	s.Headers = v
	return s
}

func (s *GetRecordingRuleResponse) SetStatusCode(v int32) *GetRecordingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRecordingRuleResponse) SetBody(v *GetRecordingRuleResponseBody) *GetRecordingRuleResponse {
	s.Body = v
	return s
}

type GetRetcodeShareUrlRequest struct {
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
}

func (s GetRetcodeShareUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRetcodeShareUrlRequest) GoString() string {
	return s.String()
}

func (s *GetRetcodeShareUrlRequest) SetPid(v string) *GetRetcodeShareUrlRequest {
	s.Pid = &v
	return s
}

type GetRetcodeShareUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetRetcodeShareUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRetcodeShareUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetRetcodeShareUrlResponseBody) SetRequestId(v string) *GetRetcodeShareUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRetcodeShareUrlResponseBody) SetUrl(v string) *GetRetcodeShareUrlResponseBody {
	s.Url = &v
	return s
}

type GetRetcodeShareUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRetcodeShareUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRetcodeShareUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRetcodeShareUrlResponse) GoString() string {
	return s.String()
}

func (s *GetRetcodeShareUrlResponse) SetHeaders(v map[string]*string) *GetRetcodeShareUrlResponse {
	s.Headers = v
	return s
}

func (s *GetRetcodeShareUrlResponse) SetStatusCode(v int32) *GetRetcodeShareUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRetcodeShareUrlResponse) SetBody(v *GetRetcodeShareUrlResponseBody) *GetRetcodeShareUrlResponse {
	s.Body = v
	return s
}

type GetSourceMapInfoRequest struct {
	AscendingSequence *bool   `json:"AscendingSequence,omitempty" xml:"AscendingSequence,omitempty"`
	Edition           *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	ID                *string `json:"ID,omitempty" xml:"ID,omitempty"`
	Keyword           *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OrderField        *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSourceMapInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSourceMapInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSourceMapInfoRequest) SetAscendingSequence(v bool) *GetSourceMapInfoRequest {
	s.AscendingSequence = &v
	return s
}

func (s *GetSourceMapInfoRequest) SetEdition(v string) *GetSourceMapInfoRequest {
	s.Edition = &v
	return s
}

func (s *GetSourceMapInfoRequest) SetID(v string) *GetSourceMapInfoRequest {
	s.ID = &v
	return s
}

func (s *GetSourceMapInfoRequest) SetKeyword(v string) *GetSourceMapInfoRequest {
	s.Keyword = &v
	return s
}

func (s *GetSourceMapInfoRequest) SetOrderField(v string) *GetSourceMapInfoRequest {
	s.OrderField = &v
	return s
}

func (s *GetSourceMapInfoRequest) SetRegionId(v string) *GetSourceMapInfoRequest {
	s.RegionId = &v
	return s
}

type GetSourceMapInfoResponseBody struct {
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceMapList []*GetSourceMapInfoResponseBodySourceMapList `json:"SourceMapList,omitempty" xml:"SourceMapList,omitempty" type:"Repeated"`
}

func (s GetSourceMapInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSourceMapInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSourceMapInfoResponseBody) SetRequestId(v string) *GetSourceMapInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSourceMapInfoResponseBody) SetSourceMapList(v []*GetSourceMapInfoResponseBodySourceMapList) *GetSourceMapInfoResponseBody {
	s.SourceMapList = v
	return s
}

type GetSourceMapInfoResponseBodySourceMapList struct {
	Fid        *string `json:"Fid,omitempty" xml:"Fid,omitempty"`
	FileName   *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Size       *string `json:"Size,omitempty" xml:"Size,omitempty"`
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetSourceMapInfoResponseBodySourceMapList) String() string {
	return tea.Prettify(s)
}

func (s GetSourceMapInfoResponseBodySourceMapList) GoString() string {
	return s.String()
}

func (s *GetSourceMapInfoResponseBodySourceMapList) SetFid(v string) *GetSourceMapInfoResponseBodySourceMapList {
	s.Fid = &v
	return s
}

func (s *GetSourceMapInfoResponseBodySourceMapList) SetFileName(v string) *GetSourceMapInfoResponseBodySourceMapList {
	s.FileName = &v
	return s
}

func (s *GetSourceMapInfoResponseBodySourceMapList) SetSize(v string) *GetSourceMapInfoResponseBodySourceMapList {
	s.Size = &v
	return s
}

func (s *GetSourceMapInfoResponseBodySourceMapList) SetUploadTime(v string) *GetSourceMapInfoResponseBodySourceMapList {
	s.UploadTime = &v
	return s
}

func (s *GetSourceMapInfoResponseBodySourceMapList) SetVersion(v string) *GetSourceMapInfoResponseBodySourceMapList {
	s.Version = &v
	return s
}

type GetSourceMapInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSourceMapInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSourceMapInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSourceMapInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSourceMapInfoResponse) SetHeaders(v map[string]*string) *GetSourceMapInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSourceMapInfoResponse) SetStatusCode(v int32) *GetSourceMapInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSourceMapInfoResponse) SetBody(v *GetSourceMapInfoResponseBody) *GetSourceMapInfoResponse {
	s.Body = v
	return s
}

type GetStackRequest struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Pid       *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RpcID     *string `json:"RpcID,omitempty" xml:"RpcID,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TraceID   *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetStackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackRequest) GoString() string {
	return s.String()
}

func (s *GetStackRequest) SetEndTime(v int64) *GetStackRequest {
	s.EndTime = &v
	return s
}

func (s *GetStackRequest) SetPid(v string) *GetStackRequest {
	s.Pid = &v
	return s
}

func (s *GetStackRequest) SetRegionId(v string) *GetStackRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackRequest) SetRpcID(v string) *GetStackRequest {
	s.RpcID = &v
	return s
}

func (s *GetStackRequest) SetStartTime(v int64) *GetStackRequest {
	s.StartTime = &v
	return s
}

func (s *GetStackRequest) SetTraceID(v string) *GetStackRequest {
	s.TraceID = &v
	return s
}

type GetStackResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StackInfo []*GetStackResponseBodyStackInfo `json:"StackInfo,omitempty" xml:"StackInfo,omitempty" type:"Repeated"`
}

func (s GetStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackResponseBody) SetRequestId(v string) *GetStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackResponseBody) SetStackInfo(v []*GetStackResponseBodyStackInfo) *GetStackResponseBody {
	s.StackInfo = v
	return s
}

type GetStackResponseBodyStackInfo struct {
	Api         *string                               `json:"Api,omitempty" xml:"Api,omitempty"`
	Duration    *int64                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Exception   *string                               `json:"Exception,omitempty" xml:"Exception,omitempty"`
	ExtInfo     *GetStackResponseBodyStackInfoExtInfo `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty" type:"Struct"`
	Line        *string                               `json:"Line,omitempty" xml:"Line,omitempty"`
	RpcId       *string                               `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	ServiceName *string                               `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime   *int64                                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetStackResponseBodyStackInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyStackInfo) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyStackInfo) SetApi(v string) *GetStackResponseBodyStackInfo {
	s.Api = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetDuration(v int64) *GetStackResponseBodyStackInfo {
	s.Duration = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetException(v string) *GetStackResponseBodyStackInfo {
	s.Exception = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetExtInfo(v *GetStackResponseBodyStackInfoExtInfo) *GetStackResponseBodyStackInfo {
	s.ExtInfo = v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetLine(v string) *GetStackResponseBodyStackInfo {
	s.Line = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetRpcId(v string) *GetStackResponseBodyStackInfo {
	s.RpcId = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetServiceName(v string) *GetStackResponseBodyStackInfo {
	s.ServiceName = &v
	return s
}

func (s *GetStackResponseBodyStackInfo) SetStartTime(v int64) *GetStackResponseBodyStackInfo {
	s.StartTime = &v
	return s
}

type GetStackResponseBodyStackInfoExtInfo struct {
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStackResponseBodyStackInfoExtInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyStackInfoExtInfo) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyStackInfoExtInfo) SetInfo(v string) *GetStackResponseBodyStackInfoExtInfo {
	s.Info = &v
	return s
}

func (s *GetStackResponseBodyStackInfoExtInfo) SetType(v string) *GetStackResponseBodyStackInfoExtInfo {
	s.Type = &v
	return s
}

type GetStackResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponse) GoString() string {
	return s.String()
}

func (s *GetStackResponse) SetHeaders(v map[string]*string) *GetStackResponse {
	s.Headers = v
	return s
}

func (s *GetStackResponse) SetStatusCode(v int32) *GetStackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackResponse) SetBody(v *GetStackResponseBody) *GetStackResponse {
	s.Body = v
	return s
}

type GetSyntheticTaskMonitorsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSyntheticTaskMonitorsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSyntheticTaskMonitorsRequest) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskMonitorsRequest) SetRegionId(v string) *GetSyntheticTaskMonitorsRequest {
	s.RegionId = &v
	return s
}

type GetSyntheticTaskMonitorsResponseBody struct {
	Code      *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetSyntheticTaskMonitorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Msg       *string                                     `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSyntheticTaskMonitorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSyntheticTaskMonitorsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskMonitorsResponseBody) SetCode(v string) *GetSyntheticTaskMonitorsResponseBody {
	s.Code = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBody) SetData(v []*GetSyntheticTaskMonitorsResponseBodyData) *GetSyntheticTaskMonitorsResponseBody {
	s.Data = v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBody) SetMsg(v string) *GetSyntheticTaskMonitorsResponseBody {
	s.Msg = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBody) SetRequestId(v string) *GetSyntheticTaskMonitorsResponseBody {
	s.RequestId = &v
	return s
}

type GetSyntheticTaskMonitorsResponseBodyData struct {
	Busy           *int64  `json:"Busy,omitempty" xml:"Busy,omitempty"`
	City           *string `json:"City,omitempty" xml:"City,omitempty"`
	CityCode       *int64  `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	ClientType     *int64  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	District       *string `json:"District,omitempty" xml:"District,omitempty"`
	NetServiceId   *int64  `json:"NetServiceId,omitempty" xml:"NetServiceId,omitempty"`
	NetServiceName *string `json:"NetServiceName,omitempty" xml:"NetServiceName,omitempty"`
}

func (s GetSyntheticTaskMonitorsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSyntheticTaskMonitorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetBusy(v int64) *GetSyntheticTaskMonitorsResponseBodyData {
	s.Busy = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetCity(v string) *GetSyntheticTaskMonitorsResponseBodyData {
	s.City = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetCityCode(v int64) *GetSyntheticTaskMonitorsResponseBodyData {
	s.CityCode = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetClientType(v int64) *GetSyntheticTaskMonitorsResponseBodyData {
	s.ClientType = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetDistrict(v string) *GetSyntheticTaskMonitorsResponseBodyData {
	s.District = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetNetServiceId(v int64) *GetSyntheticTaskMonitorsResponseBodyData {
	s.NetServiceId = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponseBodyData) SetNetServiceName(v string) *GetSyntheticTaskMonitorsResponseBodyData {
	s.NetServiceName = &v
	return s
}

type GetSyntheticTaskMonitorsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSyntheticTaskMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSyntheticTaskMonitorsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSyntheticTaskMonitorsResponse) GoString() string {
	return s.String()
}

func (s *GetSyntheticTaskMonitorsResponse) SetHeaders(v map[string]*string) *GetSyntheticTaskMonitorsResponse {
	s.Headers = v
	return s
}

func (s *GetSyntheticTaskMonitorsResponse) SetStatusCode(v int32) *GetSyntheticTaskMonitorsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSyntheticTaskMonitorsResponse) SetBody(v *GetSyntheticTaskMonitorsResponseBody) *GetSyntheticTaskMonitorsResponse {
	s.Body = v
	return s
}

type GetTraceRequest struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TraceID   *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTraceRequest) GoString() string {
	return s.String()
}

func (s *GetTraceRequest) SetEndTime(v int64) *GetTraceRequest {
	s.EndTime = &v
	return s
}

func (s *GetTraceRequest) SetRegionId(v string) *GetTraceRequest {
	s.RegionId = &v
	return s
}

func (s *GetTraceRequest) SetStartTime(v int64) *GetTraceRequest {
	s.StartTime = &v
	return s
}

func (s *GetTraceRequest) SetTraceID(v string) *GetTraceRequest {
	s.TraceID = &v
	return s
}

type GetTraceResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Spans     []*GetTraceResponseBodySpans `json:"Spans,omitempty" xml:"Spans,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBody) SetRequestId(v string) *GetTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceResponseBody) SetSpans(v []*GetTraceResponseBodySpans) *GetTraceResponseBody {
	s.Spans = v
	return s
}

type GetTraceResponseBodySpans struct {
	Children      []map[string]interface{}                 `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	Duration      *int64                                   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HaveStack     *bool                                    `json:"HaveStack,omitempty" xml:"HaveStack,omitempty"`
	LogEventList  []*GetTraceResponseBodySpansLogEventList `json:"LogEventList,omitempty" xml:"LogEventList,omitempty" type:"Repeated"`
	OperationName *string                                  `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ParentSpanId  *string                                  `json:"ParentSpanId,omitempty" xml:"ParentSpanId,omitempty"`
	ResultCode    *string                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	RpcId         *string                                  `json:"RpcId,omitempty" xml:"RpcId,omitempty"`
	RpcType       *int32                                   `json:"RpcType,omitempty" xml:"RpcType,omitempty"`
	ServiceIp     *string                                  `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string                                  `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	SpanId        *string                                  `json:"SpanId,omitempty" xml:"SpanId,omitempty"`
	TagEntryList  []*GetTraceResponseBodySpansTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	Timestamp     *int64                                   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string                                  `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s GetTraceResponseBodySpans) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpans) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpans) SetChildren(v []map[string]interface{}) *GetTraceResponseBodySpans {
	s.Children = v
	return s
}

func (s *GetTraceResponseBodySpans) SetDuration(v int64) *GetTraceResponseBodySpans {
	s.Duration = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetHaveStack(v bool) *GetTraceResponseBodySpans {
	s.HaveStack = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetLogEventList(v []*GetTraceResponseBodySpansLogEventList) *GetTraceResponseBodySpans {
	s.LogEventList = v
	return s
}

func (s *GetTraceResponseBodySpans) SetOperationName(v string) *GetTraceResponseBodySpans {
	s.OperationName = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetParentSpanId(v string) *GetTraceResponseBodySpans {
	s.ParentSpanId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetResultCode(v string) *GetTraceResponseBodySpans {
	s.ResultCode = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetRpcId(v string) *GetTraceResponseBodySpans {
	s.RpcId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetRpcType(v int32) *GetTraceResponseBodySpans {
	s.RpcType = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetServiceIp(v string) *GetTraceResponseBodySpans {
	s.ServiceIp = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetServiceName(v string) *GetTraceResponseBodySpans {
	s.ServiceName = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetSpanId(v string) *GetTraceResponseBodySpans {
	s.SpanId = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetTagEntryList(v []*GetTraceResponseBodySpansTagEntryList) *GetTraceResponseBodySpans {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpans) SetTimestamp(v int64) *GetTraceResponseBodySpans {
	s.Timestamp = &v
	return s
}

func (s *GetTraceResponseBodySpans) SetTraceID(v string) *GetTraceResponseBodySpans {
	s.TraceID = &v
	return s
}

type GetTraceResponseBodySpansLogEventList struct {
	TagEntryList []*GetTraceResponseBodySpansLogEventListTagEntryList `json:"TagEntryList,omitempty" xml:"TagEntryList,omitempty" type:"Repeated"`
	Timestamp    *int64                                               `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetTraceResponseBodySpansLogEventList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansLogEventList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansLogEventList) SetTagEntryList(v []*GetTraceResponseBodySpansLogEventListTagEntryList) *GetTraceResponseBodySpansLogEventList {
	s.TagEntryList = v
	return s
}

func (s *GetTraceResponseBodySpansLogEventList) SetTimestamp(v int64) *GetTraceResponseBodySpansLogEventList {
	s.Timestamp = &v
	return s
}

type GetTraceResponseBodySpansLogEventListTagEntryList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceResponseBodySpansLogEventListTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansLogEventListTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) SetKey(v string) *GetTraceResponseBodySpansLogEventListTagEntryList {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodySpansLogEventListTagEntryList) SetValue(v string) *GetTraceResponseBodySpansLogEventListTagEntryList {
	s.Value = &v
	return s
}

type GetTraceResponseBodySpansTagEntryList struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTraceResponseBodySpansTagEntryList) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodySpansTagEntryList) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodySpansTagEntryList) SetKey(v string) *GetTraceResponseBodySpansTagEntryList {
	s.Key = &v
	return s
}

func (s *GetTraceResponseBodySpansTagEntryList) SetValue(v string) *GetTraceResponseBodySpansTagEntryList {
	s.Value = &v
	return s
}

type GetTraceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTraceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponse) GoString() string {
	return s.String()
}

func (s *GetTraceResponse) SetHeaders(v map[string]*string) *GetTraceResponse {
	s.Headers = v
	return s
}

func (s *GetTraceResponse) SetStatusCode(v int32) *GetTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceResponse) SetBody(v *GetTraceResponseBody) *GetTraceResponse {
	s.Body = v
	return s
}

type GetTraceAppRequest struct {
	Pid      *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetTraceAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAppRequest) GoString() string {
	return s.String()
}

func (s *GetTraceAppRequest) SetPid(v string) *GetTraceAppRequest {
	s.Pid = &v
	return s
}

func (s *GetTraceAppRequest) SetRegionId(v string) *GetTraceAppRequest {
	s.RegionId = &v
	return s
}

type GetTraceAppResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceApp  *GetTraceAppResponseBodyTraceApp `json:"TraceApp,omitempty" xml:"TraceApp,omitempty" type:"Struct"`
}

func (s GetTraceAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponseBody) SetRequestId(v string) *GetTraceAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceAppResponseBody) SetTraceApp(v *GetTraceAppResponseBodyTraceApp) *GetTraceAppResponseBody {
	s.TraceApp = v
	return s
}

type GetTraceAppResponseBodyTraceApp struct {
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetTraceAppResponseBodyTraceApp) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAppResponseBodyTraceApp) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponseBodyTraceApp) SetAppId(v int64) *GetTraceAppResponseBodyTraceApp {
	s.AppId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetAppName(v string) *GetTraceAppResponseBodyTraceApp {
	s.AppName = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetCreateTime(v int64) *GetTraceAppResponseBodyTraceApp {
	s.CreateTime = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetLabels(v []*string) *GetTraceAppResponseBodyTraceApp {
	s.Labels = v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetPid(v string) *GetTraceAppResponseBodyTraceApp {
	s.Pid = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetRegionId(v string) *GetTraceAppResponseBodyTraceApp {
	s.RegionId = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetShow(v bool) *GetTraceAppResponseBodyTraceApp {
	s.Show = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetType(v string) *GetTraceAppResponseBodyTraceApp {
	s.Type = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetUpdateTime(v int64) *GetTraceAppResponseBodyTraceApp {
	s.UpdateTime = &v
	return s
}

func (s *GetTraceAppResponseBodyTraceApp) SetUserId(v string) *GetTraceAppResponseBodyTraceApp {
	s.UserId = &v
	return s
}

type GetTraceAppResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTraceAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTraceAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTraceAppResponse) GoString() string {
	return s.String()
}

func (s *GetTraceAppResponse) SetHeaders(v map[string]*string) *GetTraceAppResponse {
	s.Headers = v
	return s
}

func (s *GetTraceAppResponse) SetStatusCode(v int32) *GetTraceAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceAppResponse) SetBody(v *GetTraceAppResponseBody) *GetTraceAppResponse {
	s.Body = v
	return s
}

type ImportAppAlertRulesRequest struct {
	ContactGroupIds     *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsAutoStart         *bool   `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	Pids                *string `json:"Pids,omitempty" xml:"Pids,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
	TemplateAlertId     *string `json:"TemplateAlertId,omitempty" xml:"TemplateAlertId,omitempty"`
}

func (s ImportAppAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportAppAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesRequest) SetContactGroupIds(v string) *ImportAppAlertRulesRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetIsAutoStart(v bool) *ImportAppAlertRulesRequest {
	s.IsAutoStart = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetPids(v string) *ImportAppAlertRulesRequest {
	s.Pids = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetRegionId(v string) *ImportAppAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetTemplageAlertConfig(v string) *ImportAppAlertRulesRequest {
	s.TemplageAlertConfig = &v
	return s
}

func (s *ImportAppAlertRulesRequest) SetTemplateAlertId(v string) *ImportAppAlertRulesRequest {
	s.TemplateAlertId = &v
	return s
}

type ImportAppAlertRulesResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportAppAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportAppAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesResponseBody) SetData(v string) *ImportAppAlertRulesResponseBody {
	s.Data = &v
	return s
}

func (s *ImportAppAlertRulesResponseBody) SetRequestId(v string) *ImportAppAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

type ImportAppAlertRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImportAppAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportAppAlertRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportAppAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesResponse) SetHeaders(v map[string]*string) *ImportAppAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ImportAppAlertRulesResponse) SetStatusCode(v int32) *ImportAppAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportAppAlertRulesResponse) SetBody(v *ImportAppAlertRulesResponseBody) *ImportAppAlertRulesResponse {
	s.Body = v
	return s
}

type InstallCmsExporterRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CmsArgs    *string `json:"CmsArgs,omitempty" xml:"CmsArgs,omitempty"`
	DirectArgs *string `json:"DirectArgs,omitempty" xml:"DirectArgs,omitempty"`
	EnableTag  *bool   `json:"EnableTag,omitempty" xml:"EnableTag,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallCmsExporterRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallCmsExporterRequest) GoString() string {
	return s.String()
}

func (s *InstallCmsExporterRequest) SetClusterId(v string) *InstallCmsExporterRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallCmsExporterRequest) SetCmsArgs(v string) *InstallCmsExporterRequest {
	s.CmsArgs = &v
	return s
}

func (s *InstallCmsExporterRequest) SetDirectArgs(v string) *InstallCmsExporterRequest {
	s.DirectArgs = &v
	return s
}

func (s *InstallCmsExporterRequest) SetEnableTag(v bool) *InstallCmsExporterRequest {
	s.EnableTag = &v
	return s
}

func (s *InstallCmsExporterRequest) SetRegionId(v string) *InstallCmsExporterRequest {
	s.RegionId = &v
	return s
}

type InstallCmsExporterResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallCmsExporterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallCmsExporterResponseBody) GoString() string {
	return s.String()
}

func (s *InstallCmsExporterResponseBody) SetData(v string) *InstallCmsExporterResponseBody {
	s.Data = &v
	return s
}

func (s *InstallCmsExporterResponseBody) SetRequestId(v string) *InstallCmsExporterResponseBody {
	s.RequestId = &v
	return s
}

type InstallCmsExporterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InstallCmsExporterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallCmsExporterResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallCmsExporterResponse) GoString() string {
	return s.String()
}

func (s *InstallCmsExporterResponse) SetHeaders(v map[string]*string) *InstallCmsExporterResponse {
	s.Headers = v
	return s
}

func (s *InstallCmsExporterResponse) SetStatusCode(v int32) *InstallCmsExporterResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallCmsExporterResponse) SetBody(v *InstallCmsExporterResponseBody) *InstallCmsExporterResponse {
	s.Body = v
	return s
}

type InstallManagedPrometheusRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType     *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	KubeConfig      *string `json:"KubeConfig,omitempty" xml:"KubeConfig,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s InstallManagedPrometheusRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallManagedPrometheusRequest) GoString() string {
	return s.String()
}

func (s *InstallManagedPrometheusRequest) SetClusterId(v string) *InstallManagedPrometheusRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetClusterType(v string) *InstallManagedPrometheusRequest {
	s.ClusterType = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetKubeConfig(v string) *InstallManagedPrometheusRequest {
	s.KubeConfig = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetRegionId(v string) *InstallManagedPrometheusRequest {
	s.RegionId = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetSecurityGroupId(v string) *InstallManagedPrometheusRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetVSwitchId(v string) *InstallManagedPrometheusRequest {
	s.VSwitchId = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetVpcId(v string) *InstallManagedPrometheusRequest {
	s.VpcId = &v
	return s
}

type InstallManagedPrometheusResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InstallManagedPrometheusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallManagedPrometheusResponseBody) GoString() string {
	return s.String()
}

func (s *InstallManagedPrometheusResponseBody) SetCode(v int32) *InstallManagedPrometheusResponseBody {
	s.Code = &v
	return s
}

func (s *InstallManagedPrometheusResponseBody) SetData(v string) *InstallManagedPrometheusResponseBody {
	s.Data = &v
	return s
}

func (s *InstallManagedPrometheusResponseBody) SetMessage(v string) *InstallManagedPrometheusResponseBody {
	s.Message = &v
	return s
}

func (s *InstallManagedPrometheusResponseBody) SetRequestId(v string) *InstallManagedPrometheusResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallManagedPrometheusResponseBody) SetSuccess(v bool) *InstallManagedPrometheusResponseBody {
	s.Success = &v
	return s
}

type InstallManagedPrometheusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InstallManagedPrometheusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallManagedPrometheusResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallManagedPrometheusResponse) GoString() string {
	return s.String()
}

func (s *InstallManagedPrometheusResponse) SetHeaders(v map[string]*string) *InstallManagedPrometheusResponse {
	s.Headers = v
	return s
}

func (s *InstallManagedPrometheusResponse) SetStatusCode(v int32) *InstallManagedPrometheusResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallManagedPrometheusResponse) SetBody(v *InstallManagedPrometheusResponseBody) *InstallManagedPrometheusResponse {
	s.Body = v
	return s
}

type ListActivatedAlertsRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Filter      *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListActivatedAlertsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedAlertsRequest) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsRequest) SetCurrentPage(v int32) *ListActivatedAlertsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListActivatedAlertsRequest) SetFilter(v string) *ListActivatedAlertsRequest {
	s.Filter = &v
	return s
}

func (s *ListActivatedAlertsRequest) SetPageSize(v int32) *ListActivatedAlertsRequest {
	s.PageSize = &v
	return s
}

func (s *ListActivatedAlertsRequest) SetRegionId(v string) *ListActivatedAlertsRequest {
	s.RegionId = &v
	return s
}

type ListActivatedAlertsResponseBody struct {
	Page      *ListActivatedAlertsResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListActivatedAlertsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponseBody) SetPage(v *ListActivatedAlertsResponseBodyPage) *ListActivatedAlertsResponseBody {
	s.Page = v
	return s
}

func (s *ListActivatedAlertsResponseBody) SetRequestId(v string) *ListActivatedAlertsResponseBody {
	s.RequestId = &v
	return s
}

type ListActivatedAlertsResponseBodyPage struct {
	Alerts   []*ListActivatedAlertsResponseBodyPageAlerts `json:"Alerts,omitempty" xml:"Alerts,omitempty" type:"Repeated"`
	Page     *int32                                       `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListActivatedAlertsResponseBodyPage) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedAlertsResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponseBodyPage) SetAlerts(v []*ListActivatedAlertsResponseBodyPageAlerts) *ListActivatedAlertsResponseBodyPage {
	s.Alerts = v
	return s
}

func (s *ListActivatedAlertsResponseBodyPage) SetPage(v int32) *ListActivatedAlertsResponseBodyPage {
	s.Page = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPage) SetPageSize(v int32) *ListActivatedAlertsResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPage) SetTotal(v int32) *ListActivatedAlertsResponseBodyPage {
	s.Total = &v
	return s
}

type ListActivatedAlertsResponseBodyPageAlerts struct {
	AlertId            *string                                                   `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName          *string                                                   `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertType          *string                                                   `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	Count              *int32                                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	CreateTime         *int64                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DispatchRules      []*ListActivatedAlertsResponseBodyPageAlertsDispatchRules `json:"DispatchRules,omitempty" xml:"DispatchRules,omitempty" type:"Repeated"`
	EndsAt             *int64                                                    `json:"EndsAt,omitempty" xml:"EndsAt,omitempty"`
	ExpandFields       map[string]interface{}                                    `json:"ExpandFields,omitempty" xml:"ExpandFields,omitempty"`
	IntegrationName    *string                                                   `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	IntegrationType    *string                                                   `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	InvolvedObjectKind *string                                                   `json:"InvolvedObjectKind,omitempty" xml:"InvolvedObjectKind,omitempty"`
	InvolvedObjectName *string                                                   `json:"InvolvedObjectName,omitempty" xml:"InvolvedObjectName,omitempty"`
	Message            *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	Severity           *string                                                   `json:"Severity,omitempty" xml:"Severity,omitempty"`
	StartsAt           *int64                                                    `json:"StartsAt,omitempty" xml:"StartsAt,omitempty"`
	Status             *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListActivatedAlertsResponseBodyPageAlerts) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedAlertsResponseBodyPageAlerts) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetAlertId(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.AlertId = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetAlertName(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.AlertName = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetAlertType(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.AlertType = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetCount(v int32) *ListActivatedAlertsResponseBodyPageAlerts {
	s.Count = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetCreateTime(v int64) *ListActivatedAlertsResponseBodyPageAlerts {
	s.CreateTime = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetDispatchRules(v []*ListActivatedAlertsResponseBodyPageAlertsDispatchRules) *ListActivatedAlertsResponseBodyPageAlerts {
	s.DispatchRules = v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetEndsAt(v int64) *ListActivatedAlertsResponseBodyPageAlerts {
	s.EndsAt = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetExpandFields(v map[string]interface{}) *ListActivatedAlertsResponseBodyPageAlerts {
	s.ExpandFields = v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetIntegrationName(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.IntegrationName = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetIntegrationType(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.IntegrationType = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetInvolvedObjectKind(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.InvolvedObjectKind = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetInvolvedObjectName(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.InvolvedObjectName = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetMessage(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.Message = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetSeverity(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.Severity = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetStartsAt(v int64) *ListActivatedAlertsResponseBodyPageAlerts {
	s.StartsAt = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlerts) SetStatus(v string) *ListActivatedAlertsResponseBodyPageAlerts {
	s.Status = &v
	return s
}

type ListActivatedAlertsResponseBodyPageAlertsDispatchRules struct {
	RuleId   *int32  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListActivatedAlertsResponseBodyPageAlertsDispatchRules) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedAlertsResponseBodyPageAlertsDispatchRules) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponseBodyPageAlertsDispatchRules) SetRuleId(v int32) *ListActivatedAlertsResponseBodyPageAlertsDispatchRules {
	s.RuleId = &v
	return s
}

func (s *ListActivatedAlertsResponseBodyPageAlertsDispatchRules) SetRuleName(v string) *ListActivatedAlertsResponseBodyPageAlertsDispatchRules {
	s.RuleName = &v
	return s
}

type ListActivatedAlertsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListActivatedAlertsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListActivatedAlertsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListActivatedAlertsResponse) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsResponse) SetHeaders(v map[string]*string) *ListActivatedAlertsResponse {
	s.Headers = v
	return s
}

func (s *ListActivatedAlertsResponse) SetStatusCode(v int32) *ListActivatedAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActivatedAlertsResponse) SetBody(v *ListActivatedAlertsResponseBody) *ListActivatedAlertsResponse {
	s.Body = v
	return s
}

type ListAlertsRequest struct {
	AlertName       *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	DispatchRuleId  *int64  `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IntegrationType *string `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	Page            *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	Severity        *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	ShowActivities  *bool   `json:"ShowActivities,omitempty" xml:"ShowActivities,omitempty"`
	ShowEvents      *bool   `json:"ShowEvents,omitempty" xml:"ShowEvents,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State           *int64  `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListAlertsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertsRequest) SetAlertName(v string) *ListAlertsRequest {
	s.AlertName = &v
	return s
}

func (s *ListAlertsRequest) SetDispatchRuleId(v int64) *ListAlertsRequest {
	s.DispatchRuleId = &v
	return s
}

func (s *ListAlertsRequest) SetEndTime(v string) *ListAlertsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAlertsRequest) SetIntegrationType(v string) *ListAlertsRequest {
	s.IntegrationType = &v
	return s
}

func (s *ListAlertsRequest) SetPage(v int64) *ListAlertsRequest {
	s.Page = &v
	return s
}

func (s *ListAlertsRequest) SetSeverity(v string) *ListAlertsRequest {
	s.Severity = &v
	return s
}

func (s *ListAlertsRequest) SetShowActivities(v bool) *ListAlertsRequest {
	s.ShowActivities = &v
	return s
}

func (s *ListAlertsRequest) SetShowEvents(v bool) *ListAlertsRequest {
	s.ShowEvents = &v
	return s
}

func (s *ListAlertsRequest) SetSize(v int64) *ListAlertsRequest {
	s.Size = &v
	return s
}

func (s *ListAlertsRequest) SetStartTime(v string) *ListAlertsRequest {
	s.StartTime = &v
	return s
}

func (s *ListAlertsRequest) SetState(v int64) *ListAlertsRequest {
	s.State = &v
	return s
}

type ListAlertsResponseBody struct {
	PageBean  *ListAlertsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAlertsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBody) SetPageBean(v *ListAlertsResponseBodyPageBean) *ListAlertsResponseBody {
	s.PageBean = v
	return s
}

func (s *ListAlertsResponseBody) SetRequestId(v string) *ListAlertsResponseBody {
	s.RequestId = &v
	return s
}

type ListAlertsResponseBodyPageBean struct {
	ListAlerts []*ListAlertsResponseBodyPageBeanListAlerts `json:"ListAlerts,omitempty" xml:"ListAlerts,omitempty" type:"Repeated"`
	Page       *int64                                      `json:"Page,omitempty" xml:"Page,omitempty"`
	Size       *int64                                      `json:"Size,omitempty" xml:"Size,omitempty"`
	Total      *int64                                      `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAlertsResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBodyPageBean) SetListAlerts(v []*ListAlertsResponseBodyPageBeanListAlerts) *ListAlertsResponseBodyPageBean {
	s.ListAlerts = v
	return s
}

func (s *ListAlertsResponseBodyPageBean) SetPage(v int64) *ListAlertsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListAlertsResponseBodyPageBean) SetSize(v int64) *ListAlertsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListAlertsResponseBodyPageBean) SetTotal(v int64) *ListAlertsResponseBodyPageBean {
	s.Total = &v
	return s
}

type ListAlertsResponseBodyPageBeanListAlerts struct {
	Activities       []*ListAlertsResponseBodyPageBeanListAlertsActivities  `json:"Activities,omitempty" xml:"Activities,omitempty" type:"Repeated"`
	AlertEvents      []*ListAlertsResponseBodyPageBeanListAlertsAlertEvents `json:"AlertEvents,omitempty" xml:"AlertEvents,omitempty" type:"Repeated"`
	AlertId          *int64                                                 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName        *string                                                `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	CreateTime       *string                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DispatchRuleId   *float32                                               `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	DispatchRuleName *string                                                `json:"DispatchRuleName,omitempty" xml:"DispatchRuleName,omitempty"`
	Severity         *string                                                `json:"Severity,omitempty" xml:"Severity,omitempty"`
	State            *int64                                                 `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListAlertsResponseBodyPageBeanListAlerts) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsResponseBodyPageBeanListAlerts) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetActivities(v []*ListAlertsResponseBodyPageBeanListAlertsActivities) *ListAlertsResponseBodyPageBeanListAlerts {
	s.Activities = v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetAlertEvents(v []*ListAlertsResponseBodyPageBeanListAlertsAlertEvents) *ListAlertsResponseBodyPageBeanListAlerts {
	s.AlertEvents = v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetAlertId(v int64) *ListAlertsResponseBodyPageBeanListAlerts {
	s.AlertId = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetAlertName(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.AlertName = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetCreateTime(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.CreateTime = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetDispatchRuleId(v float32) *ListAlertsResponseBodyPageBeanListAlerts {
	s.DispatchRuleId = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetDispatchRuleName(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.DispatchRuleName = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetSeverity(v string) *ListAlertsResponseBodyPageBeanListAlerts {
	s.Severity = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlerts) SetState(v int64) *ListAlertsResponseBodyPageBeanListAlerts {
	s.State = &v
	return s
}

type ListAlertsResponseBodyPageBeanListAlertsActivities struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HandlerName *string `json:"HandlerName,omitempty" xml:"HandlerName,omitempty"`
	Time        *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Type        *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAlertsResponseBodyPageBeanListAlertsActivities) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsResponseBodyPageBeanListAlertsActivities) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) SetContent(v string) *ListAlertsResponseBodyPageBeanListAlertsActivities {
	s.Content = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) SetDescription(v string) *ListAlertsResponseBodyPageBeanListAlertsActivities {
	s.Description = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) SetHandlerName(v string) *ListAlertsResponseBodyPageBeanListAlertsActivities {
	s.HandlerName = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) SetTime(v string) *ListAlertsResponseBodyPageBeanListAlertsActivities {
	s.Time = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsActivities) SetType(v int64) *ListAlertsResponseBodyPageBeanListAlertsActivities {
	s.Type = &v
	return s
}

type ListAlertsResponseBodyPageBeanListAlertsAlertEvents struct {
	AlertName       *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations     *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GeneratorURL    *string `json:"GeneratorURL,omitempty" xml:"GeneratorURL,omitempty"`
	IntegrationName *string `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	IntegrationType *string `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	Labels          *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	ReceiveTime     *string `json:"ReceiveTime,omitempty" xml:"ReceiveTime,omitempty"`
	Severity        *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListAlertsResponseBodyPageBeanListAlertsAlertEvents) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsResponseBodyPageBeanListAlertsAlertEvents) GoString() string {
	return s.String()
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetAlertName(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.AlertName = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetAnnotations(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.Annotations = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetDescription(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.Description = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetEndTime(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.EndTime = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetGeneratorURL(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.GeneratorURL = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetIntegrationName(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.IntegrationName = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetIntegrationType(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.IntegrationType = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetLabels(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.Labels = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetReceiveTime(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.ReceiveTime = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetSeverity(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.Severity = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetStartTime(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.StartTime = &v
	return s
}

func (s *ListAlertsResponseBodyPageBeanListAlertsAlertEvents) SetState(v string) *ListAlertsResponseBodyPageBeanListAlertsAlertEvents {
	s.State = &v
	return s
}

type ListAlertsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAlertsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlertsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlertsResponse) GoString() string {
	return s.String()
}

func (s *ListAlertsResponse) SetHeaders(v map[string]*string) *ListAlertsResponse {
	s.Headers = v
	return s
}

func (s *ListAlertsResponse) SetStatusCode(v int32) *ListAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAlertsResponse) SetBody(v *ListAlertsResponseBody) *ListAlertsResponse {
	s.Body = v
	return s
}

type ListClusterFromGrafanaRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListClusterFromGrafanaRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterFromGrafanaRequest) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaRequest) SetRegionId(v string) *ListClusterFromGrafanaRequest {
	s.RegionId = &v
	return s
}

type ListClusterFromGrafanaResponseBody struct {
	PromClusterList []*ListClusterFromGrafanaResponseBodyPromClusterList `json:"PromClusterList,omitempty" xml:"PromClusterList,omitempty" type:"Repeated"`
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterFromGrafanaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterFromGrafanaResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaResponseBody) SetPromClusterList(v []*ListClusterFromGrafanaResponseBodyPromClusterList) *ListClusterFromGrafanaResponseBody {
	s.PromClusterList = v
	return s
}

func (s *ListClusterFromGrafanaResponseBody) SetRequestId(v string) *ListClusterFromGrafanaResponseBody {
	s.RequestId = &v
	return s
}

type ListClusterFromGrafanaResponseBodyPromClusterList struct {
	AgentStatus           *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName           *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterType           *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	ControllerId          *string `json:"ControllerId,omitempty" xml:"ControllerId,omitempty"`
	CreateTime            *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Extra                 *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Id                    *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstallTime           *int64  `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	IsControllerInstalled *bool   `json:"IsControllerInstalled,omitempty" xml:"IsControllerInstalled,omitempty"`
	LastHeartBeatTime     *int64  `json:"LastHeartBeatTime,omitempty" xml:"LastHeartBeatTime,omitempty"`
	NodeNum               *int32  `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	Options               *string `json:"Options,omitempty" xml:"Options,omitempty"`
	PluginsJsonArray      *string `json:"PluginsJsonArray,omitempty" xml:"PluginsJsonArray,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StateJson             *string `json:"StateJson,omitempty" xml:"StateJson,omitempty"`
	UpdateTime            *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId                *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListClusterFromGrafanaResponseBodyPromClusterList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterFromGrafanaResponseBodyPromClusterList) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetAgentStatus(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.AgentStatus = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetClusterId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ClusterId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetClusterName(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ClusterName = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetClusterType(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ClusterType = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetControllerId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.ControllerId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetCreateTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.CreateTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetExtra(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.Extra = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetId(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.Id = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetInstallTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.InstallTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetIsControllerInstalled(v bool) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.IsControllerInstalled = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetLastHeartBeatTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.LastHeartBeatTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetNodeNum(v int32) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.NodeNum = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetOptions(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.Options = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetPluginsJsonArray(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.PluginsJsonArray = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetRegionId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.RegionId = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetStateJson(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.StateJson = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetUpdateTime(v int64) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.UpdateTime = &v
	return s
}

func (s *ListClusterFromGrafanaResponseBodyPromClusterList) SetUserId(v string) *ListClusterFromGrafanaResponseBodyPromClusterList {
	s.UserId = &v
	return s
}

type ListClusterFromGrafanaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListClusterFromGrafanaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterFromGrafanaResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterFromGrafanaResponse) GoString() string {
	return s.String()
}

func (s *ListClusterFromGrafanaResponse) SetHeaders(v map[string]*string) *ListClusterFromGrafanaResponse {
	s.Headers = v
	return s
}

func (s *ListClusterFromGrafanaResponse) SetStatusCode(v int32) *ListClusterFromGrafanaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterFromGrafanaResponse) SetBody(v *ListClusterFromGrafanaResponseBody) *ListClusterFromGrafanaResponse {
	s.Body = v
	return s
}

type ListCmsInstancesRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCmsInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCmsInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListCmsInstancesRequest) SetClusterId(v string) *ListCmsInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCmsInstancesRequest) SetRegionId(v string) *ListCmsInstancesRequest {
	s.RegionId = &v
	return s
}

type ListCmsInstancesResponseBody struct {
	Data      *ListCmsInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCmsInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCmsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCmsInstancesResponseBody) SetData(v *ListCmsInstancesResponseBodyData) *ListCmsInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListCmsInstancesResponseBody) SetRequestId(v string) *ListCmsInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListCmsInstancesResponseBodyData struct {
	EnableTag *bool                                       `json:"EnableTag,omitempty" xml:"EnableTag,omitempty"`
	Products  []*ListCmsInstancesResponseBodyDataProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
}

func (s ListCmsInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCmsInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCmsInstancesResponseBodyData) SetEnableTag(v bool) *ListCmsInstancesResponseBodyData {
	s.EnableTag = &v
	return s
}

func (s *ListCmsInstancesResponseBodyData) SetProducts(v []*ListCmsInstancesResponseBodyDataProducts) *ListCmsInstancesResponseBodyData {
	s.Products = v
	return s
}

type ListCmsInstancesResponseBodyDataProducts struct {
	Descr    *string `json:"Descr,omitempty" xml:"Descr,omitempty"`
	Id       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Instance *string `json:"Instance,omitempty" xml:"Instance,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Prod     *string `json:"Prod,omitempty" xml:"Prod,omitempty"`
	Source   *string `json:"Source,omitempty" xml:"Source,omitempty"`
	State    *string `json:"State,omitempty" xml:"State,omitempty"`
	Time     *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Url      *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListCmsInstancesResponseBodyDataProducts) String() string {
	return tea.Prettify(s)
}

func (s ListCmsInstancesResponseBodyDataProducts) GoString() string {
	return s.String()
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetDescr(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Descr = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetId(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Id = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetInstance(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Instance = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetName(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Name = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetProd(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Prod = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetSource(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Source = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetState(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.State = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetTime(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Time = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetType(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Type = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetUrl(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Url = &v
	return s
}

type ListCmsInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCmsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCmsInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCmsInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListCmsInstancesResponse) SetHeaders(v map[string]*string) *ListCmsInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListCmsInstancesResponse) SetStatusCode(v int32) *ListCmsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCmsInstancesResponse) SetBody(v *ListCmsInstancesResponseBody) *ListCmsInstancesResponse {
	s.Body = v
	return s
}

type ListDashboardsRequest struct {
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType    *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	DashboardName  *string `json:"DashboardName,omitempty" xml:"DashboardName,omitempty"`
	Language       *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Product        *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RecreateSwitch *bool   `json:"RecreateSwitch,omitempty" xml:"RecreateSwitch,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListDashboardsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardsRequest) SetClusterId(v string) *ListDashboardsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDashboardsRequest) SetClusterType(v string) *ListDashboardsRequest {
	s.ClusterType = &v
	return s
}

func (s *ListDashboardsRequest) SetDashboardName(v string) *ListDashboardsRequest {
	s.DashboardName = &v
	return s
}

func (s *ListDashboardsRequest) SetLanguage(v string) *ListDashboardsRequest {
	s.Language = &v
	return s
}

func (s *ListDashboardsRequest) SetProduct(v string) *ListDashboardsRequest {
	s.Product = &v
	return s
}

func (s *ListDashboardsRequest) SetRecreateSwitch(v bool) *ListDashboardsRequest {
	s.RecreateSwitch = &v
	return s
}

func (s *ListDashboardsRequest) SetRegionId(v string) *ListDashboardsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDashboardsRequest) SetTitle(v string) *ListDashboardsRequest {
	s.Title = &v
	return s
}

type ListDashboardsResponseBody struct {
	DashboardVos []*ListDashboardsResponseBodyDashboardVos `json:"DashboardVos,omitempty" xml:"DashboardVos,omitempty" type:"Repeated"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDashboardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponseBody) SetDashboardVos(v []*ListDashboardsResponseBodyDashboardVos) *ListDashboardsResponseBody {
	s.DashboardVos = v
	return s
}

func (s *ListDashboardsResponseBody) SetRequestId(v string) *ListDashboardsResponseBody {
	s.RequestId = &v
	return s
}

type ListDashboardsResponseBodyDashboardVos struct {
	DashboardType  *string                                          `json:"DashboardType,omitempty" xml:"DashboardType,omitempty"`
	Exporter       *string                                          `json:"Exporter,omitempty" xml:"Exporter,omitempty"`
	HttpUrl        *string                                          `json:"HttpUrl,omitempty" xml:"HttpUrl,omitempty"`
	HttpsUrl       *string                                          `json:"HttpsUrl,omitempty" xml:"HttpsUrl,omitempty"`
	I18nChild      *ListDashboardsResponseBodyDashboardVosI18nChild `json:"I18nChild,omitempty" xml:"I18nChild,omitempty" type:"Struct"`
	Id             *string                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	IsArmsExporter *bool                                            `json:"IsArmsExporter,omitempty" xml:"IsArmsExporter,omitempty"`
	Kind           *string                                          `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Language       *string                                          `json:"Language,omitempty" xml:"Language,omitempty"`
	Name           *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedUpdate     *bool                                            `json:"NeedUpdate,omitempty" xml:"NeedUpdate,omitempty"`
	Tags           []*string                                        `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Time           *string                                          `json:"Time,omitempty" xml:"Time,omitempty"`
	Title          *string                                          `json:"Title,omitempty" xml:"Title,omitempty"`
	Type           *string                                          `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid            *string                                          `json:"Uid,omitempty" xml:"Uid,omitempty"`
	Url            *string                                          `json:"Url,omitempty" xml:"Url,omitempty"`
	Version        *string                                          `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListDashboardsResponseBodyDashboardVos) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsResponseBodyDashboardVos) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponseBodyDashboardVos) SetDashboardType(v string) *ListDashboardsResponseBodyDashboardVos {
	s.DashboardType = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetExporter(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Exporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetHttpUrl(v string) *ListDashboardsResponseBodyDashboardVos {
	s.HttpUrl = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetHttpsUrl(v string) *ListDashboardsResponseBodyDashboardVos {
	s.HttpsUrl = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetI18nChild(v *ListDashboardsResponseBodyDashboardVosI18nChild) *ListDashboardsResponseBodyDashboardVos {
	s.I18nChild = v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetId(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Id = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetIsArmsExporter(v bool) *ListDashboardsResponseBodyDashboardVos {
	s.IsArmsExporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetKind(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Kind = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetLanguage(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Language = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetName(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Name = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetNeedUpdate(v bool) *ListDashboardsResponseBodyDashboardVos {
	s.NeedUpdate = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTags(v []*string) *ListDashboardsResponseBodyDashboardVos {
	s.Tags = v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTime(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Time = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetTitle(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Title = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetType(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Type = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetUid(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Uid = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetUrl(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Url = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVos) SetVersion(v string) *ListDashboardsResponseBodyDashboardVos {
	s.Version = &v
	return s
}

type ListDashboardsResponseBodyDashboardVosI18nChild struct {
	DashboardType  *string   `json:"DashboardType,omitempty" xml:"DashboardType,omitempty"`
	Exporter       *string   `json:"Exporter,omitempty" xml:"Exporter,omitempty"`
	HttpUrl        *string   `json:"HttpUrl,omitempty" xml:"HttpUrl,omitempty"`
	HttpsUrl       *string   `json:"HttpsUrl,omitempty" xml:"HttpsUrl,omitempty"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	IsArmsExporter *bool     `json:"IsArmsExporter,omitempty" xml:"IsArmsExporter,omitempty"`
	Kind           *string   `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Language       *string   `json:"Language,omitempty" xml:"Language,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedUpdate     *bool     `json:"NeedUpdate,omitempty" xml:"NeedUpdate,omitempty"`
	Tags           []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Time           *string   `json:"Time,omitempty" xml:"Time,omitempty"`
	Title          *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	Type           *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid            *string   `json:"Uid,omitempty" xml:"Uid,omitempty"`
	Url            *string   `json:"Url,omitempty" xml:"Url,omitempty"`
	Version        *string   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListDashboardsResponseBodyDashboardVosI18nChild) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsResponseBodyDashboardVosI18nChild) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetDashboardType(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.DashboardType = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetExporter(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Exporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetHttpUrl(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.HttpUrl = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetHttpsUrl(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.HttpsUrl = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetId(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Id = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetIsArmsExporter(v bool) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.IsArmsExporter = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetKind(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Kind = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetLanguage(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Language = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetName(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Name = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetNeedUpdate(v bool) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.NeedUpdate = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetTags(v []*string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Tags = v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetTime(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Time = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetTitle(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Title = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetType(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Type = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetUid(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Uid = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetUrl(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Url = &v
	return s
}

func (s *ListDashboardsResponseBodyDashboardVosI18nChild) SetVersion(v string) *ListDashboardsResponseBodyDashboardVosI18nChild {
	s.Version = &v
	return s
}

type ListDashboardsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDashboardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDashboardsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardsResponse) SetHeaders(v map[string]*string) *ListDashboardsResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardsResponse) SetStatusCode(v int32) *ListDashboardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDashboardsResponse) SetBody(v *ListDashboardsResponseBody) *ListDashboardsResponse {
	s.Body = v
	return s
}

type ListDashboardsByNameRequest struct {
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType      *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	DashBoardName    *string `json:"DashBoardName,omitempty" xml:"DashBoardName,omitempty"`
	DashBoardVersion *string `json:"DashBoardVersion,omitempty" xml:"DashBoardVersion,omitempty"`
	DataSourceType   *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	GroupName        *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OnlyQuery        *bool   `json:"OnlyQuery,omitempty" xml:"OnlyQuery,omitempty"`
	ProductCode      *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDashboardsByNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsByNameRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardsByNameRequest) SetClusterId(v string) *ListDashboardsByNameRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetClusterType(v string) *ListDashboardsByNameRequest {
	s.ClusterType = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetDashBoardName(v string) *ListDashboardsByNameRequest {
	s.DashBoardName = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetDashBoardVersion(v string) *ListDashboardsByNameRequest {
	s.DashBoardVersion = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetDataSourceType(v string) *ListDashboardsByNameRequest {
	s.DataSourceType = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetGroupName(v string) *ListDashboardsByNameRequest {
	s.GroupName = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetOnlyQuery(v bool) *ListDashboardsByNameRequest {
	s.OnlyQuery = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetProductCode(v string) *ListDashboardsByNameRequest {
	s.ProductCode = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetRegionId(v string) *ListDashboardsByNameRequest {
	s.RegionId = &v
	return s
}

type ListDashboardsByNameResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDashboardsByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsByNameResponseBody) GoString() string {
	return s.String()
}

func (s *ListDashboardsByNameResponseBody) SetData(v string) *ListDashboardsByNameResponseBody {
	s.Data = &v
	return s
}

func (s *ListDashboardsByNameResponseBody) SetRequestId(v string) *ListDashboardsByNameResponseBody {
	s.RequestId = &v
	return s
}

type ListDashboardsByNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDashboardsByNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDashboardsByNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDashboardsByNameResponse) GoString() string {
	return s.String()
}

func (s *ListDashboardsByNameResponse) SetHeaders(v map[string]*string) *ListDashboardsByNameResponse {
	s.Headers = v
	return s
}

func (s *ListDashboardsByNameResponse) SetStatusCode(v int32) *ListDashboardsByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDashboardsByNameResponse) SetBody(v *ListDashboardsByNameResponseBody) *ListDashboardsByNameResponse {
	s.Body = v
	return s
}

type ListDispatchRuleRequest struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	System   *bool   `json:"System,omitempty" xml:"System,omitempty"`
}

func (s ListDispatchRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *ListDispatchRuleRequest) SetName(v string) *ListDispatchRuleRequest {
	s.Name = &v
	return s
}

func (s *ListDispatchRuleRequest) SetRegionId(v string) *ListDispatchRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ListDispatchRuleRequest) SetSystem(v bool) *ListDispatchRuleRequest {
	s.System = &v
	return s
}

type ListDispatchRuleResponseBody struct {
	DispatchRules []*ListDispatchRuleResponseBodyDispatchRules `json:"DispatchRules,omitempty" xml:"DispatchRules,omitempty" type:"Repeated"`
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDispatchRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListDispatchRuleResponseBody) SetDispatchRules(v []*ListDispatchRuleResponseBodyDispatchRules) *ListDispatchRuleResponseBody {
	s.DispatchRules = v
	return s
}

func (s *ListDispatchRuleResponseBody) SetRequestId(v string) *ListDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

type ListDispatchRuleResponseBodyDispatchRules struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleId *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	State  *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListDispatchRuleResponseBodyDispatchRules) String() string {
	return tea.Prettify(s)
}

func (s ListDispatchRuleResponseBodyDispatchRules) GoString() string {
	return s.String()
}

func (s *ListDispatchRuleResponseBodyDispatchRules) SetName(v string) *ListDispatchRuleResponseBodyDispatchRules {
	s.Name = &v
	return s
}

func (s *ListDispatchRuleResponseBodyDispatchRules) SetRuleId(v int64) *ListDispatchRuleResponseBodyDispatchRules {
	s.RuleId = &v
	return s
}

func (s *ListDispatchRuleResponseBodyDispatchRules) SetState(v string) *ListDispatchRuleResponseBodyDispatchRules {
	s.State = &v
	return s
}

type ListDispatchRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDispatchRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *ListDispatchRuleResponse) SetHeaders(v map[string]*string) *ListDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *ListDispatchRuleResponse) SetStatusCode(v int32) *ListDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDispatchRuleResponse) SetBody(v *ListDispatchRuleResponseBody) *ListDispatchRuleResponse {
	s.Body = v
	return s
}

type ListEscalationPoliciesRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Page *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListEscalationPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListEscalationPoliciesRequest) SetName(v string) *ListEscalationPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListEscalationPoliciesRequest) SetPage(v int64) *ListEscalationPoliciesRequest {
	s.Page = &v
	return s
}

func (s *ListEscalationPoliciesRequest) SetSize(v int64) *ListEscalationPoliciesRequest {
	s.Size = &v
	return s
}

type ListEscalationPoliciesResponseBody struct {
	PageBean  *ListEscalationPoliciesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEscalationPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEscalationPoliciesResponseBody) SetPageBean(v *ListEscalationPoliciesResponseBodyPageBean) *ListEscalationPoliciesResponseBody {
	s.PageBean = v
	return s
}

func (s *ListEscalationPoliciesResponseBody) SetRequestId(v string) *ListEscalationPoliciesResponseBody {
	s.RequestId = &v
	return s
}

type ListEscalationPoliciesResponseBodyPageBean struct {
	EscalationPolicies []*ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies `json:"EscalationPolicies,omitempty" xml:"EscalationPolicies,omitempty" type:"Repeated"`
	Page               *int64                                                          `json:"Page,omitempty" xml:"Page,omitempty"`
	Size               *int64                                                          `json:"Size,omitempty" xml:"Size,omitempty"`
	Total              *int64                                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEscalationPoliciesResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPoliciesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListEscalationPoliciesResponseBodyPageBean) SetEscalationPolicies(v []*ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) *ListEscalationPoliciesResponseBodyPageBean {
	s.EscalationPolicies = v
	return s
}

func (s *ListEscalationPoliciesResponseBodyPageBean) SetPage(v int64) *ListEscalationPoliciesResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListEscalationPoliciesResponseBodyPageBean) SetSize(v int64) *ListEscalationPoliciesResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListEscalationPoliciesResponseBodyPageBean) SetTotal(v int64) *ListEscalationPoliciesResponseBodyPageBean {
	s.Total = &v
	return s
}

type ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies struct {
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) GoString() string {
	return s.String()
}

func (s *ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) SetId(v int64) *ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies {
	s.Id = &v
	return s
}

func (s *ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) SetName(v string) *ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies {
	s.Name = &v
	return s
}

type ListEscalationPoliciesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEscalationPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEscalationPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEscalationPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListEscalationPoliciesResponse) SetHeaders(v map[string]*string) *ListEscalationPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListEscalationPoliciesResponse) SetStatusCode(v int32) *ListEscalationPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEscalationPoliciesResponse) SetBody(v *ListEscalationPoliciesResponseBody) *ListEscalationPoliciesResponse {
	s.Body = v
	return s
}

type ListEventBridgeIntegrationsRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Page *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListEventBridgeIntegrationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventBridgeIntegrationsRequest) GoString() string {
	return s.String()
}

func (s *ListEventBridgeIntegrationsRequest) SetName(v string) *ListEventBridgeIntegrationsRequest {
	s.Name = &v
	return s
}

func (s *ListEventBridgeIntegrationsRequest) SetPage(v int64) *ListEventBridgeIntegrationsRequest {
	s.Page = &v
	return s
}

func (s *ListEventBridgeIntegrationsRequest) SetSize(v int64) *ListEventBridgeIntegrationsRequest {
	s.Size = &v
	return s
}

type ListEventBridgeIntegrationsResponseBody struct {
	PageBean  *ListEventBridgeIntegrationsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEventBridgeIntegrationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventBridgeIntegrationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventBridgeIntegrationsResponseBody) SetPageBean(v *ListEventBridgeIntegrationsResponseBodyPageBean) *ListEventBridgeIntegrationsResponseBody {
	s.PageBean = v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBody) SetRequestId(v string) *ListEventBridgeIntegrationsResponseBody {
	s.RequestId = &v
	return s
}

type ListEventBridgeIntegrationsResponseBodyPageBean struct {
	EventBridgeIntegrations []*ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations `json:"EventBridgeIntegrations,omitempty" xml:"EventBridgeIntegrations,omitempty" type:"Repeated"`
	Page                    *int64                                                                    `json:"Page,omitempty" xml:"Page,omitempty"`
	Size                    *int64                                                                    `json:"Size,omitempty" xml:"Size,omitempty"`
	Total                   *int64                                                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEventBridgeIntegrationsResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s ListEventBridgeIntegrationsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) SetEventBridgeIntegrations(v []*ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) *ListEventBridgeIntegrationsResponseBodyPageBean {
	s.EventBridgeIntegrations = v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) SetPage(v int64) *ListEventBridgeIntegrationsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) SetSize(v int64) *ListEventBridgeIntegrationsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) SetTotal(v int64) *ListEventBridgeIntegrationsResponseBodyPageBean {
	s.Total = &v
	return s
}

type ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) String() string {
	return tea.Prettify(s)
}

func (s ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) GoString() string {
	return s.String()
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) SetDescription(v string) *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations {
	s.Description = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) SetId(v int64) *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations {
	s.Id = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) SetName(v string) *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations {
	s.Name = &v
	return s
}

type ListEventBridgeIntegrationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEventBridgeIntegrationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEventBridgeIntegrationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventBridgeIntegrationsResponse) GoString() string {
	return s.String()
}

func (s *ListEventBridgeIntegrationsResponse) SetHeaders(v map[string]*string) *ListEventBridgeIntegrationsResponse {
	s.Headers = v
	return s
}

func (s *ListEventBridgeIntegrationsResponse) SetStatusCode(v int32) *ListEventBridgeIntegrationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponse) SetBody(v *ListEventBridgeIntegrationsResponseBody) *ListEventBridgeIntegrationsResponse {
	s.Body = v
	return s
}

type ListInsightsEventsRequest struct {
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InsightsTypes *string `json:"InsightsTypes,omitempty" xml:"InsightsTypes,omitempty"`
	Pid           *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListInsightsEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInsightsEventsRequest) GoString() string {
	return s.String()
}

func (s *ListInsightsEventsRequest) SetEndTime(v string) *ListInsightsEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ListInsightsEventsRequest) SetInsightsTypes(v string) *ListInsightsEventsRequest {
	s.InsightsTypes = &v
	return s
}

func (s *ListInsightsEventsRequest) SetPid(v string) *ListInsightsEventsRequest {
	s.Pid = &v
	return s
}

func (s *ListInsightsEventsRequest) SetRegionId(v string) *ListInsightsEventsRequest {
	s.RegionId = &v
	return s
}

func (s *ListInsightsEventsRequest) SetStartTime(v string) *ListInsightsEventsRequest {
	s.StartTime = &v
	return s
}

type ListInsightsEventsResponseBody struct {
	InsightsEvents []*ListInsightsEventsResponseBodyInsightsEvents `json:"InsightsEvents,omitempty" xml:"InsightsEvents,omitempty" type:"Repeated"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInsightsEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInsightsEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInsightsEventsResponseBody) SetInsightsEvents(v []*ListInsightsEventsResponseBodyInsightsEvents) *ListInsightsEventsResponseBody {
	s.InsightsEvents = v
	return s
}

func (s *ListInsightsEventsResponseBody) SetRequestId(v string) *ListInsightsEventsResponseBody {
	s.RequestId = &v
	return s
}

type ListInsightsEventsResponseBodyInsightsEvents struct {
	Date  *int64  `json:"Date,omitempty" xml:"Date,omitempty"`
	Desc  *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Pid   *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInsightsEventsResponseBodyInsightsEvents) String() string {
	return tea.Prettify(s)
}

func (s ListInsightsEventsResponseBodyInsightsEvents) GoString() string {
	return s.String()
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetDate(v int64) *ListInsightsEventsResponseBodyInsightsEvents {
	s.Date = &v
	return s
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetDesc(v string) *ListInsightsEventsResponseBodyInsightsEvents {
	s.Desc = &v
	return s
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetLevel(v string) *ListInsightsEventsResponseBodyInsightsEvents {
	s.Level = &v
	return s
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetPid(v string) *ListInsightsEventsResponseBodyInsightsEvents {
	s.Pid = &v
	return s
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetTitle(v string) *ListInsightsEventsResponseBodyInsightsEvents {
	s.Title = &v
	return s
}

func (s *ListInsightsEventsResponseBodyInsightsEvents) SetType(v string) *ListInsightsEventsResponseBodyInsightsEvents {
	s.Type = &v
	return s
}

type ListInsightsEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInsightsEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInsightsEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInsightsEventsResponse) GoString() string {
	return s.String()
}

func (s *ListInsightsEventsResponse) SetHeaders(v map[string]*string) *ListInsightsEventsResponse {
	s.Headers = v
	return s
}

func (s *ListInsightsEventsResponse) SetStatusCode(v int32) *ListInsightsEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInsightsEventsResponse) SetBody(v *ListInsightsEventsResponseBody) *ListInsightsEventsResponse {
	s.Body = v
	return s
}

type ListIntegrationRequest struct {
	IntegrationName        *string `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	IntegrationProductType *string `json:"IntegrationProductType,omitempty" xml:"IntegrationProductType,omitempty"`
	IsDetail               *bool   `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	Page                   *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size                   *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListIntegrationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationRequest) SetIntegrationName(v string) *ListIntegrationRequest {
	s.IntegrationName = &v
	return s
}

func (s *ListIntegrationRequest) SetIntegrationProductType(v string) *ListIntegrationRequest {
	s.IntegrationProductType = &v
	return s
}

func (s *ListIntegrationRequest) SetIsDetail(v bool) *ListIntegrationRequest {
	s.IsDetail = &v
	return s
}

func (s *ListIntegrationRequest) SetPage(v int64) *ListIntegrationRequest {
	s.Page = &v
	return s
}

func (s *ListIntegrationRequest) SetSize(v int64) *ListIntegrationRequest {
	s.Size = &v
	return s
}

type ListIntegrationResponseBody struct {
	PageInfo  *ListIntegrationResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIntegrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntegrationResponseBody) SetPageInfo(v *ListIntegrationResponseBodyPageInfo) *ListIntegrationResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListIntegrationResponseBody) SetRequestId(v string) *ListIntegrationResponseBody {
	s.RequestId = &v
	return s
}

type ListIntegrationResponseBodyPageInfo struct {
	Iintegrations []*ListIntegrationResponseBodyPageInfoIintegrations `json:"Iintegrations,omitempty" xml:"Iintegrations,omitempty" type:"Repeated"`
	Page          *int64                                              `json:"Page,omitempty" xml:"Page,omitempty"`
	Size          *int64                                              `json:"Size,omitempty" xml:"Size,omitempty"`
	Total         *int64                                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListIntegrationResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListIntegrationResponseBodyPageInfo) SetIintegrations(v []*ListIntegrationResponseBodyPageInfoIintegrations) *ListIntegrationResponseBodyPageInfo {
	s.Iintegrations = v
	return s
}

func (s *ListIntegrationResponseBodyPageInfo) SetPage(v int64) *ListIntegrationResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfo) SetSize(v int64) *ListIntegrationResponseBodyPageInfo {
	s.Size = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfo) SetTotal(v int64) *ListIntegrationResponseBodyPageInfo {
	s.Total = &v
	return s
}

type ListIntegrationResponseBodyPageInfoIintegrations struct {
	ApiEndpoint            *string                                                            `json:"ApiEndpoint,omitempty" xml:"ApiEndpoint,omitempty"`
	CreateTime             *string                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IntegrationDetail      *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail `json:"IntegrationDetail,omitempty" xml:"IntegrationDetail,omitempty" type:"Struct"`
	IntegrationId          *int64                                                             `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
	IntegrationName        *string                                                            `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	IntegrationProductType *string                                                            `json:"IntegrationProductType,omitempty" xml:"IntegrationProductType,omitempty"`
	Liveness               *string                                                            `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	ShortToken             *string                                                            `json:"ShortToken,omitempty" xml:"ShortToken,omitempty"`
	State                  *bool                                                              `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListIntegrationResponseBodyPageInfoIintegrations) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationResponseBodyPageInfoIintegrations) GoString() string {
	return s.String()
}

func (s *ListIntegrationResponseBodyPageInfoIintegrations) SetApiEndpoint(v string) *ListIntegrationResponseBodyPageInfoIintegrations {
	s.ApiEndpoint = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrations) SetCreateTime(v string) *ListIntegrationResponseBodyPageInfoIintegrations {
	s.CreateTime = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrations) SetIntegrationDetail(v *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail) *ListIntegrationResponseBodyPageInfoIintegrations {
	s.IntegrationDetail = v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrations) SetIntegrationId(v int64) *ListIntegrationResponseBodyPageInfoIintegrations {
	s.IntegrationId = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrations) SetIntegrationName(v string) *ListIntegrationResponseBodyPageInfoIintegrations {
	s.IntegrationName = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrations) SetIntegrationProductType(v string) *ListIntegrationResponseBodyPageInfoIintegrations {
	s.IntegrationProductType = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrations) SetLiveness(v string) *ListIntegrationResponseBodyPageInfoIintegrations {
	s.Liveness = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrations) SetShortToken(v string) *ListIntegrationResponseBodyPageInfoIintegrations {
	s.ShortToken = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrations) SetState(v bool) *ListIntegrationResponseBodyPageInfoIintegrations {
	s.State = &v
	return s
}

type ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail struct {
	AutoRecover                *bool                    `json:"AutoRecover,omitempty" xml:"AutoRecover,omitempty"`
	Description                *string                  `json:"Description,omitempty" xml:"Description,omitempty"`
	DuplicateKey               *string                  `json:"DuplicateKey,omitempty" xml:"DuplicateKey,omitempty"`
	ExtendedFieldRedefineRules []map[string]interface{} `json:"ExtendedFieldRedefineRules,omitempty" xml:"ExtendedFieldRedefineRules,omitempty" type:"Repeated"`
	FieldRedefineRules         []map[string]interface{} `json:"FieldRedefineRules,omitempty" xml:"FieldRedefineRules,omitempty" type:"Repeated"`
	RecoverTime                *int64                   `json:"RecoverTime,omitempty" xml:"RecoverTime,omitempty"`
	Stat                       []*int64                 `json:"Stat,omitempty" xml:"Stat,omitempty" type:"Repeated"`
}

func (s ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail) GoString() string {
	return s.String()
}

func (s *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail) SetAutoRecover(v bool) *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail {
	s.AutoRecover = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail) SetDescription(v string) *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail {
	s.Description = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail) SetDuplicateKey(v string) *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail {
	s.DuplicateKey = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail) SetExtendedFieldRedefineRules(v []map[string]interface{}) *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail {
	s.ExtendedFieldRedefineRules = v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail) SetFieldRedefineRules(v []map[string]interface{}) *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail {
	s.FieldRedefineRules = v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail) SetRecoverTime(v int64) *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail {
	s.RecoverTime = &v
	return s
}

func (s *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail) SetStat(v []*int64) *ListIntegrationResponseBodyPageInfoIintegrationsIntegrationDetail {
	s.Stat = v
	return s
}

type ListIntegrationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIntegrationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIntegrationResponse) GoString() string {
	return s.String()
}

func (s *ListIntegrationResponse) SetHeaders(v map[string]*string) *ListIntegrationResponse {
	s.Headers = v
	return s
}

func (s *ListIntegrationResponse) SetStatusCode(v int32) *ListIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegrationResponse) SetBody(v *ListIntegrationResponseBody) *ListIntegrationResponse {
	s.Body = v
	return s
}

type ListNotificationPoliciesRequest struct {
	IsDetail *bool   `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Page     *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Size     *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListNotificationPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesRequest) SetIsDetail(v bool) *ListNotificationPoliciesRequest {
	s.IsDetail = &v
	return s
}

func (s *ListNotificationPoliciesRequest) SetName(v string) *ListNotificationPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListNotificationPoliciesRequest) SetPage(v int64) *ListNotificationPoliciesRequest {
	s.Page = &v
	return s
}

func (s *ListNotificationPoliciesRequest) SetRegionId(v string) *ListNotificationPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *ListNotificationPoliciesRequest) SetSize(v int64) *ListNotificationPoliciesRequest {
	s.Size = &v
	return s
}

type ListNotificationPoliciesResponseBody struct {
	PageBean  *ListNotificationPoliciesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNotificationPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBody) SetPageBean(v *ListNotificationPoliciesResponseBodyPageBean) *ListNotificationPoliciesResponseBody {
	s.PageBean = v
	return s
}

func (s *ListNotificationPoliciesResponseBody) SetRequestId(v string) *ListNotificationPoliciesResponseBody {
	s.RequestId = &v
	return s
}

type ListNotificationPoliciesResponseBodyPageBean struct {
	NotificationPolicies []*ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies `json:"NotificationPolicies,omitempty" xml:"NotificationPolicies,omitempty" type:"Repeated"`
	Page                 *int64                                                              `json:"Page,omitempty" xml:"Page,omitempty"`
	Size                 *int64                                                              `json:"Size,omitempty" xml:"Size,omitempty"`
	Total                *int64                                                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListNotificationPoliciesResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBean) SetNotificationPolicies(v []*ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) *ListNotificationPoliciesResponseBodyPageBean {
	s.NotificationPolicies = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBean) SetPage(v int64) *ListNotificationPoliciesResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBean) SetSize(v int64) *ListNotificationPoliciesResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBean) SetTotal(v int64) *ListNotificationPoliciesResponseBodyPageBean {
	s.Total = &v
	return s
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies struct {
	EscalationPolicyId *int64                                                                           `json:"EscalationPolicyId,omitempty" xml:"EscalationPolicyId,omitempty"`
	GroupRule          *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule       `json:"GroupRule,omitempty" xml:"GroupRule,omitempty" type:"Struct"`
	Id                 *int64                                                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	IntegrationId      *int64                                                                           `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
	MatchingRules      []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules `json:"MatchingRules,omitempty" xml:"MatchingRules,omitempty" type:"Repeated"`
	Name               *string                                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	NotifyRule         *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule      `json:"NotifyRule,omitempty" xml:"NotifyRule,omitempty" type:"Struct"`
	NotifyTemplate     *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate  `json:"NotifyTemplate,omitempty" xml:"NotifyTemplate,omitempty" type:"Struct"`
	Repeat             *bool                                                                            `json:"Repeat,omitempty" xml:"Repeat,omitempty"`
	RepeatInterval     *int64                                                                           `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
	SendRecoverMessage *bool                                                                            `json:"SendRecoverMessage,omitempty" xml:"SendRecoverMessage,omitempty"`
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetEscalationPolicyId(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.EscalationPolicyId = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetGroupRule(v *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.GroupRule = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetId(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.Id = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetIntegrationId(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.IntegrationId = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetMatchingRules(v []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.MatchingRules = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetName(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.Name = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetNotifyRule(v *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.NotifyRule = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetNotifyTemplate(v *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.NotifyTemplate = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetRepeat(v bool) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.Repeat = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetRepeatInterval(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.RepeatInterval = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies) SetSendRecoverMessage(v bool) *ListNotificationPoliciesResponseBodyPageBeanNotificationPolicies {
	s.SendRecoverMessage = &v
	return s
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule struct {
	GroupInterval  *int64    `json:"GroupInterval,omitempty" xml:"GroupInterval,omitempty"`
	GroupWait      *int64    `json:"GroupWait,omitempty" xml:"GroupWait,omitempty"`
	GroupingFields []*string `json:"GroupingFields,omitempty" xml:"GroupingFields,omitempty" type:"Repeated"`
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) SetGroupInterval(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule {
	s.GroupInterval = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) SetGroupWait(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule {
	s.GroupWait = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule) SetGroupingFields(v []*string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesGroupRule {
	s.GroupingFields = v
	return s
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules struct {
	MatchingConditions []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions `json:"MatchingConditions,omitempty" xml:"MatchingConditions,omitempty" type:"Repeated"`
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules) SetMatchingConditions(v []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRules {
	s.MatchingConditions = v
	return s
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) SetKey(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions {
	s.Key = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) SetOperator(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions {
	s.Operator = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions) SetValue(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesMatchingRulesMatchingConditions {
	s.Value = &v
	return s
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule struct {
	NotifyChannels  []*string                                                                                  `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	NotifyEndTime   *string                                                                                    `json:"NotifyEndTime,omitempty" xml:"NotifyEndTime,omitempty"`
	NotifyObjects   []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects `json:"NotifyObjects,omitempty" xml:"NotifyObjects,omitempty" type:"Repeated"`
	NotifyStartTime *string                                                                                    `json:"NotifyStartTime,omitempty" xml:"NotifyStartTime,omitempty"`
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) SetNotifyChannels(v []*string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule {
	s.NotifyChannels = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) SetNotifyEndTime(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule {
	s.NotifyEndTime = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) SetNotifyObjects(v []*ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule {
	s.NotifyObjects = v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule) SetNotifyStartTime(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRule {
	s.NotifyStartTime = &v
	return s
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects struct {
	NotifyObjectId   *int64  `json:"NotifyObjectId,omitempty" xml:"NotifyObjectId,omitempty"`
	NotifyObjectName *string `json:"NotifyObjectName,omitempty" xml:"NotifyObjectName,omitempty"`
	NotifyObjectType *string `json:"NotifyObjectType,omitempty" xml:"NotifyObjectType,omitempty"`
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) SetNotifyObjectId(v int64) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects {
	s.NotifyObjectId = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) SetNotifyObjectName(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects {
	s.NotifyObjectName = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects) SetNotifyObjectType(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyRuleNotifyObjects {
	s.NotifyObjectType = &v
	return s
}

type ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate struct {
	EmailContent        *string `json:"EmailContent,omitempty" xml:"EmailContent,omitempty"`
	EmailRecoverContent *string `json:"EmailRecoverContent,omitempty" xml:"EmailRecoverContent,omitempty"`
	EmailRecoverTitle   *string `json:"EmailRecoverTitle,omitempty" xml:"EmailRecoverTitle,omitempty"`
	EmailTitle          *string `json:"EmailTitle,omitempty" xml:"EmailTitle,omitempty"`
	RobotContent        *string `json:"RobotContent,omitempty" xml:"RobotContent,omitempty"`
	SmsContent          *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	SmsRecoverContent   *string `json:"SmsRecoverContent,omitempty" xml:"SmsRecoverContent,omitempty"`
	TtsContent          *string `json:"TtsContent,omitempty" xml:"TtsContent,omitempty"`
	TtsRecoverContent   *string `json:"TtsRecoverContent,omitempty" xml:"TtsRecoverContent,omitempty"`
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetEmailContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.EmailContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetEmailRecoverContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.EmailRecoverContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetEmailRecoverTitle(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.EmailRecoverTitle = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetEmailTitle(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.EmailTitle = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetRobotContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.RobotContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetSmsContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.SmsContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetSmsRecoverContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.SmsRecoverContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetTtsContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.TtsContent = &v
	return s
}

func (s *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate) SetTtsRecoverContent(v string) *ListNotificationPoliciesResponseBodyPageBeanNotificationPoliciesNotifyTemplate {
	s.TtsRecoverContent = &v
	return s
}

type ListNotificationPoliciesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNotificationPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNotificationPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNotificationPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesResponse) SetHeaders(v map[string]*string) *ListNotificationPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListNotificationPoliciesResponse) SetStatusCode(v int32) *ListNotificationPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNotificationPoliciesResponse) SetBody(v *ListNotificationPoliciesResponseBody) *ListNotificationPoliciesResponse {
	s.Body = v
	return s
}

type ListOnCallSchedulesRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Page *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListOnCallSchedulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOnCallSchedulesRequest) GoString() string {
	return s.String()
}

func (s *ListOnCallSchedulesRequest) SetName(v string) *ListOnCallSchedulesRequest {
	s.Name = &v
	return s
}

func (s *ListOnCallSchedulesRequest) SetPage(v int64) *ListOnCallSchedulesRequest {
	s.Page = &v
	return s
}

func (s *ListOnCallSchedulesRequest) SetSize(v int64) *ListOnCallSchedulesRequest {
	s.Size = &v
	return s
}

type ListOnCallSchedulesResponseBody struct {
	PageBean  *ListOnCallSchedulesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOnCallSchedulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOnCallSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOnCallSchedulesResponseBody) SetPageBean(v *ListOnCallSchedulesResponseBodyPageBean) *ListOnCallSchedulesResponseBody {
	s.PageBean = v
	return s
}

func (s *ListOnCallSchedulesResponseBody) SetRequestId(v string) *ListOnCallSchedulesResponseBody {
	s.RequestId = &v
	return s
}

type ListOnCallSchedulesResponseBodyPageBean struct {
	OnCallSchedules []*ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules `json:"OnCallSchedules,omitempty" xml:"OnCallSchedules,omitempty" type:"Repeated"`
	Page            *int64                                                    `json:"Page,omitempty" xml:"Page,omitempty"`
	Size            *int64                                                    `json:"Size,omitempty" xml:"Size,omitempty"`
	Total           *int64                                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListOnCallSchedulesResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s ListOnCallSchedulesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListOnCallSchedulesResponseBodyPageBean) SetOnCallSchedules(v []*ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) *ListOnCallSchedulesResponseBodyPageBean {
	s.OnCallSchedules = v
	return s
}

func (s *ListOnCallSchedulesResponseBodyPageBean) SetPage(v int64) *ListOnCallSchedulesResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListOnCallSchedulesResponseBodyPageBean) SetSize(v int64) *ListOnCallSchedulesResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListOnCallSchedulesResponseBodyPageBean) SetTotal(v int64) *ListOnCallSchedulesResponseBodyPageBean {
	s.Total = &v
	return s
}

type ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) String() string {
	return tea.Prettify(s)
}

func (s ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) GoString() string {
	return s.String()
}

func (s *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) SetDescription(v string) *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules {
	s.Description = &v
	return s
}

func (s *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) SetId(v int64) *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules {
	s.Id = &v
	return s
}

func (s *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) SetName(v string) *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules {
	s.Name = &v
	return s
}

type ListOnCallSchedulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOnCallSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOnCallSchedulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOnCallSchedulesResponse) GoString() string {
	return s.String()
}

func (s *ListOnCallSchedulesResponse) SetHeaders(v map[string]*string) *ListOnCallSchedulesResponse {
	s.Headers = v
	return s
}

func (s *ListOnCallSchedulesResponse) SetStatusCode(v int32) *ListOnCallSchedulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOnCallSchedulesResponse) SetBody(v *ListOnCallSchedulesResponseBody) *ListOnCallSchedulesResponse {
	s.Body = v
	return s
}

type ListPrometheusAlertRulesRequest struct {
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MatchExpressions *string `json:"MatchExpressions,omitempty" xml:"MatchExpressions,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPrometheusAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesRequest) SetClusterId(v string) *ListPrometheusAlertRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetMatchExpressions(v string) *ListPrometheusAlertRulesRequest {
	s.MatchExpressions = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetName(v string) *ListPrometheusAlertRulesRequest {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetRegionId(v string) *ListPrometheusAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetStatus(v int32) *ListPrometheusAlertRulesRequest {
	s.Status = &v
	return s
}

func (s *ListPrometheusAlertRulesRequest) SetType(v string) *ListPrometheusAlertRulesRequest {
	s.Type = &v
	return s
}

type ListPrometheusAlertRulesResponseBody struct {
	PrometheusAlertRules []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules `json:"PrometheusAlertRules,omitempty" xml:"PrometheusAlertRules,omitempty" type:"Repeated"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBody) SetPrometheusAlertRules(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) *ListPrometheusAlertRulesResponseBody {
	s.PrometheusAlertRules = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBody) SetRequestId(v string) *ListPrometheusAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRules struct {
	AlertId        *int64                                                                 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string                                                                `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	ClusterId      *string                                                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64                                                                 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string                                                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string                                                                `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message        *string                                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string                                                                `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Status         *int32                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetAlertId(v int64) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.AlertId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetAlertName(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.AlertName = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetAnnotations(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Annotations = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetClusterId(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetDispatchRuleId(v int64) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.DispatchRuleId = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetDuration(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Duration = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetExpression(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Expression = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetLabels(v []*ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Labels = v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetMessage(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Message = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetNotifyType(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.NotifyType = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetStatus(v int32) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Status = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules) SetType(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRules {
	s.Type = &v
	return s
}

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) SetName(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations) SetValue(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesAnnotations {
	s.Value = &v
	return s
}

type ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) SetName(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels) SetValue(v string) *ListPrometheusAlertRulesResponseBodyPrometheusAlertRulesLabels {
	s.Value = &v
	return s
}

type ListPrometheusAlertRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPrometheusAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPrometheusAlertRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertRulesResponse) SetHeaders(v map[string]*string) *ListPrometheusAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusAlertRulesResponse) SetStatusCode(v int32) *ListPrometheusAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusAlertRulesResponse) SetBody(v *ListPrometheusAlertRulesResponseBody) *ListPrometheusAlertRulesResponse {
	s.Body = v
	return s
}

type ListPrometheusAlertTemplatesRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPrometheusAlertTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesRequest) SetClusterId(v string) *ListPrometheusAlertTemplatesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusAlertTemplatesRequest) SetRegionId(v string) *ListPrometheusAlertTemplatesRequest {
	s.RegionId = &v
	return s
}

type ListPrometheusAlertTemplatesResponseBody struct {
	PrometheusAlertTemplates []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates `json:"PrometheusAlertTemplates,omitempty" xml:"PrometheusAlertTemplates,omitempty" type:"Repeated"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponseBody) SetPrometheusAlertTemplates(v []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) *ListPrometheusAlertTemplatesResponseBody {
	s.PrometheusAlertTemplates = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBody) SetRequestId(v string) *ListPrometheusAlertTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates struct {
	AlertName   *string                                                                        `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	Description *string                                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration    *string                                                                        `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression  *string                                                                        `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels      []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Type        *string                                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	Version     *string                                                                        `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetAlertName(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.AlertName = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetAnnotations(v []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Annotations = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetDescription(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Description = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetDuration(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Duration = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetExpression(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Expression = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetLabels(v []*ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Labels = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetType(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Type = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates) SetVersion(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplates {
	s.Version = &v
	return s
}

type ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) SetName(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations) SetValue(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesAnnotations {
	s.Value = &v
	return s
}

type ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) SetName(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels {
	s.Name = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels) SetValue(v string) *ListPrometheusAlertTemplatesResponseBodyPrometheusAlertTemplatesLabels {
	s.Value = &v
	return s
}

type ListPrometheusAlertTemplatesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPrometheusAlertTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPrometheusAlertTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusAlertTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesResponse) SetHeaders(v map[string]*string) *ListPrometheusAlertTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusAlertTemplatesResponse) SetStatusCode(v int32) *ListPrometheusAlertTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusAlertTemplatesResponse) SetBody(v *ListPrometheusAlertTemplatesResponseBody) *ListPrometheusAlertTemplatesResponse {
	s.Body = v
	return s
}

type ListPrometheusGlobalViewRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPrometheusGlobalViewRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusGlobalViewRequest) SetRegionId(v string) *ListPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

type ListPrometheusGlobalViewResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusGlobalViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusGlobalViewResponseBody) SetData(v string) *ListPrometheusGlobalViewResponseBody {
	s.Data = &v
	return s
}

func (s *ListPrometheusGlobalViewResponseBody) SetRequestId(v string) *ListPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

type ListPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPrometheusGlobalViewResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *ListPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusGlobalViewResponse) SetStatusCode(v int32) *ListPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusGlobalViewResponse) SetBody(v *ListPrometheusGlobalViewResponseBody) *ListPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

type ListPrometheusInstancesRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ShowGlobalView *bool   `json:"ShowGlobalView,omitempty" xml:"ShowGlobalView,omitempty"`
}

func (s ListPrometheusInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstancesRequest) SetRegionId(v string) *ListPrometheusInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusInstancesRequest) SetShowGlobalView(v bool) *ListPrometheusInstancesRequest {
	s.ShowGlobalView = &v
	return s
}

type ListPrometheusInstancesResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstancesResponseBody) SetData(v string) *ListPrometheusInstancesResponseBody {
	s.Data = &v
	return s
}

func (s *ListPrometheusInstancesResponseBody) SetRequestId(v string) *ListPrometheusInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListPrometheusInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPrometheusInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPrometheusInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrometheusInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstancesResponse) SetHeaders(v map[string]*string) *ListPrometheusInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListPrometheusInstancesResponse) SetStatusCode(v int32) *ListPrometheusInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrometheusInstancesResponse) SetBody(v *ListPrometheusInstancesResponseBody) *ListPrometheusInstancesResponse {
	s.Body = v
	return s
}

type ListRetcodeAppsRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListRetcodeAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRetcodeAppsRequest) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsRequest) SetRegionId(v string) *ListRetcodeAppsRequest {
	s.RegionId = &v
	return s
}

func (s *ListRetcodeAppsRequest) SetSecurityToken(v string) *ListRetcodeAppsRequest {
	s.SecurityToken = &v
	return s
}

type ListRetcodeAppsResponseBody struct {
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetcodeApps []*ListRetcodeAppsResponseBodyRetcodeApps `json:"RetcodeApps,omitempty" xml:"RetcodeApps,omitempty" type:"Repeated"`
}

func (s ListRetcodeAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRetcodeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponseBody) SetRequestId(v string) *ListRetcodeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRetcodeAppsResponseBody) SetRetcodeApps(v []*ListRetcodeAppsResponseBodyRetcodeApps) *ListRetcodeAppsResponseBody {
	s.RetcodeApps = v
	return s
}

type ListRetcodeAppsResponseBodyRetcodeApps struct {
	AppId          *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Pid            *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RetcodeAppType *string `json:"RetcodeAppType,omitempty" xml:"RetcodeAppType,omitempty"`
}

func (s ListRetcodeAppsResponseBodyRetcodeApps) String() string {
	return tea.Prettify(s)
}

func (s ListRetcodeAppsResponseBodyRetcodeApps) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetAppId(v int64) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.AppId = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetAppName(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.AppName = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetPid(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.Pid = &v
	return s
}

func (s *ListRetcodeAppsResponseBodyRetcodeApps) SetRetcodeAppType(v string) *ListRetcodeAppsResponseBodyRetcodeApps {
	s.RetcodeAppType = &v
	return s
}

type ListRetcodeAppsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRetcodeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRetcodeAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRetcodeAppsResponse) GoString() string {
	return s.String()
}

func (s *ListRetcodeAppsResponse) SetHeaders(v map[string]*string) *ListRetcodeAppsResponse {
	s.Headers = v
	return s
}

func (s *ListRetcodeAppsResponse) SetStatusCode(v int32) *ListRetcodeAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRetcodeAppsResponse) SetBody(v *ListRetcodeAppsResponseBody) *ListRetcodeAppsResponse {
	s.Body = v
	return s
}

type ListScenarioRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	Sign     *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
}

func (s ListScenarioRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioRequest) GoString() string {
	return s.String()
}

func (s *ListScenarioRequest) SetAppId(v string) *ListScenarioRequest {
	s.AppId = &v
	return s
}

func (s *ListScenarioRequest) SetName(v string) *ListScenarioRequest {
	s.Name = &v
	return s
}

func (s *ListScenarioRequest) SetRegionId(v string) *ListScenarioRequest {
	s.RegionId = &v
	return s
}

func (s *ListScenarioRequest) SetScenario(v string) *ListScenarioRequest {
	s.Scenario = &v
	return s
}

func (s *ListScenarioRequest) SetSign(v string) *ListScenarioRequest {
	s.Sign = &v
	return s
}

type ListScenarioResponseBody struct {
	ArmsScenarios []*ListScenarioResponseBodyArmsScenarios `json:"ArmsScenarios,omitempty" xml:"ArmsScenarios,omitempty" type:"Repeated"`
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListScenarioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioResponseBody) GoString() string {
	return s.String()
}

func (s *ListScenarioResponseBody) SetArmsScenarios(v []*ListScenarioResponseBodyArmsScenarios) *ListScenarioResponseBody {
	s.ArmsScenarios = v
	return s
}

func (s *ListScenarioResponseBody) SetRequestId(v string) *ListScenarioResponseBody {
	s.RequestId = &v
	return s
}

type ListScenarioResponseBodyArmsScenarios struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Extensions *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Sign       *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListScenarioResponseBodyArmsScenarios) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioResponseBodyArmsScenarios) GoString() string {
	return s.String()
}

func (s *ListScenarioResponseBodyArmsScenarios) SetAppId(v string) *ListScenarioResponseBodyArmsScenarios {
	s.AppId = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetCreateTime(v string) *ListScenarioResponseBodyArmsScenarios {
	s.CreateTime = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetExtensions(v string) *ListScenarioResponseBodyArmsScenarios {
	s.Extensions = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetId(v int64) *ListScenarioResponseBodyArmsScenarios {
	s.Id = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetName(v string) *ListScenarioResponseBodyArmsScenarios {
	s.Name = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetRegionId(v string) *ListScenarioResponseBodyArmsScenarios {
	s.RegionId = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetSign(v string) *ListScenarioResponseBodyArmsScenarios {
	s.Sign = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetUpdateTime(v string) *ListScenarioResponseBodyArmsScenarios {
	s.UpdateTime = &v
	return s
}

func (s *ListScenarioResponseBodyArmsScenarios) SetUserId(v string) *ListScenarioResponseBodyArmsScenarios {
	s.UserId = &v
	return s
}

type ListScenarioResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScenarioResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioResponse) GoString() string {
	return s.String()
}

func (s *ListScenarioResponse) SetHeaders(v map[string]*string) *ListScenarioResponse {
	s.Headers = v
	return s
}

func (s *ListScenarioResponse) SetStatusCode(v int32) *ListScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScenarioResponse) SetBody(v *ListScenarioResponseBody) *ListScenarioResponse {
	s.Body = v
	return s
}

type ListSilencePoliciesRequest struct {
	IsDetail *bool   `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Page     *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Size     *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListSilencePoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSilencePoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesRequest) SetIsDetail(v bool) *ListSilencePoliciesRequest {
	s.IsDetail = &v
	return s
}

func (s *ListSilencePoliciesRequest) SetName(v string) *ListSilencePoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListSilencePoliciesRequest) SetPage(v int64) *ListSilencePoliciesRequest {
	s.Page = &v
	return s
}

func (s *ListSilencePoliciesRequest) SetRegionId(v string) *ListSilencePoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *ListSilencePoliciesRequest) SetSize(v int64) *ListSilencePoliciesRequest {
	s.Size = &v
	return s
}

type ListSilencePoliciesResponseBody struct {
	PageBean  *ListSilencePoliciesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSilencePoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSilencePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesResponseBody) SetPageBean(v *ListSilencePoliciesResponseBodyPageBean) *ListSilencePoliciesResponseBody {
	s.PageBean = v
	return s
}

func (s *ListSilencePoliciesResponseBody) SetRequestId(v string) *ListSilencePoliciesResponseBody {
	s.RequestId = &v
	return s
}

type ListSilencePoliciesResponseBodyPageBean struct {
	Page            *int64                                                    `json:"Page,omitempty" xml:"Page,omitempty"`
	SilencePolicies []*ListSilencePoliciesResponseBodyPageBeanSilencePolicies `json:"SilencePolicies,omitempty" xml:"SilencePolicies,omitempty" type:"Repeated"`
	Size            *int64                                                    `json:"Size,omitempty" xml:"Size,omitempty"`
	Total           *int64                                                    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSilencePoliciesResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s ListSilencePoliciesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesResponseBodyPageBean) SetPage(v int64) *ListSilencePoliciesResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBean) SetSilencePolicies(v []*ListSilencePoliciesResponseBodyPageBeanSilencePolicies) *ListSilencePoliciesResponseBodyPageBean {
	s.SilencePolicies = v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBean) SetSize(v int64) *ListSilencePoliciesResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBean) SetTotal(v int64) *ListSilencePoliciesResponseBodyPageBean {
	s.Total = &v
	return s
}

type ListSilencePoliciesResponseBodyPageBeanSilencePolicies struct {
	Id            *int64                                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	MatchingRules []*ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules `json:"MatchingRules,omitempty" xml:"MatchingRules,omitempty" type:"Repeated"`
	Name          *string                                                                `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListSilencePoliciesResponseBodyPageBeanSilencePolicies) String() string {
	return tea.Prettify(s)
}

func (s ListSilencePoliciesResponseBodyPageBeanSilencePolicies) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) SetId(v int64) *ListSilencePoliciesResponseBodyPageBeanSilencePolicies {
	s.Id = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) SetMatchingRules(v []*ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules) *ListSilencePoliciesResponseBodyPageBeanSilencePolicies {
	s.MatchingRules = v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePolicies) SetName(v string) *ListSilencePoliciesResponseBodyPageBeanSilencePolicies {
	s.Name = &v
	return s
}

type ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules struct {
	MatchingConditions []*ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions `json:"MatchingConditions,omitempty" xml:"MatchingConditions,omitempty" type:"Repeated"`
}

func (s ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules) String() string {
	return tea.Prettify(s)
}

func (s ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules) SetMatchingConditions(v []*ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRules {
	s.MatchingConditions = v
	return s
}

type ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) String() string {
	return tea.Prettify(s)
}

func (s ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) SetKey(v string) *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions {
	s.Key = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) SetOperator(v string) *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions {
	s.Operator = &v
	return s
}

func (s *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions) SetValue(v string) *ListSilencePoliciesResponseBodyPageBeanSilencePoliciesMatchingRulesMatchingConditions {
	s.Value = &v
	return s
}

type ListSilencePoliciesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSilencePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSilencePoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSilencePoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesResponse) SetHeaders(v map[string]*string) *ListSilencePoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListSilencePoliciesResponse) SetStatusCode(v int32) *ListSilencePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSilencePoliciesResponse) SetBody(v *ListSilencePoliciesResponseBody) *ListSilencePoliciesResponse {
	s.Body = v
	return s
}

type ListTraceAppsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListTraceAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTraceAppsRequest) GoString() string {
	return s.String()
}

func (s *ListTraceAppsRequest) SetRegionId(v string) *ListTraceAppsRequest {
	s.RegionId = &v
	return s
}

type ListTraceAppsResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TraceApps []*ListTraceAppsResponseBodyTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
}

func (s ListTraceAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTraceAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponseBody) SetCode(v int32) *ListTraceAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetMessage(v string) *ListTraceAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetRequestId(v string) *ListTraceAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetSuccess(v bool) *ListTraceAppsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTraceAppsResponseBody) SetTraceApps(v []*ListTraceAppsResponseBodyTraceApps) *ListTraceAppsResponseBody {
	s.TraceApps = v
	return s
}

type ListTraceAppsResponseBodyTraceApps struct {
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListTraceAppsResponseBodyTraceApps) String() string {
	return tea.Prettify(s)
}

func (s ListTraceAppsResponseBodyTraceApps) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponseBodyTraceApps) SetAppId(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.AppId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetAppName(v string) *ListTraceAppsResponseBodyTraceApps {
	s.AppName = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetCreateTime(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.CreateTime = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetLabels(v []*string) *ListTraceAppsResponseBodyTraceApps {
	s.Labels = v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetPid(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Pid = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetRegionId(v string) *ListTraceAppsResponseBodyTraceApps {
	s.RegionId = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetShow(v bool) *ListTraceAppsResponseBodyTraceApps {
	s.Show = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetType(v string) *ListTraceAppsResponseBodyTraceApps {
	s.Type = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetUpdateTime(v int64) *ListTraceAppsResponseBodyTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *ListTraceAppsResponseBodyTraceApps) SetUserId(v string) *ListTraceAppsResponseBodyTraceApps {
	s.UserId = &v
	return s
}

type ListTraceAppsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTraceAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTraceAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTraceAppsResponse) GoString() string {
	return s.String()
}

func (s *ListTraceAppsResponse) SetHeaders(v map[string]*string) *ListTraceAppsResponse {
	s.Headers = v
	return s
}

func (s *ListTraceAppsResponse) SetStatusCode(v int32) *ListTraceAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTraceAppsResponse) SetBody(v *ListTraceAppsResponseBody) *ListTraceAppsResponse {
	s.Body = v
	return s
}

type ManageGetRecordingRuleRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	QueryUserId *string `json:"QueryUserId,omitempty" xml:"QueryUserId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ManageGetRecordingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ManageGetRecordingRuleRequest) GoString() string {
	return s.String()
}

func (s *ManageGetRecordingRuleRequest) SetClusterId(v string) *ManageGetRecordingRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *ManageGetRecordingRuleRequest) SetQueryUserId(v string) *ManageGetRecordingRuleRequest {
	s.QueryUserId = &v
	return s
}

func (s *ManageGetRecordingRuleRequest) SetRegionId(v string) *ManageGetRecordingRuleRequest {
	s.RegionId = &v
	return s
}

type ManageGetRecordingRuleResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManageGetRecordingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManageGetRecordingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ManageGetRecordingRuleResponseBody) SetData(v string) *ManageGetRecordingRuleResponseBody {
	s.Data = &v
	return s
}

func (s *ManageGetRecordingRuleResponseBody) SetRequestId(v string) *ManageGetRecordingRuleResponseBody {
	s.RequestId = &v
	return s
}

type ManageGetRecordingRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ManageGetRecordingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ManageGetRecordingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ManageGetRecordingRuleResponse) GoString() string {
	return s.String()
}

func (s *ManageGetRecordingRuleResponse) SetHeaders(v map[string]*string) *ManageGetRecordingRuleResponse {
	s.Headers = v
	return s
}

func (s *ManageGetRecordingRuleResponse) SetStatusCode(v int32) *ManageGetRecordingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageGetRecordingRuleResponse) SetBody(v *ManageGetRecordingRuleResponseBody) *ManageGetRecordingRuleResponse {
	s.Body = v
	return s
}

type ManageRecordingRuleRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	QueryUserId *string `json:"QueryUserId,omitempty" xml:"QueryUserId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RuleYaml    *string `json:"RuleYaml,omitempty" xml:"RuleYaml,omitempty"`
}

func (s ManageRecordingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ManageRecordingRuleRequest) GoString() string {
	return s.String()
}

func (s *ManageRecordingRuleRequest) SetClusterId(v string) *ManageRecordingRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *ManageRecordingRuleRequest) SetQueryUserId(v string) *ManageRecordingRuleRequest {
	s.QueryUserId = &v
	return s
}

func (s *ManageRecordingRuleRequest) SetRegionId(v string) *ManageRecordingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ManageRecordingRuleRequest) SetRuleYaml(v string) *ManageRecordingRuleRequest {
	s.RuleYaml = &v
	return s
}

type ManageRecordingRuleResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManageRecordingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ManageRecordingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ManageRecordingRuleResponseBody) SetData(v string) *ManageRecordingRuleResponseBody {
	s.Data = &v
	return s
}

func (s *ManageRecordingRuleResponseBody) SetRequestId(v string) *ManageRecordingRuleResponseBody {
	s.RequestId = &v
	return s
}

type ManageRecordingRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ManageRecordingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ManageRecordingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ManageRecordingRuleResponse) GoString() string {
	return s.String()
}

func (s *ManageRecordingRuleResponse) SetHeaders(v map[string]*string) *ManageRecordingRuleResponse {
	s.Headers = v
	return s
}

func (s *ManageRecordingRuleResponse) SetStatusCode(v int32) *ManageRecordingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageRecordingRuleResponse) SetBody(v *ManageRecordingRuleResponseBody) *ManageRecordingRuleResponse {
	s.Body = v
	return s
}

type OpenArmsDefaultSLRRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenArmsDefaultSLRRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsDefaultSLRRequest) GoString() string {
	return s.String()
}

func (s *OpenArmsDefaultSLRRequest) SetRegionId(v string) *OpenArmsDefaultSLRRequest {
	s.RegionId = &v
	return s
}

type OpenArmsDefaultSLRResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenArmsDefaultSLRResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsDefaultSLRResponseBody) GoString() string {
	return s.String()
}

func (s *OpenArmsDefaultSLRResponseBody) SetData(v string) *OpenArmsDefaultSLRResponseBody {
	s.Data = &v
	return s
}

func (s *OpenArmsDefaultSLRResponseBody) SetRequestId(v string) *OpenArmsDefaultSLRResponseBody {
	s.RequestId = &v
	return s
}

type OpenArmsDefaultSLRResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenArmsDefaultSLRResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenArmsDefaultSLRResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsDefaultSLRResponse) GoString() string {
	return s.String()
}

func (s *OpenArmsDefaultSLRResponse) SetHeaders(v map[string]*string) *OpenArmsDefaultSLRResponse {
	s.Headers = v
	return s
}

func (s *OpenArmsDefaultSLRResponse) SetStatusCode(v int32) *OpenArmsDefaultSLRResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenArmsDefaultSLRResponse) SetBody(v *OpenArmsDefaultSLRResponseBody) *OpenArmsDefaultSLRResponse {
	s.Body = v
	return s
}

type OpenArmsServiceSecondVersionRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OpenArmsServiceSecondVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsServiceSecondVersionRequest) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceSecondVersionRequest) SetRegionId(v string) *OpenArmsServiceSecondVersionRequest {
	s.RegionId = &v
	return s
}

func (s *OpenArmsServiceSecondVersionRequest) SetType(v string) *OpenArmsServiceSecondVersionRequest {
	s.Type = &v
	return s
}

type OpenArmsServiceSecondVersionResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenArmsServiceSecondVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsServiceSecondVersionResponseBody) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceSecondVersionResponseBody) SetOrderId(v string) *OpenArmsServiceSecondVersionResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenArmsServiceSecondVersionResponseBody) SetRequestId(v string) *OpenArmsServiceSecondVersionResponseBody {
	s.RequestId = &v
	return s
}

type OpenArmsServiceSecondVersionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenArmsServiceSecondVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenArmsServiceSecondVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenArmsServiceSecondVersionResponse) GoString() string {
	return s.String()
}

func (s *OpenArmsServiceSecondVersionResponse) SetHeaders(v map[string]*string) *OpenArmsServiceSecondVersionResponse {
	s.Headers = v
	return s
}

func (s *OpenArmsServiceSecondVersionResponse) SetStatusCode(v int32) *OpenArmsServiceSecondVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenArmsServiceSecondVersionResponse) SetBody(v *OpenArmsServiceSecondVersionResponseBody) *OpenArmsServiceSecondVersionResponse {
	s.Body = v
	return s
}

type OpenVClusterRequest struct {
	ClusterType    *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Length         *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	Product        *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RecreateSwitch *bool   `json:"RecreateSwitch,omitempty" xml:"RecreateSwitch,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenVClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenVClusterRequest) GoString() string {
	return s.String()
}

func (s *OpenVClusterRequest) SetClusterType(v string) *OpenVClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *OpenVClusterRequest) SetLength(v int32) *OpenVClusterRequest {
	s.Length = &v
	return s
}

func (s *OpenVClusterRequest) SetProduct(v string) *OpenVClusterRequest {
	s.Product = &v
	return s
}

func (s *OpenVClusterRequest) SetRecreateSwitch(v bool) *OpenVClusterRequest {
	s.RecreateSwitch = &v
	return s
}

func (s *OpenVClusterRequest) SetRegionId(v string) *OpenVClusterRequest {
	s.RegionId = &v
	return s
}

type OpenVClusterResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenVClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenVClusterResponseBody) GoString() string {
	return s.String()
}

func (s *OpenVClusterResponseBody) SetData(v string) *OpenVClusterResponseBody {
	s.Data = &v
	return s
}

func (s *OpenVClusterResponseBody) SetRequestId(v string) *OpenVClusterResponseBody {
	s.RequestId = &v
	return s
}

type OpenVClusterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenVClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenVClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenVClusterResponse) GoString() string {
	return s.String()
}

func (s *OpenVClusterResponse) SetHeaders(v map[string]*string) *OpenVClusterResponse {
	s.Headers = v
	return s
}

func (s *OpenVClusterResponse) SetStatusCode(v int32) *OpenVClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenVClusterResponse) SetBody(v *OpenVClusterResponseBody) *OpenVClusterResponse {
	s.Body = v
	return s
}

type OpenXtraceDefaultSLRRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenXtraceDefaultSLRRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenXtraceDefaultSLRRequest) GoString() string {
	return s.String()
}

func (s *OpenXtraceDefaultSLRRequest) SetRegionId(v string) *OpenXtraceDefaultSLRRequest {
	s.RegionId = &v
	return s
}

type OpenXtraceDefaultSLRResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenXtraceDefaultSLRResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenXtraceDefaultSLRResponseBody) GoString() string {
	return s.String()
}

func (s *OpenXtraceDefaultSLRResponseBody) SetData(v string) *OpenXtraceDefaultSLRResponseBody {
	s.Data = &v
	return s
}

func (s *OpenXtraceDefaultSLRResponseBody) SetRequestId(v string) *OpenXtraceDefaultSLRResponseBody {
	s.RequestId = &v
	return s
}

type OpenXtraceDefaultSLRResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *OpenXtraceDefaultSLRResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenXtraceDefaultSLRResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenXtraceDefaultSLRResponse) GoString() string {
	return s.String()
}

func (s *OpenXtraceDefaultSLRResponse) SetHeaders(v map[string]*string) *OpenXtraceDefaultSLRResponse {
	s.Headers = v
	return s
}

func (s *OpenXtraceDefaultSLRResponse) SetStatusCode(v int32) *OpenXtraceDefaultSLRResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenXtraceDefaultSLRResponse) SetBody(v *OpenXtraceDefaultSLRResponseBody) *OpenXtraceDefaultSLRResponse {
	s.Body = v
	return s
}

type QueryMetricByPageRequest struct {
	CurrentPage   *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	CustomFilters []*string                          `json:"CustomFilters,omitempty" xml:"CustomFilters,omitempty" type:"Repeated"`
	Dimensions    []*string                          `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	EndTime       *int64                             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filters       []*QueryMetricByPageRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	IntervalInSec *int32                             `json:"IntervalInSec,omitempty" xml:"IntervalInSec,omitempty"`
	Measures      []*string                          `json:"Measures,omitempty" xml:"Measures,omitempty" type:"Repeated"`
	Metric        *string                            `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Order         *string                            `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderBy       *string                            `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	PageSize      *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime     *int64                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryMetricByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricByPageRequest) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageRequest) SetCurrentPage(v int32) *QueryMetricByPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryMetricByPageRequest) SetCustomFilters(v []*string) *QueryMetricByPageRequest {
	s.CustomFilters = v
	return s
}

func (s *QueryMetricByPageRequest) SetDimensions(v []*string) *QueryMetricByPageRequest {
	s.Dimensions = v
	return s
}

func (s *QueryMetricByPageRequest) SetEndTime(v int64) *QueryMetricByPageRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMetricByPageRequest) SetFilters(v []*QueryMetricByPageRequestFilters) *QueryMetricByPageRequest {
	s.Filters = v
	return s
}

func (s *QueryMetricByPageRequest) SetIntervalInSec(v int32) *QueryMetricByPageRequest {
	s.IntervalInSec = &v
	return s
}

func (s *QueryMetricByPageRequest) SetMeasures(v []*string) *QueryMetricByPageRequest {
	s.Measures = v
	return s
}

func (s *QueryMetricByPageRequest) SetMetric(v string) *QueryMetricByPageRequest {
	s.Metric = &v
	return s
}

func (s *QueryMetricByPageRequest) SetOrder(v string) *QueryMetricByPageRequest {
	s.Order = &v
	return s
}

func (s *QueryMetricByPageRequest) SetOrderBy(v string) *QueryMetricByPageRequest {
	s.OrderBy = &v
	return s
}

func (s *QueryMetricByPageRequest) SetPageSize(v int32) *QueryMetricByPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMetricByPageRequest) SetStartTime(v int64) *QueryMetricByPageRequest {
	s.StartTime = &v
	return s
}

type QueryMetricByPageRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMetricByPageRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricByPageRequestFilters) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageRequestFilters) SetKey(v string) *QueryMetricByPageRequestFilters {
	s.Key = &v
	return s
}

func (s *QueryMetricByPageRequestFilters) SetValue(v string) *QueryMetricByPageRequestFilters {
	s.Value = &v
	return s
}

type QueryMetricByPageResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryMetricByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMetricByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageResponseBody) SetCode(v string) *QueryMetricByPageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetData(v *QueryMetricByPageResponseBodyData) *QueryMetricByPageResponseBody {
	s.Data = v
	return s
}

func (s *QueryMetricByPageResponseBody) SetMessage(v string) *QueryMetricByPageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetRequestId(v string) *QueryMetricByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetSuccess(v bool) *QueryMetricByPageResponseBody {
	s.Success = &v
	return s
}

type QueryMetricByPageResponseBodyData struct {
	Items    []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	Page     *int32                   `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int32                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryMetricByPageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricByPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageResponseBodyData) SetItems(v []map[string]interface{}) *QueryMetricByPageResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetPage(v int32) *QueryMetricByPageResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetPageSize(v int32) *QueryMetricByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetTotal(v int32) *QueryMetricByPageResponseBodyData {
	s.Total = &v
	return s
}

type QueryMetricByPageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryMetricByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryMetricByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMetricByPageResponse) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageResponse) SetHeaders(v map[string]*string) *QueryMetricByPageResponse {
	s.Headers = v
	return s
}

func (s *QueryMetricByPageResponse) SetStatusCode(v int32) *QueryMetricByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMetricByPageResponse) SetBody(v *QueryMetricByPageResponseBody) *QueryMetricByPageResponse {
	s.Body = v
	return s
}

type QueryPromInstallStatusRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryPromInstallStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPromInstallStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryPromInstallStatusRequest) SetClusterId(v string) *QueryPromInstallStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryPromInstallStatusRequest) SetRegionId(v string) *QueryPromInstallStatusRequest {
	s.RegionId = &v
	return s
}

type QueryPromInstallStatusResponseBody struct {
	Data      *QueryPromInstallStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPromInstallStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPromInstallStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPromInstallStatusResponseBody) SetData(v *QueryPromInstallStatusResponseBodyData) *QueryPromInstallStatusResponseBody {
	s.Data = v
	return s
}

func (s *QueryPromInstallStatusResponseBody) SetRequestId(v string) *QueryPromInstallStatusResponseBody {
	s.RequestId = &v
	return s
}

type QueryPromInstallStatusResponseBodyData struct {
	IsControllerInstalled *bool `json:"isControllerInstalled,omitempty" xml:"isControllerInstalled,omitempty"`
}

func (s QueryPromInstallStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryPromInstallStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPromInstallStatusResponseBodyData) SetIsControllerInstalled(v bool) *QueryPromInstallStatusResponseBodyData {
	s.IsControllerInstalled = &v
	return s
}

type QueryPromInstallStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryPromInstallStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryPromInstallStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPromInstallStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryPromInstallStatusResponse) SetHeaders(v map[string]*string) *QueryPromInstallStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryPromInstallStatusResponse) SetStatusCode(v int32) *QueryPromInstallStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPromInstallStatusResponse) SetBody(v *QueryPromInstallStatusResponseBody) *QueryPromInstallStatusResponse {
	s.Body = v
	return s
}

type QueryReleaseMetricRequest struct {
	ChangeOrderId    *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	CreateTime       *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MetricType       *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	Pid              *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	ProxyUserId      *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ReleaseEndTime   *int64  `json:"ReleaseEndTime,omitempty" xml:"ReleaseEndTime,omitempty"`
	ReleaseStartTime *int64  `json:"ReleaseStartTime,omitempty" xml:"ReleaseStartTime,omitempty"`
	Service          *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s QueryReleaseMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryReleaseMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryReleaseMetricRequest) SetChangeOrderId(v string) *QueryReleaseMetricRequest {
	s.ChangeOrderId = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetCreateTime(v int64) *QueryReleaseMetricRequest {
	s.CreateTime = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetMetricType(v string) *QueryReleaseMetricRequest {
	s.MetricType = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetPid(v string) *QueryReleaseMetricRequest {
	s.Pid = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetProxyUserId(v string) *QueryReleaseMetricRequest {
	s.ProxyUserId = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetReleaseEndTime(v int64) *QueryReleaseMetricRequest {
	s.ReleaseEndTime = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetReleaseStartTime(v int64) *QueryReleaseMetricRequest {
	s.ReleaseStartTime = &v
	return s
}

func (s *QueryReleaseMetricRequest) SetService(v string) *QueryReleaseMetricRequest {
	s.Service = &v
	return s
}

type QueryReleaseMetricResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryReleaseMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryReleaseMetricResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReleaseMetricResponseBody) SetData(v string) *QueryReleaseMetricResponseBody {
	s.Data = &v
	return s
}

func (s *QueryReleaseMetricResponseBody) SetRequestId(v string) *QueryReleaseMetricResponseBody {
	s.RequestId = &v
	return s
}

type QueryReleaseMetricResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryReleaseMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryReleaseMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryReleaseMetricResponse) GoString() string {
	return s.String()
}

func (s *QueryReleaseMetricResponse) SetHeaders(v map[string]*string) *QueryReleaseMetricResponse {
	s.Headers = v
	return s
}

func (s *QueryReleaseMetricResponse) SetStatusCode(v int32) *QueryReleaseMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReleaseMetricResponse) SetBody(v *QueryReleaseMetricResponseBody) *QueryReleaseMetricResponse {
	s.Body = v
	return s
}

type RemoveAliClusterIdsFromPrometheusGlobalViewRequest struct {
	ClusterIds          *string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty"`
	GlobalViewClusterId *string `json:"GlobalViewClusterId,omitempty" xml:"GlobalViewClusterId,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) SetClusterIds(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewRequest {
	s.ClusterIds = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) SetGlobalViewClusterId(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewRequest {
	s.GlobalViewClusterId = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) SetGroupName(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewRequest {
	s.GroupName = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) SetRegionId(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

type RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody struct {
	Data      *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) SetData(v *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody {
	s.Data = v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) SetRequestId(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

type RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData struct {
	Info    *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Msg     *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) SetInfo(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData {
	s.Info = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) SetMsg(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData {
	s.Msg = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) SetSuccess(v bool) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData {
	s.Success = &v
	return s
}

type RemoveAliClusterIdsFromPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *RemoveAliClusterIdsFromPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponse) SetStatusCode(v int32) *RemoveAliClusterIdsFromPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponse) SetBody(v *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) *RemoveAliClusterIdsFromPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

type RemoveSourcesFromPrometheusGlobalViewRequest struct {
	GlobalViewClusterId *string `json:"GlobalViewClusterId,omitempty" xml:"GlobalViewClusterId,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceNames         *string `json:"SourceNames,omitempty" xml:"SourceNames,omitempty"`
}

func (s RemoveSourcesFromPrometheusGlobalViewRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveSourcesFromPrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) SetGlobalViewClusterId(v string) *RemoveSourcesFromPrometheusGlobalViewRequest {
	s.GlobalViewClusterId = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) SetGroupName(v string) *RemoveSourcesFromPrometheusGlobalViewRequest {
	s.GroupName = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) SetRegionId(v string) *RemoveSourcesFromPrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewRequest) SetSourceNames(v string) *RemoveSourcesFromPrometheusGlobalViewRequest {
	s.SourceNames = &v
	return s
}

type RemoveSourcesFromPrometheusGlobalViewResponseBody struct {
	Data      *RemoveSourcesFromPrometheusGlobalViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveSourcesFromPrometheusGlobalViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSourcesFromPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBody) SetData(v *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) *RemoveSourcesFromPrometheusGlobalViewResponseBody {
	s.Data = v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBody) SetRequestId(v string) *RemoveSourcesFromPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

type RemoveSourcesFromPrometheusGlobalViewResponseBodyData struct {
	Info    *string `json:"Info,omitempty" xml:"Info,omitempty"`
	Msg     *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveSourcesFromPrometheusGlobalViewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RemoveSourcesFromPrometheusGlobalViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) SetInfo(v string) *RemoveSourcesFromPrometheusGlobalViewResponseBodyData {
	s.Info = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) SetMsg(v string) *RemoveSourcesFromPrometheusGlobalViewResponseBodyData {
	s.Msg = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponseBodyData) SetSuccess(v bool) *RemoveSourcesFromPrometheusGlobalViewResponseBodyData {
	s.Success = &v
	return s
}

type RemoveSourcesFromPrometheusGlobalViewResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveSourcesFromPrometheusGlobalViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSourcesFromPrometheusGlobalViewResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSourcesFromPrometheusGlobalViewResponse) GoString() string {
	return s.String()
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponse) SetHeaders(v map[string]*string) *RemoveSourcesFromPrometheusGlobalViewResponse {
	s.Headers = v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponse) SetStatusCode(v int32) *RemoveSourcesFromPrometheusGlobalViewResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSourcesFromPrometheusGlobalViewResponse) SetBody(v *RemoveSourcesFromPrometheusGlobalViewResponseBody) *RemoveSourcesFromPrometheusGlobalViewResponse {
	s.Body = v
	return s
}

type SaveTraceAppConfigRequest struct {
	Pid      *string                              `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Settings []*SaveTraceAppConfigRequestSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Repeated"`
}

func (s SaveTraceAppConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTraceAppConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigRequest) SetPid(v string) *SaveTraceAppConfigRequest {
	s.Pid = &v
	return s
}

func (s *SaveTraceAppConfigRequest) SetSettings(v []*SaveTraceAppConfigRequestSettings) *SaveTraceAppConfigRequest {
	s.Settings = v
	return s
}

type SaveTraceAppConfigRequestSettings struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SaveTraceAppConfigRequestSettings) String() string {
	return tea.Prettify(s)
}

func (s SaveTraceAppConfigRequestSettings) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigRequestSettings) SetKey(v string) *SaveTraceAppConfigRequestSettings {
	s.Key = &v
	return s
}

func (s *SaveTraceAppConfigRequestSettings) SetValue(v string) *SaveTraceAppConfigRequestSettings {
	s.Value = &v
	return s
}

type SaveTraceAppConfigResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveTraceAppConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTraceAppConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigResponseBody) SetData(v string) *SaveTraceAppConfigResponseBody {
	s.Data = &v
	return s
}

func (s *SaveTraceAppConfigResponseBody) SetRequestId(v string) *SaveTraceAppConfigResponseBody {
	s.RequestId = &v
	return s
}

type SaveTraceAppConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveTraceAppConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTraceAppConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTraceAppConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveTraceAppConfigResponse) SetHeaders(v map[string]*string) *SaveTraceAppConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveTraceAppConfigResponse) SetStatusCode(v int32) *SaveTraceAppConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTraceAppConfigResponse) SetBody(v *SaveTraceAppConfigResponseBody) *SaveTraceAppConfigResponse {
	s.Body = v
	return s
}

type SearchAlertContactRequest struct {
	ContactIds  *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertContactRequest) SetContactIds(v string) *SearchAlertContactRequest {
	s.ContactIds = &v
	return s
}

func (s *SearchAlertContactRequest) SetContactName(v string) *SearchAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactRequest) SetCurrentPage(v string) *SearchAlertContactRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertContactRequest) SetEmail(v string) *SearchAlertContactRequest {
	s.Email = &v
	return s
}

func (s *SearchAlertContactRequest) SetPageSize(v string) *SearchAlertContactRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertContactRequest) SetPhone(v string) *SearchAlertContactRequest {
	s.Phone = &v
	return s
}

func (s *SearchAlertContactRequest) SetRegionId(v string) *SearchAlertContactRequest {
	s.RegionId = &v
	return s
}

type SearchAlertContactResponseBody struct {
	PageBean  *SearchAlertContactResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponseBody) SetPageBean(v *SearchAlertContactResponseBodyPageBean) *SearchAlertContactResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchAlertContactResponseBody) SetRequestId(v string) *SearchAlertContactResponseBody {
	s.RequestId = &v
	return s
}

type SearchAlertContactResponseBodyPageBean struct {
	Contacts   []*SearchAlertContactResponseBodyPageBeanContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAlertContactResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponseBodyPageBean) SetContacts(v []*SearchAlertContactResponseBodyPageBeanContacts) *SearchAlertContactResponseBodyPageBean {
	s.Contacts = v
	return s
}

func (s *SearchAlertContactResponseBodyPageBean) SetPageNumber(v int32) *SearchAlertContactResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBean) SetPageSize(v int32) *SearchAlertContactResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBean) SetTotalCount(v int32) *SearchAlertContactResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

type SearchAlertContactResponseBodyPageBeanContacts struct {
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DingRobot   *string `json:"DingRobot,omitempty" xml:"DingRobot,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	SystemNoc   *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Webhook     *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s SearchAlertContactResponseBodyPageBeanContacts) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactResponseBodyPageBeanContacts) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetContactId(v int64) *SearchAlertContactResponseBodyPageBeanContacts {
	s.ContactId = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetContactName(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetContent(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Content = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetCreateTime(v int64) *SearchAlertContactResponseBodyPageBeanContacts {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetDingRobot(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.DingRobot = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetEmail(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Email = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetPhone(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Phone = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetSystemNoc(v bool) *SearchAlertContactResponseBodyPageBeanContacts {
	s.SystemNoc = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetUpdateTime(v int64) *SearchAlertContactResponseBodyPageBeanContacts {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetUserId(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.UserId = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetWebhook(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Webhook = &v
	return s
}

type SearchAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchAlertContactResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponse) SetHeaders(v map[string]*string) *SearchAlertContactResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertContactResponse) SetStatusCode(v int32) *SearchAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertContactResponse) SetBody(v *SearchAlertContactResponseBody) *SearchAlertContactResponse {
	s.Body = v
	return s
}

type SearchAlertContactGroupRequest struct {
	ContactGroupIds  *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	ContactId        *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName      *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	IsDetail         *bool   `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupRequest) SetContactGroupIds(v string) *SearchAlertContactGroupRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetContactGroupName(v string) *SearchAlertContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetContactId(v int64) *SearchAlertContactGroupRequest {
	s.ContactId = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetContactName(v string) *SearchAlertContactGroupRequest {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetIsDetail(v bool) *SearchAlertContactGroupRequest {
	s.IsDetail = &v
	return s
}

func (s *SearchAlertContactGroupRequest) SetRegionId(v string) *SearchAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

type SearchAlertContactGroupResponseBody struct {
	ContactGroups []*SearchAlertContactGroupResponseBodyContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Repeated"`
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponseBody) SetContactGroups(v []*SearchAlertContactGroupResponseBodyContactGroups) *SearchAlertContactGroupResponseBody {
	s.ContactGroups = v
	return s
}

func (s *SearchAlertContactGroupResponseBody) SetRequestId(v string) *SearchAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

type SearchAlertContactGroupResponseBodyContactGroups struct {
	ContactGroupId   *int64                                                      `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	ContactGroupName *string                                                     `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	Contacts         []*SearchAlertContactGroupResponseBodyContactGroupsContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	CreateTime       *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UpdateTime       *int64                                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId           *string                                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertContactGroupResponseBodyContactGroups) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupResponseBodyContactGroups) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetContactGroupId(v int64) *SearchAlertContactGroupResponseBodyContactGroups {
	s.ContactGroupId = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetContactGroupName(v string) *SearchAlertContactGroupResponseBodyContactGroups {
	s.ContactGroupName = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetContacts(v []*SearchAlertContactGroupResponseBodyContactGroupsContacts) *SearchAlertContactGroupResponseBodyContactGroups {
	s.Contacts = v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetCreateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroups {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetUpdateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroups {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetUserId(v string) *SearchAlertContactGroupResponseBodyContactGroups {
	s.UserId = &v
	return s
}

type SearchAlertContactGroupResponseBodyContactGroupsContacts struct {
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DingRobot   *string `json:"DingRobot,omitempty" xml:"DingRobot,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	SystemNoc   *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertContactGroupResponseBodyContactGroupsContacts) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupResponseBodyContactGroupsContacts) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetContactId(v int64) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.ContactId = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetContactName(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetCreateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetDingRobot(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.DingRobot = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetEmail(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.Email = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetPhone(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.Phone = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetSystemNoc(v bool) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.SystemNoc = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetUpdateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetUserId(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.UserId = &v
	return s
}

type SearchAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchAlertContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponse) SetHeaders(v map[string]*string) *SearchAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertContactGroupResponse) SetStatusCode(v int32) *SearchAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertContactGroupResponse) SetBody(v *SearchAlertContactGroupResponseBody) *SearchAlertContactGroupResponse {
	s.Body = v
	return s
}

type SearchAlertHistoriesRequest struct {
	AlertId     *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertType   *int32  `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SearchAlertHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertHistoriesRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesRequest) SetAlertId(v int64) *SearchAlertHistoriesRequest {
	s.AlertId = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetAlertType(v int32) *SearchAlertHistoriesRequest {
	s.AlertType = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetCurrentPage(v int32) *SearchAlertHistoriesRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetEndTime(v int64) *SearchAlertHistoriesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetPageSize(v int32) *SearchAlertHistoriesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetRegionId(v string) *SearchAlertHistoriesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetStartTime(v int64) *SearchAlertHistoriesRequest {
	s.StartTime = &v
	return s
}

type SearchAlertHistoriesResponseBody struct {
	PageBean  *SearchAlertHistoriesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponseBody) SetPageBean(v *SearchAlertHistoriesResponseBodyPageBean) *SearchAlertHistoriesResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchAlertHistoriesResponseBody) SetRequestId(v string) *SearchAlertHistoriesResponseBody {
	s.RequestId = &v
	return s
}

type SearchAlertHistoriesResponseBodyPageBean struct {
	AlarmHistories []*SearchAlertHistoriesResponseBodyPageBeanAlarmHistories `json:"AlarmHistories,omitempty" xml:"AlarmHistories,omitempty" type:"Repeated"`
	PageNumber     *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAlertHistoriesResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertHistoriesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponseBodyPageBean) SetAlarmHistories(v []*SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) *SearchAlertHistoriesResponseBodyPageBean {
	s.AlarmHistories = v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBean) SetPageNumber(v int32) *SearchAlertHistoriesResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBean) SetPageSize(v int32) *SearchAlertHistoriesResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBean) SetTotalCount(v int32) *SearchAlertHistoriesResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

type SearchAlertHistoriesResponseBodyPageBeanAlarmHistories struct {
	AlarmContent      *string `json:"AlarmContent,omitempty" xml:"AlarmContent,omitempty"`
	AlarmResponseCode *int32  `json:"AlarmResponseCode,omitempty" xml:"AlarmResponseCode,omitempty"`
	AlarmSources      *string `json:"AlarmSources,omitempty" xml:"AlarmSources,omitempty"`
	AlarmTime         *int64  `json:"AlarmTime,omitempty" xml:"AlarmTime,omitempty"`
	AlarmType         *int32  `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	Emails            *string `json:"Emails,omitempty" xml:"Emails,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Phones            *string `json:"Phones,omitempty" xml:"Phones,omitempty"`
	StrategyId        *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	Target            *string `json:"Target,omitempty" xml:"Target,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmContent(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmContent = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmResponseCode(v int32) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmResponseCode = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmSources(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmSources = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmTime(v int64) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmTime = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetAlarmType(v int32) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.AlarmType = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetEmails(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Emails = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetId(v int64) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Id = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetPhones(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Phones = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetStrategyId(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.StrategyId = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetTarget(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.Target = &v
	return s
}

func (s *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories) SetUserId(v string) *SearchAlertHistoriesResponseBodyPageBeanAlarmHistories {
	s.UserId = &v
	return s
}

type SearchAlertHistoriesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchAlertHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchAlertHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertHistoriesResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesResponse) SetHeaders(v map[string]*string) *SearchAlertHistoriesResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertHistoriesResponse) SetStatusCode(v int32) *SearchAlertHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertHistoriesResponse) SetBody(v *SearchAlertHistoriesResponseBody) *SearchAlertHistoriesResponse {
	s.Body = v
	return s
}

type SearchAlertRulesRequest struct {
	AppType        *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pid            *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SystemRegionId *string `json:"SystemRegionId,omitempty" xml:"SystemRegionId,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchAlertRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesRequest) SetAppType(v string) *SearchAlertRulesRequest {
	s.AppType = &v
	return s
}

func (s *SearchAlertRulesRequest) SetCurrentPage(v int32) *SearchAlertRulesRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertRulesRequest) SetPageSize(v int32) *SearchAlertRulesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertRulesRequest) SetPid(v string) *SearchAlertRulesRequest {
	s.Pid = &v
	return s
}

func (s *SearchAlertRulesRequest) SetRegionId(v string) *SearchAlertRulesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchAlertRulesRequest) SetSystemRegionId(v string) *SearchAlertRulesRequest {
	s.SystemRegionId = &v
	return s
}

func (s *SearchAlertRulesRequest) SetTitle(v string) *SearchAlertRulesRequest {
	s.Title = &v
	return s
}

func (s *SearchAlertRulesRequest) SetType(v string) *SearchAlertRulesRequest {
	s.Type = &v
	return s
}

type SearchAlertRulesResponseBody struct {
	PageBean  *SearchAlertRulesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBody) SetPageBean(v *SearchAlertRulesResponseBodyPageBean) *SearchAlertRulesResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchAlertRulesResponseBody) SetRequestId(v string) *SearchAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

type SearchAlertRulesResponseBodyPageBean struct {
	AlertRules []*SearchAlertRulesResponseBodyPageBeanAlertRules `json:"AlertRules,omitempty" xml:"AlertRules,omitempty" type:"Repeated"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBean) SetAlertRules(v []*SearchAlertRulesResponseBodyPageBeanAlertRules) *SearchAlertRulesResponseBodyPageBean {
	s.AlertRules = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) SetPageNumber(v int32) *SearchAlertRulesResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) SetPageSize(v int32) *SearchAlertRulesResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBean) SetTotalCount(v int32) *SearchAlertRulesResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRules struct {
	AlarmContext       *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext `json:"AlarmContext,omitempty" xml:"AlarmContext,omitempty" type:"Struct"`
	AlertLevel         *string                                                     `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	AlertRule          *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule    `json:"AlertRule,omitempty" xml:"AlertRule,omitempty" type:"Struct"`
	AlertTitle         *string                                                     `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	AlertType          *int32                                                      `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertVersion       *int32                                                      `json:"AlertVersion,omitempty" xml:"AlertVersion,omitempty"`
	AlertWays          []*string                                                   `json:"AlertWays,omitempty" xml:"AlertWays,omitempty" type:"Repeated"`
	Config             *string                                                     `json:"Config,omitempty" xml:"Config,omitempty"`
	ContactGroupIdList *string                                                     `json:"ContactGroupIdList,omitempty" xml:"ContactGroupIdList,omitempty"`
	ContactGroupIds    *string                                                     `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	CreateTime         *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	HostByAlertManager *bool                                                       `json:"HostByAlertManager,omitempty" xml:"HostByAlertManager,omitempty"`
	Id                 *int64                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	MetricParam        *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam  `json:"MetricParam,omitempty" xml:"MetricParam,omitempty" type:"Struct"`
	Notice             *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice       `json:"Notice,omitempty" xml:"Notice,omitempty" type:"Struct"`
	RegionId           *string                                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status             *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId             *int64                                                      `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus         *string                                                     `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	Title              *string                                                     `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTime         *int64                                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId             *string                                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRules) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRules) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlarmContext(v *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlarmContext = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertLevel(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertLevel = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertRule(v *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertRule = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertTitle = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertType(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertType = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertVersion(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertVersion = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetAlertWays(v []*string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.AlertWays = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetConfig(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Config = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetContactGroupIdList(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.ContactGroupIdList = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetContactGroupIds(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.ContactGroupIds = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetCreateTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetHostByAlertManager(v bool) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.HostByAlertManager = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetId(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Id = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetMetricParam(v *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.MetricParam = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetNotice(v *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Notice = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetRegionId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.RegionId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetStatus(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Status = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTaskId(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.TaskId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTaskStatus(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.TaskStatus = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.Title = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetUpdateTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRules) SetUserId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRules {
	s.UserId = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext struct {
	AlarmContentSubTitle *string `json:"AlarmContentSubTitle,omitempty" xml:"AlarmContentSubTitle,omitempty"`
	AlarmContentTemplate *string `json:"AlarmContentTemplate,omitempty" xml:"AlarmContentTemplate,omitempty"`
	Content              *string `json:"Content,omitempty" xml:"Content,omitempty"`
	SubTitle             *string `json:"SubTitle,omitempty" xml:"SubTitle,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetAlarmContentSubTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.AlarmContentSubTitle = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetAlarmContentTemplate(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.AlarmContentTemplate = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetContent(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.Content = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext) SetSubTitle(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlarmContext {
	s.SubTitle = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule struct {
	Operator *string                                                         `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Rules    []*SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) SetOperator(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule {
	s.Operator = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule) SetRules(v []*SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRule {
	s.Rules = v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules struct {
	Aggregates *string  `json:"Aggregates,omitempty" xml:"Aggregates,omitempty"`
	Alias      *string  `json:"Alias,omitempty" xml:"Alias,omitempty"`
	Measure    *string  `json:"Measure,omitempty" xml:"Measure,omitempty"`
	NValue     *int32   `json:"NValue,omitempty" xml:"NValue,omitempty"`
	Operator   *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value      *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetAggregates(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Aggregates = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetAlias(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Alias = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetMeasure(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Measure = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetNValue(v int32) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.NValue = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetOperator(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Operator = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules) SetValue(v float32) *SearchAlertRulesResponseBodyPageBeanAlertRulesAlertRuleRules {
	s.Value = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam struct {
	AppGroupId *string                                                                `json:"AppGroupId,omitempty" xml:"AppGroupId,omitempty"`
	AppId      *string                                                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Dimensions []*SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions `json:"Dimensions,omitempty" xml:"Dimensions,omitempty" type:"Repeated"`
	Pid        *string                                                                `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Type       *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetAppGroupId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.AppGroupId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetAppId(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.AppId = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetDimensions(v []*SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Dimensions = v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetPid(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Pid = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam) SetType(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParam {
	s.Type = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) SetKey(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	s.Key = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) SetType(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	s.Type = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions) SetValue(v string) *SearchAlertRulesResponseBodyPageBeanAlertRulesMetricParamDimensions {
	s.Value = &v
	return s
}

type SearchAlertRulesResponseBodyPageBeanAlertRulesNotice struct {
	EndTime         *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	NoticeEndTime   *int64 `json:"NoticeEndTime,omitempty" xml:"NoticeEndTime,omitempty"`
	NoticeStartTime *int64 `json:"NoticeStartTime,omitempty" xml:"NoticeStartTime,omitempty"`
	StartTime       *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetEndTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.EndTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetNoticeEndTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.NoticeEndTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetNoticeStartTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.NoticeStartTime = &v
	return s
}

func (s *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice) SetStartTime(v int64) *SearchAlertRulesResponseBodyPageBeanAlertRulesNotice {
	s.StartTime = &v
	return s
}

type SearchAlertRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchAlertRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *SearchAlertRulesResponse) SetHeaders(v map[string]*string) *SearchAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *SearchAlertRulesResponse) SetStatusCode(v int32) *SearchAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAlertRulesResponse) SetBody(v *SearchAlertRulesResponseBody) *SearchAlertRulesResponse {
	s.Body = v
	return s
}

type SearchEventsRequest struct {
	AlertId     *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertType   *int32  `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AppType     *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsTrigger   *int32  `json:"IsTrigger,omitempty" xml:"IsTrigger,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SearchEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsRequest) GoString() string {
	return s.String()
}

func (s *SearchEventsRequest) SetAlertId(v int64) *SearchEventsRequest {
	s.AlertId = &v
	return s
}

func (s *SearchEventsRequest) SetAlertType(v int32) *SearchEventsRequest {
	s.AlertType = &v
	return s
}

func (s *SearchEventsRequest) SetAppType(v string) *SearchEventsRequest {
	s.AppType = &v
	return s
}

func (s *SearchEventsRequest) SetCurrentPage(v int32) *SearchEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchEventsRequest) SetEndTime(v int64) *SearchEventsRequest {
	s.EndTime = &v
	return s
}

func (s *SearchEventsRequest) SetIsTrigger(v int32) *SearchEventsRequest {
	s.IsTrigger = &v
	return s
}

func (s *SearchEventsRequest) SetPageSize(v int32) *SearchEventsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchEventsRequest) SetPid(v string) *SearchEventsRequest {
	s.Pid = &v
	return s
}

func (s *SearchEventsRequest) SetRegionId(v string) *SearchEventsRequest {
	s.RegionId = &v
	return s
}

func (s *SearchEventsRequest) SetStartTime(v int64) *SearchEventsRequest {
	s.StartTime = &v
	return s
}

type SearchEventsResponseBody struct {
	IsTrigger *int32                            `json:"IsTrigger,omitempty" xml:"IsTrigger,omitempty"`
	PageBean  *SearchEventsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchEventsResponseBody) SetIsTrigger(v int32) *SearchEventsResponseBody {
	s.IsTrigger = &v
	return s
}

func (s *SearchEventsResponseBody) SetPageBean(v *SearchEventsResponseBodyPageBean) *SearchEventsResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchEventsResponseBody) SetRequestId(v string) *SearchEventsResponseBody {
	s.RequestId = &v
	return s
}

type SearchEventsResponseBodyPageBean struct {
	Event      []*SearchEventsResponseBodyPageBeanEvent `json:"Event,omitempty" xml:"Event,omitempty" type:"Repeated"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchEventsResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchEventsResponseBodyPageBean) SetEvent(v []*SearchEventsResponseBodyPageBeanEvent) *SearchEventsResponseBodyPageBean {
	s.Event = v
	return s
}

func (s *SearchEventsResponseBodyPageBean) SetPageNumber(v int32) *SearchEventsResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchEventsResponseBodyPageBean) SetPageSize(v int32) *SearchEventsResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchEventsResponseBodyPageBean) SetTotalCount(v int32) *SearchEventsResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

type SearchEventsResponseBodyPageBeanEvent struct {
	AlertId    *int64    `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName  *string   `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	AlertRule  *string   `json:"AlertRule,omitempty" xml:"AlertRule,omitempty"`
	AlertType  *int32    `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	EventLevel *string   `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	EventTime  *int64    `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	Id         *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Links      []*string `json:"Links,omitempty" xml:"Links,omitempty" type:"Repeated"`
	Message    *string   `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SearchEventsResponseBodyPageBeanEvent) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsResponseBodyPageBeanEvent) GoString() string {
	return s.String()
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertId(v int64) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertId = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertName(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertName = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertRule(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertRule = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetAlertType(v int32) *SearchEventsResponseBodyPageBeanEvent {
	s.AlertType = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetEventLevel(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.EventLevel = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetEventTime(v int64) *SearchEventsResponseBodyPageBeanEvent {
	s.EventTime = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetId(v int64) *SearchEventsResponseBodyPageBeanEvent {
	s.Id = &v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetLinks(v []*string) *SearchEventsResponseBodyPageBeanEvent {
	s.Links = v
	return s
}

func (s *SearchEventsResponseBodyPageBeanEvent) SetMessage(v string) *SearchEventsResponseBodyPageBeanEvent {
	s.Message = &v
	return s
}

type SearchEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchEventsResponse) GoString() string {
	return s.String()
}

func (s *SearchEventsResponse) SetHeaders(v map[string]*string) *SearchEventsResponse {
	s.Headers = v
	return s
}

func (s *SearchEventsResponse) SetStatusCode(v int32) *SearchEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchEventsResponse) SetBody(v *SearchEventsResponseBody) *SearchEventsResponse {
	s.Body = v
	return s
}

type SearchRetcodeAppByPageRequest struct {
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RetcodeAppName *string `json:"RetcodeAppName,omitempty" xml:"RetcodeAppName,omitempty"`
}

func (s SearchRetcodeAppByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchRetcodeAppByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageRequest) SetPageNumber(v int32) *SearchRetcodeAppByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetPageSize(v int32) *SearchRetcodeAppByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetRegionId(v string) *SearchRetcodeAppByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchRetcodeAppByPageRequest) SetRetcodeAppName(v string) *SearchRetcodeAppByPageRequest {
	s.RetcodeAppName = &v
	return s
}

type SearchRetcodeAppByPageResponseBody struct {
	PageBean  *SearchRetcodeAppByPageResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBody) SetPageBean(v *SearchRetcodeAppByPageResponseBodyPageBean) *SearchRetcodeAppByPageResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchRetcodeAppByPageResponseBody) SetRequestId(v string) *SearchRetcodeAppByPageResponseBody {
	s.RequestId = &v
	return s
}

type SearchRetcodeAppByPageResponseBodyPageBean struct {
	PageNumber  *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RetcodeApps []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps `json:"RetcodeApps,omitempty" xml:"RetcodeApps,omitempty" type:"Repeated"`
	TotalCount  *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetPageNumber(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetPageSize(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetRetcodeApps(v []*SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.RetcodeApps = v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBean) SetTotalCount(v int32) *SearchRetcodeAppByPageResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

type SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps struct {
	AppId          *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime     *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Pid            *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RetcodeAppType *string `json:"RetcodeAppType,omitempty" xml:"RetcodeAppType,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime     *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) String() string {
	return tea.Prettify(s)
}

func (s SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetAppId(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.AppId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetAppName(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.AppName = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetCreateTime(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.CreateTime = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetPid(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.Pid = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetRegionId(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.RegionId = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetRetcodeAppType(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.RetcodeAppType = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetType(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.Type = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetUpdateTime(v int64) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps) SetUserId(v string) *SearchRetcodeAppByPageResponseBodyPageBeanRetcodeApps {
	s.UserId = &v
	return s
}

type SearchRetcodeAppByPageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchRetcodeAppByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchRetcodeAppByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchRetcodeAppByPageResponse) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponse) SetHeaders(v map[string]*string) *SearchRetcodeAppByPageResponse {
	s.Headers = v
	return s
}

func (s *SearchRetcodeAppByPageResponse) SetStatusCode(v int32) *SearchRetcodeAppByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchRetcodeAppByPageResponse) SetBody(v *SearchRetcodeAppByPageResponseBody) *SearchRetcodeAppByPageResponse {
	s.Body = v
	return s
}

type SearchTraceAppByNameRequest struct {
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TraceAppName *string `json:"TraceAppName,omitempty" xml:"TraceAppName,omitempty"`
}

func (s SearchTraceAppByNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByNameRequest) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameRequest) SetRegionId(v string) *SearchTraceAppByNameRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByNameRequest) SetTraceAppName(v string) *SearchTraceAppByNameRequest {
	s.TraceAppName = &v
	return s
}

type SearchTraceAppByNameResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceApps []*SearchTraceAppByNameResponseBodyTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
}

func (s SearchTraceAppByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByNameResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponseBody) SetRequestId(v string) *SearchTraceAppByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBody) SetTraceApps(v []*SearchTraceAppByNameResponseBodyTraceApps) *SearchTraceAppByNameResponseBody {
	s.TraceApps = v
	return s
}

type SearchTraceAppByNameResponseBodyTraceApps struct {
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchTraceAppByNameResponseBodyTraceApps) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByNameResponseBodyTraceApps) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetAppId(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.AppId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetAppName(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.AppName = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetCreateTime(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.CreateTime = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetLabels(v []*string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Labels = v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetPid(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Pid = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetRegionId(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetShow(v bool) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Show = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetType(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.Type = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetUpdateTime(v int64) *SearchTraceAppByNameResponseBodyTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchTraceAppByNameResponseBodyTraceApps) SetUserId(v string) *SearchTraceAppByNameResponseBodyTraceApps {
	s.UserId = &v
	return s
}

type SearchTraceAppByNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTraceAppByNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTraceAppByNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByNameResponse) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByNameResponse) SetHeaders(v map[string]*string) *SearchTraceAppByNameResponse {
	s.Headers = v
	return s
}

func (s *SearchTraceAppByNameResponse) SetStatusCode(v int32) *SearchTraceAppByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTraceAppByNameResponse) SetBody(v *SearchTraceAppByNameResponseBody) *SearchTraceAppByNameResponse {
	s.Body = v
	return s
}

type SearchTraceAppByPageRequest struct {
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TraceAppName *string `json:"TraceAppName,omitempty" xml:"TraceAppName,omitempty"`
}

func (s SearchTraceAppByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageRequest) SetPageNumber(v int32) *SearchTraceAppByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetPageSize(v int32) *SearchTraceAppByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetRegionId(v string) *SearchTraceAppByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByPageRequest) SetTraceAppName(v string) *SearchTraceAppByPageRequest {
	s.TraceAppName = &v
	return s
}

type SearchTraceAppByPageResponseBody struct {
	PageBean  *SearchTraceAppByPageResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTraceAppByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBody) SetPageBean(v *SearchTraceAppByPageResponseBodyPageBean) *SearchTraceAppByPageResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchTraceAppByPageResponseBody) SetRequestId(v string) *SearchTraceAppByPageResponseBody {
	s.RequestId = &v
	return s
}

type SearchTraceAppByPageResponseBodyPageBean struct {
	PageNumber *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TraceApps  []*SearchTraceAppByPageResponseBodyPageBeanTraceApps `json:"TraceApps,omitempty" xml:"TraceApps,omitempty" type:"Repeated"`
}

func (s SearchTraceAppByPageResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetPageNumber(v int32) *SearchTraceAppByPageResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetPageSize(v int32) *SearchTraceAppByPageResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetTotalCount(v int32) *SearchTraceAppByPageResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBean) SetTraceApps(v []*SearchTraceAppByPageResponseBodyPageBeanTraceApps) *SearchTraceAppByPageResponseBodyPageBean {
	s.TraceApps = v
	return s
}

type SearchTraceAppByPageResponseBodyPageBeanTraceApps struct {
	AppId      *int64    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName    *string   `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime *int64    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Labels     []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Pid        *string   `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Show       *bool     `json:"Show,omitempty" xml:"Show,omitempty"`
	Type       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime *int64    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchTraceAppByPageResponseBodyPageBeanTraceApps) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByPageResponseBodyPageBeanTraceApps) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetAppId(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.AppId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetAppName(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.AppName = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetCreateTime(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.CreateTime = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetLabels(v []*string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Labels = v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetPid(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Pid = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetRegionId(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.RegionId = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetShow(v bool) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Show = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetType(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.Type = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetUpdateTime(v int64) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.UpdateTime = &v
	return s
}

func (s *SearchTraceAppByPageResponseBodyPageBeanTraceApps) SetUserId(v string) *SearchTraceAppByPageResponseBodyPageBeanTraceApps {
	s.UserId = &v
	return s
}

type SearchTraceAppByPageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTraceAppByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTraceAppByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTraceAppByPageResponse) GoString() string {
	return s.String()
}

func (s *SearchTraceAppByPageResponse) SetHeaders(v map[string]*string) *SearchTraceAppByPageResponse {
	s.Headers = v
	return s
}

func (s *SearchTraceAppByPageResponse) SetStatusCode(v int32) *SearchTraceAppByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTraceAppByPageResponse) SetBody(v *SearchTraceAppByPageResponseBody) *SearchTraceAppByPageResponse {
	s.Body = v
	return s
}

type SearchTracesRequest struct {
	EndTime          *int64                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExclusionFilters []*SearchTracesRequestExclusionFilters `json:"ExclusionFilters,omitempty" xml:"ExclusionFilters,omitempty" type:"Repeated"`
	MinDuration      *int64                                 `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	OperationName    *string                                `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	Pid              *string                                `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId         *string                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Reverse          *bool                                  `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ServiceIp        *string                                `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName      *string                                `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime        *int64                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tag              []*SearchTracesRequestTag              `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s SearchTracesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesRequest) SetEndTime(v int64) *SearchTracesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesRequest) SetExclusionFilters(v []*SearchTracesRequestExclusionFilters) *SearchTracesRequest {
	s.ExclusionFilters = v
	return s
}

func (s *SearchTracesRequest) SetMinDuration(v int64) *SearchTracesRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesRequest) SetOperationName(v string) *SearchTracesRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesRequest) SetPid(v string) *SearchTracesRequest {
	s.Pid = &v
	return s
}

func (s *SearchTracesRequest) SetRegionId(v string) *SearchTracesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesRequest) SetReverse(v bool) *SearchTracesRequest {
	s.Reverse = &v
	return s
}

func (s *SearchTracesRequest) SetServiceIp(v string) *SearchTracesRequest {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesRequest) SetServiceName(v string) *SearchTracesRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesRequest) SetStartTime(v int64) *SearchTracesRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesRequest) SetTag(v []*SearchTracesRequestTag) *SearchTracesRequest {
	s.Tag = v
	return s
}

type SearchTracesRequestExclusionFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesRequestExclusionFilters) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequestExclusionFilters) GoString() string {
	return s.String()
}

func (s *SearchTracesRequestExclusionFilters) SetKey(v string) *SearchTracesRequestExclusionFilters {
	s.Key = &v
	return s
}

func (s *SearchTracesRequestExclusionFilters) SetValue(v string) *SearchTracesRequestExclusionFilters {
	s.Value = &v
	return s
}

type SearchTracesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesRequestTag) GoString() string {
	return s.String()
}

func (s *SearchTracesRequestTag) SetKey(v string) *SearchTracesRequestTag {
	s.Key = &v
	return s
}

func (s *SearchTracesRequestTag) SetValue(v string) *SearchTracesRequestTag {
	s.Value = &v
	return s
}

type SearchTracesResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceInfos []*SearchTracesResponseBodyTraceInfos `json:"TraceInfos,omitempty" xml:"TraceInfos,omitempty" type:"Repeated"`
}

func (s SearchTracesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBody) SetRequestId(v string) *SearchTracesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTracesResponseBody) SetTraceInfos(v []*SearchTracesResponseBodyTraceInfos) *SearchTracesResponseBody {
	s.TraceInfos = v
	return s
}

type SearchTracesResponseBodyTraceInfos struct {
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ServiceIp     *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s SearchTracesResponseBodyTraceInfos) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponseBodyTraceInfos) GoString() string {
	return s.String()
}

func (s *SearchTracesResponseBodyTraceInfos) SetDuration(v int64) *SearchTracesResponseBodyTraceInfos {
	s.Duration = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetOperationName(v string) *SearchTracesResponseBodyTraceInfos {
	s.OperationName = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetServiceIp(v string) *SearchTracesResponseBodyTraceInfos {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetServiceName(v string) *SearchTracesResponseBodyTraceInfos {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetTimestamp(v int64) *SearchTracesResponseBodyTraceInfos {
	s.Timestamp = &v
	return s
}

func (s *SearchTracesResponseBodyTraceInfos) SetTraceID(v string) *SearchTracesResponseBodyTraceInfos {
	s.TraceID = &v
	return s
}

type SearchTracesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTracesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTracesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesResponse) GoString() string {
	return s.String()
}

func (s *SearchTracesResponse) SetHeaders(v map[string]*string) *SearchTracesResponse {
	s.Headers = v
	return s
}

func (s *SearchTracesResponse) SetStatusCode(v int32) *SearchTracesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTracesResponse) SetBody(v *SearchTracesResponseBody) *SearchTracesResponse {
	s.Body = v
	return s
}

type SearchTracesByPageRequest struct {
	EndTime          *int64                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExclusionFilters []*SearchTracesByPageRequestExclusionFilters `json:"ExclusionFilters,omitempty" xml:"ExclusionFilters,omitempty" type:"Repeated"`
	MinDuration      *int64                                       `json:"MinDuration,omitempty" xml:"MinDuration,omitempty"`
	OperationName    *string                                      `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	PageNumber       *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pid              *string                                      `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId         *string                                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Reverse          *bool                                        `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	ServiceIp        *string                                      `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName      *string                                      `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	StartTime        *int64                                       `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tags             []*SearchTracesByPageRequestTags             `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s SearchTracesByPageRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageRequest) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageRequest) SetEndTime(v int64) *SearchTracesByPageRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTracesByPageRequest) SetExclusionFilters(v []*SearchTracesByPageRequestExclusionFilters) *SearchTracesByPageRequest {
	s.ExclusionFilters = v
	return s
}

func (s *SearchTracesByPageRequest) SetMinDuration(v int64) *SearchTracesByPageRequest {
	s.MinDuration = &v
	return s
}

func (s *SearchTracesByPageRequest) SetOperationName(v string) *SearchTracesByPageRequest {
	s.OperationName = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPageNumber(v int32) *SearchTracesByPageRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPageSize(v int32) *SearchTracesByPageRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTracesByPageRequest) SetPid(v string) *SearchTracesByPageRequest {
	s.Pid = &v
	return s
}

func (s *SearchTracesByPageRequest) SetRegionId(v string) *SearchTracesByPageRequest {
	s.RegionId = &v
	return s
}

func (s *SearchTracesByPageRequest) SetReverse(v bool) *SearchTracesByPageRequest {
	s.Reverse = &v
	return s
}

func (s *SearchTracesByPageRequest) SetServiceIp(v string) *SearchTracesByPageRequest {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesByPageRequest) SetServiceName(v string) *SearchTracesByPageRequest {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesByPageRequest) SetStartTime(v int64) *SearchTracesByPageRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTracesByPageRequest) SetTags(v []*SearchTracesByPageRequestTags) *SearchTracesByPageRequest {
	s.Tags = v
	return s
}

type SearchTracesByPageRequestExclusionFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesByPageRequestExclusionFilters) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageRequestExclusionFilters) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageRequestExclusionFilters) SetKey(v string) *SearchTracesByPageRequestExclusionFilters {
	s.Key = &v
	return s
}

func (s *SearchTracesByPageRequestExclusionFilters) SetValue(v string) *SearchTracesByPageRequestExclusionFilters {
	s.Value = &v
	return s
}

type SearchTracesByPageRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchTracesByPageRequestTags) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageRequestTags) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageRequestTags) SetKey(v string) *SearchTracesByPageRequestTags {
	s.Key = &v
	return s
}

func (s *SearchTracesByPageRequestTags) SetValue(v string) *SearchTracesByPageRequestTags {
	s.Value = &v
	return s
}

type SearchTracesByPageResponseBody struct {
	PageBean  *SearchTracesByPageResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTracesByPageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBody) SetPageBean(v *SearchTracesByPageResponseBodyPageBean) *SearchTracesByPageResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchTracesByPageResponseBody) SetRequestId(v string) *SearchTracesByPageResponseBody {
	s.RequestId = &v
	return s
}

type SearchTracesByPageResponseBodyPageBean struct {
	PageNumber *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32                                              `json:"Total,omitempty" xml:"Total,omitempty"`
	TraceInfos []*SearchTracesByPageResponseBodyPageBeanTraceInfos `json:"TraceInfos,omitempty" xml:"TraceInfos,omitempty" type:"Repeated"`
}

func (s SearchTracesByPageResponseBodyPageBean) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBodyPageBean) SetPageNumber(v int32) *SearchTracesByPageResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) SetPageSize(v int32) *SearchTracesByPageResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) SetTotal(v int32) *SearchTracesByPageResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) SetTraceInfos(v []*SearchTracesByPageResponseBodyPageBeanTraceInfos) *SearchTracesByPageResponseBodyPageBean {
	s.TraceInfos = v
	return s
}

type SearchTracesByPageResponseBodyPageBeanTraceInfos struct {
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ServiceIp     *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Timestamp     *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TraceID       *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s SearchTracesByPageResponseBodyPageBeanTraceInfos) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageResponseBodyPageBeanTraceInfos) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetDuration(v int64) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.Duration = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetOperationName(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.OperationName = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetServiceIp(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetServiceName(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetTimestamp(v int64) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.Timestamp = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetTraceID(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.TraceID = &v
	return s
}

type SearchTracesByPageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTracesByPageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTracesByPageResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTracesByPageResponse) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponse) SetHeaders(v map[string]*string) *SearchTracesByPageResponse {
	s.Headers = v
	return s
}

func (s *SearchTracesByPageResponse) SetStatusCode(v int32) *SearchTracesByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTracesByPageResponse) SetBody(v *SearchTracesByPageResponseBody) *SearchTracesByPageResponse {
	s.Body = v
	return s
}

type SendTTSVerifyLinkRequest struct {
	ContactId *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Phone     *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s SendTTSVerifyLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendTTSVerifyLinkRequest) GoString() string {
	return s.String()
}

func (s *SendTTSVerifyLinkRequest) SetContactId(v int64) *SendTTSVerifyLinkRequest {
	s.ContactId = &v
	return s
}

func (s *SendTTSVerifyLinkRequest) SetPhone(v string) *SendTTSVerifyLinkRequest {
	s.Phone = &v
	return s
}

type SendTTSVerifyLinkResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendTTSVerifyLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendTTSVerifyLinkResponseBody) GoString() string {
	return s.String()
}

func (s *SendTTSVerifyLinkResponseBody) SetIsSuccess(v bool) *SendTTSVerifyLinkResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SendTTSVerifyLinkResponseBody) SetRequestId(v string) *SendTTSVerifyLinkResponseBody {
	s.RequestId = &v
	return s
}

type SendTTSVerifyLinkResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendTTSVerifyLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendTTSVerifyLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s SendTTSVerifyLinkResponse) GoString() string {
	return s.String()
}

func (s *SendTTSVerifyLinkResponse) SetHeaders(v map[string]*string) *SendTTSVerifyLinkResponse {
	s.Headers = v
	return s
}

func (s *SendTTSVerifyLinkResponse) SetStatusCode(v int32) *SendTTSVerifyLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTTSVerifyLinkResponse) SetBody(v *SendTTSVerifyLinkResponseBody) *SendTTSVerifyLinkResponse {
	s.Body = v
	return s
}

type SetRetcodeShareStatusRequest struct {
	Pid    *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Status *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetRetcodeShareStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRetcodeShareStatusRequest) GoString() string {
	return s.String()
}

func (s *SetRetcodeShareStatusRequest) SetPid(v string) *SetRetcodeShareStatusRequest {
	s.Pid = &v
	return s
}

func (s *SetRetcodeShareStatusRequest) SetStatus(v bool) *SetRetcodeShareStatusRequest {
	s.Status = &v
	return s
}

type SetRetcodeShareStatusResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetRetcodeShareStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRetcodeShareStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetRetcodeShareStatusResponseBody) SetIsSuccess(v bool) *SetRetcodeShareStatusResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *SetRetcodeShareStatusResponseBody) SetRequestId(v string) *SetRetcodeShareStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetRetcodeShareStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetRetcodeShareStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetRetcodeShareStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRetcodeShareStatusResponse) GoString() string {
	return s.String()
}

func (s *SetRetcodeShareStatusResponse) SetHeaders(v map[string]*string) *SetRetcodeShareStatusResponse {
	s.Headers = v
	return s
}

func (s *SetRetcodeShareStatusResponse) SetStatusCode(v int32) *SetRetcodeShareStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRetcodeShareStatusResponse) SetBody(v *SetRetcodeShareStatusResponseBody) *SetRetcodeShareStatusResponse {
	s.Body = v
	return s
}

type StartAlertRequest struct {
	AlertId  *string `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s StartAlertRequest) GoString() string {
	return s.String()
}

func (s *StartAlertRequest) SetAlertId(v string) *StartAlertRequest {
	s.AlertId = &v
	return s
}

func (s *StartAlertRequest) SetRegionId(v string) *StartAlertRequest {
	s.RegionId = &v
	return s
}

type StartAlertResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartAlertResponseBody) GoString() string {
	return s.String()
}

func (s *StartAlertResponseBody) SetIsSuccess(v bool) *StartAlertResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *StartAlertResponseBody) SetRequestId(v string) *StartAlertResponseBody {
	s.RequestId = &v
	return s
}

type StartAlertResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s StartAlertResponse) GoString() string {
	return s.String()
}

func (s *StartAlertResponse) SetHeaders(v map[string]*string) *StartAlertResponse {
	s.Headers = v
	return s
}

func (s *StartAlertResponse) SetStatusCode(v int32) *StartAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAlertResponse) SetBody(v *StartAlertResponseBody) *StartAlertResponse {
	s.Body = v
	return s
}

type StopAlertRequest struct {
	AlertId  *string `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAlertRequest) GoString() string {
	return s.String()
}

func (s *StopAlertRequest) SetAlertId(v string) *StopAlertRequest {
	s.AlertId = &v
	return s
}

func (s *StopAlertRequest) SetRegionId(v string) *StopAlertRequest {
	s.RegionId = &v
	return s
}

type StopAlertResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAlertResponseBody) GoString() string {
	return s.String()
}

func (s *StopAlertResponseBody) SetIsSuccess(v bool) *StopAlertResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *StopAlertResponseBody) SetRequestId(v string) *StopAlertResponseBody {
	s.RequestId = &v
	return s
}

type StopAlertResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAlertResponse) GoString() string {
	return s.String()
}

func (s *StopAlertResponse) SetHeaders(v map[string]*string) *StopAlertResponse {
	s.Headers = v
	return s
}

func (s *StopAlertResponse) SetStatusCode(v int32) *StopAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAlertResponse) SetBody(v *StopAlertResponseBody) *StopAlertResponse {
	s.Body = v
	return s
}

type TurnOnSecondSwitchRequest struct {
	Pid              *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	ProxyUserId      *string `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
	ReleaseStartTime *int64  `json:"ReleaseStartTime,omitempty" xml:"ReleaseStartTime,omitempty"`
}

func (s TurnOnSecondSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s TurnOnSecondSwitchRequest) GoString() string {
	return s.String()
}

func (s *TurnOnSecondSwitchRequest) SetPid(v string) *TurnOnSecondSwitchRequest {
	s.Pid = &v
	return s
}

func (s *TurnOnSecondSwitchRequest) SetProxyUserId(v string) *TurnOnSecondSwitchRequest {
	s.ProxyUserId = &v
	return s
}

func (s *TurnOnSecondSwitchRequest) SetReleaseStartTime(v int64) *TurnOnSecondSwitchRequest {
	s.ReleaseStartTime = &v
	return s
}

type TurnOnSecondSwitchResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TurnOnSecondSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TurnOnSecondSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *TurnOnSecondSwitchResponseBody) SetData(v string) *TurnOnSecondSwitchResponseBody {
	s.Data = &v
	return s
}

func (s *TurnOnSecondSwitchResponseBody) SetRequestId(v string) *TurnOnSecondSwitchResponseBody {
	s.RequestId = &v
	return s
}

type TurnOnSecondSwitchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TurnOnSecondSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TurnOnSecondSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s TurnOnSecondSwitchResponse) GoString() string {
	return s.String()
}

func (s *TurnOnSecondSwitchResponse) SetHeaders(v map[string]*string) *TurnOnSecondSwitchResponse {
	s.Headers = v
	return s
}

func (s *TurnOnSecondSwitchResponse) SetStatusCode(v int32) *TurnOnSecondSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *TurnOnSecondSwitchResponse) SetBody(v *TurnOnSecondSwitchResponseBody) *TurnOnSecondSwitchResponse {
	s.Body = v
	return s
}

type UninstallManagedPrometheusRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UninstallManagedPrometheusRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallManagedPrometheusRequest) GoString() string {
	return s.String()
}

func (s *UninstallManagedPrometheusRequest) SetClusterId(v string) *UninstallManagedPrometheusRequest {
	s.ClusterId = &v
	return s
}

func (s *UninstallManagedPrometheusRequest) SetClusterType(v string) *UninstallManagedPrometheusRequest {
	s.ClusterType = &v
	return s
}

func (s *UninstallManagedPrometheusRequest) SetRegionId(v string) *UninstallManagedPrometheusRequest {
	s.RegionId = &v
	return s
}

func (s *UninstallManagedPrometheusRequest) SetVpcId(v string) *UninstallManagedPrometheusRequest {
	s.VpcId = &v
	return s
}

type UninstallManagedPrometheusResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UninstallManagedPrometheusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallManagedPrometheusResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallManagedPrometheusResponseBody) SetCode(v int32) *UninstallManagedPrometheusResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallManagedPrometheusResponseBody) SetData(v string) *UninstallManagedPrometheusResponseBody {
	s.Data = &v
	return s
}

func (s *UninstallManagedPrometheusResponseBody) SetMessage(v string) *UninstallManagedPrometheusResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallManagedPrometheusResponseBody) SetRequestId(v string) *UninstallManagedPrometheusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallManagedPrometheusResponseBody) SetSuccess(v bool) *UninstallManagedPrometheusResponseBody {
	s.Success = &v
	return s
}

type UninstallManagedPrometheusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UninstallManagedPrometheusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallManagedPrometheusResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallManagedPrometheusResponse) GoString() string {
	return s.String()
}

func (s *UninstallManagedPrometheusResponse) SetHeaders(v map[string]*string) *UninstallManagedPrometheusResponse {
	s.Headers = v
	return s
}

func (s *UninstallManagedPrometheusResponse) SetStatusCode(v int32) *UninstallManagedPrometheusResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallManagedPrometheusResponse) SetBody(v *UninstallManagedPrometheusResponseBody) *UninstallManagedPrometheusResponse {
	s.Body = v
	return s
}

type UninstallPromClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UninstallPromClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallPromClusterRequest) GoString() string {
	return s.String()
}

func (s *UninstallPromClusterRequest) SetClusterId(v string) *UninstallPromClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *UninstallPromClusterRequest) SetRegionId(v string) *UninstallPromClusterRequest {
	s.RegionId = &v
	return s
}

type UninstallPromClusterResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UninstallPromClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallPromClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallPromClusterResponseBody) SetData(v bool) *UninstallPromClusterResponseBody {
	s.Data = &v
	return s
}

func (s *UninstallPromClusterResponseBody) SetRequestId(v string) *UninstallPromClusterResponseBody {
	s.RequestId = &v
	return s
}

type UninstallPromClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UninstallPromClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallPromClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallPromClusterResponse) GoString() string {
	return s.String()
}

func (s *UninstallPromClusterResponse) SetHeaders(v map[string]*string) *UninstallPromClusterResponse {
	s.Headers = v
	return s
}

func (s *UninstallPromClusterResponse) SetStatusCode(v int32) *UninstallPromClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallPromClusterResponse) SetBody(v *UninstallPromClusterResponseBody) *UninstallPromClusterResponse {
	s.Body = v
	return s
}

type UpdateAlertContactRequest struct {
	ContactId           *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName         *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	DingRobotWebhookUrl *string `json:"DingRobotWebhookUrl,omitempty" xml:"DingRobotWebhookUrl,omitempty"`
	Email               *string `json:"Email,omitempty" xml:"Email,omitempty"`
	PhoneNum            *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SystemNoc           *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
}

func (s UpdateAlertContactRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactRequest) SetContactId(v int64) *UpdateAlertContactRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateAlertContactRequest) SetContactName(v string) *UpdateAlertContactRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateAlertContactRequest) SetDingRobotWebhookUrl(v string) *UpdateAlertContactRequest {
	s.DingRobotWebhookUrl = &v
	return s
}

func (s *UpdateAlertContactRequest) SetEmail(v string) *UpdateAlertContactRequest {
	s.Email = &v
	return s
}

func (s *UpdateAlertContactRequest) SetPhoneNum(v string) *UpdateAlertContactRequest {
	s.PhoneNum = &v
	return s
}

func (s *UpdateAlertContactRequest) SetRegionId(v string) *UpdateAlertContactRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertContactRequest) SetSystemNoc(v bool) *UpdateAlertContactRequest {
	s.SystemNoc = &v
	return s
}

type UpdateAlertContactResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlertContactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactResponseBody) SetIsSuccess(v bool) *UpdateAlertContactResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateAlertContactResponseBody) SetRequestId(v string) *UpdateAlertContactResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAlertContactResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAlertContactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAlertContactResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactResponse) SetHeaders(v map[string]*string) *UpdateAlertContactResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertContactResponse) SetStatusCode(v int32) *UpdateAlertContactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertContactResponse) SetBody(v *UpdateAlertContactResponseBody) *UpdateAlertContactResponse {
	s.Body = v
	return s
}

type UpdateAlertContactGroupRequest struct {
	ContactGroupId   *int64  `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	ContactIds       *string `json:"ContactIds,omitempty" xml:"ContactIds,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAlertContactGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactGroupRequest) SetContactGroupId(v int64) *UpdateAlertContactGroupRequest {
	s.ContactGroupId = &v
	return s
}

func (s *UpdateAlertContactGroupRequest) SetContactGroupName(v string) *UpdateAlertContactGroupRequest {
	s.ContactGroupName = &v
	return s
}

func (s *UpdateAlertContactGroupRequest) SetContactIds(v string) *UpdateAlertContactGroupRequest {
	s.ContactIds = &v
	return s
}

func (s *UpdateAlertContactGroupRequest) SetRegionId(v string) *UpdateAlertContactGroupRequest {
	s.RegionId = &v
	return s
}

type UpdateAlertContactGroupResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlertContactGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactGroupResponseBody) SetIsSuccess(v bool) *UpdateAlertContactGroupResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateAlertContactGroupResponseBody) SetRequestId(v string) *UpdateAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAlertContactGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactGroupResponse) SetHeaders(v map[string]*string) *UpdateAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertContactGroupResponse) SetStatusCode(v int32) *UpdateAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertContactGroupResponse) SetBody(v *UpdateAlertContactGroupResponseBody) *UpdateAlertContactGroupResponse {
	s.Body = v
	return s
}

type UpdateAlertRuleRequest struct {
	AlertId             *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	ContactGroupIds     *string `json:"ContactGroupIds,omitempty" xml:"ContactGroupIds,omitempty"`
	IsAutoStart         *bool   `json:"IsAutoStart,omitempty" xml:"IsAutoStart,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplageAlertConfig *string `json:"TemplageAlertConfig,omitempty" xml:"TemplageAlertConfig,omitempty"`
}

func (s UpdateAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleRequest) SetAlertId(v int64) *UpdateAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetContactGroupIds(v string) *UpdateAlertRuleRequest {
	s.ContactGroupIds = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetIsAutoStart(v bool) *UpdateAlertRuleRequest {
	s.IsAutoStart = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetRegionId(v string) *UpdateAlertRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAlertRuleRequest) SetTemplageAlertConfig(v string) *UpdateAlertRuleRequest {
	s.TemplageAlertConfig = &v
	return s
}

type UpdateAlertRuleResponseBody struct {
	AlertId   *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlertRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleResponseBody) SetAlertId(v int64) *UpdateAlertRuleResponseBody {
	s.AlertId = &v
	return s
}

func (s *UpdateAlertRuleResponseBody) SetData(v string) *UpdateAlertRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAlertRuleResponseBody) SetRequestId(v string) *UpdateAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAlertRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAlertRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleResponse) SetHeaders(v map[string]*string) *UpdateAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertRuleResponse) SetStatusCode(v int32) *UpdateAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertRuleResponse) SetBody(v *UpdateAlertRuleResponseBody) *UpdateAlertRuleResponse {
	s.Body = v
	return s
}

type UpdateDispatchRuleRequest struct {
	DispatchRule *string `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateDispatchRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDispatchRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateDispatchRuleRequest) SetDispatchRule(v string) *UpdateDispatchRuleRequest {
	s.DispatchRule = &v
	return s
}

func (s *UpdateDispatchRuleRequest) SetRegionId(v string) *UpdateDispatchRuleRequest {
	s.RegionId = &v
	return s
}

type UpdateDispatchRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDispatchRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDispatchRuleResponseBody) SetRequestId(v string) *UpdateDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDispatchRuleResponseBody) SetSuccess(v bool) *UpdateDispatchRuleResponseBody {
	s.Success = &v
	return s
}

type UpdateDispatchRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDispatchRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateDispatchRuleResponse) SetHeaders(v map[string]*string) *UpdateDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateDispatchRuleResponse) SetStatusCode(v int32) *UpdateDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDispatchRuleResponse) SetBody(v *UpdateDispatchRuleResponseBody) *UpdateDispatchRuleResponse {
	s.Body = v
	return s
}

type UpdateIntegrationRequest struct {
	ApiEndpoint                *string `json:"ApiEndpoint,omitempty" xml:"ApiEndpoint,omitempty"`
	AutoRecover                *bool   `json:"AutoRecover,omitempty" xml:"AutoRecover,omitempty"`
	Description                *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DuplicateKey               *string `json:"DuplicateKey,omitempty" xml:"DuplicateKey,omitempty"`
	ExtendedFieldRedefineRules *string `json:"ExtendedFieldRedefineRules,omitempty" xml:"ExtendedFieldRedefineRules,omitempty"`
	FieldRedefineRules         *string `json:"FieldRedefineRules,omitempty" xml:"FieldRedefineRules,omitempty"`
	IntegrationId              *int64  `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
	IntegrationName            *string `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	IntegrationProductType     *string `json:"IntegrationProductType,omitempty" xml:"IntegrationProductType,omitempty"`
	Liveness                   *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	RecoverTime                *int64  `json:"RecoverTime,omitempty" xml:"RecoverTime,omitempty"`
	ShortToken                 *string `json:"ShortToken,omitempty" xml:"ShortToken,omitempty"`
	Stat                       *string `json:"Stat,omitempty" xml:"Stat,omitempty"`
	State                      *bool   `json:"State,omitempty" xml:"State,omitempty"`
}

func (s UpdateIntegrationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntegrationRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationRequest) SetApiEndpoint(v string) *UpdateIntegrationRequest {
	s.ApiEndpoint = &v
	return s
}

func (s *UpdateIntegrationRequest) SetAutoRecover(v bool) *UpdateIntegrationRequest {
	s.AutoRecover = &v
	return s
}

func (s *UpdateIntegrationRequest) SetDescription(v string) *UpdateIntegrationRequest {
	s.Description = &v
	return s
}

func (s *UpdateIntegrationRequest) SetDuplicateKey(v string) *UpdateIntegrationRequest {
	s.DuplicateKey = &v
	return s
}

func (s *UpdateIntegrationRequest) SetExtendedFieldRedefineRules(v string) *UpdateIntegrationRequest {
	s.ExtendedFieldRedefineRules = &v
	return s
}

func (s *UpdateIntegrationRequest) SetFieldRedefineRules(v string) *UpdateIntegrationRequest {
	s.FieldRedefineRules = &v
	return s
}

func (s *UpdateIntegrationRequest) SetIntegrationId(v int64) *UpdateIntegrationRequest {
	s.IntegrationId = &v
	return s
}

func (s *UpdateIntegrationRequest) SetIntegrationName(v string) *UpdateIntegrationRequest {
	s.IntegrationName = &v
	return s
}

func (s *UpdateIntegrationRequest) SetIntegrationProductType(v string) *UpdateIntegrationRequest {
	s.IntegrationProductType = &v
	return s
}

func (s *UpdateIntegrationRequest) SetLiveness(v string) *UpdateIntegrationRequest {
	s.Liveness = &v
	return s
}

func (s *UpdateIntegrationRequest) SetRecoverTime(v int64) *UpdateIntegrationRequest {
	s.RecoverTime = &v
	return s
}

func (s *UpdateIntegrationRequest) SetShortToken(v string) *UpdateIntegrationRequest {
	s.ShortToken = &v
	return s
}

func (s *UpdateIntegrationRequest) SetStat(v string) *UpdateIntegrationRequest {
	s.Stat = &v
	return s
}

func (s *UpdateIntegrationRequest) SetState(v bool) *UpdateIntegrationRequest {
	s.State = &v
	return s
}

type UpdateIntegrationResponseBody struct {
	Integration *UpdateIntegrationResponseBodyIntegration `json:"Integration,omitempty" xml:"Integration,omitempty" type:"Struct"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIntegrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationResponseBody) SetIntegration(v *UpdateIntegrationResponseBodyIntegration) *UpdateIntegrationResponseBody {
	s.Integration = v
	return s
}

func (s *UpdateIntegrationResponseBody) SetRequestId(v string) *UpdateIntegrationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIntegrationResponseBodyIntegration struct {
	ApiEndpoint                *string                  `json:"ApiEndpoint,omitempty" xml:"ApiEndpoint,omitempty"`
	AutoRecover                *bool                    `json:"AutoRecover,omitempty" xml:"AutoRecover,omitempty"`
	Description                *string                  `json:"Description,omitempty" xml:"Description,omitempty"`
	DuplicateKey               *string                  `json:"DuplicateKey,omitempty" xml:"DuplicateKey,omitempty"`
	ExtendedFieldRedefineRules []map[string]interface{} `json:"ExtendedFieldRedefineRules,omitempty" xml:"ExtendedFieldRedefineRules,omitempty" type:"Repeated"`
	FieldRedefineRules         []map[string]interface{} `json:"FieldRedefineRules,omitempty" xml:"FieldRedefineRules,omitempty" type:"Repeated"`
	IntegrationId              *int64                   `json:"IntegrationId,omitempty" xml:"IntegrationId,omitempty"`
	IntegrationName            *string                  `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	IntegrationProductType     *string                  `json:"IntegrationProductType,omitempty" xml:"IntegrationProductType,omitempty"`
	Liveness                   *string                  `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	RecoverTime                *int64                   `json:"RecoverTime,omitempty" xml:"RecoverTime,omitempty"`
	ShortToken                 *string                  `json:"ShortToken,omitempty" xml:"ShortToken,omitempty"`
	Stat                       []*int64                 `json:"Stat,omitempty" xml:"Stat,omitempty" type:"Repeated"`
	State                      *bool                    `json:"State,omitempty" xml:"State,omitempty"`
}

func (s UpdateIntegrationResponseBodyIntegration) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntegrationResponseBodyIntegration) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationResponseBodyIntegration) SetApiEndpoint(v string) *UpdateIntegrationResponseBodyIntegration {
	s.ApiEndpoint = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetAutoRecover(v bool) *UpdateIntegrationResponseBodyIntegration {
	s.AutoRecover = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetDescription(v string) *UpdateIntegrationResponseBodyIntegration {
	s.Description = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetDuplicateKey(v string) *UpdateIntegrationResponseBodyIntegration {
	s.DuplicateKey = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetExtendedFieldRedefineRules(v []map[string]interface{}) *UpdateIntegrationResponseBodyIntegration {
	s.ExtendedFieldRedefineRules = v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetFieldRedefineRules(v []map[string]interface{}) *UpdateIntegrationResponseBodyIntegration {
	s.FieldRedefineRules = v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetIntegrationId(v int64) *UpdateIntegrationResponseBodyIntegration {
	s.IntegrationId = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetIntegrationName(v string) *UpdateIntegrationResponseBodyIntegration {
	s.IntegrationName = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetIntegrationProductType(v string) *UpdateIntegrationResponseBodyIntegration {
	s.IntegrationProductType = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetLiveness(v string) *UpdateIntegrationResponseBodyIntegration {
	s.Liveness = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetRecoverTime(v int64) *UpdateIntegrationResponseBodyIntegration {
	s.RecoverTime = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetShortToken(v string) *UpdateIntegrationResponseBodyIntegration {
	s.ShortToken = &v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetStat(v []*int64) *UpdateIntegrationResponseBodyIntegration {
	s.Stat = v
	return s
}

func (s *UpdateIntegrationResponseBodyIntegration) SetState(v bool) *UpdateIntegrationResponseBodyIntegration {
	s.State = &v
	return s
}

type UpdateIntegrationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateIntegrationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIntegrationResponse) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationResponse) SetHeaders(v map[string]*string) *UpdateIntegrationResponse {
	s.Headers = v
	return s
}

func (s *UpdateIntegrationResponse) SetStatusCode(v int32) *UpdateIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIntegrationResponse) SetBody(v *UpdateIntegrationResponseBody) *UpdateIntegrationResponse {
	s.Body = v
	return s
}

type UpdatePrometheusAlertRuleRequest struct {
	AlertId        *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64  `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdatePrometheusAlertRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrometheusAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleRequest) SetAlertId(v int64) *UpdatePrometheusAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetAlertName(v string) *UpdatePrometheusAlertRuleRequest {
	s.AlertName = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetAnnotations(v string) *UpdatePrometheusAlertRuleRequest {
	s.Annotations = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetClusterId(v string) *UpdatePrometheusAlertRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetDispatchRuleId(v int64) *UpdatePrometheusAlertRuleRequest {
	s.DispatchRuleId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetDuration(v string) *UpdatePrometheusAlertRuleRequest {
	s.Duration = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetExpression(v string) *UpdatePrometheusAlertRuleRequest {
	s.Expression = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetLabels(v string) *UpdatePrometheusAlertRuleRequest {
	s.Labels = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetMessage(v string) *UpdatePrometheusAlertRuleRequest {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetNotifyType(v string) *UpdatePrometheusAlertRuleRequest {
	s.NotifyType = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetRegionId(v string) *UpdatePrometheusAlertRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleRequest) SetType(v string) *UpdatePrometheusAlertRuleRequest {
	s.Type = &v
	return s
}

type UpdatePrometheusAlertRuleResponseBody struct {
	PrometheusAlertRule *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule `json:"PrometheusAlertRule,omitempty" xml:"PrometheusAlertRule,omitempty" type:"Struct"`
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBody) SetPrometheusAlertRule(v *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) *UpdatePrometheusAlertRuleResponseBody {
	s.PrometheusAlertRule = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBody) SetRequestId(v string) *UpdatePrometheusAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule struct {
	AlertId        *int64                                                                 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertName      *string                                                                `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	Annotations    []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	ClusterId      *string                                                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DispatchRuleId *int64                                                                 `json:"DispatchRuleId,omitempty" xml:"DispatchRuleId,omitempty"`
	Duration       *string                                                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Expression     *string                                                                `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Labels         []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Message        *string                                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	NotifyType     *string                                                                `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	Status         *int32                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertId(v int64) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAlertName(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.AlertName = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetAnnotations(v []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Annotations = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetClusterId(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDispatchRuleId(v int64) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.DispatchRuleId = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetDuration(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Duration = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetExpression(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Expression = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetLabels(v []*UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Labels = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetMessage(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetNotifyType(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.NotifyType = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetStatus(v int32) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Status = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule) SetType(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRule {
	s.Type = &v
	return s
}

type UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetName(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Name = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations) SetValue(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleAnnotations {
	s.Value = &v
	return s
}

type UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetName(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Name = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels) SetValue(v string) *UpdatePrometheusAlertRuleResponseBodyPrometheusAlertRuleLabels {
	s.Value = &v
	return s
}

type UpdatePrometheusAlertRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePrometheusAlertRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePrometheusAlertRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrometheusAlertRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusAlertRuleResponse) SetHeaders(v map[string]*string) *UpdatePrometheusAlertRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrometheusAlertRuleResponse) SetStatusCode(v int32) *UpdatePrometheusAlertRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrometheusAlertRuleResponse) SetBody(v *UpdatePrometheusAlertRuleResponseBody) *UpdatePrometheusAlertRuleResponse {
	s.Body = v
	return s
}

type UpdateWebhookRequest struct {
	Body        *string `json:"Body,omitempty" xml:"Body,omitempty"`
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	HttpHeaders *string `json:"HttpHeaders,omitempty" xml:"HttpHeaders,omitempty"`
	HttpParams  *string `json:"HttpParams,omitempty" xml:"HttpParams,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
	RecoverBody *string `json:"RecoverBody,omitempty" xml:"RecoverBody,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateWebhookRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebhookRequest) GoString() string {
	return s.String()
}

func (s *UpdateWebhookRequest) SetBody(v string) *UpdateWebhookRequest {
	s.Body = &v
	return s
}

func (s *UpdateWebhookRequest) SetContactId(v int64) *UpdateWebhookRequest {
	s.ContactId = &v
	return s
}

func (s *UpdateWebhookRequest) SetContactName(v string) *UpdateWebhookRequest {
	s.ContactName = &v
	return s
}

func (s *UpdateWebhookRequest) SetHttpHeaders(v string) *UpdateWebhookRequest {
	s.HttpHeaders = &v
	return s
}

func (s *UpdateWebhookRequest) SetHttpParams(v string) *UpdateWebhookRequest {
	s.HttpParams = &v
	return s
}

func (s *UpdateWebhookRequest) SetMethod(v string) *UpdateWebhookRequest {
	s.Method = &v
	return s
}

func (s *UpdateWebhookRequest) SetRecoverBody(v string) *UpdateWebhookRequest {
	s.RecoverBody = &v
	return s
}

func (s *UpdateWebhookRequest) SetRegionId(v string) *UpdateWebhookRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWebhookRequest) SetUrl(v string) *UpdateWebhookRequest {
	s.Url = &v
	return s
}

type UpdateWebhookResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWebhookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWebhookResponseBody) SetIsSuccess(v bool) *UpdateWebhookResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateWebhookResponseBody) SetRequestId(v string) *UpdateWebhookResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWebhookResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateWebhookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWebhookResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWebhookResponse) GoString() string {
	return s.String()
}

func (s *UpdateWebhookResponse) SetHeaders(v map[string]*string) *UpdateWebhookResponse {
	s.Headers = v
	return s
}

func (s *UpdateWebhookResponse) SetStatusCode(v int32) *UpdateWebhookResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWebhookResponse) SetBody(v *UpdateWebhookResponseBody) *UpdateWebhookResponse {
	s.Body = v
	return s
}

type UploadRequest struct {
	Edition  *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	File     *string `json:"File,omitempty" xml:"File,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Pid      *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Version  *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UploadRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadRequest) GoString() string {
	return s.String()
}

func (s *UploadRequest) SetEdition(v string) *UploadRequest {
	s.Edition = &v
	return s
}

func (s *UploadRequest) SetFile(v string) *UploadRequest {
	s.File = &v
	return s
}

func (s *UploadRequest) SetFileName(v string) *UploadRequest {
	s.FileName = &v
	return s
}

func (s *UploadRequest) SetPid(v string) *UploadRequest {
	s.Pid = &v
	return s
}

func (s *UploadRequest) SetRegionId(v string) *UploadRequest {
	s.RegionId = &v
	return s
}

func (s *UploadRequest) SetVersion(v string) *UploadRequest {
	s.Version = &v
	return s
}

type UploadResponseBody struct {
	RequestId    *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadResult *UploadResponseBodyUploadResult `json:"UploadResult,omitempty" xml:"UploadResult,omitempty" type:"Struct"`
}

func (s UploadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadResponseBody) GoString() string {
	return s.String()
}

func (s *UploadResponseBody) SetRequestId(v string) *UploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadResponseBody) SetUploadResult(v *UploadResponseBodyUploadResult) *UploadResponseBody {
	s.UploadResult = v
	return s
}

type UploadResponseBodyUploadResult struct {
	Fid        *string `json:"Fid,omitempty" xml:"Fid,omitempty"`
	FileName   *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s UploadResponseBodyUploadResult) String() string {
	return tea.Prettify(s)
}

func (s UploadResponseBodyUploadResult) GoString() string {
	return s.String()
}

func (s *UploadResponseBodyUploadResult) SetFid(v string) *UploadResponseBodyUploadResult {
	s.Fid = &v
	return s
}

func (s *UploadResponseBodyUploadResult) SetFileName(v string) *UploadResponseBodyUploadResult {
	s.FileName = &v
	return s
}

func (s *UploadResponseBodyUploadResult) SetUploadTime(v string) *UploadResponseBodyUploadResult {
	s.UploadTime = &v
	return s
}

type UploadResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadResponse) GoString() string {
	return s.String()
}

func (s *UploadResponse) SetHeaders(v map[string]*string) *UploadResponse {
	s.Headers = v
	return s
}

func (s *UploadResponse) SetStatusCode(v int32) *UploadResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadResponse) SetBody(v *UploadResponseBody) *UploadResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("arms.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("arms.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("arms.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("arms.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("arms.aliyuncs.com"),
		"cn-edge-1":                   tea.String("arms.aliyuncs.com"),
		"cn-fujian":                   tea.String("arms.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("arms.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("arms.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("arms.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("arms.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("arms.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("arms.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("arms.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("arms.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("arms.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("arms.aliyuncs.com"),
		"cn-wuhan":                    tea.String("arms.aliyuncs.com"),
		"cn-yushanfang":               tea.String("arms.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("arms.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("arms.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("arms.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("arms.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("arms.aliyuncs.com"),
		"me-east-1":                   tea.String("arms.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("arms.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("arms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAliClusterIdsToPrometheusGlobalViewWithOptions(request *AddAliClusterIdsToPrometheusGlobalViewRequest, runtime *util.RuntimeOptions) (_result *AddAliClusterIdsToPrometheusGlobalViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterIds)) {
		query["ClusterIds"] = request.ClusterIds
	}

	if !tea.BoolValue(util.IsUnset(request.GlobalViewClusterId)) {
		query["GlobalViewClusterId"] = request.GlobalViewClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddAliClusterIdsToPrometheusGlobalView"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddAliClusterIdsToPrometheusGlobalViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAliClusterIdsToPrometheusGlobalView(request *AddAliClusterIdsToPrometheusGlobalViewRequest) (_result *AddAliClusterIdsToPrometheusGlobalViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAliClusterIdsToPrometheusGlobalViewResponse{}
	_body, _err := client.AddAliClusterIdsToPrometheusGlobalViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGrafanaWithOptions(request *AddGrafanaRequest, runtime *util.RuntimeOptions) (_result *AddGrafanaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Integration)) {
		query["Integration"] = request.Integration
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddGrafana"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddGrafanaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddGrafana(request *AddGrafanaRequest) (_result *AddGrafanaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddGrafanaResponse{}
	_body, _err := client.AddGrafanaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddIntegrationWithOptions(request *AddIntegrationRequest, runtime *util.RuntimeOptions) (_result *AddIntegrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Integration)) {
		query["Integration"] = request.Integration
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddIntegration"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddIntegrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddIntegration(request *AddIntegrationRequest) (_result *AddIntegrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddIntegrationResponse{}
	_body, _err := client.AddIntegrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddPrometheusGlobalViewWithOptions(request *AddPrometheusGlobalViewRequest, runtime *util.RuntimeOptions) (_result *AddPrometheusGlobalViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Clusters)) {
		query["Clusters"] = request.Clusters
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddPrometheusGlobalView"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddPrometheusGlobalViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddPrometheusGlobalView(request *AddPrometheusGlobalViewRequest) (_result *AddPrometheusGlobalViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddPrometheusGlobalViewResponse{}
	_body, _err := client.AddPrometheusGlobalViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddPrometheusGlobalViewByAliClusterIdsWithOptions(request *AddPrometheusGlobalViewByAliClusterIdsRequest, runtime *util.RuntimeOptions) (_result *AddPrometheusGlobalViewByAliClusterIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterIds)) {
		query["ClusterIds"] = request.ClusterIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddPrometheusGlobalViewByAliClusterIds"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddPrometheusGlobalViewByAliClusterIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddPrometheusGlobalViewByAliClusterIds(request *AddPrometheusGlobalViewByAliClusterIdsRequest) (_result *AddPrometheusGlobalViewByAliClusterIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddPrometheusGlobalViewByAliClusterIdsResponse{}
	_body, _err := client.AddPrometheusGlobalViewByAliClusterIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddPrometheusInstanceWithOptions(request *AddPrometheusInstanceRequest, runtime *util.RuntimeOptions) (_result *AddPrometheusInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddPrometheusInstance"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddPrometheusInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddPrometheusInstance(request *AddPrometheusInstanceRequest) (_result *AddPrometheusInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddPrometheusInstanceResponse{}
	_body, _err := client.AddPrometheusInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRecordingRuleWithOptions(request *AddRecordingRuleRequest, runtime *util.RuntimeOptions) (_result *AddRecordingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleYaml)) {
		query["RuleYaml"] = request.RuleYaml
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRecordingRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRecordingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRecordingRule(request *AddRecordingRuleRequest) (_result *AddRecordingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRecordingRuleResponse{}
	_body, _err := client.AddRecordingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppendInstancesToPrometheusGlobalViewWithOptions(request *AppendInstancesToPrometheusGlobalViewRequest, runtime *util.RuntimeOptions) (_result *AppendInstancesToPrometheusGlobalViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Clusters)) {
		query["Clusters"] = request.Clusters
	}

	if !tea.BoolValue(util.IsUnset(request.GlobalViewClusterId)) {
		query["GlobalViewClusterId"] = request.GlobalViewClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AppendInstancesToPrometheusGlobalView"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppendInstancesToPrometheusGlobalViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppendInstancesToPrometheusGlobalView(request *AppendInstancesToPrometheusGlobalViewRequest) (_result *AppendInstancesToPrometheusGlobalViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppendInstancesToPrometheusGlobalViewResponse{}
	_body, _err := client.AppendInstancesToPrometheusGlobalViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyScenarioWithOptions(tmpReq *ApplyScenarioRequest, runtime *util.RuntimeOptions) (_result *ApplyScenarioResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ApplyScenarioShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Config)) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, tea.String("Config"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigShrink)) {
		query["Config"] = request.ConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		query["Scenario"] = request.Scenario
	}

	if !tea.BoolValue(util.IsUnset(request.Sign)) {
		query["Sign"] = request.Sign
	}

	if !tea.BoolValue(util.IsUnset(request.SnDump)) {
		query["SnDump"] = request.SnDump
	}

	if !tea.BoolValue(util.IsUnset(request.SnForce)) {
		query["SnForce"] = request.SnForce
	}

	if !tea.BoolValue(util.IsUnset(request.SnStat)) {
		query["SnStat"] = request.SnStat
	}

	if !tea.BoolValue(util.IsUnset(request.SnTransfer)) {
		query["SnTransfer"] = request.SnTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateOption)) {
		query["UpdateOption"] = request.UpdateOption
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyScenario"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyScenarioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyScenario(request *ApplyScenarioRequest) (_result *ApplyScenarioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyScenarioResponse{}
	_body, _err := client.ApplyScenarioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckServiceStatusWithOptions(request *CheckServiceStatusRequest, runtime *util.RuntimeOptions) (_result *CheckServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SvcCode)) {
		query["SvcCode"] = request.SvcCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckServiceStatus"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckServiceStatus(request *CheckServiceStatusRequest) (_result *CheckServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckServiceStatusResponse{}
	_body, _err := client.CheckServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigAppWithOptions(request *ConfigAppRequest, runtime *util.RuntimeOptions) (_result *ConfigAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		query["AppIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigApp"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigApp(request *ConfigAppRequest) (_result *ConfigAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigAppResponse{}
	_body, _err := client.ConfigAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlertContactWithOptions(request *CreateAlertContactRequest, runtime *util.RuntimeOptions) (_result *CreateAlertContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.DingRobotWebhookUrl)) {
		query["DingRobotWebhookUrl"] = request.DingRobotWebhookUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNum)) {
		query["PhoneNum"] = request.PhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemNoc)) {
		query["SystemNoc"] = request.SystemNoc
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlertContact"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAlertContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlertContact(request *CreateAlertContactRequest) (_result *CreateAlertContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlertContactResponse{}
	_body, _err := client.CreateAlertContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlertContactGroupWithOptions(request *CreateAlertContactGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAlertContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupName)) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactIds)) {
		query["ContactIds"] = request.ContactIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAlertContactGroup"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAlertContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlertContactGroup(request *CreateAlertContactGroupRequest) (_result *CreateAlertContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlertContactGroupResponse{}
	_body, _err := client.CreateAlertContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDispatchRuleWithOptions(request *CreateDispatchRuleRequest, runtime *util.RuntimeOptions) (_result *CreateDispatchRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DispatchRule)) {
		query["DispatchRule"] = request.DispatchRule
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDispatchRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDispatchRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDispatchRule(request *CreateDispatchRuleRequest) (_result *CreateDispatchRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDispatchRuleResponse{}
	_body, _err := client.CreateDispatchRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateIntegrationWithOptions(request *CreateIntegrationRequest, runtime *util.RuntimeOptions) (_result *CreateIntegrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRecover)) {
		body["AutoRecover"] = request.AutoRecover
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationName)) {
		body["IntegrationName"] = request.IntegrationName
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationProductType)) {
		body["IntegrationProductType"] = request.IntegrationProductType
	}

	if !tea.BoolValue(util.IsUnset(request.RecoverTime)) {
		body["RecoverTime"] = request.RecoverTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIntegration"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIntegrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIntegration(request *CreateIntegrationRequest) (_result *CreateIntegrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIntegrationResponse{}
	_body, _err := client.CreateIntegrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrUpdateAlertRuleWithOptions(request *CreateOrUpdateAlertRuleRequest, runtime *util.RuntimeOptions) (_result *CreateOrUpdateAlertRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertCheckType)) {
		body["AlertCheckType"] = request.AlertCheckType
	}

	if !tea.BoolValue(util.IsUnset(request.AlertGroup)) {
		body["AlertGroup"] = request.AlertGroup
	}

	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		body["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.AlertName)) {
		body["AlertName"] = request.AlertName
	}

	if !tea.BoolValue(util.IsUnset(request.AlertRuleContent)) {
		body["AlertRuleContent"] = request.AlertRuleContent
	}

	if !tea.BoolValue(util.IsUnset(request.AlertStatus)) {
		body["AlertStatus"] = request.AlertStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		body["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.Annotations)) {
		body["Annotations"] = request.Annotations
	}

	if !tea.BoolValue(util.IsUnset(request.AutoAddNewApplication)) {
		body["AutoAddNewApplication"] = request.AutoAddNewApplication
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		body["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		body["Level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		body["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.MetricsKey)) {
		body["MetricsKey"] = request.MetricsKey
	}

	if !tea.BoolValue(util.IsUnset(request.MetricsType)) {
		body["MetricsType"] = request.MetricsType
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyStrategy)) {
		body["NotifyStrategy"] = request.NotifyStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Pids)) {
		body["Pids"] = request.Pids
	}

	if !tea.BoolValue(util.IsUnset(request.PromQL)) {
		body["PromQL"] = request.PromQL
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrUpdateAlertRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrUpdateAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrUpdateAlertRule(request *CreateOrUpdateAlertRuleRequest) (_result *CreateOrUpdateAlertRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrUpdateAlertRuleResponse{}
	_body, _err := client.CreateOrUpdateAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrUpdateContactWithOptions(request *CreateOrUpdateContactRequest, runtime *util.RuntimeOptions) (_result *CreateOrUpdateContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		body["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		body["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.ReissueSendNotice)) {
		body["ReissueSendNotice"] = request.ReissueSendNotice
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrUpdateContact"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrUpdateContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrUpdateContact(request *CreateOrUpdateContactRequest) (_result *CreateOrUpdateContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrUpdateContactResponse{}
	_body, _err := client.CreateOrUpdateContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrUpdateContactGroupWithOptions(request *CreateOrUpdateContactGroupRequest, runtime *util.RuntimeOptions) (_result *CreateOrUpdateContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupId)) {
		body["ContactGroupId"] = request.ContactGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroupName)) {
		body["ContactGroupName"] = request.ContactGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactIds)) {
		body["ContactIds"] = request.ContactIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrUpdateContactGroup"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrUpdateContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrUpdateContactGroup(request *CreateOrUpdateContactGroupRequest) (_result *CreateOrUpdateContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrUpdateContactGroupResponse{}
	_body, _err := client.CreateOrUpdateContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrUpdateEventBridgeIntegrationWithOptions(request *CreateOrUpdateEventBridgeIntegrationRequest, runtime *util.RuntimeOptions) (_result *CreateOrUpdateEventBridgeIntegrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKey)) {
		body["AccessKey"] = request.AccessKey
	}

	if !tea.BoolValue(util.IsUnset(request.AccessSecret)) {
		body["AccessSecret"] = request.AccessSecret
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		body["Endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.EventBusName)) {
		body["EventBusName"] = request.EventBusName
	}

	if !tea.BoolValue(util.IsUnset(request.EventBusRegionId)) {
		body["EventBusRegionId"] = request.EventBusRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrUpdateEventBridgeIntegration"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrUpdateEventBridgeIntegrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrUpdateEventBridgeIntegration(request *CreateOrUpdateEventBridgeIntegrationRequest) (_result *CreateOrUpdateEventBridgeIntegrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrUpdateEventBridgeIntegrationResponse{}
	_body, _err := client.CreateOrUpdateEventBridgeIntegrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrUpdateIMRobotWithOptions(request *CreateOrUpdateIMRobotRequest, runtime *util.RuntimeOptions) (_result *CreateOrUpdateIMRobotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CardTemplate)) {
		body["CardTemplate"] = request.CardTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.DailyNoc)) {
		body["DailyNoc"] = request.DailyNoc
	}

	if !tea.BoolValue(util.IsUnset(request.DailyNocTime)) {
		body["DailyNocTime"] = request.DailyNocTime
	}

	if !tea.BoolValue(util.IsUnset(request.EnableOutgoing)) {
		body["EnableOutgoing"] = request.EnableOutgoing
	}

	if !tea.BoolValue(util.IsUnset(request.RobotAddress)) {
		body["RobotAddress"] = request.RobotAddress
	}

	if !tea.BoolValue(util.IsUnset(request.RobotId)) {
		body["RobotId"] = request.RobotId
	}

	if !tea.BoolValue(util.IsUnset(request.RobotName)) {
		body["RobotName"] = request.RobotName
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrUpdateIMRobot"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrUpdateIMRobotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrUpdateIMRobot(request *CreateOrUpdateIMRobotRequest) (_result *CreateOrUpdateIMRobotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrUpdateIMRobotResponse{}
	_body, _err := client.CreateOrUpdateIMRobotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrUpdateNotificationPolicyWithOptions(request *CreateOrUpdateNotificationPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateOrUpdateNotificationPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EscalationPolicyId)) {
		body["EscalationPolicyId"] = request.EscalationPolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupRule)) {
		body["GroupRule"] = request.GroupRule
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationId)) {
		body["IntegrationId"] = request.IntegrationId
	}

	if !tea.BoolValue(util.IsUnset(request.MatchingRules)) {
		body["MatchingRules"] = request.MatchingRules
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyRule)) {
		body["NotifyRule"] = request.NotifyRule
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyTemplate)) {
		body["NotifyTemplate"] = request.NotifyTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Repeat)) {
		body["Repeat"] = request.Repeat
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatInterval)) {
		body["RepeatInterval"] = request.RepeatInterval
	}

	if !tea.BoolValue(util.IsUnset(request.SendRecoverMessage)) {
		body["SendRecoverMessage"] = request.SendRecoverMessage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrUpdateNotificationPolicy"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrUpdateNotificationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrUpdateNotificationPolicy(request *CreateOrUpdateNotificationPolicyRequest) (_result *CreateOrUpdateNotificationPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrUpdateNotificationPolicyResponse{}
	_body, _err := client.CreateOrUpdateNotificationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrUpdateSilencePolicyWithOptions(request *CreateOrUpdateSilencePolicyRequest, runtime *util.RuntimeOptions) (_result *CreateOrUpdateSilencePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.MatchingRules)) {
		body["MatchingRules"] = request.MatchingRules
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrUpdateSilencePolicy"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrUpdateSilencePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrUpdateSilencePolicy(request *CreateOrUpdateSilencePolicyRequest) (_result *CreateOrUpdateSilencePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrUpdateSilencePolicyResponse{}
	_body, _err := client.CreateOrUpdateSilencePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOrUpdateWebhookContactWithOptions(request *CreateOrUpdateWebhookContactRequest, runtime *util.RuntimeOptions) (_result *CreateOrUpdateWebhookContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizHeaders)) {
		body["BizHeaders"] = request.BizHeaders
	}

	if !tea.BoolValue(util.IsUnset(request.BizParams)) {
		body["BizParams"] = request.BizParams
	}

	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.Method)) {
		body["Method"] = request.Method
	}

	if !tea.BoolValue(util.IsUnset(request.RecoverBody)) {
		body["RecoverBody"] = request.RecoverBody
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookId)) {
		body["WebhookId"] = request.WebhookId
	}

	if !tea.BoolValue(util.IsUnset(request.WebhookName)) {
		body["WebhookName"] = request.WebhookName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrUpdateWebhookContact"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrUpdateWebhookContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOrUpdateWebhookContact(request *CreateOrUpdateWebhookContactRequest) (_result *CreateOrUpdateWebhookContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrUpdateWebhookContactResponse{}
	_body, _err := client.CreateOrUpdateWebhookContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePrometheusAlertRuleWithOptions(request *CreatePrometheusAlertRuleRequest, runtime *util.RuntimeOptions) (_result *CreatePrometheusAlertRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertName)) {
		query["AlertName"] = request.AlertName
	}

	if !tea.BoolValue(util.IsUnset(request.Annotations)) {
		query["Annotations"] = request.Annotations
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DispatchRuleId)) {
		query["DispatchRuleId"] = request.DispatchRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyType)) {
		query["NotifyType"] = request.NotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePrometheusAlertRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePrometheusAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePrometheusAlertRule(request *CreatePrometheusAlertRuleRequest) (_result *CreatePrometheusAlertRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePrometheusAlertRuleResponse{}
	_body, _err := client.CreatePrometheusAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRetcodeAppWithOptions(request *CreateRetcodeAppRequest, runtime *util.RuntimeOptions) (_result *CreateRetcodeAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RetcodeAppName)) {
		query["RetcodeAppName"] = request.RetcodeAppName
	}

	if !tea.BoolValue(util.IsUnset(request.RetcodeAppType)) {
		query["RetcodeAppType"] = request.RetcodeAppType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRetcodeApp"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRetcodeAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRetcodeApp(request *CreateRetcodeAppRequest) (_result *CreateRetcodeAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRetcodeAppResponse{}
	_body, _err := client.CreateRetcodeAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSyntheticTaskWithOptions(tmpReq *CreateSyntheticTaskRequest, runtime *util.RuntimeOptions) (_result *CreateSyntheticTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSyntheticTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.CommonParam))) {
		request.CommonParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.CommonParam), tea.String("CommonParam"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ExtendInterval))) {
		request.ExtendIntervalShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ExtendInterval), tea.String("ExtendInterval"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.MonitorList)) {
		request.MonitorListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MonitorList, tea.String("MonitorList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Net))) {
		request.NetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Net), tea.String("Net"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommonParamShrink)) {
		query["CommonParam"] = request.CommonParamShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendIntervalShrink)) {
		query["ExtendInterval"] = request.ExtendIntervalShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalTime)) {
		query["IntervalTime"] = request.IntervalTime
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalType)) {
		query["IntervalType"] = request.IntervalType
	}

	if !tea.BoolValue(util.IsUnset(request.IpType)) {
		query["IpType"] = request.IpType
	}

	if !tea.BoolValue(util.IsUnset(request.MonitorListShrink)) {
		query["MonitorList"] = request.MonitorListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NetShrink)) {
		query["Net"] = request.NetShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSyntheticTask"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSyntheticTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSyntheticTask(request *CreateSyntheticTaskRequest) (_result *CreateSyntheticTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSyntheticTaskResponse{}
	_body, _err := client.CreateSyntheticTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWebhookWithOptions(request *CreateWebhookRequest, runtime *util.RuntimeOptions) (_result *CreateWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		query["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.HttpHeaders)) {
		query["HttpHeaders"] = request.HttpHeaders
	}

	if !tea.BoolValue(util.IsUnset(request.HttpParams)) {
		query["HttpParams"] = request.HttpParams
	}

	if !tea.BoolValue(util.IsUnset(request.Method)) {
		query["Method"] = request.Method
	}

	if !tea.BoolValue(util.IsUnset(request.RecoverBody)) {
		query["RecoverBody"] = request.RecoverBody
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWebhook"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWebhook(request *CreateWebhookRequest) (_result *CreateWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWebhookResponse{}
	_body, _err := client.CreateWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DelAuthTokenWithOptions(request *DelAuthTokenRequest, runtime *util.RuntimeOptions) (_result *DelAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DelAuthToken"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DelAuthTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DelAuthToken(request *DelAuthTokenRequest) (_result *DelAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DelAuthTokenResponse{}
	_body, _err := client.DelAuthTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlertContactWithOptions(request *DeleteAlertContactRequest, runtime *util.RuntimeOptions) (_result *DeleteAlertContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlertContact"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAlertContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlertContact(request *DeleteAlertContactRequest) (_result *DeleteAlertContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlertContactResponse{}
	_body, _err := client.DeleteAlertContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlertContactGroupWithOptions(request *DeleteAlertContactGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteAlertContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupId)) {
		query["ContactGroupId"] = request.ContactGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlertContactGroup"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAlertContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlertContactGroup(request *DeleteAlertContactGroupRequest) (_result *DeleteAlertContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlertContactGroupResponse{}
	_body, _err := client.DeleteAlertContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlertRuleWithOptions(request *DeleteAlertRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteAlertRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlertRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlertRule(request *DeleteAlertRuleRequest) (_result *DeleteAlertRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlertRuleResponse{}
	_body, _err := client.DeleteAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlertRulesWithOptions(request *DeleteAlertRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteAlertRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertIds)) {
		query["AlertIds"] = request.AlertIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAlertRules"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlertRules(request *DeleteAlertRulesRequest) (_result *DeleteAlertRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlertRulesResponse{}
	_body, _err := client.DeleteAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCmsExporterWithOptions(request *DeleteCmsExporterRequest, runtime *util.RuntimeOptions) (_result *DeleteCmsExporterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCmsExporter"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCmsExporterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCmsExporter(request *DeleteCmsExporterRequest) (_result *DeleteCmsExporterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCmsExporterResponse{}
	_body, _err := client.DeleteCmsExporterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContactWithOptions(request *DeleteContactRequest, runtime *util.RuntimeOptions) (_result *DeleteContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteContact"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContact(request *DeleteContactRequest) (_result *DeleteContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContactResponse{}
	_body, _err := client.DeleteContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContactGroupWithOptions(request *DeleteContactGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupId)) {
		query["ContactGroupId"] = request.ContactGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteContactGroup"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContactGroup(request *DeleteContactGroupRequest) (_result *DeleteContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContactGroupResponse{}
	_body, _err := client.DeleteContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDispatchRuleWithOptions(request *DeleteDispatchRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteDispatchRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDispatchRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDispatchRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDispatchRule(request *DeleteDispatchRuleRequest) (_result *DeleteDispatchRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDispatchRuleResponse{}
	_body, _err := client.DeleteDispatchRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEventBridgeIntegrationWithOptions(request *DeleteEventBridgeIntegrationRequest, runtime *util.RuntimeOptions) (_result *DeleteEventBridgeIntegrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEventBridgeIntegration"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEventBridgeIntegrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEventBridgeIntegration(request *DeleteEventBridgeIntegrationRequest) (_result *DeleteEventBridgeIntegrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEventBridgeIntegrationResponse{}
	_body, _err := client.DeleteEventBridgeIntegrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGrafanaResourceWithOptions(request *DeleteGrafanaResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteGrafanaResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		body["ClusterName"] = request.ClusterName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGrafanaResource"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGrafanaResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGrafanaResource(request *DeleteGrafanaResourceRequest) (_result *DeleteGrafanaResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGrafanaResourceResponse{}
	_body, _err := client.DeleteGrafanaResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIMRobotWithOptions(request *DeleteIMRobotRequest, runtime *util.RuntimeOptions) (_result *DeleteIMRobotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RobotId)) {
		query["RobotId"] = request.RobotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIMRobot"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIMRobotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIMRobot(request *DeleteIMRobotRequest) (_result *DeleteIMRobotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIMRobotResponse{}
	_body, _err := client.DeleteIMRobotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIntegrationWithOptions(request *DeleteIntegrationRequest, runtime *util.RuntimeOptions) (_result *DeleteIntegrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Integration)) {
		query["Integration"] = request.Integration
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIntegration"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIntegrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIntegration(request *DeleteIntegrationRequest) (_result *DeleteIntegrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIntegrationResponse{}
	_body, _err := client.DeleteIntegrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteIntegrationsWithOptions(request *DeleteIntegrationsRequest, runtime *util.RuntimeOptions) (_result *DeleteIntegrationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIntegrations"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIntegrationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIntegrations(request *DeleteIntegrationsRequest) (_result *DeleteIntegrationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIntegrationsResponse{}
	_body, _err := client.DeleteIntegrationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNotificationPolicyWithOptions(request *DeleteNotificationPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteNotificationPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNotificationPolicy"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNotificationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNotificationPolicy(request *DeleteNotificationPolicyRequest) (_result *DeleteNotificationPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNotificationPolicyResponse{}
	_body, _err := client.DeleteNotificationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePrometheusAlertRuleWithOptions(request *DeletePrometheusAlertRuleRequest, runtime *util.RuntimeOptions) (_result *DeletePrometheusAlertRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePrometheusAlertRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePrometheusAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePrometheusAlertRule(request *DeletePrometheusAlertRuleRequest) (_result *DeletePrometheusAlertRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePrometheusAlertRuleResponse{}
	_body, _err := client.DeletePrometheusAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePrometheusGlobalViewWithOptions(request *DeletePrometheusGlobalViewRequest, runtime *util.RuntimeOptions) (_result *DeletePrometheusGlobalViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GlobalViewClusterId)) {
		query["GlobalViewClusterId"] = request.GlobalViewClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePrometheusGlobalView"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePrometheusGlobalViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePrometheusGlobalView(request *DeletePrometheusGlobalViewRequest) (_result *DeletePrometheusGlobalViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePrometheusGlobalViewResponse{}
	_body, _err := client.DeletePrometheusGlobalViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRetcodeAppWithOptions(request *DeleteRetcodeAppRequest, runtime *util.RuntimeOptions) (_result *DeleteRetcodeAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRetcodeApp"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRetcodeAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRetcodeApp(request *DeleteRetcodeAppRequest) (_result *DeleteRetcodeAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRetcodeAppResponse{}
	_body, _err := client.DeleteRetcodeAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteScenarioWithOptions(request *DeleteScenarioRequest, runtime *util.RuntimeOptions) (_result *DeleteScenarioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ScenarioId)) {
		query["ScenarioId"] = request.ScenarioId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScenario"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScenarioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteScenario(request *DeleteScenarioRequest) (_result *DeleteScenarioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScenarioResponse{}
	_body, _err := client.DeleteScenarioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSilencePolicyWithOptions(request *DeleteSilencePolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteSilencePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSilencePolicy"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSilencePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSilencePolicy(request *DeleteSilencePolicyRequest) (_result *DeleteSilencePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSilencePolicyResponse{}
	_body, _err := client.DeleteSilencePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSourceMapWithOptions(tmpReq *DeleteSourceMapRequest, runtime *util.RuntimeOptions) (_result *DeleteSourceMapResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteSourceMapShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FidList)) {
		request.FidListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FidList, tea.String("FidList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FidListShrink)) {
		query["FidList"] = request.FidListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSourceMap"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSourceMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSourceMap(request *DeleteSourceMapRequest) (_result *DeleteSourceMapResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSourceMapResponse{}
	_body, _err := client.DeleteSourceMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTraceAppWithOptions(request *DeleteTraceAppRequest, runtime *util.RuntimeOptions) (_result *DeleteTraceAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTraceApp"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTraceAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTraceApp(request *DeleteTraceAppRequest) (_result *DeleteTraceAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTraceAppResponse{}
	_body, _err := client.DeleteTraceAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWebhookContactWithOptions(request *DeleteWebhookContactRequest, runtime *util.RuntimeOptions) (_result *DeleteWebhookContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WebhookId)) {
		query["WebhookId"] = request.WebhookId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWebhookContact"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWebhookContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWebhookContact(request *DeleteWebhookContactRequest) (_result *DeleteWebhookContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWebhookContactResponse{}
	_body, _err := client.DeleteWebhookContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContactGroupsWithOptions(request *DescribeContactGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeContactGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupName)) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.IsDetail)) {
		query["IsDetail"] = request.IsDetail
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContactGroups"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContactGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContactGroups(request *DescribeContactGroupsRequest) (_result *DescribeContactGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContactGroupsResponse{}
	_body, _err := client.DescribeContactGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContactsWithOptions(request *DescribeContactsRequest, runtime *util.RuntimeOptions) (_result *DescribeContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeContacts"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContacts(request *DescribeContactsRequest) (_result *DescribeContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContactsResponse{}
	_body, _err := client.DescribeContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDispatchRuleWithOptions(request *DescribeDispatchRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeDispatchRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDispatchRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDispatchRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDispatchRule(request *DescribeDispatchRuleRequest) (_result *DescribeDispatchRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDispatchRuleResponse{}
	_body, _err := client.DescribeDispatchRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIMRobotsWithOptions(request *DescribeIMRobotsRequest, runtime *util.RuntimeOptions) (_result *DescribeIMRobotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.RobotName)) {
		query["RobotName"] = request.RobotName
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIMRobots"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIMRobotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIMRobots(request *DescribeIMRobotsRequest) (_result *DescribeIMRobotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIMRobotsResponse{}
	_body, _err := client.DescribeIMRobotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePrometheusAlertRuleWithOptions(request *DescribePrometheusAlertRuleRequest, runtime *util.RuntimeOptions) (_result *DescribePrometheusAlertRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePrometheusAlertRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePrometheusAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrometheusAlertRule(request *DescribePrometheusAlertRuleRequest) (_result *DescribePrometheusAlertRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePrometheusAlertRuleResponse{}
	_body, _err := client.DescribePrometheusAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTraceLicenseKeyWithOptions(request *DescribeTraceLicenseKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeTraceLicenseKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTraceLicenseKey"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTraceLicenseKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTraceLicenseKey(request *DescribeTraceLicenseKeyRequest) (_result *DescribeTraceLicenseKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTraceLicenseKeyResponse{}
	_body, _err := client.DescribeTraceLicenseKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeWebhookContactsWithOptions(request *DescribeWebhookContactsRequest, runtime *util.RuntimeOptions) (_result *DescribeWebhookContactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeWebhookContacts"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeWebhookContactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeWebhookContacts(request *DescribeWebhookContactsRequest) (_result *DescribeWebhookContactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeWebhookContactsResponse{}
	_body, _err := client.DescribeWebhookContactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentDownloadUrlWithOptions(request *GetAgentDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *GetAgentDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgentDownloadUrl"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAgentDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentDownloadUrl(request *GetAgentDownloadUrlRequest) (_result *GetAgentDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentDownloadUrlResponse{}
	_body, _err := client.GetAgentDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlertRulesWithOptions(request *GetAlertRulesRequest, runtime *util.RuntimeOptions) (_result *GetAlertRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertIds)) {
		query["AlertIds"] = request.AlertIds
	}

	if !tea.BoolValue(util.IsUnset(request.AlertNames)) {
		query["AlertNames"] = request.AlertNames
	}

	if !tea.BoolValue(util.IsUnset(request.AlertStatus)) {
		query["AlertStatus"] = request.AlertStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		query["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAlertRules"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlertRules(request *GetAlertRulesRequest) (_result *GetAlertRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAlertRulesResponse{}
	_body, _err := client.GetAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppApiByPageWithOptions(request *GetAppApiByPageRequest, runtime *util.RuntimeOptions) (_result *GetAppApiByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalMills)) {
		query["IntervalMills"] = request.IntervalMills
	}

	if !tea.BoolValue(util.IsUnset(request.PId)) {
		query["PId"] = request.PId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppApiByPage"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppApiByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppApiByPage(request *GetAppApiByPageRequest) (_result *GetAppApiByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppApiByPageResponse{}
	_body, _err := client.GetAppApiByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuthTokenWithOptions(request *GetAuthTokenRequest, runtime *util.RuntimeOptions) (_result *GetAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuthToken"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuthToken(request *GetAuthTokenRequest) (_result *GetAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.GetAuthTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterAllUrlWithOptions(request *GetClusterAllUrlRequest, runtime *util.RuntimeOptions) (_result *GetClusterAllUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClusterAllUrl"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterAllUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClusterAllUrl(request *GetClusterAllUrlRequest) (_result *GetClusterAllUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterAllUrlResponse{}
	_body, _err := client.GetClusterAllUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExploreUrlWithOptions(request *GetExploreUrlRequest, runtime *util.RuntimeOptions) (_result *GetExploreUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExploreUrl"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExploreUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExploreUrl(request *GetExploreUrlRequest) (_result *GetExploreUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExploreUrlResponse{}
	_body, _err := client.GetExploreUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIntegrationStateWithOptions(request *GetIntegrationStateRequest, runtime *util.RuntimeOptions) (_result *GetIntegrationStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Integration)) {
		query["Integration"] = request.Integration
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIntegrationState"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIntegrationStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIntegrationState(request *GetIntegrationStateRequest) (_result *GetIntegrationStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIntegrationStateResponse{}
	_body, _err := client.GetIntegrationStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMultipleTraceWithOptions(request *GetMultipleTraceRequest, runtime *util.RuntimeOptions) (_result *GetMultipleTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TraceIDs)) {
		query["TraceIDs"] = request.TraceIDs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMultipleTrace"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMultipleTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMultipleTrace(request *GetMultipleTraceRequest) (_result *GetMultipleTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMultipleTraceResponse{}
	_body, _err := client.GetMultipleTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOnCallSchedulesDetailWithOptions(request *GetOnCallSchedulesDetailRequest, runtime *util.RuntimeOptions) (_result *GetOnCallSchedulesDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOnCallSchedulesDetail"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOnCallSchedulesDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOnCallSchedulesDetail(request *GetOnCallSchedulesDetailRequest) (_result *GetOnCallSchedulesDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOnCallSchedulesDetailResponse{}
	_body, _err := client.GetOnCallSchedulesDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPrometheusApiTokenWithOptions(request *GetPrometheusApiTokenRequest, runtime *util.RuntimeOptions) (_result *GetPrometheusApiTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPrometheusApiToken"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPrometheusApiTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPrometheusApiToken(request *GetPrometheusApiTokenRequest) (_result *GetPrometheusApiTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPrometheusApiTokenResponse{}
	_body, _err := client.GetPrometheusApiTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPrometheusGlobalViewWithOptions(request *GetPrometheusGlobalViewRequest, runtime *util.RuntimeOptions) (_result *GetPrometheusGlobalViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GlobalViewClusterId)) {
		query["GlobalViewClusterId"] = request.GlobalViewClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPrometheusGlobalView"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPrometheusGlobalViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPrometheusGlobalView(request *GetPrometheusGlobalViewRequest) (_result *GetPrometheusGlobalViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPrometheusGlobalViewResponse{}
	_body, _err := client.GetPrometheusGlobalViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRecordingRuleWithOptions(request *GetRecordingRuleRequest, runtime *util.RuntimeOptions) (_result *GetRecordingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRecordingRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRecordingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRecordingRule(request *GetRecordingRuleRequest) (_result *GetRecordingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRecordingRuleResponse{}
	_body, _err := client.GetRecordingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRetcodeShareUrlWithOptions(request *GetRetcodeShareUrlRequest, runtime *util.RuntimeOptions) (_result *GetRetcodeShareUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRetcodeShareUrl"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRetcodeShareUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRetcodeShareUrl(request *GetRetcodeShareUrlRequest) (_result *GetRetcodeShareUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRetcodeShareUrlResponse{}
	_body, _err := client.GetRetcodeShareUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSourceMapInfoWithOptions(request *GetSourceMapInfoRequest, runtime *util.RuntimeOptions) (_result *GetSourceMapInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AscendingSequence)) {
		query["AscendingSequence"] = request.AscendingSequence
	}

	if !tea.BoolValue(util.IsUnset(request.Edition)) {
		query["Edition"] = request.Edition
	}

	if !tea.BoolValue(util.IsUnset(request.ID)) {
		query["ID"] = request.ID
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.OrderField)) {
		query["OrderField"] = request.OrderField
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSourceMapInfo"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSourceMapInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSourceMapInfo(request *GetSourceMapInfoRequest) (_result *GetSourceMapInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSourceMapInfoResponse{}
	_body, _err := client.GetSourceMapInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStackWithOptions(request *GetStackRequest, runtime *util.RuntimeOptions) (_result *GetStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RpcID)) {
		query["RpcID"] = request.RpcID
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TraceID)) {
		query["TraceID"] = request.TraceID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStack"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStack(request *GetStackRequest) (_result *GetStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackResponse{}
	_body, _err := client.GetStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSyntheticTaskMonitorsWithOptions(request *GetSyntheticTaskMonitorsRequest, runtime *util.RuntimeOptions) (_result *GetSyntheticTaskMonitorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSyntheticTaskMonitors"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSyntheticTaskMonitorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSyntheticTaskMonitors(request *GetSyntheticTaskMonitorsRequest) (_result *GetSyntheticTaskMonitorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSyntheticTaskMonitorsResponse{}
	_body, _err := client.GetSyntheticTaskMonitorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTraceWithOptions(request *GetTraceRequest, runtime *util.RuntimeOptions) (_result *GetTraceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TraceID)) {
		query["TraceID"] = request.TraceID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrace"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrace(request *GetTraceRequest) (_result *GetTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTraceResponse{}
	_body, _err := client.GetTraceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTraceAppWithOptions(request *GetTraceAppRequest, runtime *util.RuntimeOptions) (_result *GetTraceAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTraceApp"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTraceAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTraceApp(request *GetTraceAppRequest) (_result *GetTraceAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTraceAppResponse{}
	_body, _err := client.GetTraceAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportAppAlertRulesWithOptions(request *ImportAppAlertRulesRequest, runtime *util.RuntimeOptions) (_result *ImportAppAlertRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupIds)) {
		query["ContactGroupIds"] = request.ContactGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoStart)) {
		query["IsAutoStart"] = request.IsAutoStart
	}

	if !tea.BoolValue(util.IsUnset(request.Pids)) {
		query["Pids"] = request.Pids
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplageAlertConfig)) {
		query["TemplageAlertConfig"] = request.TemplageAlertConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateAlertId)) {
		query["TemplateAlertId"] = request.TemplateAlertId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportAppAlertRules"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportAppAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportAppAlertRules(request *ImportAppAlertRulesRequest) (_result *ImportAppAlertRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportAppAlertRulesResponse{}
	_body, _err := client.ImportAppAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallCmsExporterWithOptions(request *InstallCmsExporterRequest, runtime *util.RuntimeOptions) (_result *InstallCmsExporterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.CmsArgs)) {
		query["CmsArgs"] = request.CmsArgs
	}

	if !tea.BoolValue(util.IsUnset(request.DirectArgs)) {
		query["DirectArgs"] = request.DirectArgs
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTag)) {
		query["EnableTag"] = request.EnableTag
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallCmsExporter"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallCmsExporterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallCmsExporter(request *InstallCmsExporterRequest) (_result *InstallCmsExporterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallCmsExporterResponse{}
	_body, _err := client.InstallCmsExporterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallManagedPrometheusWithOptions(request *InstallManagedPrometheusRequest, runtime *util.RuntimeOptions) (_result *InstallManagedPrometheusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.KubeConfig)) {
		query["KubeConfig"] = request.KubeConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallManagedPrometheus"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallManagedPrometheusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallManagedPrometheus(request *InstallManagedPrometheusRequest) (_result *InstallManagedPrometheusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallManagedPrometheusResponse{}
	_body, _err := client.InstallManagedPrometheusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListActivatedAlertsWithOptions(request *ListActivatedAlertsRequest, runtime *util.RuntimeOptions) (_result *ListActivatedAlertsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListActivatedAlerts"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListActivatedAlertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListActivatedAlerts(request *ListActivatedAlertsRequest) (_result *ListActivatedAlertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListActivatedAlertsResponse{}
	_body, _err := client.ListActivatedAlertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlertsWithOptions(request *ListAlertsRequest, runtime *util.RuntimeOptions) (_result *ListAlertsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertName)) {
		query["AlertName"] = request.AlertName
	}

	if !tea.BoolValue(util.IsUnset(request.DispatchRuleId)) {
		query["DispatchRuleId"] = request.DispatchRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationType)) {
		query["IntegrationType"] = request.IntegrationType
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Severity)) {
		query["Severity"] = request.Severity
	}

	if !tea.BoolValue(util.IsUnset(request.ShowActivities)) {
		query["ShowActivities"] = request.ShowActivities
	}

	if !tea.BoolValue(util.IsUnset(request.ShowEvents)) {
		query["ShowEvents"] = request.ShowEvents
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAlerts"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAlertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlerts(request *ListAlertsRequest) (_result *ListAlertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlertsResponse{}
	_body, _err := client.ListAlertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterFromGrafanaWithOptions(request *ListClusterFromGrafanaRequest, runtime *util.RuntimeOptions) (_result *ListClusterFromGrafanaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterFromGrafana"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterFromGrafanaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterFromGrafana(request *ListClusterFromGrafanaRequest) (_result *ListClusterFromGrafanaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterFromGrafanaResponse{}
	_body, _err := client.ListClusterFromGrafanaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCmsInstancesWithOptions(request *ListCmsInstancesRequest, runtime *util.RuntimeOptions) (_result *ListCmsInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCmsInstances"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCmsInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCmsInstances(request *ListCmsInstancesRequest) (_result *ListCmsInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCmsInstancesResponse{}
	_body, _err := client.ListCmsInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDashboardsWithOptions(request *ListDashboardsRequest, runtime *util.RuntimeOptions) (_result *ListDashboardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.DashboardName)) {
		query["DashboardName"] = request.DashboardName
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.RecreateSwitch)) {
		query["RecreateSwitch"] = request.RecreateSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDashboards"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDashboardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDashboards(request *ListDashboardsRequest) (_result *ListDashboardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDashboardsResponse{}
	_body, _err := client.ListDashboardsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDashboardsByNameWithOptions(request *ListDashboardsByNameRequest, runtime *util.RuntimeOptions) (_result *ListDashboardsByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.DashBoardName)) {
		query["DashBoardName"] = request.DashBoardName
	}

	if !tea.BoolValue(util.IsUnset(request.DashBoardVersion)) {
		query["DashBoardVersion"] = request.DashBoardVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.OnlyQuery)) {
		query["OnlyQuery"] = request.OnlyQuery
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDashboardsByName"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDashboardsByNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDashboardsByName(request *ListDashboardsByNameRequest) (_result *ListDashboardsByNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDashboardsByNameResponse{}
	_body, _err := client.ListDashboardsByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDispatchRuleWithOptions(request *ListDispatchRuleRequest, runtime *util.RuntimeOptions) (_result *ListDispatchRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.System)) {
		query["System"] = request.System
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDispatchRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDispatchRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDispatchRule(request *ListDispatchRuleRequest) (_result *ListDispatchRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDispatchRuleResponse{}
	_body, _err := client.ListDispatchRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEscalationPoliciesWithOptions(request *ListEscalationPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListEscalationPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEscalationPolicies"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEscalationPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEscalationPolicies(request *ListEscalationPoliciesRequest) (_result *ListEscalationPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEscalationPoliciesResponse{}
	_body, _err := client.ListEscalationPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEventBridgeIntegrationsWithOptions(request *ListEventBridgeIntegrationsRequest, runtime *util.RuntimeOptions) (_result *ListEventBridgeIntegrationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEventBridgeIntegrations"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEventBridgeIntegrationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEventBridgeIntegrations(request *ListEventBridgeIntegrationsRequest) (_result *ListEventBridgeIntegrationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEventBridgeIntegrationsResponse{}
	_body, _err := client.ListEventBridgeIntegrationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInsightsEventsWithOptions(request *ListInsightsEventsRequest, runtime *util.RuntimeOptions) (_result *ListInsightsEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InsightsTypes)) {
		query["InsightsTypes"] = request.InsightsTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInsightsEvents"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInsightsEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInsightsEvents(request *ListInsightsEventsRequest) (_result *ListInsightsEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInsightsEventsResponse{}
	_body, _err := client.ListInsightsEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIntegrationWithOptions(request *ListIntegrationRequest, runtime *util.RuntimeOptions) (_result *ListIntegrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIntegration"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIntegrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIntegration(request *ListIntegrationRequest) (_result *ListIntegrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIntegrationResponse{}
	_body, _err := client.ListIntegrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNotificationPoliciesWithOptions(request *ListNotificationPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListNotificationPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsDetail)) {
		query["IsDetail"] = request.IsDetail
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNotificationPolicies"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNotificationPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNotificationPolicies(request *ListNotificationPoliciesRequest) (_result *ListNotificationPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNotificationPoliciesResponse{}
	_body, _err := client.ListNotificationPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOnCallSchedulesWithOptions(request *ListOnCallSchedulesRequest, runtime *util.RuntimeOptions) (_result *ListOnCallSchedulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOnCallSchedules"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOnCallSchedulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOnCallSchedules(request *ListOnCallSchedulesRequest) (_result *ListOnCallSchedulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOnCallSchedulesResponse{}
	_body, _err := client.ListOnCallSchedulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPrometheusAlertRulesWithOptions(request *ListPrometheusAlertRulesRequest, runtime *util.RuntimeOptions) (_result *ListPrometheusAlertRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MatchExpressions)) {
		query["MatchExpressions"] = request.MatchExpressions
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrometheusAlertRules"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrometheusAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPrometheusAlertRules(request *ListPrometheusAlertRulesRequest) (_result *ListPrometheusAlertRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrometheusAlertRulesResponse{}
	_body, _err := client.ListPrometheusAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPrometheusAlertTemplatesWithOptions(request *ListPrometheusAlertTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListPrometheusAlertTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrometheusAlertTemplates"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrometheusAlertTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPrometheusAlertTemplates(request *ListPrometheusAlertTemplatesRequest) (_result *ListPrometheusAlertTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrometheusAlertTemplatesResponse{}
	_body, _err := client.ListPrometheusAlertTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPrometheusGlobalViewWithOptions(request *ListPrometheusGlobalViewRequest, runtime *util.RuntimeOptions) (_result *ListPrometheusGlobalViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrometheusGlobalView"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrometheusGlobalViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPrometheusGlobalView(request *ListPrometheusGlobalViewRequest) (_result *ListPrometheusGlobalViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrometheusGlobalViewResponse{}
	_body, _err := client.ListPrometheusGlobalViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPrometheusInstancesWithOptions(request *ListPrometheusInstancesRequest, runtime *util.RuntimeOptions) (_result *ListPrometheusInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowGlobalView)) {
		query["ShowGlobalView"] = request.ShowGlobalView
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrometheusInstances"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrometheusInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPrometheusInstances(request *ListPrometheusInstancesRequest) (_result *ListPrometheusInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrometheusInstancesResponse{}
	_body, _err := client.ListPrometheusInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRetcodeAppsWithOptions(request *ListRetcodeAppsRequest, runtime *util.RuntimeOptions) (_result *ListRetcodeAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRetcodeApps"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRetcodeAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRetcodeApps(request *ListRetcodeAppsRequest) (_result *ListRetcodeAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRetcodeAppsResponse{}
	_body, _err := client.ListRetcodeAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScenarioWithOptions(request *ListScenarioRequest, runtime *util.RuntimeOptions) (_result *ListScenarioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Scenario)) {
		query["Scenario"] = request.Scenario
	}

	if !tea.BoolValue(util.IsUnset(request.Sign)) {
		query["Sign"] = request.Sign
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListScenario"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListScenarioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScenario(request *ListScenarioRequest) (_result *ListScenarioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScenarioResponse{}
	_body, _err := client.ListScenarioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSilencePoliciesWithOptions(request *ListSilencePoliciesRequest, runtime *util.RuntimeOptions) (_result *ListSilencePoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsDetail)) {
		query["IsDetail"] = request.IsDetail
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSilencePolicies"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSilencePoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSilencePolicies(request *ListSilencePoliciesRequest) (_result *ListSilencePoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSilencePoliciesResponse{}
	_body, _err := client.ListSilencePoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTraceAppsWithOptions(request *ListTraceAppsRequest, runtime *util.RuntimeOptions) (_result *ListTraceAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTraceApps"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTraceAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTraceApps(request *ListTraceAppsRequest) (_result *ListTraceAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTraceAppsResponse{}
	_body, _err := client.ListTraceAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ManageGetRecordingRuleWithOptions(request *ManageGetRecordingRuleRequest, runtime *util.RuntimeOptions) (_result *ManageGetRecordingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryUserId)) {
		query["QueryUserId"] = request.QueryUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ManageGetRecordingRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ManageGetRecordingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ManageGetRecordingRule(request *ManageGetRecordingRuleRequest) (_result *ManageGetRecordingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ManageGetRecordingRuleResponse{}
	_body, _err := client.ManageGetRecordingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ManageRecordingRuleWithOptions(request *ManageRecordingRuleRequest, runtime *util.RuntimeOptions) (_result *ManageRecordingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryUserId)) {
		query["QueryUserId"] = request.QueryUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleYaml)) {
		query["RuleYaml"] = request.RuleYaml
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ManageRecordingRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ManageRecordingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ManageRecordingRule(request *ManageRecordingRuleRequest) (_result *ManageRecordingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ManageRecordingRuleResponse{}
	_body, _err := client.ManageRecordingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenArmsDefaultSLRWithOptions(request *OpenArmsDefaultSLRRequest, runtime *util.RuntimeOptions) (_result *OpenArmsDefaultSLRResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenArmsDefaultSLR"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenArmsDefaultSLRResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenArmsDefaultSLR(request *OpenArmsDefaultSLRRequest) (_result *OpenArmsDefaultSLRResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenArmsDefaultSLRResponse{}
	_body, _err := client.OpenArmsDefaultSLRWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenArmsServiceSecondVersionWithOptions(request *OpenArmsServiceSecondVersionRequest, runtime *util.RuntimeOptions) (_result *OpenArmsServiceSecondVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenArmsServiceSecondVersion"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenArmsServiceSecondVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenArmsServiceSecondVersion(request *OpenArmsServiceSecondVersionRequest) (_result *OpenArmsServiceSecondVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenArmsServiceSecondVersionResponse{}
	_body, _err := client.OpenArmsServiceSecondVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenVClusterWithOptions(request *OpenVClusterRequest, runtime *util.RuntimeOptions) (_result *OpenVClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		query["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	if !tea.BoolValue(util.IsUnset(request.RecreateSwitch)) {
		query["RecreateSwitch"] = request.RecreateSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenVCluster"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenVClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenVCluster(request *OpenVClusterRequest) (_result *OpenVClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenVClusterResponse{}
	_body, _err := client.OpenVClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OpenXtraceDefaultSLRWithOptions(request *OpenXtraceDefaultSLRRequest, runtime *util.RuntimeOptions) (_result *OpenXtraceDefaultSLRResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenXtraceDefaultSLR"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenXtraceDefaultSLRResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenXtraceDefaultSLR(request *OpenXtraceDefaultSLRRequest) (_result *OpenXtraceDefaultSLRResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenXtraceDefaultSLRResponse{}
	_body, _err := client.OpenXtraceDefaultSLRWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMetricByPageWithOptions(request *QueryMetricByPageRequest, runtime *util.RuntimeOptions) (_result *QueryMetricByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.CustomFilters)) {
		query["CustomFilters"] = request.CustomFilters
	}

	if !tea.BoolValue(util.IsUnset(request.Dimensions)) {
		query["Dimensions"] = request.Dimensions
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.IntervalInSec)) {
		query["IntervalInSec"] = request.IntervalInSec
	}

	if !tea.BoolValue(util.IsUnset(request.Measures)) {
		query["Measures"] = request.Measures
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		query["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMetricByPage"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMetricByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMetricByPage(request *QueryMetricByPageRequest) (_result *QueryMetricByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMetricByPageResponse{}
	_body, _err := client.QueryMetricByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryPromInstallStatusWithOptions(request *QueryPromInstallStatusRequest, runtime *util.RuntimeOptions) (_result *QueryPromInstallStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPromInstallStatus"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPromInstallStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryPromInstallStatus(request *QueryPromInstallStatusRequest) (_result *QueryPromInstallStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPromInstallStatusResponse{}
	_body, _err := client.QueryPromInstallStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryReleaseMetricWithOptions(request *QueryReleaseMetricRequest, runtime *util.RuntimeOptions) (_result *QueryReleaseMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeOrderId)) {
		query["ChangeOrderId"] = request.ChangeOrderId
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		query["CreateTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyUserId)) {
		query["ProxyUserId"] = request.ProxyUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseEndTime)) {
		query["ReleaseEndTime"] = request.ReleaseEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseStartTime)) {
		query["ReleaseStartTime"] = request.ReleaseStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Service)) {
		query["Service"] = request.Service
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryReleaseMetric"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryReleaseMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryReleaseMetric(request *QueryReleaseMetricRequest) (_result *QueryReleaseMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryReleaseMetricResponse{}
	_body, _err := client.QueryReleaseMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAliClusterIdsFromPrometheusGlobalViewWithOptions(request *RemoveAliClusterIdsFromPrometheusGlobalViewRequest, runtime *util.RuntimeOptions) (_result *RemoveAliClusterIdsFromPrometheusGlobalViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterIds)) {
		query["ClusterIds"] = request.ClusterIds
	}

	if !tea.BoolValue(util.IsUnset(request.GlobalViewClusterId)) {
		query["GlobalViewClusterId"] = request.GlobalViewClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveAliClusterIdsFromPrometheusGlobalView"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAliClusterIdsFromPrometheusGlobalViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAliClusterIdsFromPrometheusGlobalView(request *RemoveAliClusterIdsFromPrometheusGlobalViewRequest) (_result *RemoveAliClusterIdsFromPrometheusGlobalViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveAliClusterIdsFromPrometheusGlobalViewResponse{}
	_body, _err := client.RemoveAliClusterIdsFromPrometheusGlobalViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSourcesFromPrometheusGlobalViewWithOptions(request *RemoveSourcesFromPrometheusGlobalViewRequest, runtime *util.RuntimeOptions) (_result *RemoveSourcesFromPrometheusGlobalViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GlobalViewClusterId)) {
		query["GlobalViewClusterId"] = request.GlobalViewClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceNames)) {
		query["SourceNames"] = request.SourceNames
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveSourcesFromPrometheusGlobalView"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveSourcesFromPrometheusGlobalViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSourcesFromPrometheusGlobalView(request *RemoveSourcesFromPrometheusGlobalViewRequest) (_result *RemoveSourcesFromPrometheusGlobalViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveSourcesFromPrometheusGlobalViewResponse{}
	_body, _err := client.RemoveSourcesFromPrometheusGlobalViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTraceAppConfigWithOptions(request *SaveTraceAppConfigRequest, runtime *util.RuntimeOptions) (_result *SaveTraceAppConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.Settings)) {
		query["Settings"] = request.Settings
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveTraceAppConfig"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveTraceAppConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTraceAppConfig(request *SaveTraceAppConfigRequest) (_result *SaveTraceAppConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTraceAppConfigResponse{}
	_body, _err := client.SaveTraceAppConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchAlertContactWithOptions(request *SearchAlertContactRequest, runtime *util.RuntimeOptions) (_result *SearchAlertContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactIds)) {
		query["ContactIds"] = request.ContactIds
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchAlertContact"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAlertContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchAlertContact(request *SearchAlertContactRequest) (_result *SearchAlertContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchAlertContactResponse{}
	_body, _err := client.SearchAlertContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchAlertContactGroupWithOptions(request *SearchAlertContactGroupRequest, runtime *util.RuntimeOptions) (_result *SearchAlertContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupIds)) {
		query["ContactGroupIds"] = request.ContactGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroupName)) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.IsDetail)) {
		query["IsDetail"] = request.IsDetail
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchAlertContactGroup"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAlertContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchAlertContactGroup(request *SearchAlertContactGroupRequest) (_result *SearchAlertContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchAlertContactGroupResponse{}
	_body, _err := client.SearchAlertContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchAlertHistoriesWithOptions(request *SearchAlertHistoriesRequest, runtime *util.RuntimeOptions) (_result *SearchAlertHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		query["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchAlertHistories"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAlertHistoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchAlertHistories(request *SearchAlertHistoriesRequest) (_result *SearchAlertHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchAlertHistoriesResponse{}
	_body, _err := client.SearchAlertHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchAlertRulesWithOptions(request *SearchAlertRulesRequest, runtime *util.RuntimeOptions) (_result *SearchAlertRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemRegionId)) {
		query["SystemRegionId"] = request.SystemRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchAlertRules"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAlertRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchAlertRules(request *SearchAlertRulesRequest) (_result *SearchAlertRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchAlertRulesResponse{}
	_body, _err := client.SearchAlertRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchEventsWithOptions(request *SearchEventsRequest, runtime *util.RuntimeOptions) (_result *SearchEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.AlertType)) {
		query["AlertType"] = request.AlertType
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsTrigger)) {
		query["IsTrigger"] = request.IsTrigger
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchEvents"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchEvents(request *SearchEventsRequest) (_result *SearchEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchEventsResponse{}
	_body, _err := client.SearchEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchRetcodeAppByPageWithOptions(request *SearchRetcodeAppByPageRequest, runtime *util.RuntimeOptions) (_result *SearchRetcodeAppByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RetcodeAppName)) {
		query["RetcodeAppName"] = request.RetcodeAppName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchRetcodeAppByPage"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchRetcodeAppByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchRetcodeAppByPage(request *SearchRetcodeAppByPageRequest) (_result *SearchRetcodeAppByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchRetcodeAppByPageResponse{}
	_body, _err := client.SearchRetcodeAppByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTraceAppByNameWithOptions(request *SearchTraceAppByNameRequest, runtime *util.RuntimeOptions) (_result *SearchTraceAppByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TraceAppName)) {
		query["TraceAppName"] = request.TraceAppName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTraceAppByName"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTraceAppByNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTraceAppByName(request *SearchTraceAppByNameRequest) (_result *SearchTraceAppByNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTraceAppByNameResponse{}
	_body, _err := client.SearchTraceAppByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTraceAppByPageWithOptions(request *SearchTraceAppByPageRequest, runtime *util.RuntimeOptions) (_result *SearchTraceAppByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TraceAppName)) {
		query["TraceAppName"] = request.TraceAppName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTraceAppByPage"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTraceAppByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTraceAppByPage(request *SearchTraceAppByPageRequest) (_result *SearchTraceAppByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTraceAppByPageResponse{}
	_body, _err := client.SearchTraceAppByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTracesWithOptions(request *SearchTracesRequest, runtime *util.RuntimeOptions) (_result *SearchTracesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExclusionFilters)) {
		query["ExclusionFilters"] = request.ExclusionFilters
	}

	if !tea.BoolValue(util.IsUnset(request.MinDuration)) {
		query["MinDuration"] = request.MinDuration
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceIp)) {
		query["ServiceIp"] = request.ServiceIp
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTraces"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTracesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTraces(request *SearchTracesRequest) (_result *SearchTracesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTracesResponse{}
	_body, _err := client.SearchTracesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTracesByPageWithOptions(request *SearchTracesByPageRequest, runtime *util.RuntimeOptions) (_result *SearchTracesByPageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExclusionFilters)) {
		query["ExclusionFilters"] = request.ExclusionFilters
	}

	if !tea.BoolValue(util.IsUnset(request.MinDuration)) {
		query["MinDuration"] = request.MinDuration
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Reverse)) {
		query["Reverse"] = request.Reverse
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceIp)) {
		query["ServiceIp"] = request.ServiceIp
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTracesByPage"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTracesByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTracesByPage(request *SearchTracesByPageRequest) (_result *SearchTracesByPageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTracesByPageResponse{}
	_body, _err := client.SearchTracesByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendTTSVerifyLinkWithOptions(request *SendTTSVerifyLinkRequest, runtime *util.RuntimeOptions) (_result *SendTTSVerifyLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		body["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["Phone"] = request.Phone
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendTTSVerifyLink"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendTTSVerifyLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendTTSVerifyLink(request *SendTTSVerifyLinkRequest) (_result *SendTTSVerifyLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendTTSVerifyLinkResponse{}
	_body, _err := client.SendTTSVerifyLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetRetcodeShareStatusWithOptions(request *SetRetcodeShareStatusRequest, runtime *util.RuntimeOptions) (_result *SetRetcodeShareStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetRetcodeShareStatus"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetRetcodeShareStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetRetcodeShareStatus(request *SetRetcodeShareStatusRequest) (_result *SetRetcodeShareStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetRetcodeShareStatusResponse{}
	_body, _err := client.SetRetcodeShareStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartAlertWithOptions(request *StartAlertRequest, runtime *util.RuntimeOptions) (_result *StartAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartAlert"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartAlert(request *StartAlertRequest) (_result *StartAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartAlertResponse{}
	_body, _err := client.StartAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopAlertWithOptions(request *StopAlertRequest, runtime *util.RuntimeOptions) (_result *StopAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopAlert"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopAlert(request *StopAlertRequest) (_result *StopAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAlertResponse{}
	_body, _err := client.StopAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TurnOnSecondSwitchWithOptions(request *TurnOnSecondSwitchRequest, runtime *util.RuntimeOptions) (_result *TurnOnSecondSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TurnOnSecondSwitch"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TurnOnSecondSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TurnOnSecondSwitch(request *TurnOnSecondSwitchRequest) (_result *TurnOnSecondSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TurnOnSecondSwitchResponse{}
	_body, _err := client.TurnOnSecondSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallManagedPrometheusWithOptions(request *UninstallManagedPrometheusRequest, runtime *util.RuntimeOptions) (_result *UninstallManagedPrometheusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallManagedPrometheus"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallManagedPrometheusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallManagedPrometheus(request *UninstallManagedPrometheusRequest) (_result *UninstallManagedPrometheusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UninstallManagedPrometheusResponse{}
	_body, _err := client.UninstallManagedPrometheusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallPromClusterWithOptions(request *UninstallPromClusterRequest, runtime *util.RuntimeOptions) (_result *UninstallPromClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallPromCluster"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallPromClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallPromCluster(request *UninstallPromClusterRequest) (_result *UninstallPromClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UninstallPromClusterResponse{}
	_body, _err := client.UninstallPromClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAlertContactWithOptions(request *UpdateAlertContactRequest, runtime *util.RuntimeOptions) (_result *UpdateAlertContactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.DingRobotWebhookUrl)) {
		query["DingRobotWebhookUrl"] = request.DingRobotWebhookUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNum)) {
		query["PhoneNum"] = request.PhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemNoc)) {
		query["SystemNoc"] = request.SystemNoc
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlertContact"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAlertContactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAlertContact(request *UpdateAlertContactRequest) (_result *UpdateAlertContactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAlertContactResponse{}
	_body, _err := client.UpdateAlertContactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAlertContactGroupWithOptions(request *UpdateAlertContactGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateAlertContactGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactGroupId)) {
		query["ContactGroupId"] = request.ContactGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroupName)) {
		query["ContactGroupName"] = request.ContactGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ContactIds)) {
		query["ContactIds"] = request.ContactIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlertContactGroup"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAlertContactGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAlertContactGroup(request *UpdateAlertContactGroupRequest) (_result *UpdateAlertContactGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAlertContactGroupResponse{}
	_body, _err := client.UpdateAlertContactGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAlertRuleWithOptions(request *UpdateAlertRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateAlertRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroupIds)) {
		query["ContactGroupIds"] = request.ContactGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.IsAutoStart)) {
		query["IsAutoStart"] = request.IsAutoStart
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplageAlertConfig)) {
		query["TemplageAlertConfig"] = request.TemplageAlertConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAlertRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAlertRule(request *UpdateAlertRuleRequest) (_result *UpdateAlertRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAlertRuleResponse{}
	_body, _err := client.UpdateAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDispatchRuleWithOptions(request *UpdateDispatchRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateDispatchRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DispatchRule)) {
		query["DispatchRule"] = request.DispatchRule
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDispatchRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDispatchRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDispatchRule(request *UpdateDispatchRuleRequest) (_result *UpdateDispatchRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDispatchRuleResponse{}
	_body, _err := client.UpdateDispatchRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateIntegrationWithOptions(request *UpdateIntegrationRequest, runtime *util.RuntimeOptions) (_result *UpdateIntegrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiEndpoint)) {
		body["ApiEndpoint"] = request.ApiEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRecover)) {
		body["AutoRecover"] = request.AutoRecover
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DuplicateKey)) {
		body["DuplicateKey"] = request.DuplicateKey
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendedFieldRedefineRules)) {
		body["ExtendedFieldRedefineRules"] = request.ExtendedFieldRedefineRules
	}

	if !tea.BoolValue(util.IsUnset(request.FieldRedefineRules)) {
		body["FieldRedefineRules"] = request.FieldRedefineRules
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationId)) {
		body["IntegrationId"] = request.IntegrationId
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationName)) {
		body["IntegrationName"] = request.IntegrationName
	}

	if !tea.BoolValue(util.IsUnset(request.IntegrationProductType)) {
		body["IntegrationProductType"] = request.IntegrationProductType
	}

	if !tea.BoolValue(util.IsUnset(request.Liveness)) {
		body["Liveness"] = request.Liveness
	}

	if !tea.BoolValue(util.IsUnset(request.RecoverTime)) {
		body["RecoverTime"] = request.RecoverTime
	}

	if !tea.BoolValue(util.IsUnset(request.ShortToken)) {
		body["ShortToken"] = request.ShortToken
	}

	if !tea.BoolValue(util.IsUnset(request.Stat)) {
		body["Stat"] = request.Stat
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		body["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIntegration"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIntegrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateIntegration(request *UpdateIntegrationRequest) (_result *UpdateIntegrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIntegrationResponse{}
	_body, _err := client.UpdateIntegrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePrometheusAlertRuleWithOptions(request *UpdatePrometheusAlertRuleRequest, runtime *util.RuntimeOptions) (_result *UpdatePrometheusAlertRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlertId)) {
		query["AlertId"] = request.AlertId
	}

	if !tea.BoolValue(util.IsUnset(request.AlertName)) {
		query["AlertName"] = request.AlertName
	}

	if !tea.BoolValue(util.IsUnset(request.Annotations)) {
		query["Annotations"] = request.Annotations
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DispatchRuleId)) {
		query["DispatchRuleId"] = request.DispatchRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Expression)) {
		query["Expression"] = request.Expression
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		query["Message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyType)) {
		query["NotifyType"] = request.NotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePrometheusAlertRule"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePrometheusAlertRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePrometheusAlertRule(request *UpdatePrometheusAlertRuleRequest) (_result *UpdatePrometheusAlertRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePrometheusAlertRuleResponse{}
	_body, _err := client.UpdatePrometheusAlertRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWebhookWithOptions(request *UpdateWebhookRequest, runtime *util.RuntimeOptions) (_result *UpdateWebhookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		query["Body"] = request.Body
	}

	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.ContactName)) {
		query["ContactName"] = request.ContactName
	}

	if !tea.BoolValue(util.IsUnset(request.HttpHeaders)) {
		query["HttpHeaders"] = request.HttpHeaders
	}

	if !tea.BoolValue(util.IsUnset(request.HttpParams)) {
		query["HttpParams"] = request.HttpParams
	}

	if !tea.BoolValue(util.IsUnset(request.Method)) {
		query["Method"] = request.Method
	}

	if !tea.BoolValue(util.IsUnset(request.RecoverBody)) {
		query["RecoverBody"] = request.RecoverBody
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWebhook"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWebhookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWebhook(request *UpdateWebhookRequest) (_result *UpdateWebhookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWebhookResponse{}
	_body, _err := client.UpdateWebhookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadWithOptions(request *UploadRequest, runtime *util.RuntimeOptions) (_result *UploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Edition)) {
		query["Edition"] = request.Edition
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		query["Pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.File)) {
		body["File"] = request.File
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Upload"),
		Version:     tea.String("2019-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Upload(request *UploadRequest) (_result *UploadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadResponse{}
	_body, _err := client.UploadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
