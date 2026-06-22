# Ejecutar

Run es el servicio de ejecución de contenedores de HolaCloud. Gestiona contenedores Docker a través de una interfaz compatible con Docker Registry v2, permitiéndote desplegar y ejecutar aplicaciones en entornos aislados.

## Características

### Compatible con Docker Registry v2
Run es un registro Docker Registry v2 completamente compatible. Sube y baja imágenes usando comandos Docker estándar — sin necesidad de herramientas personalizadas.

```bash
docker login run.hola.cloud
docker build -t run.hola.cloud/mi-proyecto/mi-app:latest .
docker push run.hola.cloud/mi-proyecto/mi-app:latest
```

### Administración desde la Consola
Gestiona contenedores, imágenes y despliegues a través de la Consola de HolaCloud. Inicia, detén y reinicia contenedores, visualiza registros y monitorea el uso de recursos — todo desde una interfaz web.

### Variables de Entorno
Configura tus contenedores con variables de entorno. Establécelas en el momento del despliegue a través de la Consola o la API para despliegues flexibles basados en configuración.

```bash
curl -X POST "https://api.hola.cloud/v1/run/deploy" \
  -H "Authorization: Bearer tu-token" \
  -d '{
    "image": "run.hola.cloud/mi-proyecto/mi-app:latest",
    "env": {
      "DATABASE_URL": "postgres://...",
      "LOG_LEVEL": "debug"
    }
  }'
```

### Volúmenes
Adjunta volúmenes persistentes a tus contenedores para cargas de trabajo con estado. Los volúmenes están respaldados por el almacenamiento distribuido de HolaCloud y sobreviven a reinicios del contenedor.

## Casos de Uso

### Despliegue Completo de Aplicaciones
Despliega aplicaciones web, APIs y microservicios como contenedores Docker. Run gestiona el enrutamiento de red, las verificaciones de salud y los reinicios automáticos.

### Entornos de Desarrollo
Crea entornos de desarrollo aislados para cada rama o desarrollador. Usa Run para crear contenedores efímeros que coincidan con tu configuración de producción.

### Pipelines CI/CD
Integra Run en tu flujo de trabajo CI/CD. Sube imágenes al registro como parte del proceso de compilación, luego despliégalas automáticamente en entornos de staging o producción.

## Primeros Pasos

1. **Autentícate** en el registro:
   ```bash
   docker login run.hola.cloud
   ```

2. **Compila y sube** tu imagen:
   ```bash
   docker build -t run.hola.cloud/mi-proyecto/mi-app:latest .
   docker push run.hola.cloud/mi-proyecto/mi-app:latest
   ```

3. **Despliega** a través de la Consola o la API:
   ```bash
   curl -X POST "https://api.hola.cloud/v1/run/deploy" \
     -H "Authorization: Bearer tu-token" \
     -d '{"image": "run.hola.cloud/mi-proyecto/mi-app:latest"}'
   ```

Tu contenedor se iniciará y estará disponible en el endpoint asignado.
