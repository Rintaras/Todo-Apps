# Backend タスク一覧

`api-document.yaml`（OpenAPI）とフロント（`Frontend/js/api.js`）を契約の正とし、Go で API を実装するときのチェックリストです。順序は参考用で、自分の学習ペースに合わせてよいです。

## プロジェクト基盤

- [ ] `go mod init` とモジュール名の決定
- [ ] `main` のエントリ（例: `cmd/server/main.go`）と設定の読み込み（環境変数 / `.env` は `env-sample` を参考）
- [ ] HTTP サーバー起動（ポート例: `8080`）、グレースフルシャットダウン（任意）
- [ ] ルーティング（`chi` / `echo` / `net/http` など好みのもの）と `/api` プレフィックス配下に Todo ルートをまとめる

## レイヤード構成（例）

- [ ] **ドメイン / モデル**: `Todo`（id, title, completed, createdAt, updatedAt など）を構造体で定義
- [ ] **リポジトリ（インターフェース）**: 一覧・作成・1件取得・更新・削除のメソッドを宣言（実装は MySQL などに差し替え可能に）
- [ ] **ユースケース / サービス層**: バリデーション・ビジネスルール（タイトル必須・最大長など）をここに集約
- [ ] **ハンドラ（プレゼンテーション）**: JSON の受け渡し、ステータスコード、OpenAPI どおりのレスポンス形に揃える

## API 実装（OpenAPI 準拠）

ベースパスは `http://localhost:8080/api` を想定（`servers` とフロントの `config.js` を一致させる）。

- [ ] `GET /api/todos` → `200`、ボディ `{ "todos": [ ... ] }`
- [ ] `POST /api/todos` → リクエスト `{ "title": "..." }`、`201` + 作成された `Todo`、不正時は `400` + `Error`
- [ ] `GET /api/todos/{id}` → `200` または `404`
- [ ] `PATCH /api/todos/{id}` → 部分更新（`title` / `completed`）、`200` または `404`
- [ ] `DELETE /api/todos/{id}` → `204` または `404`

## 永続化（MySQL 想定 / docker-compose）

- [ ] 接続設定（ホスト `mysql`（Compose 内）または `localhost`、DB 名・ユーザーは `env-sample` と整合）
- [ ] マイグレーションまたは初期 DDL（`todos` テーブル: id, title, completed, timestamps など）
- [ ] リポジトリ実装（プリペアドステートメント、トランザクションが必要ならユースケース側で）

## 横断的な関心事

- [ ] **CORS**: フロントを別ポート（例: `3000`）から開く場合、`Access-Control-Allow-Origin` 等の設定
- [ ] **ログ**: リクエスト ID、エラー時のスタック／メッセージ（本番では情報漏れに注意）
- [ ] **バリデーション**: タイトル必須・最大長（OpenAPI の `maxLength: 500` に合わせる）
- [ ] **404 / 400 の JSON 形**: `components/schemas/Error`（例: `{ "message": "..." }`）に揃える

## 品質・運用（任意だが学習におすすめ）

- [ ] ハンドラまたはユースケースのユニットテスト / 結合テスト
- [ ] OpenAPI と実装の乖離チェック（手動で Swagger UI と突き合わせ、または codegen / バリデータ利用）
- [ ] コンテナ化する場合はリポジトリ直下に `Dockerfile` を追加し、`docker-compose` の `go` サービスと接続（別タスク）

## 参照

- API 契約: `api-document.yaml`
- Swagger UI: `docker compose up -d swagger-ui` → http://localhost:8081
- フロントが期待するパス: `Frontend/js/api.js`
