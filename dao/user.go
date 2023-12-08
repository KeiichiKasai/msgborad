package dao

import (
	"bufio"
	"fmt"
	"main.go/model"
	"main.go/utils"
	"os"
	"strings"
	"sync"
)

const UserFilePath = "user.txt"

var Redis = map[string]string{
	"Lanshan": "袁神启动！",
}

func InitDatabase() {
	f, err := os.OpenFile(UserFilePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("打开文件错误")
		panic(err)
	} else {
		defer f.Close()
	}
	r := bufio.NewReader(f)
	//写入Redis
	for {
		temp, _, e := r.ReadLine()
		if e != nil {
			break
		}
		data := strings.Split(string(temp), " ")
		Redis[data[0]] = data[1]
	}
}

// Refresh 写入MySQL
func Refresh(username, password, phone string) {
	f, err := os.OpenFile(UserFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("写入数据错误")
		panic(err)
	} else {
		defer f.Close()
	}
	w := bufio.NewWriter(f)

	_, err = w.WriteString(username + " " + password + " " + phone + "\n")

	if err != nil {
		fmt.Println("写入数据错误")
		panic(err)
		return
	}
	w.Flush()

}

func GetPasswordByUsername(username string) (string, error) {
	password, ok := Redis[username]
	if !ok {
		return "", utils.NotExistUser
	}
	return password, nil
}

func InsertUser(user model.User) {
	var m sync.Mutex
	m.Lock()
	Redis[user.Username] = user.Password
	m.Unlock()
	Refresh(user.Username, user.Password, user.Phone)
}

func SearchUser(username string) (string, error) {
	f, err := os.Open(UserFilePath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		temp := strings.Split(text, " ")
		if temp[0] == username {
			return text, nil
		}
	}
	if scanner.Err() != nil {
		return "", err
	}
	return "", err
}

func UpdateUser(user model.User) {
	f, err := os.OpenFile(UserFilePath, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	pos := int64(0)
	r := bufio.NewReader(f)
	for {
		l, e := r.ReadString('\n')
		fmt.Println(l)
		if e != nil {
			//读到末尾
			fmt.Println("Read file error!", e)
			break

		}
		temp := strings.Split(l, " ")

		if temp[0] == user.Username {
			data := []byte(user.Username + " " + user.Password + " " + user.Phone + "\n")
			_, err2 := f.WriteAt(data, pos)
			if err2 != nil {
				fmt.Println(err2)
			}
			break
		}
		pos += int64(len(temp))
	}
	Redis[user.Username] = user.Password
}
