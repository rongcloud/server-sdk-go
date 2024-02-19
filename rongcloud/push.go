package rongcloud

import (
	"context"
	"encoding/json"
	"fmt"
)

type PushCustomRequest struct {
	// [必传] 目标平台（操作系统），可以为 ios、android 其中一个或全部。全部填写时给给 Android、iOS 两个平台推送消息。
	Platform []string `json:"platform,omitempty"`
	// [必传] 推送条件。支持按用户标签推送（tag 、tag_or）、按应用包名推送（packageName）和按指定平台全部推送（is_to_all）。注意：如果推送条件中 is_to_all 为 true，则忽略其他推送条件。
	Audience *PushCustomAudience `json:"audience,omitempty"`
	// [必传] 按平台（操作系统）指定推送内容。
	Notification *PushCustomNotification `json:"notification,omitempty"`
}

type PushCustomAudience struct {
	// 用户标签数组，标签之间为 AND 关系。数组中最多包含 20 个标签。
	Tag []string `json:"tag,omitempty"`
	// 用户标签数组，标签之间为 OR 关系。数组中最多包含 20 个标签。
	TagOr []string `json:"tag_or,omitempty"`
	// 应用包名。与 tag 或 tag_or 为逻辑与（AND）关系。
	PackageName *string `json:"packageName,omitempty"`
	// [必传] 是否按指定平台（操作系统）全部推送。true 表示全部推送，此时其他推送条件字段均无效。false 表示按其他推送条件进行推送。
	IsToAll *bool `json:"is_to_all,omitempty"`
	// [必传] 在按用户标签推送场景下，可通过 tagItems 实现复杂与或非逻辑。在 tagItems 包含有效内容的情况下，tag、tag_or 字段无效。
	TagItems []*PushCustomAudienceTagItem `json:"tagItems,omitempty"`
}

type PushCustomAudienceTagItem struct {
	// [必传] 用户标签数组。
	Tags []string `json:"tags,omitempty"`
	// [必传] 是否对 tags 数组的运算结果进行非运算。默认为 false。
	IsNot *bool `json:"isNot,omitempty"`
	// [必传] tags 数组内标签之间的运算符。
	TagsOperator *string `json:"tagsOperator,omitempty"`
	// [必传] tagItems 数组内当前 Object 与上一个 Object 之间的运算符。注意：首个 Object 内的 itemsOperator 未被使用，为无效字段。
	ItemsOperator *string `json:"itemsOperator,omitempty"`
}

type PushCustomNotification struct {
	// 通知栏显示标题，最长不超过 50 个字符。
	Title *string `json:"title,omitempty"`
	// 默认推送通知内容。
	Alert *string `json:"alert,omitempty"`
	// 设置 iOS 平台下的推送及附加信息。
	IOS *PushCustomNotificationIOS `json:"ios,omitempty"`
	// 设置 Android 平台下的推送及附加信息。
	Android *PushAndroid `json:"android,omitempty"`
}

type PushCustomNotificationIOS struct {
	// 通知栏显示的推送标题，仅针对 iOS 平台，支持 iOS 8.2 及以上版本，参数在 ios 节点下设置，详细可参考“设置 iOS 推送标题请求示例”，此属性优先级高于 notification 下的 title。
	Title *string `json:"title,omitempty"`
	// 针对 iOS 平台，静默推送是 iOS7 之后推出的一种推送方式。 允许应用在收到通知后在后台运行一段代码，且能够马上执行。详情请查看知识库文档。1 表示为开启，0 表示为关闭，默认为 0
	ContentAvailable *int `json:"contentAvailable,omitempty"`
	// 应用角标，仅针对 iOS 平台；不填时，表示不改变角标数；为 0 或负数时，表示 App 角标上的数字清零；否则传相应数字表示把角标数改为指定的数字，最大不超过 9999，参数在 ios 节点下设置，详细可参考“设置 iOS 角标数 HTTP 请求示例”。
	Badge *int `json:"badge,omitempty"`
	// iOS 平台通知栏分组 ID，相同的 thread-id 推送分一组，单组超过 5 条推送会折叠展示
	ThreadId *string `json:"thread-id,omitempty"`
	// iOS 平台，从 iOS10 开始支持，设置后设备收到有相同 ID 的消息，会合并成一条
	ApnsCollapseId *string `json:"apns-collapse-id,omitempty"`
	// iOS 富文本推送的类型开发者自己定义，自己在 App 端进行解析判断，与 richMediaUri 一起使用，当设置 category 后，推送时默认携带 mutable-content 进行推送，属性值为 1。
	Category *string `json:"category,omitempty"`
	// iOS 富文本推送内容的 URL，与 category 一起使用。
	RichMediaUri *string `json:"richMediaUri,omitempty"`
	// 适用于 iOS 15 及之后的系统。取值为 passive，active（默认），time-sensitive，或 critical，取值说明详见对应的 APNs 的 interruption-level 字段。在 iOS 15 及以上版本中，系统的 “定时推送摘要”、“专注模式” 都可能导致重要的推送通知（例如余额变化）无法及时被用户感知的情况，可考虑设置该字段。
	InterruptionLevel *string `json:"interruption-level,omitempty"`
	// 附加信息，如果开发者自己需要，可以自己在 App 端进行解析。
	Extras interface{} `json:"extras,omitempty"`
}

