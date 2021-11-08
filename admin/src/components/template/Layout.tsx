import React from 'react'
import { Outlet } from 'react-router'
import * as UI from "@chakra-ui/react"
import RouterConfig from '../../config/RouterConfig'
import { NavLink } from 'react-router-dom'

const Layout = () => {
  return (
    <UI.Flex>
      <UI.Box
        h="100vh"
        py={16}
        px={4}
        bg="gray.700"
      >
        <UI.Stack spacing={8}>
          {RouterConfig.map(e => {
            return (
            <>{e.path !== "/login" && e.children?.map(el => {
              return (
                <>{el.children?.map(elm => (
                  <>{elm.icon && 
                    <NavLink
                      to={`${elm.path}`}
                    >
                      <UI.Flex
                        alignItems="center"
                        justifyContent="center"
                      >
                        <UI.Box
                          mr={3}
                          fontSize="22px"
                          color="white"
                        >
                          {elm.icon}
                        </UI.Box>
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
