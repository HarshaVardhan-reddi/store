# Store API

A simple Go-based multi-tenant application that exposes APIs to manage stores, users, and their products. Built as a learning project to improve Go skills and experiment with different ORMs.

## Tech Stack

- **Language:** Go  
- **Database:** MySQL  
- **Runner:** Docker  

## Overview

This application provides REST APIs for:

- Managing **Stores**
- Managing **Products** under each store
- Managing **Users** associated with stores
- Handling **Authentication & Authorization** using JWT

Each store acts as a separate tenant, allowing isolated user and product management within the same system.

## Features

- Multi-tenant architecture (multiple stores → multiple products & users)  
- CRUD operations for stores  
- CRUD operations for products  
- CRUD operations for users  
- **JWT-based Authentication & Authorization**  
  - Secure login for users  
  - Store-level authorization checks  
  - Protected routes for product and store operations  
- Clean and modular project structure  
- ORM-friendly design

## Purpose

This project is created to:

- Enhance Go development skills  
- Explore and compare different Go ORMs  
- Practice secure API development  
- Learn JWT-based auth flows  
- Build Dockerized service architecture
