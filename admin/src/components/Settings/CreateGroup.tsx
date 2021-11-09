import React, { useState } from 'react'
import * as UI from "@chakra-ui/react"
import { useForm, SubmitHandler } from "react-hook-form";
import axios from 'axios';
import { Toast } from '../template';

type GroupType = {
  group_name: string
  group_category: string
}

const CreateGroup = () => {
  const { isOpen, onOpen, onClose } = UI.useDisclosure()
  const toast = UI.useToast()

  const [isError, setIsError] = useState<boolean>(false)

  const { register, handleSubmit } = useForm<GroupType>();
  const onSubmit: SubmitHandler<GroupType> = data => {
    axios.post('settings', data)
      .then(() => {
        onClose()
        setIsError(false)
        
        toast({
          position: "bottom-right",
          render: () => (
            <Toast />
          )
        })
      })
      .catch((e) => {
        console.log(e.response)
        setIsError(true)
      })
  };

  return (
    <>
      <UI.Button onClick={onOpen}>新規グループ作成</UI.Button>

      <UI.Modal isOpen={isOpen} onClose={onClose}>
        <UI.ModalOverlay />
        <UI.ModalContent>
          <UI.ModalHeader>UI.Modal Title</UI.ModalHeader>
          <UI.ModalCloseButton />
            <form onSubmit={handleSubmit(onSubmit)}>
              <UI.ModalBody>
                <UI.Stack spacing={4}>
                  <UI.FormControl id="group_name">
                    <UI.FormLabel>Group Name</UI.FormLabel>
                    <UI.Input type="text" {...register("group_name")} />
                    <UI.FormHelperText>グループ名（例：ステータス関係）</UI.FormHelperText>
                  </UI.FormControl>
      
                  <UI.FormControl id="group_category">
                    <UI.FormLabel>Group Category</UI.FormLabel>
                    <UI.Input type="text" {...register("group_category")} />
                    <UI.FormHelperText>グループスラッグ ユニークな英語（例：parameter）</UI.FormHelperText>
                  </UI.FormControl>
                  {isError && <UI.Text color="red.600">作成に失敗しました。</UI.Text>}
                </UI.Stack>
              </UI.ModalBody>
              <UI.ModalFooter>
                <UI.Button
                  type="submit"
                  colorScheme="blue"
                  mr={3}
                >
                  作成
                </UI.Button>
                <UI.Button variant="ghost" onClick={onClose}>閉じる</UI.Button>
              </UI.ModalFooter>
            </form>
        </UI.ModalContent>
      </UI.Modal>
    </>
  )
}

export default CreateGroup
