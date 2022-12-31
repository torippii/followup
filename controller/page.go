package controller

import (
	"hrmos-bulk-register/model"
	"hrmos-bulk-register/util"
	"log"

	"github.com/sclevine/agouti"
)

const (
	TARGET_URL = "https://hoge.fuga/"
	USER_AGENT = "--user-agent='Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36'"
)

type Controller struct {
	Page *agouti.Page
}

// Webドライバの取得
func GetDriver() *agouti.WebDriver {
	options := agouti.ChromeOptions(
		"args", []string{
			"--headless",    // ヘッドレスモード ブラウザ表示させずに実行するため
			"--disable-gpu", // windowsで使用する場合に必要なオプション
			USER_AGENT,      // UAにheadlessの記載が出ないよう通常時のUAに書き換え
			"--disable-blink-features=AutomationControlled", // webドライバで動作しているとtrueになるのでfalseに書き換え
			"--window-size=1200,1000",
		})
	return agouti.ChromeDriver(agouti.Browser("chrome"), options)
}

// コントローラの生成
func NewController(driver *agouti.WebDriver) *Controller {
	p, err := driver.NewPage()
	if err != nil {
		log.Fatalln(err)
	}

	return &Controller{
		Page: p,
	}
}

// ログイン処理
func (c *Controller) Login(u model.User) {

	// ページに遷移
	if err := c.Page.Navigate(TARGET_URL); err != nil {
		log.Fatalln(err)
	}

	// ログイン
	c.Page.FindByID("user_login_id").Fill(u.ID)
	c.Page.FindByID("user_password").Fill(u.PW)
	c.Page.FindByName("Submit").Click()
}

// 勤怠入力処理
func (c *Controller) BulkRegister() {

	// ホーム画面から月の勤怠一覧画面へ遷移
	c.Page.FindByID("work_navi").Click()

	// CSVを読み込む
	attendanceList := util.LoadCSV()

	// 各日付の入力行を取得し、対象日付があれば入力画面へ遷移する
	dates := c.Page.FindByID("editGraphTable").AllByClass("cellDate")
	dcnt, _ := c.Page.FindByID("editGraphTable").AllByClass("cellDate").Count()
	for _, v := range attendanceList {

		// ターゲットの日付を取得
		for j := 0; j < dcnt; j++ {
			date, _ := dates.At(j).FindByClass("date").Text()

			if v.Date != date {
				continue
			}
			// ターゲット日付の勤務入力画面へ遷移
			dates.At(j).FindByClass("view_work").Find("a").Click()

			// 出勤開始時刻を入力
			c.Page.FindByID("work_start_at_str").Clear()
			c.Page.FindByID("work_start_at_str").SendKeys(v.WorkStart)

			c.Page.FindByID("work_end_at_str").Clear()
			c.Page.FindByID("work_end_at_str").SendKeys(v.WorkEnd)

			// 登録実行
			c.Page.FindByClass("tableHeader").FindByName("commit").Click()

			break
		}
	}
}
