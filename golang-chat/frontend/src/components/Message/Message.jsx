import React from 'react'
import { useState } from 'react'
import './Message.css'

const Message = (props) => {
  const message = useState(JSON.parse(props.message))[0];
  return (
    <div className='Message'>{message.body}</div>
  )
}

export default Message