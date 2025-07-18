# 🐹 Golang Basic CRUD API

Este es un proyecto básico en Go que implementa un CRUD (Create, Read, Update, Delete) de tareas usando la librería `gorilla/mux` como router HTTP.

Incluye:
- Manejo de rutas RESTful
- Lectura y escritura de JSON
- Uso de `map[int]task` como almacenamiento en memoria
- Hot reload con `CompileDaemon` (opcional)

---

## 📦 Instalación

1. **Inicializar módulo de Go**

```bash
go mod init golang-api
```

2. **Instalar Gorilla Mux**

```bash
go get github.com/gorilla/mux@latest
```

3. **Instalar CompileDaemon (opcional, para desarrollo en caliente)**

```bash
go install github.com/githubnemo/CompileDaemon@latest
```

4. **Correr el servidor con hot reload**

```bash
CompileDaemon -command="./golang-api"
```

> ⚠️ Asegurate de tener `$HOME/go/bin` en tu `PATH` para poder ejecutar `CompileDaemon`.

---

## 🚀 Endpoints de la API

Todos los endpoints devuelven y reciben datos en formato **JSON**.

### `GET /`
Devuelve un mensaje de bienvenida.

```bash
curl http://localhost:3000/
```

---

### `GET /tasks`
Devuelve todas las tareas almacenadas.

```bash
curl http://localhost:3000/tasks
```

---

### `POST /task`
Crea una nueva tarea. Requiere un JSON con `name` y `content`.

```bash
curl -X POST http://localhost:3000/task \
     -H "Content-Type: application/json" \
     -d '{"name": "Nueva tarea", "content": "Contenido de la tarea"}'
```

---

### `GET /task/{id}`
Devuelve una tarea específica por su ID.

```bash
curl http://localhost:3000/task/1
```

---

### `PUT /task/{id}`
Actualiza una tarea existente. Requiere un JSON con `name` y `content`.

```bash
curl -X PUT http://localhost:3000/task/1 \
     -H "Content-Type: application/json" \
     -d '{"name": "Tarea actualizada", "content": "Contenido actualizado"}'
```

---

### `DELETE /task/{id}`
Elimina una tarea por su ID.

```bash
curl -X DELETE http://localhost:3000/task/1
```

---


## 🧑‍💻 Autor

Lautaro Vissio — [GitHub](https://github.com/vissio)

---

## 📜 Licencia

MIT — libre para copiar, modificar y compartir.
