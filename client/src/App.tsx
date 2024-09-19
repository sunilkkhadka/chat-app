import React, { useEffect, useState } from "react";

// Define the WebSocket type as WebSocket | null because it can be initially null
const App: React.FC = () => {
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [messages, setMessages] = useState<string[]>([]);
  const [input, setInput] = useState<string>("");

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8000/");
    setSocket(ws);

    ws.onmessage = (event: MessageEvent) => {
      setMessages((prev) => [...prev, event.data]);
    };

    // Cleanup the WebSocket connection when the component unmounts
    return () => {
      ws.close();
    };
  }, []);

  const sendMessage = () => {
    if (socket && input) {
      socket.send(input);
      setInput(""); // Clear the input after sending
    }
  };

  return (
    <div>
      <h1>Chat</h1>
      <div>
        {messages.map((msg, index) => (
          <p key={index}>{msg}</p>
        ))}
      </div>
      <input value={input} onChange={(e) => setInput(e.target.value)} />
      <button onClick={sendMessage}>Send</button>
    </div>
  );
};

export default App;
