package service

import (
	"fmt"
	"main.go/dao"
	"main.go/model"
	"strings"
	"time"
)

func LeaveMessage(username, context string) error {
	current := time.Now().String()
	cur := strings.Split(current, ".")[0]
	m := model.Message{
		Username:  username,
		WriteTime: cur,
		Context:   context,
	}
	err := dao.InsertMessage(m)
	if err != nil {
		return err
	}
	return nil
}

func GetMessage() ([]model.Message, error) {
	var msg []model.Message
	rawData, err := dao.SearchMessage()
	fmt.Println(rawData)
	fmt.Println("--------")
	fmt.Println(rawData[0])
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(rawData); i++ {
		data := strings.Split(rawData[i], " ")
		m := model.Message{
			Username:  data[0],
			WriteTime: data[1] + " " + data[2],
			Context:   data[3],
		}
		msg = append(msg, m)
	}
	return msg, nil
}
