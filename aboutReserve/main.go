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
	"sort"
	"os"
	"log"

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
		AssignTo: &countP,
	}
	window := declarative.MainWindow{
		Title: "testWindow",
		Size:  declarative.Size{Width: 600, Height: 400},
		Layout: declarative.VBox{
			Margins: declarative.Margins{Left: 50, Top: 50, Right: 50, Bottom: 50},
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
					fmt.Println("内容", i)
				},
			},
			declarative.RadioButtonGroup{
				Optional: false,
				Buttons: []declarative.RadioButton{
					declarative.RadioButton{
						Text: "测试",
					},
					declarative.RadioButton{
						Text: "开发",
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

type Condom struct {
	Index   int
	Name    string
	Price   int
	checked bool
}

type CondomModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Condom
}

func (m *CondomModel) RowCount() int {
	return len(m.items)
}

func (m *CondomModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index
	case 1:
		return item.Name
	case 2:
		return item.Price
	}
	panic("unexpected col")
}

func (m *CondomModel) Checked(row int) bool {
	return m.items[row].checked
}

func (m *CondomModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked
	return nil
}

func (m *CondomModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *CondomModel) Len() int {
	return len(m.items)
}

func (m *CondomModel) Less(i, j int) bool {
	a, b := m.items[i], m.items[j]

	c := func(ls bool) bool {
		if m.sortOrder == walk.SortAscending {
			return ls
		}

		return !ls
	}

	switch m.sortColumn {
	case 0:
		return c(a.Index < b.Index)
	case 1:
		return c(a.Name < b.Name)
	case 2:
		return c(a.Price < b.Price)
	}

	panic("unreachable")
}

func (m *CondomModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

func NewCondomModel() *CondomModel {
	m := new(CondomModel)
	m.items = make([]*Condom, 3)

	m.items[0] = &Condom{
		Index: 0,
		Name:  "golang",
		Price: 20,
	}

	m.items[1] = &Condom{
		Index: 1,
		Name:  "python",
		Price: 18,
	}

	m.items[2] = &Condom{
		Index: 2,
		Name:  "JavaScript",
		Price: 19,
	}

	return m
}

type CondomMainWindow struct {
	*walk.MainWindow
	model *CondomModel
	tv    *walk.TableView
}

func createTableView() {
	mw := &CondomMainWindow{model: NewCondomModel()}

	declarative.MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Condom展示",
		Size:     declarative.Size{800, 600},
		Layout:   declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Composite{
				Layout: declarative.HBox{MarginsZero: true},
				Children: []declarative.Widget{
					declarative.HSpacer{},
					declarative.PushButton{
						Text: "Add",
						OnClicked: func() {
							mw.model.items = append(mw.model.items, &Condom{
								Index: mw.model.Len() + 1,
								Name:  "第六感",
								Price: mw.model.Len() * 5,
							})
							mw.model.PublishRowsReset()
							mw.tv.SetSelectedIndexes([]int{})
						},
					},
					declarative.PushButton{
						Text: "Delete",
						OnClicked: func() {
							var items []*Condom
							remove := mw.tv.SelectedIndexes()
							for i, x := range mw.model.items {
								removeOk := false
								for _, j := range remove {
									if i == j {
										removeOk = true
									}
								}
								if !removeOk {
									items = append(items, x)
								}
							}
							mw.model.items = items
							mw.model.PublishRowsReset()
							mw.tv.SetSelectedIndexes([]int{})
						},
					},
					declarative.PushButton{
						Text: "ExecChecked",
						OnClicked: func() {
							for _, x := range mw.model.items {
								if x.checked {
									fmt.Printf("checked: %v\n", x)
								}
							}
							fmt.Println()
						},
					},
					declarative.PushButton{
						Text: "AddPriceChecked",
						OnClicked: func() {
							for i, x := range mw.model.items {
								if x.checked {
									x.Price++
									mw.model.PublishRowChanged(i)
								}
							}
						},
					},
				},
			},
			declarative.Composite{
				Layout: declarative.VBox{},
				ContextMenuItems: []declarative.MenuItem{
					declarative.Action{
						Text:        "I&nfo",
						OnTriggered: mw.tv_ItemActivated,
					},
					declarative.Action{
						Text: "E&xit",
						OnTriggered: func() {
							mw.Close()
						},
					},
				},
				Children: []declarative.Widget{
					declarative.TableView{
						AssignTo:         &mw.tv,
						CheckBoxes:       true,
						ColumnsOrderable: true,
						MultiSelection:   true,
						Columns: []declarative.TableViewColumn{
							{Title: "编号"},
							{Title: "名称"},
							{Title: "价格"},
						},
						Model: mw.model,
						OnCurrentIndexChanged: func() {
							i := mw.tv.CurrentIndex()
							if 0 <= i {
								fmt.Printf("OnCurrentIndexChanged: %v\n", mw.model.items[i].Name)
							}
						},
						OnItemActivated: mw.tv_ItemActivated,
					},
				},
			},
		},
	}.Run()
}

