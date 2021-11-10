import React, { useEffect, useState } from 'react'
import * as UI from "@chakra-ui/react"
import { SubmitHandler, useForm } from 'react-hook-form';
import SettingState, { SettingType } from '../../atoms/SettingsState';
import axios from 'axios';
import { Toast } from '../template';
import { useSetRecoilState } from 'recoil';
import IsRedirect from '../../atoms/IsRedirect';

type Props = {
  onClose: any
  setting: SettingType
}

const SettingItemEdit = (props: Props) => {
  const { register, setValue, handleSubmit } = useForm<SettingType>();
  const [isError, setIsError] = useState<boolean>(false)
  const toast = UI.useToast()
  const setRedirect = useSetRecoilState(IsRedirect)
  const setSettings = useSetRecoilState(SettingState)

  useEffect(() => {
    setValue("setting_name", props.setting.setting_name)
    setValue("setting_label", props.setting.setting_label)
    setValue("setting_type", props.setting.setting_type)
    setValue("setting_value", props.setting.setting_value)
  }, [props.setting])

  const onSubmit: SubmitHandler<SettingType> = data => {
    axios.put(`settings/${props.setting.ID}`, data)
      .then(() => {
        props.onClose()
        setIsError(false)
        
        toast({
          position: "bottom-right",
          render: () => (
            <Toast />
          )
        })
        
        axios.get('settings')
        .then(e => setSettings(e.data))
        .catch(() => {
          setRedirect(true)
        })
      })
      .catch((e) => {
        console.log(e.response)
        setIsError(true)
      })
  };

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)}>
        <UI.ModalBody>
          <UI.Stack spacing={4}>
            <UI.FormControl id="group_name">
              <UI.FormLabel>設定名</UI.FormLabel>
              <UI.Input type="text" {...register("setting_name")} />
            </UI.FormControl>

            <UI.FormControl id="setting_label">
              <UI.FormLabel>カラム</UI.FormLabel>
              <UI.Input type="text" {...register("setting_label")} />
            </UI.FormControl>

            <UI.FormControl id="setting_type">
              <UI.FormLabel>型</UI.FormLabel>
                <UI.Select {...register("setting_type")} >
                  <option value="string">String</option>
                  <option value="int">Number</option>
                  <option value="boolean">Boolean</option>
                </UI.Select>
            </UI.FormControl>

            <UI.FormControl id="setting_value">
              <UI.FormLabel>値</UI.FormLabel>
              <UI.Input type="text" {...register("setting_value")} />
            </UI.FormControl>
            {isError && <UI.Text color="red.600">更新に失敗しました。</UI.Text>}
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
          <UI.Button variant="ghost" onClick={props.onClose}>閉じる</UI.Button>
        </UI.ModalFooter>
      </form>
    </>
  )
}

export default SettingItemEdit
