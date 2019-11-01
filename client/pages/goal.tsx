import { Layout } from "antd";
import { Header } from "../component/layout/Header";
import { SideBar } from "../component/layout/SideBar";
export default () => {
  return (
    <Layout>
      <Header />
      <Layout>
        <SideBar />
      </Layout>
    </Layout>
  );
};
