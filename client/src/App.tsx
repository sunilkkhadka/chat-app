import React, { useEffect, useState } from "react";

type Message = {
  sender_id: string;
  receiver_id: string;
  message: string;
  chat_type: string;
};

const App: React.FC = () => {
  const randomId = `user-${Date.now()}-${Math.floor(Math.random() * 10000)}`;

  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [status, setStatus] = useState("");
  const [messages, setMessages] = useState<Message[]>([]);
  const [input, setInput] = useState<Message>({
    sender_id: randomId,
    receiver_id: "",
    message: "",
    chat_type: "private",
  });

  console.log("Messages", messages);

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8000/");
    ws.onopen = () => {
      console.log("Connected Successfully");
      ws.send(JSON.stringify({ sender_id: randomId }));
    };

    setSocket(ws);

    ws.onerror = (event) => {
      console.log("Socket error:", error);
    };

    ws.onmessage = (event: MessageEvent) => {
      const response = JSON.parse(event.data);
      console.log(response);
      if ("status" in response) {
        setStatus(response.status);
      } else {
        const response: Message = JSON.parse(event.data);
        setMessages((prev) => [...prev, response]);
      }
    };

    return () => {
      if (ws.readyState === WebSocket.OPEN) {
        ws.close();
      }
    };
  }, []);

  const sendMessage = () => {
    console.log("send clicked");
    if (socket && input) {
      console.log("inpu message", input);
      socket.send(JSON.stringify(input));
      setInput((prev) => {
        return { ...prev, text: "" };
      });
    }
  };

  return (
    <div>
      <h1>Chat</h1>
      <strong style={{ color: "green", marginBottom: "10px" }}>{status}</strong>
      <div>
        {messages.map((msg, index) => (
          <p key={index}>
            <strong>{msg.sender_id}: &nbsp;</strong>
            {msg.message}
          </p>
        ))}
      </div>
      <div className="input__area">
        <input
          type="text"
          placeholder="sender id"
          value={input.sender_id}
          onChange={(e) =>
            setInput((prev) => {
              return {
                ...prev,
                sender_id: e.target.value,
              };
            })
          }
        />
        <input
          value={input.receiver_id}
          placeholder="receiver id"
          onChange={(e) =>
            setInput((prev) => {
              return {
                ...prev,
                receiver_id: e.target.value,
              };
            })
          }
        />
        <input
          value={input.message}
          placeholder="text"
          onChange={(e) =>
            setInput((prev) => {
              return {
                ...prev,
                message: e.target.value,
              };
            })
          }
        />
      </div>
      <button onClick={sendMessage}>Send</button>
    </div>
  );
};

export default App;
