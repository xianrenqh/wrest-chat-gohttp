package deliver

import (
	"errors"
	"fmt"
	"github.com/opentdp/wrest-chat/dbase/point"
	"github.com/opentdp/wrest-chat/dbase/pointlist"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/opentdp/go-helper/logman"
)

func Send(deliver, content string) error {

	content = strings.TrimSpace(content)
	delivers := strings.Split(deliver, "\n")

	for _, dr := range delivers {
		logman.Warn("deliver "+dr, "content", content)
		// 解析参数
		args := strings.Split(strings.TrimSpace(dr), ",")
		if len(args) < 2 {
			return errors.New("deliver is error")
		}
		// 分渠道投递
		switch args[0] {
		case "wechat":
			time.Sleep(1 * time.Second)
			wechatMessage(args[1:], content)
		}
	}

	return nil

}

// 操作积分
func UpdateOrCreatePoints(Wxid string, RoomId string, Type int32, sendPoint int32, Sign int32, Desc string) int32 {
	//写入积分明细表
	pointlist.Create(&pointlist.CreateParam{
		Wxid:   Wxid,
		Roomid: RoomId,
		Type:   Type,
		Point:  sendPoint,
		Sign:   Sign,
		Desc:   Desc,
	})

	//写入总积分表
	upPoint, _ := point.Fetch(&point.FetchParam{Wxid: Wxid, Roomid: RoomId})

	var newPoint int32

	if upPoint.Rd > 0 {
		if Sign == 1 {
			newPoint = upPoint.Point + sendPoint
		} else {
			newPoint = upPoint.Point - sendPoint
			if newPoint < 1 {
				//删除
				err := point.Delete(&point.DeleteParam{Rd: upPoint.Rd})
				if err != nil {
					return 0
				}
				return 0
			}
		}
		ret := point.Update(&point.UpdateParam{
			Rd:     upPoint.Rd,
			Wxid:   Wxid,
			Roomid: RoomId,
			Point:  newPoint,
		})
		fmt.Println(ret)
	} else {
		newPoint = sendPoint
		point.Create(&point.CreateParam{
			Wxid:   Wxid,
			Roomid: RoomId,
			Point:  newPoint,
		})
	}
	return newPoint
}

func GetFiles(picUrlList string, Type int32) (string, error) {
	resp, err := http.Get(picUrlList)
	if err != nil {
		return fmt.Sprintf("请求失败：%s", err), nil
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Sprintf("读取响应失败：%s", err), nil
		}

		// 将图片内容保存到文件
		var saveDir = "public/FileCache/"
		savePath := saveDir + strconv.FormatInt(time.Now().UnixNano()/1e6, 10) + ".jpg"
		if Type == 2 {
			savePath = saveDir + strconv.FormatInt(time.Now().UnixNano()/1e6, 10) + ".mp4"
		}

		if _, err := os.Stat(saveDir); os.IsNotExist(err) {
			err = os.MkdirAll(saveDir, 0755)
			if err != nil {
				return fmt.Sprintf("创建文件夹失败：%s", err), nil
			}
		}

		err = ioutil.WriteFile(savePath, bodyBytes, 0755)
		if err != nil {
			return fmt.Sprintf("保存图片失败：%s", err), nil
		}
		return savePath, nil
	} else {
		return fmt.Sprintf("请求失败，状态码：%s", resp.StatusCode), nil
	}
}
