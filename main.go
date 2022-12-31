package main

import (
	"fmt"
	"log"

	"hrmos-bulk-register/controller"
	"hrmos-bulk-register/util"

	"github.com/sclevine/agouti"
)

func main(){
	fmt.Println("HRMOS Bulk Register")

	//ログイン情報の読み込み
	u := util.LoadSettingJson()

	// Webドライバの取得
	var d *agouti.WebDriver = controller.GetDriver()
	if err := d.Start(); err != nil {
		log.Fatalln(err)
	}
	defer d.Stop()

	// ページのコントローラ生成
	c := controller.NewController(d)

	// ログイン
	c.Login(u)

	// 勤怠登録
	c.BulkRegister()


	// 終了時にブラウザも閉じるため終了して良いかを確認
	var input string
	for {
		fmt.Print("処理が完了しました。終了します。 (y/n) > ")
		fmt.Scan(&input)
		if input == "y" || input == "Y" {
			break
		}
	}
}