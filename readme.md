# 🚀 Go + React + PostgreSQL Starter Template

Минималистичный шаблон для полноценного веб-приложения с горячей перезагрузкой, проксированием и готовой инфраструктурой.

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)](https://golang.org/)
[![React](https://img.shields.io/badge/React-18+-61DAFB?logo=react)](https://react.dev/)
[![Docker](https://img.shields.io/badge/Docker-3.8+-2496ED?logo=docker)](https://www.docker.com/)

![Architecture Diagram](https://via.placeholder.com/800x400.png?text=Go+➔+React+➔+PostgreSQL+%7C+Caddy+as+reverse+proxy)

---

## ✨ Особенности

- 🛠 **Hot-reload** для Go (Air) и React (Vite)
- 🔄 **Два окружения**: Development & Production
- 📦 Готовая **Docker-инфраструктура**
- 🔒 Автоматическое проксирование через **Caddy**
- 🐘 Поддержка **PostgreSQL** с персистентным хранилищем
- 🌐 Оптимизированные сетевые настройки

---

## 🧰 Технологический стек

| Компонент          | Технологии                  |
| ------------------ | --------------------------- |
| **Бэкенд**         | Go 1.24.1, Air, PostgreSQL  |
| **Фронтенд**       | React 18, Vite, TypeScript  |
| **Инфраструктура** | Docker, Caddy, Alpine Linux |

---

## 🚀 Быстрый старт

1. Склонируйте репозиторий:

```bash
git clone https://github.com/yourusername/go-react-template
cd go-react-template
```

````

2. Запустите в нужном режиме:

```bash
# Development mode
docker-compose -f docker-compose.dev.yml up --build

# Production mode
docker-compose -f docker-compose.prod.yml up --build -d
```

3. Доступные эндпоинты:

- Фронтенд: [http://localhost](http://localhost)
- Бэкенд API: [http://http://localhost/api/\*](http://localhost/api/*)
- PostgreSQL: [postgres://user:password@localhost:5432/dbname](postgres://user:password@localhost:5432/dbname)

## 🛠 Конфигурация

### 🔧 Переменные окружения

Создайте `.env` файл в корне проекта:

#### backend

```bash
DB_HOST=postgres
DB_USER=user
DB_PASSWORD=password
DB_NAME=dbname
DB_PORT=5432
```

#### frontend

```bash
VITE_API_BASE_URL=/api
```

### Особенности продакшн-режима:

- 🏎️ Go-приложение в статическом бинарнике
- 🖼️ React с оптимизированной статической сборкой
- 🔐 Готовый к HTTPS конфиг Caddy (раскомментируйте в `Caddyfile.prod`)
- 🗄️ Персистентные volumes для БД

---

```

```
````
