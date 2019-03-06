package cdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateLiveAppSnapshotConfig invokes the cdn.UpdateLiveAppSnapshotConfig API synchronously
// api document: https://help.aliyun.com/api/cdn/updateliveappsnapshotconfig.html
func (client *Client) UpdateLiveAppSnapshotConfig(request *UpdateLiveAppSnapshotConfigRequest) (response *UpdateLiveAppSnapshotConfigResponse, err error) {
	response = CreateUpdateLiveAppSnapshotConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLiveAppSnapshotConfigWithChan invokes the cdn.UpdateLiveAppSnapshotConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/updateliveappsnapshotconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateLiveAppSnapshotConfigWithChan(request *UpdateLiveAppSnapshotConfigRequest) (<-chan *UpdateLiveAppSnapshotConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateLiveAppSnapshotConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLiveAppSnapshotConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// UpdateLiveAppSnapshotConfigWithCallback invokes the cdn.UpdateLiveAppSnapshotConfig API asynchronously
// api document: https://help.aliyun.com/api/cdn/updateliveappsnapshotconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateLiveAppSnapshotConfigWithCallback(request *UpdateLiveAppSnapshotConfigRequest, callback func(response *UpdateLiveAppSnapshotConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLiveAppSnapshotConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateLiveAppSnapshotConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// UpdateLiveAppSnapshotConfigRequest is the request struct for api UpdateLiveAppSnapshotConfig
type UpdateLiveAppSnapshotConfigRequest struct {
	*requests.RpcRequest
	TimeInterval       requests.Integer `position:"Query" name:"TimeInterval"`
	OssBucket          string           `position:"Query" name:"OssBucket"`
	AppName            string           `position:"Query" name:"AppName"`
	SecurityToken      string           `position:"Query" name:"SecurityToken"`
	DomainName         string           `position:"Query" name:"DomainName"`
	OssEndpoint        string           `position:"Query" name:"OssEndpoint"`
	SequenceOssObject  string           `position:"Query" name:"SequenceOssObject"`
	OverwriteOssObject string           `position:"Query" name:"OverwriteOssObject"`
	OwnerId            requests.Integer `position:"Query" name:"OwnerId"`
}

// UpdateLiveAppSnapshotConfigResponse is the response struct for api UpdateLiveAppSnapshotConfig
type UpdateLiveAppSnapshotConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateLiveAppSnapshotConfigRequest creates a request to invoke UpdateLiveAppSnapshotConfig API
func CreateUpdateLiveAppSnapshotConfigRequest() (request *UpdateLiveAppSnapshotConfigRequest) {
	request = &UpdateLiveAppSnapshotConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "UpdateLiveAppSnapshotConfig", "", "")
	return
}

// CreateUpdateLiveAppSnapshotConfigResponse creates a response to parse from UpdateLiveAppSnapshotConfig response
func CreateUpdateLiveAppSnapshotConfigResponse() (response *UpdateLiveAppSnapshotConfigResponse) {
	response = &UpdateLiveAppSnapshotConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}