package domain

// SqMember ...
type Member struct {
	ID uint32 `json:"id"`
	// Mobile 手机号
	Mobile string `json:"mobile"`
	// Role 用户级别3为军级团长，5为高级团长 7:团长 9:普通用户
	Role uint8 `json:"role"`
	// Role1 公司
	Role1 uint32 `json:"role1"`
	// Role3 军团长
	Role3 uint32 `json:"role3"`
	// Role5 高级团长
	Role5 uint32 `json:"role5"`
	// Role7 团长
	Role7 uint32 `json:"role7"`
	// Role9 普通用户
	Role9 uint32 `json:"role9"`
	// Username 用户名
	Username string `json:"username"`
	// Nickname 昵称
	Nickname string `json:"nickname"`
	// Realname 真实姓名
	Realname string `json:"realname"`
	// XcxOpenid 小程序openid
	XcxOpenid string `json:"xcx_openid"`
	// SqUnionid 关联社区系统的unionid
	SqUnionid string `json:"sq_unionid"`
	// Unionid unionid
	Unionid string `json:"unionid"`
	// Openid 公众号openid
	Openid string `json:"openid"`
	// AppOpenid app_openid
	AppOpenid string `json:"app_openid"`
	// Password 密码
	Password string `json:"password"`
	// Salt 登录密钥
	Salt string `json:"salt"`
	// Avatar 头像
	Avatar string `json:"avatar"`
	// UpgradeStatus 是否满足升级条件 0=不满足 1=满足
	UpgradeStatus uint8 `json:"upgrade_status"`
	// LastUpgradeTime 最后升级时间
	LastUpgradeTime uint32 `json:"last_upgrade_time"`
	// Status 账户状态, -1=已被删除 1=正常，2待审核，3已拒绝，4待激活，9拉黑
	Status uint8 `json:"status"`
	// TUserid 推荐人ID
	TUserid uint32 `json:"t_userid"`
	// TUsername 推荐人姓名
	TUsername string `json:"t_username"`
	// TRole 推荐人角色
	TRole uint32 `json:"t_role"`
	// Parentid 上级ID
	Parentid uint32 `json:"parentid"`
	// PRole 上级角色
	PRole uint32 `json:"p_role"`
	// PUsername 上级姓名
	PUsername string `json:"p_username"`
	// PjRole 平级级别
	PjRole int32 `json:"pj_role"`
	// PjUserid 平级ID
	PjUserid int32 `json:"pj_userid"`
	// PjUsername 平级用户名
	PjUsername string `json:"pj_username"`
	// Isactivate 0未激活，1已激活
	Isactivate int8 `json:"isactivate"`
	// InviteCode 邀请码 10000+用户编号
	InviteCode string `json:"invite_code"`
	// RegisterTime 注册时间
	RegisterTime int32 `json:"register_time"`
	// Wechat 微信号
	Wechat string `json:"wechat"`
	// Teamid 团队ID
	Teamid int32 `json:"teamid"`
	// Bondmoney 累积保证金金额
	Bondmoney float32 `json:"bondmoney"`
	// Money 余额
	Money float32 `json:"money"`
	// Integral 用户积分
	Integral int32 `json:"integral"`
	// ReceiveIntegral 用户已兑换积分
	ReceiveIntegral int32 `json:"receive_integral"`
	// Recharge 充值余额
	Recharge float32 `json:"recharge"`
	// MoneyRecharge 累计充值金额
	MoneyRecharge float32 `json:"money_recharge"`
	// MoneyBonus 累计提成
	MoneyBonus float32 `json:"money_bonus"`
	// Resource 用户来源 1=普通注册 2=社区转换 3=邀请注册 4=盖诺代理
	Resource uint8 `json:"resource"`
	// Createtime 创建时间
	Createtime uint32 `json:"createtime"`
	// Lastupdatetime 最后更新时间
	Lastupdatetime uint32 `json:"lastupdatetime"`
	// Newrole 升级新等级 isactivate=0时有效
	Newrole uint32 `json:"newrole"`
	// PingjiIndex 1:正常 2:违背
	PingjiIndex uint8 `json:"pingji_index"`
	// Effective 是否有效团长-完成200订单,1为是
	Effective uint8 `json:"effective"`
	// ReceiveStatus 面膜领取状态
	ReceiveStatus int8 `json:"receive_status"`
	// GynoRole 盖诺等级
	GynoRole int32 `json:"gyno_role"`
	// GynoID 盖诺代理ID
	GynoID int32 `json:"gyno_id"`
	// CashMoney 代金券余额
	CashMoney float32 `json:"cash_money"`
	// IDCardNo 身份证号码
	IDCardNo string `json:"id_card_no"`
	// IDCardName 身份证姓名
	IDCardName string `json:"id_card_name"`
	// IDCardImgTop 身份证照片正面
	IDCardImgTop string `json:"id_card_img_top"`
	// IDCardImgBottom 身份证照片反面
	IDCardImgBottom string `json:"id_card_img_bottom"`
	// IsCertification 是否实名认证 0：否 1：是
	IsCertification int8 `json:"is_certification"`
	// IDCardTimeStart 身份证有效期开始时间
	IDCardTimeStart int32 `json:"id_card_time_start"`
	// IDCardTimeEnd 身份证有效期结束时间
	IDCardTimeEnd int32 `json:"id_card_time_end"`
	// ActivityMsgOpen 推荐消息开关 0：打开 1：关闭
	ActivityMsgOpen int8 `json:"activity_msg_open"`
}
