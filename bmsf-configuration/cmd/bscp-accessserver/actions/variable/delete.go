/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.,
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package template

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/viper"

	pb "bk-bscp/internal/protocol/accessserver"
	pbcommon "bk-bscp/internal/protocol/common"
	pbtemplateserver "bk-bscp/internal/protocol/templateserver"
	"bk-bscp/pkg/common"
	"bk-bscp/pkg/logger"
)

// DeleteAction delete a variable object
type DeleteAction struct {
	viper          *viper.Viper
	templateClient pbtemplateserver.TemplateClient
	req            *pb.DeleteVariableReq
	resp           *pb.DeleteVariableResp
}

// NewDeleteAction creates new DeleteAction
func NewDeleteAction(viper *viper.Viper, templateClient pbtemplateserver.TemplateClient,
	req *pb.DeleteVariableReq, resp *pb.DeleteVariableResp) *DeleteAction {
	action := &DeleteAction{viper: viper, templateClient: templateClient, req: req, resp: resp}

	action.resp.Seq = req.Seq
	action.resp.ErrCode = pbcommon.ErrCode_E_OK
	action.resp.ErrMsg = "OK"

	return action
}

// Err setup error code message in response and return the error.
func (act *DeleteAction) Err(errCode pbcommon.ErrCode, errMsg string) error {
	act.resp.ErrCode = errCode
	act.resp.ErrMsg = errMsg
	return errors.New(errMsg)
}

// Input handles the input messages.
func (act *DeleteAction) Input() error {
	if err := act.verify(); err != nil {
		return act.Err(pbcommon.ErrCode_E_AS_PARAMS_INVALID, err.Error())
	}
	return nil
}

// Output handles the output messages.
func (act *DeleteAction) Output() error {
	// do nothing.
	return nil
}

func (act *DeleteAction) verify() error {
	if err := common.VerifyID(act.req.Bid, "bid"); err != nil {
		return err
	}

	if err := common.VerifyID(act.req.Vid, "vid"); err != nil {
		return err
	}

	if err := common.VerifyVariableType(act.req.Type); err != nil {
		return err
	}

	if err := common.VerifyNormalName(act.req.Operator, "operator"); err != nil {
		return err
	}

	return nil
}

func (act *DeleteAction) deleteVariable() (pbcommon.ErrCode, string) {
	req := &pbtemplateserver.DeleteVariableReq{
		Seq:      act.req.Seq,
		Bid:      act.req.Bid,
		Vid:      act.req.Vid,
		Type:     act.req.Type,
		Operator: act.req.Operator,
	}

	ctx, cancel := context.WithTimeout(context.Background(), act.viper.GetDuration("templateserver.calltimeout"))
	defer cancel()

	logger.V(2).Infof("DeleteVariableReq[%d]| request to templateserver DeleteVariableReq, %+v", req.Seq, req)

	resp, err := act.templateClient.DeleteVariable(ctx, req)
	if err != nil {
		return pbcommon.ErrCode_E_AS_SYSTEM_UNKONW, fmt.Sprintf("request to templateserver DeleteVariable, %+v", err)
	}
	if resp.ErrCode != pbcommon.ErrCode_E_OK {
		return resp.ErrCode, resp.ErrMsg
	}

	return pbcommon.ErrCode_E_OK, "OK"
}

// Do do action
func (act *DeleteAction) Do() error {

	// delete variable
	if errCode, errMsg := act.deleteVariable(); errCode != pbcommon.ErrCode_E_OK {
		return act.Err(errCode, errMsg)
	}
	return nil
}