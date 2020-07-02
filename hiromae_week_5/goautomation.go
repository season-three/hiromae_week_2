package main

import (
	"fmt"
	"time"

	"github.com/sclevine/agouti"
	"github.com/season-three/hiromae_week_5/secret"
)

func main() {
	fmt.Println(secret.Myusername())
	fmt.Println(secret.Mypassword())

	url1 := "https://www.linkedin.com/search/results/people/?facetGeoRegion=%5B%22jp%3A0%22%5D&keywords=Golang%20AND%20Docker%20AND%20GCP&origin=GLOBAL_SEARCH_HEADER"
	url2 := ""
	url3 := ""

	driver := agouti.ChromeDriver()

	//ドライバの起動＋スタート
	err := driver.Start()
	if err != nil {
		fmt.Println("Failed to start driver")
	}
	//defer driver.Stop()

	//クロームの起動＋ページを開く
	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		fmt.Println("Failed to open new page")
	}

	//ログインページ開く
	err = page.Navigate(url1)
	if err != nil {
		fmt.Println("Failed to open login page")
	}

	//もしログイン画面が出てきたら
	page.FindByClass("main__sign-in-link").Click()

	//usernameとpasswordのIDを取得
	username := page.FindByID("username")
	password := page.FindByID("password")

	//usernameとpasswordを入力
	username.Click()
	username.Fill(secret.Myusername())
	time.Sleep(3 * time.Second)
	username.Click()
	password.Fill(secret.Mypassword())
	time.Sleep(3 * time.Second)

	//サインインボタンを押す
	err = page.FindByButton("サインイン").Submit()
	if err != nil {
		fmt.Println("Failed to login")
	}
	time.Sleep(8 * time.Second)

	//つながり申請する
	count := 0

	for i := 0; i < 3; i++ {
		err = page.FindByButton("つながりを申請").Click()
		page.FindByButton("完了").Click()
		count++
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		s.ScrollFinger(0, 200)
		time.Sleep(2 * time.Second)
	} else {
		page.FindByButton("次へ").Submit()
		time.Sleep(5 * time.Second)
	}

	/*

		//別の単語で検索 or 新しいURLで検索

		//プロフィール見てくれた人につながり申請(スクロールの方法が分かったら)
		// URL := https://www.linkedin.com/me/profile-views/urn:li:wvmp:summary/

		//いいね1日5つの投稿に押す？
		url10 := "https://www.linkedin.com/feed/"

		err = page.Navigate(url10)
		if err != nil {
			fmt.Println("Failed to open feed page")
		}
		for i := 0; i < 5; i++ {
			page.FindByClass("artdeco-button__text react-button__text").Click()
		}

		//繋がり申請があれば承認押す
		url11 := "https://www.linkedin.com/mynetwork/"

		err = page.Navigate(url11)
		if err != nil {
			fmt.Println("Failed to open accept page")
		}
		//申請されてる分だけ押すにはどうすれば良いか？
		for {
			err := page.FindByClass("artdeco-button__text").Click()
			if err != nil {
				break
			}
		}


	*/
}
