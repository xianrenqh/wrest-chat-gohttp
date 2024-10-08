package deliver

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/opentdp/go-helper/dborm"
	"github.com/opentdp/wrest-chat/dbase/point"
	"github.com/opentdp/wrest-chat/dbase/pointlist"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Send(deliver, content string) error {

	content = strings.TrimSpace(content)
	delivers := strings.Split(deliver, "\n")

	for _, dr := range delivers {
		log.Warn().Str("content", content).Msg("deliver " + dr)
		// 解析参数
		args := strings.Split(strings.TrimSpace(dr), ",")
		if len(args) < 2 {
			return errors.New("接收器错误...")
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
func UpdateOrCreatePoints(Wxid string, RoomId string, Type int32, sendPoint int, Sign int32, Desc string) int {
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

	var newPoint int

	if upPoint.Rd > 0 {
		if Sign == 1 {
			newPoint = upPoint.Point + sendPoint
		} else {
			newPoint = upPoint.Point - sendPoint
			if newPoint < 1 {
				newPoint = 0
			}
		}
		sql := "UPDATE point SET Point = ? WHERE Rd = ?"
		dborm.Db.Exec(sql, newPoint, upPoint.Rd)
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

func JudgeEqualListWord(content string, picKeyWords []string) bool {
	words := strings.Split(content, " ") // 假设 content 中的单词由空格分隔
	for _, word := range words {
		found := false
		for _, keyWord := range picKeyWords {
			if word == keyWord {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// 对字符串进行两次 MD5 加密
func DoubleMD5(s string) string {
	// 第一次 MD5 加密
	hash1 := md5.Sum([]byte(s))
	firstMd5 := hex.EncodeToString(hash1[:])

	// 第二次 MD5 加密
	hash2 := md5.Sum([]byte(firstMd5))
	secondMd5 := hex.EncodeToString(hash2[:])

	return secondMd5
}
