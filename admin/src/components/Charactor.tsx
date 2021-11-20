import React, { useEffect } from 'react'
import * as UI from "@chakra-ui/react"
import axios from 'axios'
import { useRecoilState, useSetRecoilState } from 'recoil'
import IsRedirect from '../atoms/IsRedirect'
import CharactorState, { CharactorType } from '../atoms/CharactorState'
import { useForm } from 'react-hook-form'

const Charactor = () => {
  const [charactor, setCharactor] = useRecoilState(CharactorState)
  const setRedirect = useSetRecoilState(IsRedirect)
  const { isOpen, onOpen, onClose } = UI.useDisclosure()
  const { register, handleSubmit, setValue } = useForm<CharactorType>()

  useEffect(() => {
    axios.get('/character')
      .then(e => {
        setCharactor(e.data)
      })
      .catch(() => {
        setRedirect(true)
      })
  }, [])

  const onSubmit = handleSubmit(data => {
    axios.put(`character/${data.ID}`, data)
      .then(() => {
        axios.get('/character')
        .then(e => {
          setCharactor(e.data)
          onClose()
        })
        .catch(() => {
          setRedirect(true)
        })
      })
      .catch(e => console.log(e.response))
  })

  return (
    <>
      <UI.Heading
        size="md"
      >管理者一覧</UI.Heading>

      <UI.Table
        mt={4}
        variant="striped"
      >
        <UI.Thead>
          <UI.Tr>
            <UI.Th>イメージ</UI.Th>
            <UI.Th>アイコン</UI.Th>
            <UI.Th>キャラクター名</UI.Th>
            <UI.Th>英語表記</UI.Th>
          </UI.Tr>
        </UI.Thead>
        <UI.Tbody>
          {charactor.map(e => {
            return (
              <UI.Tr
                key={e.ID}
                onClick={() => {
                  setValue("ID", `${e.ID}`)
                  setValue("english", e.english)
                  setValue("name", e.name)
                  setValue("img", e.img)
                  setValue("icon", e.icon)
                  onOpen()
                }}
              >
                <UI.Td>
                  <UI.Box
                    w={16}
                    h={16}
                    mx="auto"
                  >
                    <UI.Image
                      src={`https://api.game-boomin.net/api/v1/image/${e.img}`}
                      w={16}
                      h={16}
                      objectFit="cover"
                      objectPosition="top center"
                    />
                  </UI.Box>
                </UI.Td>
                <UI.Td>
                  <UI.Box
                    w={16}
                    h={16}
                    mx="auto"
                  >
                    <UI.Image src={`https://api.game-boomin.net/api/v1/image/${e.icon}`} 
                      w={16}
                      h={16}
                      objectFit="cover"
                      objectPosition="top center"
                    />
                  </UI.Box>
                </UI.Td>
                <UI.Td>
                  {e.name}
                </UI.Td>
                <UI.Td>
                  {e.english}
                </UI.Td>
              </UI.Tr>
            )
          })}
        </UI.Tbody>
      </UI.Table>

      <UI.Modal closeOnOverlayClick={false} isOpen={isOpen} onClose={onClose}>
        <UI.ModalOverlay />
        <UI.ModalContent>
          <UI.ModalHeader>キャラクター編集</UI.ModalHeader>
          <UI.ModalCloseButton />
          <UI.ModalBody>
            <form onSubmit={onSubmit}>
              <UI.Stack
                spacing={4}
              >
                <UI.FormControl>
                  <UI.FormLabel>画像名</UI.FormLabel>
                  <UI.Input type="text" {...register("img")} />
                </UI.FormControl>
                
                <UI.FormControl>
                  <UI.FormLabel>アイコン名</UI.FormLabel>
                  <UI.Input type="text" {...register("icon")} />
                </UI.FormControl>

                <UI.FormControl>
                  <UI.FormLabel>名前</UI.FormLabel>
                  <UI.Input type="text" {...register("name")} />
                </UI.FormControl>

                <UI.FormControl>
                  <UI.FormLabel>英語名</UI.FormLabel>
                  <UI.Input type="text" {...register("english")} />
                </UI.FormControl>
              </UI.Stack>
              <UI.Stack mt={4} py={4} direction="row" spacing={2}>
                <UI.Button mr={3} onClick={onClose}>Close</UI.Button>
                <UI.Button type="submit" colorScheme="blue">Save</UI.Button>
              </UI.Stack>
            </form>
          </UI.ModalBody>
        </UI.ModalContent>
      </UI.Modal>
    </>
  )
}

export default Charactor
