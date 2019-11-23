import { Layout } from "antd";
import { Header } from "../component/layout/Header";
import { SideBar } from "../component/layout/SideBar";
import { Task } from "../component/Task";
import { Top } from "../component/Top";
export default () => {
  const { Content } = Layout;
  return (
    <Layout>
      <Header />
      <Layout>
        <SideBar />
        <Content
          style={{
            margin: "24px 16px",
            padding: 24,
            background: "#fff",
            minHeight: 280
          }}
        >
          <Task />
        </Content>
      </Layout>
    </Layout>
  );
};
