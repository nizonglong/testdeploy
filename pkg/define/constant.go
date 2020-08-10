package define

const (
    IntZero  = 0
    IntOne   = 1
    IntTwo   = 2
    IntThree = 3

    StrGameID = "game_id"
    StrType   = "type"
)

const (
    ErrParam = "參數錯誤"
    ErrDB    = "數據庫操作出錯"
    ErrCache = "緩存操作出錯"
    StatusOk = "successful"
)

// db连接关键字
const (
    QiPaiDB     = "qipaidb"
    StatDB      = "statdb"
    ConfigCache = "gameconfigcache"
    // 中间件
    MiddlewareUserCache = "usercache"
    MiddlewareLimit     = "limit"

    StrConfig = "config"
)

// 数据库表名
const (
    TableGameConfig = "game_config"
)

// 游戏配置类型
const (
    ConfigTypeLoading           = 1  // 加载界面加载文字配置
    ConfigTypeMarquee           = 2  // 跑马灯配置
    ConfigTypeADCoin            = 3  // 观看广告后金币配置
    ConfigTypeShopBanner        = 4  // 商城轮播图配置
    ConfigTypeEmail             = 5  // 邮箱配置
    ConfigTypeCommunity         = 6  // 社区内容配置
    ConfigTypeShopDailyFreeCoin = 7  // 商城每日免费领取金币配置，数额次数
    ConfigTypeADCount           = 8  // 观看广告次数配置
    ConfigTypeTimeZone          = 9  // 时区配置
    ConfigTypeUnLock            = 10 // 德州解锁聚宝盆，每日任务
    ConfigTypeBonus             = 11 // 聚宝盆能量加成
    ConfigTypeGuild             = 12 // 新手引导
)

// 错误提示
const (
    RetryLaterError      = "稍後重試"
    DataNotFit           = "数据不符合要求"
    OperationFrequentTip = "操作頻繁，請稍等"
    WatchADLaterTip      = "請稍後觀看廣告"
)

// 年月日常量
const (
    YMDLayout         = "2006-01-02"
    LayoutDayForCache = "20060102"
)
