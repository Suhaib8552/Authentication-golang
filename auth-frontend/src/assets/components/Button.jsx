import React from 'react'

function Button({Text,onClick}) {
  return (
    <>
    <button style={{padding:"2% 4%",fontSize:"16px"}} onClick={onClick}>{Text}</button>
    </>
  )
}

export default Button