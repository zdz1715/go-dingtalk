package dingtalk

import (
	"context"
	"fmt"
	"net/http"
)

type RobotsService service

type ListRobotsInGroupOptions struct {
	OpenConversationId *string `json:"openConversationId,omitempty"`
}

type Robot struct {
	RobotCode       string    `json:"robotCode"`
	Name            string    `json:"name"`
	DownloadIconURL string    `json:"downloadIconURL"`
	OpenRobotType   RobotType `json:"openRobotType"`
}

type ListRobotsInGroupResult struct {
	ChatbotInstanceVOList []*Robot `json:"chatbotInstanceVOList"`
}

// ListRobotsInGroup gets a list of robot in group
// API docs: https://open.dingtalk.com/document/orgapp/obtain-the-list-of-robots-in-the-group
func (s *RobotsService) ListRobotsInGroup(ctx context.Context, opts *ListRobotsInGroupOptions) (*ListRobotsInGroupResult, error) {
	const apiEndpoint = "/v1.0/robot/getBotListInGroup"
	var respBody ListRobotsInGroupResult
	if err := s.client.InvokeByToken(ctx, http.MethodPost, apiEndpoint, opts, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}

type RobotTextMsgOptions struct {
	Content *string `json:"content,omitempty"`
}

type RobotMarkdownMsgOptions struct {
	Title *string `json:"title,omitempty"`
	Text  *string `json:"text,omitempty"`
}

type RobotLinkMsgOptions struct {
	MessageUrl *string `json:"messageUrl,omitempty"`
	Title      *string `json:"title,omitempty"`
	PicUrl     *string `json:"picUrl,omitempty"`
	Text       *string `json:"text,omitempty"`
}

type RobotSendOptionsAt struct {
	IsAtAll   *bool    `json:"isAtAll,omitempty"`
	AtMobiles []string `json:"atMobiles,omitempty"`
	AtUserIds []string `json:"atUserIds,omitempty"`
}

type RobotActionCardMsgOptions struct {
	HideAvatar     *string                         `json:"hideAvatar,omitempty"`
	BtnOrientation *string                         `json:"btnOrientation,omitempty"`
	SingleURL      *string                         `json:"singleURL,omitempty"`
	SingleTitle    *string                         `json:"singleTitle,omitempty"`
	Text           *string                         `json:"text,omitempty"`
	Title          *string                         `json:"title,omitempty"`
	Btns           []*RobotActionCardMsgOptionsBtn `json:"btns,omitempty"`
}

type RobotActionCardMsgOptionsBtn struct {
	Title     *string `json:"title,omitempty"`
	ActionURL *string `json:"actionURL,omitempty"`
}

type RobotFeedCardMsgOptions struct {
	Links []*RobotFeedCardMsgOptionsLink `json:"links,omitempty"`
}

type RobotFeedCardMsgOptionsLink struct {
	MessageURL *string `json:"messageURL,omitempty"`
	Title      *string `json:"title,omitempty"`
	PicURL     *string `json:"picURL,omitempty"`
}

type RobotSendOptions struct {
	Msgtype    *RobotMsgType              `json:"msgtype,omitempty"`
	At         *RobotSendOptionsAt        `json:"at,omitempty"`
	Link       *RobotLinkMsgOptions       `json:"link,omitempty"`
	Text       *RobotTextMsgOptions       `json:"text,omitempty"`
	Markdown   *RobotMarkdownMsgOptions   `json:"markdown,omitempty"`
	ActionCard *RobotActionCardMsgOptions `json:"actionCard,omitempty"`
	FeedCard   *RobotFeedCardMsgOptions   `json:"feedCard,omitempty"`
}

// Send messages in group by custom robots
// API docs: https://open.dingtalk.com/document/orgapp/custom-robots-send-group-messages
func (s *RobotsService) Send(ctx context.Context, accessToken string, opts *RobotSendOptions) error {
	var apiEndpoint = fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", accessToken)
	return s.client.Invoke(ctx, http.MethodPost, apiEndpoint, opts, nil, "")
}
