import react from "react";
import "./button.css";

function Button({ children }) {
  return (
    <>
      <button className="Button">{children}</button>
    </>
  );
}

export default Button;
