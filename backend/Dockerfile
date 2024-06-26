# Utilizar la imagen oficial de Go con una versión específica
FROM golang:1.22-alpine as builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el archivo go.mod y go.sum al directorio de trabajo
COPY go.mod go.sum ./

# Descargar las dependencias de Go
RUN go mod download

# Copiar el resto del código del backend al contenedor
COPY . .

# Compilar el binario del backend
RUN go build -o main .

# Imagen base para la aplicación compilada
FROM alpine:latest

# Establecer el directorio de trabajo para la aplicación final
WORKDIR /root/

# Copiar la aplicación compilada desde la imagen de compilación
COPY --from=builder /app/main .

# Exponer el puerto en el que escucha tu aplicación
EXPOSE 8080

# Definir el comando de inicio del contenedor
CMD ["./main"]
