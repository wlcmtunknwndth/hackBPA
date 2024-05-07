# API

## Auth

Для авторизации между страницами и элементами сайта используются jwt-токены. 
Токен один -- Access. Хранится, как Cookie под именем "access". 
Предусмотрена операция Refresh. Access выдается при успешных /register и /login и 
по умолчанию действует 4 минуты. Остальные запросы лишь проверют, есть ли 
токен в Cookie, и обновляют(Refresh) его при необходимости, если TimeToLive 
еще не вытек. Пользователи хранятся в базе данных. Структура User:

```Go
type User struct{
    Username string `json:"username"`
    Password string `json:"password"`
	Age      uint8  `json:"age"`
	Gender   bool   `json:"gender"` // True -- женщина, False -- Мужчина 
    isAdmin  boolean //структура не десериализуется.
```

Роль администратора нельзя задать через запрос. 
Роль задаётся только через базу данных. При входе по правам администратора 
информация присваивается через запрос от ф-ции к БД, а после кодируется
в jwt-токен, поэтому права не пропадают.

### POST /login

{
    "username": "string_value",
    "password": "string_value"
}

### POST /register

{
    "username": "string_value",
    "password": "string_value"
}

### POST /logout

empty body {}. В данном случае лишь стирается access-токен.

### DELETE /delete_user
Для удаления пользователя должен быть jwt-токен администратора.

{ "username": "string_value" }

## EVENTS

Event JSON: 

```JSON
{
  "id": uint,
  "price": uint,
  "restrictions": uint,
  "date": "timestamp as string",
  "feature": "deaf",
  "city": "moscow",
  "address": "Malay Ordinka, 3",
  "name": "Mayhem",
  "img_path": "/path/to/event/image/folder",
  "description": "chainsaw gutsfuck"
}
```

"restriction" is an age restriction, where num means lower bound
"img_path" -- картинки для ивента будут в папке ./data/events/<id>, пронумерованной от 1
"feature" -- для каждой особенности будет свой номер в бд, но для фронта:
    1. disability - "проблемы с мобильностью"
    2. deaf - "проблемы с слухом"
    3. blind - "проблемы c зрением"
    4. neuro - "нейроотличия"


### GET /event?id=

```JSON
{
    "id": uint,
    "price": uint,
    "restrictions": uint,
    "date": "timestamp as string",
    "feature": "deaf",
    "city":"moscow",
    "address":"Malaya Ordinka, 3",
    "name": "Mayhem",
    "img_path": "/path/to/event/image/folder",
    "description": "chainsaw gutsfuck"
}
```

### GET /events?filter=["deaf", "blind" or other features]&ordering=["date", "price"]&order=["ascending", "descending"]&location=['moscow', etc]

```JSON
{
  "events": [
    {
      "id": uint,
      "price": uint,
      "restrictions": uint,
      "date": "timestamp as string",
      "feature": "deaf",
      "city":"moscow",
      "address":"Malaya Ordinka, 3",
      "name": "Mayhem",
      "img_path": "/path/to/event/image/folder",
      "description": "chainsaw gutsfuck"
    }
   ]       
}
```

### POST /create_event

```JSON
{
    "price": uint,
    "restrictions": uint,
    "date": "timestamp as string",
    "feature": "deaf",
    "city":"moscow",
    "address":"Malaya Ordinka, 3",
    "name": "Mayhem",
    "description": "chainsaw gutsfuck"
}
```

При создании возвращается id нового ивента.

id присваивается автоматически для того, чтобы не возникало
проблем

Предположительно, будет рут для загрузки именно фотографий в папку 
с определенный id

### PATCH /patch_event

### POST /create_event?id=<id>

```JSON
{
    "price": uint,
    "restrictions": uint,
    "date": "timestamp as string",
    "feature": "deaf",
    "city":"moscow",
    "address":"Malaya Ordinka, 3",
    "name": "Mayhem",
    "description": "chainsaw gutsfuck"
}
```