import React, { useContext, useState } from "react";
import Input from "../components/Input";
import Button from "../components/Button";
import { Link, useNavigate } from "react-router-dom";
import api_url from "../api/api";
import { authLogin } from "../api/auth";
import { AuthContext } from "../context/AuthContext";

function Login() {

  const navigate=useNavigate()

  const {login}=useContext(AuthContext)

  const [data,setData]=useState({
    email:"",
    password:""
  })

  async function handleClick(){

    console.log(data.email,data.password)

    try {

    const res=await authLogin(data)

    // const resdata=await res.json();

    // if(!res.ok){
    //   console.log(resdata.error)
    //   return
    // }

    login()

    navigate('/profile')
      
    } catch (error) {

      console.log(error)
      
    }

  }

  return (
    <div style={{display:"flex",alignItems:"center",justifyContent:"center",flexDirection:"column",padding:"8%",gap:"20px",border:"1px solid black",height:"35vh"}}>
      <h2 style={{marginBottom:"20%"}}>Login Page</h2>
      <Input Type="email" 
      onChange={(e)=>
        setData({
          ...data,
          email:e.target.value
        })
      }
      placeholder="Enter the email..." required={true}/>

      <Input Type="password" 
      onChange={(e)=>
        setData({
          ...data,
          password:e.target.value
        })
      }
      placeholder="Enter the password..." required={true}/>

      <Button Text="Login" onClick={handleClick} />
      <p style={{fontSize:"18px",fontWeight:"400"}}>Don't have an account? <Link to="/register">Register</Link></p>
    </div>
  );
}

export default Login;
