//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"
)

// DeviceProfileBasicInfoRequest defines the Request Content for PATCH UpdateDeviceProfileBasicInfo DTO.
// This object and its properties correspond to the DeviceProfileBasicInfoRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.2.0#/DeviceProfileBasicInfoRequest
type DeviceProfileBasicInfoRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	BasicInfo             dtos.UpdateDeviceProfileBasicInfo `json:"basicinfo"`
}

// Validate satisfies the Validator interface
func (d DeviceProfileBasicInfoRequest) Validate() error {
	err := common.Validate(d)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateDeviceRequest type
func (d *DeviceProfileBasicInfoRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		BasicInfo dtos.UpdateDeviceProfileBasicInfo
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*d = DeviceProfileBasicInfoRequest(alias)

	// validate DeviceProfileBasicInfoRequest DTO
	if err := d.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceDeviceProfileModelBasicInfoFieldsWithDTO replace existing DeviceProfile's basic info fields with DTO patch
func ReplaceDeviceProfileModelBasicInfoFieldsWithDTO(dp *models.DeviceProfile, patch dtos.UpdateDeviceProfileBasicInfo) {
	if patch.Description != nil {
		dp.Description = *patch.Description
	}
	if patch.Manufacturer != nil {
		dp.Manufacturer = *patch.Manufacturer
	}
	if patch.Model != nil {
		dp.Model = *patch.Model
	}
	if patch.Labels != nil {
		dp.Labels = patch.Labels
	}
}
