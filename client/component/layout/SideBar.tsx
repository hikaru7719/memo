import { Icon, Layout, Menu } from "antd";
import Link from "next/link";
export const SideBar = () => {
  const { Sider } = Layout;
  const { SubMenu } = Menu;
  return (
    <Sider trigger={null} style={{ height: 800 }}>
      <Menu
        theme="dark"
        mode="inline"
        defaultOpenKeys={["sub1", "sub2", "sub3"]}
      >
        <SubMenu
          key={"sub1"}
          title={
            <span>
              <Icon type="star" />
              <span>Goal</span>
            </span>
          }
        >
          <Menu.Item key="1">
            <Link href="/">
              <span>目標の一覧</span>
            </Link>
          </Menu.Item>
          <Menu.Item key="2">
            <Link href="/">
              <span>目標の登録</span>
            </Link>
          </Menu.Item>
        </SubMenu>
        <SubMenu
          key={"sub2"}
          title={
            <span>
              <Icon type="edit" />
              <span>Task</span>
            </span>
          }
        >
          <Menu.Item key="1">
            <Link href="/">
              <span>完了タスクの一覧</span>
            </Link>
          </Menu.Item>
          <Menu.Item key="2">
            <Link href="/">
              <span>タスクの登録</span>
            </Link>
          </Menu.Item>
        </SubMenu>
        <SubMenu
          key={"sub3"}
          title={
            <span>
              <Icon type="profile" />
              <span>Memo</span>
            </span>
          }
        >
          <Menu.Item key="1">
            <Link href="/">
              <span>メモの一覧</span>
            </Link>
          </Menu.Item>
          <Menu.Item key="2">
            <Link href="/">
              <span>メモの登録</span>
            </Link>
          </Menu.Item>
        </SubMenu>
      </Menu>
    </Sider>
  );
};
