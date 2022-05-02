# test_go
Golang test

## Description
Upon running a web server will start listening on localhost:8081  
See available handlers in [api\api.yaml](api\api.yaml)

## Building/running
Requirements - go 1.18

### Windows
**build:**  
```
go build -o ./build/app.exe
```

**run:**  
```
go run ./
```
or after build
```
build\app.exe
```

**tests:**  
```
go test ./...
```

<details> 
  <summary>Task requirements (in russian)</summary>
В решении можно использовать любые стандартные библиотеки.
Описать структуру клиента, у которого есть идентификатор и баланс. 
Описать структуру лидерборда, у которого есть методы:

1. добавление клиента в лидерборд,
2. печать лидерборда,
2. метод, который по идентификатору клиента, отыскивает:
   а) другого клиента с балансом, которых максимально близок к его балансу сверху
   б) другого клиента с балансом, которых максимально близок к его балансу снизу
   (пример: балансы 10 50 20 40 30, для третьего клиента с балансом 20 нужно вывести первого с балансом 10 и пятого с балансом 30)

Все три метода должны быть доступны через http api (REST, RPC — не принципиально).
Лидерборд должен жить в памяти (не нужно сохранять стейт между перезапусками) и в начале каждого часа добавлять всем клиентам на баланс фиксированную сумму.
Юнит-тесты по усмотрению.
</details>


