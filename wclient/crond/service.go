package crond

import (
	"errors"
	"github.com/opentdp/go-helper/command"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"

	"github.com/opentdp/wrest-chat/dbase/cronjob"
	"github.com/opentdp/wrest-chat/dbase/tables"
	"github.com/opentdp/wrest-chat/wclient"
	"github.com/opentdp/wrest-chat/wclient/aichat"
	"github.com/opentdp/wrest-chat/wclient/deliver"
)

var crontab *cron.Cron

func Daemon() {

	log.Info().Msg("正在启用定时任务，请稍后...")

	crontab = cron.New(cron.WithSeconds())

	jobs, err := cronjob.FetchAll(&cronjob.FetchAllParam{})
	if err != nil || len(jobs) == 0 {
		return
	}

	for _, job := range jobs {
		AttachJob(job)
	}

	crontab.Start()
	log.Info().Msg("定时任务启用成功，请稍后...")
}

// 触发计划任务

func Execute(id uint) error {

	job, _ := cronjob.Fetch(&cronjob.FetchParam{Rd: id})

	if job.Content == "" {
		return errors.New("content is empty")
	}
	if job.Deliver == "-" {
		return errors.New("deliver is empty")
	}

	log.Info().Msg("cron:run " + job.Name)
	// 如果需要记录额外的信息，可以使用如下方式
	log.Info().Str("message", "cron:run "+job.Name).Str("entryId", strconv.FormatInt(job.EntryId, 10)).Send()

	// 发送文本内容
	if job.Type == "TEXT" {
		return deliver.Send(job.Deliver, job.Content)
	}

	// 发送AI生成的文本
	if job.Type == "AI" {
		wc := wclient.Register()
		if wc == nil {
			log.Error().Msg("cron:ai出错。wclient is nil")
		}
		self := wc.CmdClient.GetSelfInfo()
		data := aichat.Text(job.Content, self.Wxid, "")
		return deliver.Send(job.Deliver, data)
	}

	// 执行命令获取结果
	output, err := command.Exec(&command.ExecPayload{
		Name:          "cron: " + job.Name,
		CommandType:   job.Type,
		WorkDirectory: job.Directory,
		Content:       job.Content,
		Timeout:       job.Timeout,
	})
	if err != nil {
		log.Warn().Msg("任务运行错误：" + job.Name)
		return err
	}

	// 发送命令执行结果
	log.Warn().Msg("cron:run " + job.Name)
	if output != "" {
		return deliver.Send(job.Deliver, output)
	}

	return nil

}

// 激活计划任务

func AttachJob(job *tables.Cronjob) error {

	cmd := func(id uint) func() {
		return func() { Execute(id) }
	}

	sepc := []string{
		job.Second, job.Minute, job.Hour, job.DayOfMonth, job.Month, job.DayOfWeek,
	}

	entryId, err := crontab.AddFunc(strings.Join(sepc, " "), cmd(job.Rd))
	if err != nil {
		return err
	}

	log.Info().Msg("任务名称： " + job.Name + " " + string(entryId))
	err = cronjob.Update(&cronjob.UpdateParam{
		Rd:      job.Rd,
		EntryId: int64(entryId),
	})

	return err

}

// 管理生命周期

func NewById(rd uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Rd: rd})

	if err == nil && job.Rd > 0 {
		AttachJob(job)
	}

}

func UndoById(rd uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Rd: rd})

	if err == nil && job.Rd > 0 {
		log.Info().Msg("cron:remove " + job.Name + string(job.EntryId))
		crontab.Remove(cron.EntryID(job.EntryId))
	}

}

func RedoById(rd uint) {

	job, err := cronjob.Fetch(&cronjob.FetchParam{Rd: rd})

	if err == nil && job.Rd > 0 {
		log.Info().Msg("cron:update " + job.Name + string(job.EntryId))
		crontab.Remove(cron.EntryID(job.EntryId))
		AttachJob(job)
	}

}

// 获取执行状态

type JobStatus struct {
	EntryId  int64 `json:"entry_id"`
	NextTime int64 `json:"next_time"`
	PrevTime int64 `json:"prev_time"`
}

func GetEntries() map[uint]JobStatus {

	list := map[uint]JobStatus{}

	jobs, err := cronjob.FetchAll(&cronjob.FetchAllParam{})
	if err != nil || len(jobs) == 0 {
		return list
	}

	for _, job := range jobs {
		entry := crontab.Entry(cron.EntryID(job.EntryId))
		list[job.Rd] = JobStatus{
			EntryId:  int64(entry.ID),
			NextTime: entry.Next.Unix(),
			PrevTime: entry.Prev.Unix(),
		}
	}

	return list

}
