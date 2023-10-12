package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowPathObjectByIdRequest Request Object
type ShowPathObjectByIdRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType ShowPathObjectByIdRequestDlmType `json:"Dlm-Type"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 目录编号
	CatalogId string `json:"catalog_id"`

	// limit
	Limit *int32 `json:"limit,omitempty"`

	// offset
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowPathObjectByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPathObjectByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowPathObjectByIdRequest", string(data)}, " ")
}

type ShowPathObjectByIdRequestDlmType struct {
	value string
}

type ShowPathObjectByIdRequestDlmTypeEnum struct {
	SHARED    ShowPathObjectByIdRequestDlmType
	EXCLUSIVE ShowPathObjectByIdRequestDlmType
}

func GetShowPathObjectByIdRequestDlmTypeEnum() ShowPathObjectByIdRequestDlmTypeEnum {
	return ShowPathObjectByIdRequestDlmTypeEnum{
		SHARED: ShowPathObjectByIdRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowPathObjectByIdRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowPathObjectByIdRequestDlmType) Value() string {
	return c.value
}

func (c ShowPathObjectByIdRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPathObjectByIdRequestDlmType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
