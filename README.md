Запуск: docker compose up --build --force-recreate

Проверка:

    создание таска: curl -X POST http://localhost:8080/tasks
    чтение таска: curl http://localhost:8080/tasks/{id}

Значение изменится через 3 минуты 