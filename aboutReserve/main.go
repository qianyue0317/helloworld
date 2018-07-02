package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"math/rand"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"github.com/lxn/walk/declarative"
	"github.com/lxn/walk"
	"strconv"
	"strings"
)

var nameList []string

func openMySql() (*sql.DB, error) {
	fmt.Println("数据库host:", DB_HOST)
	db, e := sql.Open("mysql", "root:dongyan@tcp("+DB_HOST+":3306)/rainbow?charset=utf8")
	if e != nil {
		panic(e)
		fmt.Println("打开数据库失败")
		return nil, e
	}
	fmt.Println("连接到数据库")
	return db, nil
}

// 生成随机生日的方法
func generateBirthday() string {
	startDate := time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)
	endDate := time.Date(2018, 6, 5, 0, 0, 0, 0, time.Local)
	startMill := startDate.Unix()
	endMill := endDate.Unix()
	randMill := rand.Int63n(endMill-startMill) + startMill
	randDate := time.Unix(randMill, 0)
	return randDate.Format("2006-01-02")
}

func generateName(count int) []string {
	results := []string{}
	if count > 3 {
		fmt.Println("最多创建三个宝宝")
		return results
	}
	nameListLength := len(nameList)
	for i := 0; i < count; i++ {
		results = append(results, nameList[rand.Intn(nameListLength)%nameListLength])
	}
	return results
}

func getUserIdWithOpenId(db *sql.DB) string {
	rows, e := db.Query("select id from user where user.wx_openid=?", wxOpenId)
	if e != nil {
		fmt.Println("查询用户出错", e)
		return ""
	}
	for rows.Next() {
		var id string
		rows.Scan(&id)
		return id
	}
	return ""
}

func addBaby(name string) (interface{}, bool) {
	client := &http.Client{}
	body := map[string]interface{}{}
	//body:=make(url.Values)
	body["baby_sex"] = 1
	body["baby_birthday"] = generateBirthday()
	body["hospital_id"] = 10001
	body["baby_name"] = name
	body["user_id"] = 2043
	str, jsonError := json.Marshal(body)
	if jsonError != nil {
		return nil, false
	}
	strTemp := string(str)
	fmt.Println("请求Body:", strTemp)

	//v:=url.Values{}
	//v.Add("baby_sex", "1")
	//v.Add("baby_birthday", generateBirthday())
	//v.Add("hospital_id", "10001")
	//v.Add("baby_name", name)
	//v.Add("user_id", "2043")
	//u:=ioutil.NopCloser(strings.NewReader(v.Encode()))

	request, e := http.NewRequest("POST", /*"http://localhost:5000/testPost"*/ ADD_URL, bytes.NewReader(str))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	response, err := client.Do(request)
	if e != nil {
		return nil, false
	}

	//response, err := http.Post(ADD_URL, "application/x-www-form-urlencoded", u)

	defer response.Body.Close()
	if err != nil {
		return nil, false
	}
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, false
	}
	fmt.Println(string(resBody))
	return 1, true
}


/*创建窗口*/
func createWindow() {

	var countP *walk.TextEdit

	count := declarative.TextEdit{
		AssignTo:&countP,
	}
	window := declarative.MainWindow{
		Title:  "testWindow",
		Size:   declarative.Size{Width: 600, Height: 400},
		Layout: declarative.VBox{
			Margins:declarative.Margins{Left: 50, Top: 50, Right: 50, Bottom: 50},
		},
		Children: []declarative.Widget{
			declarative.HSplitter{
				Children: []declarative.Widget{
					count,
					declarative.TextEdit{},
				},
			},
			declarative.PushButton{
				Text: "Scream",
				OnClicked: func() {
					fmt.Println("点击到了")
					i, e := strconv.Atoi(countP.Text())
					if e != nil {
						fmt.Println("请输入数字")
					}
					fmt.Println("内容",i)
				},
			},
			declarative.RadioButtonGroup{
				Optional:false,
				Buttons:[]declarative.RadioButton{
					declarative.RadioButton{
						Text:"测试",
					},
					declarative.RadioButton{
						Text:"开发",
					},
				},
			},
		},
		OnSizeChanged: func() {
			countP.SetHeight(100)

		},
	}
	window.Run()
}


func runMainWindow() {
	var inTE, outTE *walk.TextEdit

	declarative.MainWindow{
		Title:   "SCREAMO",
		MinSize: declarative.Size{600, 400},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.HSplitter{
				Children: []declarative.Widget{
					declarative.TextEdit{AssignTo: &inTE},
					declarative.TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			declarative.PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}


func runLoginMainWindow() {
	var usernameTE, passwordTE *walk.LineEdit
	declarative.MainWindow{
		Title:   "登录",
		MinSize: declarative.Size{270, 290},
		Layout:  declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Composite{
				Layout: declarative.Grid{Columns: 2, Spacing: 10},
				Children: []declarative.Widget{
					declarative.VSplitter{
						Children: []declarative.Widget{
							declarative.Label{
								Text: "用户名",
							},
						},
					},
					declarative.VSplitter{
						Children: []declarative.Widget{
							declarative.LineEdit{
								MinSize:  declarative.Size{160, 0},
								AssignTo: &usernameTE,
							},
						},
					},
					declarative.VSplitter{
						Children: []declarative.Widget{
							declarative.Label{MaxSize: declarative.Size{160, 40},
								Text: "密码",
							},
						},
					},
					declarative.VSplitter{
						Children: []declarative.Widget{
							declarative.LineEdit{
								MinSize:  declarative.Size{160, 0},
								AssignTo: &passwordTE,
							},
						},
					},
				},
			},

			declarative.PushButton{
				Text:    "登录",
				MinSize: declarative.Size{120, 50},
				OnClicked: func() {
					if usernameTE.Text() == "" {
						var tmp walk.Form
						walk.MsgBox(tmp, "用户名为空", "", walk.MsgBoxIconInformation)
						return
					}
					if passwordTE.Text() == "" {
						var tmp walk.Form
						walk.MsgBox(tmp, "密码为空", "", walk.MsgBoxIconInformation)
						return
					}
				},
			},
		},
	}.Run()
}

func main() {
	//createWindow()
	//runMainWindow()
	runLoginMainWindow()
}
