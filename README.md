# HTTP routes 

## POST: /Register
input: {"username": string, "password": string}

## POST: /Login
input: {"username": string, "password": string}

## POST: /Logout

input: {} (empty)

Авторизация производится через логин-пароль, администраторы
прописываются через БД. 

## GET: /getMessage
input: {"id": uint}

output: {"message": string}

## GET: /getTask
input: {"id": uint}

output: {
"id": 146099,
"message": "Привет! Покупателю 310597 не возвращаются деньги: статус завис на confirmed. 356911",
"user_id": 205592,
"channel_id": 347308,
"create_at": 1693353138227,
"update_at": 1693353138227,
"root_id": null,
"edit_at": 0,
"delete_at": 0,
"is_pinned": false,
"original_id": ""
},
