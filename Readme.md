## Распределенная система на Go
Что может быть лучше практики - только еще больше практики.
Работая архитектором, нельзя быть только теоретиком.
Пробую на практических примерах реализовывать алгоритмы согласованности и некий аналог ***WAL***. 
Сервис поддерживат два интерфейса - ***REST*** и ***gRPC***.
В качестве http handler'а используетм Gorilla.mux (возможно потом на GIN заменю)

Пакет ***internal/log***:
- **storage.go** - некая обертка над файлом, в который пишется лог
- **index.go** - собственно название отражает суть - это индекс для упрощения обращений к логам
- **config.go** - задает некоторые настройки для лога, в частности настройки размеров сегмента
- **segment.go** - объединяет индекс и хранилище. Когда необходимо добавить запись в лог, сегмент добавляет запись в sorage и создает запись в index.