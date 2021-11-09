import React, { useEffect, useState } from 'react'
import * as UI from "@chakra-ui/react"
import axios from 'axios'
import * as Icon from "react-icons/ri";
import AdministratorState from '../atoms/AdministratorState'
import { useRecoilState, useSetRecoilState } from 'recoil'
import IsRedirect from '../atoms/IsRedirect';
import { useForm } from 'react-hook-form';
import { Toast } from './template';

type AuthType = {
  user: string
  password: string
  password_confirm: string
}

const Administrator = () => {
  const toast = UI.useToast()
  const [admin, setAdmin] = useRecoilState(AdministratorState)
  const setRedirect = useSetRecoilState(IsRedirect)

  useEffect(() => {
    axios.get('/admins')
      .then(e => {
        setAdmin(e.data)
      })
      .catch(() => {
        setRedirect(true)
      })
  }, [])

  const { isOpen, onOpen, onClose } = UI.useDisclosure()

  const { register, handleSubmit } = useForm<AuthType>()
  const [isError, setIsError] = useState<boolean>(false)
  
  const onSubmit = handleSubmit(data => {
    const auth = async () => {
      await axios.post('register', data)
        .then(() => {
          onClose()
          setIsError(false)
          
          toast({
            position: "bottom-right",
            render: () => (
              <Toast />
            )
          })

          axios.get('/admins')
            .then(e => {
              setAdmin(e.data)
            })
            .catch(() => {
              setRedirect(true)
            })
        })
        .catch((e) => {
          console.log(e.response)
          setIsError(true)
        })
    }
    auth()
  })

  return (
    <>
      <UI.Heading
        size="md"
      >管理者一覧</UI.Heading>
      
      <UI.Flex>
        <UI.Spacer />
        <UI.Button onClick={onOpen}>管理者アカウント追加</UI.Button>
      </UI.Flex>

      <UI.Table
        mt={4}
        variant="striped"
      >
        <UI.Thead>
          <UI.Tr>
            <UI.Th>ユーザー名</UI.Th>
            <UI.Th>作成日</UI.Th>
            <UI.Th>操作</UI.Th>
          </UI.Tr>
        </UI.Thead>
        <UI.Tbody>
          {admin.map(e => {
            return (
              <UI.Tr key={e.ID}>
                <UI.Td>{e.user}</UI.Td>
                <UI.Td>{e.CreatedAt}</UI.Td>
                <UI.Td>
                  <UI.Stack
                    direction="row"
                    spacing={2}
                  >
                    <UI.IconButton
                      colorScheme="blue"
                      aria-label="Edit"
                      icon={<Icon.RiEdit2Fill />}
                    />
                    <UI.IconButton
                      colorScheme="red"
                      aria-label="Delete"
                      icon={<Icon.RiDeleteBin6Fill />}
                    />
                  </UI.Stack>
                </UI.Td>
              </UI.Tr>
            )
          })}
        </UI.Tbody>
      </UI.Table>

      <UI.Modal isOpen={isOpen} onClose={onClose}>
        <UI.ModalOverlay />
        <UI.ModalContent>
          <UI.ModalHeader>UI.Modal Title</UI.ModalHeader>
          <UI.ModalCloseButton />
          <UI.ModalBody>
            <form onSubmit={onSubmit}>
              <UI.Box
                p={8}>
                <UI.Stack spacing={4}>
                  <UI.FormControl id="User">
                    <UI.FormLabel>User ID</UI.FormLabel>
                    <UI.Input type="text" {...register("user")} />
                  </UI.FormControl>
                  <UI.FormControl id="password">
                    <UI.FormLabel>Password</UI.FormLabel>
                    <UI.Input type="password" {...register("password")} />
                  </UI.FormControl>
                  <UI.FormControl id="password_comfirm">
                    <UI.FormLabel>Password_comfirm</UI.FormLabel>
                    <UI.Input type="password" {...register("password_confirm")} />
                  </UI.FormControl>
                  <UI.Stack spacing={10}>
                    <UI.Stack
                      direction={{ base: 'column', sm: 'row' }}
                      align={'start'}
                      justify={'space-between'}>
                        {isError && <UI.Text color={'red.600'}>Error</UI.Text>}
                    </UI.Stack>
                    <UI.Button
                      type="submit"
                      bg={'blue.400'}
                      color={'white'}
                      _hover={{
                        bg: 'blue.500',
                      }}>
                      Create new account
                    </UI.Button>
                  </UI.Stack>
                </UI.Stack>
              </UI.Box>
            </form>
          </UI.ModalBody>

          <UI.ModalFooter>
            <UI.Button variant="ghost" onClick={onClose}>Cancel</UI.Button>
          </UI.ModalFooter>
        </UI.ModalContent>
      </UI.Modal>
    </>
  )
}

export default Administrator
