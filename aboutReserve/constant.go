package main

import "fmt"

var DevEnv = 0
var TestEnv = 1

var ENV = TestEnv
var wxOpenId = "ooycV0VlhBXT4wiJHRV4lHHYeUXA"

//测试环境
var TEST_ADD_URL = "https://rainbow5.arxanfintech.com/rainbow/v1/baby_add"
var TEST_RESERVE_URL = "https://rainbow5.arxanfintech.com/rainbow/v1/reserve"
var TEST_DELETE_URL = "https://rainbow5.arxanfintech.com/rainbow/v1/baby_delete"
var TEST_DB_HOST = "124.42.118.162"
//开发环境
var DEV_ADD_URL = "https://rainbow3.arxanfintech.com/rainbow/v1/baby_add"
var DEV_RESERVE_URL = "https://rainbow3.arxanfintech.com/rainbow/v1/reserve"
var DEV_DELETE_URL = "https://rainbow3.arxanfintech.com/rainbow/v1/baby_delete"
var DEV_DB_HOST = "124.42.118.198"

var ADD_URL = ""
var RESERVE_URL = ""
var DELETE_URL = ""
var DB_HOST = ""

func config() {
	fmt.Println("开始配置")
	if ENV == TestEnv {
		ADD_URL = TEST_ADD_URL
		RESERVE_URL = TEST_RESERVE_URL
		DELETE_URL = TEST_DELETE_URL
		DB_HOST = TEST_DB_HOST
	} else if ENV == DevEnv {
		ADD_URL = DEV_ADD_URL
		RESERVE_URL = DEV_RESERVE_URL
		DELETE_URL = DEV_DELETE_URL
		DB_HOST = DEV_DB_HOST
	}
}
