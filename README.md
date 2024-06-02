# Proyecto-Arq-Soft
# Proyecto-Arq-Soft
# Proyecto de Gestión de Cursos

## Descripción
Este proyecto es una aplicación web para la gestión de cursos, donde los usuarios pueden inscribirse en diferentes cursos y ver los detalles de cada uno. El proyecto está dividido en un backend desarrollado en Golang y un frontend desarrollado en React.

## Características
- Autenticación de usuarios (alumnos y administradores)
- Listado de cursos disponibles
- Detalles del curso
- Inscripción en cursos
- Visualización de cursos inscritos

## Requisitos Previos
Asegúrate de tener instalados los siguientes programas:
- [Node.js](https://nodejs.org/) (versión 14 o superior)
- [npm](https://www.npmjs.com/)
- [Go](https://golang.org/) (versión 1.16 o superior)
- [Docker](https://www.docker.com/products/docker-desktop)

## Instalación
1. Clona el repositorio:
   ```bash
   git clone https://github.com/tu-usuario/tu-repositorio.git
   cd tu-repositorio

2. Configura el backend:
cd backend
go mod download

3.Configura el frontend:
cd ../frontend
npm install

## Uso
## Localmente
1. Inicia el backend:
cd backend
go run main.go

2. Inicia el frontend:
cd ../frontend
npm start

## Con Docker
1. Construye los contenedores:
docker-compose build

2. Inicia los contenedores:
docker-compose up
La aplicación frontend estará disponible en http://localhost:5000 y el backend en http://localhost:8080.

## Scripts Disponibles
Frontend
1. `npm start`: Inicia el servidor de desarrollo.
2. `npm run build`: Compila la aplicación para producción.
3. `npm test`: Ejecuta las pruebas.
Backend
1. `go run main.go`: Inicia el servidor backend.

## Configuración de Docker
El archivo docker-compose.yml está configurado para orquestar el frontend y el backend utilizando Docker.

docker-compose.yml: 

version: '3.8'

services:
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - GO_ENV=production

  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    ports:
      - "5000:5000"
    environment:
      - NODE_ENV=production
    depends_on:
      - backend
      
## Estructura del Proyecto

Proyecto-Arq-Soft/
│
├── backend/           # Código del backend en Golang
│   ├── controllers/
│   ├── domain/
│   ├── DTOs/
│   ├── middleware/
│   ├── router/
│   ├── services/
│   ├── go.mod
│   ├── go.sum
│   └── main.go
│
├── frontend/          # Código del frontend en React
│   ├── node_modules/
│   ├── public/
│   ├── src/
│   ├── .gitignore
│   ├── package.json
│   ├── package-lock.json
│   └── Dockerfile
│
├── docker-compose.yml # Configuración de Docker Compose
└── README.md          # Archivo README


## Contribuciones
¡Las contribuciones son bienvenidas! Por favor, abre un issue o un pull request para discutir cualquier cambio importante.

## Contacto
Para preguntas o comentarios, puedes contactarnos en:

Correo: 2201183@ucc.edu.ar, 2207479@ucc.edu.ar, 2205721@ucc.edu.ar
