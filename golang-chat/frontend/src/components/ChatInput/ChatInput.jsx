import React from 'react'
import './ChatInput.css'

const ChatInput = (props) => {
  return (
    <div className='ChatInput'>
      <input onKeyDown={props.send} placeholder="Type a message... Hit enter to send" />
    </div>
  )
}

export default ChatInput