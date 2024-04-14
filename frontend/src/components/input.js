import * as Form from "@radix-ui/react-form";
import styles from "./input.module.css";
import Button from "./button";

function Input({ children }) {
  return (
    <Form.Root className={styles["FormRoot"]}>
      <Form.Field className={styles["FormField"]} name="input">
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "baseline",
            justifyContent: "space-between",
          }}
        >
          <Form.Message className={styles["FormMessage"]} match="valueMissing">
            {children}
          </Form.Message>
        </div>
        <Form.Control asChild>
          <input className={styles["Input"]} required />
        </Form.Control>
      </Form.Field>
      <Form.Submit asChild>
        <Button>✓</Button>
      </Form.Submit>
    </Form.Root>
  );
}

export default Input;
