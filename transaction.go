package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"log"
	"net/http"
)

// 模拟数据存储
type user struct {
	id    string
	money int
}

var userdata []*user

func init() {
	userdata = append(userdata, &user{id: "Tom", money: 10000})
	userdata = append(userdata, &user{id: "John", money: 10000})
	userdata = append(userdata, &user{id: "hack", money: 0})
}

// 模拟交易系统API
func main() {
	http.HandleFunc("/transaction", transactionHandler)
	log.Println("Transaction Server Running....")
	http.ListenAndServe(":1999", nil)
}

// 假设这是交易处理器
func transactionHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/transaction" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch req.Method {
	case "GET":
		http.ServeFile(w, req, "form.html")
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		if e := transaction(req.FormValue("Id"), req.FormValue("toId"), req.FormValue("money")); e != nil {
			fmt.Fprintf(w,"%s",e.Error())
			return
		}
		fmt.Fprintf(w,"transaction successful.")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func transaction(Id, toId, money string) error {

	if len(Id) <= 0 || len(toId) <= 0 || cast.ToInt(money) <= 0 {
		return errors.New("form invalid parameter")
	}
	var user, recipient *user
	for _, u := range userdata {
		if u.id == Id {
			user = u
			break
		}
	}
	for _, r := range userdata {
		if r.id == toId {
			recipient = r
			break
		}
	}
	// 用户不存在
	if user == nil || recipient == nil {
		return errors.New("user does not exist")
	}
	// 有钱才能转账
	if user.money < 0 || user.money < cast.ToInt(money) {
		return errors.New("your balance is insufficient")
	}
	user.money = user.money - cast.ToInt(money)
	recipient.money = recipient.money + cast.ToInt(money)
	log.Printf("用户 %s 向用户 %s 转账 %s 元.", Id, toId, money)
	return nil
}
