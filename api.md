# API

## Auth(в разработке)

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
isAdmin  boolean //структура не десериализуется.
```

Роль администратора нельзя задать через запрос. 
Роль задаётся только через базу данных. При входе по правам администратора 
информация присваивается через запрос от ф-ции к БД, а после кодируется
в jwt-токен, поэтому права не пропадают. 


}

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