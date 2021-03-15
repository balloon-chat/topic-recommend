# Recommend Topic API

## 概要

作成された話題の中から、注目の話題を取得するサービスです。

## 準備

- Firebase RealtimeDatabaseのURLを指定してください
  ```bash
  export FIREBASE_DATABASE_URL="https://firebaseproject.firebaseio.com/"
  ```

## ローカル(CI)環境で実行する場合の手順

1. Firebase RealtimeDatabaseを作成
2. Firebase > 設定 > プロジェクトの設定 > サービスアカウント から秘密鍵を作成
3. 秘密鍵のパスを指定

  ```bash
  export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/service-account-file.json"
  ```
