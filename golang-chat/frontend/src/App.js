import { useEffect, useState } from "react";
import { connect, sendMsg } from "./api";
import "./App.css";
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput/ChatInput";
import Header from "./components/Header/Header";

function App() {
  const [chatLog, setChatLog] = useState([]);

  useEffect(() => {
    connect((msg) => {
      console.log("new message");
      setChatLog((prev) => [...prev, msg]);
    });
  }, []);

  const send = (event) => {
    if (event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  };

  return (
    <div className="app">
      <Header />
      <ChatHistory chatLog={chatLog} />
      <ChatInput send={send} />
    </div>
  );
}

export default App;
