import * as Form from "@radix-ui/react-form";
import "./input.css";
import Button from "./button";

function Input({ children }) {
  return (
    <Form.Root className="FormRoot">
      <Form.Field className="FormField" name="input">
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "baseline",
            justifyContent: "space-between",
          }}
        >
          <Form.Message className="FormMessage" match="valueMissing">
            {children}
          </Form.Message>
        </div>
        <Form.Control asChild>
          <input className="Input" required />
        </Form.Control>
      </Form.Field>
      <Form.Submit asChild>
        <Button>✓</Button>
      </Form.Submit>
    </Form.Root>
  );
}

export default Input;
