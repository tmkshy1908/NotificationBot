#!/bin/bash
gcloud sql instances create notification-db --database-version=POSTGRES_12 --tier=db-f1-micro  --region=asia-northeast1
