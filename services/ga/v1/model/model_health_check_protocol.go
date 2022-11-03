package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 健康检查的协议。
type HealthCheckProtocol struct {
	value string
}

type HealthCheckProtocolEnum struct {
	TCP HealthCheckProtocol
}

func GetHealthCheckProtocolEnum() HealthCheckProtocolEnum {
	return HealthCheckProtocolEnum{
		TCP: HealthCheckProtocol{
			value: "TCP",
		},
	}
}

func (c HealthCheckProtocol) Value() string {
	return c.value
}

func (c HealthCheckProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HealthCheckProtocol) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
