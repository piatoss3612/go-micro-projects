import React from 'react'
import Message from '../Message/Message'
import './ChatHistory.css'

const ChatHistory = (props) => {
  const messages = props.chatLog.map(msg => <Message key={msg.timeStamp} message={msg.data} />);

  return (
    <div className='ChatHistory'>
      <h2>Chat History</h2>
      {messages}
    </div>
  )
}

export default ChatHistory