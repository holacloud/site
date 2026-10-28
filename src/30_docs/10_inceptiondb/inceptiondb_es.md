# InceptionDB

InceptionDB es una base de datos de documentos JSON que se usa por HTTP. No hay que instalar ningún driver ni definir un esquema: envías JSON y recibes JSON, desde cualquier lenguaje que sepa hacer una petición HTTP.

## Cómo se organiza

| Concepto | Qué es |
|---|---|
| **Base de datos** | Tu espacio. Contiene colecciones y tiene sus propias API keys. Se crea en la [consola](https://inceptiondb.hola.cloud/). |
| **Colección** | Un grupo de documentos con nombre, como `customers` u `orders`. Los documentos de una colección pueden tener campos distintos. |
| **Documento** | Un objeto JSON. Todo documento tiene un `id`; si no lo envías, InceptionDB lo genera. |
| **Índice** | Opcional. Hace rápidas las búsquedas y los rangos sobre un campo, y puede impedir valores repetidos. |
| **API key** | Una clave y un secreto que tu aplicación envía en las cabeceras `Api-Key` y `Api-Secret` para acceder a los datos de una base de datos. |

Todas las peticiones van a `https://inceptiondb.hola.cloud`, y los ejemplos leen tus credenciales de tres variables de entorno, así que puedes copiarlos y pegarlos tal cual:

```bash
export INCEPTIONDB_DATABASE="your-database-id"
export INCEPTIONDB_API_KEY="your-api-key"
export INCEPTIONDB_API_SECRET="your-api-secret"
```

## Por dónde empezar

- **[Primeros pasos](/es/docs/inceptiondb/v2/primeros-pasos/)**: de cero a tu primera consulta en cinco minutos.
- **[Colecciones](/es/docs/inceptiondb/v2/manejando-colecciones/)**: crear colecciones, valores por defecto y tamaño.
- **[Trabajando con datos](/es/docs/inceptiondb/v2/trabajando-con-datos/)**: insertar, leer, actualizar y borrar documentos.
- **[Consultas](/es/docs/inceptiondb/v2/consultas/)**: filtros, operadores, campos anidados, arrays y paginación.
- **[Índices](/es/docs/inceptiondb/v2/indices/)**: consultas rápidas y valores únicos.
- **[API keys](/es/docs/inceptiondb/v2/claves-api/)**: crear bases de datos y claves, y mantenerlas seguras.
- **[Referencia del API](/es/docs/inceptiondb/v2/api-reference/)**: cada endpoint con su petición y su respuesta.

## Para qué encaja bien

- **Datos de aplicación**: usuarios, ajustes, carritos y pedidos: cualquier JSON cuya forma cambia con el tiempo.
- **Contenido**: artículos, páginas y catálogos de productos con campos que varían de un elemento a otro.
- **Eventos y telemetría**: lecturas de dispositivos y analítica de producto escritas a gran velocidad.
- **Prototipos**: empieza a guardar datos en minutos y añade índices cuando las consultas estén claras.
