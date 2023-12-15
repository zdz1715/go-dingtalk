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

type SendOptionsText struct {
	Content *string `json:"content,omitempty"`
}

type SendOptionsMarkdown struct {
	Title *string `json:"title,omitempty"`
	Text  *string `json:"text,omitempty"`
}

type SendOptionsLink struct {
	MessageUrl *string `json:"messageUrl,omitempty"`
	Title      *string `json:"title,omitempty"`
	PicUrl     *string `json:"picUrl,omitempty"`
	Text       *string `json:"text,omitempty"`
}

type SendOptionsAt struct {
	IsAtAll   *bool    `json:"isAtAll,omitempty"`
	AtMobiles []string `json:"atMobiles,omitempty"`
	AtUserIds []string `json:"atUserIds,omitempty"`
}

type SendOptionsActionCard struct {
	HideAvatar     *string                    `json:"hideAvatar,omitempty"`
	BtnOrientation *string                    `json:"btnOrientation,omitempty"`
	SingleURL      *string                    `json:"singleURL,omitempty"`
	SingleTitle    *string                    `json:"singleTitle,omitempty"`
	Text           *string                    `json:"text,omitempty"`
	Title          *string                    `json:"title,omitempty"`
	Btns           []SendOptionsActionCardBtn `json:"btns,omitempty"`
}

type SendOptionsActionCardBtn struct {
	Title     *string `json:"title,omitempty"`
	ActionURL *string `json:"actionURL,omitempty"`
}

type SendOptionsFeedCard struct {
	Links []SendOptionsFeedCardLink `json:"links,omitempty"`
}

type SendOptionsFeedCardLink struct {
	MessageURL *string `json:"messageURL,omitempty"`
	Title      *string `json:"title,omitempty"`
	PicURL     *string `json:"picURL,omitempty"`
}

type SendOptions struct {
	Msgtype    *MsgType               `json:"msgtype,omitempty"`
	At         *SendOptionsAt         `json:"at,omitempty"`
	Link       *SendOptionsLink       `json:"link,omitempty"`
	Text       *SendOptionsText       `json:"text,omitempty"`
	Markdown   *SendOptionsMarkdown   `json:"markdown,omitempty"`
	ActionCard *SendOptionsActionCard `json:"actionCard,omitempty"`
	FeedCard   *SendOptionsFeedCard   `json:"feedCard,omitempty"`
}

// Send messages in group by custom robots
// API docs: https://open.dingtalk.com/document/orgapp/custom-robots-send-group-messages
func (s *RobotsService) Send(ctx context.Context, accessToken string, opts *SendOptions) error {
	var apiEndpoint = fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", accessToken)
	return s.client.Invoke(ctx, http.MethodPost, apiEndpoint, opts, nil, "")
}
