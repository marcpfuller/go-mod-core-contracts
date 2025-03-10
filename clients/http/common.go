//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

type commonClient struct {
	baseUrl string
}

// NewCommonClient creates an instance of CommonClient
func NewCommonClient(baseUrl string) interfaces.CommonClient {
	return &commonClient{
		baseUrl: baseUrl,
	}
}

func (cc *commonClient) Configuration(ctx context.Context) (dtoCommon.ConfigResponse, errors.EdgeX) {
	cr := dtoCommon.ConfigResponse{}
	err := utils.GetRequest(ctx, &cr, cc.baseUrl, common.ApiConfigRoute, nil)
	if err != nil {
		return cr, errors.NewCommonEdgeXWrapper(err)
	}
	return cr, nil
}

func (cc *commonClient) Ping(ctx context.Context) (dtoCommon.PingResponse, errors.EdgeX) {
	pr := dtoCommon.PingResponse{}
	err := utils.GetRequest(ctx, &pr, cc.baseUrl, common.ApiPingRoute, nil)
	if err != nil {
		return pr, errors.NewCommonEdgeXWrapper(err)
	}
	return pr, nil
}

func (cc *commonClient) Version(ctx context.Context) (dtoCommon.VersionResponse, errors.EdgeX) {
	vr := dtoCommon.VersionResponse{}
	err := utils.GetRequest(ctx, &vr, cc.baseUrl, common.ApiVersionRoute, nil)
	if err != nil {
		return vr, errors.NewCommonEdgeXWrapper(err)
	}
	return vr, nil
}

func (cc *commonClient) AddSecret(ctx context.Context, request dtoCommon.SecretRequest) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	err = utils.PostRequestWithRawData(ctx, &res, cc.baseUrl, common.ApiSecretRoute, nil, request)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}
