import React, { useContext, useEffect, useState } from 'react'
import Button from '../components/Button'
import api_url from '../api/api'
import { useNavigate } from 'react-router-dom'
import { authLogout, authProfile } from '../api/auth'
import { AuthContext } from '../context/AuthContext'

function Profile() {

  const [profdata,setProfdata]=useState(null)

  const navigate=useNavigate()

  const {logout}=useContext(AuthContext)

  useEffect(()=>{
    handleFetch()
  },[])

  async function handleFetch(){

    try {


      

      const res=await authProfile()

      const resdata=await res.json()

      console.log(resdata)

      if(!res.ok){
        alert(resdata.error)
        navigate('/login')
        return
      }

      setProfdata(resdata)
      
    } catch (error) {

      console.log(error)
      
    }

  }


  async function handleLogout(){

    

    await authLogout()

    logout()

    navigate('/login')

  }



  return (
    <>
    <div style={{width:"50vw",maxHeight:"50vh",display:"flex",alignItems:"center",justifyContent:"center",flexDirection:"column"}}>
    <h2 style={{marginBottom:"5%"}}>Profile Page</h2>

    <div style={{fontSize:"18px",marginBottom:"5%"}}>
      {profdata && (
    <>
      <p>ID: {profdata.id}</p>
      <p>Username: {profdata.username}</p>
      <p>Email: {profdata.email}</p>
    </>
  )}
  </div>

    <Button Text="Logout" onClick={handleLogout} />

    </div>

    </>
  )
}

export default Profile