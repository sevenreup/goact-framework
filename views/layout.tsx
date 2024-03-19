import "./main.css"

const Layout = ({ children }) => {
  return (
    <div className="p-2">
      <h1>My App</h1>
      {children}
    </div>
  );
};

export default Layout;