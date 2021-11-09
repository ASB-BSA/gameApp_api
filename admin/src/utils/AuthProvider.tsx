import React from 'react'
import { Navigate, Outlet } from "react-router-dom"
import { useRecoilValue } from 'recoil'
import IsRedirect from '../atoms/IsRedirect'

const AuthProvider = () => {
  const redirect = useRecoilValue(IsRedirect)
  if (redirect) {
    return <Navigate to="/login" replace />;
  }
  
  return (
    <>
      <Outlet />
    </>
  )
}

export default AuthProvider
