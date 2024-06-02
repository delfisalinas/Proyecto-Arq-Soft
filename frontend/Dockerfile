# Establecer la imagen base
FROM node:14-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el package.json y package-lock.json
COPY package*.json ./

# Instalar dependencias
RUN npm install

# Copiar el resto del código del frontend
COPY . .

# Compilar la aplicación
RUN npm run build

# Instalar un servidor web para servir la aplicación compilada
RUN npm install -g serve

# Exponer el puerto en el que el servidor escucha
EXPOSE 5000

# Comando para ejecutar el servidor web
CMD ["serve", "-s", "build"]
