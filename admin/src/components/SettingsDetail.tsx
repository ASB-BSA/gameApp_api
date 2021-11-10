import React, { useEffect, useState } from 'react'
import * as UI from "@chakra-ui/react"
import axios from 'axios'
import { useRecoilState, useSetRecoilState } from 'recoil'
import IsRedirect from '../atoms/IsRedirect'
import SettingState, { SettingGroup, SettingType } from '../atoms/SettingsState'
import { useParams } from 'react-router-dom'
import { useForm } from 'react-hook-form'
import * as Icon from "react-icons/ri"
import SettingItemEdit from './Settings/SettingItemEdit'

const SettingsDetail = () => {
  let { id } = useParams();
  const [settings, setSettings] = useRecoilState(SettingState)
  const setRedirect = useSetRecoilState(IsRedirect)

  const { isOpen, onOpen, onClose } = UI.useDisclosure()
  const { register, handleSubmit, reset } = useForm<SettingType>({
    defaultValues: {
      "group_id": 0,
      "setting_label": "",
      "setting_name": "",
      "setting_type": "string",
      "setting_value": "",
    }
  });
  
  const [item, setItem] = useState<SettingGroup|undefined>({
    ID: 0,
    group_name: "",
    group_category: "",
    settings: [],
  })

  useEffect(() => {
    if (settings.length > 0) {
      setItem(settings.find(f => f.group_category === id))
    }
  }, [settings])

  useEffect(() => {
    axios.get('settings')
      .then(e => setSettings(e.data))
      .catch(() => {
        setRedirect(true)
      })
  }, [])
  
  const onSubmit = handleSubmit(data => {
    data.group_id = String(item?.ID)
    axios.post(`settings/${data.group_id}`, data)
      .then(() => {
        reset()
        axios.get('settings')
          .then(e => setSettings(e.data))
          .catch(() => {
            setRedirect(true)
          })
      })
      .catch(e => {
        console.log(e.response)
      })
  })

  const [edit, setEdit] = useState<SettingType>({
    "ID": 0,
    "group_id": 0,
    "setting_label": "",
    "setting_name": "",
    "setting_type": "string",
    "setting_value": "",
  })

  return (
    <>
      <UI.Heading
        size="md"
      >{item?.group_name}設定</UI.Heading>

      <form onSubmit={onSubmit}>
        <UI.Table
          variant="striped"
          mt={4}
        >
          <UI.Thead>
            <UI.Tr>
              <UI.Th>設定名</UI.Th>
              <UI.Th>カラム名(英語)</UI.Th>
              <UI.Th>型</UI.Th>
              <UI.Th>値</UI.Th>
              <UI.Th>操作</UI.Th>
            </UI.Tr>
          </UI.Thead>
          <UI.Tbody>
            {item?.settings ?
              item?.settings.map((e, i) => {
                return (
                  <UI.Tr key={i}>
                    <UI.Td>{e.setting_name}</UI.Td>
                    <UI.Td>{e.setting_label}</UI.Td>
                    <UI.Td>{e.setting_type}</UI.Td>
                    <UI.Td>{e.setting_value}</UI.Td>
                    <UI.Td>
                      <UI.Stack
                        direction="row"
                        spacing={2}
                      >
                        <UI.IconButton
                          colorScheme="blue"
                          aria-label="Edit"
                          onClick={() => {
                            setEdit(e)
                            onOpen()
                          }}
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
              })
            : <UI.Tr><UI.Td colSpan={5}>なにもありません</UI.Td></UI.Tr>}
            <UI.Tr>
              <UI.Td><UI.Input type="text" {...register("setting_name")} /></UI.Td>
              <UI.Td><UI.Input type="text" {...register("setting_label")} /></UI.Td>
              <UI.Td>
                <UI.Select {...register("setting_type")} >
                  <option value="string">String</option>
                  <option value="int">Number</option>
                  <option value="boolean">Boolean</option>
                </UI.Select>
              </UI.Td>
              <UI.Td><UI.Input type="text" {...register("setting_value")} /></UI.Td>
              <UI.Td>
                <UI.Button type="submit">追加</UI.Button>
              </UI.Td>
            </UI.Tr>
          </UI.Tbody>
        </UI.Table>
      </form>
      
      <UI.Modal isOpen={isOpen} onClose={onClose}>
        <UI.ModalOverlay />
        <UI.ModalContent>
          <UI.ModalHeader>UI.Modal Title</UI.ModalHeader>
          <SettingItemEdit setting={edit} onClose={onClose} />
          <UI.ModalCloseButton />
        </UI.ModalContent>
      </UI.Modal>
    </>
  )
}

export default SettingsDetail