type PushAndroid struct {
	Honor  *PushAndroidHonor `json:"honor,omitempty"`
	HW     *PushAndroidHW    `json:"hw,omitempty"`
	Oppo   *PushAndroidOppo  `json:"oppo,omitempty"`
	Vivo   *PushAndroidVivo  `json:"vivo,omitempty"`
	Fcm    *PushAndroidFcm   `json:"fcm,omitempty"`
	Extras interface{}       `json:"extras,omitempty"`
}

type PushAndroidHonor struct {
	// 荣耀通知栏消息优先级，取值：
	//  NORMAL（服务与通讯类消息）
	//  LOW（咨询营销类消息）。若资讯营销类消息发送时带图片，图片不会展示。
	Importance *string `json:"importance,omitempty"`
	// 荣耀推送自定义通知栏消息右侧的大图标 URL，若不设置，则不展示通知栏右侧图标。
	//  URL 使用的协议必须是 HTTPS 协议，取值样例：https://example.com/image.png。
	//  图标文件须小于 512KB，图标建议规格大小：40dp x 40dp，弧角大小为 8dp。
	//  超出建议规格大小的图标会存在图片压缩或显示不全的情况。
	Image *string `json:"image,omitempty"`
}

type PushAndroidHW struct {
	// 华为推送通知渠道的 ID。详见自定义通知渠道。
	ChannelId *string `json:"channelId,omitempty"`
	// 华为推送通知栏消息优先级，取值 NORMAL、LOW，默认为 NORMAL 重要消息。
	Importance *string `json:"importance,omitempty"`
	// 华为推送自定义的通知栏消息右侧大图标 URL，如果不设置，则不展示通知栏右侧图标。URL 使用的协议必须是 HTTPS 协议，取值样例：https://example.com/image.png。图标文件须小于 512KB，图标建议规格大小：40dp x 40dp，弧角大小为 8dp，超出建议规格大小的图标会存在图片压缩或显示不全的情况。
	Image *string `json:"image,omitempty"`
	// 华为推送通道的消息自分类标识，category 取值必须为大写字母，例如 IM。App 根据华为要求完成自分类权益申请 或 申请特殊权限 后可传入该字段有效。详见华为推送官方文档消息分类标准。该字段优先级高于开发者后台为 App Key 下的应用标识配置的华为推送 Category。
	Category *string `json:"category,omitempty"`
}

type PushAndroidMi struct {
	// 小米推送通知渠道的 ID。详见小米推送消息分类新规。
	ChannelId *string `json:"channelId,omitempty"`
	// （由于小米官方已停止支持该能力，该字段已失效）消息右侧图标 URL，如果不设置，则不展示通知栏右侧图标。国内版仅 MIUI12 以上版本支持，以下版本均不支持；国际版支持。图片要求：大小120 * 120px，格式为 png 或者 jpg 格式。
	// Deprecated
	LargeIconUri *string `json:"large_icon_uri,omitempty"`
}

type PushAndroidOppo struct {
	// oppo 推送通知渠道的 ID。详见推送私信通道申请。
	ChannelId *string `json:"channelId,omitempty"`
}

type PushAndroidVivo struct {
	// VIVO 推送服务的消息类别。可选值 0（运营消息，默认值） 和 1（系统消息）。该参数对应 VIVO 推送服务的 classification 字段，见 VIVO 推送消息分类说明 。该字段优先级高于开发者后台为 App Key 下的应用标识配置的 VIVO 推送通道类型。
	Classification *string `json:"classification,omitempty"`
	// VIVO 推送服务的消息二级分类。例如 IM（即时消息）。该参数对应 VIVO 推送服务的 category 字段。详细的 category 取值请参见 VIVO 推送消息分类说明 。如果指定 category ，必须同时传入与当前二级分类匹配的 classification 字段的值（系统消息场景或运营消息场景）。请注意遵照 VIVO 官方要求，确保二级分类（category）取值属于 VIVO 系统消息场景或运营消息场景下允许发送的内容。该字段优先级高于开发者后台为 App Key 下的应用标识配置的 VIVO 推送 Category。
	Category *string `json:"category,omitempty"`
}

