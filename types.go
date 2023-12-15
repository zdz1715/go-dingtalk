package dingtalk

type AuthType uint8

type Language string

const (
	ZH_CN Language = "zh_CN"
	EN_US Language = "en_US"
)

// RobotType
// docs: https://open.dingtalk.com/document/orgapp/obtain-the-list-of-robots-in-the-group
type RobotType uint8

const (
	PublicRobotType RobotType = iota + 1
	InternalRobotType
	ThirdPartiesRobotType
	SceneGroupRobotType
)

// MsgType custom robot msg type
// docs: https://open.dingtalk.com/document/orgapp/custom-robots-send-group-messages#title-72m-8ag-pqw
type MsgType string

const (
	TextMsgType       MsgType = "text"
	LinkMsgType       MsgType = "link"
	MarkdownMsgType   MsgType = "markdown"
	ActionCardMsgType MsgType = "actionCard"
	FeedCardMsgType   MsgType = "feedCard"
)
