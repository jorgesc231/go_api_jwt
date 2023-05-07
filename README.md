# Ejemplo de API con autenticacion usando JWT

Ejemplo de API REST con autenticacion de usuarios usando tokens [JWT](https://jwt.io/).

Usa [Fiber](https://docs.gofiber.io/) como framework para el backend y [gorm](https://gorm.io/) como ORM para la base de datos.

Usa **jwt-go** para los tokens `jwt`.

Usa **SQLite** como base de datos de ejemplo.

## Tiene los siguientes endpoints:


### Ping

```
GET /api/ping
```

Ejemplo:


```
http://127.0.0.1:8000/api/ping
```

Retorna un codigo 200 y un JSON con el mensaje: *"success"*.


### Registro

```
POST /api/register
```

Ejemplo:

```
http://127.0.0.1:8000/api/register
```

Enviar en el body un JSON con el siguiente contenido:

```
{
    "name": "prueba",
    "email": "prueba@mail.com",
    "password": "123456789"
}
```

Retorna un codigo 200 con un JSON con el id de usuario, nombre y email.  
Si el usuario ya existe retorna un codigo 400 y un json con el mensaje *Already in use*.


### Inicio de Sesion

```
POST /api/login
```

Ejemplo:


```
http://127.0.0.1:8000/api/login
```

Enviar en el body un JSON con el siguiente contenido:

```
{
    "email": "prueba@mail.com",
    "password": "123456789"
}
```

Retorna un codigo 200 y crea una Cookie con el token *jwt* si se inicio sesion correctamente.  
Si la cuenta no existe retorna un codigo 404 y un JSON con el mensaje *"Account not found"*.        

Si la contrasena es incorrecta retorna un codigo 400 y un JSON con el mensaje *"Incorrect Password"*.


### Obtener datos del usuario (requiere estar autenticado)

```
GET /api/user
```

Ejemplo:

```
http://127.0.0.1:8000/api/user
```

Si no se esta auntenticado retorna un codigo 401 y un JSON con el mensaje *"Not Authenticated"*.  

Si se esta auntenticado retorna un codigo 200 y un JSON con el id de usuario, nombre y email.


### Cierre de Sesion

```
POST /api/logout
```

Ejemplo:

```
http://127.0.0.1:8000/api/logout
```

Invalida y elimina la Cookie con el token *jwt* y retorna un codigo 200.


*Se le podria agregar cualquier otro tipo de funcionalidad.*

## Compilar

1. Instalar Go:

[https://go.dev/dl/]()

2. Clonar repositorio:

```
git clone https://github.com/jorgesc231/go_api_jwt.git
cd go_api_jwt
```

3. Compilar:


```
go build main.go
```

4. Ejecutar:

```
./main
```



## Compilar para usar una raspberry pi como servidor


[Tutorial](https://github.com/jorgesc231/tutoriales_raspberry_pi/tree/master/go_cross_rpi1)


## Aplicacion que consume la API


[https://github.com/jorgesc231/api_jwt_client]()