type PushAndroidFcm struct {
	// Google FCM 推送通知渠道的 ID。应用程序必须先创建一个具有此频道 ID 的频道，然后才能收到具有此频道 ID 的任何通知。更多信息请参见 Android 官方文档。
	ChannelId *string `json:"channelId,omitempty"`
	// Google FCM 推送中可以折叠的一组消息的标识符，以便在可以恢复传递时仅发送最后一条消息。
	CollapseKey *string `json:"collapse_key,omitempty"`
	// Google FCM 推送自定义的通知栏消息右侧图标 URL，如果不设置，则不展示通知栏右侧图标。
	//  图片的大小上限为 1MB。
	//  要求开发者后台 FCM 推送配置为证书与通知消息方式。
	ImageUrl *string `json:"imageUrl,omitempty"`
}

type PushCustomResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Id                 string `json:"id"` // 推送唯一标识。
}

// PushCustom 推送Plus专用推送接口,发送不落地通知
// More details see https://doc.rongcloud.cn/imserver/server/v1/push/push-plus
func (rc *RongCloud) PushCustom(ctx context.Context, req *PushCustomRequest) (*PushCustomResponse, error) {
	path := "/push/custom.json"
	resp := &PushCustomResponse{}
	httpResp, err := rc.postJson(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type PushUserRequest struct {
	// [必传] 用户 ID，每次发送时最多发送 100 个用户。
	UserIds []string `json:"userIds,omitempty"`
	// [必传] 按操作系统类型推送消息内容。
	Notification *PushUserNotification `json:"notification,omitempty"`
}

type PushUserNotification struct {
	// 通知栏显示标题，最长不超过 50 个字符。
	Title *string `json:"title,omitempty"`
	// [必传] 推送消息内容。
	PushContent *string                  `json:"pushContent,omitempty"`
	IOS         *PushUserNotificationIOS `json:"ios,omitempty"`
	Android     *PushAndroid             `json:"android,omitempty"`
}

// PushUserNotificationIOS
// Not same with PushCustomNotificationIOS, no Title field
type PushUserNotificationIOS struct {
	// 针对 iOS 平台，静默推送是 iOS7 之后推出的一种推送方式。 允许应用在收到通知后在后台运行一段代码，且能够马上执行。详情请查看知识库文档。1 表示为开启，0 表示为关闭，默认为 0
	ContentAvailable *int `json:"contentAvailable,omitempty"`
	// 应用角标，仅针对 iOS 平台；不填时，表示不改变角标数；为 0 或负数时，表示 App 角标上的数字清零；否则传相应数字表示把角标数改为指定的数字，最大不超过 9999，参数在 ios 节点下设置，详细可参考“设置 iOS 角标数 HTTP 请求示例”。
	Badge *int `json:"badge,omitempty"`
	// iOS 平台通知栏分组 ID，相同的 thread-id 推送分一组，单组超过 5 条推送会折叠展示
	ThreadId *string `json:"thread-id,omitempty"`
	// iOS 平台，从 iOS10 开始支持，设置后设备收到有相同 ID 的消息，会合并成一条
	ApnsCollapseId *string `json:"apns-collapse-id,omitempty"`
	// iOS 富文本推送的类型开发者自己定义，自己在 App 端进行解析判断，与 RichMediaUri 一起使用，当设置 Category 后，推送时默认携带 mutable-content 进行推送，属性值为 1。
	Category *string `json:"category,omitempty"`
	// iOS 富文本推送内容的 URL，与 category 一起使用。
	RichMediaUri *string `json:"richMediaUri,omitempty"`
	// 适用于 iOS 15 及之后的系统。取值为 passive，active（默认），time-sensitive，或 critical，取值说明详见对应的 APNs 的 interruption-level 字段。在 iOS 15 及以上版本中，系统的 “定时推送摘要”、“专注模式” 都可能导致重要的推送通知（例如余额变化）无法及时被用户感知的情况，可考虑设置该字段。
	InterruptionLevel *string `json:"interruption-level,omitempty"`
	// 附加信息，如果开发者自己需要，可以自己在 App 端进行解析。
	Extras interface{} `json:"extras,omitempty"`
}

type PushUserResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Id                 string `json:"id"` // 推送唯一标识。
}

// PushUser 发送指定用户不落地通知
// More details see https://doc.rongcloud.cn/imserver/server/v1/system/send-push-by-user
func (rc *RongCloud) PushUser(ctx context.Context, req *PushUserRequest) (*PushUserResponse, error) {
	path := "/push/user.json"
	resp := &PushUserResponse{}
	httpResp, err := rc.postJson(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}

type PushRequest struct {
	// [必传] 目标操作系统，iOS、Android 最少传递一个。如果需要给两个系统推送消息时，则需要全部填写，发送时如目标用户在 Web 端登录也会收到此条消息。
	Platform []string `json:"platform,omitempty"`
	// [必传] 发送人用户 ID。
	// 注意：发送消息所使用的用户 ID 必须已获取过用户 Token，否则消息一旦触发离线推送，通知内无法正确显示发送者的用户信息。
	FromUserid *string `json:"fromuserid,omitempty"`
	// [必传] 推送条件。支持按用户 ID 推送，按用户标签推送（tag 、tag_or）、按应用包名推送（packageName）和按指定平台全部推送（is_to_all）。注意：如果推送条件中 is_to_all 为 true，则忽略其他推送条件。
	Audience *PushAudience `json:"audience,omitempty"`
	// [必传] 消息类型参数的SDK封装, 例如: TXTMsg(文本消息), HQVCMsg(高清语音消息)
	Message RCMsg `json:"message,omitempty"`
	// 按操作系统类型推送通知内容，如 platform 中设置了给 iOS 和 Android 系统推送消息，而在 notification 中只设置了 iOS 的推送内容，则 Android 的推送内容为最初 alert 设置的内容。
	Notification *PushNotification `json:"notification,omitempty"`
}

func (r *PushRequest) MarshalJSON() ([]byte, error) {
	req := shadowPushRequest{
		Platform:     r.Platform,
		FromUserid:   r.FromUserid,
		Audience:     r.Audience,
		Notification: r.Notification,
	}
	if r.Message != nil {
		content, err := r.Message.ToString()
		if err != nil {
			return nil, fmt.Errorf("%s RCMsg.ToString() error %s", r.Message.ObjectName(), err)
		}
		req.Message = &shadowPushRequestMessage{
			ObjectName: r.Message.ObjectName(),
			Content:    content,
		}
	}
	return json.Marshal(req)
}

type shadowPushRequest struct {
	Platform     []string                  `json:"platform,omitempty"`
	FromUserid   *string                   `json:"fromuserid,omitempty"`
	Audience     *PushAudience             `json:"audience,omitempty"`
	Message      *shadowPushRequestMessage `json:"message,omitempty"`
	Notification *PushNotification         `json:"notification,omitempty"`
}

type shadowPushRequestMessage struct {
	ObjectName string `json:"objectName,omitempty"`
	Content    string `json:"content,omitempty"`
}

type PushAudience struct {
	// 用户标签，每次发送时最多发送 20 个标签，标签之间为 AND 的关系，is_to_all 为 true 时参数无效。
	Tag []string `json:"tag,omitempty"`
	// 用户标签，每次发送时最多发送 20 个标签，标签之间为 OR 的关系，is_to_all 为 true 时参数无效，tag_or 同 tag 参数可以同时存在。
	TagOr []string `json:"tag_or,omitempty"`
	// 用户 ID，每次发送时最多发送 1000 个用户，如果 tag 和 userid 两个条件同时存在时，则以 userid 为准，如果 userid 有值时，则 platform 参数无效，is_to_all 为 true 时参数无效。
	Userid []string `json:"userid,omitempty"`
	// 应用包名，is_to_all 为 true 时，此参数无效。与 tag、tag_or 同时存在时为 And 的关系，向同时满足条件的用户推送。与 userid 条件同时存在时，以 userid 为准进行推送。
	PackageName *string `json:"packageName,omitempty"`
	// 是否全部推送，false 表示按 tag 、tag_or 或 userid 条件推送，true 表示向所有用户推送，tag、tag_or 和 userid 条件无效。
	IsToAll *bool `json:"is_to_all,omitempty"`
}

type PushNotification struct {
	// 通知栏显示标题，最长不超过 50 个字符。
	Title *string `json:"title,omitempty"`
	// 是否越过客户端配置，强制在推送通知内显示通知内容（pushContent）。默认值 0 表示不强制，1 表示强制。
	// 说明：客户端设备可设置在接收推送通知时仅显示类似「您收到了一条通知」的提醒。从服务端发送消息时，可通过设置forceShowPushContent 为 1 越过该配置，强制客户端针在此条消息的推送通知中显示推送内容。
	ForceShowPushContent *int `json:"forceShowPushContent,omitempty"`
	// 推送通知内容。注意，如果此处 Alert 不传，则必须在 IOS.Alert 和 Android.Alert 分别指定 IOS 和 Android 下的推送通知内容，否则无法正常推送。一旦指定了各平台推送内容，则推送内容以对应平台系统的 alert 为准。如果都不填写，则无法发起推送。
	Alert *string `json:"alert,omitempty"`
	// 设置 IOS 平台下的推送及附加信息。
	IOS *PushNotificationIOS `json:"ios,omitempty"`
	// 设置 Android 平台下的推送及附加信息
	Android *PushNotificationAndroid `json:"android"`
}

type PushNotificationIOS struct {
	// 	通知栏显示的推送标题，仅针对 iOS 平台，支持 iOS 8.2 及以上版本。该属性优先级高于 PushNotification.Title。
	Title *string `json:"title,omitempty"`
	// 针对 iOS 平台，静默推送是 iOS7 之后推出的一种推送方式。 允许应用在收到通知后在后台运行一段代码，且能够马上执行。详情请查看知识库文档。1 表示为开启，0 表示为关闭，默认为 0
	ContentAvailable *int `json:"contentAvailable,omitempty"`
	// 推送通知内容，传入后默认的推送通知内容失效。
	Alert *string `json:"alert,omitempty"`
	// 应用角标，仅针对 iOS 平台；不填时，表示不改变角标数；为 0 或负数时，表示 App 角标上的数字清零；否则传相应数字表示把角标数改为指定的数字，最大不超过 9999，参数在 ios 节点下设置，详细可参考“设置 iOS 角标数 HTTP 请求示例”。
	Badge *int `json:"badge,omitempty"`
	// iOS 平台通知栏分组 ID，相同的 thread-id 推送分一组，单组超过 5 条推送会折叠展示
	ThreadId *string `json:"thread-id,omitempty"`
	// iOS 平台，从 iOS10 开始支持，设置后设备收到有相同 ID 的消息，会合并成一条
	ApnsCollapseId *string `json:"apns-collapse-id,omitempty"`
	// iOS 富文本推送的类型开发者自己定义，自己在 App 端进行解析判断，与 RichMediaUri 一起使用，当设置 Category 后，推送时默认携带 mutable-content 进行推送，属性值为 1。
	Category *string `json:"category,omitempty"`
	// iOS 富文本推送内容的 URL，与 category 一起使用。
	RichMediaUri *string `json:"richMediaUri,omitempty"`
	// 适用于 iOS 15 及之后的系统。取值为 passive，active（默认），time-sensitive，或 critical，取值说明详见对应的 APNs 的 interruption-level 字段。在 iOS 15 及以上版本中，系统的 “定时推送摘要”、“专注模式” 都可能导致重要的推送通知（例如余额变化）无法及时被用户感知的情况，可考虑设置该字段。
	InterruptionLevel *string `json:"interruption-level,omitempty"`
	// 附加信息，如果开发者自己需要，可以自己在 App 端进行解析。
	Extras interface{} `json:"extras,omitempty"`
}

type PushNotificationAndroid struct {
	// Android 平台下推送通知内容，传入后默认的推送通知内容失效。
	Alert  *string           `json:"alert,omitempty"`
	Honor  *PushAndroidHonor `json:"honor,omitempty"`
	HW     *PushAndroidHW    `json:"hw,omitempty"`
	Oppo   *PushAndroidOppo  `json:"oppo,omitempty"`
	Vivo   *PushAndroidVivo  `json:"vivo,omitempty"`
	Fcm    *PushAndroidFcm   `json:"fcm,omitempty"`
	Extras interface{}       `json:"extras,omitempty"`
}

type PushResponse struct {
	CodeResult
	httpResponseGetter `json:"-"`
	Id                 string `json:"id"` // 推送唯一标识。
}

// Push 发送全量用户不落地通知
// More details see https://doc.rongcloud.cn/imserver/server/v1/system/send-push-to-all
func (rc *RongCloud) Push(ctx context.Context, req *PushRequest) (*PushResponse, error) {
	path := "/push.json"
	resp := &PushResponse{}
	httpResp, err := rc.postJson(ctx, path, req, &resp)
	resp.httpResponseGetter = newRawHttpResponseGetter(httpResp)
	return resp, err
}