func (mw *CondomMainWindow) tv_ItemActivated() {
	msg := ``
	for _, i := range mw.tv.SelectedIndexes() {
		msg = msg + "\n" + mw.model.items[i].Name
	}
	walk.MsgBox(mw, "title", msg, walk.MsgBoxIconInformation)
}


type MyMainWindow struct {
	*walk.MainWindow
	edit *walk.TextEdit

	path string
}

/**
	文件选择器
 */
func showFileSelect() {
	mw := &MyMainWindow{}
	MW := declarative.MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "文件选择对话框",
		MinSize:  declarative.Size{150, 200},
		Size:     declarative.Size{300, 400},
		Layout:   declarative.VBox{},
		Children: []declarative.Widget{
			declarative.TextEdit{
				AssignTo: &mw.edit,
			},
			declarative.PushButton{
				Text:      "打开",
				OnClicked: mw.pbClicked,
			},
		},
	}
	if _, err := MW.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func (mw *MyMainWindow) pbClicked() {

	dlg := walk.FileDialog{}
	dlg.FilePath = mw.path
	dlg.Title = "Select File"
	dlg.Filter = "Exe files (*.exe)|*.exe|All files (*.*)|*.*"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		mw.edit.AppendText("Error : File Open\r\n")
		return
	} else if !ok {
		mw.edit.AppendText("Cancel\r\n")
		return
	}
	mw.path = dlg.FilePath
	s := fmt.Sprintf("Select : %s\r\n", mw.path)
	mw.edit.AppendText(s)
}

func fileSearch() {
	mw := &FileSearchMainWindow{}

	if _, err := (declarative.MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "SearchBox",
		//Icon:     "test.ico",
		MinSize:  declarative.Size{300, 400},
		Layout:   declarative.VBox{},
		Children: []declarative.Widget{
			declarative.GroupBox{
				Layout: declarative.HBox{},
				Children: []declarative.Widget{
					declarative.LineEdit{
						AssignTo: &mw.searchBox,
					},
					declarative.PushButton{
						Text:      "检索",
						OnClicked: mw.clicked,
					},
				},
			},
			declarative.TextEdit{
				AssignTo: &mw.textArea,
			},
			declarative.ListBox{
				AssignTo: &mw.results,
				Row:      5,
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}

}

type FileSearchMainWindow struct {
	*walk.MainWindow
	searchBox *walk.LineEdit
	textArea  *walk.TextEdit
	results   *walk.ListBox
}

func (mw *FileSearchMainWindow) clicked() {
	word := mw.searchBox.Text()
	text := mw.textArea.Text()
	model := []string{}
	for _, i := range search(text, word) {
		model = append(model, fmt.Sprintf("%d检索成功", i))
	}
	log.Print(model)
	mw.results.SetModel(model)
}

func search(text, word string) (result []int) {
	result = []int{}
	i := 0
	for j, _ := range text {
		if strings.HasPrefix(text[j:], word) {
			log.Print(i)
			result = append(result, i)
		}
		i += 1
	}
	return
}

func main() {
	//createWindow()
	//runMainWindow()
	//runLoginMainWindow()
	//createTableView()
	//showFileSelect()
	fileSearch()
}
