import React from 'react'
import * as UI from '@chakra-ui/react'
import { SettingGroup } from '../../atoms/SettingsState'
import { RiSettings3Fill } from 'react-icons/ri'
import { useNavigate } from 'react-router-dom'

type Props = {
  "data": SettingGroup
}

const SettingListItem = (props: Props) => {
  const navigate = useNavigate()
  
  return (
    <>
      <UI.LinkBox>
        <UI.Box>
          <UI.Flex
            alignItems="center"
            justifyContent="center"
            w={8}
            h={8}
            mb={4}
            rounded="full"
            color={UI.useColorModeValue(`gray.600`, `gray.100`)}
            bg={UI.useColorModeValue(`gray.100`, `gray.600`)}
          >
            <RiSettings3Fill />
          </UI.Flex>
          <UI.LinkOverlay onClick={() => {
            navigate(`/setting/${props.data.group_category}`)
          }}>
            <UI.Heading
              mb={2}
              size="md"
              lineHeight="shorter"
              color={UI.useColorModeValue("gray.900", "white.900")}
            >
              {props.data.group_name}
            </UI.Heading>
          </UI.LinkOverlay>
          <UI.Text
            fontSize="sm"
            color={UI.useColorModeValue("gray.500", "gray.300")}
          >
            {props.data.group_name}の設定画面
          </UI.Text>
        </UI.Box>
      </UI.LinkBox>
    </>
  )
}

export default SettingListItem
