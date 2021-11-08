import React from 'react'
import { Outlet } from 'react-router'
import * as UI from "@chakra-ui/react"
import RouterConfig from '../../config/RouterConfig'
import { NavLink } from 'react-router-dom'

const Layout = () => {
  return (
    <UI.Flex>
      <UI.Box
        w={64}
        h="100vh"
        p={8}
        bg="gray.700"
      >
        <UI.Stack spacing={4}>
          {RouterConfig.map(e => {
            return (
            <>{e.path !== "/login" && e.children?.map(el => {
              console.log(el)
              return (
                <>{el.children?.map(elm => (
                  <>{elm.icon && 
                    <NavLink
                      to={`${elm.path}`}
                    >
                      <UI.Flex
                        alignItems="center"
                      >
                        <UI.Box
                          mr={3}
                          fontSize="22px"
                          color="white"
                        >
                          {elm.icon}
                        </UI.Box>
                        <UI.Text
                          fontSize="13px"
                          color="white"
                        >{elm.name}</UI.Text>
                      </UI.Flex>
                    </NavLink>
                    }
                  </>))
                }</>
              )})
            }</>
          )})}
        </UI.Stack>
      </UI.Box>
      <UI.Box
        flex={1}
        p={8}
      >
        <Outlet />
      </UI.Box>
    </UI.Flex>
  )
}

export default Layout
