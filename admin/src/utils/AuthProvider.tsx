import React from 'react'
import { Outlet } from 'react-router'

const AuthProvider = () => {
  return (
    <>
      <Outlet />
    </>
  )
}

export default AuthProvider
