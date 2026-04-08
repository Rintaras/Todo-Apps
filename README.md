# Todo-Apps

Go でサーバーサイド（API・レイヤード構成など）を学ぶためのリポジトリです。フロントは静的テンプレートのみ用意してあり、バックエンドは `api-document.yaml`（OpenAPI）と Swagger UI を見ながら自分で実装する想定です。

## 技術スタック（現状）

| 領域 | 技術 |
|------|------|
| フロント | HTML5、CSS3、バニラ JavaScript（ビルドツールなし） |
| API 契約 | OpenAPI 3.0（`api-document.yaml`） |
| ローカル表示 | Docker Compose、`nginx:alpine`（静的配信）、`swaggerapi/swagger-ui` |
| データストア（学習用コンテナ） | MySQL 8、Redis 7、`redis-commander`（任意の管理 UI） |
| バックエンド（学習用） | Go 1.25+、Gin、GORM（`jinzhu/gorm`）、MySQL ドライバ |
| Go モジュール | `Backend/go.mod`（`module todo-apps/backend`）、参考用に `Backend/db/conn.go`（`database/sql`）もあり |

## バックエンド（DB 接続）

- API サーバー（`Backend/server`）は起動時に **`Config.LoadDotEnv()`** で、カレントディレクトリから親ディレクトリへ辿って見つかった **`.env`** を読み込みます（シェルで `export` していなくても動かせます）。
- 必須の環境変数: `MYSQL_USER` / `MYSQL_HOST` / `MYSQL_DATABASE`（`MYSQL_PASSWORD` は空でも可）。`MYSQL_PORT` を省略した場合は **`3306`** とみなします。
- Docker の MySQL をホストマシンから叩くときは `MYSQL_HOST=127.0.0.1`。**アプリを同じ Compose ネットワーク内のコンテナで動かす**ときは `MYSQL_HOST=mysql` にしてください。
- `docker compose up -d mysql` で DB が立ち上がってから `go run ./server` してください。

- TablePlusの場合
```bash
open -a TablePlus "mysql://root:password@127.0.0.1:3306/TodoApp?name=Todo+App&statusColor=007F3D&env=Local"
```

```bash
cd Backend && go build ./...
```

## フロントエンド

- 配置: `Frontend/`
- エントリ: `Frontend/index.html`
- API ベース URL: `Frontend/js/config.js` の `apiBaseUrl`（既定は `/api`）
- `api.js` は OpenAPI のパス（`/todos` など）に合わせた `fetch` ラッパーです。
- API が未接続のときはデモ用データで UI だけ動作確認できます。

### 静的ファイルの確認

```bash
docker compose up -d web
```

ブラウザで http://localhost:3000 を開きます。

## Swagger UI

```bash
docker compose up -d swagger-ui
```

http://localhost:8081 で `api-document.yaml` の内容を参照できます。

ブラウザの「Try it out」は **Swagger（8081）→ API（8080）** となり別オリジンになるため、API 側で **CORS を有効化**しています（`Backend/server/Routes/Routes.go`）。それでも `Failed to fetch` のときは、**`go run ./server` が 8080 で起動しているか**を先に確認してください。

## OpenAPI

- ファイル: `api-document.yaml`
- サーバ URL の例: `http://localhost:8080/api`（Go 実装時に合わせて変更してください）
- 一覧レスポンスは `{ "todos": [...] }` 形式です。フロントは配列直返しにも対応しています。

## データベース初期化

- DDL: `Backend/db/init/ddl.sql`（データベース `TodoApp`、テーブル `todos`）
- `docker compose up` で MySQL コンテナを**初めて**立ち上げるとき、上記が `/docker-entrypoint-initdb.d` 経由で自動実行されます（既存の `db-data` ボリュームがある場合は実行されないため、変更後は `docker compose down -v` でボリュームを消すか、手で `mysql` に流し込んでください）。

## 環境変数サンプル

`env-sample` の内容をコピーしてリポジトリルートに `.env` を作成してください（`.gitignore` で除外済み）。`Backend` 直下から `go run ./server` しても、親の `.env` が読み込まれます。

## Makefile（任意）

プロジェクトルートの `Makefile` によく使うコマンドを追加できます。
