package model

import (
	"encoding/json"

	"strings"
)

//
type IdCardRequestBody struct {
	// 与url二选一  图像数据，base64编码，要求base64编码后大小不超过10MB。图片最小边不小于15px，最长边不超过8000px。支持JPEG、JPG、PNG、BMP、TIFF格式。  图片文件Base64编码字符串，点击[这里](https://support.huaweicloud.com/ocr_faq/ocr_01_0032.html)查看详细获取方式。

	Image *string `json:"image,omitempty"`
	// 与image二选一  图片的URL路径，目前支持：  - 公网http/https url  - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。  > 说明：  - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。  - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。

	Url *string `json:"url,omitempty"`
	//  - front：身份证正面。  - back：身份证背面。  > 说明： 如果参数值为空或无该参数，系统自动识别，建议填写，准确率更高。

	Side *string `json:"side,omitempty"`
	// 返回校验身份证号等信息的开关，默认false，可选值如下所示：  - true：返回校验信息  - false：不反回校验信息

	ReturnVerification *bool `json:"return_verification,omitempty"`
}

func (o IdCardRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IdCardRequestBody struct{}"
	}

	return strings.Join([]string{"IdCardRequestBody", string(data)}, " ")
}