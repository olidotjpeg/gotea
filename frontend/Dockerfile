# Stage 1: Base image with dependencies
FROM node:20.13.1-bookworm AS base

# Set the working directory
WORKDIR /frontend

# Copy package.json and package-lock.json to the working directory
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code to the working directory
COPY . .

# Stage 2: Development image
FROM base AS devfrontend

# Set the entrypoint to list the contents of the root directory
ENTRYPOINT ["npm", "run", "dev"]