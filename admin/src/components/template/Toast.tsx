import React from 'react'
import { RiCheckboxCircleFill } from 'react-icons/ri'
import * as UI from "@chakra-ui/react"

const Toast = () => {
  return (
    <UI.Flex
      maxW="sm"
      w="full"
      mx="auto"
      bg={UI.useColorModeValue("white", "gray.800")}
      shadow="md"
      rounded="lg"
      overflow="hidden"
    >
      <UI.Flex justifyContent="center" alignItems="center" w={12} bg="green.500">
        <RiCheckboxCircleFill color="white" />
      </UI.Flex>

      <UI.Box mx={-3} py={2} px={4}>
        <UI.Box mx={3}>
          <UI.chakra.span
            color={UI.useColorModeValue("green.500", "green.400")}
            fontWeight="bold"
          >
            Success
          </UI.chakra.span>
          <UI.chakra.p
            color={UI.useColorModeValue("gray.600", "gray.200")}
            fontSize="sm"
          >
            The process has been completed successfully.
          </UI.chakra.p>
        </UI.Box>
      </UI.Box>
    </UI.Flex>
  )
}

export default Toast
