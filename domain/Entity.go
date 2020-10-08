package domain

// WdMemberList ...
type MemberList struct {
	// MemberListID 主键id
	MemberListID int `json:"id" binding:"required" form:"id"`
	// MemberListUsername 账号
	MemberListUsername string `json:"-"`
	// Role 用户角色
	Role uint8 `json:"role"`
	// MemberListPwd 密码
	MemberListPwd string `json:"-"`
	// MemberListSalt 密钥
	MemberListSalt string `json:"-"`
	// MemberListGroupid 组ID
	MemberListGroupid int32 `json:"-"`
	// MemberListNickname 昵称
	MemberListNickname string `json:"-"`
	// MemberListProvince 省ID
	MemberListProvince int32 `json:"-"`
	// MemberListCity 市ID
	MemberListCity int32 `json:"-"`
	// MemberListTown 区ID
	MemberListTown int32 `json:"-"`
	// MemberListSex 性别
	MemberListSex string `json:"-"`
	// MemberListHeadpic 头像
	MemberListHeadpic string `json:"-"`
	// MemberListTel 电话
	MemberListTel string `json:"-"`
	// MemberListWx 微信
	MemberListWx string `json:"-"`
	// MemberListEmail 邮箱
	MemberListEmail string `json:"-"`
	// MemberListOpen 是否开启
	MemberListOpen int32 `json:"-"`
	// MemberListAddtime 注册时间
	MemberListAddtime int32 `json:"-"`
	// MemberListFrom 从哪来
	MemberListFrom string `json:"-"`
	// UserURL 用户URL
	UserURL string `json:"-"`
	// Birthday 生日
	Birthday string `json:"-"`
	// MemberListGroupcaiwuid 财务组id
	MemberListGroupcaiwuid int32 `json:"-"`
	// Signture 签名
	Signture string `json:"-"`
	// LastLoginIP 上次登陆IP
	LastLoginIP string `json:"-"`
	// LastLoginTime 上次登陆时间
	LastLoginTime int32 `json:"-"`
	// UserActivationKey 用户激活KEY
	UserActivationKey string `json:"-"`
	// UserStatus 用户状态（0未处理1拒绝2通过）
	UserStatus int32 `json:"user_status"`
	// Isblack 是否黑名单
	Isblack int32 `json:"isblack"`
	// Score 分数
	Score int32 `json:"-"`
	// Coin 金币
	Coin int32 `json:"-"`
	// IDType 1 身份证，2港澳台通行证，3护照
	IDType int32 `json:"-"`
	// IDCard 身份证
	IDCard string `json:"-"`
	// IDImages 证件照
	IDImages string `json:"-"`
	// Photo 照片
	Photo string `json:"-"`
	// Realname 姓名
	Realname string `json:"realname"`
	// Bond 保证金
	Bond int32 `json:"-"`
	// Parentid 上级ID
	Parentid int32 `json:"parentid"`
	// PRole 上级角色
	PRole int32 `json:"p_role"`
	// PUsername 上级姓名
	PUsername string `json:"p_username"`
	// TUserid 推荐人ID
	TUserid int `json:"t_userid"`
	// TRole 推荐人角色
	TRole int32 `json:"t_role"`
	// TUsername 推荐人姓名
	TUsername string `json:"t_username"`
	// InvitationCode 邀请码
	InvitationCode string `json:"-"`
	// MyCodeqr 邀请二维码图片
	MyCodeqr string `json:"-"`
	// MyCode 我的邀请码
	MyCode string `json:"-"`
	// Role9 公司
	Role9 int32 `json:"-"`
	// Role10 官方
	Role10 int32 `json:"role10"`
	// Role11 荣誉官方
	Role11 int32 `json:"role11"`
	// Role12 总代
	Role12 int32 `json:"role12"`
	// Role13 省级
	Role13 int32 `json:"role13"`
	// Role14 市级
	Role14 int32 `json:"role14"`
	// Role15 特约
	Role15 int32 `json:"role15"`
	// Role16 无
	Role16 int32 `json:"role16"`
	// Country 国家
	Country string `json:"-"`
	// Province 省
	Province string `json:"-"`
	// City 市
	City string `json:"-"`
	// Area 县区
	Area string `json:"-"`
	// Address 详细地址
	Address string `json:"-"`
	// Memberno 代理编号
	Memberno int32 `json:"-"`
	// Isdelete 是否逻辑删除
	Isdelete int32 `json:"-"`
	// Isown 是不是属于公司直管 1 为是
	Isown int32 `json:"-"`
	// Reason 审核拒绝理由
	Reason string `json:"-"`
	// Job 1兼职2全职
	Job int32 `json:"-"`
	// MyPaybackIncome 我的总返点收入
	MyPaybackIncome float32 `json:"-"`
	// MyPaybackExpend 我的总返点支出
	MyPaybackExpend float32 `json:"-"`
	// CardPic1 身份证正面照
	CardPic1 string `json:"-"`
	// CardPic2 身份证反面照
	CardPic2 string `json:"-"`
	// Lastchangeroletime 最后一次改变级别时间
	Lastchangeroletime int32 `json:"-"`
	// RealnameRequest 实名认证申请
	RealnameRequest int32 `json:"-"`
	// Isrealname 2是通过实名
	Isrealname int32 `json:"-"`
	// Totalguige 总进货规格数
	Totalguige int32 `json:"-"`
	// Allyeji 所有业绩
	Allyeji float32 `json:"-"`
	// Selfyeji 自己的销量
	Selfyeji float32 `json:"-"`
	// Pingjiyeji 平级业绩
	Pingjiyeji float32 `json:"-"`
	// Shoudan 1首单2非首单3锁定非首单
	Shoudan int32 `json:"-"`
	// Taishoudan 是否肽首单 1是0否
	Taishoudan uint32 `json:"-"`
	// Money 余额
	Money float32 `json:"-"`
	// Profit 总利润
	Profit float32 `json:"-"`
	// Profitsale 零售利润
	Profitsale float32 `json:"-"`
	// Profitstore 出货利润
	Profitstore float32 `json:"-"`
	// Isold 0 不是，1是平移会员
	Isold int32 `json:"-"`
	// Lastjinhuotime 最后进货时间
	Lastjinhuotime int32 `json:"-"`
	// Isactive 0，未激活，1激活
	Isactive int32 `json:"isactive"`
	// Ishalf 0 不半价，1 半价门槛
	Ishalf int32 `json:"-"`
	// Teamid 团队ID
	Teamid int32 `json:"-"`
	// Nofandian 是否不返点
	Nofandian int32 `json:"-"`
	// Alloryeji 出货零售总业绩
	Alloryeji float32 `json:"-"`
	// Selfoutyeji 自己出货总业绩
	Selfoutyeji float32 `json:"-"`
	// Pingjioutyeji 平级总出货业绩
	Pingjioutyeji float32 `json:"-"`
	// Selfresaleyeji 自己零售总业绩
	Selfresaleyeji float32 `json:"-"`
	// Pingjiresaleyeji 平级零售总业绩
	Pingjiresaleyeji float32 `json:"-"`
	// VerifyBookURL 保存生成后的证书路径
	VerifyBookURL string `json:"-"`
	// GqTotal 枸杞累计进货
	GqTotal int32 `json:"-"`
	// GqUnreceived 枸杞未领取
	GqUnreceived int32 `json:"-"`
	// GqReceived 枸杞已领取
	GqReceived int32 `json:"-"`
	// Qichu 期初
	Qichu int32 `json:"-"`
	// CTotal 推荐人数（总数）
	CTotal int32 `json:"-"`
	// Oldrole12 升级前的总代
	Oldrole12 int32 `json:"-"`
	// Firstjinhuotime 第一次进货时间
	Firstjinhuotime int32 `json:"-"`
	// Firstlingshotime 第一次零售时间
	Firstlingshotime int32 `json:"-"`
	// MyJifu 集福二维码图片
	MyJifu string `json:"-"`
	// Integral 稽查积分
	Integral int32 `json:"-"`
	// LearningIntegral 学习积分
	LearningIntegral int32 `json:"-"`
	// Allhuangouprice 总换购金额
	Allhuangouprice float32 `json:"-"`
	// Xbfydprice 学霸护眼灯总换购金额
	Xbfydprice float32 `json:"-"`
	// Ajgprice 富硒玫瑰阿胶糕总换购金额
	Ajgprice float32 `json:"-"`
	// Isazhalf 阿祖苗药是否半门槛：0 不半门槛，1 半门槛
	Isazhalf uint32 `json:"-"`
	// WelcomeName 店铺欢迎语
	WelcomeName string `json:"-"`
	// StoreName 店铺名称
	StoreName string `json:"-"`
	// StoreAvatar 店铺头像
	StoreAvatar string `json:"-"`
	// StoreCodeqr 店铺二维码
	StoreCodeqr string `json:"-"`
	// StoreOpenid 店铺绑定小程序openid
	StoreOpenid string `json:"-"`
	// DianjiCount 人气
	DianjiCount int64 `json:"-"`
	// KefuTel 客服电话
	KefuTel string `json:"-"`
	// KefuWx 客服微信
	KefuWx string `json:"-"`
	// KefuCodeqr 客服二维码
	KefuCodeqr string `json:"-"`
	// Yfcount 免邮次数
	Yfcount int32 `json:"-"`
	// Jb 特邀嘉宾
	Jb int32 `json:"-"`
	// MemberListFromid 从哪来用户id
	MemberListFromid uint32 `json:"-"`
	// MergeMoney BBC合并金额
	MergeMoney float32 `json:"-"`
	// StoreMoney 代理库存余额
	StoreMoney float32 `json:"-"`
	// CanExchange 是否允许换货 1是0否
	CanExchange uint8 `json:"-"`
}
