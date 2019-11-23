import { Card, Icon } from "antd";
import { Row, Col } from "antd";
import { Typography } from "antd";
import { Button } from "antd";
import * as React from "react";
export default () => {
  const Title = Typography.Title;
  return (
    <div style={{ marginTop: 200 }}>
      <Row type="flex" justify="center">
        <Col span={8}>
          <Card style={{ width: 500 }}>
            <div style={{ textAlign: "center" }}>
              <Title level={2}>アカウントを利用してログイン</Title>
              <Button type="primary" block>
                <Icon type="twitter" />
                Login with Twitter
              </Button>
              <Button type="danger" block style={{ marginTop: 15 }}>
                <Icon type="google" />
                Login with Google
              </Button>
              <Button
                block
                style={{
                  marginTop: 15,
                  backgroundColor: "#000",
                  color: "#fff"
                }}
              >
                <Icon type="github" />
                Login with GitHub
              </Button>
            </div>
          </Card>
        </Col>
      </Row>
    </div>
  );
};
