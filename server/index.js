const express = require('express');
const http = require('http');
const { Server } = require('socket.io');
const cors = require('cors');

const app = express();
const server = http.createServer(app);
const io = new Server(server, {
  cors: { origin: '*' }
});

app.use(cors());
app.get('/', (req, res) => res.send('Chat server running!'));

io.on('connection', (socket) => {
  console.log('User connected:', socket.id);

  socket.on('chat', (msg) => {
    io.emit('chat', msg);
  });

  socket.on('disconnect', () => {
    console.log('User disconnected:', socket.id);
  });
});

server.listen(5000, () => console.log('Server running on port 5000'));
