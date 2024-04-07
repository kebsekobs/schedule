import React, { useState } from "react";
import * as Select from "@radix-ui/react-select";
import {
  CheckIcon,
  ChevronDownIcon,
  ChevronUpIcon,
} from "@radix-ui/react-icons";
import "./select.css";

function SelectComp({ list, placeholder, name }) {
  const [isToggled, setIsToggled] = useState("closed");

  return (
    <div>
      <Select.Root
        onOpenChange={(e) => setIsToggled(e === true ? "open" : "closed")}
      >
        <Select.Trigger
          className="SelectTrigger"
          aria-label="Food"
          data-state={isToggled}
        >
          <Select.Value placeholder={placeholder} />
          <Select.Icon className="SelectIcon">
            <ChevronDownIcon />
          </Select.Icon>
        </Select.Trigger>
        <Select.Content className="SelectContent">
          <Select.ScrollUpButton className="SelectScrollButton">
            <ChevronUpIcon />
          </Select.ScrollUpButton>
          <Select.Viewport className="SelectViewport">
            <Select.Group>
              <Select.Label className="SelectLabel">{name}</Select.Label>
              {list.map((item, id) => (
                <SelectItem value={item.id} id={id}>
                  {item.id}
                </SelectItem>
              ))}
            </Select.Group>
          </Select.Viewport>
          <Select.ScrollDownButton className="SelectScrollButton">
            <ChevronDownIcon />
          </Select.ScrollDownButton>
        </Select.Content>
      </Select.Root>
    </div>
  );
}

const SelectItem = React.forwardRef(
  ({ children, className, id, ...props }, forwardedRef) => {
    return (
      <Select.Item className="" {...props} ref={forwardedRef} key={id}>
        <Select.ItemText>{children}</Select.ItemText>
        <Select.ItemIndicator className="SelectItemIndicator">
          <CheckIcon />
        </Select.ItemIndicator>
      </Select.Item>
    );
  }
);

export default SelectComp;
