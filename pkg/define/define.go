package define

type GameConfig struct {
	ID         int16  `json:"id" form:"id"`
	Data       string `json:"data" form:"data"`
	Type       int16  `json:"type" form:"type"`
	GameID     int16  `json:"game_id" form:"game_id"`
	GameAreaID int16  `json:"game_area_id" form:"game_area_id"`
}

type DeleteGameConfigRequest struct {
	GameID       int16 `json:"game_id" form:"game_id" binding:"required"`
	GameConfigID int16 `json:"game_config_id" form:"game_config_id" binding:"required"`
}

type GetGameConfigRequest struct {
	GameID int16 `json:"game_id" form:"game_id" binding:"required"`
	Type   int16 `json:"type" form:"type" binding:"required"`
}

type GetShopFreeCoinRequest struct {
	GameID int16 `json:"game_id" form:"game_id" binding:"required"`
	UserID int32 `json:"user_id" form:"user_id"`
}

/**
 * ------------------------------------------------
 * 游戏配置，从数据库的data进行转换
 * ------------------------------------------------
 */
// 加载界面加载文字配置
type LoadingWordConfig struct {
	ID      int16  `json:"id" form:"id"`
	Content string `json:"content" form:"content"` // 加载文字内容
}

// 轮播图
type BannerConfig struct {
	Title    string `json:"title" form:"title"`
	ImgUrl   string `json:"img_url" form:"img_url"`   // 图片地址
	JumpUrl  string `json:"jump_url" form:"jump_url"` // 跳转url
	Position int16  `json:"position" form:"position"` // 轮播图位置
}

// 观看广告后金币配置
type ADWatchCoinConfig struct {
	ReceiveCount int32 `json:"receive_count"`
	GoldCount    int32 `json:"gold_count"`
}

// 观看广告次数配置
type ADWatchCountConfig struct {
	MinPayLimit     int16 `json:"min_pay_limit"`
	MaxPayLimit     int16 `json:"max_pay_limit"`
	MinPlayDay      int16 `json:"min_play_day"`
	MaxPlayDay      int16 `json:"max_play_day"`
	MaxReceiveCount int32 `json:"max_receive_count"`
}

// 跑马灯，横幅
type MarqueeConfig struct {
	ID          int16  `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	PropType    int32  `json:"prop_type" form:"prop_type"`
	ElementType int32  `json:"element_type" form:"element_type"`
	PropNum     int32  `json:"prop_num" form:"prop_num"`
}

// 邮件内容配置
type EmailConfig struct {
	ID        int16  `json:"id" form:"id"`
	Title     string `json:"title" form:"title"`
	Type      int16  `json:"type"`
	Content   string `json:"content" form:"content"` // 邮件内容
	Props     []Prop `json:"props"`
	Signature string `json:"signature"` // 邮件签名
}

type Prop struct {
	PropType    int32 `json:"prop_type"`
	ElementType int32 `json:"element_type"`
	AwardTotal  int32 `json:"award_total"`
	PropName    string `json:"prop_name"`
}

// 社区内容配置
type CommunityConfig struct {
	ID      int16  `json:"id" form:"id"`
	Icon    string `json:"icon" form:"icon"`         // 社区图标配置
	Content string `json:"content" form:"content"`   // 内容
	JumpUrl string `json:"jump_url" form:"jump_url"` // 跳转url
	Props   []Prop `json:"props"`
}

// 时区配置
type TimeZone struct {
	ID                    int16  `json:"id" form:"id"`
	CountryName           string `json:"country_name"`
	ProvincialCapitalCode string `json:"provincial_capital_code"`
}

// 新手引导配置
type NewGuildConfig struct {
    ID    int16  `json:"id"`
    Type  int16  `json:"type"`
    Props []Prop `json:"props"`
}