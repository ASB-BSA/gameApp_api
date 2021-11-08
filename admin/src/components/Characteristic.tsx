import React from 'react'
import * as UI from "@chakra-ui/react"
import { useRecoilState } from 'recoil';
import CharacteristicState from '../atoms/CharacteristicState';

const Characteristic = () => {
  const [characteristics, setCharacteristics] = useRecoilState(CharacteristicState);

  return (
    <>
      <UI.Heading
        size="md"
      >特技管理</UI.Heading>
      <UI.Table variant="striped" colorScheme="gray">
        <UI.Thead>
          <UI.Tr>
            <UI.Th>技名</UI.Th>
            <UI.Th>into</UI.Th>
            <UI.Th>multiply by</UI.Th>
          </UI.Tr>
        </UI.Thead>
        <UI.Tbody>
          {
            characteristics.map(e => {
              return (
                <UI.Tr>
                  <UI.Td>inches</UI.Td>
                  <UI.Td>millimeUI.tres (mm)</UI.Td>
                  <UI.Td>25.4</UI.Td>
                </UI.Tr>
              )
            })
          }
          <UI.Tr>
            <UI.Td>inches</UI.Td>
            <UI.Td>millimeUI.tres (mm)</UI.Td>
            <UI.Td>25.4</UI.Td>
          </UI.Tr>
        </UI.Tbody>
      </UI.Table>
    </>
  )
}

export default Characteristic
