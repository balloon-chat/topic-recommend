#!/bin/bash
# 必ずプロジェクトのルートディレクトリ上で実行する。
gcloud functions deploy topic-recommend \
  --region asia-northeast1 \
  --entry-point RecommendTopics \
  --runtime go113 \
  --trigger-http \
  --set-env-vars FIREBASE_DATABASE_URL="https://balloon-6bad2-default-rtdb.firebaseio.com/"
