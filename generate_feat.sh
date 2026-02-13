#!/bin/bash

# Cek jika argumen tidak ada
if [ -z "$1" ]; then
  echo "❌ Tolong masukkan nama fitur!"
  echo "Contoh: ./generate-fitur.sh device"
  exit 1
fi

FEATURE_NAME=$1
BASE_DIR="./internal/features/$FEATURE_NAME"

# Cek jika folder sudah ada
if [ -d "$BASE_DIR" ]; then
  echo "❌ Fitur '$FEATURE_NAME' sudah ada!"
  exit 1
fi

# Buat struktur folder
mkdir -p $BASE_DIR/delivery/http
mkdir -p $BASE_DIR/dto
mkdir -p $BASE_DIR/repository
mkdir -p $BASE_DIR/service

# Fungsi buat file dengan package name
create_file_with_package() {
    FILE_PATH=$1
    PACKAGE_NAME=$2
    echo "package $PACKAGE_NAME" > "$FILE_PATH"
}

# File delivery/http
create_file_with_package "$BASE_DIR/delivery/http/delivery_config.go" "http"
create_file_with_package "$BASE_DIR/delivery/http/${FEATURE_NAME}_delivery.go" "http"
create_file_with_package "$BASE_DIR/delivery/http/${FEATURE_NAME}_routes.go" "http"

# File dto
create_file_with_package "$BASE_DIR/dto/req.go" "dto"
create_file_with_package "$BASE_DIR/dto/res.go" "dto"

# File repository
create_file_with_package "$BASE_DIR/repository/${FEATURE_NAME}_repository.go" "repository"

# File service
create_file_with_package "$BASE_DIR/service/${FEATURE_NAME}_service.go" "service"
create_file_with_package "$BASE_DIR/service/service_config.go" "service"

# File interfaces di root fitur
create_file_with_package "$BASE_DIR/delivery_interface.go" "$FEATURE_NAME"
create_file_with_package "$BASE_DIR/repository_interface.go" "$FEATURE_NAME"
create_file_with_package "$BASE_DIR/service_interface.go" "$FEATURE_NAME"

echo "✅ Template fitur '$FEATURE_NAME' berhasil dibuat di $BASE_DIR"
