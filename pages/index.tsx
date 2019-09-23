import React, { useState } from "react";
import {
  TextField,
  Stack,
  IStackProps,
  PrimaryButton
} from "office-ui-fabric-react";

const Example: React.FC = () => {
  const columnProps: Partial<IStackProps> = {
    tokens: { childrenGap: 25 },
    styles: { root: { width: 800 } }
  };
  const columnProps2: Partial<IStackProps> = {
    styles: { root: { width: 100 } }
  };
  const [count, setCount] = useState(0);
  return (
    <Stack horizontal horizontalAlign="center" tokens={{ childrenGap: 50 }}>
      <Stack {...columnProps}>
        <TextField label="title" />
        <TextField label="memo" multiline rows={3} />
        <Stack.Item align="center" styles={columnProps2}>
          <PrimaryButton text="Primary" />
        </Stack.Item>
      </Stack>
    </Stack>
  );
};

export default Example;
