# API keys

Tus aplicaciones acceden a una base de datos con una **API key**: una clave y un secreto que se envían en dos cabeceras. Las bases de datos, las API keys y los owners se gestionan en la [consola de InceptionDB](https://inceptiondb.hola.cloud/), con la sesión de tu cuenta de hola.cloud.

## Crear una base de datos

En la consola, abre el selector de base de datos y elige **New database**. Está lista en un segundo. Su ID aparece en la página de resumen: es el valor de `INCEPTIONDB_DATABASE`.

## Crear una API key

1. Elige la base de datos y abre **API keys & settings**.
2. Pulsa **New API key** y ponle el nombre de la aplicación que la va a usar, como `backend-production`.
3. Copia el secreto. Se muestra **una sola vez**: InceptionDB solo guarda un hash.

La consola te da las credenciales listas para usar, por ejemplo como variables de entorno:

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

Una API key da acceso a todas las colecciones de su base de datos, y solo a esa base de datos.

## Usarla

Envía la clave y el secreto en las cabeceras `Api-Key` y `Api-Secret` de cada petición:

```bash
curl "https://inceptiondb.hola.cloud/v1/databases/$INCEPTIONDB_DATABASE/collections" \
  -H "Api-Key: $INCEPTIONDB_API_KEY" \
  -H "Api-Secret: $INCEPTIONDB_API_SECRET"
```

```json
[{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":3}]
```

Una clave o un secreto incorrectos responden `403 Forbidden`.

## Mantenla segura

- Guarda el secreto en tu servidor, en variables de entorno o en un gestor de secretos. Nunca en el código, en un repositorio ni en una página web o app móvil: cualquiera que lo vea puede leer y cambiar tus datos.
- Usa una API key por aplicación o entorno. Así puedes revocar una sin afectar a las demás.
- Si un secreto se filtra, crea una clave nueva, despliégala y **revoca** la antigua en **API keys & settings**. Las claves revocadas dejan de funcionar al instante.

## Owners

Los owners son los usuarios de hola.cloud que gestionan una base de datos desde la consola: sus colecciones, API keys y otros owners. Se añaden y se quitan en **API keys & settings**. Una base de datos siempre conserva al menos un owner, y no puedes quitarte a ti mismo.

## Operaciones de cuenta y el API

Crear, renombrar y borrar bases de datos, y gestionar API keys y owners, se hace en la consola. También son endpoints HTTP (ver la [referencia del API](/es/docs/inceptiondb/v2/api-reference/)), pero necesitan la sesión de un usuario identificado: una API key no puede llamarlos. Cada acción de la consola muestra su petición: pulsa las bolas de la derecha de la pantalla.
