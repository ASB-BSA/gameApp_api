import React, { useEffect, useState } from 'react'
import {
  Flex,
  Box,
  FormControl,
  FormLabel,
  Input,
  Stack,
  Button,
  Heading,
  Text,
  useColorModeValue,
} from '@chakra-ui/react';
import { useForm } from 'react-hook-form';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

type AuthType = {
  user: string
  password: string
}

const Login = () => {
  const { register, handleSubmit } = useForm<AuthType>();
  const navigate = useNavigate()
  const [isError, setIsError] = useState<boolean>(false)
  
  const onSubmit = handleSubmit(data => {
    const auth = async () => {
      await axios.post('login', data)
        .then((e) => {
          console.log(e.data)
          setIsError(false)
          navigate('/', { replace: true })
        })
        .catch(() => setIsError(true))
    }
    auth()
  })

  return (
    <Flex
      minH={'100vh'}
      align={'center'}
      justify={'center'}
      bg={useColorModeValue('gray.50', 'gray.800')}>
      <Stack spacing={8} mx={'auto'} maxW={'lg'} py={12} px={6}>
        <Stack align={'center'}>
          <Heading fontSize={'4xl'}>Sign in to your account</Heading>
          <Text fontSize={'lg'} color={'gray.600'}>
            Boomin Games Admin Control Room
          </Text>
        </Stack>
        <form onSubmit={onSubmit}>
          <Box
            rounded={'lg'}
            bg={useColorModeValue('white', 'gray.700')}
            boxShadow={'lg'}
            p={8}>
            <Stack spacing={4}>
              <FormControl id="User">
                <FormLabel>User ID</FormLabel>
                <Input type="text" {...register("user")} />
              </FormControl>
              <FormControl id="password">
                <FormLabel>Password</FormLabel>
                <Input type="password" {...register("password")} />
              </FormControl>
              <Stack spacing={10}>
                <Stack
                  direction={{ base: 'column', sm: 'row' }}
                  align={'start'}
                  justify={'space-between'}>
                    {isError && <Text color={'red.600'}>Error</Text>}
                </Stack>
                <Button
                  type="submit"
                  bg={'blue.400'}
                  color={'white'}
                  _hover={{
                    bg: 'blue.500',
                  }}>
                  Sign in
                </Button>
              </Stack>
            </Stack>
          </Box>
        </form>
      </Stack>
    </Flex>
  )
}

export default Login
