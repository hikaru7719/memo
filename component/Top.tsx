import { Row, Card, Col } from "antd";
export const Top = () => {
  return (
    <>
      <div style={{ width: "100%", height: "250px", background: "#1890ff" }}>
        <div style={{ textAlign: "center", color: "#fff" }}>
          <p style={{ marginTop: "30px", fontSize: "130px" }}>Your Note ...</p>
        </div>
      </div>
      <div style={{ height: "250px", marginTop: "30px" }}>
        <Row gutter={16}>
          <Col key={1} span="8">
            <Card>Hoge</Card>
          </Col>
          <Col key={2} span="8">
            <Card>Hoge</Card>
          </Col>
          <Col key={3} span="8">
            <Card>Hoge</Card>
          </Col>
        </Row>
      </div>
    </>
  );
};
