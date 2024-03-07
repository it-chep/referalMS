# **Микросервис "Реферальная система"**


# API Endpoints Documentation

## New Referal

### URL
/api/v1/new_referal/

### Description
This endpoint is used to create a new referal.

### Headers
```json
{
  "Admin-Login": "string", //required; Login for User Admin
  "Admin-Password": "string", //required; Password for User Admin
  "Admin-Token": "string" //required; Intergation Token 
}
```

### Body
```json
{
  "tg_id": "integer", //required
  "id_in_integration_service": "integer", 
  "name": "string", //required
  "username": "string" //required
}
```

### Returns
```json
{
  "status": 200,
  "error": "", 
  "referal_link": "string"
}
```

## New User

### URL
/api/v1/new_user/

### Description
This endpoint is used to create a new user.

### Headers
```json
{
  "Admin-Login": "string", //required; Login for User Admin
  "Admin-Password": "string", //required; Password for User Admin
  "Admin-Token": "string" //required; Intergation Token 
}
```

### Body
```json
{
    "tg_id": "integer", //required
    "in_service_id": "integer", 
    "name": "string", //required
    "username": "string", //required
    "referal_link": "string"
}
```

### Returns
```json
{
  "status": 200,
  "error": ""
}
```

## Get Statistic

### URL
/api/v1/get_statistic/

### Description
This endpoint is used to get statistics for a referal.

### Headers
```json
{
  "Admin-Login": "string", //required; Login for User Admin
  "Admin-Password": "string", //required; Password for User Admin
  "Admin-Token": "string" //required; Intergation Token 
}
```

### Body
```json
{
  "tg_id": "integer",
  "in_service_id": "integer"
}
```
### Returns
```json
{
    "status": 200,
    "error": "",
    "all_users": "integer", // count users for all time
    "last_n_days": "integer" // count users for last integer days
}
```

## Get Winners

### URL
/api/v1/get_winners/

### Description
This endpoint is used to get the best referals

### Headers
```json
{
  "Admin-Login": "string", //required; Login for User Admin
  "Admin-Password": "string", //required; Password for User Admin
  "Admin-Token": "string" //required; Intergation Token 
}
```

### Body
```json
{
    "limit": "integer", // limit best refeals
    "days_interval": "integer" // days interval: current date - integer days
}
```

### Returns
```json
{
    "status": "200",
    "error": "",
    "users": [
        {
            "name": "string",
            "users_count": "integer",
            "username": "string",
            "in_service_id": "integer"
        }
    ]
}
```
