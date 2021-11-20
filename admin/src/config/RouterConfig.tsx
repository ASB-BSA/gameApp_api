import { Administrator, Characteristic, Charactor, Dashboard, Login, Setting, SettingsDetail } from "../components"
import { Layout } from "../components/template"
import * as Icon from "react-icons/ri";
import AuthProvider from "../utils/AuthProvider"

type RouterType = {
  name: string|null
  path: string|null
  component: JSX.Element
  icon: JSX.Element|null
  children: RouterType[]|null
}

const RouterConfig: RouterType[] = [
  {
    "name": "ログイン",
    "path": "/login",
    "component": <Login />,
    "icon": null,
    "children": null
  },
  {
    "name": null,
    "path": null,
    "component": <AuthProvider />,
    "icon": null,
    "children": [
      {
        "name": null,
        "path": null,
        "component": <Layout />,
        "icon": null,
        "children": [
          {
            "name": "ダッシュボード",
            "path": "/",
            "component": <Dashboard />,
            "icon": <Icon.RiDashboardFill />,
            "children": null
          },
          {
            "name": "ゲーム設定",
            "path": "/setting",
            "component": <Setting />,
            "icon": <Icon.RiListSettingsFill />,
            "children": null
          },
          {
            "name": "ゲーム設定 - 詳細",
            "path": "/setting/:id",
            "component": <SettingsDetail />,
            "icon": null,
            "children": null
          },
          {
            "name": "特技管理",
            "path": "/characteristic",
            "component": <Characteristic />,
            "icon": <Icon.RiFireFill />,
            "children": null
          },
          {
            "name": "キャラクター管理",
            "path": "/character",
            "component": <Charactor />,
            "icon": <Icon.RiBearSmileFill />,
            "children": null
          },
          {
            "name": "管理者",
            "path": "/administrator",
            "component": <Administrator />,
            "icon": <Icon.RiAdminFill />,
            "children": null
          },
        ]
      },
    ]
  }
]

export default RouterConfig