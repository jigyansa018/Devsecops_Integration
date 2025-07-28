#!/bin/bash
echo "📦 Installing dependencies..."
cd /home/ec2-user/chat-app

# Backend
cd server
npm install

# Frontend
cd ../src
npm install
npm run build
