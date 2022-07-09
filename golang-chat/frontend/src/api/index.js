var socket = new WebSocket("ws://localhost:9000/ws");

export const connect = (cb) => {
  console.log("connecting");

  socket.onopen = () => {
    console.log("successfully connected");
  };

  socket.onmessage = (msg) => {
    console.log("message from websocket:", msg);
    cb(msg);
  };

  socket.onclose = (event) => {
    console.log("socket connection closed:", event);
  };

  socket.onerror = (error) => {
    console.log("socket error:", error);
  };
};

export const sendMsg = (msg) => {
  console.log("sending message:", msg);
  socket.send(msg);
};
