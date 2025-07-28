#!/bin/bash
echo "🚀 Starting server..."
cd /home/ec2-user/chat-app/server
nohup node index.js > app.log 2>&1 &
