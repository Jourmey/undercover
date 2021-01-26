package common

const (
	Role_Normal     = "Normal"     // 正常
	Role_Undercover = "Undercover" // 卧底

	GameStage_Start = "Start" // 准备阶段
	GameStage_Vote  = "Vote"  // 投票阶段
	GameStage_Game  = "Game"  // 游戏阶段
	GameStage_Over  = "Over"  // 完成阶段
)

type Ping struct {
	Value string
}

// 用户
type User struct {
	Openid string // 标识
	No     string // 序号
	Name   string // 用户名
	Status int    // 状态
	RoomId string // 房间号
}

// 词语
type Word struct {
	Id          int    // Id
	Word        string // 词
	AnotherWord string // 另一个词
}

// 消息
type GameMessage struct {
	UserId   string // 用户Id
	UserName string // 用户名称
	RoomId   string // 房间id
	Msg      string // 消息
	Status   int    // 状态
	Data     interface{}
	Type     string
}

func NewErrorGameMessage(e error) *GameMessage {
	m := &GameMessage{
		Msg:    e.Error(),
		Status: 0,
	}
	return m
}
func NewSuccessGameMessage(msg string) *GameMessage {
	m := &GameMessage{
		Msg:    msg,
		Status: 1,
	}
	return m
}

func (g *GameMessage) WithType(t string) *GameMessage {
	g.Type = t
	return g
}
func (g *GameMessage) WithData(t interface{}) *GameMessage {
	g.Data = t
	return g
}

type Login struct {
	UserName string // 用户名称
	UserId   string // 用户名称
}

//type Logout struct {
//	UserName string // 用户名称
//	UserId   string // 用户名称
//}

type UserOut struct {
	UserId string
}

type Room struct {
	Id               int
	CreateUserId     string // 用户ID
	RoomId           string // 房间ID
	Password         string // 房间密码
	TotalNumber      string // 总人数
	Number           int    // 当前人数
	UndercoverNumber string // 卧底人数
	IsPrepare        bool   //是否准备
}

type RoomOut struct {
	RoomId string
	UserId string
}

// 游戏
type Game struct {
	Round            int              // 回合数
	SurvivalUserList map[string]*User // 存活用户列表
	Keyword          *Keyword         // 词语
	UndercoverNum    int              // 卧底数量
	Stage            string           // 阶段
	ActionTime       int              // 操作时间
	VoteTime         int              // 投票等待时间 (秒)
	RoomId           string           // 房间号
	WinRole          string           // 胜利方
}

type Vote struct {
	Round            int
	UserId           string
	VotePlayerNumber string
	IsPrepare        bool
	RoomId           string
	GameId           string
}

type Keyword struct {
	Code           string
	NormalWord     string
	UndercoverWord string
	Vension        int64
}

type KeywordResult struct {
	Keyword      string
	Stage        string
	CreateUserId string
}
