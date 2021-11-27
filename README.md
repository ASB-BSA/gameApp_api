
# Readme.md

Readme

## APIエンドポイント

### 前提

```env
BASEURL=https://domain/api/v1
```

### 権限

| 値 | 意味 |
|:--:| -- |
| × | 権限無し |
| 〇 | ユーザーログインのみ |
| ◎ | 対戦者のみ |
| - | 管理者のみ |

### API一覧

| エンドポイント | 値 | 権限 | メソッド | 詳細 |
|--|--|:--:|:--:|--|
| /image/:id | XXX.png | × | GET | 画像リソースを取得する |
| /user | { name: string } | × | POST | ユーザー登録 |
| /settings/:id | XXX | 〇 | GET | 登録されている設定取得 |
| /user | - | 〇 | GET | ユーザー情報取得 |
| /character | - | 〇 | GET | ゲーム内のキャラクターリソースを取得する |
| /team/:id | 編集するキャラID | 〇 | PUT | チームキャラクター情報更新 |
| /characteristic | - | 〇 | GET | 特技一覧を取得 |
| /room | - | 〇 | GET | 対戦ルーム作成 |
| /room | { roomNumber: 000000 } | 〇 | POST | ルームに参加 |
