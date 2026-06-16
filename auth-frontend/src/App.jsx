
//import './App.css'
import Input from './components/Input'
import Login from './pages/Login'
import Register from './pages/Register'
import Profile from './pages/Profile'
import {Routes,Route, Navigate} from 'react-router-dom'
import Protectedroute from './components/Protectedroute'

function App() {


  return (
    <>
    <div style={{display:"flex",alignItems:"center",justifyContent:"center",height:"100vh",width:"100vw"}}>
    
    
      


      <Routes>
        <Route path='/' element={<Navigate to="/login" />} />
        <Route path='/login' element={<Login/>} />
        <Route path='/register' element={<Register/>} />
        <Route path='/profile' element={
          <Protectedroute>
          <Profile />
          </Protectedroute>
          } />
      </Routes>

      </div>

      </>
  )
}

export default App
