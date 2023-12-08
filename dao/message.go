package dao

import (
	"bufio"
	"fmt"
	"main.go/model"
	"os"
)

const MessageFilePath = "message.txt"

func InsertMessage(m model.Message) error {
	f, err := os.OpenFile(MessageFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("openFile failed")
		return err
	}
	data := m.Username + " " + m.WriteTime + " " + m.Context + "\n"
	w := bufio.NewWriter(f)
	_, err = w.WriteString(data)
	if err != nil {
		fmt.Println("write in failed")
		return err
	}
	w.Flush()
	return nil
}
func SearchMessage() ([]string, error) {
	data := make([]string, 0)
	f, err := os.Open(MessageFilePath)
	if err != nil {
		fmt.Println("open file failed")
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}
	return data, nil

}
