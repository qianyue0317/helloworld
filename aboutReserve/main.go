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
	body["baby_sex"] = 1
	body["baby_birthday"] = generateBirthday()
	body["default_hospitail_id"] = 10001
	body["baby_name"] = name
	body["user_id"] = 2043
	str, jsonError := json.Marshal(body)
	if jsonError != nil {
		return nil, false
	}
	strTemp := string(str)
	fmt.Println("请求Body:", strTemp)
	request, e := http.NewRequest("POST", /*"http://localhost:5000/testPost"*/ ADD_URL, bytes.NewReader(str))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if e != nil {
		return nil, false
	}
	response, err := client.Do(request)
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

func main() {
	// 初始化姓名列表
	nameList = initNameList()
	config()
	//fmt.Println(generateBirthday())
	db, e := openMySql()
	if e != nil {
		return
	}
	fmt.Println(getUserIdWithOpenId(db))
	addBaby("qian")
	//var env string
	//fmt.Print("请输入环境(1:测试,2开发):")
	//fmt.Scanf("%s", &env)
	//fmt.Println("您输入的内容为:", env)
}
