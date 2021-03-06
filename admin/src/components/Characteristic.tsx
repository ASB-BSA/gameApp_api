import React, { useEffect } from 'react'
import * as UI from "@chakra-ui/react"
import * as Icon from "react-icons/ri"
import { useRecoilState, useSetRecoilState } from 'recoil';
import CharacteristicState, { CharacteristicStateType } from '../atoms/CharacteristicState';
import { useForm } from 'react-hook-form';
import ConvertToText from '../utils/Characteristic';
import IsRedirect from '../atoms/IsRedirect';
import axios from 'axios';

const Characteristic = () => {
  const [characteristics, setCharacteristics] = useRecoilState(CharacteristicState);
  const setRedirect = useSetRecoilState(IsRedirect)

  useEffect(() => {
    axios.get('characteristic')
      .then(e => setCharacteristics(e.data))
      .catch(() => {
        setRedirect(true)
      })
  }, [])

  const { register, setValue, handleSubmit, reset } = useForm<CharacteristicStateType>({
    defaultValues: {
      name: "",
      timing: 'start',
      conditionsParameter: 'hp',
      conditionsValue: 0,
      conditionsExpression: '>',
      toWhom: 'myself',
      parameter: 'attack',
      happen: '+',
      howMuch: 0,
    }
  });

  const onSubmit = handleSubmit(data => {
    data.conditionsValue = String(data.conditionsValue)
    data.howMuch = String(data.howMuch)

    setCharacteristics([...characteristics, data])
    axios.post(`characteristic`, data)
      .then(() => {
        reset()
        axios.get('characteristic')
          .then(e => setCharacteristics(e.data))
          .catch(() => {
            setRedirect(true)
          })
      })
      .catch(e => {
        console.log(e.response)
      })
  })

  return (
    <>
      <UI.Heading
        size="md"
      >特技管理</UI.Heading>
      <form onSubmit={onSubmit}>
        <UI.Table
          variant="striped"
          mt={4}
        >
          <UI.Thead>
            <UI.Tr>
              <UI.Th>技名</UI.Th>
              <UI.Th colSpan={8}>効果</UI.Th>
              <UI.Th>操作</UI.Th>
            </UI.Tr>
          </UI.Thead>
          <UI.Tbody>
            {
              characteristics.map((e, i) => {
                return (
                  <UI.Tr key={i}>
                    <UI.Td>{e.name}</UI.Td>
                    <UI.Td colSpan={8}>{ConvertToText(e)}</UI.Td>
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
              })
            }
            <UI.Tr>
              <UI.Td><UI.Input type="text" {...register("name")} /></UI.Td>
              <UI.Td>
                <UI.Select {...register("timing")} >
                  <option value="start">ターン開始時</option>
                  <option value="attack">攻撃した時</option>
                  <option value="damage">ダメージを受けたとき</option>
                  <option value="end">ターン終了時</option>
                </UI.Select>
              </UI.Td>
              <UI.Td>
                <UI.Select {...register("conditionsParameter")} >
                  <option value="hp">体力</option>
                  <option value="damage">ダメージ</option>
                </UI.Select>
              </UI.Td>
              <UI.Td>
                <UI.Flex alignItems="flex-end">
                  <UI.NumberInput onChange={e => setValue("conditionsValue", Number(e))}>
                    <UI.NumberInputField />
                  </UI.NumberInput>
                  <UI.Text>%</UI.Text>
                </UI.Flex>
              </UI.Td>
              <UI.Td>
                <UI.Select {...register("conditionsExpression")} >
                  <option value=">">以下</option>
                  <option value="<">以上</option>
                  <option value="=">に</option>
                </UI.Select>
              </UI.Td>
              <UI.Td>
                <UI.Select {...register("toWhom")} >
                  <option value="myself">自分に</option>
                  <option value="all_allies">味方全体に</option>
                  <option value="random_allies">ランダムで味方単体に</option>
                  <option value="all_enemies">相手全体に</option>
                  <option value="ramdom_enemies">ランダムで相手単体に</option>
                </UI.Select>
              </UI.Td>
              <UI.Td>
                <UI.Select {...register("parameter")} >
                  <option value="attack">攻撃力</option>
                  <option value="defence">防御力</option>
                  <option value="critical_rate">クリティカル率</option>
                  <option value="agility">素早さ</option>
                  <option value="hp">HP</option>
                  <option value="mp">MP</option>
                </UI.Select>
              </UI.Td>
              <UI.Td>
                <UI.NumberInput onChange={e => setValue("howMuch", Number(e))}>
                  <UI.NumberInputField />
                </UI.NumberInput>
              </UI.Td>
              <UI.Td>
                <UI.Select {...register("happen")} >
                  <option value="+">アップ</option>
                  <option value="-">ダウン</option>
                  <option value="=">に</option>
                </UI.Select>
              </UI.Td>
              <UI.Td>
                <UI.Button type="submit">追加</UI.Button>
              </UI.Td>
            </UI.Tr>
          </UI.Tbody>
        </UI.Table>
      </form>
    </>
  )
}

export default Characteristic
