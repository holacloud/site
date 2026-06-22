# Administrando Colas

Tailon expone colas en memoria.

Usa `POST /v1/queues` para crear, `GET /v1/queues` para listar IDs de colas como `[]string`, y `GET /v1/queues/{queue_id}` para recuperar `name`, `len`, `reads` y `writes`.

El handler de delete no está implementado.
