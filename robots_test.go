package dingtalk

import (
	"context"
	"testing"

	"github.com/zdz1715/go-utils/goutils"

	"github.com/zdz1715/ghttp"
)

func TestRobotsService_ListRobotsInGroup(t *testing.T) {
	client, err := NewClient(testInternalAppCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(ghttp.DefaultDebug),
		},
	})

	if err != nil {
		t.Fatal(err)
	}
	// todo: 404!!! ??????
	// curl -I -X POST https://api.dingtalk.com/v1.0/robot/getBotListInGroup
	reply, err := client.Robot.ListRobotsInGroup(context.Background(), &ListRobotsInGroupOptions{
		OpenConversationId: goutils.Ptr("1"),
	})

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%v", reply)
	}
}

func TestRobotsService_Send(t *testing.T) {
	client, err := NewClient(testInternalAppCredential, &Options{
		ClientOpts: []ghttp.ClientOption{
			ghttp.WithDebug(ghttp.DefaultDebug),
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	err = client.Robot.Send(context.Background(),
		"9ec4930cc8cc4ec543e2a771d58680dd14811b58b1e11674a51880784afd039f",
		&RobotSendOptions{
			Msgtype: goutils.Ptr(MarkdownRobotMsgType),

			// text
			Text: &RobotTextMsgOptions{
				Content: goutils.Ptr("it is text msg"),
			},

			// link
			//Msgtype: goutils.Ptr(LinkMsgType),
			Link: &RobotLinkMsgOptions{
				MessageUrl: goutils.Ptr("https://www.baidu.com"),
				Title:      goutils.Ptr("百度一下，你就知道"),
				Text:       goutils.Ptr("点击"),
				PicUrl:     goutils.Ptr("https://img1.baidu.com/it/u=3511929181,2776771648&fm=253&fmt=auto&app=138&f=JPEG?w=584&h=500"),
			},

			//actionCard
			ActionCard: &RobotActionCardMsgOptions{
				Title: goutils.Ptr("乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身"),
				Text: goutils.Ptr(`![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)
### 乔布斯 20 年前想打造的苹果咖啡厅
Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划`),
				BtnOrientation: goutils.Ptr("0"),
				//SingleTitle:    goutils.Ptr("阅读全文"),
				//SingleURL:      goutils.Ptr("https://www.dingtalk.com/"),
				Btns: []*RobotActionCardMsgOptionsBtn{
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

			// feedCard
			FeedCard: &RobotFeedCardMsgOptions{
				Links: []*RobotFeedCardMsgOptionsLink{
					{
						MessageURL: goutils.Ptr("https://www.baidu.com"),
						Title:      goutils.Ptr("百度一下，你就知道-111"),
						PicURL:     goutils.Ptr("https://img1.baidu.com/it/u=3511929181,2776771648&fm=253&fmt=auto&app=138&f=JPEG?w=584&h=500"),
					},
					{
						MessageURL: goutils.Ptr("https://www.baidu.com"),
						Title:      goutils.Ptr("百度一下，你就知道-222"),
						PicURL:     goutils.Ptr("https://img1.baidu.com/it/u=3511929181,2776771648&fm=253&fmt=auto&app=138&f=JPEG?w=584&h=500"),
					},
				},
			},

			// markdown
			Markdown: &RobotMarkdownMsgOptions{
				Title: goutils.Ptr("it is markdown"),
				Text: goutils.Ptr(`
#### 杭州天气 @150XXXXXXXX
> 9度，西北风1级，空气良89，相对温度73%
> ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)
> ###### 10点20分发布 [天气](https://www.dingtalk.com)`),
			},

			At: &RobotSendOptionsAt{
				IsAtAll: goutils.Ptr(false),
				//AtMobiles: []string{""},
			},
		})

	if err != nil {
		t.Error(err)
	}

}
