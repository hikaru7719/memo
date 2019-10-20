import { Row, Card, Col, Icon, Typography } from "antd";
export const Top = () => {
  const { Title, Text } = Typography;
  return (
    <>
      <div style={{ width: "100%", backgroundColor: "#fff" }}>
        <div style={{ textAlign: "center" }}>
          <Title
            style={{
              marginTop: "10px",
              marginBottom: "30px",
              fontSize: "80px"
            }}
          >
            Memo ...
          </Title>
        </div>
      </div>
      <div
        style={{
          width: "100%",
          height: "300px",
          backgroundImage: `url("/static/farm_and_cow.jpg")`,
          marginTop: "-20px"
        }}
      >
        <div style={{ textAlign: "center" }}>
          <Text>
            Memoは個人の目標やタスクを記録するアプリです。個人のポートフォリをとして活用できます。
            長期的に目標やタスクをトレースしてくれます。
          </Text>
        </div>
      </div>
      <div style={{ height: "400px", marginTop: "30px" }}>
        <Row type="flex" gutter={16} justify="space-around" align="top">
          <Col key={1} span={6}>
            <Card hoverable>
              <Typography>
                <Title style={{ fontSize: "30px" }}>
                  <Icon type="star" style={{ margin: "2px" }} />
                  Goal
                </Title>
                <Text>個人の月間目標や年間目標を記録します。</Text>
              </Typography>
            </Card>
          </Col>
          <Col key={2} span={6}>
            <Card hoverable>
              <Typography>
                <Title style={{ fontSize: "30px" }}>
                  <Icon type="edit" style={{ margin: "2px" }} />
                  Finishd Task
                </Title>
                <Text>完了したタスクを記録します。</Text>
              </Typography>
            </Card>
          </Col>
          <Col key={3} span={6}>
            <Card hoverable>
              <Typography>
                <Title style={{ fontSize: "30px" }}>
                  <Icon type="profile" style={{ margin: "2px" }} />
                  Note
                </Title>
                <Text>日々の気づきのメモを記録します。</Text>
              </Typography>
            </Card>
          </Col>
        </Row>
      </div>
    </>
  );
};
