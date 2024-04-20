import "./button.css";

function Button({ children, type='', onClick=null, styleFeature='' }) {
  return (
    <>
      <button type={type} className={`btn ${styleFeature}`} onClick={onClick}>{children}</button>
    </>
  );
}

export default Button;
