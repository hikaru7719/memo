import { Form, Input, Select, Typography, Tag, Icon, Button } from "antd";
import axios from "axios";
import TextArea from "antd/lib/input/TextArea";
import { CustomTag } from "./CustomTag";
import * as React from "react";

export const Task = () => {
  const [loading, setLoading] = React.useState(false);
  const handleClick = async (e: React.MouseEvent) => {
    setLoading(true);
    const resp = await axios.post("http://gateway:8081/healthcheck");
    console.log(resp.data);
    setTimeout(() => {
      setLoading(false);
    }, 5000);
  };
  const { Option } = Select;
  return (
    <>
      <div style={{ padding: 50, background: "#fff", minHeight: 360 }}>
        <Typography>
          <Typography.Title>タスクの登録</Typography.Title>
        </Typography>
        <Form>
          <Form.Item label="タスク名">
            <Input placeholder="タスク1" />
          </Form.Item>
          <Form.Item label="説明">
            <TextArea
              rows={3}
              placeholder="今日中に終わらせないといけないタスクです。"
            />
          </Form.Item>
          <Form.Item label="ステータス">
            <Select defaultValue="TODO">
              <Option value="TODO">ToDo</Option>
              <Option value="PROGRESS">Progress</Option>
              <Option value="DONE">Done</Option>
            </Select>
          </Form.Item>
          <Form.Item label="タグ">
            <CustomTag />
          </Form.Item>
          <div style={{ textAlign: "center" }}>
            <Button type="primary" loading={loading} onClick={handleClick}>
              登録
            </Button>
          </div>
        </Form>
      </div>
    </>
  );
};
