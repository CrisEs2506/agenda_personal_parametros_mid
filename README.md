# Parametros_API
API MID entre CRUDs de [Agenda_Personal](https://github.com/CrisEs2506/agenda_personal_crud) y [Parametros](https://github.com/CrisEs2506/parametros_crud) elaborado con lenguaje Goland, el framework BeeGo y el motor de bases de datos relacional PostgreSQL que realiza una petición GET para traer todos los contactos de la Agenda_Personal y los registros con nombre Cargos en la tabla tipo_parametros de Parametros.
## Tecnologías Implementadas
- Golang
- Beego
- PostgreSQL
## Instalación
> Dirigirse a la carpeta de go en la que quiere clonar el proyecto 

`cd go/src`
> Clonar el repositorio

`git clone https://github.com/CrisEs2506/agenda_personal_parametros_mid.git`
> Entrar en la carpeta del proyecto clonado

`cd agenda_personal_parametros_mid`
> Instalar dependencias del proyecto

`go get`
> Crear archivo para las variables de entorno .env
```
export AGENDA_PERSONAL_MID_PARAMETROS_PROTOCOL_ADMIN = [nombre de protocolo, en este caso http]
export AGENDA_PERSONAL_MID_PARAMETROS_AGENDA_PERSONAL_URL = [url en donde esta corriendo el CRUD Agenda_Personal]
export AGENDA_PERSONAL_MID_PARAMETROS_PARAMETROS_URL = [url en donde esta corriendo el CRUD Parametros]
```
## Ejecución del Proyecto
`bee run`