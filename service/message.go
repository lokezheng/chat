package service

import "chat/model"

type MessageService struct {
}

func (s *MessageService) AddMessage(msg model.Message) (err error) {
	_, err = DbEngin.InsertOne(&msg)
	return err
}

func (s *MessageService) FindMessages(id int64, cmd int) []model.Message {
	tmp := make([]model.Message, 0)
	DbEngin.Where("dstid = ? and cmd = ? and createat >= now()-interval 10 minute", id, cmd).Find(&tmp)
	return tmp
}

//获取最近的50条消息
func (s *MessageService) FindMessagesLimit(id int64, cmd int) []model.Message {
	tmp := make([]model.Message, 0)
	DbEngin.Where("dstid = ? and cmd = ? ", id, cmd).Limit(50).OrderBy("id desc").Find(&tmp)
	return tmp
}
