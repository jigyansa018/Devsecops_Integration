import React from 'react';

const ChatBox = ({ messages }) => (
  <div className="chat-box">
    {messages.map((msg, idx) => (
      <div key={idx} className="chat-message">{msg}</div>
    ))}
  </div>
);

export default ChatBox;
