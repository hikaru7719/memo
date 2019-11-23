import { Icon, Input, Tag } from "antd";
import { TweenOneGroup } from "rc-tween-one";
import * as React from "react";

const colors = [
  "magenta",
  "red",
  "volcano",
  "orange",
  "gold",
  "lime",
  "green",
  "cyan",
  "blue",
  "geekblue",
  "purple"
];

export const CustomTag = () => {
  const [visable, setVisable] = React.useState(false);
  const [inputValue, setInputValue] = React.useState("");
  const [tags, setTags] = React.useState<string[]>([]);
  const inputRef = React.useRef<Input>(null);

  React.useEffect(() => {
    console.log(inputRef);
  }, [inputRef]);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value);
  };

  const addTags = () => {
    if (inputValue && !tags.find(t => t === inputValue)) {
      setTags([...tags, inputValue]);
    }
    setInputValue("");
    setVisable(false);
  };

  const handleClose = (removedTag: string) => {
    const t = tags.filter(tag => tag !== removedTag);
    setTags(t);
    setVisable(false);
  };

  const showInput = (e: React.MouseEvent) => {
    setVisable(true);
  };

  React.useEffect(() => {
    if (visable && inputRef && inputRef.current) {
      inputRef.current.focus();
    }
  }, [visable]);

  const forMap = (n: number, tag: string) => {
    const tagElem = (
      <Tag
        color={colors[n % 11]}
        closable
        onClose={(e: React.ChangeEvent<HTMLInputElement>) => {
          e.preventDefault();
          handleClose(tag);
        }}
      >
        {tag}
      </Tag>
    );
    return (
      <span key={n} style={{ display: "inline-block" }}>
        {tagElem}
      </span>
    );
  };

  const tagChild = tags.map((tag, n) => {
    return forMap(n, tag);
  });
  return (
    <div>
      <div style={{ marginBottom: 16 }}>
        <TweenOneGroup
          enter={{
            scale: 0.8,
            opacity: 0,
            type: "from",
            duration: 100
          }}
          leave={{ opacity: 0, width: 0, scale: 0, duration: 200 }}
          appear={false}
        >
          {tagChild}
        </TweenOneGroup>
      </div>
      {visable && (
        <Input
          ref={inputRef}
          type="text"
          size="small"
          style={{ width: 78 }}
          value={inputValue}
          onChange={handleInputChange}
          onBlur={addTags}
          onPressEnter={addTags}
        />
      )}
      {!visable && (
        <Tag
          onClick={showInput}
          style={{ background: "#fff", borderStyle: "dashed" }}
        >
          <Icon type="plus" /> New Tag
        </Tag>
      )}
    </div>
  );
};
