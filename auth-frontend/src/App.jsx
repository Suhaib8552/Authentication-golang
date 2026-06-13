
//import './App.css'
import Input from './assets/components/Input'
import Login from './assets/pages/Login'
import Register from './assets/pages/Register'
import Profile from './assets/pages/Profile'
import {Routes,Route, Navigate} from 'react-router-dom'

function App() {


  return (
    <>
    <div style={{display:"flex",alignItems:"center",justifyContent:"center",height:"100vh",width:"100vw"}}>
    
    
      


      <Routes>
        <Route path='/' element={<Navigate to="/login" />} />
        <Route path='/login' element={<Login/>} />
        <Route path='/register' element={<Register/>} />
        <Route path='/profile' element={<Profile />} />
      </Routes>

      </div>

      </>
  )
}

export default App
