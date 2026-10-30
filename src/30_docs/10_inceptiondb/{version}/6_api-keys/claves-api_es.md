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
[{"defaults":{"id":"uuid()"},"indexes":0,"name":"customers","total":2}]
```

Una clave o un secreto incorrectos responden `403 Forbidden`.

## Mantenla segura

- InceptionDB es un API de servidor a servidor. No lo llames desde código que se ejecuta en un navegador o en una app móvil: ahí cualquiera puede leer el secreto. Si una página web necesita datos, haz que llame a tu propio backend, que guarda el secreto y llama a InceptionDB.
- Guarda el secreto en tu servidor, en variables de entorno o en un gestor de secretos. Nunca en el código, en un repositorio ni en una página web o app móvil: cualquiera que lo vea puede leer y cambiar tus datos.
- Usa una API key por aplicación o entorno. Así puedes revocar una sin afectar a las demás.
- Si un secreto se filtra, crea una clave nueva, despliégala y **revoca** la antigua en **API keys & settings**. Las claves revocadas dejan de funcionar al instante.

## Owners

Los owners son los usuarios de hola.cloud que gestionan una base de datos desde la consola: sus colecciones, API keys y otros owners. Se añaden y se quitan en **API keys & settings**. Una base de datos siempre conserva al menos un owner, y no puedes quitarte a ti mismo.

## Operaciones de cuenta: solo en la consola

Crear, renombrar y borrar bases de datos, y gestionar API keys y owners, se hace en la consola. La [referencia del API](/es/docs/inceptiondb/v2/api-reference/) lista sus endpoints para explicar lo que hace la consola, pero no están pensados para automatizar:

- `Api-Key` y `Api-Secret` no funcionan con ellos. Una API key solo sirve para las colecciones de su base de datos.
- Necesitan la sesión de un usuario identificado en la consola, y no hay una forma admitida de obtenerla fuera de la consola.

## Ver el código de una acción de la consola

Cada acción de la consola es una petición HTTP, y la consola muestra cómo hacerla desde tu código. Pulsa **API calls**, con el icono `{ }`, al pie del menú lateral. Lista las últimas peticiones; elige una para verla en curl, Go, Python y los demás lenguajes, lista para copiar. Cada petición nueva aparece además como un punto de color en la columna estrecha del borde derecho de la consola; al pulsar un punto se abre la misma vista.
