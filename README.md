# Todo-Apps

Go でサーバーサイド（API・レイヤード構成など）を学ぶためのリポジトリです。フロントは静的テンプレートのみ用意してあり、バックエンドは `api-document.yaml`（OpenAPI）と Swagger UI を見ながら自分で実装する想定です。

## 技術スタック（現状）

| 領域 | 技術 |
|------|------|
| フロント | HTML5、CSS3、バニラ JavaScript（ビルドツールなし） |
| API 契約 | OpenAPI 3.0（`api-document.yaml`） |
| ローカル表示 | Docker Compose、`nginx:alpine`（静的配信）、`swaggerapi/swagger-ui` |
| データストア（学習用コンテナ） | MySQL 8、Redis 7、`redis-commander`（任意の管理 UI） |

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

## OpenAPI

- ファイル: `api-document.yaml`
- サーバ URL の例: `http://localhost:8080/api`（Go 実装時に合わせて変更してください）
- 一覧レスポンスは `{ "todos": [...] }` 形式です。フロントは配列直返しにも対応しています。

## 環境変数サンプル

`env-sample` を参考に、MySQL 接続などアプリ用の `.env` を自分で整えてください。

## Makefile（任意）

プロジェクトルートの `Makefile` によく使うコマンドを追加できます。
