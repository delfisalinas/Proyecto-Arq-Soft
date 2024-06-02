# Establecer la imagen base
FROM golang:1.16-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del backend al directorio de trabajo
COPY . .

# Descargar dependencias e instalar
RUN go mod download

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto en el que la aplicación escucha
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
