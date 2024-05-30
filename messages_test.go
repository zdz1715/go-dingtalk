package dingtalk

import (
	"context"
	"testing"

	"github.com/zdz1715/go-utils/goutils"

	"github.com/zdz1715/ghttp"
)

func TestMessagesService_AsyncSendCorpConversation(t *testing.T) {
	client, err := NewClient(testInternalAppCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(ghttp.DefaultDebug),
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.Message.AsyncSendCorpConversation(context.Background(), &AsyncSendCorpConversationOptions{
		AgentId:    goutils.Ptr("2824804683"),
		UseridList: goutils.Ptr("125024135535014255"),
		Msg: &MsgOptions{
			Msgtype: goutils.Ptr(FileMsgType),

			Text: &TextMsgOptions{
				Content: goutils.Ptr("111111"),
			},

			Image: &ImageMsgOptions{
				MediaId: goutils.Ptr("@lADOADmaWMzazQKA"),
			},

			Voice: &VoiceMsgOptions{
				MediaId:  goutils.Ptr("@lADOADmaWMzazQKA"),
				Duration: goutils.Ptr("20"),
			},

			File: &FileMsgOptions{
				MediaId: goutils.Ptr("@lADOADmaWMzazQKA"),
			},

			Link: &LinkMsgOptions{
				MessageUrl: goutils.Ptr("https://www.baidu.com"),
				Title:      goutils.Ptr("百度一下，你就知道"),
				Text:       goutils.Ptr("点击"),
				PicUrl:     goutils.Ptr("https://img1.baidu.com/it/u=3511929181,2776771648&fm=253&fmt=auto&app=138&f=JPEG?w=584&h=500"),
			},

			OA: &OAMsgOptions{
				MessageUrl:   goutils.Ptr("https://www.baidu.com"),
				PcMessageUrl: goutils.Ptr("https://www.baidu.com"),
				Head: &OAMsgOptionsHead{
					Text:    goutils.Ptr("头部标题"),
					Bgcolor: goutils.Ptr("FFBBBBBB"),
					StatusBar: &OAMsgOptionsHeadStatusBar{
						StatusValue: goutils.Ptr("StatusValue"),
						StatusBg:    goutils.Ptr("FFBBBBBB"),
					},
				},
				Body: &OAMsgOptionsBody{
					Title: goutils.Ptr("正文标题"),
					Form: []*OAMsgOptionsBodyForm{
						{
							Key:   goutils.Ptr("KEY1"),
							Value: goutils.Ptr("VALUE1"),
						},
						{
							Key:   goutils.Ptr("KEY2"),
							Value: goutils.Ptr("VALUE2"),
						},
					},
					Rich: &OAMsgOptionsBodyRich{
						Num:  goutils.Ptr("110.120"),
						Unit: goutils.Ptr("$"),
					},
					Content:   goutils.Ptr("大段文本大段文本大段文本大段文本大段文本大段文本"),
					Image:     goutils.Ptr("@lADOADmaWMzazQKA"),
					FileCount: goutils.Ptr("3"),
					Author:    goutils.Ptr("李四"),
				},
			},

			Markdown: &MarkdownMsgOptions{
				Title: goutils.Ptr("it is markdown"),
				Text: goutils.Ptr(`
#### 杭州天气 @150XXXXXXXX
> 9度，西北风1级，空气良89，相对温度73%
> ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)
> ###### 10点20分发布 [天气](https://www.dingtalk.com)`),
			},

			ActionCard: &ActionCardMsgOptions{
				Title: goutils.Ptr("乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身"),
				Markdown: goutils.Ptr(`![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)
### 乔布斯 20 年前想打造的苹果咖啡厅
Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划`),
				BtnOrientation: goutils.Ptr("0"),
				//SingleTitle:    goutils.Ptr("阅读全文"),
				//SingleURL:      goutils.Ptr("https://www.dingtalk.com/"),
				BtnJsonList: []*ActionCardMsgOptionsBtn{
					{
						Title:     goutils.Ptr("内容不错"),
						ActionURL: goutils.Ptr("https://www.dingtalk.com/"),
					},
					{
						Title:     goutils.Ptr("不感兴趣"),
						ActionURL: goutils.Ptr("https://www.dingtalk.com/"),
					},
				},
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%+v", reply)
	}

}
