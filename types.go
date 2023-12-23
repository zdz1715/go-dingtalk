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

// RobotMsgType custom robot msg type
// docs: https://open.dingtalk.com/document/orgapp/custom-robots-send-group-messages#title-72m-8ag-pqw
type RobotMsgType string

const (
	TextRobotMsgType       RobotMsgType = "text"
	LinkRobotMsgType       RobotMsgType = "link"
	MarkdownRobotMsgType   RobotMsgType = "markdown"
	ActionCardRobotMsgType RobotMsgType = "actionCard"
	FeedCardRobotMsgType   RobotMsgType = "feedCard"
)

// MsgType msg type
// docs: https://open.dingtalk.com/document/isvapp/message-types-and-data-format?spm=ding_open_doc.document.0.0.25e12580NcSonF#title-x16-76n-jpg
type MsgType string

const (
	TextMsgType       MsgType = "text"
	ImageMsgType      MsgType = "image"
	VoiceMsgType      MsgType = "voice"
	FileMsgType       MsgType = "file"
	LinkMsgType       MsgType = "link"
	OAMsgType         MsgType = "oa"
	MarkdownMsgType   MsgType = "markdown"
	ActionCardMsgType MsgType = "action_card"
)
