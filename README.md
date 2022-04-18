# L0

## WRK тест
    Running 10s test @ http://localhost:8080/90gds8EkcY3O0mVTt4H
    400 goroutine(s) running concurrently
    68462 requests in 10.011016026s, 147.03MB read
| Index | Value|
| ------------- | ----------|
|Requests/sec: | 6838.67 |
| Transfer/sec: |  14.69MB  |
| Avg Req Time: | 58.490935ms|
| Fastest Request: | 0s  |
| Slowest Request:  | 411.0189ms|
| Number of Errors:  | 0  |

Задание:
В сервисе:
1. Подключение и подписка на канал в nats-streaming
2. Полученные данные писать в Postgres
3. Так же полученные данные сохранить in memory в сервисе (Кеш)
4. В случае падения сервиса восстанавливать Кеш из Postgres
5. Поднять http сервер и выдавать данные по id из кеша
6. Сделать простейший интерфейс отображения полученных данных, для
их запроса по id
