# 🚀 Go + React + MySQL Starter Template

Минималистичный шаблон для полноценного веб-приложения с горячей перезагрузкой, проксированием и готовой инфраструктурой.

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)](https://golang.org/)
[![React](https://img.shields.io/badge/React-18+-61DAFB?logo=react)](https://react.dev/)
[![Docker](https://img.shields.io/badge/Docker-3.8+-2496ED?logo=docker)](https://www.docker.com/)

---

## ✨ Особенности

- 🛠 **Hot-reload** для Go (Air) и React (Vite)
- 🔄 **Два окружения**: Development & Production
- 📦 Готовая **Docker-инфраструктура**
- 🔒 Автоматическое проксирование через **Caddy**
- 🐘 Поддержка **MySQL** с персистентным хранилищем
- 🌐 Оптимизированные сетевые настройки

---

## 🧰 Технологический стек

| Компонент          | Технологии                  |
| ------------------ | --------------------------- |
| **Бэкенд**         | Go 1.24.1, Air, MySQL  |
| **Фронтенд**       | React 18, Vite, TypeScript  |
| **Инфраструктура** | Docker, Caddy, Alpine Linux |

---

## 🚀 Быстрый старт

1. Склонируйте репозиторий:

```bash
git clone https://github.com/yourusername/go-react-template
cd go-react-template
```

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
- MySQL: [Mys://user:password@localhost:5432/dbname](Mys://user:password@localhost:5432/dbname)

### Особенности продакшн-режима:

- 🏎️ Go-приложение в статическом бинарнике
- 🖼️ React с оптимизированной статической сборкой
- 🔐 Готовый к HTTPS конфиг Caddy (раскомментируйте в `Caddyfile.prod`)
- 🗄️ Персистентные volumes для БД

---
