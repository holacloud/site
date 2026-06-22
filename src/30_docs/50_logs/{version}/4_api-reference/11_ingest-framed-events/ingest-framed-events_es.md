
# Ingest Framed Events

Ingresa eventos enmarcados usando el protocolo logframe. Este endpoint acepta un protocolo binario con prefijos de longitud de metadatos y mensaje.

## Autenticación

Requiere credenciales de datos:

- `X-Instantlogs-Event-Secret: <secret>` o `Authorization: Bearer <secret>`

## Cuerpo de la Solicitud

El cuerpo de la solicitud utiliza el formato del protocolo logframe:

```
len_metadata:len_message:metadata_json:message\n
```

Cada marco consiste en:
- `len_metadata` — Longitud del JSON de metadatos (como número ASCII)
- `len_message` — Longitud del mensaje (como número ASCII)
- `metadata_json` — Objeto JSON con metadatos del evento
- `message` — El contenido del mensaje de log
- `\n` — Delimitador de nueva línea

## Solicitud

```bash
# Marco: metadata={"service":"web"} + message="Usuario conectado"
# len_metadata=18, len_message=17
# Formato: 18:17:{"service":"web"}Usuario conectado\n
printf '18:17:{"service":"web"}Usuario conectado\n' | \
curl -X POST "https://api.hola.cloud/v1/ingest/events" \
  -H "X-Instantlogs-Event-Secret: lgs_abc123def456" \
  -H "Content-Type: application/octet-stream" \
  --data-binary @-
```

```http
POST /v1/ingest/events HTTP/1.1
Host: api.hola.cloud
X-Instantlogs-Event-Secret: lgs_abc123def456
Content-Type: application/octet-stream

18:17:{"service":"web"}Usuario conectado\n
```

## Respuesta

```json
{
  "ingested": 1
}
```

## Códigos de Error

| Código | Descripción |
|--------|-------------|
| 400 | Marco malformado — prefijo de longitud inválido o delimitador faltante |
| 401 | Secreto de evento faltante o inválido |
| 413 | Carga útil demasiado grande |
