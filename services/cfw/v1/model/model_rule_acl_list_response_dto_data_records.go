package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// items
type RuleAclListResponseDtoDataRecords struct {

	// 规则id
	RuleId *string `json:"rule_id,omitempty"`

	// 地址类型0 ipv4,1 ipv6
	AddressType *int32 `json:"address_type,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	Sequence *OrderRuleAclDto `json:"sequence,omitempty"`

	// 规则方向0：外到内1：内到外
	Direction *RuleAclListResponseDtoDataRecordsDirection `json:"direction,omitempty"`

	// 动作0：permit,1：deny
	ActionType *int32 `json:"action_type,omitempty"`

	// 规则下发状态 0：禁用,1：启用
	Status *int32 `json:"status,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 长连接时长小时
	LongConnectTimeHour *int64 `json:"long_connect_time_hour,omitempty"`

	// 长连接时长分钟
	LongConnectTimeMinute *int64 `json:"long_connect_time_minute,omitempty"`

	// 长连接时长秒
	LongConnectTimeSecond *int64 `json:"long_connect_time_second,omitempty"`

	// 长连接时长
	LongConnectTime *int64 `json:"long_connect_time,omitempty"`

	// 长连接支持
	LongConnectEnable *int32 `json:"long_connect_enable,omitempty"`

	Source *RuleAddressDto `json:"source,omitempty"`

	Destination *RuleAddressDto `json:"destination,omitempty"`

	Service *RuleServiceDto `json:"service,omitempty"`
}

func (o RuleAclListResponseDtoDataRecords) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleAclListResponseDtoDataRecords struct{}"
	}

	return strings.Join([]string{"RuleAclListResponseDtoDataRecords", string(data)}, " ")
}

type RuleAclListResponseDtoDataRecordsDirection struct {
	value int32
}

type RuleAclListResponseDtoDataRecordsDirectionEnum struct {
	E_0 RuleAclListResponseDtoDataRecordsDirection
	E_1 RuleAclListResponseDtoDataRecordsDirection
}

func GetRuleAclListResponseDtoDataRecordsDirectionEnum() RuleAclListResponseDtoDataRecordsDirectionEnum {
	return RuleAclListResponseDtoDataRecordsDirectionEnum{
		E_0: RuleAclListResponseDtoDataRecordsDirection{
			value: 0,
		}, E_1: RuleAclListResponseDtoDataRecordsDirection{
			value: 1,
		},
	}
}

func (c RuleAclListResponseDtoDataRecordsDirection) Value() int32 {
	return c.value
}

func (c RuleAclListResponseDtoDataRecordsDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuleAclListResponseDtoDataRecordsDirection) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}