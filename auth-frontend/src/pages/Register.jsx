import React, { useState } from 'react'
import Input from '../components/Input'
import Button from '../components/Button'
import { Link, useNavigate } from 'react-router-dom'
import api_url from '../api/api';
import { authRegister } from '../api/auth';


function Register() {

  const navigate=useNavigate();

  const [data,setData]=useState({
    username:"",
    email:"",
    password:""
  })

  async function handleClick(){

    console.log(data.username,data.email,data.password)

    try {


      const res=await authRegister(data)

      if(!res.ok){
        console.log(data.error);
        return;
      }
      alert("Registration successful");
      navigate('/login')

    } catch (error) {

      console.log(error)
      
    }

  }

  return (
    <div style={{display:"flex",alignItems:"center",justifyContent:"center",flexDirection:"column",padding:"8%",gap:"20px",border:"1px solid black",height:"35vh"}}>
      
      <h2 style={{marginBottom:"20%"}}>Register Page</h2>

      <Input Type="text"
      onChange={(e)=>
        setData({
          ...data,username:e.target.value,
        })
      }
       placeholder="Enter the username..." 
       required={true} />

      <Input Type="email" placeholder="Enter the email..."
      onChange={(e)=>
        setData({
          ...data,email:e.target.value,
        })
      }
      required={true}/>

      <Input Type="password" placeholder="Enter the password..." 
      onChange={(e)=>
        setData({
          ...data,password:e.target.value,
        })
      }
      required={true}/>

      <Button Text="Register"  onClick={handleClick} />

      <p style={{fontSize:"18px",fontWeight:"400"}}>Already have an account? <Link to="/login">Login</Link></p>
    </div>
  )
}

export default Register