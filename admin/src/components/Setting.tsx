import React, { useEffect } from 'react'
import * as UI from "@chakra-ui/react"
import { CreateGroup, SettingListItem } from './Settings'
import axios from 'axios'
import { useRecoilState, useSetRecoilState } from 'recoil'
import SettingState from '../atoms/SettingsState'
import IsRedirect from '../atoms/IsRedirect'

const Setting = () => {
  const [settings, setSettings] = useRecoilState(SettingState)
  const setRedirect = useSetRecoilState(IsRedirect)

  useEffect(() => {
    axios.get('settings')
      .then(e => setSettings(e.data))
      .catch(() => {
        setRedirect(true)
      })
  }, [])

  return (
    <>
      <UI.Heading
        size="md"
      >ゲーム設定</UI.Heading>

      <UI.Flex>
        <UI.Spacer />
        <CreateGroup />
      </UI.Flex>

      <UI.Flex
        justifyContent="center"
        alignItems="center"
        py={8}
      >
        <UI.Box
          px={8}
          py={20}
          mx="auto"
        >
          <UI.Box textAlign="center">
            <UI.Heading fontWeight="extrabold">各種設定</UI.Heading>
            <UI.Text
              mt={4}
              maxW="2xl"
              fontSize="xl"
              mx={{ lg: "auto" }}
              color={UI.useColorModeValue("gray.500", "gray.400")}
            >
              Boomin Fighters の各種設定ページです
            </UI.Text>
          </UI.Box>

          <UI.SimpleGrid
            columns={{ base: 1, sm: 2, md: 3, lg: 4 }}
            spacingX={{ base: 16, lg: 24 }}
            spacingY={20}
            mt={12}
          >
            {settings.map((e, i) => (
              <SettingListItem key={i} data={e} />
            ))}
          </UI.SimpleGrid>
        </UI.Box>
      </UI.Flex>
    </>
  )
}

export default Setting
