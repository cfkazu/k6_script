# k6_script
go上からk6を簡単に使いたいなと思ったので作成。割と見切り発車だが……

# usage
```
go run . (テストしたいホスト名が書かれたtxtファイル名) (テストの種類)
```
とすることで、txtファイルに書かれたホスト名に順次テストを実施していく。
テストの種類はいま
```
spike:spikeテストを実行
stress:stressテストを実行
```
しかなく、またそれらの詳細はそれぞれ対応する.jsファイルで確認する必要がある。
それぞれのパラメータをもっと簡単にいじれるようにしたい。

#txtファイルについて
```
test.k6.io,default
test.k6.io/news.php,news
test.k6.io/contacts.php,contacts
```
のように記す。ホスト名,タグ名の順。
