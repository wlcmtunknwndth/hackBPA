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
    "location": "string",
    "name": "string",
    "description": "string",
    "img_path": "string",
    "date": "timestamp as string",
    "features": {
        "disability": bool,
        "deaf": bool
        "blind": bool
        "neural": bool
    }
}
```

"restriction" is an age restriction, where num means lower bound

### GET /event?id=

```JSON
{
    "id": uint,
    "price": uint,
    "restrictions": uint,
    "location": "string",
    "name": "string",
    "description": "string",
    "img_path": "/path/to/event/image/folder",
    "date": "timestamp as string",
    "features": {
        "disability": bool,
        "deaf": bool,
        "blind": bool,
        "neural": bool
    }
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
      "location": "string",
      "name": "string",
      "description": "string",
      "img_path": "string",
      "date": "timestamp as string",
      "features": {
        "disability": bool,
        "deaf": bool
        "blind": bool
        "neural": bool
      }
    }
   ]       
}
```

### GET