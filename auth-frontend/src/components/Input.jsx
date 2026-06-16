import React from 'react'

function Input({Type,placeholder,required,onChange}) {
  return (
    <input style={{padding:"2%",fontSize:"18px",borderRadius:"2%"}}
    onChange={onChange}
     type={Type} placeholder={placeholder} required={required}/>
  )
}

export default Input