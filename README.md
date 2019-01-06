■go
・go導入にあたりしたこと
1:goenvインストール（goのバージョン管理）
brew install goenv

2:goenv環境変数設定。.zshrcに以下を追記
export PATH="$HOME/.goenv/bin:$PATH"
eval "$(goenv init -)"

3:goインストール可能バージョン確認
goenv install -l ・・・これでバージョン一覧が表示される

4:goインストール
goenv install 1.11.2

5:現在使用できるバージョン確認
goenv versions

6:インストールしたバージョンを使えるように設定
goenv global 1.11.2

7:goバージョン確認
go version

8:Go環境変数設定。.zshrcに以下を追記
export GOPATH=$HOME/go/package:$HOME/go/workspace
export PATH=$HOME/go/package/bin:$HOME/go/workspace/bin:$PATH

上記のようにpackageとworkspaceを書くことで、自分のプロジェクトと外部からとってきたパッケージが1つのディレクトリに混在する事を防ぐ

■Echo（Goフレームワーク）
・バリデーション
Echoにはバリデーション機能が組み込まれてない。
go-playgroundというサードパーティのライブラリを組み込んだ。

＞取り込み手順
go get gopkg.in/go-playground/validator.v9
以上

■vim-go
プロジェクト（$GOPATH）から作り直した場合、プロジェクト配下で適当にvim開いて「GoInstallBinaries」を実行
vim-goが依存しているパッケージをインストール
補完とかできるようになる
参考サイト：
ttps://qiita.com/koara-local/items/6c886eccfb459159c431
