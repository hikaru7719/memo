import express from "express";
import next from "next";
import * as library from "office-ui-fabric-react/lib-commonjs/Utilities";
import * as responsiveLib from "office-ui-fabric-react/lib-commonjs/utilities/decorators/withResponsiveMode";

library.setSSR(true);
library.setRTL(true);
responsiveLib.setResponsiveMode(responsiveLib.ResponsiveMode.large);

const port = 3000;
const dev = process.env.NODE_ENV !== "production";
const renderer = next({ dev });
const handle = renderer.getRequestHandler();
const app = express();
app.get("*", async (req: express.Request, res: express.Response) => {
  return handle(req, res);
});

renderer.prepare().then(() => {
  app.listen(port, (err?: Error) => {
    if (err) {
      throw err;
    }
    console.log("app start");
  });
});
