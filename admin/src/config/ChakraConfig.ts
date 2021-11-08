import { extendTheme, ThemeConfig } from "@chakra-ui/react"

const config :ThemeConfig = {
  useSystemColorMode: false,
  initialColorMode: "light",
}

const fonts = {
  heading: "Noto Sans JP",
  body: "Noto Sans JP"
}

const theme = extendTheme({ config, fonts })

export default theme