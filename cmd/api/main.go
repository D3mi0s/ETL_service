package main

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

/*
etl-project/                  # Корень проекта
├── cmd/                      # Входные точки приложения
│   └── api/                  # Сервер API
│       └── main.go           # Основной файл сервера
├── internal/                 # Внутренние пакеты (не для импорта извне)
│   ├── minio/                # Логика работы с MinIO
│   │   └── client.go         # Инициализация MinIO клиента
│   ├── postgres/             # Логика работы с PostgreSQL
│   │   └── repository.go     # Запросы к БД
│   └── handlers/             # HTTP-обработчики
│       └── upload.go         # Обработчик загрузки файлов
├── migrations/               # SQL-миграции
│   └── 001_init_tables.sql   # Скрипт создания таблиц
├── configs/                  # Конфигурационные файлы
│   └── config.yaml           # Настройки окружения
├── docker-compose.yml        # Docker-контейнеры
├── Dockerfile                # Сборка образа для API (опционально)
├── go.mod                    # Файл зависимостей Go
├── go.sum                    # Хеши зависимостей
└── README.md                 # Инструкции по запуску
*/

// Конфигурация объектного хранилища MinIO, которая должна совпадать с docker-compose.yml
const (
	MinioEndpoint  = "localhost:9000"
	MinioAccessKey = "minioadmin"
	MinioSecretKey = "minioadmin"
	MinioBucket    = "raw_data"
)

func main() {

}

// Инициализация подключения к MinIO
func initMinIO() *minio.Client {
	/*
		Создаем клиента с учетными данными
		minio.New() - это функция библиотеки minio-go. Используется для создания нового клиента, который будет
		использоваться для выполнения операций, например, загрузки, скачивании или удалении объектов на
		сервере MinIO.
		Принимает следующие аршументы:
		1. endpoint (string) - это URL адрес MinIO сервера
		2. accessKey (string) - имя пользователя в системе (логин)
		3. secretKey (string) - пароль пользователя в системе
		4. sercure (bool) - указывает, следует ли использовать HTTPS (true) или HTTP (false)

		В более новых версиях minio-go рекомендуется использовать структуру minio.Options.
	*/
	client, err := minio.New(MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKey, MinioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("Ошибка создания MinIO клиента: %v", err)
	}

	return client
}
