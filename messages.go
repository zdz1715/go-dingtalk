package dingtalk

import (
	"context"
	"net/http"
)

type MessagesService service

type AsyncSendCorpConversationOptions struct {
	AgentId    *string     `json:"agent_id,omitempty"`
	UseridList *string     `json:"userid_list,omitempty"`
	DeptIdList *string     `json:"dept_id_list,omitempty"`
	ToAllUser  *bool       `json:"to_all_user,omitempty"`
	Msg        *MsgOptions `json:"msg,omitempty" query:"msg"`
}

type MsgOptions struct {
	Msgtype    *MsgType              `json:"msgtype,omitempty"`
	Text       *TextMsgOptions       `json:"text,omitempty"`
	Image      *ImageMsgOptions      `json:"image,omitempty"`
	Voice      *VoiceMsgOptions      `json:"voice,omitempty"`
	File       *FileMsgOptions       `json:"file,omitempty"`
	Link       *LinkMsgOptions       `json:"link,omitempty"`
	OA         *OAMsgOptions         `json:"oa,omitempty"`
	Markdown   *MarkdownMsgOptions   `json:"markdown,omitempty"`
	ActionCard *ActionCardMsgOptions `json:"action_card,omitempty"`
}

type TextMsgOptions struct {
	Content *string `json:"content,omitempty"`
}

type ImageMsgOptions struct {
	MediaId *string `json:"media_id,omitempty"`
}

type VoiceMsgOptions struct {
	MediaId  *string `json:"media_id,omitempty"`
	Duration *string `json:"duration,omitempty"`
}

type FileMsgOptions struct {
	MediaId *string `json:"media_id,omitempty"`
}

type LinkMsgOptions struct {
	MessageUrl *string `json:"messageUrl,omitempty"`
	Title      *string `json:"title,omitempty"`
	PicUrl     *string `json:"picUrl,omitempty"`
	Text       *string `json:"text,omitempty"`
}

type MarkdownMsgOptions struct {
	Title *string `json:"title,omitempty"`
	Text  *string `json:"text,omitempty"`
}

type ActionCardMsgOptionsBtn struct {
	Title     *string `json:"title,omitempty"`
	ActionURL *string `json:"action_url,omitempty"`
}

type ActionCardMsgOptions struct {
	Title          *string                    `json:"title,omitempty"`
	Markdown       *string                    `json:"markdown,omitempty"`
	SingleTitle    *string                    `json:"single_title,omitempty"`
	SingleURL      *string                    `json:"single_url,omitempty"`
	BtnOrientation *string                    `json:"btn_orientation,omitempty"`
	BtnJsonList    []*ActionCardMsgOptionsBtn `json:"btn_json_list,omitempty"`
}

type OAMsgOptionsHeadStatusBar struct {
	StatusValue *string `json:"status_value,omitempty"`
	StatusBg    *string `json:"status_bg,omitempty"`
}

type OAMsgOptionsHead struct {
	Bgcolor   *string                    `json:"bgcolor,omitempty"`
	Text      *string                    `json:"text,omitempty"`
	StatusBar *OAMsgOptionsHeadStatusBar `json:"status_bar,omitempty"`
}

type OAMsgOptionsBodyForm struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type OAMsgOptionsBodyRich struct {
	Num  *string `json:"num,omitempty"`
	Unit *string `json:"unit,omitempty"`
}

type OAMsgOptions struct {
	MessageUrl   *string           `json:"message_url,omitempty"`
	PcMessageUrl *string           `json:"pc_message_url,omitempty"`
	Head         *OAMsgOptionsHead `json:"head,omitempty"`
	Body         *OAMsgOptionsBody `json:"body,omitempty"`
}

type OAMsgOptionsBody struct {
	Title     *string                 `json:"title,omitempty"`
	Author    *string                 `json:"author,omitempty"`
	Content   *string                 `json:"content,omitempty"`
	Image     *string                 `json:"image,omitempty"`
	FileCount *string                 `json:"file_count,omitempty"`
	Form      []*OAMsgOptionsBodyForm `json:"form,omitempty"`
	Rich      *OAMsgOptionsBodyRich   `json:"rich,omitempty"`
}

type Task struct {
	TaskId int64 `json:"task_id"`
}

func (s *MessagesService) AsyncSendCorpConversation(ctx context.Context, opts *AsyncSendCorpConversationOptions) (*Task, error) {
	const apiEndpoint = "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2"
	var respBody Task
	if err := s.client.InvokeByToken(ctx, http.MethodPost, apiEndpoint, opts, &respBody); err != nil {
		return nil, err
	}
	return &respBody, nil
}
