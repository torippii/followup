# How To Use

## 前提
・Gitがインストール済み
・Chromeブラウザがインストール済み
・Go(v1.18以上推奨)がインストール済み

## 1.ソースコードを取得する

### 1-1.下記コマンドで取得する。
git clone https://github.com/torippii/followup.git

## 2.接続情報を設定する

### 2-1.followup/controller/page.go にて接続先のURLを設定する。
例）TARGET_URL = "https://hoge.fuga.co/fugahoge/login/"

### 2-2.ログインIDとPWを設定する
followup/setting.json に自身のログインIDとPWを設定する。

## 3.Webドライバーのパスを通す
利用環境により異なります。以下Windowsの例となるため適宜読み替えてください。

### 3-1.Webドライバーをダウンロード
ChromeブラウザとWebドライバーのメジャーバージョンは合わせてください。
ブラウザのバージョンは 「︙」＞「ヘルプ」＞「GoogleChromeについて」から確認可能です。

下記サイトからダウンロードします。
https://sites.google.com/chromium.org/driver/

### 3-2.環境変数の設定
zipをダウンロード後、解凍をおこない任意のフォルダへ配置します。
配置先のパスを環境変数のPathに追加してください。

Pathに追加後、コマンドプロンプトを開き chromedriver を実行し「ChromeDriver was started successcully」が末尾に出力されていれば設定完了です。
Ctrl+cでChromeDriverを終了します。

## 4.実行ファイルの作成

### 4-1.ビルド場所へ移動
コマンドプロンプトにてプロジェクトルート(followup のフォルダ) へ移動します。

### 4-2.ビルド
go build を実行します。

## 5.実行

### 5-1.勤怠データの用意
attendance.csv に勤怠データを入力します。
左から順に[日にち(yy),出勤(hhmm),退勤(hhmm),休憩開始(hhmm),休憩終了(hhmm)]です。
先頭に#が着くとコメントアウト行とみなされます。

### 5-2.実行
実行ファイルをダブルクリックで実行可能です。
登録までとなり、申請までは行いません。
勤務開始と勤務終了の時間だけ登録されます。休憩時間は登録されません。(そこは作ってませんが勤務開始などと同じ要領で作れるのでご自身で修正して動かしてみるのもよいと思います！)


