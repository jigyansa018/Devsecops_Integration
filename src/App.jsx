import React, { useState, useEffect } from 'react';
import socket from './socket';
import ChatBox from './ChatBox';

function App() {
  const [messages, setMessages] = useState([]);
  const [text, setText] = useState('');

  useEffect(() => {
    socket.on('chat', (msg) => {
      setMessages((prev) => [...prev, msg]);
    });

    return () => socket.off('chat');
  }, []);

  const sendMessage = (e) => {
    e.preventDefault();
    if (text.trim()) {
      socket.emit('chat', text);
      setText('');
    }
  };

  return (
    <div className="app">
      <h1>💬 Chat App</h1>
      <ChatBox messages={messages} />
      <form onSubmit={sendMessage}>
        <input
          value={text}
          onChange={(e) => setText(e.target.value)}
          placeholder="Type a message"
        />
        <button type="submit">Send</button>
      </form>
    </div>
  );
}

export default App;
