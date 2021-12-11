
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
| /battle | { battleId: id } | 〇 | POST | 対戦情報のIDを元に認証 |


## バトル処理フロー

### ルーム参加まで

1. ルーム作成（6桁のルーム番号を返す）
2. ルーム番号を持って /room POST を叩く

### ルーム認証からバトル画面移行まで（参加者）

1. ルーム番号が存在していたらそのルームを閉じる
2. 両者の対戦用のチームデータを作成
3. 作成した情報をもとに対戦情報を作成
4. 参加者に対戦情報を持った認証トークンを付与
5. Pusherでルーム待機者に対戦情報のIDを送信
6. 対戦情報を返す
7. 待機画面に遷移

### ルーム作成からバトル画面移行まで（作成者）

1. Pusherで対戦情報のIDが送られてきたら、それをもとに `/battle` POST を叩く
2. 対戦情報が存在するか、参加しているかをチェック
3. 対戦情報を持った認証トークンを付与
4. 対戦情報を返す
5. 待機画面に遷移

### 対戦の流れ

1. opponentTeams, userTeams のTeamsから`Agility`の高い順に並び替え
