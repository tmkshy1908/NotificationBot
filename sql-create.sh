#!/bin/bash
INSTANCE_NAME="notification-db"
if gcloud sql instances describe $INSTANCE_NAME &> /dev/null; then
    echo "インスタンス ${INSTANCE_NAME} は既に存在します。"
else
    echo "インスタンス ${INSTANCE_NAME} は存在しません。作成を開始します。"
    gcloud sql instances create notification-db --database-version=POSTGRES_12 --tier=db-f1-micro  --region=asia-northeast1
fi
